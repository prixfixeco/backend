// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: recipe_steps_get_one_by_id.sql

package generated

import (
	"context"
	"database/sql"
	"time"
)

const GetRecipeStepByID = `-- name: GetRecipeStepByID :one
SELECT
	recipe_steps.id,
	recipe_steps.index,
	valid_preparations.id,
	valid_preparations.name,
	valid_preparations.description,
	valid_preparations.icon_path,
	valid_preparations.yields_nothing,
	valid_preparations.restrict_to_ingredients,
	valid_preparations.zero_ingredients_allowable,
	valid_preparations.past_tense,
	valid_preparations.created_at,
	valid_preparations.last_updated_at,
	valid_preparations.archived_at,
	recipe_steps.minimum_estimated_time_in_seconds,
	recipe_steps.maximum_estimated_time_in_seconds,
	recipe_steps.minimum_temperature_in_celsius,
	recipe_steps.maximum_temperature_in_celsius,
	recipe_steps.notes,
	recipe_steps.explicit_instructions,
	recipe_steps.optional,
	recipe_steps.created_at,
	recipe_steps.last_updated_at,
	recipe_steps.archived_at,
	recipe_steps.belongs_to_recipe
FROM recipe_steps
	JOIN recipes ON recipe_steps.belongs_to_recipe=recipes.id
	JOIN valid_preparations ON recipe_steps.preparation_id=valid_preparations.id
WHERE recipe_steps.archived_at IS NULL
	AND recipe_steps.id = $1
`

type GetRecipeStepByIDRow struct {
	CreatedAt                     time.Time      `db:"created_at"`
	CreatedAt_2                   time.Time      `db:"created_at_2"`
	LastUpdatedAt_2               sql.NullTime   `db:"last_updated_at_2"`
	ArchivedAt_2                  sql.NullTime   `db:"archived_at_2"`
	LastUpdatedAt                 sql.NullTime   `db:"last_updated_at"`
	ArchivedAt                    sql.NullTime   `db:"archived_at"`
	IconPath                      string         `db:"icon_path"`
	Description                   string         `db:"description"`
	ExplicitInstructions          string         `db:"explicit_instructions"`
	PastTense                     string         `db:"past_tense"`
	Name                          string         `db:"name"`
	ID_2                          string         `db:"id_2"`
	Notes                         string         `db:"notes"`
	ID                            string         `db:"id"`
	BelongsToRecipe               string         `db:"belongs_to_recipe"`
	MinimumTemperatureInCelsius   sql.NullString `db:"minimum_temperature_in_celsius"`
	MaximumTemperatureInCelsius   sql.NullString `db:"maximum_temperature_in_celsius"`
	MaximumEstimatedTimeInSeconds sql.NullInt64  `db:"maximum_estimated_time_in_seconds"`
	MinimumEstimatedTimeInSeconds sql.NullInt64  `db:"minimum_estimated_time_in_seconds"`
	Index                         int32          `db:"index"`
	ZeroIngredientsAllowable      bool           `db:"zero_ingredients_allowable"`
	Optional                      bool           `db:"optional"`
	RestrictToIngredients         bool           `db:"restrict_to_ingredients"`
	YieldsNothing                 bool           `db:"yields_nothing"`
}

func (q *Queries) GetRecipeStepByID(ctx context.Context, db DBTX, id string) (*GetRecipeStepByIDRow, error) {
	row := db.QueryRowContext(ctx, GetRecipeStepByID, id)
	var i GetRecipeStepByIDRow
	err := row.Scan(
		&i.ID,
		&i.Index,
		&i.ID_2,
		&i.Name,
		&i.Description,
		&i.IconPath,
		&i.YieldsNothing,
		&i.RestrictToIngredients,
		&i.ZeroIngredientsAllowable,
		&i.PastTense,
		&i.CreatedAt,
		&i.LastUpdatedAt,
		&i.ArchivedAt,
		&i.MinimumEstimatedTimeInSeconds,
		&i.MaximumEstimatedTimeInSeconds,
		&i.MinimumTemperatureInCelsius,
		&i.MaximumTemperatureInCelsius,
		&i.Notes,
		&i.ExplicitInstructions,
		&i.Optional,
		&i.CreatedAt_2,
		&i.LastUpdatedAt_2,
		&i.ArchivedAt_2,
		&i.BelongsToRecipe,
	)
	return &i, err
}
