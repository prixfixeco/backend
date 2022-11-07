// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: valid_ingredients_search.sql

package generated

import (
	"context"
	"database/sql"
	"time"
)

const SearchForValidIngredient = `-- name: SearchForValidIngredient :exec
SELECT
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
	valid_ingredients.archived_at
FROM valid_ingredients
WHERE valid_ingredients.name ILIKE $1
	AND valid_ingredients.archived_at IS NULL
LIMIT 50
`

type SearchForValidIngredientRow struct {
	CreatedAt                               time.Time      `db:"created_at"`
	ArchivedAt                              sql.NullTime   `db:"archived_at"`
	LastUpdatedAt                           sql.NullTime   `db:"last_updated_at"`
	Name                                    string         `db:"name"`
	Description                             string         `db:"description"`
	Warning                                 string         `db:"warning"`
	StorageInstructions                     string         `db:"storage_instructions"`
	PluralName                              string         `db:"plural_name"`
	IconPath                                string         `db:"icon_path"`
	ID                                      string         `db:"id"`
	MaximumIdealStorageTemperatureInCelsius sql.NullString `db:"maximum_ideal_storage_temperature_in_celsius"`
	MinimumIdealStorageTemperatureInCelsius sql.NullString `db:"minimum_ideal_storage_temperature_in_celsius"`
	IsLiquid                                sql.NullBool   `db:"is_liquid"`
	ContainsShellfish                       bool           `db:"contains_shellfish"`
	AnimalFlesh                             bool           `db:"animal_flesh"`
	Volumetric                              bool           `db:"volumetric"`
	ContainsWheat                           bool           `db:"contains_wheat"`
	ContainsSoy                             bool           `db:"contains_soy"`
	AnimalDerived                           bool           `db:"animal_derived"`
	ContainsTreeNut                         bool           `db:"contains_tree_nut"`
	RestrictToPreparations                  bool           `db:"restrict_to_preparations"`
	ContainsFish                            bool           `db:"contains_fish"`
	ContainsSesame                          bool           `db:"contains_sesame"`
	ContainsPeanut                          bool           `db:"contains_peanut"`
	ContainsDairy                           bool           `db:"contains_dairy"`
	ContainsEgg                             bool           `db:"contains_egg"`
	ContainsGluten                          bool           `db:"contains_gluten"`
}

func (q *Queries) SearchForValidIngredient(ctx context.Context, db DBTX, name string) error {
	_, err := db.ExecContext(ctx, SearchForValidIngredient, name)
	return err
}
