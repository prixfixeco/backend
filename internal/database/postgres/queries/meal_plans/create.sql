-- name: CreateMealPlan :exec

INSERT INTO meal_plans (id,notes,status,voting_deadline,belongs_to_household,created_by_user) VALUES ($1,$2,$3,$4,$5,$6);