resource "digitalocean_kubernetes_cluster" "dev" {
  name         = "dev"
  region       = local.region
  auto_upgrade = true
  version      = "1.21.9-do.0"

  maintenance_policy {
    start_time = "04:00"
    day        = "sunday"
  }

  node_pool {
    name       = "default"
    size       = "s-1vcpu-1gb"
    node_count = 3
  }
}