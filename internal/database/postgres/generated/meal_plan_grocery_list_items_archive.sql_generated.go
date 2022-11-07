// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: meal_plan_grocery_list_items_archive.sql

package generated

import (
	"context"
)

const ArchiveMealPlanGroceryListItem = `-- name: ArchiveMealPlanGroceryListItem :exec
UPDATE meal_plan_grocery_list_items SET completed_at = NOW() WHERE completed_at IS NULL AND id = $1
`

func (q *Queries) ArchiveMealPlanGroceryListItem(ctx context.Context, db DBTX, id string) error {
	_, err := db.ExecContext(ctx, ArchiveMealPlanGroceryListItem, id)
	return err
}
