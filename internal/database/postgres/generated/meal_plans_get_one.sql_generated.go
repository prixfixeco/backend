// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: meal_plans_get_one.sql

package generated

import (
	"context"
	"database/sql"
	"time"
)

const GetMealPlan = `-- name: GetMealPlan :exec
SELECT
	meal_plans.id,
	meal_plans.notes,
	meal_plans.status,
    meal_plans.voting_deadline,
    meal_plans.grocery_list_initialized,
    meal_plans.tasks_created,
	meal_plans.created_at,
	meal_plans.last_updated_at,
	meal_plans.archived_at,
	meal_plans.belongs_to_household
FROM meal_plans
WHERE meal_plans.archived_at IS NULL
	AND meal_plans.id = $1
	AND meal_plans.belongs_to_household = $2
`

type GetMealPlanParams struct {
	ID                 string `db:"id"`
	BelongsToHousehold string `db:"belongs_to_household"`
}

type GetMealPlanRow struct {
	CreatedAt              time.Time      `db:"created_at"`
	VotingDeadline         time.Time      `db:"voting_deadline"`
	LastUpdatedAt          sql.NullTime   `db:"last_updated_at"`
	ArchivedAt             sql.NullTime   `db:"archived_at"`
	ID                     string         `db:"id"`
	Notes                  string         `db:"notes"`
	Status                 MealPlanStatus `db:"status"`
	BelongsToHousehold     string         `db:"belongs_to_household"`
	GroceryListInitialized bool           `db:"grocery_list_initialized"`
	TasksCreated           bool           `db:"tasks_created"`
}

func (q *Queries) GetMealPlan(ctx context.Context, arg *GetMealPlanParams) error {
	_, err := q.exec(ctx, q.getMealPlanStmt, GetMealPlan, arg.ID, arg.BelongsToHousehold)
	return err
}
