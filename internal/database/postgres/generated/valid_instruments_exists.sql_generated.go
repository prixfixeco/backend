// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: valid_instruments_exists.sql

package generated

import (
	"context"
)

const ValidInstrumentExists = `-- name: ValidInstrumentExists :exec
SELECT EXISTS ( SELECT valid_instruments.id FROM valid_instruments WHERE valid_instruments.archived_at IS NULL AND valid_instruments.id = $1 )
`

func (q *Queries) ValidInstrumentExists(ctx context.Context, db DBTX, id string) error {
	_, err := db.ExecContext(ctx, ValidInstrumentExists, id)
	return err
}
