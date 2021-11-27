locals {
  database_username = "api"
}

resource "random_password" "database_password" {
  length           = 64
  special          = true
  override_special = "!#$%&*()-_=+[]{}<>:?"
}

resource "aws_rds_cluster" "api_database" {
  cluster_identifier              = "dev-db"
  engine                          = "aurora-postgresql"
  database_name                   = "prixfixe"
  enabled_cloudwatch_logs_exports = ["postgresql"]

  engine_mode = "serverless"
  scaling_configuration {
    auto_pause               = true
    min_capacity             = 2
    max_capacity             = 2
    seconds_until_auto_pause = 300
    timeout_action           = "ForceApplyCapacityChange"
  }

  master_username         = local.database_username
  master_password         = random_password.database_password.result
  backup_retention_period = 7
  storage_encrypted       = true
  preferred_backup_window = "01:00-05:00"

  tags = merge(var.default_tags, {})
}

resource "aws_ssm_parameter" "database_url" {
  name  = "PRIXFIXE_DATABASE_URL"
  type  = "String"
  value = format("postgres://%s:%s@%s:%d/prixfixe", local.database_username, random_password.database_password.result, aws_rds_cluster.api_database.endpoint, aws_rds_cluster.api_database.port)

  tags = merge(var.default_tags, {})
}

provider "postgresql" {
  host            = aws_rds_cluster.api_database.endpoint
  port            = aws_rds_cluster.api_database.port
  username        = local.database_username
  password        = random_password.database_password.result
  sslmode         = "require"
  connect_timeout = 15
}

resource "postgresql_database" "prixfixe" {
  name = "prixfixe"
}