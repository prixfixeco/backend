// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: meal_plans_get_finalized_without_grocery_list_init.sql

package generated

import (
	"context"
	"database/sql"
	"time"
)

const GetFinalizedMealPlansWithoutInitializedGroceryLists = `-- name: GetFinalizedMealPlansWithoutInitializedGroceryLists :many
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
	AND meal_plans.grocery_list_initialized = 'finalized'
    AND meal_plans.grocery_list_initialized IS FALSE
GROUP BY meal_plans.id
ORDER BY meal_plans.id
`

type GetFinalizedMealPlansWithoutInitializedGroceryListsRow struct {
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

func (q *Queries) GetFinalizedMealPlansWithoutInitializedGroceryLists(ctx context.Context, db DBTX) ([]*GetFinalizedMealPlansWithoutInitializedGroceryListsRow, error) {
	rows, err := db.QueryContext(ctx, GetFinalizedMealPlansWithoutInitializedGroceryLists)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*GetFinalizedMealPlansWithoutInitializedGroceryListsRow{}
	for rows.Next() {
		var i GetFinalizedMealPlansWithoutInitializedGroceryListsRow
		if err := rows.Scan(
			&i.ID,
			&i.Notes,
			&i.Status,
			&i.VotingDeadline,
			&i.GroceryListInitialized,
			&i.TasksCreated,
			&i.CreatedAt,
			&i.LastUpdatedAt,
			&i.ArchivedAt,
			&i.BelongsToHousehold,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
