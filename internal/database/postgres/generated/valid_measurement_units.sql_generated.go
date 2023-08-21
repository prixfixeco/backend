// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: valid_measurement_units.sql

package generated

import (
	"context"
	"database/sql"
	"time"

	"github.com/lib/pq"
)

const archiveValidMeasurementUnit = `-- name: ArchiveValidMeasurementUnit :exec

UPDATE valid_measurement_units SET archived_at = NOW() WHERE archived_at IS NULL AND id = $1
`

func (q *Queries) ArchiveValidMeasurementUnit(ctx context.Context, db DBTX, id string) error {
	_, err := db.ExecContext(ctx, archiveValidMeasurementUnit, id)
	return err
}

const checkValidMeasurementUnitExistence = `-- name: CheckValidMeasurementUnitExistence :one

SELECT EXISTS ( SELECT valid_measurement_units.id FROM valid_measurement_units WHERE valid_measurement_units.archived_at IS NULL AND valid_measurement_units.id = $1 )
`

func (q *Queries) CheckValidMeasurementUnitExistence(ctx context.Context, db DBTX, id string) (bool, error) {
	row := db.QueryRowContext(ctx, checkValidMeasurementUnitExistence, id)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const createValidMeasurementUnit = `-- name: CreateValidMeasurementUnit :exec

INSERT INTO valid_measurement_units
(id,"name",description,volumetric,icon_path,universal,metric,imperial,plural_name,slug)
VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
`

type CreateValidMeasurementUnitParams struct {
	ID          string
	Name        string
	Description string
	IconPath    string
	PluralName  string
	Slug        string
	Volumetric  sql.NullBool
	Universal   bool
	Metric      bool
	Imperial    bool
}

func (q *Queries) CreateValidMeasurementUnit(ctx context.Context, db DBTX, arg *CreateValidMeasurementUnitParams) error {
	_, err := db.ExecContext(ctx, createValidMeasurementUnit,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.Volumetric,
		arg.IconPath,
		arg.Universal,
		arg.Metric,
		arg.Imperial,
		arg.PluralName,
		arg.Slug,
	)
	return err
}

const getRandomValidMeasurementUnit = `-- name: GetRandomValidMeasurementUnit :one

SELECT
	valid_measurement_units.id,
	valid_measurement_units.name,
	valid_measurement_units.description,
	valid_measurement_units.volumetric,
	valid_measurement_units.icon_path,
	valid_measurement_units.universal,
	valid_measurement_units.metric,
	valid_measurement_units.imperial,
	valid_measurement_units.slug,
	valid_measurement_units.plural_name,
	valid_measurement_units.created_at,
	valid_measurement_units.last_updated_at,
	valid_measurement_units.archived_at
FROM valid_measurement_units
WHERE valid_measurement_units.archived_at IS NULL
	ORDER BY random() LIMIT 1
`

type GetRandomValidMeasurementUnitRow struct {
	CreatedAt     time.Time
	ArchivedAt    sql.NullTime
	LastUpdatedAt sql.NullTime
	PluralName    string
	Name          string
	Description   string
	ID            string
	IconPath      string
	Slug          string
	Volumetric    sql.NullBool
	Imperial      bool
	Metric        bool
	Universal     bool
}

func (q *Queries) GetRandomValidMeasurementUnit(ctx context.Context, db DBTX) (*GetRandomValidMeasurementUnitRow, error) {
	row := db.QueryRowContext(ctx, getRandomValidMeasurementUnit)
	var i GetRandomValidMeasurementUnitRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Volumetric,
		&i.IconPath,
		&i.Universal,
		&i.Metric,
		&i.Imperial,
		&i.Slug,
		&i.PluralName,
		&i.CreatedAt,
		&i.LastUpdatedAt,
		&i.ArchivedAt,
	)
	return &i, err
}

const getValidMeasurementUnit = `-- name: GetValidMeasurementUnit :one

SELECT
	valid_measurement_units.id,
	valid_measurement_units.name,
	valid_measurement_units.description,
	valid_measurement_units.volumetric,
	valid_measurement_units.icon_path,
	valid_measurement_units.universal,
	valid_measurement_units.metric,
	valid_measurement_units.imperial,
	valid_measurement_units.slug,
	valid_measurement_units.plural_name,
	valid_measurement_units.created_at,
	valid_measurement_units.last_updated_at,
	valid_measurement_units.archived_at
FROM valid_measurement_units
WHERE valid_measurement_units.archived_at IS NULL
	AND valid_measurement_units.id = $1
`

type GetValidMeasurementUnitRow struct {
	CreatedAt     time.Time
	ArchivedAt    sql.NullTime
	LastUpdatedAt sql.NullTime
	PluralName    string
	Name          string
	Description   string
	ID            string
	IconPath      string
	Slug          string
	Volumetric    sql.NullBool
	Imperial      bool
	Metric        bool
	Universal     bool
}

func (q *Queries) GetValidMeasurementUnit(ctx context.Context, db DBTX, id string) (*GetValidMeasurementUnitRow, error) {
	row := db.QueryRowContext(ctx, getValidMeasurementUnit, id)
	var i GetValidMeasurementUnitRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Volumetric,
		&i.IconPath,
		&i.Universal,
		&i.Metric,
		&i.Imperial,
		&i.Slug,
		&i.PluralName,
		&i.CreatedAt,
		&i.LastUpdatedAt,
		&i.ArchivedAt,
	)
	return &i, err
}

const getValidMeasurementUnitByID = `-- name: GetValidMeasurementUnitByID :one

SELECT valid_measurement_units.id,
       valid_measurement_units.name,
       valid_measurement_units.description,
       valid_measurement_units.volumetric,
       valid_measurement_units.icon_path,
       valid_measurement_units.universal,
       valid_measurement_units.metric,
       valid_measurement_units.imperial,
       valid_measurement_units.slug,
       valid_measurement_units.plural_name,
       valid_measurement_units.created_at,
       valid_measurement_units.last_updated_at,
       valid_measurement_units.archived_at
  FROM valid_measurement_units
 WHERE valid_measurement_units.archived_at IS NULL
`

type GetValidMeasurementUnitByIDRow struct {
	CreatedAt     time.Time
	ArchivedAt    sql.NullTime
	LastUpdatedAt sql.NullTime
	PluralName    string
	Name          string
	Description   string
	ID            string
	IconPath      string
	Slug          string
	Volumetric    sql.NullBool
	Imperial      bool
	Metric        bool
	Universal     bool
}

func (q *Queries) GetValidMeasurementUnitByID(ctx context.Context, db DBTX) (*GetValidMeasurementUnitByIDRow, error) {
	row := db.QueryRowContext(ctx, getValidMeasurementUnitByID)
	var i GetValidMeasurementUnitByIDRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Volumetric,
		&i.IconPath,
		&i.Universal,
		&i.Metric,
		&i.Imperial,
		&i.Slug,
		&i.PluralName,
		&i.CreatedAt,
		&i.LastUpdatedAt,
		&i.ArchivedAt,
	)
	return &i, err
}

const getValidMeasurementUnits = `-- name: GetValidMeasurementUnits :many

SELECT
    valid_measurement_units.id,
    valid_measurement_units.name,
    valid_measurement_units.description,
    valid_measurement_units.volumetric,
    valid_measurement_units.icon_path,
    valid_measurement_units.universal,
    valid_measurement_units.metric,
    valid_measurement_units.imperial,
    valid_measurement_units.slug,
    valid_measurement_units.plural_name,
    valid_measurement_units.created_at,
    valid_measurement_units.last_updated_at,
    valid_measurement_units.archived_at,
    (
        SELECT
            COUNT(valid_measurement_units.id)
        FROM
            valid_measurement_units
        WHERE
            valid_measurement_units.archived_at IS NULL
            AND valid_measurement_units.created_at > COALESCE($1, (SELECT NOW() - interval '999 years'))
            AND valid_measurement_units.created_at < COALESCE($2, (SELECT NOW() + interval '999 years'))
            AND (
                valid_measurement_units.last_updated_at IS NULL
                OR valid_measurement_units.last_updated_at > COALESCE($3, (SELECT NOW() - interval '999 years'))
            )
            AND (
                valid_measurement_units.last_updated_at IS NULL
                OR valid_measurement_units.last_updated_at < COALESCE($4, (SELECT NOW() + interval '999 years'))
            )
    ) as filtered_count,
    (
        SELECT
            COUNT(valid_measurement_units.id)
        FROM
            valid_measurement_units
        WHERE
            valid_measurement_units.archived_at IS NULL
    ) as total_count
FROM
    valid_measurement_units
WHERE
    valid_measurement_units.archived_at IS NULL
    AND valid_measurement_units.created_at > COALESCE($1, (SELECT NOW() - interval '999 years'))
    AND valid_measurement_units.created_at < COALESCE($2, (SELECT NOW() + interval '999 years'))
    AND (
        valid_measurement_units.last_updated_at IS NULL
        OR valid_measurement_units.last_updated_at > COALESCE($3, (SELECT NOW() - interval '999 years'))
    )
    AND (
        valid_measurement_units.last_updated_at IS NULL
        OR valid_measurement_units.last_updated_at < COALESCE($4, (SELECT NOW() + interval '999 years'))
    )
GROUP BY
    valid_measurement_units.id
ORDER BY
    valid_measurement_units.id
OFFSET $5
LIMIT $6
`

type GetValidMeasurementUnitsParams struct {
	CreatedBefore sql.NullTime
	CreatedAfter  sql.NullTime
	UpdatedBefore sql.NullTime
	UpdatedAfter  sql.NullTime
	QueryOffset   sql.NullInt32
	QueryLimit    sql.NullInt32
}

type GetValidMeasurementUnitsRow struct {
	CreatedAt     time.Time
	LastUpdatedAt sql.NullTime
	ArchivedAt    sql.NullTime
	Name          string
	Description   string
	ID            string
	IconPath      string
	Slug          string
	PluralName    string
	TotalCount    int64
	FilteredCount int64
	Volumetric    sql.NullBool
	Imperial      bool
	Metric        bool
	Universal     bool
}

func (q *Queries) GetValidMeasurementUnits(ctx context.Context, db DBTX, arg *GetValidMeasurementUnitsParams) ([]*GetValidMeasurementUnitsRow, error) {
	rows, err := db.QueryContext(ctx, getValidMeasurementUnits,
		arg.CreatedBefore,
		arg.CreatedAfter,
		arg.UpdatedBefore,
		arg.UpdatedAfter,
		arg.QueryOffset,
		arg.QueryLimit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*GetValidMeasurementUnitsRow{}
	for rows.Next() {
		var i GetValidMeasurementUnitsRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Volumetric,
			&i.IconPath,
			&i.Universal,
			&i.Metric,
			&i.Imperial,
			&i.Slug,
			&i.PluralName,
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

const getValidMeasurementUnitsNeedingIndexing = `-- name: GetValidMeasurementUnitsNeedingIndexing :many

SELECT valid_measurement_units.id
  FROM valid_measurement_units
 WHERE (valid_measurement_units.archived_at IS NULL)
       AND (
			(
				valid_measurement_units.last_indexed_at IS NULL
			)
			OR valid_measurement_units.last_indexed_at
				< now() - '24 hours'::INTERVAL
		)
`

func (q *Queries) GetValidMeasurementUnitsNeedingIndexing(ctx context.Context, db DBTX) ([]string, error) {
	rows, err := db.QueryContext(ctx, getValidMeasurementUnitsNeedingIndexing)
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

const getValidMeasurementUnitsWithIDs = `-- name: GetValidMeasurementUnitsWithIDs :many

SELECT
	valid_measurement_units.id,
	valid_measurement_units.name,
	valid_measurement_units.description,
	valid_measurement_units.volumetric,
	valid_measurement_units.icon_path,
	valid_measurement_units.universal,
	valid_measurement_units.metric,
	valid_measurement_units.imperial,
	valid_measurement_units.slug,
	valid_measurement_units.plural_name,
	valid_measurement_units.created_at,
	valid_measurement_units.last_updated_at,
	valid_measurement_units.archived_at
FROM valid_measurement_units
WHERE valid_measurement_units.archived_at IS NULL
	AND valid_measurement_units.id = ANY($1::text[])
`

type GetValidMeasurementUnitsWithIDsRow struct {
	CreatedAt     time.Time
	ArchivedAt    sql.NullTime
	LastUpdatedAt sql.NullTime
	PluralName    string
	Name          string
	Description   string
	ID            string
	IconPath      string
	Slug          string
	Volumetric    sql.NullBool
	Imperial      bool
	Metric        bool
	Universal     bool
}

func (q *Queries) GetValidMeasurementUnitsWithIDs(ctx context.Context, db DBTX, ids []string) ([]*GetValidMeasurementUnitsWithIDsRow, error) {
	rows, err := db.QueryContext(ctx, getValidMeasurementUnitsWithIDs, pq.Array(ids))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*GetValidMeasurementUnitsWithIDsRow{}
	for rows.Next() {
		var i GetValidMeasurementUnitsWithIDsRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Volumetric,
			&i.IconPath,
			&i.Universal,
			&i.Metric,
			&i.Imperial,
			&i.Slug,
			&i.PluralName,
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

const searchForValidMeasurementUnits = `-- name: SearchForValidMeasurementUnits :many

SELECT
	valid_measurement_units.id,
	valid_measurement_units.name,
	valid_measurement_units.description,
	valid_measurement_units.volumetric,
	valid_measurement_units.icon_path,
	valid_measurement_units.universal,
	valid_measurement_units.metric,
	valid_measurement_units.imperial,
	valid_measurement_units.slug,
	valid_measurement_units.plural_name,
	valid_measurement_units.created_at,
	valid_measurement_units.last_updated_at,
	valid_measurement_units.archived_at
FROM valid_measurement_units
WHERE (valid_measurement_units.name ILIKE '%' || $1::text || '%' OR valid_measurement_units.universal is TRUE)
AND valid_measurement_units.archived_at IS NULL
LIMIT 50
`

type SearchForValidMeasurementUnitsRow struct {
	CreatedAt     time.Time
	ArchivedAt    sql.NullTime
	LastUpdatedAt sql.NullTime
	PluralName    string
	Name          string
	Description   string
	ID            string
	IconPath      string
	Slug          string
	Volumetric    sql.NullBool
	Imperial      bool
	Metric        bool
	Universal     bool
}

func (q *Queries) SearchForValidMeasurementUnits(ctx context.Context, db DBTX, query string) ([]*SearchForValidMeasurementUnitsRow, error) {
	rows, err := db.QueryContext(ctx, searchForValidMeasurementUnits, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*SearchForValidMeasurementUnitsRow{}
	for rows.Next() {
		var i SearchForValidMeasurementUnitsRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Volumetric,
			&i.IconPath,
			&i.Universal,
			&i.Metric,
			&i.Imperial,
			&i.Slug,
			&i.PluralName,
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

const searchValidMeasurementUnitsByIngredientID = `-- name: SearchValidMeasurementUnitsByIngredientID :many

SELECT
	DISTINCT(valid_measurement_units.id),
	valid_measurement_units.name,
	valid_measurement_units.description,
	valid_measurement_units.volumetric,
	valid_measurement_units.icon_path,
	valid_measurement_units.universal,
	valid_measurement_units.metric,
	valid_measurement_units.imperial,
	valid_measurement_units.slug,
	valid_measurement_units.plural_name,
	valid_measurement_units.created_at,
	valid_measurement_units.last_updated_at,
	valid_measurement_units.archived_at,
	(
	SELECT
	  COUNT(valid_measurement_units.id)
	FROM
	  valid_measurement_units
	WHERE
	    valid_measurement_units.archived_at IS NULL
        AND valid_measurement_units.created_at > COALESCE($1, (SELECT NOW() - interval '999 years'))
        AND valid_measurement_units.created_at < COALESCE($2, (SELECT NOW() + interval '999 years'))
        AND (
            valid_measurement_units.last_updated_at IS NULL
            OR valid_measurement_units.last_updated_at > COALESCE($3, (SELECT NOW() - interval '999 years'))
        )
        AND (
            valid_measurement_units.last_updated_at IS NULL
            OR valid_measurement_units.last_updated_at < COALESCE($4, (SELECT NOW() + interval '999 years'))
        )
	    AND (
	        valid_ingredient_measurement_units.valid_ingredient_id = $5
	        OR valid_measurement_units.universal = true
	    )
	) as filtered_count,
	(
	    SELECT
	        COUNT(valid_measurement_units.id)
	    FROM
	        valid_measurement_units
	    WHERE
	        valid_measurement_units.archived_at IS NULL
	) as total_count
FROM valid_measurement_units
	JOIN valid_ingredient_measurement_units ON valid_ingredient_measurement_units.valid_measurement_unit_id = valid_measurement_units.id
	JOIN valid_ingredients ON valid_ingredient_measurement_units.valid_ingredient_id = valid_ingredients.id
WHERE
	(
	    valid_ingredient_measurement_units.valid_ingredient_id = $5
	    OR valid_measurement_units.universal = true
	)
    AND valid_measurement_units.archived_at IS NULL
    AND valid_ingredients.archived_at IS NULL
    AND valid_ingredient_measurement_units.archived_at IS NULL
    AND valid_measurement_units.created_at > COALESCE($1, (SELECT NOW() - interval '999 years'))
    AND valid_measurement_units.created_at < COALESCE($2, (SELECT NOW() + interval '999 years'))
    AND (
        valid_measurement_units.last_updated_at IS NULL
        OR valid_measurement_units.last_updated_at > COALESCE($3, (SELECT NOW() - interval '999 years'))
    )
    AND (
        valid_measurement_units.last_updated_at IS NULL
        OR valid_measurement_units.last_updated_at < COALESCE($4, (SELECT NOW() + interval '999 years'))
    )
	LIMIT $7
	OFFSET $6
`

type SearchValidMeasurementUnitsByIngredientIDParams struct {
	CreatedBefore     sql.NullTime
	CreatedAfter      sql.NullTime
	UpdatedBefore     sql.NullTime
	UpdatedAfter      sql.NullTime
	ValidIngredientID string
	QueryOffset       sql.NullInt32
	QueryLimit        sql.NullInt32
}

type SearchValidMeasurementUnitsByIngredientIDRow struct {
	CreatedAt     time.Time
	LastUpdatedAt sql.NullTime
	ArchivedAt    sql.NullTime
	Name          string
	Description   string
	ID            string
	IconPath      string
	Slug          string
	PluralName    string
	TotalCount    int64
	FilteredCount int64
	Volumetric    sql.NullBool
	Imperial      bool
	Metric        bool
	Universal     bool
}

func (q *Queries) SearchValidMeasurementUnitsByIngredientID(ctx context.Context, db DBTX, arg *SearchValidMeasurementUnitsByIngredientIDParams) ([]*SearchValidMeasurementUnitsByIngredientIDRow, error) {
	rows, err := db.QueryContext(ctx, searchValidMeasurementUnitsByIngredientID,
		arg.CreatedBefore,
		arg.CreatedAfter,
		arg.UpdatedBefore,
		arg.UpdatedAfter,
		arg.ValidIngredientID,
		arg.QueryOffset,
		arg.QueryLimit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*SearchValidMeasurementUnitsByIngredientIDRow{}
	for rows.Next() {
		var i SearchValidMeasurementUnitsByIngredientIDRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Volumetric,
			&i.IconPath,
			&i.Universal,
			&i.Metric,
			&i.Imperial,
			&i.Slug,
			&i.PluralName,
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

const updateValidMeasurementUnit = `-- name: UpdateValidMeasurementUnit :exec

UPDATE valid_measurement_units SET
	name = $1,
	description = $2,
	volumetric = $3,
	icon_path = $4,
	universal = $5,
	metric = $6,
	imperial = $7,
	slug = $8,
	plural_name = $9,
	last_updated_at = NOW()
WHERE archived_at IS NULL AND id = $10
`

type UpdateValidMeasurementUnitParams struct {
	Name        string
	Description string
	IconPath    string
	Slug        string
	PluralName  string
	ID          string
	Volumetric  sql.NullBool
	Universal   bool
	Metric      bool
	Imperial    bool
}

func (q *Queries) UpdateValidMeasurementUnit(ctx context.Context, db DBTX, arg *UpdateValidMeasurementUnitParams) error {
	_, err := db.ExecContext(ctx, updateValidMeasurementUnit,
		arg.Name,
		arg.Description,
		arg.Volumetric,
		arg.IconPath,
		arg.Universal,
		arg.Metric,
		arg.Imperial,
		arg.Slug,
		arg.PluralName,
		arg.ID,
	)
	return err
}

const updateValidMeasurementUnitLastIndexedAt = `-- name: UpdateValidMeasurementUnitLastIndexedAt :exec

UPDATE valid_measurement_units SET last_indexed_at = NOW() WHERE id = $1 AND archived_at IS NULL
`

func (q *Queries) UpdateValidMeasurementUnitLastIndexedAt(ctx context.Context, db DBTX, id string) error {
	_, err := db.ExecContext(ctx, updateValidMeasurementUnitLastIndexedAt, id)
	return err
}
