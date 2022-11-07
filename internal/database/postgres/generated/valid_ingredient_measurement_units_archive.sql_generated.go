// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: valid_ingredient_measurement_units_archive.sql

package generated

import (
	"context"
)

const ArchiveValidIngredientMeasurementUnit = `-- name: ArchiveValidIngredientMeasurementUnit :exec
UPDATE valid_ingredient_measurement_units SET archived_at = NOW() WHERE archived_at IS NULL AND id = $1
`

func (q *Queries) ArchiveValidIngredientMeasurementUnit(ctx context.Context, db DBTX, id string) error {
	_, err := db.ExecContext(ctx, ArchiveValidIngredientMeasurementUnit, id)
	return err
}
