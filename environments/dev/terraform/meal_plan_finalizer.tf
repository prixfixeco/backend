resource "aws_ecr_repository" "meal_plan_finalizer" {
  name = "meal_plan_finalizer"
  # do not set image_tag_mutability to "IMMUTABLE", or else we cannot use :latest tags.

  image_scanning_configuration {
    scan_on_push = true
  }
}

resource "aws_cloudwatch_log_group" "meal_plan_finalizer" {
  name              = "/ecs/dev_meal_plan_finalizer"
  retention_in_days = local.log_retention_period_in_days
}

resource "aws_cloudwatch_log_group" "finalizer_sidecar" {
  name              = "/ecs/dev-meal-plan-finalizer-telemetry-collector-sidecar"
  retention_in_days = local.log_retention_period_in_days
}


resource "aws_iam_role" "meal_plan_finalizer_task_execution_role" {
  name               = "meal-plan-finalizer-task-execution-role"
  assume_role_policy = data.aws_iam_policy_document.ecs_task_execution_assume_role.json
}

resource "aws_iam_role" "meal_plan_finalizer_task_role" {
  name = "meal-plan-finalizer-task-role"

  assume_role_policy = data.aws_iam_policy_document.ecs_task_assume_role.json

  inline_policy {
    name   = "allow_sqs_queue_access"
    policy = data.aws_iam_policy_document.allow_to_manipulate_queues.json
  }

  inline_policy {
    name   = "allow_ssm_access"
    policy = data.aws_iam_policy_document.allow_parameter_store_access.json
  }

  inline_policy {
    name   = "allow_decrypt_ssm_parameters"
    policy = data.aws_iam_policy_document.allow_to_decrypt_parameters.json
  }

  inline_policy {
    name   = "ECS-AWSOTel"
    policy = data.aws_iam_policy_document.opentelemetry_collector_policy.json
  }
}

resource "aws_ecs_task_definition" "meal_plan_finalizer" {
  family = "meal_plan_finalizer"

  container_definitions = jsonencode([
    {
      name : "otel-collector",
      image : format("%s:latest", aws_ecr_repository.otel_collector.repository_url)
      essential : true,
      logConfiguration : {
        logDriver : "awslogs",
        options : {
          awslogs-region : local.aws_region,
          awslogs-group : aws_cloudwatch_log_group.finalizer_sidecar.name,
          awslogs-create-group : "true",
          awslogs-stream-prefix : "otel-collector"
        }
      }
    },
    {
      name  = "meal_plan_finalizer",
      image = format("%s:latest", aws_ecr_repository.meal_plan_finalizer.repository_url),
      essential : true,
      logConfiguration : {
        logDriver : "awslogs",
        options : {
          awslogs-region : local.aws_region,
          awslogs-group : aws_cloudwatch_log_group.meal_plan_finalizer.name,
          awslogs-stream-prefix : "ecs",
        },
      },
    },
  ])

  execution_role_arn = aws_iam_role.meal_plan_finalizer_task_execution_role.arn
  task_role_arn      = aws_iam_role.meal_plan_finalizer_task_role.arn

  # These are the minimum values for Fargate containers.
  cpu                      = 256
  memory                   = 512
  requires_compatibilities = ["FARGATE"]

  network_mode = "awsvpc"
}

resource "aws_ecs_service" "meal_plan_finalizer" {
  name                               = "meal_plan_finalizer"
  task_definition                    = aws_ecs_task_definition.meal_plan_finalizer.arn
  cluster                            = aws_ecs_cluster.dev.id
  launch_type                        = "FARGATE"
  deployment_maximum_percent         = 200
  deployment_minimum_healthy_percent = 100
  desired_count                      = 1

  deployment_controller {
    type = "ECS"
  }

  deployment_circuit_breaker {
    enable   = true
    rollback = true
  }

  network_configuration {
    security_groups = [
      aws_security_group.meal_plan_finalizer.id,
    ]

    subnets = concat(
      [for x in aws_subnet.public_subnets : x.id],
      [for x in aws_subnet.private_subnets : x.id],
    )
  }
}

resource "aws_security_group" "meal_plan_finalizer" {
  name        = "meal_plan_finalizer"
  description = "meal plan finalizer"
  vpc_id      = aws_vpc.main.id

  egress {
    from_port        = 0
    to_port          = 0
    protocol         = "-1"
    cidr_blocks      = ["0.0.0.0/0"]
    ipv6_cidr_blocks = ["::/0"]
  }
}
