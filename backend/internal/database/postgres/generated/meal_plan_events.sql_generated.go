// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: meal_plan_events.sql

package generated

import (
	"context"
	"database/sql"
	"time"
)

const archiveMealPlanEvent = `-- name: ArchiveMealPlanEvent :execrows
UPDATE meal_plan_events SET archived_at = NOW() WHERE archived_at IS NULL AND id = $1 AND belongs_to_meal_plan = $2
`

type ArchiveMealPlanEventParams struct {
	ID                string
	BelongsToMealPlan string
}

func (q *Queries) ArchiveMealPlanEvent(ctx context.Context, db DBTX, arg *ArchiveMealPlanEventParams) (int64, error) {
	result, err := db.ExecContext(ctx, archiveMealPlanEvent, arg.ID, arg.BelongsToMealPlan)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

const checkMealPlanEventExistence = `-- name: CheckMealPlanEventExistence :one
SELECT EXISTS (
	SELECT meal_plan_events.id
	FROM meal_plan_events
	WHERE meal_plan_events.archived_at IS NULL
		AND meal_plan_events.id = $1
		AND meal_plan_events.belongs_to_meal_plan = $2
)
`

type CheckMealPlanEventExistenceParams struct {
	ID         string
	MealPlanID string
}

func (q *Queries) CheckMealPlanEventExistence(ctx context.Context, db DBTX, arg *CheckMealPlanEventExistenceParams) (bool, error) {
	row := db.QueryRowContext(ctx, checkMealPlanEventExistence, arg.ID, arg.MealPlanID)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const createMealPlanEvent = `-- name: CreateMealPlanEvent :exec
INSERT INTO meal_plan_events (
	id,
	notes,
	starts_at,
	ends_at,
	meal_name,
	belongs_to_meal_plan
) VALUES (
	$1,
	$2,
	$3,
	$4,
	$5,
	$6
)
`

type CreateMealPlanEventParams struct {
	ID                string
	Notes             string
	StartsAt          time.Time
	EndsAt            time.Time
	MealName          MealName
	BelongsToMealPlan string
}

func (q *Queries) CreateMealPlanEvent(ctx context.Context, db DBTX, arg *CreateMealPlanEventParams) error {
	_, err := db.ExecContext(ctx, createMealPlanEvent,
		arg.ID,
		arg.Notes,
		arg.StartsAt,
		arg.EndsAt,
		arg.MealName,
		arg.BelongsToMealPlan,
	)
	return err
}

const getAllMealPlanEventsForMealPlan = `-- name: GetAllMealPlanEventsForMealPlan :many
SELECT
	meal_plan_events.id,
	meal_plan_events.notes,
	meal_plan_events.starts_at,
	meal_plan_events.ends_at,
	meal_plan_events.meal_name,
	meal_plan_events.belongs_to_meal_plan,
	meal_plan_events.created_at,
	meal_plan_events.last_updated_at,
	meal_plan_events.archived_at
FROM meal_plan_events
WHERE
	meal_plan_events.archived_at IS NULL
	AND meal_plan_events.belongs_to_meal_plan = $1
`

func (q *Queries) GetAllMealPlanEventsForMealPlan(ctx context.Context, db DBTX, mealPlanID string) ([]*MealPlanEvents, error) {
	rows, err := db.QueryContext(ctx, getAllMealPlanEventsForMealPlan, mealPlanID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*MealPlanEvents{}
	for rows.Next() {
		var i MealPlanEvents
		if err := rows.Scan(
			&i.ID,
			&i.Notes,
			&i.StartsAt,
			&i.EndsAt,
			&i.MealName,
			&i.BelongsToMealPlan,
			&i.CreatedAt,
			&i.LastUpdatedAt,
			&i.ArchivedAt,
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

const getMealPlanEvent = `-- name: GetMealPlanEvent :one
SELECT
	meal_plan_events.id,
	meal_plan_events.notes,
	meal_plan_events.starts_at,
	meal_plan_events.ends_at,
	meal_plan_events.meal_name,
	meal_plan_events.belongs_to_meal_plan,
	meal_plan_events.created_at,
	meal_plan_events.last_updated_at,
	meal_plan_events.archived_at
FROM meal_plan_events
WHERE meal_plan_events.archived_at IS NULL
	AND meal_plan_events.id = $1
	AND meal_plan_events.belongs_to_meal_plan = $2
`

type GetMealPlanEventParams struct {
	ID                string
	BelongsToMealPlan string
}

func (q *Queries) GetMealPlanEvent(ctx context.Context, db DBTX, arg *GetMealPlanEventParams) (*MealPlanEvents, error) {
	row := db.QueryRowContext(ctx, getMealPlanEvent, arg.ID, arg.BelongsToMealPlan)
	var i MealPlanEvents
	err := row.Scan(
		&i.ID,
		&i.Notes,
		&i.StartsAt,
		&i.EndsAt,
		&i.MealName,
		&i.BelongsToMealPlan,
		&i.CreatedAt,
		&i.LastUpdatedAt,
		&i.ArchivedAt,
	)
	return &i, err
}

const getMealPlanEvents = `-- name: GetMealPlanEvents :many
SELECT
	meal_plan_events.id,
	meal_plan_events.notes,
	meal_plan_events.starts_at,
	meal_plan_events.ends_at,
	meal_plan_events.meal_name,
	meal_plan_events.belongs_to_meal_plan,
	meal_plan_events.created_at,
	meal_plan_events.last_updated_at,
	meal_plan_events.archived_at,
	(
		SELECT COUNT(meal_plan_events.id)
		FROM meal_plan_events
		WHERE meal_plan_events.archived_at IS NULL
			AND
			meal_plan_events.created_at > COALESCE($1, (SELECT NOW() - '999 years'::INTERVAL))
			AND meal_plan_events.created_at < COALESCE($2, (SELECT NOW() + '999 years'::INTERVAL))
			AND (
				meal_plan_events.last_updated_at IS NULL
				OR meal_plan_events.last_updated_at > COALESCE($3, (SELECT NOW() - '999 years'::INTERVAL))
			)
			AND (
				meal_plan_events.last_updated_at IS NULL
				OR meal_plan_events.last_updated_at < COALESCE($4, (SELECT NOW() + '999 years'::INTERVAL))
			)
			AND (NOT COALESCE($5, false)::boolean OR meal_plan_events.archived_at = NULL)
			AND meal_plan_events.belongs_to_meal_plan = $6
	) AS filtered_count,
	(
		SELECT COUNT(meal_plan_events.id)
		FROM meal_plan_events
		WHERE meal_plan_events.archived_at IS NULL
			AND meal_plan_events.belongs_to_meal_plan = $6
	) AS total_count
FROM meal_plan_events
WHERE
	meal_plan_events.archived_at IS NULL
	AND meal_plan_events.created_at > COALESCE($1, (SELECT NOW() - '999 years'::INTERVAL))
	AND meal_plan_events.created_at < COALESCE($2, (SELECT NOW() + '999 years'::INTERVAL))
	AND (
		meal_plan_events.last_updated_at IS NULL
		OR meal_plan_events.last_updated_at > COALESCE($4, (SELECT NOW() - '999 years'::INTERVAL))
	)
	AND (
		meal_plan_events.last_updated_at IS NULL
		OR meal_plan_events.last_updated_at < COALESCE($3, (SELECT NOW() + '999 years'::INTERVAL))
	)
			AND (NOT COALESCE($5, false)::boolean OR meal_plan_events.archived_at = NULL)
	AND meal_plan_events.belongs_to_meal_plan = $6
GROUP BY meal_plan_events.id
ORDER BY meal_plan_events.id
LIMIT $8
OFFSET $7
`

type GetMealPlanEventsParams struct {
	CreatedAfter    sql.NullTime
	CreatedBefore   sql.NullTime
	UpdatedBefore   sql.NullTime
	UpdatedAfter    sql.NullTime
	MealPlanID      string
	QueryOffset     sql.NullInt32
	QueryLimit      sql.NullInt32
	IncludeArchived sql.NullBool
}

type GetMealPlanEventsRow struct {
	ID                string
	Notes             string
	StartsAt          time.Time
	EndsAt            time.Time
	MealName          MealName
	BelongsToMealPlan string
	CreatedAt         time.Time
	LastUpdatedAt     sql.NullTime
	ArchivedAt        sql.NullTime
	FilteredCount     int64
	TotalCount        int64
}

func (q *Queries) GetMealPlanEvents(ctx context.Context, db DBTX, arg *GetMealPlanEventsParams) ([]*GetMealPlanEventsRow, error) {
	rows, err := db.QueryContext(ctx, getMealPlanEvents,
		arg.CreatedAfter,
		arg.CreatedBefore,
		arg.UpdatedBefore,
		arg.UpdatedAfter,
		arg.IncludeArchived,
		arg.MealPlanID,
		arg.QueryOffset,
		arg.QueryLimit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*GetMealPlanEventsRow{}
	for rows.Next() {
		var i GetMealPlanEventsRow
		if err := rows.Scan(
			&i.ID,
			&i.Notes,
			&i.StartsAt,
			&i.EndsAt,
			&i.MealName,
			&i.BelongsToMealPlan,
			&i.CreatedAt,
			&i.LastUpdatedAt,
			&i.ArchivedAt,
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

const mealPlanEventIsEligibleForVoting = `-- name: MealPlanEventIsEligibleForVoting :one
SELECT EXISTS (
	SELECT meal_plan_events.id
	FROM meal_plan_events
		JOIN meal_plans ON meal_plan_events.belongs_to_meal_plan = meal_plans.id
	WHERE
		meal_plan_events.archived_at IS NULL
		AND meal_plans.id = $1
		AND meal_plans.status = 'awaiting_votes'
		AND meal_plans.archived_at IS NULL
		AND meal_plan_events.id = $2
		AND meal_plan_events.archived_at IS NULL
)
`

type MealPlanEventIsEligibleForVotingParams struct {
	MealPlanID      string
	MealPlanEventID string
}

func (q *Queries) MealPlanEventIsEligibleForVoting(ctx context.Context, db DBTX, arg *MealPlanEventIsEligibleForVotingParams) (bool, error) {
	row := db.QueryRowContext(ctx, mealPlanEventIsEligibleForVoting, arg.MealPlanID, arg.MealPlanEventID)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const updateMealPlanEvent = `-- name: UpdateMealPlanEvent :execrows
UPDATE meal_plan_events SET
	notes = $1,
	starts_at = $2,
	ends_at = $3,
	meal_name = $4,
	belongs_to_meal_plan = $5,
	last_updated_at = NOW()
WHERE archived_at IS NULL
	AND id = $6
`

type UpdateMealPlanEventParams struct {
	Notes             string
	StartsAt          time.Time
	EndsAt            time.Time
	MealName          MealName
	BelongsToMealPlan string
	ID                string
}

func (q *Queries) UpdateMealPlanEvent(ctx context.Context, db DBTX, arg *UpdateMealPlanEventParams) (int64, error) {
	result, err := db.ExecContext(ctx, updateMealPlanEvent,
		arg.Notes,
		arg.StartsAt,
		arg.EndsAt,
		arg.MealName,
		arg.BelongsToMealPlan,
		arg.ID,
	)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
