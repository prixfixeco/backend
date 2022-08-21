// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: valid_instruments.sql

package generated

import (
	"context"
)

const ArchiveValidInstrument = `-- name: ArchiveValidInstrument :exec
UPDATE valid_instruments SET archived_on = extract(epoch FROM NOW()) WHERE archived_on IS NULL AND id = $1
`

func (q *Queries) ArchiveValidInstrument(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, ArchiveValidInstrument, id)
	return err
}

const CreateValidInstrument = `-- name: CreateValidInstrument :exec
INSERT INTO valid_instruments (id,name,variant,description,icon_path) VALUES ($1,$2,$3,$4,$5)
`

type CreateValidInstrumentParams struct {
	ID          string
	Name        string
	Variant     string
	Description string
	IconPath    string
}

func (q *Queries) CreateValidInstrument(ctx context.Context, arg *CreateValidInstrumentParams) error {
	_, err := q.db.ExecContext(ctx, CreateValidInstrument,
		arg.ID,
		arg.Name,
		arg.Variant,
		arg.Description,
		arg.IconPath,
	)
	return err
}

const GetRandomValidInstrument = `-- name: GetRandomValidInstrument :one
SELECT
    valid_instruments.id,
    valid_instruments.name,
    valid_instruments.variant,
    valid_instruments.description,
    valid_instruments.icon_path,
    valid_instruments.created_on,
    valid_instruments.last_updated_on,
    valid_instruments.archived_on
FROM valid_instruments
WHERE valid_instruments.archived_on IS NULL
ORDER BY random() LIMIT 1
`

func (q *Queries) GetRandomValidInstrument(ctx context.Context) (*ValidInstruments, error) {
	row := q.db.QueryRowContext(ctx, GetRandomValidInstrument)
	var i ValidInstruments
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Variant,
		&i.Description,
		&i.IconPath,
		&i.CreatedOn,
		&i.LastUpdatedOn,
		&i.ArchivedOn,
	)
	return &i, err
}

const GetTotalValidInstrumentCount = `-- name: GetTotalValidInstrumentCount :one
SELECT COUNT(valid_instruments.id) FROM valid_instruments WHERE valid_instruments.archived_on IS NULL
`

func (q *Queries) GetTotalValidInstrumentCount(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, GetTotalValidInstrumentCount)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const GetValidInstrument = `-- name: GetValidInstrument :one
SELECT
    valid_instruments.id,
    valid_instruments.name,
    valid_instruments.variant,
    valid_instruments.description,
    valid_instruments.icon_path,
    valid_instruments.created_on,
    valid_instruments.last_updated_on,
    valid_instruments.archived_on
FROM valid_instruments
WHERE valid_instruments.archived_on IS NULL
AND valid_instruments.id = $1
`

func (q *Queries) GetValidInstrument(ctx context.Context, id string) (*ValidInstruments, error) {
	row := q.db.QueryRowContext(ctx, GetValidInstrument, id)
	var i ValidInstruments
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Variant,
		&i.Description,
		&i.IconPath,
		&i.CreatedOn,
		&i.LastUpdatedOn,
		&i.ArchivedOn,
	)
	return &i, err
}

const SearchForValidInstruments = `-- name: SearchForValidInstruments :many
SELECT
    valid_instruments.id,
    valid_instruments.name,
    valid_instruments.variant,
    valid_instruments.description,
    valid_instruments.icon_path,
    valid_instruments.created_on,
    valid_instruments.last_updated_on,
    valid_instruments.archived_on
FROM valid_instruments
WHERE valid_instruments.archived_on IS NULL
  AND valid_instruments.name ILIKE $1
  LIMIT 50
`

func (q *Queries) SearchForValidInstruments(ctx context.Context, name string) ([]*ValidInstruments, error) {
	rows, err := q.db.QueryContext(ctx, SearchForValidInstruments, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*ValidInstruments
	for rows.Next() {
		var i ValidInstruments
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Variant,
			&i.Description,
			&i.IconPath,
			&i.CreatedOn,
			&i.LastUpdatedOn,
			&i.ArchivedOn,
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

const UpdateValidInstrument = `-- name: UpdateValidInstrument :exec
UPDATE valid_instruments SET name = $1, variant = $2, description = $3, icon_path = $4, last_updated_on = extract(epoch FROM NOW()) WHERE archived_on IS NULL AND id = $5
`

type UpdateValidInstrumentParams struct {
	Name        string
	Variant     string
	Description string
	IconPath    string
	ID          string
}

func (q *Queries) UpdateValidInstrument(ctx context.Context, arg *UpdateValidInstrumentParams) error {
	_, err := q.db.ExecContext(ctx, UpdateValidInstrument,
		arg.Name,
		arg.Variant,
		arg.Description,
		arg.IconPath,
		arg.ID,
	)
	return err
}

const ValidInstrumentExists = `-- name: ValidInstrumentExists :one
SELECT EXISTS ( SELECT valid_instruments.id FROM valid_instruments WHERE valid_instruments.archived_on IS NULL AND valid_instruments.id = $1 )
`

func (q *Queries) ValidInstrumentExists(ctx context.Context, id string) (bool, error) {
	row := q.db.QueryRowContext(ctx, ValidInstrumentExists, id)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}
