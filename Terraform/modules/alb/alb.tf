resource "aws_alb" "alb" {
  name            = "${var.cluster_name}-alb"
  security_groups = ["${aws_security_group.alb.id}"]
  subnets         = var.subnet_ids
  tags = {
    Name = "${var.cluster_name}-alb"
  }
}

resource "aws_alb_target_group" "default" {
  name                 = "${aws_alb.alb.name}-default"
  port                 = 80
  protocol             = "HTTP"
  vpc_id               = var.vpc_id

  tags = {
    Environment = var.environment
  }
}

resource "aws_alb_listener" "http" {
  load_balancer_arn = aws_alb.alb.id
  port              = "80"
  protocol          = "HTTP"

  default_action {
    target_group_arn = aws_alb_target_group.default.id
    type             = "forward"
  }
}