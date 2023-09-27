// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: valid_preparations.sql

package generated

import (
	"context"
	"database/sql"
	"time"

	"github.com/lib/pq"
)

const archiveValidPreparation = `-- name: ArchiveValidPreparation :execrows

UPDATE valid_preparations SET archived_at = NOW() WHERE archived_at IS NULL AND id = $1
`

func (q *Queries) ArchiveValidPreparation(ctx context.Context, db DBTX, id string) (int64, error) {
	result, err := db.ExecContext(ctx, archiveValidPreparation, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

const checkValidPreparationExistence = `-- name: CheckValidPreparationExistence :one

SELECT EXISTS (
	SELECT valid_preparations.id
	FROM valid_preparations
	WHERE valid_preparations.archived_at IS NULL
		AND valid_preparations.id = $1
)
`

func (q *Queries) CheckValidPreparationExistence(ctx context.Context, db DBTX, id string) (bool, error) {
	row := db.QueryRowContext(ctx, checkValidPreparationExistence, id)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const createValidPreparation = `-- name: CreateValidPreparation :exec

INSERT INTO valid_preparations (
	id,
	name,
	description,
	icon_path,
	yields_nothing,
	restrict_to_ingredients,
	past_tense,
	slug,
	minimum_ingredient_count,
	maximum_ingredient_count,
	minimum_instrument_count,
	maximum_instrument_count,
	temperature_required,
	time_estimate_required,
	condition_expression_required,
	consumes_vessel,
	only_for_vessels,
	minimum_vessel_count,
	maximum_vessel_count
) VALUES (
	$1,
	$2,
	$3,
	$4,
	$5,
	$6,
	$7,
	$8,
	$9,
	$10,
	$11,
	$12,
	$13,
	$14,
	$15,
	$16,
	$17,
	$18,
	$19
)
`

type CreateValidPreparationParams struct {
	Name                        string
	Description                 string
	IconPath                    string
	ID                          string
	PastTense                   string
	Slug                        string
	MaximumVesselCount          sql.NullInt32
	MaximumIngredientCount      sql.NullInt32
	MaximumInstrumentCount      sql.NullInt32
	MinimumVesselCount          int32
	MinimumIngredientCount      int32
	MinimumInstrumentCount      int32
	YieldsNothing               bool
	TimeEstimateRequired        bool
	ConditionExpressionRequired bool
	ConsumesVessel              bool
	OnlyForVessels              bool
	TemperatureRequired         bool
	RestrictToIngredients       bool
}

func (q *Queries) CreateValidPreparation(ctx context.Context, db DBTX, arg *CreateValidPreparationParams) error {
	_, err := db.ExecContext(ctx, createValidPreparation,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.IconPath,
		arg.YieldsNothing,
		arg.RestrictToIngredients,
		arg.PastTense,
		arg.Slug,
		arg.MinimumIngredientCount,
		arg.MaximumIngredientCount,
		arg.MinimumInstrumentCount,
		arg.MaximumInstrumentCount,
		arg.TemperatureRequired,
		arg.TimeEstimateRequired,
		arg.ConditionExpressionRequired,
		arg.ConsumesVessel,
		arg.OnlyForVessels,
		arg.MinimumVesselCount,
		arg.MaximumVesselCount,
	)
	return err
}

const getRandomValidPreparation = `-- name: GetRandomValidPreparation :one

SELECT
	valid_preparations.id,
	valid_preparations.name,
	valid_preparations.description,
	valid_preparations.icon_path,
	valid_preparations.yields_nothing,
	valid_preparations.restrict_to_ingredients,
	valid_preparations.past_tense,
	valid_preparations.slug,
	valid_preparations.minimum_ingredient_count,
	valid_preparations.maximum_ingredient_count,
	valid_preparations.minimum_instrument_count,
	valid_preparations.maximum_instrument_count,
	valid_preparations.temperature_required,
	valid_preparations.time_estimate_required,
	valid_preparations.condition_expression_required,
	valid_preparations.consumes_vessel,
	valid_preparations.only_for_vessels,
	valid_preparations.minimum_vessel_count,
	valid_preparations.maximum_vessel_count,
	valid_preparations.last_indexed_at,
	valid_preparations.created_at,
	valid_preparations.last_updated_at,
	valid_preparations.archived_at
FROM valid_preparations
WHERE valid_preparations.archived_at IS NULL
ORDER BY RANDOM() LIMIT 1
`

type GetRandomValidPreparationRow struct {
	CreatedAt                   time.Time
	LastUpdatedAt               sql.NullTime
	ArchivedAt                  sql.NullTime
	LastIndexedAt               sql.NullTime
	Name                        string
	IconPath                    string
	ID                          string
	PastTense                   string
	Description                 string
	Slug                        string
	MaximumIngredientCount      sql.NullInt32
	MaximumInstrumentCount      sql.NullInt32
	MaximumVesselCount          sql.NullInt32
	MinimumVesselCount          int32
	MinimumInstrumentCount      int32
	MinimumIngredientCount      int32
	RestrictToIngredients       bool
	OnlyForVessels              bool
	ConsumesVessel              bool
	ConditionExpressionRequired bool
	TimeEstimateRequired        bool
	TemperatureRequired         bool
	YieldsNothing               bool
}

func (q *Queries) GetRandomValidPreparation(ctx context.Context, db DBTX) (*GetRandomValidPreparationRow, error) {
	row := db.QueryRowContext(ctx, getRandomValidPreparation)
	var i GetRandomValidPreparationRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.IconPath,
		&i.YieldsNothing,
		&i.RestrictToIngredients,
		&i.PastTense,
		&i.Slug,
		&i.MinimumIngredientCount,
		&i.MaximumIngredientCount,
		&i.MinimumInstrumentCount,
		&i.MaximumInstrumentCount,
		&i.TemperatureRequired,
		&i.TimeEstimateRequired,
		&i.ConditionExpressionRequired,
		&i.ConsumesVessel,
		&i.OnlyForVessels,
		&i.MinimumVesselCount,
		&i.MaximumVesselCount,
		&i.LastIndexedAt,
		&i.CreatedAt,
		&i.LastUpdatedAt,
		&i.ArchivedAt,
	)
	return &i, err
}

const getValidPreparation = `-- name: GetValidPreparation :one

SELECT
	valid_preparations.id,
	valid_preparations.name,
	valid_preparations.description,
	valid_preparations.icon_path,
	valid_preparations.yields_nothing,
	valid_preparations.restrict_to_ingredients,
	valid_preparations.past_tense,
	valid_preparations.slug,
	valid_preparations.minimum_ingredient_count,
	valid_preparations.maximum_ingredient_count,
	valid_preparations.minimum_instrument_count,
	valid_preparations.maximum_instrument_count,
	valid_preparations.temperature_required,
	valid_preparations.time_estimate_required,
	valid_preparations.condition_expression_required,
	valid_preparations.consumes_vessel,
	valid_preparations.only_for_vessels,
	valid_preparations.minimum_vessel_count,
	valid_preparations.maximum_vessel_count,
	valid_preparations.last_indexed_at,
	valid_preparations.created_at,
	valid_preparations.last_updated_at,
	valid_preparations.archived_at
FROM valid_preparations
WHERE valid_preparations.archived_at IS NULL
AND valid_preparations.id = $1
`

type GetValidPreparationRow struct {
	CreatedAt                   time.Time
	LastUpdatedAt               sql.NullTime
	ArchivedAt                  sql.NullTime
	LastIndexedAt               sql.NullTime
	Name                        string
	IconPath                    string
	ID                          string
	PastTense                   string
	Description                 string
	Slug                        string
	MaximumIngredientCount      sql.NullInt32
	MaximumInstrumentCount      sql.NullInt32
	MaximumVesselCount          sql.NullInt32
	MinimumVesselCount          int32
	MinimumInstrumentCount      int32
	MinimumIngredientCount      int32
	RestrictToIngredients       bool
	OnlyForVessels              bool
	ConsumesVessel              bool
	ConditionExpressionRequired bool
	TimeEstimateRequired        bool
	TemperatureRequired         bool
	YieldsNothing               bool
}

func (q *Queries) GetValidPreparation(ctx context.Context, db DBTX, id string) (*GetValidPreparationRow, error) {
	row := db.QueryRowContext(ctx, getValidPreparation, id)
	var i GetValidPreparationRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.IconPath,
		&i.YieldsNothing,
		&i.RestrictToIngredients,
		&i.PastTense,
		&i.Slug,
		&i.MinimumIngredientCount,
		&i.MaximumIngredientCount,
		&i.MinimumInstrumentCount,
		&i.MaximumInstrumentCount,
		&i.TemperatureRequired,
		&i.TimeEstimateRequired,
		&i.ConditionExpressionRequired,
		&i.ConsumesVessel,
		&i.OnlyForVessels,
		&i.MinimumVesselCount,
		&i.MaximumVesselCount,
		&i.LastIndexedAt,
		&i.CreatedAt,
		&i.LastUpdatedAt,
		&i.ArchivedAt,
	)
	return &i, err
}

const getValidPreparations = `-- name: GetValidPreparations :many

SELECT
	valid_preparations.id,
	valid_preparations.name,
	valid_preparations.description,
	valid_preparations.icon_path,
	valid_preparations.yields_nothing,
	valid_preparations.restrict_to_ingredients,
	valid_preparations.past_tense,
	valid_preparations.slug,
	valid_preparations.minimum_ingredient_count,
	valid_preparations.maximum_ingredient_count,
	valid_preparations.minimum_instrument_count,
	valid_preparations.maximum_instrument_count,
	valid_preparations.temperature_required,
	valid_preparations.time_estimate_required,
	valid_preparations.condition_expression_required,
	valid_preparations.consumes_vessel,
	valid_preparations.only_for_vessels,
	valid_preparations.minimum_vessel_count,
	valid_preparations.maximum_vessel_count,
	valid_preparations.last_indexed_at,
	valid_preparations.created_at,
	valid_preparations.last_updated_at,
	valid_preparations.archived_at,
	(
		SELECT COUNT(valid_preparations.id)
		FROM valid_preparations
		WHERE valid_preparations.archived_at IS NULL
			AND valid_preparations.created_at > COALESCE($1, (SELECT NOW() - '999 years'::INTERVAL))
			AND valid_preparations.created_at < COALESCE($2, (SELECT NOW() + '999 years'::INTERVAL))
			AND (
				valid_preparations.last_updated_at IS NULL
				OR valid_preparations.last_updated_at > COALESCE($3, (SELECT NOW() - '999 years'::INTERVAL))
			)
			AND (
				valid_preparations.last_updated_at IS NULL
				OR valid_preparations.last_updated_at < COALESCE($4, (SELECT NOW() + '999 years'::INTERVAL))
			)
	) AS filtered_count,
	(
		SELECT COUNT(valid_preparations.id)
		FROM valid_preparations
		WHERE valid_preparations.archived_at IS NULL
	) AS total_count
FROM valid_preparations
WHERE
	valid_preparations.archived_at IS NULL
	AND valid_preparations.created_at > COALESCE($1, (SELECT NOW() - '999 years'::INTERVAL))
	AND valid_preparations.created_at < COALESCE($2, (SELECT NOW() + '999 years'::INTERVAL))
	AND (
		valid_preparations.last_updated_at IS NULL
		OR valid_preparations.last_updated_at > COALESCE($4, (SELECT NOW() - '999 years'::INTERVAL))
	)
	AND (
		valid_preparations.last_updated_at IS NULL
		OR valid_preparations.last_updated_at < COALESCE($3, (SELECT NOW() + '999 years'::INTERVAL))
	)
GROUP BY valid_preparations.id
ORDER BY valid_preparations.id
LIMIT $6
OFFSET $5
`

type GetValidPreparationsParams struct {
	CreatedAfter  sql.NullTime
	CreatedBefore sql.NullTime
	UpdatedBefore sql.NullTime
	UpdatedAfter  sql.NullTime
	QueryOffset   sql.NullInt32
	QueryLimit    sql.NullInt32
}

type GetValidPreparationsRow struct {
	CreatedAt                   time.Time
	ArchivedAt                  sql.NullTime
	LastUpdatedAt               sql.NullTime
	LastIndexedAt               sql.NullTime
	ID                          string
	Name                        string
	Description                 string
	IconPath                    string
	PastTense                   string
	Slug                        string
	TotalCount                  int64
	FilteredCount               int64
	MaximumVesselCount          sql.NullInt32
	MaximumInstrumentCount      sql.NullInt32
	MaximumIngredientCount      sql.NullInt32
	MinimumVesselCount          int32
	MinimumInstrumentCount      int32
	MinimumIngredientCount      int32
	TimeEstimateRequired        bool
	ConditionExpressionRequired bool
	ConsumesVessel              bool
	OnlyForVessels              bool
	TemperatureRequired         bool
	RestrictToIngredients       bool
	YieldsNothing               bool
}

func (q *Queries) GetValidPreparations(ctx context.Context, db DBTX, arg *GetValidPreparationsParams) ([]*GetValidPreparationsRow, error) {
	rows, err := db.QueryContext(ctx, getValidPreparations,
		arg.CreatedAfter,
		arg.CreatedBefore,
		arg.UpdatedBefore,
		arg.UpdatedAfter,
		arg.QueryOffset,
		arg.QueryLimit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*GetValidPreparationsRow{}
	for rows.Next() {
		var i GetValidPreparationsRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.IconPath,
			&i.YieldsNothing,
			&i.RestrictToIngredients,
			&i.PastTense,
			&i.Slug,
			&i.MinimumIngredientCount,
			&i.MaximumIngredientCount,
			&i.MinimumInstrumentCount,
			&i.MaximumInstrumentCount,
			&i.TemperatureRequired,
			&i.TimeEstimateRequired,
			&i.ConditionExpressionRequired,
			&i.ConsumesVessel,
			&i.OnlyForVessels,
			&i.MinimumVesselCount,
			&i.MaximumVesselCount,
			&i.LastIndexedAt,
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

const getValidPreparationsNeedingIndexing = `-- name: GetValidPreparationsNeedingIndexing :many

SELECT valid_preparations.id
FROM valid_preparations
WHERE valid_preparations.archived_at IS NULL
	AND (
	valid_preparations.last_indexed_at IS NULL
	OR valid_preparations.last_indexed_at < NOW() - '24 hours'::INTERVAL
)
`

func (q *Queries) GetValidPreparationsNeedingIndexing(ctx context.Context, db DBTX) ([]string, error) {
	rows, err := db.QueryContext(ctx, getValidPreparationsNeedingIndexing)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []string{}
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		items = append(items, id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getValidPreparationsWithIDs = `-- name: GetValidPreparationsWithIDs :many

SELECT
	valid_preparations.id,
	valid_preparations.name,
	valid_preparations.description,
	valid_preparations.icon_path,
	valid_preparations.yields_nothing,
	valid_preparations.restrict_to_ingredients,
	valid_preparations.past_tense,
	valid_preparations.slug,
	valid_preparations.minimum_ingredient_count,
	valid_preparations.maximum_ingredient_count,
	valid_preparations.minimum_instrument_count,
	valid_preparations.maximum_instrument_count,
	valid_preparations.temperature_required,
	valid_preparations.time_estimate_required,
	valid_preparations.condition_expression_required,
	valid_preparations.consumes_vessel,
	valid_preparations.only_for_vessels,
	valid_preparations.minimum_vessel_count,
	valid_preparations.maximum_vessel_count,
	valid_preparations.last_indexed_at,
	valid_preparations.created_at,
	valid_preparations.last_updated_at,
	valid_preparations.archived_at
FROM valid_preparations
WHERE valid_preparations.archived_at IS NULL
	AND valid_preparations.id = ANY($1::text[])
`

type GetValidPreparationsWithIDsRow struct {
	CreatedAt                   time.Time
	LastUpdatedAt               sql.NullTime
	ArchivedAt                  sql.NullTime
	LastIndexedAt               sql.NullTime
	Name                        string
	IconPath                    string
	ID                          string
	PastTense                   string
	Description                 string
	Slug                        string
	MaximumIngredientCount      sql.NullInt32
	MaximumInstrumentCount      sql.NullInt32
	MaximumVesselCount          sql.NullInt32
	MinimumVesselCount          int32
	MinimumInstrumentCount      int32
	MinimumIngredientCount      int32
	RestrictToIngredients       bool
	OnlyForVessels              bool
	ConsumesVessel              bool
	ConditionExpressionRequired bool
	TimeEstimateRequired        bool
	TemperatureRequired         bool
	YieldsNothing               bool
}

func (q *Queries) GetValidPreparationsWithIDs(ctx context.Context, db DBTX, ids []string) ([]*GetValidPreparationsWithIDsRow, error) {
	rows, err := db.QueryContext(ctx, getValidPreparationsWithIDs, pq.Array(ids))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*GetValidPreparationsWithIDsRow{}
	for rows.Next() {
		var i GetValidPreparationsWithIDsRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.IconPath,
			&i.YieldsNothing,
			&i.RestrictToIngredients,
			&i.PastTense,
			&i.Slug,
			&i.MinimumIngredientCount,
			&i.MaximumIngredientCount,
			&i.MinimumInstrumentCount,
			&i.MaximumInstrumentCount,
			&i.TemperatureRequired,
			&i.TimeEstimateRequired,
			&i.ConditionExpressionRequired,
			&i.ConsumesVessel,
			&i.OnlyForVessels,
			&i.MinimumVesselCount,
			&i.MaximumVesselCount,
			&i.LastIndexedAt,
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

const searchForValidPreparations = `-- name: SearchForValidPreparations :many

SELECT
	valid_preparations.id,
	valid_preparations.name,
	valid_preparations.description,
	valid_preparations.icon_path,
	valid_preparations.yields_nothing,
	valid_preparations.restrict_to_ingredients,
	valid_preparations.past_tense,
	valid_preparations.slug,
	valid_preparations.minimum_ingredient_count,
	valid_preparations.maximum_ingredient_count,
	valid_preparations.minimum_instrument_count,
	valid_preparations.maximum_instrument_count,
	valid_preparations.temperature_required,
	valid_preparations.time_estimate_required,
	valid_preparations.condition_expression_required,
	valid_preparations.consumes_vessel,
	valid_preparations.only_for_vessels,
	valid_preparations.minimum_vessel_count,
	valid_preparations.maximum_vessel_count,
	valid_preparations.last_indexed_at,
	valid_preparations.created_at,
	valid_preparations.last_updated_at,
	valid_preparations.archived_at
FROM valid_preparations
WHERE valid_preparations.name ILIKE '%' || $1::text || '%'
	AND valid_preparations.archived_at IS NULL
LIMIT 50
`

type SearchForValidPreparationsRow struct {
	CreatedAt                   time.Time
	LastUpdatedAt               sql.NullTime
	ArchivedAt                  sql.NullTime
	LastIndexedAt               sql.NullTime
	Name                        string
	IconPath                    string
	ID                          string
	PastTense                   string
	Description                 string
	Slug                        string
	MaximumIngredientCount      sql.NullInt32
	MaximumInstrumentCount      sql.NullInt32
	MaximumVesselCount          sql.NullInt32
	MinimumVesselCount          int32
	MinimumInstrumentCount      int32
	MinimumIngredientCount      int32
	RestrictToIngredients       bool
	OnlyForVessels              bool
	ConsumesVessel              bool
	ConditionExpressionRequired bool
	TimeEstimateRequired        bool
	TemperatureRequired         bool
	YieldsNothing               bool
}

func (q *Queries) SearchForValidPreparations(ctx context.Context, db DBTX, nameQuery string) ([]*SearchForValidPreparationsRow, error) {
	rows, err := db.QueryContext(ctx, searchForValidPreparations, nameQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*SearchForValidPreparationsRow{}
	for rows.Next() {
		var i SearchForValidPreparationsRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.IconPath,
			&i.YieldsNothing,
			&i.RestrictToIngredients,
			&i.PastTense,
			&i.Slug,
			&i.MinimumIngredientCount,
			&i.MaximumIngredientCount,
			&i.MinimumInstrumentCount,
			&i.MaximumInstrumentCount,
			&i.TemperatureRequired,
			&i.TimeEstimateRequired,
			&i.ConditionExpressionRequired,
			&i.ConsumesVessel,
			&i.OnlyForVessels,
			&i.MinimumVesselCount,
			&i.MaximumVesselCount,
			&i.LastIndexedAt,
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

const updateValidPreparation = `-- name: UpdateValidPreparation :execrows

UPDATE valid_preparations SET
	name = $1,
	description = $2,
	icon_path = $3,
	yields_nothing = $4,
	restrict_to_ingredients = $5,
	past_tense = $6,
	slug = $7,
	minimum_ingredient_count = $8,
	maximum_ingredient_count = $9,
	minimum_instrument_count = $10,
	maximum_instrument_count = $11,
	temperature_required = $12,
	time_estimate_required = $13,
	condition_expression_required = $14,
	consumes_vessel = $15,
	only_for_vessels = $16,
	minimum_vessel_count = $17,
	maximum_vessel_count = $18,
	last_updated_at = NOW()
WHERE archived_at IS NULL
	AND id = $19
`

type UpdateValidPreparationParams struct {
	Description                 string
	IconPath                    string
	ID                          string
	Name                        string
	PastTense                   string
	Slug                        string
	MaximumIngredientCount      sql.NullInt32
	MaximumInstrumentCount      sql.NullInt32
	MaximumVesselCount          sql.NullInt32
	MinimumVesselCount          int32
	MinimumIngredientCount      int32
	MinimumInstrumentCount      int32
	RestrictToIngredients       bool
	ConditionExpressionRequired bool
	ConsumesVessel              bool
	OnlyForVessels              bool
	TimeEstimateRequired        bool
	TemperatureRequired         bool
	YieldsNothing               bool
}

func (q *Queries) UpdateValidPreparation(ctx context.Context, db DBTX, arg *UpdateValidPreparationParams) (int64, error) {
	result, err := db.ExecContext(ctx, updateValidPreparation,
		arg.Name,
		arg.Description,
		arg.IconPath,
		arg.YieldsNothing,
		arg.RestrictToIngredients,
		arg.PastTense,
		arg.Slug,
		arg.MinimumIngredientCount,
		arg.MaximumIngredientCount,
		arg.MinimumInstrumentCount,
		arg.MaximumInstrumentCount,
		arg.TemperatureRequired,
		arg.TimeEstimateRequired,
		arg.ConditionExpressionRequired,
		arg.ConsumesVessel,
		arg.OnlyForVessels,
		arg.MinimumVesselCount,
		arg.MaximumVesselCount,
		arg.ID,
	)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

const updateValidPreparationLastIndexedAt = `-- name: UpdateValidPreparationLastIndexedAt :execrows

UPDATE valid_preparations SET last_indexed_at = NOW() WHERE id = $1 AND archived_at IS NULL
`

func (q *Queries) UpdateValidPreparationLastIndexedAt(ctx context.Context, db DBTX, id string) (int64, error) {
	result, err := db.ExecContext(ctx, updateValidPreparationLastIndexedAt, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
