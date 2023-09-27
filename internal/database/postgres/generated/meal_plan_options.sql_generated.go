// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: meal_plan_options.sql

package generated

import (
	"context"
	"database/sql"
	"time"
)

const archiveMealPlanOption = `-- name: ArchiveMealPlanOption :execrows

UPDATE meal_plan_options SET
	archived_at = NOW()
WHERE archived_at IS NULL
	AND belongs_to_meal_plan_event = $1
	AND id = $2
`

type ArchiveMealPlanOptionParams struct {
	ID                     string
	BelongsToMealPlanEvent sql.NullString
}

func (q *Queries) ArchiveMealPlanOption(ctx context.Context, db DBTX, arg *ArchiveMealPlanOptionParams) (int64, error) {
	result, err := db.ExecContext(ctx, archiveMealPlanOption, arg.BelongsToMealPlanEvent, arg.ID)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

const checkMealPlanOptionExistence = `-- name: CheckMealPlanOptionExistence :one

SELECT EXISTS (
	SELECT meal_plan_options.id
	FROM meal_plan_options
		JOIN meal_plan_events ON meal_plan_options.belongs_to_meal_plan_event = meal_plan_events.id
		JOIN meal_plans ON meal_plan_events.belongs_to_meal_plan = meal_plans.id
	WHERE meal_plan_options.archived_at IS NULL
		AND meal_plan_options.belongs_to_meal_plan_event = $1
		AND meal_plan_options.id = $2
		AND meal_plan_events.archived_at IS NULL
		AND meal_plan_events.belongs_to_meal_plan = $3
		AND meal_plan_events.id = $1
		AND meal_plans.archived_at IS NULL
		AND meal_plans.id = $3
)
`

type CheckMealPlanOptionExistenceParams struct {
	MealPlanOptionID string
	MealPlanID       string
	MealPlanEventID  sql.NullString
}

func (q *Queries) CheckMealPlanOptionExistence(ctx context.Context, db DBTX, arg *CheckMealPlanOptionExistenceParams) (bool, error) {
	row := db.QueryRowContext(ctx, checkMealPlanOptionExistence, arg.MealPlanEventID, arg.MealPlanOptionID, arg.MealPlanID)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const createMealPlanOption = `-- name: CreateMealPlanOption :exec

INSERT INTO meal_plan_options (
	id,
	assigned_cook,
	assigned_dishwasher,
	meal_id,
	notes,
	meal_scale,
	belongs_to_meal_plan_event,
	chosen
) VALUES (
	$1,
	$2,
	$3,
	$4,
	$5,
	$6,
	$7,
	$8
)
`

type CreateMealPlanOptionParams struct {
	ID                     string
	AssignedCook           sql.NullString
	AssignedDishwasher     sql.NullString
	MealID                 string
	Notes                  string
	MealScale              string
	BelongsToMealPlanEvent sql.NullString
	Chosen                 bool
}

func (q *Queries) CreateMealPlanOption(ctx context.Context, db DBTX, arg *CreateMealPlanOptionParams) error {
	_, err := db.ExecContext(ctx, createMealPlanOption,
		arg.ID,
		arg.AssignedCook,
		arg.AssignedDishwasher,
		arg.MealID,
		arg.Notes,
		arg.MealScale,
		arg.BelongsToMealPlanEvent,
		arg.Chosen,
	)
	return err
}

const finalizeMealPlanOption = `-- name: FinalizeMealPlanOption :exec

UPDATE meal_plan_options SET
	chosen = (belongs_to_meal_plan_event = $1 AND id = $2),
	tiebroken = $3
WHERE archived_at IS NULL
	AND belongs_to_meal_plan_event = $4
	AND id = $2
`

type FinalizeMealPlanOptionParams struct {
	ID                     string
	MealPlanEventID        sql.NullString
	BelongsToMealPlanEvent sql.NullString
	Tiebroken              bool
}

func (q *Queries) FinalizeMealPlanOption(ctx context.Context, db DBTX, arg *FinalizeMealPlanOptionParams) error {
	_, err := db.ExecContext(ctx, finalizeMealPlanOption,
		arg.MealPlanEventID,
		arg.ID,
		arg.Tiebroken,
		arg.BelongsToMealPlanEvent,
	)
	return err
}

const getAllMealPlanOptionsForMealPlanEvent = `-- name: GetAllMealPlanOptionsForMealPlanEvent :many

SELECT
	meal_plan_options.id,
	meal_plan_options.assigned_cook,
	meal_plan_options.assigned_dishwasher,
	meal_plan_options.chosen,
	meal_plan_options.tiebroken,
	meal_plan_options.meal_scale,
	meal_plan_options.meal_id,
	meal_plan_options.notes,
	meal_plan_options.created_at,
	meal_plan_options.last_updated_at,
	meal_plan_options.archived_at,
	meal_plan_options.belongs_to_meal_plan_event,
	meals.id as meal_id,
	meals.name as meal_name,
	meals.description as meal_description,
	meals.min_estimated_portions as meal_min_estimated_portions,
	meals.max_estimated_portions as meal_max_estimated_portions,
	meals.eligible_for_meal_plans as meal_eligible_for_meal_plans,
	meals.created_at as meal_created_at,
	meals.last_updated_at as meal_last_updated_at,
	meals.archived_at as meal_archived_at,
	meals.created_by_user as meal_created_by_user
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

type GetAllMealPlanOptionsForMealPlanEventParams struct {
	MealPlanID      string
	MealPlanEventID sql.NullString
}

type GetAllMealPlanOptionsForMealPlanEventRow struct {
	CreatedAt                time.Time
	MealCreatedAt            time.Time
	MealArchivedAt           sql.NullTime
	MealLastUpdatedAt        sql.NullTime
	ArchivedAt               sql.NullTime
	LastUpdatedAt            sql.NullTime
	MealScale                string
	MealMinEstimatedPortions string
	MealID                   string
	ID                       string
	MealCreatedByUser        string
	Notes                    string
	MealID_2                 string
	MealName                 string
	MealDescription          string
	MealMaxEstimatedPortions sql.NullString
	BelongsToMealPlanEvent   sql.NullString
	AssignedDishwasher       sql.NullString
	AssignedCook             sql.NullString
	MealEligibleForMealPlans bool
	Chosen                   bool
	Tiebroken                bool
}

func (q *Queries) GetAllMealPlanOptionsForMealPlanEvent(ctx context.Context, db DBTX, arg *GetAllMealPlanOptionsForMealPlanEventParams) ([]*GetAllMealPlanOptionsForMealPlanEventRow, error) {
	rows, err := db.QueryContext(ctx, getAllMealPlanOptionsForMealPlanEvent, arg.MealPlanEventID, arg.MealPlanID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*GetAllMealPlanOptionsForMealPlanEventRow{}
	for rows.Next() {
		var i GetAllMealPlanOptionsForMealPlanEventRow
		if err := rows.Scan(
			&i.ID,
			&i.AssignedCook,
			&i.AssignedDishwasher,
			&i.Chosen,
			&i.Tiebroken,
			&i.MealScale,
			&i.MealID,
			&i.Notes,
			&i.CreatedAt,
			&i.LastUpdatedAt,
			&i.ArchivedAt,
			&i.BelongsToMealPlanEvent,
			&i.MealID_2,
			&i.MealName,
			&i.MealDescription,
			&i.MealMinEstimatedPortions,
			&i.MealMaxEstimatedPortions,
			&i.MealEligibleForMealPlans,
			&i.MealCreatedAt,
			&i.MealLastUpdatedAt,
			&i.MealArchivedAt,
			&i.MealCreatedByUser,
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

const getMealPlanOption = `-- name: GetMealPlanOption :one

SELECT
	meal_plan_options.id,
	meal_plan_options.assigned_cook,
	meal_plan_options.assigned_dishwasher,
	meal_plan_options.chosen,
	meal_plan_options.tiebroken,
	meal_plan_options.meal_scale,
	meal_plan_options.meal_id,
	meal_plan_options.notes,
	meal_plan_options.created_at,
	meal_plan_options.last_updated_at,
	meal_plan_options.archived_at,
	meal_plan_options.belongs_to_meal_plan_event,
	meals.id as meal_id,
	meals.name as meal_name,
	meals.description as meal_description,
	meals.min_estimated_portions as meal_min_estimated_portions,
	meals.max_estimated_portions as meal_max_estimated_portions,
	meals.eligible_for_meal_plans as meal_eligible_for_meal_plans,
	meals.created_at as meal_created_at,
	meals.last_updated_at as meal_last_updated_at,
	meals.archived_at as meal_archived_at,
	meals.created_by_user as meal_created_by_user
FROM meal_plan_options
	JOIN meal_plan_events ON meal_plan_options.belongs_to_meal_plan_event = meal_plan_events.id
	JOIN meal_plans ON meal_plan_events.belongs_to_meal_plan = meal_plans.id
	JOIN meals ON meal_plan_options.meal_id = meals.id
WHERE meal_plan_options.archived_at IS NULL
	AND meal_plan_options.belongs_to_meal_plan_event = $1
	AND meal_plan_options.id = $2
	AND meal_plan_events.id = $1
	AND meal_plan_events.belongs_to_meal_plan = $3
	AND meal_plans.archived_at IS NULL
	AND meal_plans.id = $3
`

type GetMealPlanOptionParams struct {
	MealPlanOptionID string
	MealPlanID       string
	MealPlanEventID  sql.NullString
}

type GetMealPlanOptionRow struct {
	CreatedAt                time.Time
	MealCreatedAt            time.Time
	MealArchivedAt           sql.NullTime
	MealLastUpdatedAt        sql.NullTime
	ArchivedAt               sql.NullTime
	LastUpdatedAt            sql.NullTime
	MealScale                string
	MealMinEstimatedPortions string
	MealID                   string
	ID                       string
	MealCreatedByUser        string
	Notes                    string
	MealID_2                 string
	MealName                 string
	MealDescription          string
	MealMaxEstimatedPortions sql.NullString
	BelongsToMealPlanEvent   sql.NullString
	AssignedDishwasher       sql.NullString
	AssignedCook             sql.NullString
	MealEligibleForMealPlans bool
	Chosen                   bool
	Tiebroken                bool
}

func (q *Queries) GetMealPlanOption(ctx context.Context, db DBTX, arg *GetMealPlanOptionParams) (*GetMealPlanOptionRow, error) {
	row := db.QueryRowContext(ctx, getMealPlanOption, arg.MealPlanEventID, arg.MealPlanOptionID, arg.MealPlanID)
	var i GetMealPlanOptionRow
	err := row.Scan(
		&i.ID,
		&i.AssignedCook,
		&i.AssignedDishwasher,
		&i.Chosen,
		&i.Tiebroken,
		&i.MealScale,
		&i.MealID,
		&i.Notes,
		&i.CreatedAt,
		&i.LastUpdatedAt,
		&i.ArchivedAt,
		&i.BelongsToMealPlanEvent,
		&i.MealID_2,
		&i.MealName,
		&i.MealDescription,
		&i.MealMinEstimatedPortions,
		&i.MealMaxEstimatedPortions,
		&i.MealEligibleForMealPlans,
		&i.MealCreatedAt,
		&i.MealLastUpdatedAt,
		&i.MealArchivedAt,
		&i.MealCreatedByUser,
	)
	return &i, err
}

const getMealPlanOptionByID = `-- name: GetMealPlanOptionByID :one

SELECT
	meal_plan_options.id,
	meal_plan_options.assigned_cook,
	meal_plan_options.assigned_dishwasher,
	meal_plan_options.chosen,
	meal_plan_options.tiebroken,
	meal_plan_options.meal_scale,
	meal_plan_options.meal_id,
	meal_plan_options.notes,
	meal_plan_options.created_at,
	meal_plan_options.last_updated_at,
	meal_plan_options.archived_at,
	meal_plan_options.belongs_to_meal_plan_event,
	meals.id as meal_id,
	meals.name as meal_name,
	meals.description as meal_description,
	meals.min_estimated_portions as meal_min_estimated_portions,
	meals.max_estimated_portions as meal_max_estimated_portions,
	meals.eligible_for_meal_plans as meal_eligible_for_meal_plans,
	meals.created_at as meal_created_at,
	meals.last_updated_at as meal_last_updated_at,
	meals.archived_at as meal_archived_at,
	meals.created_by_user as meal_created_by_user
FROM meal_plan_options
	JOIN meal_plan_events ON meal_plan_options.belongs_to_meal_plan_event = meal_plan_events.id
	JOIN meal_plans ON meal_plan_events.belongs_to_meal_plan = meal_plans.id
	JOIN meals ON meal_plan_options.meal_id = meals.id
WHERE meal_plan_options.archived_at IS NULL
	AND meal_plan_options.id = $1
`

type GetMealPlanOptionByIDRow struct {
	CreatedAt                time.Time
	MealCreatedAt            time.Time
	MealArchivedAt           sql.NullTime
	MealLastUpdatedAt        sql.NullTime
	ArchivedAt               sql.NullTime
	LastUpdatedAt            sql.NullTime
	MealScale                string
	MealMinEstimatedPortions string
	MealID                   string
	ID                       string
	MealCreatedByUser        string
	Notes                    string
	MealID_2                 string
	MealName                 string
	MealDescription          string
	MealMaxEstimatedPortions sql.NullString
	BelongsToMealPlanEvent   sql.NullString
	AssignedDishwasher       sql.NullString
	AssignedCook             sql.NullString
	MealEligibleForMealPlans bool
	Chosen                   bool
	Tiebroken                bool
}

func (q *Queries) GetMealPlanOptionByID(ctx context.Context, db DBTX, mealPlanOptionID string) (*GetMealPlanOptionByIDRow, error) {
	row := db.QueryRowContext(ctx, getMealPlanOptionByID, mealPlanOptionID)
	var i GetMealPlanOptionByIDRow
	err := row.Scan(
		&i.ID,
		&i.AssignedCook,
		&i.AssignedDishwasher,
		&i.Chosen,
		&i.Tiebroken,
		&i.MealScale,
		&i.MealID,
		&i.Notes,
		&i.CreatedAt,
		&i.LastUpdatedAt,
		&i.ArchivedAt,
		&i.BelongsToMealPlanEvent,
		&i.MealID_2,
		&i.MealName,
		&i.MealDescription,
		&i.MealMinEstimatedPortions,
		&i.MealMaxEstimatedPortions,
		&i.MealEligibleForMealPlans,
		&i.MealCreatedAt,
		&i.MealLastUpdatedAt,
		&i.MealArchivedAt,
		&i.MealCreatedByUser,
	)
	return &i, err
}

const getMealPlanOptions = `-- name: GetMealPlanOptions :many

SELECT
	meal_plan_options.id,
	meal_plan_options.assigned_cook,
	meal_plan_options.assigned_dishwasher,
	meal_plan_options.chosen,
	meal_plan_options.tiebroken,
	meal_plan_options.meal_scale,
	meal_plan_options.meal_id,
	meal_plan_options.notes,
	meal_plan_options.created_at,
	meal_plan_options.last_updated_at,
	meal_plan_options.archived_at,
	meal_plan_options.belongs_to_meal_plan_event,
	meals.id as meal_id,
	meals.name as meal_name,
	meals.description as meal_description,
	meals.min_estimated_portions as meal_min_estimated_portions,
	meals.max_estimated_portions as meal_max_estimated_portions,
	meals.eligible_for_meal_plans as meal_eligible_for_meal_plans,
	meals.created_at as meal_created_at,
	meals.last_updated_at as meal_last_updated_at,
	meals.archived_at as meal_archived_at,
	meals.created_by_user as meal_created_by_user,
	(
		SELECT
			COUNT(meal_plan_options.id)
		FROM
			meal_plan_options
		WHERE
			meal_plan_options.archived_at IS NULL
			AND meal_plan_options.belongs_to_meal_plan_event = $1
			AND meal_plan_options.created_at > COALESCE($2, (SELECT NOW() - interval '999 years'))
			AND meal_plan_options.created_at < COALESCE($3, (SELECT NOW() + interval '999 years'))
			AND (meal_plan_options.last_updated_at IS NULL OR meal_plan_options.last_updated_at > COALESCE($4, (SELECT NOW() - interval '999 years')))
			AND (meal_plan_options.last_updated_at IS NULL OR meal_plan_options.last_updated_at < COALESCE($5, (SELECT NOW() + interval '999 years')))
	) as filtered_count,
	(
		SELECT COUNT(meal_plan_options.id)
		FROM meal_plan_options
		WHERE meal_plan_options.archived_at IS NULL
	) as total_count
FROM meal_plan_options
	JOIN meal_plan_events ON meal_plan_options.belongs_to_meal_plan_event = meal_plan_events.id
	JOIN meal_plans ON meal_plan_events.belongs_to_meal_plan = meal_plans.id
	JOIN meals ON meal_plan_options.meal_id = meals.id
WHERE
	meal_plan_options.archived_at IS NULL
	AND meal_plan_options.belongs_to_meal_plan_event = $1
	AND meal_plan_events.id = $1
	AND meal_plan_events.belongs_to_meal_plan = $6
	AND meal_plans.archived_at IS NULL
	AND meal_plans.id = $6
OFFSET $7
LIMIT $8
`

type GetMealPlanOptionsParams struct {
	CreatedAfter    sql.NullTime
	CreatedBefore   sql.NullTime
	UpdatedAfter    sql.NullTime
	UpdatedBefore   sql.NullTime
	MealPlanID      string
	MealPlanEventID sql.NullString
	QueryOffset     sql.NullInt32
	QueryLimit      sql.NullInt32
}

type GetMealPlanOptionsRow struct {
	CreatedAt                time.Time
	MealCreatedAt            time.Time
	MealArchivedAt           sql.NullTime
	MealLastUpdatedAt        sql.NullTime
	ArchivedAt               sql.NullTime
	LastUpdatedAt            sql.NullTime
	MealID                   string
	MealMinEstimatedPortions string
	ID                       string
	MealScale                string
	MealCreatedByUser        string
	Notes                    string
	MealID_2                 string
	MealName                 string
	MealDescription          string
	BelongsToMealPlanEvent   sql.NullString
	MealMaxEstimatedPortions sql.NullString
	AssignedDishwasher       sql.NullString
	AssignedCook             sql.NullString
	FilteredCount            int64
	TotalCount               int64
	MealEligibleForMealPlans bool
	Chosen                   bool
	Tiebroken                bool
}

func (q *Queries) GetMealPlanOptions(ctx context.Context, db DBTX, arg *GetMealPlanOptionsParams) ([]*GetMealPlanOptionsRow, error) {
	rows, err := db.QueryContext(ctx, getMealPlanOptions,
		arg.MealPlanEventID,
		arg.CreatedAfter,
		arg.CreatedBefore,
		arg.UpdatedAfter,
		arg.UpdatedBefore,
		arg.MealPlanID,
		arg.QueryOffset,
		arg.QueryLimit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*GetMealPlanOptionsRow{}
	for rows.Next() {
		var i GetMealPlanOptionsRow
		if err := rows.Scan(
			&i.ID,
			&i.AssignedCook,
			&i.AssignedDishwasher,
			&i.Chosen,
			&i.Tiebroken,
			&i.MealScale,
			&i.MealID,
			&i.Notes,
			&i.CreatedAt,
			&i.LastUpdatedAt,
			&i.ArchivedAt,
			&i.BelongsToMealPlanEvent,
			&i.MealID_2,
			&i.MealName,
			&i.MealDescription,
			&i.MealMinEstimatedPortions,
			&i.MealMaxEstimatedPortions,
			&i.MealEligibleForMealPlans,
			&i.MealCreatedAt,
			&i.MealLastUpdatedAt,
			&i.MealArchivedAt,
			&i.MealCreatedByUser,
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

const updateMealPlanOption = `-- name: UpdateMealPlanOption :execrows

UPDATE meal_plan_options
SET
	assigned_cook = $1,
	assigned_dishwasher = $2,
	meal_id = $3,
	notes = $4,
	meal_scale = $5,
	last_updated_at = NOW()
WHERE archived_at IS NULL
	AND belongs_to_meal_plan_event = $6
	AND id = $7
`

type UpdateMealPlanOptionParams struct {
	MealID             string
	Notes              string
	MealScale          string
	MealPlanOptionID   string
	AssignedCook       sql.NullString
	AssignedDishwasher sql.NullString
	MealPlanEventID    sql.NullString
}

func (q *Queries) UpdateMealPlanOption(ctx context.Context, db DBTX, arg *UpdateMealPlanOptionParams) (int64, error) {
	result, err := db.ExecContext(ctx, updateMealPlanOption,
		arg.AssignedCook,
		arg.AssignedDishwasher,
		arg.MealID,
		arg.Notes,
		arg.MealScale,
		arg.MealPlanEventID,
		arg.MealPlanOptionID,
	)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
