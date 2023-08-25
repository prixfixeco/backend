// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: recipe_media.sql

package generated

import (
	"context"
	"database/sql"
	"time"
)

const archiveRecipeMedia = `-- name: ArchiveRecipeMedia :execrows

UPDATE recipe_media SET archived_at = NOW() WHERE archived_at IS NULL AND id = $1
`

func (q *Queries) ArchiveRecipeMedia(ctx context.Context, db DBTX, id string) (int64, error) {
	result, err := db.ExecContext(ctx, archiveRecipeMedia, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

const checkRecipeMediaExistence = `-- name: CheckRecipeMediaExistence :one

SELECT EXISTS ( SELECT recipe_media.id FROM recipe_media WHERE recipe_media.archived_at IS NULL AND recipe_media.id = $1 )
`

func (q *Queries) CheckRecipeMediaExistence(ctx context.Context, db DBTX, id string) (bool, error) {
	row := db.QueryRowContext(ctx, checkRecipeMediaExistence, id)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const createRecipeMedia = `-- name: CreateRecipeMedia :exec

INSERT INTO recipe_media (id,belongs_to_recipe,belongs_to_recipe_step,mime_type,internal_path,external_path,"index")
	VALUES ($1,$2,$3,$4,$5,$6,$7)
`

type CreateRecipeMediaParams struct {
	ID                  string
	MimeType            string
	InternalPath        string
	ExternalPath        string
	BelongsToRecipe     sql.NullString
	BelongsToRecipeStep sql.NullString
	Index               int32
}

func (q *Queries) CreateRecipeMedia(ctx context.Context, db DBTX, arg *CreateRecipeMediaParams) error {
	_, err := db.ExecContext(ctx, createRecipeMedia,
		arg.ID,
		arg.BelongsToRecipe,
		arg.BelongsToRecipeStep,
		arg.MimeType,
		arg.InternalPath,
		arg.ExternalPath,
		arg.Index,
	)
	return err
}

const getRecipeMedia = `-- name: GetRecipeMedia :one

SELECT
	recipe_media.id,
	recipe_media.belongs_to_recipe,
	recipe_media.belongs_to_recipe_step,
	recipe_media.mime_type,
	recipe_media.internal_path,
	recipe_media.external_path,
	recipe_media.index,
	recipe_media.created_at,
	recipe_media.last_updated_at,
	recipe_media.archived_at
FROM recipe_media
WHERE recipe_media.archived_at IS NULL
	AND recipe_media.id = $1
`

type GetRecipeMediaRow struct {
	CreatedAt           time.Time
	LastUpdatedAt       sql.NullTime
	ArchivedAt          sql.NullTime
	ID                  string
	MimeType            string
	InternalPath        string
	ExternalPath        string
	BelongsToRecipe     sql.NullString
	BelongsToRecipeStep sql.NullString
	Index               int32
}

func (q *Queries) GetRecipeMedia(ctx context.Context, db DBTX, id string) (*GetRecipeMediaRow, error) {
	row := db.QueryRowContext(ctx, getRecipeMedia, id)
	var i GetRecipeMediaRow
	err := row.Scan(
		&i.ID,
		&i.BelongsToRecipe,
		&i.BelongsToRecipeStep,
		&i.MimeType,
		&i.InternalPath,
		&i.ExternalPath,
		&i.Index,
		&i.CreatedAt,
		&i.LastUpdatedAt,
		&i.ArchivedAt,
	)
	return &i, err
}

const getRecipeMediaForRecipe = `-- name: GetRecipeMediaForRecipe :many

SELECT
	recipe_media.id,
	recipe_media.belongs_to_recipe,
	recipe_media.belongs_to_recipe_step,
	recipe_media.mime_type,
	recipe_media.internal_path,
	recipe_media.external_path,
	recipe_media.index,
	recipe_media.created_at,
	recipe_media.last_updated_at,
	recipe_media.archived_at
FROM recipe_media
WHERE recipe_media.belongs_to_recipe = $1
	AND recipe_media.belongs_to_recipe_step IS NULL
	AND recipe_media.archived_at IS NULL
GROUP BY recipe_media.id
ORDER BY recipe_media.id
`

type GetRecipeMediaForRecipeRow struct {
	CreatedAt           time.Time
	LastUpdatedAt       sql.NullTime
	ArchivedAt          sql.NullTime
	ID                  string
	MimeType            string
	InternalPath        string
	ExternalPath        string
	BelongsToRecipe     sql.NullString
	BelongsToRecipeStep sql.NullString
	Index               int32
}

func (q *Queries) GetRecipeMediaForRecipe(ctx context.Context, db DBTX, belongsToRecipe sql.NullString) ([]*GetRecipeMediaForRecipeRow, error) {
	rows, err := db.QueryContext(ctx, getRecipeMediaForRecipe, belongsToRecipe)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*GetRecipeMediaForRecipeRow{}
	for rows.Next() {
		var i GetRecipeMediaForRecipeRow
		if err := rows.Scan(
			&i.ID,
			&i.BelongsToRecipe,
			&i.BelongsToRecipeStep,
			&i.MimeType,
			&i.InternalPath,
			&i.ExternalPath,
			&i.Index,
			&i.CreatedAt,
			&i.LastUpdatedAt,
			&i.ArchivedAt,
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

const getRecipeMediaForRecipeStep = `-- name: GetRecipeMediaForRecipeStep :many

SELECT
	recipe_media.id,
	recipe_media.belongs_to_recipe,
	recipe_media.belongs_to_recipe_step,
	recipe_media.mime_type,
	recipe_media.internal_path,
	recipe_media.external_path,
	recipe_media.index,
	recipe_media.created_at,
	recipe_media.last_updated_at,
	recipe_media.archived_at
FROM recipe_media
WHERE recipe_media.belongs_to_recipe = $1
	AND recipe_media.belongs_to_recipe_step = $2
	AND recipe_media.archived_at IS NULL
GROUP BY recipe_media.id
ORDER BY recipe_media.id
`

type GetRecipeMediaForRecipeStepParams struct {
	BelongsToRecipe     sql.NullString
	BelongsToRecipeStep sql.NullString
}

type GetRecipeMediaForRecipeStepRow struct {
	CreatedAt           time.Time
	LastUpdatedAt       sql.NullTime
	ArchivedAt          sql.NullTime
	ID                  string
	MimeType            string
	InternalPath        string
	ExternalPath        string
	BelongsToRecipe     sql.NullString
	BelongsToRecipeStep sql.NullString
	Index               int32
}

func (q *Queries) GetRecipeMediaForRecipeStep(ctx context.Context, db DBTX, arg *GetRecipeMediaForRecipeStepParams) ([]*GetRecipeMediaForRecipeStepRow, error) {
	rows, err := db.QueryContext(ctx, getRecipeMediaForRecipeStep, arg.BelongsToRecipe, arg.BelongsToRecipeStep)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*GetRecipeMediaForRecipeStepRow{}
	for rows.Next() {
		var i GetRecipeMediaForRecipeStepRow
		if err := rows.Scan(
			&i.ID,
			&i.BelongsToRecipe,
			&i.BelongsToRecipeStep,
			&i.MimeType,
			&i.InternalPath,
			&i.ExternalPath,
			&i.Index,
			&i.CreatedAt,
			&i.LastUpdatedAt,
			&i.ArchivedAt,
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

const updateRecipeMedia = `-- name: UpdateRecipeMedia :execrows

UPDATE recipe_media
SET
	belongs_to_recipe = $1,
	belongs_to_recipe_step = $2,
	mime_type = $3,
	internal_path = $4,
	external_path = $5,
	"index" = $6,
	last_updated_at = NOW()
WHERE archived_at IS NULL
	AND id = $6
`

type UpdateRecipeMediaParams struct {
	MimeType            string
	InternalPath        string
	ExternalPath        string
	BelongsToRecipe     sql.NullString
	BelongsToRecipeStep sql.NullString
	Index               int32
}

func (q *Queries) UpdateRecipeMedia(ctx context.Context, db DBTX, arg *UpdateRecipeMediaParams) (int64, error) {
	result, err := db.ExecContext(ctx, updateRecipeMedia,
		arg.BelongsToRecipe,
		arg.BelongsToRecipeStep,
		arg.MimeType,
		arg.InternalPath,
		arg.ExternalPath,
		arg.Index,
	)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
