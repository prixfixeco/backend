// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: recipes.sql

package generated

import (
	"context"
	"database/sql"
)

const ArchiveRecipe = `-- name: ArchiveRecipe :exec
UPDATE recipes SET archived_on = extract(epoch FROM NOW()) WHERE archived_on IS NULL AND created_by_user = $1 AND id = $2
`

type ArchiveRecipeParams struct {
	CreatedByUser string
	ID            string
}

func (q *Queries) ArchiveRecipe(ctx context.Context, arg *ArchiveRecipeParams) error {
	_, err := q.db.ExecContext(ctx, ArchiveRecipe, arg.CreatedByUser, arg.ID)
	return err
}

const CreateRecipe = `-- name: CreateRecipe :exec
INSERT INTO recipes (id,name,source,description,inspired_by_recipe_id,created_by_user) VALUES ($1,$2,$3,$4,$5,$6)
`

type CreateRecipeParams struct {
	ID                 string
	Name               string
	Source             string
	Description        string
	InspiredByRecipeID sql.NullString
	CreatedByUser      string
}

func (q *Queries) CreateRecipe(ctx context.Context, arg *CreateRecipeParams) error {
	_, err := q.db.ExecContext(ctx, CreateRecipe,
		arg.ID,
		arg.Name,
		arg.Source,
		arg.Description,
		arg.InspiredByRecipeID,
		arg.CreatedByUser,
	)
	return err
}

const GetRecipeByID = `-- name: GetRecipeByID :many
SELECT
    recipes.id,
    recipes.name,
    recipes.source,
    recipes.description,
    recipes.inspired_by_recipe_id,
    recipes.created_on,
    recipes.last_updated_on,
    recipes.archived_on,
    recipes.created_by_user,
    recipe_steps.id,
    recipe_steps.index,
    valid_preparations.id,
    valid_preparations.name,
    valid_preparations.description,
    valid_preparations.icon_path,
    valid_preparations.created_on,
    valid_preparations.last_updated_on,
    valid_preparations.archived_on,
    recipe_steps.minimum_estimated_time_in_seconds,
    recipe_steps.maximum_estimated_time_in_seconds,
    recipe_steps.minimum_temperature_in_celsius,
    recipe_steps.maximum_temperature_in_celsius,
    recipe_steps.notes,
    recipe_steps.optional,
    recipe_steps.created_on,
    recipe_steps.last_updated_on,
    recipe_steps.archived_on,
    recipe_steps.belongs_to_recipe
FROM recipes
         FULL OUTER JOIN recipe_steps ON recipes.id=recipe_steps.belongs_to_recipe
         FULL OUTER JOIN valid_preparations ON recipe_steps.preparation_id=valid_preparations.id
WHERE recipes.archived_on IS NULL
  AND recipes.id = $1
ORDER BY recipe_steps.index
`

type GetRecipeByIDRow struct {
	ID                            sql.NullString
	Name                          sql.NullString
	Source                        sql.NullString
	Description                   sql.NullString
	InspiredByRecipeID            sql.NullString
	CreatedOn                     sql.NullInt64
	LastUpdatedOn                 sql.NullInt64
	ArchivedOn                    sql.NullInt64
	CreatedByUser                 sql.NullString
	ID_2                          sql.NullString
	Index                         sql.NullInt32
	ID_3                          sql.NullString
	Name_2                        sql.NullString
	Description_2                 sql.NullString
	IconPath                      sql.NullString
	CreatedOn_2                   sql.NullInt64
	LastUpdatedOn_2               sql.NullInt64
	ArchivedOn_2                  sql.NullInt64
	MinimumEstimatedTimeInSeconds sql.NullInt64
	MaximumEstimatedTimeInSeconds sql.NullInt64
	MinimumTemperatureInCelsius   sql.NullInt32
	MaximumTemperatureInCelsius   sql.NullInt32
	Notes                         sql.NullString
	Optional                      sql.NullBool
	CreatedOn_3                   sql.NullInt64
	LastUpdatedOn_3               sql.NullInt64
	ArchivedOn_3                  sql.NullInt64
	BelongsToRecipe               sql.NullString
}

func (q *Queries) GetRecipeByID(ctx context.Context, id string) ([]*GetRecipeByIDRow, error) {
	rows, err := q.db.QueryContext(ctx, GetRecipeByID, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetRecipeByIDRow
	for rows.Next() {
		var i GetRecipeByIDRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Source,
			&i.Description,
			&i.InspiredByRecipeID,
			&i.CreatedOn,
			&i.LastUpdatedOn,
			&i.ArchivedOn,
			&i.CreatedByUser,
			&i.ID_2,
			&i.Index,
			&i.ID_3,
			&i.Name_2,
			&i.Description_2,
			&i.IconPath,
			&i.CreatedOn_2,
			&i.LastUpdatedOn_2,
			&i.ArchivedOn_2,
			&i.MinimumEstimatedTimeInSeconds,
			&i.MaximumEstimatedTimeInSeconds,
			&i.MinimumTemperatureInCelsius,
			&i.MaximumTemperatureInCelsius,
			&i.Notes,
			&i.Optional,
			&i.CreatedOn_3,
			&i.LastUpdatedOn_3,
			&i.ArchivedOn_3,
			&i.BelongsToRecipe,
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

const GetRecipeByIDAndAuthorID = `-- name: GetRecipeByIDAndAuthorID :many
SELECT
    recipes.id,
    recipes.name,
    recipes.source,
    recipes.description,
    recipes.inspired_by_recipe_id,
    recipes.created_on,
    recipes.last_updated_on,
    recipes.archived_on,
    recipes.created_by_user,
    recipe_steps.id,
    recipe_steps.index,
    valid_preparations.id,
    valid_preparations.name,
    valid_preparations.description,
    valid_preparations.icon_path,
    valid_preparations.created_on,
    valid_preparations.last_updated_on,
    valid_preparations.archived_on,
    recipe_steps.minimum_estimated_time_in_seconds,
    recipe_steps.maximum_estimated_time_in_seconds,
    recipe_steps.minimum_temperature_in_celsius,
    recipe_steps.maximum_temperature_in_celsius,
    recipe_steps.notes,
    recipe_steps.optional,
    recipe_steps.created_on,
    recipe_steps.last_updated_on,
    recipe_steps.archived_on,
    recipe_steps.belongs_to_recipe
FROM recipes
         FULL OUTER JOIN recipe_steps ON recipes.id=recipe_steps.belongs_to_recipe
         FULL OUTER JOIN valid_preparations ON recipe_steps.preparation_id=valid_preparations.id
WHERE recipes.archived_on IS NULL
  AND recipes.id = $1
  AND recipes.created_by_user = $2
ORDER BY recipe_steps.index
`

type GetRecipeByIDAndAuthorIDParams struct {
	ID            string
	CreatedByUser string
}

type GetRecipeByIDAndAuthorIDRow struct {
	ID                            sql.NullString
	Name                          sql.NullString
	Source                        sql.NullString
	Description                   sql.NullString
	InspiredByRecipeID            sql.NullString
	CreatedOn                     sql.NullInt64
	LastUpdatedOn                 sql.NullInt64
	ArchivedOn                    sql.NullInt64
	CreatedByUser                 sql.NullString
	ID_2                          sql.NullString
	Index                         sql.NullInt32
	ID_3                          sql.NullString
	Name_2                        sql.NullString
	Description_2                 sql.NullString
	IconPath                      sql.NullString
	CreatedOn_2                   sql.NullInt64
	LastUpdatedOn_2               sql.NullInt64
	ArchivedOn_2                  sql.NullInt64
	MinimumEstimatedTimeInSeconds sql.NullInt64
	MaximumEstimatedTimeInSeconds sql.NullInt64
	MinimumTemperatureInCelsius   sql.NullInt32
	MaximumTemperatureInCelsius   sql.NullInt32
	Notes                         sql.NullString
	Optional                      sql.NullBool
	CreatedOn_3                   sql.NullInt64
	LastUpdatedOn_3               sql.NullInt64
	ArchivedOn_3                  sql.NullInt64
	BelongsToRecipe               sql.NullString
}

func (q *Queries) GetRecipeByIDAndAuthorID(ctx context.Context, arg *GetRecipeByIDAndAuthorIDParams) ([]*GetRecipeByIDAndAuthorIDRow, error) {
	rows, err := q.db.QueryContext(ctx, GetRecipeByIDAndAuthorID, arg.ID, arg.CreatedByUser)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetRecipeByIDAndAuthorIDRow
	for rows.Next() {
		var i GetRecipeByIDAndAuthorIDRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Source,
			&i.Description,
			&i.InspiredByRecipeID,
			&i.CreatedOn,
			&i.LastUpdatedOn,
			&i.ArchivedOn,
			&i.CreatedByUser,
			&i.ID_2,
			&i.Index,
			&i.ID_3,
			&i.Name_2,
			&i.Description_2,
			&i.IconPath,
			&i.CreatedOn_2,
			&i.LastUpdatedOn_2,
			&i.ArchivedOn_2,
			&i.MinimumEstimatedTimeInSeconds,
			&i.MaximumEstimatedTimeInSeconds,
			&i.MinimumTemperatureInCelsius,
			&i.MaximumTemperatureInCelsius,
			&i.Notes,
			&i.Optional,
			&i.CreatedOn_3,
			&i.LastUpdatedOn_3,
			&i.ArchivedOn_3,
			&i.BelongsToRecipe,
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

const GetTotalRecipesCount = `-- name: GetTotalRecipesCount :one
SELECT COUNT(recipes.id) FROM recipes WHERE recipes.archived_on IS NULL
`

func (q *Queries) GetTotalRecipesCount(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, GetTotalRecipesCount)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const RecipeExists = `-- name: RecipeExists :one
SELECT EXISTS ( SELECT recipes.id FROM recipes WHERE recipes.archived_on IS NULL AND recipes.id = $1 )
`

func (q *Queries) RecipeExists(ctx context.Context, id string) (bool, error) {
	row := q.db.QueryRowContext(ctx, RecipeExists, id)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const UpdateRecipe = `-- name: UpdateRecipe :exec
UPDATE recipes SET name = $1, source = $2, description = $3, inspired_by_recipe_id = $4, last_updated_on = extract(epoch FROM NOW()) WHERE archived_on IS NULL AND created_by_user = $5 AND id = $6
`

type UpdateRecipeParams struct {
	Name               string
	Source             string
	Description        string
	InspiredByRecipeID sql.NullString
	CreatedByUser      string
	ID                 string
}

func (q *Queries) UpdateRecipe(ctx context.Context, arg *UpdateRecipeParams) error {
	_, err := q.db.ExecContext(ctx, UpdateRecipe,
		arg.Name,
		arg.Source,
		arg.Description,
		arg.InspiredByRecipeID,
		arg.CreatedByUser,
		arg.ID,
	)
	return err
}
