// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: webhooks.sql

package generated

import (
	"context"
)

const ArchiveWebhook = `-- name: ArchiveWebhook :exec
UPDATE webhooks SET
	last_updated_on = extract(epoch FROM NOW()),
	archived_on = extract(epoch FROM NOW())
WHERE archived_on IS NULL
AND belongs_to_household = $1
AND id = $2
`

type ArchiveWebhookParams struct {
	BelongsToHousehold string
	ID                 string
}

func (q *Queries) ArchiveWebhook(ctx context.Context, arg *ArchiveWebhookParams) error {
	_, err := q.db.ExecContext(ctx, ArchiveWebhook, arg.BelongsToHousehold, arg.ID)
	return err
}

const CreateWebhook = `-- name: CreateWebhook :exec
INSERT INTO webhooks (id,name,content_type,url,method,events,data_types,topics,belongs_to_household) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
`

type CreateWebhookParams struct {
	ID                 string
	Name               string
	ContentType        string
	Url                string
	Method             string
	Events             string
	DataTypes          string
	Topics             string
	BelongsToHousehold string
}

func (q *Queries) CreateWebhook(ctx context.Context, arg *CreateWebhookParams) error {
	_, err := q.db.ExecContext(ctx, CreateWebhook,
		arg.ID,
		arg.Name,
		arg.ContentType,
		arg.Url,
		arg.Method,
		arg.Events,
		arg.DataTypes,
		arg.Topics,
		arg.BelongsToHousehold,
	)
	return err
}

const GetAllWebhooksCount = `-- name: GetAllWebhooksCount :one
SELECT COUNT(webhooks.id) FROM webhooks WHERE webhooks.archived_on IS NULL
`

func (q *Queries) GetAllWebhooksCount(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, GetAllWebhooksCount)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const GetWebhook = `-- name: GetWebhook :one
SELECT webhooks.id, webhooks.name, webhooks.content_type, webhooks.url, webhooks.method, webhooks.events, webhooks.data_types, webhooks.topics, webhooks.created_on, webhooks.last_updated_on, webhooks.archived_on, webhooks.belongs_to_household FROM webhooks WHERE webhooks.archived_on IS NULL AND webhooks.belongs_to_household = $1 AND webhooks.id = $2
`

type GetWebhookParams struct {
	BelongsToHousehold string
	ID                 string
}

func (q *Queries) GetWebhook(ctx context.Context, arg *GetWebhookParams) (*Webhooks, error) {
	row := q.db.QueryRowContext(ctx, GetWebhook, arg.BelongsToHousehold, arg.ID)
	var i Webhooks
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.ContentType,
		&i.Url,
		&i.Method,
		&i.Events,
		&i.DataTypes,
		&i.Topics,
		&i.CreatedOn,
		&i.LastUpdatedOn,
		&i.ArchivedOn,
		&i.BelongsToHousehold,
	)
	return &i, err
}

const WebhookExists = `-- name: WebhookExists :one
SELECT EXISTS ( SELECT webhooks.id FROM webhooks WHERE webhooks.archived_on IS NULL AND webhooks.belongs_to_household = $1 AND webhooks.id = $2 )
`

type WebhookExistsParams struct {
	BelongsToHousehold string
	ID                 string
}

func (q *Queries) WebhookExists(ctx context.Context, arg *WebhookExistsParams) (bool, error) {
	row := q.db.QueryRowContext(ctx, WebhookExists, arg.BelongsToHousehold, arg.ID)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}
