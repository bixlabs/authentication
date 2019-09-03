[
  {
    "name": "go-authenticator",
    "image": "${app_image}",
    "cpu": ${fargate_cpu},
    "memory": ${fargate_memory},
    "networkMode": "awsvpc",
    "logConfiguration": {
        "logDriver": "awslogs",
        "options": {
          "awslogs-group": "/ecs/ga",
          "awslogs-region": "${aws_region}",
          "awslogs-stream-prefix": "ecs"
        }
    },
    "environment": [
      {
        "name": "AUTH_SERVER_PORT",
        "value": "${app_port}"
      },
      {
        "name": "AUTH_SERVER_TOKEN_EXPIRATION",
        "value": "${app_server_token_expiration}"
      },
      {
        "name": "AUTH_SERVER_SECRET",
        "value": "${app_server_secret}"
      },
      {
        "name": "AUTH_SERVER_RESET_PASSWORD_MAX",
        "value": "${app_reset_password_max}"
      },
      {
        "name": "AUTH_SERVER_RESET_PASSWORD_MIN",
        "value": "${app_reset_password_min}"
      },
      {
        "name": "AUTH_SERVER_DATABASE_NAME",
        "value": "${app_server_database_name}"
      },
      {
        "name": "AUTH_SERVER_DATABASE_USER",
        "value": "${app_server_database_user}"
      },
      {
        "name": "AUTH_SERVER_DATABASE_PASSWORD",
        "value": "${app_server_database_password}"
      },
      {
        "name": "AUTH_SERVER_DATABASE_SALT",
        "value": "${app_server_database_salt}"
      }
    ],
    "portMappings": [
      {
        "containerPort": ${app_port},
        "hostPort": ${app_port}
      }
    ]
  }
]