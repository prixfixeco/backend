terraform {
  required_version = "1.9.5"

  cloud {
    organization = "dinnerdonebetter"

    workspaces {
      name = "dev-admin-app"
    }
  }

  required_providers {
    cloudflare = {
      source  = "cloudflare/cloudflare"
      version = "4.40.0"
    }
    google = {
      source  = "hashicorp/google"
      version = "5.41.0"
    }
    archive = {
      source  = "hashicorp/archive"
      version = "2.5.0"
    }
    kubernetes = {
      source  = "hashicorp/kubernetes"
      version = "2.35.1"
    }
  }
}