resource "aws_ssm_parameter" "sqs-queue" {
  name  = "${var.project_name}_sqs_url"
  type  = "String"
  value = aws_sqs_queue.app_queue.url
}

resource "aws_secretsmanager_secret" "api-token" {
  name = var.aws_secret_name
}

resource "aws_secretsmanager_secret_version" "api-token" {
  secret_id     = aws_secretsmanager_secret.api-token.id
  secret_string = var.api_token
}

resource "aws_sqs_queue" "app_queue" {
  name                      = "${var.project_name}_queue"
  max_message_size          = 2048
  message_retention_seconds = 3600
  receive_wait_time_seconds = 5
}