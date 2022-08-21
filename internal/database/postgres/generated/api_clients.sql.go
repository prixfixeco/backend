// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: api_clients.sql

package generated

import (
	"context"
	"database/sql"
)

const ArchiveAPIClient = `-- name: ArchiveAPIClient :exec
	UPDATE api_clients SET
		last_updated_on = extract(epoch FROM NOW()),
		archived_on = extract(epoch FROM NOW())
	WHERE archived_on IS NULL
	AND belongs_to_user = $1 AND id = $2
`

type ArchiveAPIClientParams struct {
	BelongsToUser string
	ID            string
}

func (q *Queries) ArchiveAPIClient(ctx context.Context, arg *ArchiveAPIClientParams) error {
	_, err := q.db.ExecContext(ctx, ArchiveAPIClient, arg.BelongsToUser, arg.ID)
	return err
}

const CreateAPIClient = `-- name: CreateAPIClient :exec
	INSERT INTO api_clients (id,name,client_id,secret_key,belongs_to_user) VALUES ($1,$2,$3,$4,$5)
`

type CreateAPIClientParams struct {
	ID            string
	Name          sql.NullString
	ClientID      string
	SecretKey     []byte
	BelongsToUser string
}

func (q *Queries) CreateAPIClient(ctx context.Context, arg *CreateAPIClientParams) error {
	_, err := q.db.ExecContext(ctx, CreateAPIClient,
		arg.ID,
		arg.Name,
		arg.ClientID,
		arg.SecretKey,
		arg.BelongsToUser,
	)
	return err
}

const GetAPIClientByClientID = `-- name: GetAPIClientByClientID :one
	SELECT
		api_clients.id,
		api_clients.name,
		api_clients.client_id,
		api_clients.secret_key,
		api_clients.created_on,
		api_clients.last_updated_on,
		api_clients.archived_on,
		api_clients.belongs_to_user
	FROM api_clients
	WHERE api_clients.archived_on IS NULL
	AND api_clients.client_id = $1
`

type GetAPIClientByClientIDRow struct {
	ID            string
	Name          sql.NullString
	ClientID      string
	SecretKey     []byte
	CreatedOn     int64
	LastUpdatedOn sql.NullInt64
	ArchivedOn    sql.NullInt64
	BelongsToUser string
}

func (q *Queries) GetAPIClientByClientID(ctx context.Context, clientID string) (*GetAPIClientByClientIDRow, error) {
	row := q.db.QueryRowContext(ctx, GetAPIClientByClientID, clientID)
	var i GetAPIClientByClientIDRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.ClientID,
		&i.SecretKey,
		&i.CreatedOn,
		&i.LastUpdatedOn,
		&i.ArchivedOn,
		&i.BelongsToUser,
	)
	return &i, err
}

const GetAPIClientByDatabaseID = `-- name: GetAPIClientByDatabaseID :one
	SELECT
		api_clients.id,
		api_clients.name,
		api_clients.client_id,
		api_clients.secret_key,
		api_clients.created_on,
		api_clients.last_updated_on,
		api_clients.archived_on,
		api_clients.belongs_to_user
	FROM api_clients
	WHERE api_clients.archived_on IS NULL
	AND api_clients.belongs_to_user = $1
	AND api_clients.id = $2
`

type GetAPIClientByDatabaseIDParams struct {
	BelongsToUser string
	ID            string
}

type GetAPIClientByDatabaseIDRow struct {
	ID            string
	Name          sql.NullString
	ClientID      string
	SecretKey     []byte
	CreatedOn     int64
	LastUpdatedOn sql.NullInt64
	ArchivedOn    sql.NullInt64
	BelongsToUser string
}

func (q *Queries) GetAPIClientByDatabaseID(ctx context.Context, arg *GetAPIClientByDatabaseIDParams) (*GetAPIClientByDatabaseIDRow, error) {
	row := q.db.QueryRowContext(ctx, GetAPIClientByDatabaseID, arg.BelongsToUser, arg.ID)
	var i GetAPIClientByDatabaseIDRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.ClientID,
		&i.SecretKey,
		&i.CreatedOn,
		&i.LastUpdatedOn,
		&i.ArchivedOn,
		&i.BelongsToUser,
	)
	return &i, err
}

const GetTotalAPIClientCount = `-- name: GetTotalAPIClientCount :exec
	SELECT COUNT(api_clients.id) FROM api_clients WHERE api_clients.archived_on IS NULL
`

func (q *Queries) GetTotalAPIClientCount(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, GetTotalAPIClientCount)
	return err
}
