terraform {
  backend "remote" {
    organization = "prixfixe"

    workspaces {
      name = "dev-API"
    }
  }
  required_providers {
    cloudflare = {
      source  = "cloudflare/cloudflare"
      version = "~> 2.0"
    }
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.0"
    }
  }
}

variable "default_tags" {
  default = {
    Environment = "dev"
    Terraform   = "true"
  }
  description = "default resource tags"
  type        = map(string)
}