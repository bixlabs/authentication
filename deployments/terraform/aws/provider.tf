provider "aws" {
  shared_credentials_file = "$HOME/.aws/credentials"
  profile    = "default"
  region     =  "${var.aws_region}"
  version = "~> 2.24"
}

provider "template" {
  version = "~> 2.1"
}