terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.0"
    }
  }
}

provider "aws" {
  region  = var.aws_region
  access_key = "${var.aws_access_key}"
  secret_key = "${var.aws_secret_key}"
  version = "~> 3.0"
}

terraform {
  backend "s3" {
    bucket = "terraform--ecs"
    key    = "sample/terraform.tfstate"
    region = "ap-south-1"
  }
}