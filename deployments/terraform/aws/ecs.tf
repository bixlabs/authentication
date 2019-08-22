resource "aws_ecs_cluster" "main" {
  name = "${var.aws_ecs_cluster_name}"
}

data "template_file" "go_authenticator_api" {
  template = file("${path.module}/templates/ecs/go_authenticator_api.json.tpl")

  vars = {
    app_image                    = "${var.app_image}"
    app_port                     = "${var.auth_server_port}"
    app_server_token_expiration  = "${var.auth_server_token_expiration}"
    app_server_secret            = "${var.auth_server_secret}"
    app_reset_password_max       = "${var.auth_server_reset_password_max}"
    app_reset_password_min       = "${var.auth_server_reset_password_min}"
    app_server_database_name     = "${var.auth_server_database_name}"
    app_server_database_user     = "${var.auth_server_database_user}"
    app_server_database_password = "${var.auth_server_database_password}"
    app_server_database_salt     = "${var.auth_server_database_salt}"
    fargate_cpu                  = "${var.fargate_cpu}"
    fargate_memory               = "${var.fargate_memory}"
    aws_region                   = "${var.aws_region}"
  }
}

resource "aws_ecs_task_definition" "app" {
  family                   = "ga-task"
  execution_role_arn       = "${aws_iam_role.ecs_task_execution_role.arn}"
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  cpu                      = "${var.fargate_cpu}"
  memory                   = "${var.fargate_memory}"
  container_definitions    = "${data.template_file.go_authenticator_api.rendered}"
}

resource "aws_ecs_service" "main" {
  name            = "ga-service"
  cluster         = "${aws_ecs_cluster.main.id}"
  task_definition = "${aws_ecs_task_definition.app.arn}"
  desired_count   = "${var.app_count}"
  launch_type     = "FARGATE"

  network_configuration {
    security_groups  = ["${aws_security_group.ecs_tasks.id}"]
    subnets          = "${aws_subnet.private.*.id}"
    assign_public_ip = true
  }

  load_balancer {
    target_group_arn = "${aws_alb_target_group.app.id}"
    container_name   = "go-authenticator"
    container_port   = "${var.auth_server_port}"
  }

  depends_on = [aws_alb_listener.front_end, aws_iam_role_policy_attachment.ecs_task_execution_role]
}