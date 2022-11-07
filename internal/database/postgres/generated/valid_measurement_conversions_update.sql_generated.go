// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: valid_measurement_conversions_update.sql

package generated

import (
	"context"
	"database/sql"
)

const UpdateValidMeasurementConversion = `-- name: UpdateValidMeasurementConversion :exec
UPDATE valid_measurement_conversions
SET
    from_unit = $1,
    to_unit = $2,
    only_for_ingredient = $3,
    modifier = $4,
    notes = $5,
	last_updated_at = NOW()
WHERE archived_at IS NULL
	AND id = $6
`

type UpdateValidMeasurementConversionParams struct {
	FromUnit          string         `db:"from_unit"`
	ToUnit            string         `db:"to_unit"`
	Modifier          string         `db:"modifier"`
	Notes             string         `db:"notes"`
	ID                string         `db:"id"`
	OnlyForIngredient sql.NullString `db:"only_for_ingredient"`
}

func (q *Queries) UpdateValidMeasurementConversion(ctx context.Context, db DBTX, arg *UpdateValidMeasurementConversionParams) error {
	_, err := db.ExecContext(ctx, UpdateValidMeasurementConversion,
		arg.FromUnit,
		arg.ToUnit,
		arg.OnlyForIngredient,
		arg.Modifier,
		arg.Notes,
		arg.ID,
	)
	return err
}
