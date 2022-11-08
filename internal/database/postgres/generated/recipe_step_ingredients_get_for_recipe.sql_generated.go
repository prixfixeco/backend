// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: recipe_step_ingredients_get_for_recipe.sql

package generated

import (
	"context"
	"database/sql"
	"time"
)

const GetRecipeStepIngredientForRecipe = `-- name: GetRecipeStepIngredientForRecipe :many
SELECT
	recipe_step_ingredients.id,
	recipe_step_ingredients.name,
	recipe_step_ingredients.optional,
	valid_ingredients.id,
	valid_ingredients.name,
	valid_ingredients.description,
	valid_ingredients.warning,
	valid_ingredients.contains_egg,
	valid_ingredients.contains_dairy,
	valid_ingredients.contains_peanut,
	valid_ingredients.contains_tree_nut,
	valid_ingredients.contains_soy,
	valid_ingredients.contains_wheat,
	valid_ingredients.contains_shellfish,
	valid_ingredients.contains_sesame,
	valid_ingredients.contains_fish,
	valid_ingredients.contains_gluten,
	valid_ingredients.animal_flesh,
	valid_ingredients.volumetric,
	valid_ingredients.is_liquid,
	valid_ingredients.icon_path,
	valid_ingredients.animal_derived,
	valid_ingredients.plural_name,
	valid_ingredients.restrict_to_preparations,
	valid_ingredients.minimum_ideal_storage_temperature_in_celsius,
	valid_ingredients.maximum_ideal_storage_temperature_in_celsius,
	valid_ingredients.storage_instructions,
	valid_ingredients.created_at,
	valid_ingredients.last_updated_at,
	valid_ingredients.archived_at,
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
	recipe_step_ingredients.minimum_quantity_value,
	recipe_step_ingredients.maximum_quantity_value,
	recipe_step_ingredients.quantity_notes,
	recipe_step_ingredients.product_of_recipe_step,
	recipe_step_ingredients.recipe_step_product_id,
	recipe_step_ingredients.ingredient_notes,
	recipe_step_ingredients.created_at,
	recipe_step_ingredients.last_updated_at,
	recipe_step_ingredients.archived_at,
	recipe_step_ingredients.belongs_to_recipe_step
FROM
	recipe_step_ingredients
	JOIN recipe_steps ON recipe_step_ingredients.belongs_to_recipe_step = recipe_steps.id
	JOIN recipes ON recipe_steps.belongs_to_recipe = recipes.id
	LEFT JOIN valid_ingredients ON recipe_step_ingredients.ingredient_id = valid_ingredients.id
	JOIN valid_measurement_units ON recipe_step_ingredients.measurement_unit = valid_measurement_units.id
WHERE
	recipe_step_ingredients.archived_at IS NULL
	AND recipes.id = $1
GROUP BY
	recipe_step_ingredients.id,
	valid_measurement_units.id,
	valid_ingredients.id
ORDER BY
	recipe_step_ingredients.id
`

type GetRecipeStepIngredientForRecipeRow struct {
	CreatedAt_2                             time.Time      `db:"created_at_2"`
	CreatedAt_3                             time.Time      `db:"created_at_3"`
	CreatedAt                               time.Time      `db:"created_at"`
	LastUpdatedAt                           sql.NullTime   `db:"last_updated_at"`
	ArchivedAt                              sql.NullTime   `db:"archived_at"`
	LastUpdatedAt_2                         sql.NullTime   `db:"last_updated_at_2"`
	ArchivedAt_2                            sql.NullTime   `db:"archived_at_2"`
	LastUpdatedAt_3                         sql.NullTime   `db:"last_updated_at_3"`
	ArchivedAt_3                            sql.NullTime   `db:"archived_at_3"`
	MinimumQuantityValue                    string         `db:"minimum_quantity_value"`
	IngredientNotes                         string         `db:"ingredient_notes"`
	QuantityNotes                           string         `db:"quantity_notes"`
	MaximumQuantityValue                    string         `db:"maximum_quantity_value"`
	Warning                                 string         `db:"warning"`
	Description                             string         `db:"description"`
	Name_2                                  string         `db:"name_2"`
	PluralName_2                            string         `db:"plural_name_2"`
	IconPath_2                              string         `db:"icon_path_2"`
	Description_2                           string         `db:"description_2"`
	Name_3                                  string         `db:"name_3"`
	IconPath                                string         `db:"icon_path"`
	ID_3                                    string         `db:"id_3"`
	PluralName                              string         `db:"plural_name"`
	ID_2                                    string         `db:"id_2"`
	Name                                    string         `db:"name"`
	StorageInstructions                     string         `db:"storage_instructions"`
	ID                                      string         `db:"id"`
	BelongsToRecipeStep                     string         `db:"belongs_to_recipe_step"`
	MinimumIdealStorageTemperatureInCelsius sql.NullString `db:"minimum_ideal_storage_temperature_in_celsius"`
	RecipeStepProductID                     sql.NullString `db:"recipe_step_product_id"`
	MaximumIdealStorageTemperatureInCelsius sql.NullString `db:"maximum_ideal_storage_temperature_in_celsius"`
	Volumetric_2                            sql.NullBool   `db:"volumetric_2"`
	IsLiquid                                sql.NullBool   `db:"is_liquid"`
	ContainsSesame                          bool           `db:"contains_sesame"`
	AnimalFlesh                             bool           `db:"animal_flesh"`
	Universal                               bool           `db:"universal"`
	Metric                                  bool           `db:"metric"`
	Imperial                                bool           `db:"imperial"`
	ContainsGluten                          bool           `db:"contains_gluten"`
	AnimalDerived                           bool           `db:"animal_derived"`
	ContainsFish                            bool           `db:"contains_fish"`
	Volumetric                              bool           `db:"volumetric"`
	ContainsShellfish                       bool           `db:"contains_shellfish"`
	ContainsWheat                           bool           `db:"contains_wheat"`
	ContainsSoy                             bool           `db:"contains_soy"`
	ContainsTreeNut                         bool           `db:"contains_tree_nut"`
	RestrictToPreparations                  bool           `db:"restrict_to_preparations"`
	ContainsPeanut                          bool           `db:"contains_peanut"`
	ContainsDairy                           bool           `db:"contains_dairy"`
	ContainsEgg                             bool           `db:"contains_egg"`
	Optional                                bool           `db:"optional"`
	ProductOfRecipeStep                     bool           `db:"product_of_recipe_step"`
}

func (q *Queries) GetRecipeStepIngredientForRecipe(ctx context.Context, id string) ([]*GetRecipeStepIngredientForRecipeRow, error) {
	rows, err := q.query(ctx, q.getRecipeStepIngredientForRecipeStmt, GetRecipeStepIngredientForRecipe, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*GetRecipeStepIngredientForRecipeRow{}
	for rows.Next() {
		var i GetRecipeStepIngredientForRecipeRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Optional,
			&i.ID_2,
			&i.Name_2,
			&i.Description,
			&i.Warning,
			&i.ContainsEgg,
			&i.ContainsDairy,
			&i.ContainsPeanut,
			&i.ContainsTreeNut,
			&i.ContainsSoy,
			&i.ContainsWheat,
			&i.ContainsShellfish,
			&i.ContainsSesame,
			&i.ContainsFish,
			&i.ContainsGluten,
			&i.AnimalFlesh,
			&i.Volumetric,
			&i.IsLiquid,
			&i.IconPath,
			&i.AnimalDerived,
			&i.PluralName,
			&i.RestrictToPreparations,
			&i.MinimumIdealStorageTemperatureInCelsius,
			&i.MaximumIdealStorageTemperatureInCelsius,
			&i.StorageInstructions,
			&i.CreatedAt,
			&i.LastUpdatedAt,
			&i.ArchivedAt,
			&i.ID_3,
			&i.Name_3,
			&i.Description_2,
			&i.Volumetric_2,
			&i.IconPath_2,
			&i.Universal,
			&i.Metric,
			&i.Imperial,
			&i.PluralName_2,
			&i.CreatedAt_2,
			&i.LastUpdatedAt_2,
			&i.ArchivedAt_2,
			&i.MinimumQuantityValue,
			&i.MaximumQuantityValue,
			&i.QuantityNotes,
			&i.ProductOfRecipeStep,
			&i.RecipeStepProductID,
			&i.IngredientNotes,
			&i.CreatedAt_3,
			&i.LastUpdatedAt_3,
			&i.ArchivedAt_3,
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
