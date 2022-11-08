// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: valid_measurement_conversions_get_all_to_measurement_unit.sql

package generated

import (
	"context"
	"database/sql"
	"time"
)

const GetValidMeasurementConversionsToMeasurementUnit = `-- name: GetValidMeasurementConversionsToMeasurementUnit :many
SELECT
    valid_measurement_conversions.id,
    valid_measurement_units_from.id,
    valid_measurement_units_from.name,
    valid_measurement_units_from.description,
    valid_measurement_units_from.volumetric,
    valid_measurement_units_from.icon_path,
    valid_measurement_units_from.universal,
    valid_measurement_units_from.metric,
    valid_measurement_units_from.imperial,
    valid_measurement_units_from.plural_name,
    valid_measurement_units_from.created_at,
    valid_measurement_units_from.last_updated_at,
    valid_measurement_units_from.archived_at,
    valid_measurement_units_to.id,
    valid_measurement_units_to.name,
    valid_measurement_units_to.description,
    valid_measurement_units_to.volumetric,
    valid_measurement_units_to.icon_path,
    valid_measurement_units_to.universal,
    valid_measurement_units_to.metric,
    valid_measurement_units_to.imperial,
    valid_measurement_units_to.plural_name,
    valid_measurement_units_to.created_at,
    valid_measurement_units_to.last_updated_at,
    valid_measurement_units_to.archived_at,
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
    valid_measurement_conversions.modifier,
    valid_measurement_conversions.notes,
    valid_measurement_conversions.created_at,
    valid_measurement_conversions.last_updated_at,
    valid_measurement_conversions.archived_at
FROM valid_measurement_conversions
         LEFT JOIN valid_ingredients ON valid_measurement_conversions.only_for_ingredient = valid_ingredients.id
         JOIN valid_measurement_units AS valid_measurement_units_from ON valid_measurement_conversions.from_unit = valid_measurement_units_from.id
         JOIN valid_measurement_units AS valid_measurement_units_to ON valid_measurement_conversions.to_unit = valid_measurement_units_to.id
WHERE valid_measurement_conversions.archived_at IS NULL
  AND valid_measurement_units_from.archived_at IS NULL
  AND valid_measurement_units_to.id = $1
  AND valid_measurement_units_to.archived_at IS NULL
`

type GetValidMeasurementConversionsToMeasurementUnitRow struct {
	CreatedAt_4                             time.Time      `db:"created_at_4"`
	CreatedAt_3                             time.Time      `db:"created_at_3"`
	CreatedAt                               time.Time      `db:"created_at"`
	CreatedAt_2                             time.Time      `db:"created_at_2"`
	ArchivedAt_2                            sql.NullTime   `db:"archived_at_2"`
	LastUpdatedAt_2                         sql.NullTime   `db:"last_updated_at_2"`
	ArchivedAt                              sql.NullTime   `db:"archived_at"`
	ArchivedAt_4                            sql.NullTime   `db:"archived_at_4"`
	LastUpdatedAt_3                         sql.NullTime   `db:"last_updated_at_3"`
	ArchivedAt_3                            sql.NullTime   `db:"archived_at_3"`
	LastUpdatedAt_4                         sql.NullTime   `db:"last_updated_at_4"`
	LastUpdatedAt                           sql.NullTime   `db:"last_updated_at"`
	IconPath_3                              string         `db:"icon_path_3"`
	ID_3                                    string         `db:"id_3"`
	Modifier                                string         `db:"modifier"`
	Description_2                           string         `db:"description_2"`
	PluralName                              string         `db:"plural_name"`
	IconPath_2                              string         `db:"icon_path_2"`
	Notes                                   string         `db:"notes"`
	ID                                      string         `db:"id"`
	StorageInstructions                     string         `db:"storage_instructions"`
	PluralName_2                            string         `db:"plural_name_2"`
	IconPath                                string         `db:"icon_path"`
	Description                             string         `db:"description"`
	Name                                    string         `db:"name"`
	ID_4                                    string         `db:"id_4"`
	Name_3                                  string         `db:"name_3"`
	Description_3                           string         `db:"description_3"`
	ID_2                                    string         `db:"id_2"`
	PluralName_3                            string         `db:"plural_name_3"`
	Warning                                 string         `db:"warning"`
	Name_2                                  string         `db:"name_2"`
	MaximumIdealStorageTemperatureInCelsius sql.NullString `db:"maximum_ideal_storage_temperature_in_celsius"`
	MinimumIdealStorageTemperatureInCelsius sql.NullString `db:"minimum_ideal_storage_temperature_in_celsius"`
	IsLiquid                                sql.NullBool   `db:"is_liquid"`
	Volumetric                              sql.NullBool   `db:"volumetric"`
	Volumetric_2                            sql.NullBool   `db:"volumetric_2"`
	ContainsEgg                             bool           `db:"contains_egg"`
	ContainsGluten                          bool           `db:"contains_gluten"`
	AnimalFlesh                             bool           `db:"animal_flesh"`
	Volumetric_3                            bool           `db:"volumetric_3"`
	ContainsFish                            bool           `db:"contains_fish"`
	ContainsDairy                           bool           `db:"contains_dairy"`
	AnimalDerived                           bool           `db:"animal_derived"`
	ContainsSesame                          bool           `db:"contains_sesame"`
	RestrictToPreparations                  bool           `db:"restrict_to_preparations"`
	ContainsShellfish                       bool           `db:"contains_shellfish"`
	ContainsWheat                           bool           `db:"contains_wheat"`
	Imperial_2                              bool           `db:"imperial_2"`
	Metric_2                                bool           `db:"metric_2"`
	Universal_2                             bool           `db:"universal_2"`
	ContainsSoy                             bool           `db:"contains_soy"`
	Metric                                  bool           `db:"metric"`
	Imperial                                bool           `db:"imperial"`
	Universal                               bool           `db:"universal"`
	ContainsTreeNut                         bool           `db:"contains_tree_nut"`
	ContainsPeanut                          bool           `db:"contains_peanut"`
}

func (q *Queries) GetValidMeasurementConversionsToMeasurementUnit(ctx context.Context, id string) ([]*GetValidMeasurementConversionsToMeasurementUnitRow, error) {
	rows, err := q.query(ctx, q.getValidMeasurementConversionsToMeasurementUnitStmt, GetValidMeasurementConversionsToMeasurementUnit, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*GetValidMeasurementConversionsToMeasurementUnitRow{}
	for rows.Next() {
		var i GetValidMeasurementConversionsToMeasurementUnitRow
		if err := rows.Scan(
			&i.ID,
			&i.ID_2,
			&i.Name,
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
			&i.ID_3,
			&i.Name_2,
			&i.Description_2,
			&i.Volumetric_2,
			&i.IconPath_2,
			&i.Universal_2,
			&i.Metric_2,
			&i.Imperial_2,
			&i.PluralName_2,
			&i.CreatedAt_2,
			&i.LastUpdatedAt_2,
			&i.ArchivedAt_2,
			&i.ID_4,
			&i.Name_3,
			&i.Description_3,
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
			&i.Volumetric_3,
			&i.IsLiquid,
			&i.IconPath_3,
			&i.AnimalDerived,
			&i.PluralName_3,
			&i.RestrictToPreparations,
			&i.MinimumIdealStorageTemperatureInCelsius,
			&i.MaximumIdealStorageTemperatureInCelsius,
			&i.StorageInstructions,
			&i.CreatedAt_3,
			&i.LastUpdatedAt_3,
			&i.ArchivedAt_3,
			&i.Modifier,
			&i.Notes,
			&i.CreatedAt_4,
			&i.LastUpdatedAt_4,
			&i.ArchivedAt_4,
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
