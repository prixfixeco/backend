-- name: CreateRecipePrepTaskStep :exec

INSERT INTO recipe_prep_task_steps (
	id,
	belongs_to_recipe_prep_task,
	belongs_to_recipe_step,
	satisfies_recipe_step
) VALUES (
	sqlc.arg(id),
	sqlc.arg(belongs_to_recipe_prep_task),
	sqlc.arg(belongs_to_recipe_step),
	sqlc.arg(satisfies_recipe_step)
);
