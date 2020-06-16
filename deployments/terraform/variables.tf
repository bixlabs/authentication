variable "aws_region" {
  description = "The AWS region things are created in"
  default     = "us-west-2"
}

variable "ecs_task_execution_role_name" {
  description = "ECS task execution role name"
  default = "myEcsTaskExecutionRole"
}

variable "ecs_auto_scale_role_name" {
  description = "ECS auto scale role Name"
  default = "myEcsAutoScaleRole"
}

variable "az_count" {
  description = "Number of AZs to cover in a given region"
  default     = "2"
}

variable "app_image" {
  description = "Docker image to run in the ECS cluster"
  default     = "bixlabs/go-authenticator:0.1"
}

variable "auth_server_port" { // Can be override by TF_VAR_auth_server_port ENV var
  description = "Port exposed by the docker image to redirect traffic to"
  default     = 9000
}

variable "auth_server_token_expiration" {
  description = "Token expiration time"
  default     = 3600
}

variable "auth_server_secret" {
  description = "Server secret"
  default     = ""
}

variable "auth_server_reset_password_max" {
  description = "Reset password max"
  default     = 99999
}

variable "auth_server_reset_password_min" {
  description = "Reset password min"
  default     = 10000
}


variable "auth_server_database_name" {
  description = "Database name"
  default     = "sqlite.s3db"
}


variable "auth_server_database_user" {
  description = "Database user"
  default     = "admin"
}


variable "auth_server_database_password" {
  description = "Database password"
  default     = "admin"
}


variable "auth_server_database_salt" {
  description = "Database salt"
  default     = "salted"
}

variable "app_count" {
  description = "Number of docker containers to run"
  default     = 1
}

// TODO: Create a healtcheck endpoint and update health_check_path variable
variable "health_check_path" {
  default = "/healthcheck"
}

variable "fargate_cpu" {
  description = "Fargate instance CPU units to provision (1 vCPU = 1024 CPU units)"
  default     = "1024"
}

variable "fargate_memory" {
  description = "Fargate instance memory to provision (in MiB)"
  default     = "2048"
}