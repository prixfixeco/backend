// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: recipe_step_products_get_one.sql

package generated

import (
	"context"
	"database/sql"
	"time"
)

const GetRecipeStepProduct = `-- name: GetRecipeStepProduct :one
SELECT
	recipe_step_products.id,
	recipe_step_products.name,
	recipe_step_products.type,
	valid_measurement_units.id,
	valid_measurement_units.name,
	valid_measurement_units.description,
	valid_measurement_units.volumetric,
	valid_measurement_units.icon_path,
	valid_measurement_units.universal,
	valid_measurement_units.metric,
	valid_measurement_units.imperial,
	valid_measurement_units.plural_name,
	valid_measurement_units.created_at,
	valid_measurement_units.last_updated_at,
	valid_measurement_units.archived_at,
	recipe_step_products.minimum_quantity_value,
	recipe_step_products.maximum_quantity_value,
	recipe_step_products.quantity_notes,
	recipe_step_products.compostable,
	recipe_step_products.maximum_storage_duration_in_seconds,
	recipe_step_products.minimum_storage_temperature_in_celsius,
	recipe_step_products.maximum_storage_temperature_in_celsius,
	recipe_step_products.storage_instructions,
	recipe_step_products.created_at,
	recipe_step_products.last_updated_at,
	recipe_step_products.archived_at,
	recipe_step_products.belongs_to_recipe_step
FROM recipe_step_products
	JOIN recipe_steps ON recipe_step_products.belongs_to_recipe_step=recipe_steps.id
	JOIN recipes ON recipe_steps.belongs_to_recipe=recipes.id
	JOIN valid_measurement_units ON recipe_step_products.measurement_unit=valid_measurement_units.id
WHERE recipe_step_products.archived_at IS NULL
	AND recipe_step_products.belongs_to_recipe_step = $1
	AND recipe_step_products.id = $2
	AND recipe_steps.archived_at IS NULL
	AND recipe_steps.belongs_to_recipe = $3
	AND recipe_steps.id = $4
	AND recipes.archived_at IS NULL
	AND recipes.id = $5
`

type GetRecipeStepProductParams struct {
	BelongsToRecipeStep string `db:"belongs_to_recipe_step"`
	ID                  string `db:"id"`
	BelongsToRecipe     string `db:"belongs_to_recipe"`
	ID_2                string `db:"id_2"`
	ID_3                string `db:"id_3"`
}

type GetRecipeStepProductRow struct {
	CreatedAt                          time.Time             `db:"created_at"`
	CreatedAt_2                        time.Time             `db:"created_at_2"`
	ArchivedAt_2                       sql.NullTime          `db:"archived_at_2"`
	ArchivedAt                         sql.NullTime          `db:"archived_at"`
	LastUpdatedAt                      sql.NullTime          `db:"last_updated_at"`
	LastUpdatedAt_2                    sql.NullTime          `db:"last_updated_at_2"`
	QuantityNotes                      string                `db:"quantity_notes"`
	IconPath                           string                `db:"icon_path"`
	Description                        string                `db:"description"`
	Name_2                             string                `db:"name_2"`
	StorageInstructions                string                `db:"storage_instructions"`
	PluralName                         string                `db:"plural_name"`
	ID_2                               string                `db:"id_2"`
	Type                               RecipeStepProductType `db:"type"`
	Name                               string                `db:"name"`
	MinimumQuantityValue               string                `db:"minimum_quantity_value"`
	MaximumQuantityValue               string                `db:"maximum_quantity_value"`
	ID                                 string                `db:"id"`
	BelongsToRecipeStep                string                `db:"belongs_to_recipe_step"`
	MaximumStorageTemperatureInCelsius sql.NullString        `db:"maximum_storage_temperature_in_celsius"`
	MinimumStorageTemperatureInCelsius sql.NullString        `db:"minimum_storage_temperature_in_celsius"`
	MaximumStorageDurationInSeconds    sql.NullInt32         `db:"maximum_storage_duration_in_seconds"`
	Volumetric                         sql.NullBool          `db:"volumetric"`
	Compostable                        bool                  `db:"compostable"`
	Imperial                           bool                  `db:"imperial"`
	Metric                             bool                  `db:"metric"`
	Universal                          bool                  `db:"universal"`
}

func (q *Queries) GetRecipeStepProduct(ctx context.Context, arg *GetRecipeStepProductParams) (*GetRecipeStepProductRow, error) {
	row := q.queryRow(ctx, q.getRecipeStepProductStmt, GetRecipeStepProduct,
		arg.BelongsToRecipeStep,
		arg.ID,
		arg.BelongsToRecipe,
		arg.ID_2,
		arg.ID_3,
	)
	var i GetRecipeStepProductRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Type,
		&i.ID_2,
		&i.Name_2,
		&i.Description,
		&i.Volumetric,
		&i.IconPath,
		&i.Universal,
		&i.Metric,
		&i.Imperial,
		&i.PluralName,
		&i.CreatedAt,
		&i.LastUpdatedAt,
		&i.ArchivedAt,
		&i.MinimumQuantityValue,
		&i.MaximumQuantityValue,
		&i.QuantityNotes,
		&i.Compostable,
		&i.MaximumStorageDurationInSeconds,
		&i.MinimumStorageTemperatureInCelsius,
		&i.MaximumStorageTemperatureInCelsius,
		&i.StorageInstructions,
		&i.CreatedAt_2,
		&i.LastUpdatedAt_2,
		&i.ArchivedAt_2,
		&i.BelongsToRecipeStep,
	)
	return &i, err
}
