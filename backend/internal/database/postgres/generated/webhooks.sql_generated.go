// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: webhooks.sql

package generated

import (
	"context"
	"database/sql"
	"time"
)

const archiveWebhook = `-- name: ArchiveWebhook :execrows
UPDATE webhooks SET
	archived_at = NOW()
WHERE archived_at IS NULL
	AND id = $1
	AND belongs_to_household = $2
`

type ArchiveWebhookParams struct {
	ID                 string
	BelongsToHousehold string
}

func (q *Queries) ArchiveWebhook(ctx context.Context, db DBTX, arg *ArchiveWebhookParams) (int64, error) {
	result, err := db.ExecContext(ctx, archiveWebhook, arg.ID, arg.BelongsToHousehold)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

const checkWebhookExistence = `-- name: CheckWebhookExistence :one
SELECT EXISTS(
	SELECT webhooks.id
	FROM webhooks
	WHERE webhooks.archived_at IS NULL
	AND webhooks.id = $1
	AND webhooks.belongs_to_household = $2
)
`

type CheckWebhookExistenceParams struct {
	ID                 string
	BelongsToHousehold string
}

func (q *Queries) CheckWebhookExistence(ctx context.Context, db DBTX, arg *CheckWebhookExistenceParams) (bool, error) {
	row := db.QueryRowContext(ctx, checkWebhookExistence, arg.ID, arg.BelongsToHousehold)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const createWebhook = `-- name: CreateWebhook :exec
INSERT INTO webhooks (
	id,
	name,
	content_type,
	url,
	method,
	belongs_to_household
) VALUES (
	$1,
	$2,
	$3,
	$4,
	$5,
	$6
)
`

type CreateWebhookParams struct {
	ID                 string
	Name               string
	ContentType        string
	URL                string
	Method             string
	BelongsToHousehold string
}

func (q *Queries) CreateWebhook(ctx context.Context, db DBTX, arg *CreateWebhookParams) error {
	_, err := db.ExecContext(ctx, createWebhook,
		arg.ID,
		arg.Name,
		arg.ContentType,
		arg.URL,
		arg.Method,
		arg.BelongsToHousehold,
	)
	return err
}

const getWebhook = `-- name: GetWebhook :many
SELECT
	webhooks.id as webhook_id,
	webhooks.name as webhook_name,
	webhooks.content_type as webhook_content_type,
	webhooks.url as webhook_url,
	webhooks.method as webhook_method,
	webhook_trigger_events.id as webhook_trigger_event_id,
	webhook_trigger_events.trigger_event as webhook_trigger_event_trigger_event,
	webhook_trigger_events.belongs_to_webhook as webhook_trigger_event_belongs_to_webhook,
	webhook_trigger_events.created_at as webhook_trigger_event_created_at,
	webhook_trigger_events.archived_at as webhook_trigger_event_archived_at,
	webhooks.created_at as webhook_created_at,
	webhooks.last_updated_at as webhook_last_updated_at,
	webhooks.archived_at as webhook_archived_at,
	webhooks.belongs_to_household as webhook_belongs_to_household
FROM webhooks
	JOIN webhook_trigger_events ON webhooks.id = webhook_trigger_events.belongs_to_webhook
WHERE webhook_trigger_events.archived_at IS NULL
	AND webhooks.archived_at IS NULL
	AND webhooks.belongs_to_household = $1
	AND webhooks.id = $2
`

type GetWebhookParams struct {
	BelongsToHousehold string
	ID                 string
}

type GetWebhookRow struct {
	WebhookID                           string
	WebhookName                         string
	WebhookContentType                  string
	WebhookUrl                          string
	WebhookMethod                       string
	WebhookTriggerEventID               string
	WebhookTriggerEventTriggerEvent     WebhookEvent
	WebhookTriggerEventBelongsToWebhook string
	WebhookTriggerEventCreatedAt        time.Time
	WebhookTriggerEventArchivedAt       sql.NullTime
	WebhookCreatedAt                    time.Time
	WebhookLastUpdatedAt                sql.NullTime
	WebhookArchivedAt                   sql.NullTime
	WebhookBelongsToHousehold           string
}

func (q *Queries) GetWebhook(ctx context.Context, db DBTX, arg *GetWebhookParams) ([]*GetWebhookRow, error) {
	rows, err := db.QueryContext(ctx, getWebhook, arg.BelongsToHousehold, arg.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*GetWebhookRow{}
	for rows.Next() {
		var i GetWebhookRow
		if err := rows.Scan(
			&i.WebhookID,
			&i.WebhookName,
			&i.WebhookContentType,
			&i.WebhookUrl,
			&i.WebhookMethod,
			&i.WebhookTriggerEventID,
			&i.WebhookTriggerEventTriggerEvent,
			&i.WebhookTriggerEventBelongsToWebhook,
			&i.WebhookTriggerEventCreatedAt,
			&i.WebhookTriggerEventArchivedAt,
			&i.WebhookCreatedAt,
			&i.WebhookLastUpdatedAt,
			&i.WebhookArchivedAt,
			&i.WebhookBelongsToHousehold,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getWebhooksForHousehold = `-- name: GetWebhooksForHousehold :many
SELECT
	webhooks.id,
	webhooks.name,
	webhooks.content_type,
	webhooks.url,
	webhooks.method,
	webhook_trigger_events.id,
	webhook_trigger_events.trigger_event,
	webhook_trigger_events.belongs_to_webhook,
	webhook_trigger_events.created_at,
	webhook_trigger_events.archived_at,
	webhooks.created_at,
	webhooks.last_updated_at,
	webhooks.archived_at,
	webhooks.belongs_to_household,
	(
		SELECT COUNT(webhooks.id)
		FROM webhooks
		WHERE webhooks.archived_at IS NULL
			AND
			webhooks.created_at > COALESCE($1, (SELECT NOW() - '999 years'::INTERVAL))
			AND webhooks.created_at < COALESCE($2, (SELECT NOW() + '999 years'::INTERVAL))
			AND (
				webhooks.last_updated_at IS NULL
				OR webhooks.last_updated_at > COALESCE($3, (SELECT NOW() - '999 years'::INTERVAL))
			)
			AND (
				webhooks.last_updated_at IS NULL
				OR webhooks.last_updated_at < COALESCE($4, (SELECT NOW() + '999 years'::INTERVAL))
			)
			AND (NOT COALESCE($5, false)::boolean OR webhooks.archived_at = NULL)
			AND webhooks.belongs_to_household = $6
	) AS filtered_count,
	(
		SELECT COUNT(webhooks.id)
		FROM webhooks
		WHERE webhooks.archived_at IS NULL
			AND webhooks.belongs_to_household = $6
			AND webhook_trigger_events.archived_at IS NULL
	) AS total_count
FROM webhooks
	JOIN webhook_trigger_events ON webhooks.id = webhook_trigger_events.belongs_to_webhook
WHERE webhooks.archived_at IS NULL
	AND webhooks.created_at > COALESCE($1, (SELECT NOW() - '999 years'::INTERVAL))
	AND webhooks.created_at < COALESCE($2, (SELECT NOW() + '999 years'::INTERVAL))
	AND (
		webhooks.last_updated_at IS NULL
		OR webhooks.last_updated_at > COALESCE($4, (SELECT NOW() - '999 years'::INTERVAL))
	)
	AND (
		webhooks.last_updated_at IS NULL
		OR webhooks.last_updated_at < COALESCE($3, (SELECT NOW() + '999 years'::INTERVAL))
	)
			AND (NOT COALESCE($5, false)::boolean OR webhooks.archived_at = NULL)
	AND webhooks.belongs_to_household = $6
	AND webhook_trigger_events.archived_at IS NULL
LIMIT $8
OFFSET $7
`

type GetWebhooksForHouseholdParams struct {
	CreatedAfter       sql.NullTime
	CreatedBefore      sql.NullTime
	UpdatedBefore      sql.NullTime
	UpdatedAfter       sql.NullTime
	BelongsToHousehold string
	QueryOffset        sql.NullInt32
	QueryLimit         sql.NullInt32
	IncludeArchived    sql.NullBool
}

type GetWebhooksForHouseholdRow struct {
	ID                 string
	Name               string
	ContentType        string
	URL                string
	Method             string
	ID_2               string
	TriggerEvent       WebhookEvent
	BelongsToWebhook   string
	CreatedAt          time.Time
	ArchivedAt         sql.NullTime
	CreatedAt_2        time.Time
	LastUpdatedAt      sql.NullTime
	ArchivedAt_2       sql.NullTime
	BelongsToHousehold string
	FilteredCount      int64
	TotalCount         int64
}

func (q *Queries) GetWebhooksForHousehold(ctx context.Context, db DBTX, arg *GetWebhooksForHouseholdParams) ([]*GetWebhooksForHouseholdRow, error) {
	rows, err := db.QueryContext(ctx, getWebhooksForHousehold,
		arg.CreatedAfter,
		arg.CreatedBefore,
		arg.UpdatedBefore,
		arg.UpdatedAfter,
		arg.IncludeArchived,
		arg.BelongsToHousehold,
		arg.QueryOffset,
		arg.QueryLimit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*GetWebhooksForHouseholdRow{}
	for rows.Next() {
		var i GetWebhooksForHouseholdRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.ContentType,
			&i.URL,
			&i.Method,
			&i.ID_2,
			&i.TriggerEvent,
			&i.BelongsToWebhook,
			&i.CreatedAt,
			&i.ArchivedAt,
			&i.CreatedAt_2,
			&i.LastUpdatedAt,
			&i.ArchivedAt_2,
			&i.BelongsToHousehold,
			&i.FilteredCount,
			&i.TotalCount,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getWebhooksForHouseholdAndEvent = `-- name: GetWebhooksForHouseholdAndEvent :many
SELECT
	webhooks.id,
	webhooks.name,
	webhooks.content_type,
	webhooks.url,
	webhooks.method,
	webhooks.created_at,
	webhooks.last_updated_at,
	webhooks.archived_at,
	webhooks.belongs_to_household
FROM webhooks
	JOIN webhook_trigger_events ON webhooks.id = webhook_trigger_events.belongs_to_webhook
WHERE webhook_trigger_events.archived_at IS NULL
	AND webhook_trigger_events.trigger_event = $1
	AND webhooks.belongs_to_household = $2
	AND webhooks.archived_at IS NULL
`

type GetWebhooksForHouseholdAndEventParams struct {
	TriggerEvent       WebhookEvent
	BelongsToHousehold string
}

func (q *Queries) GetWebhooksForHouseholdAndEvent(ctx context.Context, db DBTX, arg *GetWebhooksForHouseholdAndEventParams) ([]*Webhooks, error) {
	rows, err := db.QueryContext(ctx, getWebhooksForHouseholdAndEvent, arg.TriggerEvent, arg.BelongsToHousehold)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*Webhooks{}
	for rows.Next() {
		var i Webhooks
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.ContentType,
			&i.URL,
			&i.Method,
			&i.CreatedAt,
			&i.LastUpdatedAt,
			&i.ArchivedAt,
			&i.BelongsToHousehold,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
