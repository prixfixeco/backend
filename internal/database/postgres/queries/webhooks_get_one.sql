SELECT webhooks.id, webhooks.name, webhooks.content_type, webhooks.url, webhooks.method, webhooks.events, webhooks.data_types, webhooks.topics, webhooks.created_at, webhooks.last_updated_at, webhooks.archived_at, webhooks.belongs_to_household FROM webhooks WHERE webhooks.archived_at IS NULL AND webhooks.belongs_to_household = $1 AND webhooks.id = $2;
