// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: meal_plan_tasks_create.sql

package generated

import (
	"context"
	"database/sql"
)

const CreateMealPlanTask = `-- name: CreateMealPlanTask :exec
INSERT INTO meal_plan_tasks (id,status,status_explanation,creation_explanation,belongs_to_meal_plan_option,belongs_to_recipe_prep_task,assigned_to_user)
VALUES ($1,$2,$3,$4,$5,$6,$7)
`

type CreateMealPlanTaskParams struct {
	ID                      string         `db:"id"`
	Status                  PrepStepStatus `db:"status"`
	StatusExplanation       string         `db:"status_explanation"`
	CreationExplanation     string         `db:"creation_explanation"`
	BelongsToMealPlanOption string         `db:"belongs_to_meal_plan_option"`
	BelongsToRecipePrepTask string         `db:"belongs_to_recipe_prep_task"`
	AssignedToUser          sql.NullString `db:"assigned_to_user"`
}

func (q *Queries) CreateMealPlanTask(ctx context.Context, arg *CreateMealPlanTaskParams) error {
	_, err := q.exec(ctx, q.createMealPlanTaskStmt, CreateMealPlanTask,
		arg.ID,
		arg.Status,
		arg.StatusExplanation,
		arg.CreationExplanation,
		arg.BelongsToMealPlanOption,
		arg.BelongsToRecipePrepTask,
		arg.AssignedToUser,
	)
	return err
}
