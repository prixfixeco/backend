// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: valid_preparations.sql

package generated

import (
	"context"
)

const ArchiveValidPreparation = `-- name: ArchiveValidPreparation :exec
UPDATE valid_preparations SET archived_on = extract(epoch FROM NOW()) WHERE archived_on IS NULL AND id = $1
`

func (q *Queries) ArchiveValidPreparation(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, ArchiveValidPreparation, id)
	return err
}

const CreateValidPreparation = `-- name: CreateValidPreparation :exec
INSERT INTO valid_preparations (id,name,description,icon_path) VALUES ($1,$2,$3,$4)
`

type CreateValidPreparationParams struct {
	ID          string
	Name        string
	Description string
	IconPath    string
}

func (q *Queries) CreateValidPreparation(ctx context.Context, arg *CreateValidPreparationParams) error {
	_, err := q.db.ExecContext(ctx, CreateValidPreparation,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.IconPath,
	)
	return err
}

const GetRandomValidPreparation = `-- name: GetRandomValidPreparation :one
SELECT
    valid_preparations.id,
    valid_preparations.name,
    valid_preparations.description,
    valid_preparations.icon_path,
    valid_preparations.created_on,
    valid_preparations.last_updated_on,
    valid_preparations.archived_on
FROM valid_preparations
WHERE valid_preparations.archived_on IS NULL
ORDER BY random() LIMIT 1
`

func (q *Queries) GetRandomValidPreparation(ctx context.Context) (*ValidPreparations, error) {
	row := q.db.QueryRowContext(ctx, GetRandomValidPreparation)
	var i ValidPreparations
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.IconPath,
		&i.CreatedOn,
		&i.LastUpdatedOn,
		&i.ArchivedOn,
	)
	return &i, err
}

const GetTotalValidPreparationsCount = `-- name: GetTotalValidPreparationsCount :one
SELECT COUNT(valid_preparations.id) FROM valid_preparations WHERE valid_preparations.archived_on IS NULL
`

func (q *Queries) GetTotalValidPreparationsCount(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, GetTotalValidPreparationsCount)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const GetValidPreparation = `-- name: GetValidPreparation :one
SELECT
    valid_preparations.id,
    valid_preparations.name,
    valid_preparations.description,
    valid_preparations.icon_path,
    valid_preparations.created_on,
    valid_preparations.last_updated_on,
    valid_preparations.archived_on
FROM valid_preparations
WHERE valid_preparations.archived_on IS NULL
  AND valid_preparations.id = $1
`

func (q *Queries) GetValidPreparation(ctx context.Context, id string) (*ValidPreparations, error) {
	row := q.db.QueryRowContext(ctx, GetValidPreparation, id)
	var i ValidPreparations
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.IconPath,
		&i.CreatedOn,
		&i.LastUpdatedOn,
		&i.ArchivedOn,
	)
	return &i, err
}

const UpdateValidPreparation = `-- name: UpdateValidPreparation :exec
UPDATE valid_preparations SET name = $1, description = $2, icon_path = $3, last_updated_on = extract(epoch FROM NOW()) WHERE archived_on IS NULL AND id = $4
`

type UpdateValidPreparationParams struct {
	Name        string
	Description string
	IconPath    string
	ID          string
}

func (q *Queries) UpdateValidPreparation(ctx context.Context, arg *UpdateValidPreparationParams) error {
	_, err := q.db.ExecContext(ctx, UpdateValidPreparation,
		arg.Name,
		arg.Description,
		arg.IconPath,
		arg.ID,
	)
	return err
}

const ValidPreparationExists = `-- name: ValidPreparationExists :one
SELECT EXISTS ( SELECT valid_preparations.id FROM valid_preparations WHERE valid_preparations.archived_on IS NULL AND valid_preparations.id = $1 )
`

func (q *Queries) ValidPreparationExists(ctx context.Context, id string) (bool, error) {
	row := q.db.QueryRowContext(ctx, ValidPreparationExists, id)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const ValidPreparationsSearch = `-- name: ValidPreparationsSearch :many
SELECT valid_preparations.id, valid_preparations.name, valid_preparations.description, valid_preparations.icon_path, valid_preparations.created_on, valid_preparations.last_updated_on, valid_preparations.archived_on FROM valid_preparations WHERE valid_preparations.archived_on IS NULL AND valid_preparations.name ILIKE $1 LIMIT 50
`

func (q *Queries) ValidPreparationsSearch(ctx context.Context, name string) ([]*ValidPreparations, error) {
	rows, err := q.db.QueryContext(ctx, ValidPreparationsSearch, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*ValidPreparations
	for rows.Next() {
		var i ValidPreparations
		if err := rows.Scan(
			&i.ID,
			&i.Name,
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
