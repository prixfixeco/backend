// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: household_instrument_ownerships.sql

package generated

import (
	"context"
	"database/sql"
	"time"
)

const archiveHouseholdInstrumentOwnership = `-- name: ArchiveHouseholdInstrumentOwnership :exec

UPDATE household_instrument_ownerships SET archived_at = NOW() WHERE archived_at IS NULL AND id = $1 AND belongs_to_household = $2
`

type ArchiveHouseholdInstrumentOwnershipParams struct {
	ID                 string
	BelongsToHousehold string
}

func (q *Queries) ArchiveHouseholdInstrumentOwnership(ctx context.Context, db DBTX, arg *ArchiveHouseholdInstrumentOwnershipParams) error {
	_, err := db.ExecContext(ctx, archiveHouseholdInstrumentOwnership, arg.ID, arg.BelongsToHousehold)
	return err
}

const checkHouseholdInstrumentOwnershipExistence = `-- name: CheckHouseholdInstrumentOwnershipExistence :one

SELECT EXISTS ( SELECT household_instrument_ownerships.id FROM household_instrument_ownerships WHERE household_instrument_ownerships.archived_at IS NULL AND household_instrument_ownerships.id = $1 AND household_instrument_ownerships.belongs_to_household = $2 )
`

type CheckHouseholdInstrumentOwnershipExistenceParams struct {
	ID                 string
	BelongsToHousehold string
}

func (q *Queries) CheckHouseholdInstrumentOwnershipExistence(ctx context.Context, db DBTX, arg *CheckHouseholdInstrumentOwnershipExistenceParams) (bool, error) {
	row := db.QueryRowContext(ctx, checkHouseholdInstrumentOwnershipExistence, arg.ID, arg.BelongsToHousehold)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const createHouseholdInstrumentOwnership = `-- name: CreateHouseholdInstrumentOwnership :exec

INSERT INTO household_instrument_ownerships (id,notes,quantity,valid_instrument_id,belongs_to_household) VALUES ($1,$2,$3,$4,$5)
`

type CreateHouseholdInstrumentOwnershipParams struct {
	ID                 string
	Notes              string
	ValidInstrumentID  string
	BelongsToHousehold string
	Quantity           int32
}

func (q *Queries) CreateHouseholdInstrumentOwnership(ctx context.Context, db DBTX, arg *CreateHouseholdInstrumentOwnershipParams) error {
	_, err := db.ExecContext(ctx, createHouseholdInstrumentOwnership,
		arg.ID,
		arg.Notes,
		arg.Quantity,
		arg.ValidInstrumentID,
		arg.BelongsToHousehold,
	)
	return err
}

const getHouseholdInstrumentOwnership = `-- name: GetHouseholdInstrumentOwnership :one

SELECT
	household_instrument_ownerships.id,
	household_instrument_ownerships.notes,
	household_instrument_ownerships.quantity,
	valid_instruments.id as valid_instrument_id,
	valid_instruments.name as valid_instrument_name,
	valid_instruments.plural_name as valid_instrument_plural_name,
	valid_instruments.description as valid_instrument_description,
	valid_instruments.icon_path as valid_instrument_icon_path,
	valid_instruments.usable_for_storage as valid_instrument_usable_for_storage,
	valid_instruments.display_in_summary_lists as valid_instrument_display_in_summary_lists,
	valid_instruments.include_in_generated_instructions as valid_instrument_include_in_generated_instructions,
	valid_instruments.slug as valid_instrument_slug,
	valid_instruments.created_at as valid_instrument_created_at,
	valid_instruments.last_updated_at as valid_instrument_last_updated_at,
	valid_instruments.archived_at as valid_instrument_archived_at,
	household_instrument_ownerships.belongs_to_household,
	household_instrument_ownerships.created_at,
	household_instrument_ownerships.last_updated_at,
	household_instrument_ownerships.archived_at
FROM household_instrument_ownerships
INNER JOIN valid_instruments ON household_instrument_ownerships.valid_instrument_id = valid_instruments.id
WHERE household_instrument_ownerships.archived_at IS NULL
	AND household_instrument_ownerships.id = $1
	AND household_instrument_ownerships.belongs_to_household = $2
`

type GetHouseholdInstrumentOwnershipParams struct {
	ID                 string
	BelongsToHousehold string
}

type GetHouseholdInstrumentOwnershipRow struct {
	ValidInstrumentCreatedAt                      time.Time
	CreatedAt                                     time.Time
	ArchivedAt                                    sql.NullTime
	LastUpdatedAt                                 sql.NullTime
	ValidInstrumentArchivedAt                     sql.NullTime
	ValidInstrumentLastUpdatedAt                  sql.NullTime
	ValidInstrumentName                           string
	ValidInstrumentIconPath                       string
	ValidInstrumentSlug                           string
	ValidInstrumentDescription                    string
	ValidInstrumentPluralName                     string
	ID                                            string
	BelongsToHousehold                            string
	ValidInstrumentID                             string
	Notes                                         string
	Quantity                                      int32
	ValidInstrumentUsableForStorage               bool
	ValidInstrumentDisplayInSummaryLists          bool
	ValidInstrumentIncludeInGeneratedInstructions bool
}

func (q *Queries) GetHouseholdInstrumentOwnership(ctx context.Context, db DBTX, arg *GetHouseholdInstrumentOwnershipParams) (*GetHouseholdInstrumentOwnershipRow, error) {
	row := db.QueryRowContext(ctx, getHouseholdInstrumentOwnership, arg.ID, arg.BelongsToHousehold)
	var i GetHouseholdInstrumentOwnershipRow
	err := row.Scan(
		&i.ID,
		&i.Notes,
		&i.Quantity,
		&i.ValidInstrumentID,
		&i.ValidInstrumentName,
		&i.ValidInstrumentPluralName,
		&i.ValidInstrumentDescription,
		&i.ValidInstrumentIconPath,
		&i.ValidInstrumentUsableForStorage,
		&i.ValidInstrumentDisplayInSummaryLists,
		&i.ValidInstrumentIncludeInGeneratedInstructions,
		&i.ValidInstrumentSlug,
		&i.ValidInstrumentCreatedAt,
		&i.ValidInstrumentLastUpdatedAt,
		&i.ValidInstrumentArchivedAt,
		&i.BelongsToHousehold,
		&i.CreatedAt,
		&i.LastUpdatedAt,
		&i.ArchivedAt,
	)
	return &i, err
}

const getHouseholdInstrumentOwnerships = `-- name: GetHouseholdInstrumentOwnerships :many

SELECT
  household_instrument_ownerships.id,
  household_instrument_ownerships.notes,
  household_instrument_ownerships.quantity,
  valid_instruments.id,
  valid_instruments.name,
  valid_instruments.plural_name,
  valid_instruments.description,
  valid_instruments.icon_path,
  valid_instruments.usable_for_storage,
  valid_instruments.display_in_summary_lists,
  valid_instruments.include_in_generated_instructions,
  valid_instruments.slug,
  valid_instruments.created_at,
  valid_instruments.last_updated_at,
  valid_instruments.archived_at,
  household_instrument_ownerships.belongs_to_household,
  household_instrument_ownerships.created_at,
  household_instrument_ownerships.last_updated_at,
  household_instrument_ownerships.archived_at,
  (
    SELECT
      COUNT(household_instrument_ownerships.id)
    FROM
      household_instrument_ownerships
    WHERE
      household_instrument_ownerships.belongs_to_household = $1
      AND household_instrument_ownerships.archived_at IS NULL
      AND household_instrument_ownerships.created_at > COALESCE($2, (SELECT NOW() - interval '999 years'))
      AND household_instrument_ownerships.created_at < COALESCE($3, (SELECT NOW() + interval '999 years'))
      AND (
        household_instrument_ownerships.last_updated_at IS NULL
        OR household_instrument_ownerships.last_updated_at > COALESCE($4, (SELECT NOW() - interval '999 years'))
      )
      AND (
        household_instrument_ownerships.last_updated_at IS NULL
        OR household_instrument_ownerships.last_updated_at < COALESCE($5, (SELECT NOW() + interval '999 years'))
      )
  ) as filtered_count,
  (
    SELECT
      COUNT(household_instrument_ownerships.id)
    FROM
      household_instrument_ownerships
    WHERE
      household_instrument_ownerships.belongs_to_household = $1
      AND household_instrument_ownerships.archived_at IS NULL
  ) as total_count
FROM
  household_instrument_ownerships
  JOIN valid_instruments ON valid_instruments.id = household_instrument_ownerships.valid_instrument_id
WHERE household_instrument_ownerships.belongs_to_household = $1
  AND household_instrument_ownerships.archived_at IS NULL
  AND household_instrument_ownerships.created_at > COALESCE($2, (SELECT NOW() - interval '999 years'))
  AND household_instrument_ownerships.created_at < COALESCE($3, (SELECT NOW() + interval '999 years'))
  AND (
    household_instrument_ownerships.last_updated_at IS NULL
    OR household_instrument_ownerships.last_updated_at > COALESCE($4, (SELECT NOW() - interval '999 years'))
  )
  AND (
    household_instrument_ownerships.last_updated_at IS NULL
    OR household_instrument_ownerships.last_updated_at < COALESCE($5, (SELECT NOW() + interval '999 years'))
  )
GROUP BY
  household_instrument_ownerships.id,
  valid_instruments.id
ORDER BY
  household_instrument_ownerships.id
OFFSET $6
LIMIT $7
`

type GetHouseholdInstrumentOwnershipsParams struct {
	BelongsToHousehold string
	CreatedAt          time.Time
	CreatedAt_2        time.Time
	LastUpdatedAt      sql.NullTime
	LastUpdatedAt_2    sql.NullTime
	Offset             int32
	Limit              int32
}

type GetHouseholdInstrumentOwnershipsRow struct {
	CreatedAt                      time.Time
	CreatedAt_2                    time.Time
	ArchivedAt_2                   sql.NullTime
	LastUpdatedAt_2                sql.NullTime
	ArchivedAt                     sql.NullTime
	LastUpdatedAt                  sql.NullTime
	BelongsToHousehold             string
	ID_2                           string
	Notes                          string
	IconPath                       string
	ID                             string
	Slug                           string
	Description                    string
	PluralName                     string
	Name                           string
	FilteredCount                  int64
	TotalCount                     int64
	Quantity                       int32
	IncludeInGeneratedInstructions bool
	DisplayInSummaryLists          bool
	UsableForStorage               bool
}

func (q *Queries) GetHouseholdInstrumentOwnerships(ctx context.Context, db DBTX, arg *GetHouseholdInstrumentOwnershipsParams) ([]*GetHouseholdInstrumentOwnershipsRow, error) {
	rows, err := db.QueryContext(ctx, getHouseholdInstrumentOwnerships,
		arg.BelongsToHousehold,
		arg.CreatedAt,
		arg.CreatedAt_2,
		arg.LastUpdatedAt,
		arg.LastUpdatedAt_2,
		arg.Offset,
		arg.Limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*GetHouseholdInstrumentOwnershipsRow{}
	for rows.Next() {
		var i GetHouseholdInstrumentOwnershipsRow
		if err := rows.Scan(
			&i.ID,
			&i.Notes,
			&i.Quantity,
			&i.ID_2,
			&i.Name,
			&i.PluralName,
			&i.Description,
			&i.IconPath,
			&i.UsableForStorage,
			&i.DisplayInSummaryLists,
			&i.IncludeInGeneratedInstructions,
			&i.Slug,
			&i.CreatedAt,
			&i.LastUpdatedAt,
			&i.ArchivedAt,
			&i.BelongsToHousehold,
			&i.CreatedAt_2,
			&i.LastUpdatedAt_2,
			&i.ArchivedAt_2,
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

const updateHouseholdInstrumentOwnership = `-- name: UpdateHouseholdInstrumentOwnership :exec

UPDATE household_instrument_ownerships
SET
	notes = $1,
	quantity = $2,
	valid_instrument_id = $3,
	last_updated_at = NOW()
WHERE archived_at IS NULL
	AND id = $4
	AND household_instrument_ownerships.belongs_to_household = $5
`

type UpdateHouseholdInstrumentOwnershipParams struct {
	Notes              string
	ValidInstrumentID  string
	ID                 string
	BelongsToHousehold string
	Quantity           int32
}

func (q *Queries) UpdateHouseholdInstrumentOwnership(ctx context.Context, db DBTX, arg *UpdateHouseholdInstrumentOwnershipParams) error {
	_, err := db.ExecContext(ctx, updateHouseholdInstrumentOwnership,
		arg.Notes,
		arg.Quantity,
		arg.ValidInstrumentID,
		arg.ID,
		arg.BelongsToHousehold,
	)
	return err
}
