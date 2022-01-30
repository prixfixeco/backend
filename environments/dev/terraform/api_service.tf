resource "digitalocean_app" "prixfixe" {
  spec {
    name   = "dev"
    region = "nyc3"

    domain {
      name = "api.prixfixe.dev"
    }

    service {
      name               = "api"
      environment_slug   = "go"
      instance_count     = 1
      instance_size_slug = "professional-xs"

      image {
        registry_type = "DOCR"
        registry      = ""
        repository    = "api_server"
      }

      http_port = 8000

      routes {
        path = "/"
      }

      run_command = "/server"
    }

    database {
      name       = "dev"
      engine     = "PG"
      production = false
      version    = "12"
    }
  }

  depends_on = [
    digitalocean_container_registry.dev,
  ]
}