resource "digitalocean_database_cluster" "database" {
  name       = "prixfixe-dev"
  engine     = "pg"
  version    = "13"
  size       = "db-s-1vcpu-1gb"
  region     = local.region
  node_count = 1
}

resource "digitalocean_project_resources" "dev_db" {
  project = digitalocean_project.prixfixe_dev.id
  resources = [
    format("do:dbaas:%s", digitalocean_database_cluster.database.id),
  ]
}

resource "cloudflare_record" "database_dot_prixfixe_dot_dev" {
  zone_id         = var.CLOUDFLARE_ZONE_ID
  name            = "database.prixfixe.dev"
  value           = digitalocean_database_cluster.database.host
  type            = "CNAME"
  proxied         = true
  allow_overwrite = true
  ttl             = 1
}
