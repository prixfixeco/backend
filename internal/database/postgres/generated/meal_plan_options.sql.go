// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: meal_plan_options.sql

package generated

import (
	"context"
	"database/sql"
)

const ArchiveMealPlanOption = `-- name: ArchiveMealPlanOption :exec
UPDATE meal_plan_options SET archived_on = extract(epoch FROM NOW()) WHERE archived_on IS NULL AND belongs_to_meal_plan = $1 AND id = $2
`

type ArchiveMealPlanOptionParams struct {
	BelongsToMealPlan string
	ID                string
}

func (q *Queries) ArchiveMealPlanOption(ctx context.Context, arg *ArchiveMealPlanOptionParams) error {
	_, err := q.db.ExecContext(ctx, ArchiveMealPlanOption, arg.BelongsToMealPlan, arg.ID)
	return err
}

const FinalizeMealPlanOption = `-- name: FinalizeMealPlanOption :exec
UPDATE meal_plan_options SET chosen = (belongs_to_meal_plan = $1 AND id = $2), tiebroken = $3 WHERE archived_on IS NULL AND belongs_to_meal_plan = $1 AND id = $2
`

type FinalizeMealPlanOptionParams struct {
	BelongsToMealPlan string
	ID                string
	Tiebroken         bool
}

func (q *Queries) FinalizeMealPlanOption(ctx context.Context, arg *FinalizeMealPlanOptionParams) error {
	_, err := q.db.ExecContext(ctx, FinalizeMealPlanOption, arg.BelongsToMealPlan, arg.ID, arg.Tiebroken)
	return err
}

const GetMealPlanOptionQuery = `-- name: GetMealPlanOptionQuery :one
SELECT
	meal_plan_options.id,
	meal_plan_options.day,
	meal_plan_options.meal_name,
	meal_plan_options.chosen,
	meal_plan_options.tiebroken,
	meal_plan_options.meal_id,
	meal_plan_options.notes,
	meal_plan_options.created_on,
	meal_plan_options.last_updated_on,
	meal_plan_options.archived_on,
	meal_plan_options.belongs_to_meal_plan,
	meals.id,
	meals.name,
	meals.description,
	meals.created_on,
	meals.last_updated_on,
	meals.archived_on,
	meals.created_by_user
FROM meal_plan_options
JOIN meal_plans ON meal_plan_options.belongs_to_meal_plan=meal_plans.id
JOIN meals ON meal_plan_options.meal_id=meals.id
WHERE meal_plan_options.archived_on IS NULL
AND meal_plan_options.belongs_to_meal_plan = $1
AND meal_plan_options.id = $2
AND meal_plans.archived_on IS NULL
AND meal_plans.id = $3
`

type GetMealPlanOptionQueryParams struct {
	BelongsToMealPlan string
	ID                string
	ID_2              string
}

type GetMealPlanOptionQueryRow struct {
	ID                string
	Day               int32
	MealName          MealName
	Chosen            bool
	Tiebroken         bool
	MealID            string
	Notes             string
	CreatedOn         int64
	LastUpdatedOn     sql.NullInt64
	ArchivedOn        sql.NullInt64
	BelongsToMealPlan string
	ID_2              string
	Name              string
	Description       string
	CreatedOn_2       int64
	LastUpdatedOn_2   sql.NullInt64
	ArchivedOn_2      sql.NullInt64
	CreatedByUser     string
}

func (q *Queries) GetMealPlanOptionQuery(ctx context.Context, arg *GetMealPlanOptionQueryParams) (*GetMealPlanOptionQueryRow, error) {
	row := q.db.QueryRowContext(ctx, GetMealPlanOptionQuery, arg.BelongsToMealPlan, arg.ID, arg.ID_2)
	var i GetMealPlanOptionQueryRow
	err := row.Scan(
		&i.ID,
		&i.Day,
		&i.MealName,
		&i.Chosen,
		&i.Tiebroken,
		&i.MealID,
		&i.Notes,
		&i.CreatedOn,
		&i.LastUpdatedOn,
		&i.ArchivedOn,
		&i.BelongsToMealPlan,
		&i.ID_2,
		&i.Name,
		&i.Description,
		&i.CreatedOn_2,
		&i.LastUpdatedOn_2,
		&i.ArchivedOn_2,
		&i.CreatedByUser,
	)
	return &i, err
}

const GetTotalMealPlanOptionsCount = `-- name: GetTotalMealPlanOptionsCount :one
SELECT COUNT(meal_plan_options.id) FROM meal_plan_options WHERE meal_plan_options.archived_on IS NULL
`

func (q *Queries) GetTotalMealPlanOptionsCount(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, GetTotalMealPlanOptionsCount)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const MealPlanOptionCreation = `-- name: MealPlanOptionCreation :exec
INSERT INTO meal_plan_options (id,day,meal_name,meal_id,notes,belongs_to_meal_plan) VALUES ($1,$2,$3,$4,$5,$6)
`

type MealPlanOptionCreationParams struct {
	ID                string
	Day               int32
	MealName          MealName
	MealID            string
	Notes             string
	BelongsToMealPlan string
}

func (q *Queries) MealPlanOptionCreation(ctx context.Context, arg *MealPlanOptionCreationParams) error {
	_, err := q.db.ExecContext(ctx, MealPlanOptionCreation,
		arg.ID,
		arg.Day,
		arg.MealName,
		arg.MealID,
		arg.Notes,
		arg.BelongsToMealPlan,
	)
	return err
}

const MealPlanOptionExists = `-- name: MealPlanOptionExists :one
SELECT EXISTS ( SELECT meal_plan_options.id FROM meal_plan_options JOIN meal_plans ON meal_plan_options.belongs_to_meal_plan=meal_plans.id WHERE meal_plan_options.archived_on IS NULL AND meal_plan_options.belongs_to_meal_plan = $1 AND meal_plan_options.id = $2 AND meal_plans.archived_on IS NULL AND meal_plans.id = $3 )
`

type MealPlanOptionExistsParams struct {
	BelongsToMealPlan string
	ID                string
	ID_2              string
}

func (q *Queries) MealPlanOptionExists(ctx context.Context, arg *MealPlanOptionExistsParams) (bool, error) {
	row := q.db.QueryRowContext(ctx, MealPlanOptionExists, arg.BelongsToMealPlan, arg.ID, arg.ID_2)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const UpdateMealPlanOption = `-- name: UpdateMealPlanOption :exec
UPDATE meal_plan_options SET day = $1, meal_id = $2, meal_name = $3, notes = $4, last_updated_on = extract(epoch FROM NOW()) WHERE archived_on IS NULL AND belongs_to_meal_plan = $5 AND id = $6
`

type UpdateMealPlanOptionParams struct {
	Day               int32
	MealID            string
	MealName          MealName
	Notes             string
	BelongsToMealPlan string
	ID                string
}

func (q *Queries) UpdateMealPlanOption(ctx context.Context, arg *UpdateMealPlanOptionParams) error {
	_, err := q.db.ExecContext(ctx, UpdateMealPlanOption,
		arg.Day,
		arg.MealID,
		arg.MealName,
		arg.Notes,
		arg.BelongsToMealPlan,
		arg.ID,
	)
	return err
}
