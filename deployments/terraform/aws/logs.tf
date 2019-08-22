# Set up CloudWatch group and log stream and retain logs for 30 days
resource "aws_cloudwatch_log_group" "go_authenticator_api_log_group" {
  name              = "/ecs/ga"
  retention_in_days = 30

  tags = {
    Name = "ga-log-group"
  }
}

resource "aws_cloudwatch_log_stream" "go_authenticator_api_log_log_stream" {
  name           = "ga-log-stream"
  log_group_name = "${aws_cloudwatch_log_group.go_authenticator_api_log_group.name}"
}