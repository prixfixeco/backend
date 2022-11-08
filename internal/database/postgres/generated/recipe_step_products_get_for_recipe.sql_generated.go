// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: recipe_step_products_get_for_recipe.sql

package generated

import (
	"context"
	"database/sql"
	"time"
)

const GetRecipeStepProductsForRecipe = `-- name: GetRecipeStepProductsForRecipe :many
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
	LEFT OUTER JOIN valid_measurement_units ON recipe_step_products.measurement_unit=valid_measurement_units.id
WHERE recipe_step_products.archived_at IS NULL
	AND recipe_steps.archived_at IS NULL
	AND recipe_steps.belongs_to_recipe = $1
	AND recipes.archived_at IS NULL
	AND recipes.id = $1
`

type GetRecipeStepProductsForRecipeRow struct {
	CreatedAt_2                        time.Time             `db:"created_at_2"`
	CreatedAt                          sql.NullTime          `db:"created_at"`
	ArchivedAt_2                       sql.NullTime          `db:"archived_at_2"`
	LastUpdatedAt_2                    sql.NullTime          `db:"last_updated_at_2"`
	ArchivedAt                         sql.NullTime          `db:"archived_at"`
	LastUpdatedAt                      sql.NullTime          `db:"last_updated_at"`
	Name                               string                `db:"name"`
	Type                               RecipeStepProductType `db:"type"`
	StorageInstructions                string                `db:"storage_instructions"`
	QuantityNotes                      string                `db:"quantity_notes"`
	MaximumQuantityValue               string                `db:"maximum_quantity_value"`
	MinimumQuantityValue               string                `db:"minimum_quantity_value"`
	ID                                 string                `db:"id"`
	BelongsToRecipeStep                string                `db:"belongs_to_recipe_step"`
	ID_2                               sql.NullString        `db:"id_2"`
	Name_2                             sql.NullString        `db:"name_2"`
	Description                        sql.NullString        `db:"description"`
	IconPath                           sql.NullString        `db:"icon_path"`
	MaximumStorageTemperatureInCelsius sql.NullString        `db:"maximum_storage_temperature_in_celsius"`
	MinimumStorageTemperatureInCelsius sql.NullString        `db:"minimum_storage_temperature_in_celsius"`
	PluralName                         sql.NullString        `db:"plural_name"`
	MaximumStorageDurationInSeconds    sql.NullInt32         `db:"maximum_storage_duration_in_seconds"`
	Volumetric                         sql.NullBool          `db:"volumetric"`
	Universal                          sql.NullBool          `db:"universal"`
	Metric                             sql.NullBool          `db:"metric"`
	Imperial                           sql.NullBool          `db:"imperial"`
	Compostable                        bool                  `db:"compostable"`
}

func (q *Queries) GetRecipeStepProductsForRecipe(ctx context.Context, belongsToRecipe string) ([]*GetRecipeStepProductsForRecipeRow, error) {
	rows, err := q.query(ctx, q.getRecipeStepProductsForRecipeStmt, GetRecipeStepProductsForRecipe, belongsToRecipe)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*GetRecipeStepProductsForRecipeRow{}
	for rows.Next() {
		var i GetRecipeStepProductsForRecipeRow
		if err := rows.Scan(
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
