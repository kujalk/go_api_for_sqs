resource "aws_security_group" "ecs-service" {
  name        = "${var.project_name}_ECS_Service"
  description = "To allow Traffic to ECS Service"
  vpc_id      = var.vpc_id


  tags = {
    Name = "${var.project_name}_ECS_Service"
  }

  ingress {
    description = "Traffic Allow for API"
    from_port   = 8080
    to_port     = 8080
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    description = "Outside"
    from_port   = 0
    to_port     = 0
    protocol    = -1
    cidr_blocks = ["0.0.0.0/0"]
  }
}
