package postgres

import (
	"context"
	"database/sql/driver"
	"errors"
	"testing"

	"github.com/dinnerdonebetter/backend/pkg/types"
	"github.com/dinnerdonebetter/backend/pkg/types/fakes"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func buildMockRowsFromMealPlanTasksWithRecipePrepTaskSteps(includeCounts bool, filteredCount uint64, mealPlanTasks ...*types.MealPlanTask) *sqlmock.Rows {
	columns := []string{
		"meal_plan_tasks.id",
		"meal_plan_options.id",
		"meal_plan_options.assigned_cook",
		"meal_plan_options.assigned_dishwasher",
		"meal_plan_options.chosen",
		"meal_plan_options.tiebroken",
		"meal_plan_options.meal_scale",
		"meal_plan_options.meal_id",
		"meal_plan_options.notes",
		"meal_plan_options.created_at",
		"meal_plan_options.last_updated_at",
		"meal_plan_options.archived_at",
		"meal_plan_options.belongs_to_meal_plan_event",
		"recipe_prep_tasks.id",
		"recipe_prep_tasks.name",
		"recipe_prep_tasks.description",
		"recipe_prep_tasks.notes",
		"recipe_prep_tasks.optional",
		"recipe_prep_tasks.explicit_storage_instructions",
		"recipe_prep_tasks.minimum_time_buffer_before_recipe_in_seconds",
		"recipe_prep_tasks.maximum_time_buffer_before_recipe_in_seconds",
		"recipe_prep_tasks.storage_type",
		"recipe_prep_tasks.minimum_storage_temperature_in_celsius",
		"recipe_prep_tasks.maximum_storage_temperature_in_celsius",
		"recipe_prep_tasks.belongs_to_recipe",
		"recipe_prep_tasks.created_at",
		"recipe_prep_tasks.last_updated_at",
		"recipe_prep_tasks.archived_at",
		"recipe_prep_task_steps.id",
		"recipe_prep_task_steps.belongs_to_recipe_step",
		"recipe_prep_task_steps.belongs_to_recipe_prep_task",
		"recipe_prep_task_steps.satisfies_recipe_step",
		"meal_plan_tasks.created_at",
		"meal_plan_tasks.last_updated_at",
		"meal_plan_tasks.completed_at",
		"meal_plan_tasks.status",
		"meal_plan_tasks.creation_explanation",
		"meal_plan_tasks.status_explanation",
		"meal_plan_tasks.assigned_to_user",
	}

	if includeCounts {
		columns = append(columns, "filtered_count", "total_count")
	}

	exampleRows := sqlmock.NewRows(columns)

	for _, x := range mealPlanTasks {
		for _, y := range x.RecipePrepTask.TaskSteps {
			rowValues := []driver.Value{
				&x.ID,
				&x.MealPlanOption.ID,
				&x.MealPlanOption.AssignedCook,
				&x.MealPlanOption.AssignedDishwasher,
				&x.MealPlanOption.Chosen,
				&x.MealPlanOption.TieBroken,
				&x.MealPlanOption.MealScale,
				&x.MealPlanOption.Meal.ID,
				&x.MealPlanOption.Notes,
				&x.MealPlanOption.CreatedAt,
				&x.MealPlanOption.LastUpdatedAt,
				&x.MealPlanOption.ArchivedAt,
				&x.MealPlanOption.BelongsToMealPlanEvent,
				&x.RecipePrepTask.ID,
				&x.RecipePrepTask.Name,
				&x.RecipePrepTask.Description,
				&x.RecipePrepTask.Notes,
				&x.RecipePrepTask.Optional,
				&x.RecipePrepTask.ExplicitStorageInstructions,
				&x.RecipePrepTask.MinimumTimeBufferBeforeRecipeInSeconds,
				&x.RecipePrepTask.MaximumTimeBufferBeforeRecipeInSeconds,
				&x.RecipePrepTask.StorageType,
				&x.RecipePrepTask.MinimumStorageTemperatureInCelsius,
				&x.RecipePrepTask.MaximumStorageTemperatureInCelsius,
				&x.RecipePrepTask.BelongsToRecipe,
				&x.RecipePrepTask.CreatedAt,
				&x.RecipePrepTask.LastUpdatedAt,
				&x.RecipePrepTask.ArchivedAt,
				&y.ID,
				&y.BelongsToRecipeStep,
				&y.BelongsToRecipePrepTask,
				&y.SatisfiesRecipeStep,
				&x.CreatedAt,
				&x.LastUpdatedAt,
				&x.CompletedAt,
				&x.Status,
				&x.CreationExplanation,
				&x.StatusExplanation,
				&x.AssignedToUser,
			}

			if includeCounts {
				rowValues = append(rowValues, filteredCount, len(mealPlanTasks))
			}

			exampleRows.AddRow(rowValues...)
		}
	}

	return exampleRows
}

func TestQuerier_MealPlanTaskExists(T *testing.T) {
	T.Parallel()

	T.Run("with invalid meal plan ID", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()

		exampleMealPlanTaskID := fakes.BuildFakeID()

		c, db := buildTestClient(t)

		actual, err := c.MealPlanTaskExists(ctx, "", exampleMealPlanTaskID)
		assert.Error(t, err)
		assert.False(t, actual)

		mock.AssertExpectationsForObjects(t, db)
	})

	T.Run("with invalid meal plan task ID", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()

		exampleMealPlanID := fakes.BuildFakeID()

		c, db := buildTestClient(t)

		actual, err := c.MealPlanTaskExists(ctx, exampleMealPlanID, "")
		assert.Error(t, err)
		assert.False(t, actual)

		mock.AssertExpectationsForObjects(t, db)
	})
}

func TestQuerier_GetMealPlanTask(T *testing.T) {
	T.Parallel()

	T.Run("standard", func(t *testing.T) {
		t.Parallel()

		exampleMealPlanTask := fakes.BuildFakeMealPlanTask()

		ctx := context.Background()
		c, db := buildTestClient(t)

		getMealPlanTaskByIDArgs := []any{
			exampleMealPlanTask.ID,
		}

		db.ExpectQuery(formatQueryForSQLMock(getMealPlanTasksQuery)).
			WithArgs(interfaceToDriverValue(getMealPlanTaskByIDArgs)...).
			WillReturnRows(buildMockRowsFromMealPlanTasksWithRecipePrepTaskSteps(false, 0, exampleMealPlanTask))

		actual, err := c.GetMealPlanTask(ctx, exampleMealPlanTask.ID)
		assert.NoError(t, err)
		assert.Equal(t, exampleMealPlanTask, actual)

		mock.AssertExpectationsForObjects(t, db)
	})

	T.Run("with invalid meal plan ID", func(t *testing.T) {
		t.Parallel()

		exampleMealPlanTaskID := fakes.BuildFakeID()

		ctx := context.Background()
		c, db := buildTestClient(t)

		actual, err := c.GetMealPlanTask(ctx, exampleMealPlanTaskID)
		assert.Error(t, err)
		assert.Nil(t, actual)

		mock.AssertExpectationsForObjects(t, db)
	})

	T.Run("with invalid meal plan task ID", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		c, db := buildTestClient(t)

		actual, err := c.GetMealPlanTask(ctx, "")
		assert.Error(t, err)
		assert.Nil(t, actual)

		mock.AssertExpectationsForObjects(t, db)
	})

	T.Run("with invalid response from database", func(t *testing.T) {
		t.Parallel()

		exampleMealPlanTask := fakes.BuildFakeMealPlanTask()

		ctx := context.Background()
		c, db := buildTestClient(t)

		getMealPlanTaskByIDArgs := []any{
			exampleMealPlanTask.ID,
		}

		db.ExpectQuery(formatQueryForSQLMock(getMealPlanTasksQuery)).
			WithArgs(interfaceToDriverValue(getMealPlanTaskByIDArgs)...).
			WillReturnRows(buildErroneousMockRow())

		actual, err := c.GetMealPlanTask(ctx, exampleMealPlanTask.ID)
		assert.Error(t, err)
		assert.Nil(t, actual)

		mock.AssertExpectationsForObjects(t, db)
	})
}

func TestQuerier_createMealPlanTask(T *testing.T) {
	T.Parallel()

	T.Run("with nil input", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		c, db := buildTestClient(t)

		db.ExpectBegin()
		tx, err := c.db.BeginTx(ctx, nil)
		require.NoError(t, err)
		require.NotNil(t, tx)

		actual, err := c.createMealPlanTask(ctx, tx, nil)
		assert.Error(t, err)
		assert.Nil(t, actual)

		mock.AssertExpectationsForObjects(t, db)
	})
}

func TestQuerier_CreateMealPlanTask(T *testing.T) {
	T.Parallel()

	T.Run("with nil input", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		c, db := buildTestClient(t)

		actual, err := c.CreateMealPlanTask(ctx, nil)
		assert.Error(t, err)
		assert.Nil(t, actual)

		mock.AssertExpectationsForObjects(t, db)
	})
}

func TestQuerier_GetMealPlanTasksForMealPlan(T *testing.T) {
	T.Parallel()

	T.Run("standard", func(t *testing.T) {
		t.Parallel()

		exampleMealPlanID := fakes.BuildFakeID()
		exampleMealPlanTasks := fakes.BuildFakeMealPlanTaskList()

		ctx := context.Background()
		c, db := buildTestClient(t)

		getMealPlanTaskByIDArgs := []any{
			exampleMealPlanID,
		}

		db.ExpectQuery(formatQueryForSQLMock(listMealPlanTasksForMealPlanQuery)).
			WithArgs(interfaceToDriverValue(getMealPlanTaskByIDArgs)...).
			WillReturnRows(buildMockRowsFromMealPlanTasksWithRecipePrepTaskSteps(false, 0, exampleMealPlanTasks.Data...))

		actual, err := c.GetMealPlanTasksForMealPlan(ctx, exampleMealPlanID)
		assert.NoError(t, err)
		assert.Equal(t, exampleMealPlanTasks.Data, actual)

		mock.AssertExpectationsForObjects(t, db)
	})

	T.Run("with missing meal plan ID", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		c, db := buildTestClient(t)

		actual, err := c.GetMealPlanTasksForMealPlan(ctx, "")
		assert.Error(t, err)
		assert.Nil(t, actual)

		mock.AssertExpectationsForObjects(t, db)
	})

	T.Run("with error executing query", func(t *testing.T) {
		t.Parallel()

		exampleMealPlanID := fakes.BuildFakeID()
		exampleMealPlanTasks := fakes.BuildFakeMealPlanTaskList()
		for i := range exampleMealPlanTasks.Data {
			exampleMealPlanTasks.Data[i].MealPlanOption = types.MealPlanOption{}
			exampleMealPlanTasks.Data[i].RecipePrepTask = types.RecipePrepTask{}
		}

		ctx := context.Background()
		c, db := buildTestClient(t)

		getMealPlanTaskByIDArgs := []any{
			exampleMealPlanID,
		}

		db.ExpectQuery(formatQueryForSQLMock(listMealPlanTasksForMealPlanQuery)).
			WithArgs(interfaceToDriverValue(getMealPlanTaskByIDArgs)...).
			WillReturnError(errors.New("blah"))

		actual, err := c.GetMealPlanTasksForMealPlan(ctx, exampleMealPlanID)
		assert.Error(t, err)
		assert.Nil(t, actual)

		mock.AssertExpectationsForObjects(t, db)
	})

	T.Run("with invalid response from database", func(t *testing.T) {
		t.Parallel()

		exampleMealPlanID := fakes.BuildFakeID()
		exampleMealPlanTasks := fakes.BuildFakeMealPlanTaskList()
		for i := range exampleMealPlanTasks.Data {
			exampleMealPlanTasks.Data[i].MealPlanOption = types.MealPlanOption{}
			exampleMealPlanTasks.Data[i].RecipePrepTask = types.RecipePrepTask{}
		}

		ctx := context.Background()
		c, db := buildTestClient(t)

		getMealPlanTaskByIDArgs := []any{
			exampleMealPlanID,
		}

		db.ExpectQuery(formatQueryForSQLMock(listMealPlanTasksForMealPlanQuery)).
			WithArgs(interfaceToDriverValue(getMealPlanTaskByIDArgs)...).
			WillReturnRows(buildErroneousMockRow())

		actual, err := c.GetMealPlanTasksForMealPlan(ctx, exampleMealPlanID)
		assert.Error(t, err)
		assert.Nil(t, actual)

		mock.AssertExpectationsForObjects(t, db)
	})
}

func TestQuerier_MarkMealPlanAsHavingTasksCreated(T *testing.T) {
	T.Parallel()

	T.Run("standard", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		c, db := buildTestClient(t)

		exampleMealPlan := fakes.BuildFakeMealPlan()

		markMealPlanOptionAsHavingStepsCreatedArgs := []any{
			exampleMealPlan.ID,
		}

		db.ExpectExec(formatQueryForSQLMock(markMealPlanAsHavingStepsCreatedQuery)).
			WithArgs(interfaceToDriverValue(markMealPlanOptionAsHavingStepsCreatedArgs)...).
			WillReturnResult(newArbitraryDatabaseResult())

		assert.NoError(t, c.MarkMealPlanAsHavingTasksCreated(ctx, exampleMealPlan.ID))

		mock.AssertExpectationsForObjects(t, db)
	})

	T.Run("with empty meal plan ID", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		c, db := buildTestClient(t)

		assert.Error(t, c.MarkMealPlanAsHavingTasksCreated(ctx, ""))

		mock.AssertExpectationsForObjects(t, db)
	})

	T.Run("with error executing query", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		c, db := buildTestClient(t)

		exampleMealPlan := fakes.BuildFakeMealPlan()

		markMealPlanOptionAsHavingStepsCreatedArgs := []any{
			exampleMealPlan.ID,
		}

		db.ExpectExec(formatQueryForSQLMock(markMealPlanAsHavingStepsCreatedQuery)).
			WithArgs(interfaceToDriverValue(markMealPlanOptionAsHavingStepsCreatedArgs)...).
			WillReturnError(errors.New("blah"))

		assert.Error(t, c.MarkMealPlanAsHavingTasksCreated(ctx, exampleMealPlan.ID))

		mock.AssertExpectationsForObjects(t, db)
	})
}

func TestQuerier_MarkMealPlanAsHavingGroceryListInitialized(T *testing.T) {
	T.Parallel()

	T.Run("with empty meal plan ID", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		c, db := buildTestClient(t)

		assert.Error(t, c.MarkMealPlanAsHavingGroceryListInitialized(ctx, ""))

		mock.AssertExpectationsForObjects(t, db)
	})
}

func TestQuerier_ChangeMealPlanTaskStatus(T *testing.T) {
	T.Parallel()

	T.Run("standard", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		c, _ := buildTestClient(t)

		assert.Error(t, c.ChangeMealPlanTaskStatus(ctx, nil))
	})
}
