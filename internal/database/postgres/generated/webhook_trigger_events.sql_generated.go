// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: webhook_trigger_events.sql

package generated

import (
	"context"
)

const createWebhookTriggerEvent = `-- name: CreateWebhookTriggerEvent :exec

INSERT INTO webhook_trigger_events (
	id,
	trigger_event,
	belongs_to_webhook
) VALUES (
	$1,
	$2,
	$3
)
`

type CreateWebhookTriggerEventParams struct {
	ID               string
	TriggerEvent     WebhookEvent
	BelongsToWebhook string
}

func (q *Queries) CreateWebhookTriggerEvent(ctx context.Context, db DBTX, arg *CreateWebhookTriggerEventParams) error {
	_, err := db.ExecContext(ctx, createWebhookTriggerEvent, arg.ID, arg.TriggerEvent, arg.BelongsToWebhook)
	return err
}
