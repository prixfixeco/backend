-- name: MealPlanEventIsEligibleForVoting :one

SELECT
  EXISTS (
    SELECT
      meal_plan_events.id
    FROM
      meal_plan_events
      JOIN meal_plans ON meal_plan_events.belongs_to_meal_plan = meal_plans.id
    WHERE
      meal_plan_events.archived_at IS NULL
      AND meal_plans.id = sqlc.arg(meal_plan_id)
      AND meal_plans.status = 'awaiting_votes'
      AND meal_plans.archived_at IS NULL
      AND meal_plan_events.id = sqlc.arg(meal_plan_event_id)
      AND meal_plan_events.archived_at IS NULL
  );