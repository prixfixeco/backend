// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: recipes_update.sql

package generated

import (
	"context"
	"database/sql"
)

const UpdateRecipe = `-- name: UpdateRecipe :exec
UPDATE recipes SET name = $1, source = $2, description = $3, inspired_by_recipe_id = $4, yields_portions = $5, seal_of_approval = $6, last_updated_at = NOW() WHERE archived_at IS NULL AND created_by_user = $7 AND id = $8
`

type UpdateRecipeParams struct {
	Name               string         `db:"name"`
	Source             string         `db:"source"`
	Description        string         `db:"description"`
	CreatedByUser      string         `db:"created_by_user"`
	ID                 string         `db:"id"`
	InspiredByRecipeID sql.NullString `db:"inspired_by_recipe_id"`
	YieldsPortions     int32          `db:"yields_portions"`
	SealOfApproval     bool           `db:"seal_of_approval"`
}

func (q *Queries) UpdateRecipe(ctx context.Context, arg *UpdateRecipeParams) error {
	_, err := q.exec(ctx, q.updateRecipeStmt, UpdateRecipe,
		arg.Name,
		arg.Source,
		arg.Description,
		arg.InspiredByRecipeID,
		arg.YieldsPortions,
		arg.SealOfApproval,
		arg.CreatedByUser,
		arg.ID,
	)
	return err
}
