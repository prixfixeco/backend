// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: meal_plans.sql

package generated

import (
	"context"
	"database/sql"
	"time"
)

const archiveMealPlan = `-- name: ArchiveMealPlan :execrows

UPDATE meal_plans SET archived_at = NOW() WHERE archived_at IS NULL AND belongs_to_household = $1 AND id = $2
`

type ArchiveMealPlanParams struct {
	BelongsToHousehold string
	ID                 string
}

func (q *Queries) ArchiveMealPlan(ctx context.Context, db DBTX, arg *ArchiveMealPlanParams) (int64, error) {
	result, err := db.ExecContext(ctx, archiveMealPlan, arg.BelongsToHousehold, arg.ID)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

const checkMealPlanExistence = `-- name: CheckMealPlanExistence :one

SELECT EXISTS ( SELECT meal_plans.id FROM meal_plans WHERE meal_plans.archived_at IS NULL AND meal_plans.id = $1 AND meal_plans.belongs_to_household = $2 )
`

type CheckMealPlanExistenceParams struct {
	MealPlanID  string
	HouseholdID string
}

func (q *Queries) CheckMealPlanExistence(ctx context.Context, db DBTX, arg *CheckMealPlanExistenceParams) (bool, error) {
	row := db.QueryRowContext(ctx, checkMealPlanExistence, arg.MealPlanID, arg.HouseholdID)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const createMealPlan = `-- name: CreateMealPlan :exec

INSERT INTO meal_plans (id,notes,status,voting_deadline,belongs_to_household,created_by_user) VALUES ($1,$2,$3,$4,$5,$6)
`

type CreateMealPlanParams struct {
	ID                 string
	Notes              string
	Status             MealPlanStatus
	VotingDeadline     time.Time
	BelongsToHousehold string
	CreatedByUser      string
}

func (q *Queries) CreateMealPlan(ctx context.Context, db DBTX, arg *CreateMealPlanParams) error {
	_, err := db.ExecContext(ctx, createMealPlan,
		arg.ID,
		arg.Notes,
		arg.Status,
		arg.VotingDeadline,
		arg.BelongsToHousehold,
		arg.CreatedByUser,
	)
	return err
}

const finalizeMealPlan = `-- name: FinalizeMealPlan :exec

UPDATE meal_plans SET status = $1 WHERE archived_at IS NULL AND id = $2
`

type FinalizeMealPlanParams struct {
	Status MealPlanStatus
	ID     string
}

func (q *Queries) FinalizeMealPlan(ctx context.Context, db DBTX, arg *FinalizeMealPlanParams) error {
	_, err := db.ExecContext(ctx, finalizeMealPlan, arg.Status, arg.ID)
	return err
}

const getExpiredAndUnresolvedMealPlans = `-- name: GetExpiredAndUnresolvedMealPlans :many

SELECT
    meal_plans.id,
    meal_plans.notes,
    meal_plans.status,
    meal_plans.voting_deadline,
    meal_plans.grocery_list_initialized,
    meal_plans.tasks_created,
    meal_plans.election_method,
    meal_plans.created_at,
    meal_plans.last_updated_at,
    meal_plans.archived_at,
    meal_plans.belongs_to_household,
    meal_plans.created_by_user
FROM meal_plans
WHERE meal_plans.archived_at IS NULL
	AND meal_plans.status = 'awaiting_votes'
	AND voting_deadline < now()
GROUP BY meal_plans.id
ORDER BY meal_plans.id
`

type GetExpiredAndUnresolvedMealPlansRow struct {
	VotingDeadline         time.Time
	CreatedAt              time.Time
	LastUpdatedAt          sql.NullTime
	ArchivedAt             sql.NullTime
	ID                     string
	Notes                  string
	Status                 MealPlanStatus
	ElectionMethod         ValidElectionMethod
	BelongsToHousehold     string
	CreatedByUser          string
	GroceryListInitialized bool
	TasksCreated           bool
}

func (q *Queries) GetExpiredAndUnresolvedMealPlans(ctx context.Context, db DBTX) ([]*GetExpiredAndUnresolvedMealPlansRow, error) {
	rows, err := db.QueryContext(ctx, getExpiredAndUnresolvedMealPlans)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*GetExpiredAndUnresolvedMealPlansRow{}
	for rows.Next() {
		var i GetExpiredAndUnresolvedMealPlansRow
		if err := rows.Scan(
			&i.ID,
			&i.Notes,
			&i.Status,
			&i.VotingDeadline,
			&i.GroceryListInitialized,
			&i.TasksCreated,
			&i.ElectionMethod,
			&i.CreatedAt,
			&i.LastUpdatedAt,
			&i.ArchivedAt,
			&i.BelongsToHousehold,
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

const getFinalizedMealPlansForPlanning = `-- name: GetFinalizedMealPlansForPlanning :many

SELECT
    meal_plans.id as meal_plan_id,
    meal_plan_options.id as meal_plan_option_id,
    meals.id as meal_id,
    meal_plan_events.id as meal_plan_event_id,
    meal_components.recipe_id as recipe_id
FROM
    meal_plan_options
    JOIN meal_plan_events ON meal_plan_options.belongs_to_meal_plan_event = meal_plan_events.id
    JOIN meal_plans ON meal_plan_events.belongs_to_meal_plan = meal_plans.id
    JOIN meal_components ON meal_plan_options.meal_id = meal_components.meal_id
    JOIN meals ON meal_plan_options.meal_id = meals.id
WHERE
    meal_plans.archived_at IS NULL
    AND meal_plans.status = 'finalized'
    AND meal_plan_options.chosen IS TRUE
    AND meal_plans.tasks_created IS FALSE
GROUP BY
    meal_plans.id,
    meal_plan_options.id,
    meals.id,
    meal_plan_events.id,
    meal_components.recipe_id
ORDER BY
    meal_plans.id
`

type GetFinalizedMealPlansForPlanningRow struct {
	MealPlanID       string
	MealPlanOptionID string
	MealID           string
	MealPlanEventID  string
	RecipeID         string
}

func (q *Queries) GetFinalizedMealPlansForPlanning(ctx context.Context, db DBTX) ([]*GetFinalizedMealPlansForPlanningRow, error) {
	rows, err := db.QueryContext(ctx, getFinalizedMealPlansForPlanning)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*GetFinalizedMealPlansForPlanningRow{}
	for rows.Next() {
		var i GetFinalizedMealPlansForPlanningRow
		if err := rows.Scan(
			&i.MealPlanID,
			&i.MealPlanOptionID,
			&i.MealID,
			&i.MealPlanEventID,
			&i.RecipeID,
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

const getFinalizedMealPlansWithoutGroceryListInit = `-- name: GetFinalizedMealPlansWithoutGroceryListInit :many

SELECT
	meal_plans.id,
	meal_plans.belongs_to_household
FROM meal_plans
WHERE meal_plans.archived_at IS NULL
	AND meal_plans.status = 'finalized'
	AND meal_plans.grocery_list_initialized IS FALSE
`

type GetFinalizedMealPlansWithoutGroceryListInitRow struct {
	ID                 string
	BelongsToHousehold string
}

func (q *Queries) GetFinalizedMealPlansWithoutGroceryListInit(ctx context.Context, db DBTX) ([]*GetFinalizedMealPlansWithoutGroceryListInitRow, error) {
	rows, err := db.QueryContext(ctx, getFinalizedMealPlansWithoutGroceryListInit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*GetFinalizedMealPlansWithoutGroceryListInitRow{}
	for rows.Next() {
		var i GetFinalizedMealPlansWithoutGroceryListInitRow
		if err := rows.Scan(&i.ID, &i.BelongsToHousehold); err != nil {
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

const getMealPlan = `-- name: GetMealPlan :one

SELECT
    meal_plans.id,
    meal_plans.notes,
    meal_plans.status,
    meal_plans.voting_deadline,
    meal_plans.grocery_list_initialized,
    meal_plans.tasks_created,
    meal_plans.election_method,
    meal_plans.created_at,
    meal_plans.last_updated_at,
    meal_plans.archived_at,
    meal_plans.belongs_to_household,
    meal_plans.created_by_user
FROM meal_plans
WHERE meal_plans.archived_at IS NULL
  AND meal_plans.id = $1
  AND meal_plans.belongs_to_household = $2
`

type GetMealPlanParams struct {
	ID                 string
	BelongsToHousehold string
}

type GetMealPlanRow struct {
	VotingDeadline         time.Time
	CreatedAt              time.Time
	LastUpdatedAt          sql.NullTime
	ArchivedAt             sql.NullTime
	ID                     string
	Notes                  string
	Status                 MealPlanStatus
	ElectionMethod         ValidElectionMethod
	BelongsToHousehold     string
	CreatedByUser          string
	GroceryListInitialized bool
	TasksCreated           bool
}

func (q *Queries) GetMealPlan(ctx context.Context, db DBTX, arg *GetMealPlanParams) (*GetMealPlanRow, error) {
	row := db.QueryRowContext(ctx, getMealPlan, arg.ID, arg.BelongsToHousehold)
	var i GetMealPlanRow
	err := row.Scan(
		&i.ID,
		&i.Notes,
		&i.Status,
		&i.VotingDeadline,
		&i.GroceryListInitialized,
		&i.TasksCreated,
		&i.ElectionMethod,
		&i.CreatedAt,
		&i.LastUpdatedAt,
		&i.ArchivedAt,
		&i.BelongsToHousehold,
		&i.CreatedByUser,
	)
	return &i, err
}

const getMealPlanPastVotingDeadline = `-- name: GetMealPlanPastVotingDeadline :one

SELECT
	meal_plans.id,
	meal_plans.notes,
	meal_plans.status,
	meal_plans.voting_deadline,
	meal_plans.grocery_list_initialized,
	meal_plans.tasks_created,
	meal_plans.election_method,
	meal_plans.created_at,
	meal_plans.last_updated_at,
	meal_plans.archived_at,
	meal_plans.belongs_to_household,
	meal_plans.created_by_user
FROM meal_plans
WHERE meal_plans.archived_at IS NULL
	AND meal_plans.id = $1
	AND meal_plans.belongs_to_household = $2
	AND meal_plans.status = 'awaiting_votes'
	AND NOW() > meal_plans.voting_deadline
`

type GetMealPlanPastVotingDeadlineParams struct {
	MealPlanID  string
	HouseholdID string
}

type GetMealPlanPastVotingDeadlineRow struct {
	VotingDeadline         time.Time
	CreatedAt              time.Time
	LastUpdatedAt          sql.NullTime
	ArchivedAt             sql.NullTime
	ID                     string
	Notes                  string
	Status                 MealPlanStatus
	ElectionMethod         ValidElectionMethod
	BelongsToHousehold     string
	CreatedByUser          string
	GroceryListInitialized bool
	TasksCreated           bool
}

func (q *Queries) GetMealPlanPastVotingDeadline(ctx context.Context, db DBTX, arg *GetMealPlanPastVotingDeadlineParams) (*GetMealPlanPastVotingDeadlineRow, error) {
	row := db.QueryRowContext(ctx, getMealPlanPastVotingDeadline, arg.MealPlanID, arg.HouseholdID)
	var i GetMealPlanPastVotingDeadlineRow
	err := row.Scan(
		&i.ID,
		&i.Notes,
		&i.Status,
		&i.VotingDeadline,
		&i.GroceryListInitialized,
		&i.TasksCreated,
		&i.ElectionMethod,
		&i.CreatedAt,
		&i.LastUpdatedAt,
		&i.ArchivedAt,
		&i.BelongsToHousehold,
		&i.CreatedByUser,
	)
	return &i, err
}

const getMealPlans = `-- name: GetMealPlans :many

SELECT
    meal_plans.id,
    meal_plans.notes,
    meal_plans.status,
    meal_plans.voting_deadline,
    meal_plans.grocery_list_initialized,
    meal_plans.tasks_created,
    meal_plans.election_method,
    meal_plans.created_at,
    meal_plans.last_updated_at,
    meal_plans.archived_at,
    meal_plans.belongs_to_household,
    meal_plans.created_by_user,
    (
        SELECT
            COUNT(meal_plans.id)
        FROM
            meal_plans
        WHERE
            meal_plans.archived_at IS NULL
            AND meal_plans.belongs_to_household = $1
            AND meal_plans.belongs_to_household = $1
            AND meal_plans.created_at > COALESCE($2, (SELECT NOW() - interval '999 years'))
            AND meal_plans.created_at < COALESCE($3, (SELECT NOW() + interval '999 years'))
            AND (
                meal_plans.last_updated_at IS NULL
                OR meal_plans.last_updated_at > COALESCE($4, (SELECT NOW() - interval '999 years'))
            )
            AND (
                meal_plans.last_updated_at IS NULL
                OR meal_plans.last_updated_at < COALESCE($5, (SELECT NOW() + interval '999 years'))
            )
    ) as filtered_count,
    (
        SELECT
            COUNT(meal_plans.id)
        FROM
            meal_plans
        WHERE
            meal_plans.archived_at IS NULL
            AND meal_plans.belongs_to_household = $1
    ) as total_count
FROM meal_plans
WHERE meal_plans.archived_at IS NULL
    AND meal_plans.belongs_to_household = $1
    AND meal_plans.belongs_to_household = $1
    AND meal_plans.created_at > COALESCE($2, (SELECT NOW() - interval '999 years'))
    AND meal_plans.created_at < COALESCE($3, (SELECT NOW() + interval '999 years'))
    AND (
        meal_plans.last_updated_at IS NULL
        OR meal_plans.last_updated_at > COALESCE($4, (SELECT NOW() - interval '999 years'))
    )
    AND (
        meal_plans.last_updated_at IS NULL
        OR meal_plans.last_updated_at < COALESCE($5, (SELECT NOW() + interval '999 years'))
    )
    OFFSET $6
    LIMIT $7
`

type GetMealPlansParams struct {
	HouseholdID   string
	CreatedAfter  sql.NullTime
	CreatedBefore sql.NullTime
	UpdatedAfter  sql.NullTime
	UpdatedBefore sql.NullTime
	QueryOffset   sql.NullInt32
	QueryLimit    sql.NullInt32
}

type GetMealPlansRow struct {
	VotingDeadline         time.Time
	CreatedAt              time.Time
	ArchivedAt             sql.NullTime
	LastUpdatedAt          sql.NullTime
	Status                 MealPlanStatus
	ElectionMethod         ValidElectionMethod
	Notes                  string
	ID                     string
	BelongsToHousehold     string
	CreatedByUser          string
	FilteredCount          int64
	TotalCount             int64
	GroceryListInitialized bool
	TasksCreated           bool
}

func (q *Queries) GetMealPlans(ctx context.Context, db DBTX, arg *GetMealPlansParams) ([]*GetMealPlansRow, error) {
	rows, err := db.QueryContext(ctx, getMealPlans,
		arg.HouseholdID,
		arg.CreatedAfter,
		arg.CreatedBefore,
		arg.UpdatedAfter,
		arg.UpdatedBefore,
		arg.QueryOffset,
		arg.QueryLimit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*GetMealPlansRow{}
	for rows.Next() {
		var i GetMealPlansRow
		if err := rows.Scan(
			&i.ID,
			&i.Notes,
			&i.Status,
			&i.VotingDeadline,
			&i.GroceryListInitialized,
			&i.TasksCreated,
			&i.ElectionMethod,
			&i.CreatedAt,
			&i.LastUpdatedAt,
			&i.ArchivedAt,
			&i.BelongsToHousehold,
			&i.CreatedByUser,
			&i.FilteredCount,
			&i.TotalCount,
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

const markMealPlanAsGroceryListInitialized = `-- name: MarkMealPlanAsGroceryListInitialized :exec

UPDATE meal_plans
SET
	grocery_list_initialized = 'true',
	last_updated_at = NOW()
WHERE archived_at IS NULL
	AND id = $1
`

func (q *Queries) MarkMealPlanAsGroceryListInitialized(ctx context.Context, db DBTX, id string) error {
	_, err := db.ExecContext(ctx, markMealPlanAsGroceryListInitialized, id)
	return err
}

const markMealPlanAsPrepTasksCreated = `-- name: MarkMealPlanAsPrepTasksCreated :exec

UPDATE meal_plans
SET
	tasks_created = 'true',
	last_updated_at = NOW()
WHERE archived_at IS NULL
	AND id = $1
`

func (q *Queries) MarkMealPlanAsPrepTasksCreated(ctx context.Context, db DBTX, id string) error {
	_, err := db.ExecContext(ctx, markMealPlanAsPrepTasksCreated, id)
	return err
}

const updateMealPlan = `-- name: UpdateMealPlan :execrows

UPDATE meal_plans SET notes = $1, status = $2, voting_deadline = $3, last_updated_at = NOW() WHERE archived_at IS NULL AND belongs_to_household = $4 AND id = $5
`

type UpdateMealPlanParams struct {
	Notes              string
	Status             MealPlanStatus
	VotingDeadline     time.Time
	BelongsToHousehold string
	ID                 string
}

func (q *Queries) UpdateMealPlan(ctx context.Context, db DBTX, arg *UpdateMealPlanParams) (int64, error) {
	result, err := db.ExecContext(ctx, updateMealPlan,
		arg.Notes,
		arg.Status,
		arg.VotingDeadline,
		arg.BelongsToHousehold,
		arg.ID,
	)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
