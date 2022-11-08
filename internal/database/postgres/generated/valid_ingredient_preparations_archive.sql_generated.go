// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: valid_ingredient_preparations_archive.sql

package generated

import (
	"context"
)

const ArchiveValidIngredientPreparation = `-- name: ArchiveValidIngredientPreparation :exec
UPDATE valid_ingredient_preparations SET archived_at = NOW() WHERE archived_at IS NULL AND id = $1
`

func (q *Queries) ArchiveValidIngredientPreparation(ctx context.Context, id string) error {
	_, err := q.exec(ctx, q.archiveValidIngredientPreparationStmt, ArchiveValidIngredientPreparation, id)
	return err
}
