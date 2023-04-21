resource "google_pubsub_topic" "outbound_emails_topic" {
  name = "outbound_emails"
}

resource "google_project_iam_custom_role" "outbound_emailer_role" {
  role_id     = "outbound_emailer_role"
  title       = "Outbound emailer role"
  description = "An IAM role for the outbound emailer"
  permissions = [
    "secretmanager.versions.access",
    "pubsub.topics.list",
    "pubsub.subscriptions.consume",
    "pubsub.subscriptions.create",
    "pubsub.subscriptions.delete",
  ]
}

resource "google_storage_bucket" "outbound_emailer_bucket" {
  name     = "outbound-emailer-cloud-function"
  location = "US"
}

data "archive_file" "outbound_emailer_function" {
  type        = "zip"
  source_dir  = "${path.module}/outbound_emailer_cloud_function"
  output_path = "${path.module}/outbound_emailer_cloud_function.zip"
}

resource "google_storage_bucket_object" "outbound_emailer_archive" {
  name   = format("outbound_emailer_function-%s.zip", data.archive_file.outbound_emailer_function.output_md5)
  bucket = google_storage_bucket.outbound_emailer_bucket.name
  source = "${path.module}/outbound_emailer_cloud_function.zip"
}

resource "google_service_account" "outbound_emailer_user_service_account" {
  account_id   = "outbound-emailer-worker"
  display_name = "Outbound Emailer Worker"
}

# Permissions on the service account used by the function and Eventarc trigger
resource "google_project_iam_member" "outbound_emailer_invoking" {
  project = local.project_id
  role    = "roles/run.invoker"
  member  = "serviceAccount:${google_service_account.outbound_emailer_user_service_account.email}"
}

resource "google_project_iam_member" "outbound_emailer_secret_accessor" {
  project    = local.project_id
  role       = "roles/secretmanager.secretAccessor"
  member     = "serviceAccount:${google_service_account.outbound_emailer_user_service_account.email}"
  depends_on = [google_project_iam_member.outbound_emailer_invoking]
}

resource "google_project_iam_member" "outbound_emailer_event_receiving" {
  project    = local.project_id
  role       = "roles/eventarc.eventReceiver"
  member     = "serviceAccount:${google_service_account.outbound_emailer_user_service_account.email}"
  depends_on = [google_project_iam_member.outbound_emailer_secret_accessor]
}

resource "google_project_iam_member" "outbound_emailer_artifactregistry_reader" {
  project    = local.project_id
  role       = "roles/artifactregistry.reader"
  member     = "serviceAccount:${google_service_account.outbound_emailer_user_service_account.email}"
  depends_on = [google_project_iam_member.outbound_emailer_event_receiving]
}

resource "google_project_iam_member" "outbound_emailer_user" {
  project = local.project_id
  role    = google_project_iam_custom_role.outbound_emailer_role.id
  member  = format("serviceAccount:%s", google_service_account.outbound_emailer_user_service_account.email)
}

resource "google_cloudfunctions2_function" "outbound_emailer" {
  depends_on = [
    google_project_iam_member.outbound_emailer_event_receiving,
    google_project_iam_member.outbound_emailer_artifactregistry_reader,
  ]

  name        = "outbound-emailer"
  location    = local.gcp_region
  description = format("Outbound Emailer (%s)", data.archive_file.outbound_emailer_function.output_md5)

  build_config {
    runtime     = local.go_runtime
    entry_point = "SendEmail"

    source {
      storage_source {
        bucket = google_storage_bucket.outbound_emailer_bucket.name
        object = google_storage_bucket_object.outbound_emailer_archive.name
      }
    }
  }

  service_config {
    available_memory               = "128Mi"
    ingress_settings               = "ALLOW_INTERNAL_ONLY"
    all_traffic_on_latest_revision = true
    service_account_email          = google_service_account.outbound_emailer_user_service_account.email

    environment_variables = {
      PF_ENVIRONMENT = local.environment,
      # TODO: use the outbound_emailer_user for this, currently it has permission denied for accessing tables
      # https://dba.stackexchange.com/questions/53914/permission-denied-for-relation-table
      # https://www.postgresql.org/docs/13/sql-alterdefaultprivileges.html
      PRIXFIXE_DATABASE_USER                     = google_sql_user.api_user.name,
      PRIXFIXE_DATABASE_NAME                     = local.database_name,
      // NOTE: if you're creating a cloud function or server for the first time, terraform cannot configure the database connection.
      // You have to go into the Cloud Run interface and deploy a new revision with a database connection, which will persist upon further deployments.
      PRIXFIXE_DATABASE_INSTANCE_CONNECTION_NAME = google_sql_database_instance.dev.connection_name,
      GOOGLE_CLOUD_SECRET_STORE_PREFIX           = format("projects/%d/secrets", data.google_project.project.number)
      GOOGLE_CLOUD_PROJECT_ID                    = data.google_project.project.project_id
    }

    secret_environment_variables {
      key        = "PRIXFIXE_SENDGRID_API_TOKEN"
      project_id = local.project_id
      secret     = google_secret_manager_secret.sendgrid_api_token.secret_id
      version    = "latest"
    }

    secret_environment_variables {
      key        = "PRIXFIXE_SEGMENT_API_TOKEN"
      project_id = local.project_id
      secret     = google_secret_manager_secret.segment_api_token.secret_id
      version    = "latest"
    }
  }

  event_trigger {
    trigger_region        = local.gcp_region
    event_type            = local.pubsub_topic_publish_event
    pubsub_topic          = google_pubsub_topic.outbound_emails_topic.id
    retry_policy          = "RETRY_POLICY_RETRY"
    service_account_email = google_service_account.outbound_emailer_user_service_account.email
  }
}
