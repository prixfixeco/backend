// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: meal_plan_options_get_for_meal_plan_event.sql

package generated

import (
	"context"
	"database/sql"
	"time"
)

const GetMealPlanOptionsForMealPlanEvent = `-- name: GetMealPlanOptionsForMealPlanEvent :many
SELECT
	meal_plan_options.id,
	meal_plan_options.assigned_cook,
	meal_plan_options.assigned_dishwasher,
	meal_plan_options.chosen,
	meal_plan_options.tiebroken,
	meal_plan_options.meal_id,
	meal_plan_options.notes,
	meal_plan_options.created_at,
	meal_plan_options.last_updated_at,
	meal_plan_options.archived_at,
	meal_plan_options.belongs_to_meal_plan_event,
	meals.id,
	meals.name,
	meals.description,
	meals.created_at,
	meals.last_updated_at,
	meals.archived_at,
	meals.created_by_user
FROM meal_plan_options
	JOIN meal_plan_events ON meal_plan_options.belongs_to_meal_plan_event = meal_plan_events.id
	JOIN meal_plans ON meal_plan_events.belongs_to_meal_plan = meal_plans.id
	JOIN meals ON meal_plan_options.meal_id = meals.id
WHERE
	meal_plan_options.archived_at IS NULL
	AND meal_plan_options.belongs_to_meal_plan_event = $1
	AND meal_plan_events.id = $1
	AND meal_plan_events.belongs_to_meal_plan = $2
	AND meal_plans.archived_at IS NULL
	AND meal_plans.id = $2
`

type GetMealPlanOptionsForMealPlanEventParams struct {
	BelongsToMealPlan      string         `db:"belongs_to_meal_plan"`
	BelongsToMealPlanEvent sql.NullString `db:"belongs_to_meal_plan_event"`
}

type GetMealPlanOptionsForMealPlanEventRow struct {
	CreatedAt              time.Time      `db:"created_at"`
	CreatedAt_2            time.Time      `db:"created_at_2"`
	LastUpdatedAt_2        sql.NullTime   `db:"last_updated_at_2"`
	ArchivedAt_2           sql.NullTime   `db:"archived_at_2"`
	LastUpdatedAt          sql.NullTime   `db:"last_updated_at"`
	ArchivedAt             sql.NullTime   `db:"archived_at"`
	ID_2                   string         `db:"id_2"`
	Description            string         `db:"description"`
	MealID                 string         `db:"meal_id"`
	Notes                  string         `db:"notes"`
	Name                   string         `db:"name"`
	ID                     string         `db:"id"`
	CreatedByUser          string         `db:"created_by_user"`
	BelongsToMealPlanEvent sql.NullString `db:"belongs_to_meal_plan_event"`
	AssignedDishwasher     sql.NullString `db:"assigned_dishwasher"`
	AssignedCook           sql.NullString `db:"assigned_cook"`
	Tiebroken              bool           `db:"tiebroken"`
	Chosen                 bool           `db:"chosen"`
}

func (q *Queries) GetMealPlanOptionsForMealPlanEvent(ctx context.Context, arg *GetMealPlanOptionsForMealPlanEventParams) ([]*GetMealPlanOptionsForMealPlanEventRow, error) {
	rows, err := q.query(ctx, q.getMealPlanOptionsForMealPlanEventStmt, GetMealPlanOptionsForMealPlanEvent, arg.BelongsToMealPlanEvent, arg.BelongsToMealPlan)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*GetMealPlanOptionsForMealPlanEventRow{}
	for rows.Next() {
		var i GetMealPlanOptionsForMealPlanEventRow
		if err := rows.Scan(
			&i.ID,
			&i.AssignedCook,
			&i.AssignedDishwasher,
			&i.Chosen,
			&i.Tiebroken,
			&i.MealID,
			&i.Notes,
			&i.CreatedAt,
			&i.LastUpdatedAt,
			&i.ArchivedAt,
			&i.BelongsToMealPlanEvent,
			&i.ID_2,
			&i.Name,
			&i.Description,
			&i.CreatedAt_2,
			&i.LastUpdatedAt_2,
			&i.ArchivedAt_2,
			&i.CreatedByUser,
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
