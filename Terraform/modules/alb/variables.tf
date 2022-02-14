variable "allowed_cidr_blocks" {
  type = list(string)
}

variable "cluster_name" {
  description = "The name to use to create the cluster and the resources. Only alphanumeric characters and dash allowed (e.g. 'my-cluster')"
}
variable "aws_region" {
  default     = "us-east-1"
  description = "The AWS region where to launch the cluster and releated resources"
}

variable "vpc_id" {
  default     = "invalid-vpc"
  type        = string
  description = "A VPC for loadbalancer"
}

variable "subnet_ids" {
  description = "A list of subnet IDs for LoadBalancer"
  type = list(string)
}

variable "environment" {}