// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: meal_plan_options_create.sql

package generated

import (
	"context"
	"database/sql"
)

const CreateMealPlanOption = `-- name: CreateMealPlanOption :exec
INSERT INTO meal_plan_options (id,assigned_cook,assigned_dishwasher,meal_id,notes,belongs_to_meal_plan_event,chosen) VALUES ($1,$2,$3,$4,$5,$6,$7)
`

type CreateMealPlanOptionParams struct {
	ID                     string         `db:"id"`
	AssignedCook           sql.NullString `db:"assigned_cook"`
	AssignedDishwasher     sql.NullString `db:"assigned_dishwasher"`
	MealID                 string         `db:"meal_id"`
	Notes                  string         `db:"notes"`
	BelongsToMealPlanEvent sql.NullString `db:"belongs_to_meal_plan_event"`
	Chosen                 bool           `db:"chosen"`
}

func (q *Queries) CreateMealPlanOption(ctx context.Context, db DBTX, arg *CreateMealPlanOptionParams) error {
	_, err := db.ExecContext(ctx, CreateMealPlanOption,
		arg.ID,
		arg.AssignedCook,
		arg.AssignedDishwasher,
		arg.MealID,
		arg.Notes,
		arg.BelongsToMealPlanEvent,
		arg.Chosen,
	)
	return err
}
