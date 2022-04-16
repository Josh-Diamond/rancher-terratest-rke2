# Rancher specific variable section.

variable "rancher_api_url" {
  type        = string
  description = "The URL of your Rancher server that is being tested."
}

variable "rancher_admin_bearer_token" {
  type        = string
  description = "The bearer token for the Rancher admin user."
}

# AWS specific variables.

variable "aws_region" {
  type        = string
  description = "The region to use."
}

variable "aws_security_group_name" {
  type        = string
  description = "The security group name to be used."
}

variable "aws_subnet_id" {
  type        = string
  description = "The subnet ID to use."
}

variable "aws_vpc_id" {
  type        = string
  description = "The VPC to use."
}

variable "aws_zone_letter" {
  type        = string
  description = "The zone letter to use."
}
