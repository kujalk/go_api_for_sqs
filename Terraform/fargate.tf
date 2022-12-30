resource "aws_ecs_cluster" "fargate" {
  name = "${var.project_name}_fargate_cluster"
}

resource "aws_ecs_cluster_capacity_providers" "fargate" {
  cluster_name = aws_ecs_cluster.fargate.name

  capacity_providers = ["FARGATE"]

  default_capacity_provider_strategy {
    base              = 1
    weight            = 100
    capacity_provider = "FARGATE"
  }
}

resource "aws_ecs_task_definition" "task" {
  family                   = "${var.project_name}_ecs_task"
  requires_compatibilities = ["FARGATE"]
  network_mode             = "awsvpc"
  execution_role_arn       = aws_iam_role.ecs_role.arn
  task_role_arn            = aws_iam_role.ecs_role.arn
  cpu                      = 256
  memory                   = 512

  container_definitions = <<EOF
[
  {
    "name": "go-api",
    "image": "${var.image_url}",
    "cpu": 256,
    "memory": 512,
    "secrets": [{
      "name": "SQS_QUEUE_URL",
      "valueFrom": "${aws_ssm_parameter.sqs-queue.arn}"
    }]
  }
]
EOF
}


resource "aws_ecs_service" "go-api" {
  name            = "${var.project_name}_ecs_service"
  cluster         = aws_ecs_cluster.fargate.id
  task_definition = aws_ecs_task_definition.task.arn
  desired_count   = 1

  network_configuration {
    subnets          = [var.public_subnet1_id, var.public_subnet2_id]
    assign_public_ip = true
    security_groups  = [aws_security_group.ecs-service.id]
  }
}