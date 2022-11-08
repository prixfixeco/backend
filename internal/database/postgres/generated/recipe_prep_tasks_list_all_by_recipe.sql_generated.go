// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: recipe_prep_tasks_list_all_by_recipe.sql

package generated

import (
	"context"
	"database/sql"
)

const GetRecipePrepTasksForRecipe = `-- name: GetRecipePrepTasksForRecipe :many
SELECT
    recipe_prep_tasks.id,
    recipe_prep_tasks.notes,
    recipe_prep_tasks.explicit_storage_instructions,
    recipe_prep_tasks.minimum_time_buffer_before_recipe_in_seconds,
    recipe_prep_tasks.maximum_time_buffer_before_recipe_in_seconds,
    recipe_prep_tasks.storage_type,
    recipe_prep_tasks.minimum_storage_temperature_in_celsius,
    recipe_prep_tasks.maximum_storage_temperature_in_celsius,
    recipe_prep_tasks.belongs_to_recipe,
    recipe_prep_tasks.created_at,
    recipe_prep_tasks.last_updated_at,
    recipe_prep_tasks.archived_at,
    recipe_prep_task_steps.id,
    recipe_prep_task_steps.belongs_to_recipe_step,
    recipe_prep_task_steps.belongs_to_recipe_prep_task,
    recipe_prep_task_steps.satisfies_recipe_step
FROM recipe_prep_tasks
     FULL OUTER JOIN recipe_prep_task_steps ON recipe_prep_task_steps.belongs_to_recipe_prep_task=recipe_prep_tasks.id
     FULL OUTER JOIN recipe_steps ON recipe_prep_task_steps.belongs_to_recipe_step=recipe_steps.id
     FULL OUTER JOIN recipes ON recipe_prep_tasks.belongs_to_recipe=recipes.id
WHERE recipe_prep_tasks.archived_at IS NULL
  AND recipe_steps.archived_at IS NULL
  AND recipes.archived_at IS NULL
  AND recipes.id = $1
  AND recipe_steps.belongs_to_recipe = $1
`

type GetRecipePrepTasksForRecipeRow struct {
	CreatedAt                              sql.NullTime             `db:"created_at"`
	ArchivedAt                             sql.NullTime             `db:"archived_at"`
	LastUpdatedAt                          sql.NullTime             `db:"last_updated_at"`
	ID                                     sql.NullString           `db:"id"`
	BelongsToRecipeStep                    sql.NullString           `db:"belongs_to_recipe_step"`
	StorageType                            NullStorageContainerType `db:"storage_type"`
	MinimumStorageTemperatureInCelsius     sql.NullString           `db:"minimum_storage_temperature_in_celsius"`
	MaximumStorageTemperatureInCelsius     sql.NullString           `db:"maximum_storage_temperature_in_celsius"`
	BelongsToRecipe                        sql.NullString           `db:"belongs_to_recipe"`
	BelongsToRecipePrepTask                sql.NullString           `db:"belongs_to_recipe_prep_task"`
	ExplicitStorageInstructions            sql.NullString           `db:"explicit_storage_instructions"`
	Notes                                  sql.NullString           `db:"notes"`
	ID_2                                   sql.NullString           `db:"id_2"`
	MaximumTimeBufferBeforeRecipeInSeconds sql.NullInt32            `db:"maximum_time_buffer_before_recipe_in_seconds"`
	MinimumTimeBufferBeforeRecipeInSeconds sql.NullInt32            `db:"minimum_time_buffer_before_recipe_in_seconds"`
	SatisfiesRecipeStep                    sql.NullBool             `db:"satisfies_recipe_step"`
}

func (q *Queries) GetRecipePrepTasksForRecipe(ctx context.Context, id string) ([]*GetRecipePrepTasksForRecipeRow, error) {
	rows, err := q.query(ctx, q.getRecipePrepTasksForRecipeStmt, GetRecipePrepTasksForRecipe, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*GetRecipePrepTasksForRecipeRow{}
	for rows.Next() {
		var i GetRecipePrepTasksForRecipeRow
		if err := rows.Scan(
			&i.ID,
			&i.Notes,
			&i.ExplicitStorageInstructions,
			&i.MinimumTimeBufferBeforeRecipeInSeconds,
			&i.MaximumTimeBufferBeforeRecipeInSeconds,
			&i.StorageType,
			&i.MinimumStorageTemperatureInCelsius,
			&i.MaximumStorageTemperatureInCelsius,
			&i.BelongsToRecipe,
			&i.CreatedAt,
			&i.LastUpdatedAt,
			&i.ArchivedAt,
			&i.ID_2,
			&i.BelongsToRecipeStep,
			&i.BelongsToRecipePrepTask,
			&i.SatisfiesRecipeStep,
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
