package postgres

import (
	"context"
	"database/sql"
	"testing"

	"github.com/dinnerdonebetter/backend/pkg/types"
	"github.com/dinnerdonebetter/backend/pkg/types/converters"
	"github.com/dinnerdonebetter/backend/pkg/types/fakes"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func buildMealPlanForIntegrationTest(userID string, meal *types.Meal) *types.MealPlan {
	exampleMealPlan := fakes.BuildFakeMealPlan()
	exampleMealPlan.CreatedByUser = userID

	exampleMealPlan.Events = []*types.MealPlanEvent{
		buildMealPlanEventForIntegrationTest(meal),
	}

	// only one event means it's immediately finalized
	exampleMealPlan.Status = string(types.MealPlanStatusFinalized)

	return exampleMealPlan
}

func createMealPlanForTest(t *testing.T, ctx context.Context, exampleMealPlan *types.MealPlan, dbc *Querier) *types.MealPlan {
	t.Helper()

	// create
	dbInput := converters.ConvertMealPlanToMealPlanDatabaseCreationInput(exampleMealPlan)

	created, err := dbc.CreateMealPlan(ctx, dbInput)
	assert.NoError(t, err)
	require.NotNil(t, created)

	exampleMealPlan.CreatedAt = created.CreatedAt
	for i := range created.Events {
		assert.Equal(t, created.Events[i].ID, exampleMealPlan.Events[i].ID)
		exampleMealPlan.Events[i].CreatedAt = created.Events[i].CreatedAt
		exampleMealPlan.Events[i].StartsAt = created.Events[i].StartsAt
		exampleMealPlan.Events[i].EndsAt = created.Events[i].EndsAt
		exampleMealPlan.Events[i].BelongsToMealPlan = created.Events[i].BelongsToMealPlan

		for j := range created.Events[i].Options {
			assert.Equal(t, created.Events[i].Options[j].ID, exampleMealPlan.Events[i].Options[j].ID)
			assert.Equal(t, created.Events[i].Options[j].Meal.ID, exampleMealPlan.Events[i].Options[j].Meal.ID)
			exampleMealPlan.Events[i].Options[j] = created.Events[i].Options[j]
		}

		assert.Equal(t, created.Events[i].Options, exampleMealPlan.Events[i].Options)
	}
	assert.Equal(t, exampleMealPlan, created)

	mealPlan, err := dbc.GetMealPlan(ctx, created.ID, created.BelongsToHousehold)
	require.NoError(t, err)

	exampleMealPlan.CreatedAt = mealPlan.CreatedAt
	exampleMealPlan.VotingDeadline = mealPlan.VotingDeadline
	for i := range mealPlan.Events {
		assert.Equal(t, mealPlan.Events[i].ID, exampleMealPlan.Events[i].ID)
		exampleMealPlan.Events[i].CreatedAt = mealPlan.Events[i].CreatedAt
		exampleMealPlan.Events[i].StartsAt = mealPlan.Events[i].StartsAt
		exampleMealPlan.Events[i].EndsAt = mealPlan.Events[i].EndsAt
		exampleMealPlan.Events[i].BelongsToMealPlan = mealPlan.Events[i].BelongsToMealPlan

		for j := range mealPlan.Events[i].Options {
			assert.Equal(t, mealPlan.Events[i].Options[j].ID, exampleMealPlan.Events[i].Options[j].ID)
			assert.Equal(t, mealPlan.Events[i].Options[j].Meal.ID, exampleMealPlan.Events[i].Options[j].Meal.ID)
			exampleMealPlan.Events[i].Options[j] = mealPlan.Events[i].Options[j]
		}

		assert.Equal(t, mealPlan.Events[i].Options, exampleMealPlan.Events[i].Options)
	}

	require.Equal(t, exampleMealPlan.CreatedAt, mealPlan.CreatedAt)
	require.Equal(t, exampleMealPlan.VotingDeadline, mealPlan.VotingDeadline)
	require.Equal(t, exampleMealPlan.ArchivedAt, mealPlan.ArchivedAt)
	require.Equal(t, exampleMealPlan.LastUpdatedAt, mealPlan.LastUpdatedAt)
	require.Equal(t, exampleMealPlan.ID, mealPlan.ID)
	require.Equal(t, exampleMealPlan.Status, mealPlan.Status)
	require.Equal(t, exampleMealPlan.Notes, mealPlan.Notes)
	require.Equal(t, exampleMealPlan.ElectionMethod, mealPlan.ElectionMethod)
	require.Equal(t, exampleMealPlan.BelongsToHousehold, mealPlan.BelongsToHousehold)
	require.Equal(t, exampleMealPlan.CreatedByUser, mealPlan.CreatedByUser)
	require.Equal(t, exampleMealPlan.GroceryListInitialized, mealPlan.GroceryListInitialized)
	require.Equal(t, exampleMealPlan.TasksCreated, mealPlan.TasksCreated)

	for i := range mealPlan.Events {
		require.Equal(t, exampleMealPlan.Events[i].CreatedAt, mealPlan.Events[i].CreatedAt)
		require.Equal(t, exampleMealPlan.Events[i].StartsAt, mealPlan.Events[i].StartsAt)
		require.Equal(t, exampleMealPlan.Events[i].EndsAt, mealPlan.Events[i].EndsAt)
		require.Equal(t, exampleMealPlan.Events[i].ArchivedAt, mealPlan.Events[i].ArchivedAt)
		require.Equal(t, exampleMealPlan.Events[i].LastUpdatedAt, mealPlan.Events[i].LastUpdatedAt)
		require.Equal(t, exampleMealPlan.Events[i].MealName, mealPlan.Events[i].MealName)
		require.Equal(t, exampleMealPlan.Events[i].Notes, mealPlan.Events[i].Notes)
		require.Equal(t, exampleMealPlan.Events[i].BelongsToMealPlan, mealPlan.Events[i].BelongsToMealPlan)
		require.Equal(t, exampleMealPlan.Events[i].ID, mealPlan.Events[i].ID)

		for j := range mealPlan.Events[i].Options {
			require.Equal(t, exampleMealPlan.Events[i].Options[j].CreatedAt, mealPlan.Events[i].Options[j].CreatedAt)
			require.Equal(t, exampleMealPlan.Events[i].Options[j].LastUpdatedAt, mealPlan.Events[i].Options[j].LastUpdatedAt)
			require.Equal(t, exampleMealPlan.Events[i].Options[j].AssignedCook, mealPlan.Events[i].Options[j].AssignedCook)
			require.Equal(t, exampleMealPlan.Events[i].Options[j].ArchivedAt, mealPlan.Events[i].Options[j].ArchivedAt)
			require.Equal(t, exampleMealPlan.Events[i].Options[j].AssignedDishwasher, mealPlan.Events[i].Options[j].AssignedDishwasher)
			require.Equal(t, exampleMealPlan.Events[i].Options[j].Notes, mealPlan.Events[i].Options[j].Notes)
			require.Equal(t, exampleMealPlan.Events[i].Options[j].BelongsToMealPlanEvent, mealPlan.Events[i].Options[j].BelongsToMealPlanEvent)
			require.Equal(t, exampleMealPlan.Events[i].Options[j].ID, mealPlan.Events[i].Options[j].ID)
			require.Equal(t, exampleMealPlan.Events[i].Options[j].Votes, mealPlan.Events[i].Options[j].Votes)
			require.Equal(t, exampleMealPlan.Events[i].Options[j].Meal, mealPlan.Events[i].Options[j].Meal)
			require.Equal(t, exampleMealPlan.Events[i].Options[j].MealScale, mealPlan.Events[i].Options[j].MealScale)
			require.Equal(t, exampleMealPlan.Events[i].Options[j].Chosen, mealPlan.Events[i].Options[j].Chosen)
			require.Equal(t, exampleMealPlan.Events[i].Options[j].TieBroken, mealPlan.Events[i].Options[j].TieBroken)
		}
	}

	assert.Equal(t, exampleMealPlan, mealPlan)

	return mealPlan
}

func TestQuerier_Integration_MealPlans(t *testing.T) {
	if !runningContainerTests {
		t.SkipNow()
	}

	ctx := context.Background()
	dbc, container := buildDatabaseClientForTest(t, ctx)

	databaseURI, err := container.ConnectionString(ctx)
	require.NoError(t, err)
	require.NotEmpty(t, databaseURI)

	defer func(t *testing.T) {
		t.Helper()
		assert.NoError(t, container.Terminate(ctx))
	}(t)

	user := createUserForTest(t, ctx, nil, dbc)
	householdID, err := dbc.GetDefaultHouseholdIDForUser(ctx, user.ID)
	require.NoError(t, err)
	require.NotEmpty(t, householdID)

	recipe := createRecipeForTest(t, ctx, nil, dbc, true)
	meal := createMealForTest(t, ctx, buildMealForIntegrationTest(user.ID, recipe), dbc)

	exampleMealPlan := buildMealPlanForIntegrationTest(user.ID, meal)
	exampleMealPlan.BelongsToHousehold = householdID
	createdMealPlans := []*types.MealPlan{}

	// create
	createdMealPlans = append(createdMealPlans, createMealPlanForTest(t, ctx, exampleMealPlan, dbc))

	// create more
	for i := 0; i < exampleQuantity; i++ {
		input := buildMealPlanForIntegrationTest(user.ID, meal)
		input.BelongsToHousehold = householdID
		createdMealPlans = append(createdMealPlans, createMealPlanForTest(t, ctx, input, dbc))
	}

	// fetch as list
	mealPlans, err := dbc.GetMealPlans(ctx, householdID, nil)
	assert.NoError(t, err)
	assert.NotEmpty(t, mealPlans.Data)
	assert.Equal(t, len(createdMealPlans), len(mealPlans.Data))

	// delete
	for _, mealPlan := range createdMealPlans {
		assert.NoError(t, dbc.ArchiveMealPlan(ctx, mealPlan.ID, householdID))

		var exists bool
		exists, err = dbc.MealPlanExists(ctx, mealPlan.ID, householdID)
		assert.NoError(t, err)
		assert.False(t, exists)

		var y *types.MealPlan
		y, err = dbc.GetMealPlan(ctx, mealPlan.ID, householdID)
		assert.Nil(t, y)
		assert.Error(t, err)
		assert.ErrorIs(t, err, sql.ErrNoRows)
	}
}
