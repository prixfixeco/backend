// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: valid_ingredient_measurement_units_get_one.sql

package generated

import (
	"context"
	"database/sql"
	"time"
)

const GetValidIngredientMeasurementUnit = `-- name: GetValidIngredientMeasurementUnit :exec
SELECT
	valid_ingredient_measurement_units.id,
	valid_ingredient_measurement_units.notes,
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
	valid_ingredient_measurement_units.minimum_allowable_quantity,
	valid_ingredient_measurement_units.maximum_allowable_quantity,
	valid_ingredient_measurement_units.created_at,
	valid_ingredient_measurement_units.last_updated_at,
	valid_ingredient_measurement_units.archived_at
FROM valid_ingredient_measurement_units
	JOIN valid_measurement_units ON valid_ingredient_measurement_units.valid_measurement_unit_id = valid_measurement_units.id
	JOIN valid_ingredients ON valid_ingredient_measurement_units.valid_ingredient_id = valid_ingredients.id
WHERE valid_ingredient_measurement_units.archived_at IS NULL
	AND valid_ingredient_measurement_units.id = $1
`

type GetValidIngredientMeasurementUnitRow struct {
	CreatedAt_3                             time.Time      `db:"created_at_3"`
	CreatedAt_2                             time.Time      `db:"created_at_2"`
	CreatedAt                               time.Time      `db:"created_at"`
	LastUpdatedAt_2                         sql.NullTime   `db:"last_updated_at_2"`
	ArchivedAt_2                            sql.NullTime   `db:"archived_at_2"`
	ArchivedAt                              sql.NullTime   `db:"archived_at"`
	LastUpdatedAt                           sql.NullTime   `db:"last_updated_at"`
	LastUpdatedAt_3                         sql.NullTime   `db:"last_updated_at_3"`
	ArchivedAt_3                            sql.NullTime   `db:"archived_at_3"`
	MaximumAllowableQuantity                string         `db:"maximum_allowable_quantity"`
	MinimumAllowableQuantity                string         `db:"minimum_allowable_quantity"`
	IconPath                                string         `db:"icon_path"`
	ID                                      string         `db:"id"`
	Description                             string         `db:"description"`
	ID_3                                    string         `db:"id_3"`
	Name_2                                  string         `db:"name_2"`
	Description_2                           string         `db:"description_2"`
	Warning                                 string         `db:"warning"`
	Name                                    string         `db:"name"`
	ID_2                                    string         `db:"id_2"`
	Notes                                   string         `db:"notes"`
	StorageInstructions                     string         `db:"storage_instructions"`
	PluralName_2                            string         `db:"plural_name_2"`
	IconPath_2                              string         `db:"icon_path_2"`
	PluralName                              string         `db:"plural_name"`
	MaximumIdealStorageTemperatureInCelsius sql.NullString `db:"maximum_ideal_storage_temperature_in_celsius"`
	MinimumIdealStorageTemperatureInCelsius sql.NullString `db:"minimum_ideal_storage_temperature_in_celsius"`
	IsLiquid                                sql.NullBool   `db:"is_liquid"`
	Volumetric                              sql.NullBool   `db:"volumetric"`
	ContainsSesame                          bool           `db:"contains_sesame"`
	Volumetric_2                            bool           `db:"volumetric_2"`
	ContainsShellfish                       bool           `db:"contains_shellfish"`
	AnimalDerived                           bool           `db:"animal_derived"`
	ContainsSoy                             bool           `db:"contains_soy"`
	RestrictToPreparations                  bool           `db:"restrict_to_preparations"`
	AnimalFlesh                             bool           `db:"animal_flesh"`
	ContainsGluten                          bool           `db:"contains_gluten"`
	ContainsTreeNut                         bool           `db:"contains_tree_nut"`
	ContainsPeanut                          bool           `db:"contains_peanut"`
	ContainsDairy                           bool           `db:"contains_dairy"`
	ContainsEgg                             bool           `db:"contains_egg"`
	ContainsFish                            bool           `db:"contains_fish"`
	Imperial                                bool           `db:"imperial"`
	Metric                                  bool           `db:"metric"`
	Universal                               bool           `db:"universal"`
	ContainsWheat                           bool           `db:"contains_wheat"`
}

func (q *Queries) GetValidIngredientMeasurementUnit(ctx context.Context, db DBTX, id string) error {
	_, err := db.ExecContext(ctx, GetValidIngredientMeasurementUnit, id)
	return err
}
