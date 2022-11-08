// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: valid_preparation_instruments_archive.sql

package generated

import (
	"context"
)

const ArchiveValidPreparationInstrument = `-- name: ArchiveValidPreparationInstrument :exec
UPDATE valid_preparation_instruments SET archived_at = NOW() WHERE archived_at IS NULL AND id = $1
`

func (q *Queries) ArchiveValidPreparationInstrument(ctx context.Context, id string) error {
	_, err := q.exec(ctx, q.archiveValidPreparationInstrumentStmt, ArchiveValidPreparationInstrument, id)
	return err
}
