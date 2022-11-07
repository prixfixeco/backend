// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: recipe_steps_archive.sql

package generated

import (
	"context"
)

const ArchiveRecipeStep = `-- name: ArchiveRecipeStep :exec
UPDATE recipe_steps SET archived_at = NOW() WHERE archived_at IS NULL AND belongs_to_recipe = $1 AND id = $2
`

type ArchiveRecipeStepParams struct {
	BelongsToRecipe string `db:"belongs_to_recipe"`
	ID              string `db:"id"`
}

func (q *Queries) ArchiveRecipeStep(ctx context.Context, db DBTX, arg *ArchiveRecipeStepParams) error {
	_, err := db.ExecContext(ctx, ArchiveRecipeStep, arg.BelongsToRecipe, arg.ID)
	return err
}
