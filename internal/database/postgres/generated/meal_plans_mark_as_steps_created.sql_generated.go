// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: meal_plans_mark_as_steps_created.sql

package generated

import (
	"context"
)

const MarkMealPlanTasksAsCreated = `-- name: MarkMealPlanTasksAsCreated :exec
UPDATE meal_plans
SET
    tasks_created = 'true',
	last_updated_at = NOW()
WHERE archived_at IS NULL
	AND id = $1
`

func (q *Queries) MarkMealPlanTasksAsCreated(ctx context.Context, id string) error {
	_, err := q.exec(ctx, q.markMealPlanTasksAsCreatedStmt, MarkMealPlanTasksAsCreated, id)
	return err
}
