// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: webhooks_create.sql

package generated

import (
	"context"
)

const CreateWebhook = `-- name: CreateWebhook :exec
INSERT INTO
    webhooks (
    id,
    name,
    content_type,
    url,
    method,
    belongs_to_household
)
VALUES
    ($1, $2, $3, $4, $5, $6)
`

type CreateWebhookParams struct {
	ID                 string `db:"id"`
	Name               string `db:"name"`
	ContentType        string `db:"content_type"`
	Url                string `db:"url"`
	Method             string `db:"method"`
	BelongsToHousehold string `db:"belongs_to_household"`
}

func (q *Queries) CreateWebhook(ctx context.Context, arg *CreateWebhookParams) error {
	_, err := q.exec(ctx, q.createWebhookStmt, CreateWebhook,
		arg.ID,
		arg.Name,
		arg.ContentType,
		arg.Url,
		arg.Method,
		arg.BelongsToHousehold,
	)
	return err
}
