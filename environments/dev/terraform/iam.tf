resource "aws_iam_role" "worker_lambda_role" {
  name = "Worker"

  inline_policy {
    name = "allow_sqs_queue_access"

    policy = jsonencode({
      Version = "2012-10-17",
      Statement = [
        {
          Action = [
            "sqs:SendMessage",
            "sqs:ReceiveMessage",
          ],
          Effect   = "Allow",
          Resource = "*",
        }
      ]
    })
  }

  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Effect = "Allow",
        Principal = {
          Service = "lambda.amazonaws.com",
        },
        Action = "sts:AssumeRole",
      },
    ]
  })
}

resource "aws_iam_role" "server_lambda_role" {
  name = "APIServer"

  inline_policy {
    name = "allow_sqs_queue_access"

    policy = jsonencode({
      Version = "2012-10-17",
      Statement = [
        {
          Action = [
            "sqs:SendMessage",
            "sqs:ReceiveMessage",
          ],
          Effect   = "Allow",
          Resource = "*",
        }
      ]
    })
  }

  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Effect = "Allow",
        Principal = {
          Service = "ec2.amazonaws.com",
        },
        Action = "sts:AssumeRole",
      },
      {
        Effect = "Allow",
        Principal = {
          Service = "lambda.amazonaws.com",
        },
        Action = "sts:AssumeRole",
      }
    ]
  })
}