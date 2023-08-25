// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: recipe_step_vessels.sql

package generated

import (
	"context"
	"database/sql"
	"time"
)

const archiveRecipeStepVessel = `-- name: ArchiveRecipeStepVessel :execrows

UPDATE recipe_step_vessels SET archived_at = NOW() WHERE archived_at IS NULL AND belongs_to_recipe_step = $1 AND id = $2
`

type ArchiveRecipeStepVesselParams struct {
	BelongsToRecipeStep string
	ID                  string
}

func (q *Queries) ArchiveRecipeStepVessel(ctx context.Context, db DBTX, arg *ArchiveRecipeStepVesselParams) (int64, error) {
	result, err := db.ExecContext(ctx, archiveRecipeStepVessel, arg.BelongsToRecipeStep, arg.ID)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

const checkRecipeStepVesselExistence = `-- name: CheckRecipeStepVesselExistence :one

SELECT EXISTS (
    SELECT recipe_step_vessels.id
    FROM recipe_step_vessels
        JOIN recipe_steps ON recipe_step_vessels.belongs_to_recipe_step=recipe_steps.id
        JOIN recipes ON recipe_steps.belongs_to_recipe=recipes.id
    WHERE recipe_step_vessels.archived_at IS NULL
        AND recipe_step_vessels.belongs_to_recipe_step = $1
        AND recipe_step_vessels.id = $2
        AND recipe_steps.archived_at IS NULL
        AND recipe_steps.belongs_to_recipe = $3
        AND recipe_steps.id = $1
        AND recipes.archived_at IS NULL
        AND recipes.id = $3
)
`

type CheckRecipeStepVesselExistenceParams struct {
	RecipeStepID       string
	RecipeStepVesselID string
	RecipeID           string
}

func (q *Queries) CheckRecipeStepVesselExistence(ctx context.Context, db DBTX, arg *CheckRecipeStepVesselExistenceParams) (bool, error) {
	row := db.QueryRowContext(ctx, checkRecipeStepVesselExistence, arg.RecipeStepID, arg.RecipeStepVesselID, arg.RecipeID)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const createRecipeStepVessel = `-- name: CreateRecipeStepVessel :exec

INSERT INTO recipe_step_vessels
(id,"name",notes,belongs_to_recipe_step,recipe_step_product_id,valid_vessel_id,vessel_predicate,minimum_quantity,maximum_quantity,unavailable_after_step)
VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
`

type CreateRecipeStepVesselParams struct {
	ID                   string
	Name                 string
	Notes                string
	BelongsToRecipeStep  string
	VesselPredicate      string
	RecipeStepProductID  sql.NullString
	ValidVesselID        sql.NullString
	MaximumQuantity      sql.NullInt32
	MinimumQuantity      int32
	UnavailableAfterStep bool
}

func (q *Queries) CreateRecipeStepVessel(ctx context.Context, db DBTX, arg *CreateRecipeStepVesselParams) error {
	_, err := db.ExecContext(ctx, createRecipeStepVessel,
		arg.ID,
		arg.Name,
		arg.Notes,
		arg.BelongsToRecipeStep,
		arg.RecipeStepProductID,
		arg.ValidVesselID,
		arg.VesselPredicate,
		arg.MinimumQuantity,
		arg.MaximumQuantity,
		arg.UnavailableAfterStep,
	)
	return err
}

const getRecipeStepVessel = `-- name: GetRecipeStepVessel :one

SELECT
    recipe_step_vessels.id,
	valid_vessels.id as valid_vessel_id,
    valid_vessels.name as valid_vessel_name,
    valid_vessels.plural_name as valid_vessel_plural_name,
    valid_vessels.description as valid_vessel_description,
    valid_vessels.icon_path as valid_vessel_icon_path,
    valid_vessels.usable_for_storage as valid_vessel_usable_for_storage,
    valid_vessels.slug as valid_vessel_slug,
    valid_vessels.display_in_summary_lists as valid_vessel_display_in_summary_lists,
    valid_vessels.include_in_generated_instructions as valid_vessel_include_in_generated_instructions,
    valid_vessels.capacity as valid_vessel_capacity,
	valid_measurement_units.id as valid_measurement_unit_id,
	valid_measurement_units.name as valid_measurement_unit_name,
	valid_measurement_units.description as valid_measurement_unit_description,
	valid_measurement_units.volumetric as valid_measurement_unit_volumetric,
	valid_measurement_units.icon_path as valid_measurement_unit_icon_path,
	valid_measurement_units.universal as valid_measurement_unit_universal,
	valid_measurement_units.metric as valid_measurement_unit_metric,
	valid_measurement_units.imperial as valid_measurement_unit_imperial,
	valid_measurement_units.slug as valid_measurement_unit_slug,
	valid_measurement_units.plural_name as valid_measurement_unit_plural_name,
	valid_measurement_units.created_at as valid_measurement_unit_created_at,
	valid_measurement_units.last_updated_at as valid_measurement_unit_last_updated_at,
	valid_measurement_units.archived_at as valid_measurement_unit_archived_at,
    valid_vessels.width_in_millimeters as valid_vessel_width_in_millimeters,
    valid_vessels.length_in_millimeters as valid_vessel_length_in_millimeters,
    valid_vessels.height_in_millimeters as valid_vessel_height_in_millimeters,
    valid_vessels.shape as valid_vessel_shape,
    valid_vessels.created_at as valid_vessel_created_at,
    valid_vessels.last_updated_at as valid_vessel_last_updated_at,
    valid_vessels.archived_at as valid_vessel_archived_at,
    recipe_step_vessels.name,
    recipe_step_vessels.notes,
    recipe_step_vessels.belongs_to_recipe_step,
    recipe_step_vessels.recipe_step_product_id,
    recipe_step_vessels.vessel_predicate,
    recipe_step_vessels.minimum_quantity,
    recipe_step_vessels.maximum_quantity,
    recipe_step_vessels.unavailable_after_step,
    recipe_step_vessels.created_at,
    recipe_step_vessels.last_updated_at,
    recipe_step_vessels.archived_at
FROM recipe_step_vessels
	LEFT JOIN valid_vessels ON recipe_step_vessels.valid_vessel_id=valid_vessels.id
    LEFT JOIN valid_measurement_units ON valid_vessels.capacity_unit=valid_measurement_units.id
	JOIN recipe_steps ON recipe_step_vessels.belongs_to_recipe_step=recipe_steps.id
	JOIN recipes ON recipe_steps.belongs_to_recipe=recipes.id
WHERE recipe_step_vessels.archived_at IS NULL
	AND recipe_step_vessels.belongs_to_recipe_step = $1
	AND recipe_step_vessels.id = $2
	AND recipe_steps.archived_at IS NULL
	AND recipe_steps.belongs_to_recipe = $3
	AND recipe_steps.id = $1
	AND recipes.archived_at IS NULL
	AND recipes.id = $3
`

type GetRecipeStepVesselParams struct {
	RecipeStepID       string
	RecipeStepVesselID string
	RecipeID           string
}

type GetRecipeStepVesselRow struct {
	CreatedAt                                 time.Time
	ValidVesselLastUpdatedAt                  sql.NullTime
	LastUpdatedAt                             sql.NullTime
	ArchivedAt                                sql.NullTime
	ValidVesselArchivedAt                     sql.NullTime
	ValidMeasurementUnitCreatedAt             sql.NullTime
	ValidVesselCreatedAt                      sql.NullTime
	ValidMeasurementUnitArchivedAt            sql.NullTime
	ValidMeasurementUnitLastUpdatedAt         sql.NullTime
	VesselPredicate                           string
	BelongsToRecipeStep                       string
	Notes                                     string
	Name                                      string
	ID                                        string
	ValidVesselCapacity                       sql.NullString
	ValidVesselIconPath                       sql.NullString
	ValidVesselID                             sql.NullString
	ValidVesselName                           sql.NullString
	ValidVesselPluralName                     sql.NullString
	ValidMeasurementUnitSlug                  sql.NullString
	ValidMeasurementUnitPluralName            sql.NullString
	ValidVesselDescription                    sql.NullString
	ValidMeasurementUnitDescription           sql.NullString
	ValidMeasurementUnitName                  sql.NullString
	ValidVesselWidthInMillimeters             sql.NullString
	ValidVesselLengthInMillimeters            sql.NullString
	ValidVesselHeightInMillimeters            sql.NullString
	ValidVesselShape                          NullVesselShape
	ValidMeasurementUnitID                    sql.NullString
	RecipeStepProductID                       sql.NullString
	ValidMeasurementUnitIconPath              sql.NullString
	ValidVesselSlug                           sql.NullString
	MaximumQuantity                           sql.NullInt32
	MinimumQuantity                           int32
	ValidVesselUsableForStorage               sql.NullBool
	ValidVesselDisplayInSummaryLists          sql.NullBool
	ValidVesselIncludeInGeneratedInstructions sql.NullBool
	ValidMeasurementUnitVolumetric            sql.NullBool
	ValidMeasurementUnitImperial              sql.NullBool
	ValidMeasurementUnitMetric                sql.NullBool
	ValidMeasurementUnitUniversal             sql.NullBool
	UnavailableAfterStep                      bool
}

func (q *Queries) GetRecipeStepVessel(ctx context.Context, db DBTX, arg *GetRecipeStepVesselParams) (*GetRecipeStepVesselRow, error) {
	row := db.QueryRowContext(ctx, getRecipeStepVessel, arg.RecipeStepID, arg.RecipeStepVesselID, arg.RecipeID)
	var i GetRecipeStepVesselRow
	err := row.Scan(
		&i.ID,
		&i.ValidVesselID,
		&i.ValidVesselName,
		&i.ValidVesselPluralName,
		&i.ValidVesselDescription,
		&i.ValidVesselIconPath,
		&i.ValidVesselUsableForStorage,
		&i.ValidVesselSlug,
		&i.ValidVesselDisplayInSummaryLists,
		&i.ValidVesselIncludeInGeneratedInstructions,
		&i.ValidVesselCapacity,
		&i.ValidMeasurementUnitID,
		&i.ValidMeasurementUnitName,
		&i.ValidMeasurementUnitDescription,
		&i.ValidMeasurementUnitVolumetric,
		&i.ValidMeasurementUnitIconPath,
		&i.ValidMeasurementUnitUniversal,
		&i.ValidMeasurementUnitMetric,
		&i.ValidMeasurementUnitImperial,
		&i.ValidMeasurementUnitSlug,
		&i.ValidMeasurementUnitPluralName,
		&i.ValidMeasurementUnitCreatedAt,
		&i.ValidMeasurementUnitLastUpdatedAt,
		&i.ValidMeasurementUnitArchivedAt,
		&i.ValidVesselWidthInMillimeters,
		&i.ValidVesselLengthInMillimeters,
		&i.ValidVesselHeightInMillimeters,
		&i.ValidVesselShape,
		&i.ValidVesselCreatedAt,
		&i.ValidVesselLastUpdatedAt,
		&i.ValidVesselArchivedAt,
		&i.Name,
		&i.Notes,
		&i.BelongsToRecipeStep,
		&i.RecipeStepProductID,
		&i.VesselPredicate,
		&i.MinimumQuantity,
		&i.MaximumQuantity,
		&i.UnavailableAfterStep,
		&i.CreatedAt,
		&i.LastUpdatedAt,
		&i.ArchivedAt,
	)
	return &i, err
}

const getRecipeStepVessels = `-- name: GetRecipeStepVessels :many

SELECT
    recipe_step_vessels.id,
    valid_vessels.id as valid_vessel_id,
    valid_vessels.name as valid_vessel_name,
    valid_vessels.plural_name as valid_vessel_plural_name,
    valid_vessels.description as valid_vessel_description,
    valid_vessels.icon_path as valid_vessel_icon_path,
    valid_vessels.usable_for_storage as valid_vessel_usable_for_storage,
    valid_vessels.slug as valid_vessel_slug,
    valid_vessels.display_in_summary_lists as valid_vessel_display_in_summary_lists,
    valid_vessels.include_in_generated_instructions as valid_vessel_include_in_generated_instructions,
    valid_vessels.capacity as valid_vessel_capacity,
    valid_measurement_units.id as valid_measurement_unit_id,
    valid_measurement_units.name as valid_measurement_unit_name,
    valid_measurement_units.description as valid_measurement_unit_description,
    valid_measurement_units.volumetric as valid_measurement_unit_volumetric,
    valid_measurement_units.icon_path as valid_measurement_unit_icon_path,
    valid_measurement_units.universal as valid_measurement_unit_universal,
    valid_measurement_units.metric as valid_measurement_unit_metric,
    valid_measurement_units.imperial as valid_measurement_unit_imperial,
    valid_measurement_units.slug as valid_measurement_unit_slug,
    valid_measurement_units.plural_name as valid_measurement_unit_plural_name,
    valid_measurement_units.created_at as valid_measurement_unit_created_at,
    valid_measurement_units.last_updated_at as valid_measurement_unit_last_updated_at,
    valid_measurement_units.archived_at as valid_measurement_unit_archived_at,
    valid_vessels.width_in_millimeters as valid_vessel_width_in_millimeters,
    valid_vessels.length_in_millimeters as valid_vessel_length_in_millimeters,
    valid_vessels.height_in_millimeters as valid_vessel_height_in_millimeters,
    valid_vessels.shape as valid_vessel_shape,
    valid_vessels.created_at as valid_vessel_created_at,
    valid_vessels.last_updated_at as valid_vessel_last_updated_at,
    valid_vessels.archived_at as valid_vessel_archived_at,
    recipe_step_vessels.name,
    recipe_step_vessels.notes,
    recipe_step_vessels.belongs_to_recipe_step,
    recipe_step_vessels.recipe_step_product_id,
    recipe_step_vessels.vessel_predicate,
    recipe_step_vessels.minimum_quantity,
    recipe_step_vessels.maximum_quantity,
    recipe_step_vessels.unavailable_after_step,
    recipe_step_vessels.created_at,
    recipe_step_vessels.last_updated_at,
    recipe_step_vessels.archived_at,
    (
        SELECT
            COUNT(recipe_step_vessels.id)
        FROM
            recipe_step_vessels
        WHERE
            recipe_step_vessels.archived_at IS NULL
          AND recipe_step_vessels.created_at > COALESCE($1, (SELECT NOW() - interval '999 years'))
          AND recipe_step_vessels.created_at < COALESCE($2, (SELECT NOW() + interval '999 years'))
          AND (recipe_step_vessels.last_updated_at IS NULL OR recipe_step_vessels.last_updated_at > COALESCE($3, (SELECT NOW() - interval '999 years')))
          AND (recipe_step_vessels.last_updated_at IS NULL OR recipe_step_vessels.last_updated_at < COALESCE($4, (SELECT NOW() + interval '999 years')))
    ) as filtered_count,
    (
        SELECT
            COUNT(recipe_step_vessels.id)
        FROM
            recipe_step_vessels
        WHERE
            recipe_step_vessels.archived_at IS NULL
    ) as total_count
FROM recipe_step_vessels
     LEFT JOIN valid_vessels ON recipe_step_vessels.valid_vessel_id=valid_vessels.id
     LEFT JOIN valid_measurement_units ON valid_vessels.capacity_unit=valid_measurement_units.id
     JOIN recipe_steps ON recipe_step_vessels.belongs_to_recipe_step=recipe_steps.id
     JOIN recipes ON recipe_steps.belongs_to_recipe=recipes.id
WHERE recipe_step_vessels.archived_at IS NULL
    AND recipe_step_vessels.belongs_to_recipe_step = $5
    AND recipe_step_vessels.created_at > COALESCE($1, (SELECT NOW() - interval '999 years'))
    AND recipe_step_vessels.created_at < COALESCE($2, (SELECT NOW() + interval '999 years'))
    AND (recipe_step_vessels.last_updated_at IS NULL OR recipe_step_vessels.last_updated_at > COALESCE($3, (SELECT NOW() - interval '999 years')))
    AND (recipe_step_vessels.last_updated_at IS NULL OR recipe_step_vessels.last_updated_at < COALESCE($4, (SELECT NOW() + interval '999 years')))
    AND recipe_steps.belongs_to_recipe = $6
    AND recipe_steps.archived_at IS NULL
    AND recipe_steps.id = $5
    AND recipes.archived_at IS NULL
    AND recipes.id = $6
    OFFSET $7
    LIMIT $8
`

type GetRecipeStepVesselsParams struct {
	CreatedAfter  sql.NullTime
	CreatedBefore sql.NullTime
	UpdatedAfter  sql.NullTime
	UpdatedBefore sql.NullTime
	RecipeStepID  string
	RecipeID      string
	QueryOffset   sql.NullInt32
	QueryLimit    sql.NullInt32
}

type GetRecipeStepVesselsRow struct {
	CreatedAt                                 time.Time
	ValidMeasurementUnitCreatedAt             sql.NullTime
	ArchivedAt                                sql.NullTime
	LastUpdatedAt                             sql.NullTime
	ValidVesselArchivedAt                     sql.NullTime
	ValidVesselLastUpdatedAt                  sql.NullTime
	ValidVesselCreatedAt                      sql.NullTime
	ValidMeasurementUnitArchivedAt            sql.NullTime
	ValidMeasurementUnitLastUpdatedAt         sql.NullTime
	BelongsToRecipeStep                       string
	Name                                      string
	VesselPredicate                           string
	ID                                        string
	Notes                                     string
	RecipeStepProductID                       sql.NullString
	ValidVesselLengthInMillimeters            sql.NullString
	ValidVesselCapacity                       sql.NullString
	ValidMeasurementUnitName                  sql.NullString
	ValidVesselID                             sql.NullString
	ValidMeasurementUnitSlug                  sql.NullString
	ValidMeasurementUnitPluralName            sql.NullString
	ValidVesselName                           sql.NullString
	ValidVesselPluralName                     sql.NullString
	ValidVesselSlug                           sql.NullString
	ValidVesselWidthInMillimeters             sql.NullString
	ValidMeasurementUnitID                    sql.NullString
	ValidVesselHeightInMillimeters            sql.NullString
	ValidVesselShape                          NullVesselShape
	ValidMeasurementUnitDescription           sql.NullString
	ValidVesselIconPath                       sql.NullString
	ValidVesselDescription                    sql.NullString
	ValidMeasurementUnitIconPath              sql.NullString
	FilteredCount                             int64
	TotalCount                                int64
	MaximumQuantity                           sql.NullInt32
	MinimumQuantity                           int32
	ValidMeasurementUnitVolumetric            sql.NullBool
	ValidVesselUsableForStorage               sql.NullBool
	ValidVesselDisplayInSummaryLists          sql.NullBool
	ValidVesselIncludeInGeneratedInstructions sql.NullBool
	ValidMeasurementUnitImperial              sql.NullBool
	ValidMeasurementUnitMetric                sql.NullBool
	ValidMeasurementUnitUniversal             sql.NullBool
	UnavailableAfterStep                      bool
}

func (q *Queries) GetRecipeStepVessels(ctx context.Context, db DBTX, arg *GetRecipeStepVesselsParams) ([]*GetRecipeStepVesselsRow, error) {
	rows, err := db.QueryContext(ctx, getRecipeStepVessels,
		arg.CreatedAfter,
		arg.CreatedBefore,
		arg.UpdatedAfter,
		arg.UpdatedBefore,
		arg.RecipeStepID,
		arg.RecipeID,
		arg.QueryOffset,
		arg.QueryLimit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*GetRecipeStepVesselsRow{}
	for rows.Next() {
		var i GetRecipeStepVesselsRow
		if err := rows.Scan(
			&i.ID,
			&i.ValidVesselID,
			&i.ValidVesselName,
			&i.ValidVesselPluralName,
			&i.ValidVesselDescription,
			&i.ValidVesselIconPath,
			&i.ValidVesselUsableForStorage,
			&i.ValidVesselSlug,
			&i.ValidVesselDisplayInSummaryLists,
			&i.ValidVesselIncludeInGeneratedInstructions,
			&i.ValidVesselCapacity,
			&i.ValidMeasurementUnitID,
			&i.ValidMeasurementUnitName,
			&i.ValidMeasurementUnitDescription,
			&i.ValidMeasurementUnitVolumetric,
			&i.ValidMeasurementUnitIconPath,
			&i.ValidMeasurementUnitUniversal,
			&i.ValidMeasurementUnitMetric,
			&i.ValidMeasurementUnitImperial,
			&i.ValidMeasurementUnitSlug,
			&i.ValidMeasurementUnitPluralName,
			&i.ValidMeasurementUnitCreatedAt,
			&i.ValidMeasurementUnitLastUpdatedAt,
			&i.ValidMeasurementUnitArchivedAt,
			&i.ValidVesselWidthInMillimeters,
			&i.ValidVesselLengthInMillimeters,
			&i.ValidVesselHeightInMillimeters,
			&i.ValidVesselShape,
			&i.ValidVesselCreatedAt,
			&i.ValidVesselLastUpdatedAt,
			&i.ValidVesselArchivedAt,
			&i.Name,
			&i.Notes,
			&i.BelongsToRecipeStep,
			&i.RecipeStepProductID,
			&i.VesselPredicate,
			&i.MinimumQuantity,
			&i.MaximumQuantity,
			&i.UnavailableAfterStep,
			&i.CreatedAt,
			&i.LastUpdatedAt,
			&i.ArchivedAt,
			&i.FilteredCount,
			&i.TotalCount,
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

const getRecipeStepVesselsForRecipe = `-- name: GetRecipeStepVesselsForRecipe :many

SELECT
    recipe_step_vessels.id,
    valid_vessels.id as valid_vessel_id,
    valid_vessels.name as valid_vessel_name,
    valid_vessels.plural_name as valid_vessel_plural_name,
    valid_vessels.description as valid_vessel_description,
    valid_vessels.icon_path as valid_vessel_icon_path,
    valid_vessels.usable_for_storage as valid_vessel_usable_for_storage,
    valid_vessels.slug as valid_vessel_slug,
    valid_vessels.display_in_summary_lists as valid_vessel_display_in_summary_lists,
    valid_vessels.include_in_generated_instructions as valid_vessel_include_in_generated_instructions,
    valid_vessels.capacity as valid_vessel_capacity,
    valid_measurement_units.id as valid_measurement_unit_id,
    valid_measurement_units.name as valid_measurement_unit_name,
    valid_measurement_units.description as valid_measurement_unit_description,
    valid_measurement_units.volumetric as valid_measurement_unit_volumetric,
    valid_measurement_units.icon_path as valid_measurement_unit_icon_path,
    valid_measurement_units.universal as valid_measurement_unit_universal,
    valid_measurement_units.metric as valid_measurement_unit_metric,
    valid_measurement_units.imperial as valid_measurement_unit_imperial,
    valid_measurement_units.slug as valid_measurement_unit_slug,
    valid_measurement_units.plural_name as valid_measurement_unit_plural_name,
    valid_measurement_units.created_at as valid_measurement_unit_created_at,
    valid_measurement_units.last_updated_at as valid_measurement_unit_last_updated_at,
    valid_measurement_units.archived_at as valid_measurement_unit_archived_at,
    valid_vessels.width_in_millimeters as valid_vessel_width_in_millimeters,
    valid_vessels.length_in_millimeters as valid_vessel_length_in_millimeters,
    valid_vessels.height_in_millimeters as valid_vessel_height_in_millimeters,
    valid_vessels.shape as valid_vessel_shape,
    valid_vessels.created_at as valid_vessel_created_at,
    valid_vessels.last_updated_at as valid_vessel_last_updated_at,
    valid_vessels.archived_at as valid_vessel_archived_at,
    recipe_step_vessels.name,
    recipe_step_vessels.notes,
    recipe_step_vessels.belongs_to_recipe_step,
    recipe_step_vessels.recipe_step_product_id,
    recipe_step_vessels.vessel_predicate,
    recipe_step_vessels.minimum_quantity,
    recipe_step_vessels.maximum_quantity,
    recipe_step_vessels.unavailable_after_step,
    recipe_step_vessels.created_at,
    recipe_step_vessels.last_updated_at,
    recipe_step_vessels.archived_at
FROM recipe_step_vessels
    LEFT JOIN valid_vessels ON recipe_step_vessels.valid_vessel_id=valid_vessels.id
    LEFT JOIN valid_measurement_units ON valid_vessels.capacity_unit=valid_measurement_units.id
    JOIN recipe_steps ON recipe_step_vessels.belongs_to_recipe_step=recipe_steps.id
    JOIN recipes ON recipe_steps.belongs_to_recipe=recipes.id
WHERE recipe_step_vessels.archived_at IS NULL
    AND recipe_steps.archived_at IS NULL
    AND recipe_steps.belongs_to_recipe = $1
    AND recipes.archived_at IS NULL
    AND recipes.id = $1
`

type GetRecipeStepVesselsForRecipeRow struct {
	CreatedAt                                 time.Time
	ValidVesselLastUpdatedAt                  sql.NullTime
	LastUpdatedAt                             sql.NullTime
	ArchivedAt                                sql.NullTime
	ValidVesselArchivedAt                     sql.NullTime
	ValidMeasurementUnitCreatedAt             sql.NullTime
	ValidVesselCreatedAt                      sql.NullTime
	ValidMeasurementUnitArchivedAt            sql.NullTime
	ValidMeasurementUnitLastUpdatedAt         sql.NullTime
	VesselPredicate                           string
	BelongsToRecipeStep                       string
	Notes                                     string
	Name                                      string
	ID                                        string
	ValidVesselCapacity                       sql.NullString
	ValidVesselIconPath                       sql.NullString
	ValidVesselID                             sql.NullString
	ValidVesselName                           sql.NullString
	ValidVesselPluralName                     sql.NullString
	ValidMeasurementUnitSlug                  sql.NullString
	ValidMeasurementUnitPluralName            sql.NullString
	ValidVesselDescription                    sql.NullString
	ValidMeasurementUnitDescription           sql.NullString
	ValidMeasurementUnitName                  sql.NullString
	ValidVesselWidthInMillimeters             sql.NullString
	ValidVesselLengthInMillimeters            sql.NullString
	ValidVesselHeightInMillimeters            sql.NullString
	ValidVesselShape                          NullVesselShape
	ValidMeasurementUnitID                    sql.NullString
	RecipeStepProductID                       sql.NullString
	ValidMeasurementUnitIconPath              sql.NullString
	ValidVesselSlug                           sql.NullString
	MaximumQuantity                           sql.NullInt32
	MinimumQuantity                           int32
	ValidVesselUsableForStorage               sql.NullBool
	ValidVesselDisplayInSummaryLists          sql.NullBool
	ValidVesselIncludeInGeneratedInstructions sql.NullBool
	ValidMeasurementUnitVolumetric            sql.NullBool
	ValidMeasurementUnitImperial              sql.NullBool
	ValidMeasurementUnitMetric                sql.NullBool
	ValidMeasurementUnitUniversal             sql.NullBool
	UnavailableAfterStep                      bool
}

func (q *Queries) GetRecipeStepVesselsForRecipe(ctx context.Context, db DBTX, belongsToRecipe string) ([]*GetRecipeStepVesselsForRecipeRow, error) {
	rows, err := db.QueryContext(ctx, getRecipeStepVesselsForRecipe, belongsToRecipe)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*GetRecipeStepVesselsForRecipeRow{}
	for rows.Next() {
		var i GetRecipeStepVesselsForRecipeRow
		if err := rows.Scan(
			&i.ID,
			&i.ValidVesselID,
			&i.ValidVesselName,
			&i.ValidVesselPluralName,
			&i.ValidVesselDescription,
			&i.ValidVesselIconPath,
			&i.ValidVesselUsableForStorage,
			&i.ValidVesselSlug,
			&i.ValidVesselDisplayInSummaryLists,
			&i.ValidVesselIncludeInGeneratedInstructions,
			&i.ValidVesselCapacity,
			&i.ValidMeasurementUnitID,
			&i.ValidMeasurementUnitName,
			&i.ValidMeasurementUnitDescription,
			&i.ValidMeasurementUnitVolumetric,
			&i.ValidMeasurementUnitIconPath,
			&i.ValidMeasurementUnitUniversal,
			&i.ValidMeasurementUnitMetric,
			&i.ValidMeasurementUnitImperial,
			&i.ValidMeasurementUnitSlug,
			&i.ValidMeasurementUnitPluralName,
			&i.ValidMeasurementUnitCreatedAt,
			&i.ValidMeasurementUnitLastUpdatedAt,
			&i.ValidMeasurementUnitArchivedAt,
			&i.ValidVesselWidthInMillimeters,
			&i.ValidVesselLengthInMillimeters,
			&i.ValidVesselHeightInMillimeters,
			&i.ValidVesselShape,
			&i.ValidVesselCreatedAt,
			&i.ValidVesselLastUpdatedAt,
			&i.ValidVesselArchivedAt,
			&i.Name,
			&i.Notes,
			&i.BelongsToRecipeStep,
			&i.RecipeStepProductID,
			&i.VesselPredicate,
			&i.MinimumQuantity,
			&i.MaximumQuantity,
			&i.UnavailableAfterStep,
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

const updateRecipeStepVessel = `-- name: UpdateRecipeStepVessel :execrows

UPDATE recipe_step_vessels SET
	name = $1,
	notes = $2,
	belongs_to_recipe_step = $3,
	recipe_step_product_id = $4,
	valid_vessel_id = $5,
	vessel_predicate = $6,
	minimum_quantity = $7,
    maximum_quantity = $8,
    unavailable_after_step = $9,
	last_updated_at = NOW()
WHERE archived_at IS NULL
	AND belongs_to_recipe_step = $3
	AND id = $10
`

type UpdateRecipeStepVesselParams struct {
	Name                 string
	Notes                string
	RecipeStepID         string
	VesselPredicate      string
	ID                   string
	RecipeStepProductID  sql.NullString
	ValidVesselID        sql.NullString
	MaximumQuantity      sql.NullInt32
	MinimumQuantity      int32
	UnavailableAfterStep bool
}

func (q *Queries) UpdateRecipeStepVessel(ctx context.Context, db DBTX, arg *UpdateRecipeStepVesselParams) (int64, error) {
	result, err := db.ExecContext(ctx, updateRecipeStepVessel,
		arg.Name,
		arg.Notes,
		arg.RecipeStepID,
		arg.RecipeStepProductID,
		arg.ValidVesselID,
		arg.VesselPredicate,
		arg.MinimumQuantity,
		arg.MaximumQuantity,
		arg.UnavailableAfterStep,
		arg.ID,
	)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
