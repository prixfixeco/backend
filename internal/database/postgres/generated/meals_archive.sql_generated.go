// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: meals_archive.sql

package generated

import (
	"context"
)

const ArchiveMeal = `-- name: ArchiveMeal :exec
UPDATE meals SET archived_at = NOW() WHERE archived_at IS NULL AND created_by_user = $1 AND id = $2
`

type ArchiveMealParams struct {
	CreatedByUser string `db:"created_by_user"`
	ID            string `db:"id"`
}

func (q *Queries) ArchiveMeal(ctx context.Context, db DBTX, arg *ArchiveMealParams) error {
	_, err := db.ExecContext(ctx, ArchiveMeal, arg.CreatedByUser, arg.ID)
	return err
}
