// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: recipe_step_ingredients_archive.sql

package generated

import (
	"context"
)

const ArchiveRecipeStepIngredient = `-- name: ArchiveRecipeStepIngredient :exec
UPDATE recipe_step_ingredients SET archived_at = NOW() WHERE archived_at IS NULL AND belongs_to_recipe_step = $1 AND id = $2
`

type ArchiveRecipeStepIngredientParams struct {
	BelongsToRecipeStep string `db:"belongs_to_recipe_step"`
	ID                  string `db:"id"`
}

func (q *Queries) ArchiveRecipeStepIngredient(ctx context.Context, db DBTX, arg *ArchiveRecipeStepIngredientParams) error {
	_, err := db.ExecContext(ctx, ArchiveRecipeStepIngredient, arg.BelongsToRecipeStep, arg.ID)
	return err
}
