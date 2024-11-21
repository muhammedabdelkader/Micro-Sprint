provider "aws" {
  region = "eu-west-1"
}

# Variable to select a VPC by name
variable "selected_vpc_name" {
  description = "The name of the VPC to use for further operations"
  type        = string 
}

# Fetch the VPC by Name
data "aws_vpc" "selected_vpc" {
  filter {
    name   = "tag:Name"
    values = [var.selected_vpc_name]
  }
}

# Fetch all subnets in the VPC
data "aws_subnets" "vpc_subnets" {
  filter {
    name   = "vpc-id"
    values = [data.aws_vpc.selected_vpc.id]
  }
}

# Fetch details for each subnet
data "aws_subnet" "details" {
  for_each = toset(data.aws_subnets.vpc_subnets.ids)
  id       = each.value
}

# Output all allocated CIDR blocks
output "allocated_cidr_blocks" {
  value = [
    for subnet in data.aws_subnet.details :
    subnet.cidr_block
  ]
}

# Available CIDR block calculation (placeholder for custom logic)
output "available_cidr_ranges" {
  value = "Custom logic required: Subtract allocated CIDRs from VPC CIDR"
}

output "vpc_cidr" {
  value = data.aws_vpc.selected_vpc.cidr_block
  
}