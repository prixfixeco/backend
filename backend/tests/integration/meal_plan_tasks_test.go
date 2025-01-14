package integration

import (
	"testing"

	"github.com/dinnerdonebetter/backend/internal/observability/tracing"
	"github.com/dinnerdonebetter/backend/pkg/types"
	"github.com/dinnerdonebetter/backend/pkg/types/converters"
	"github.com/dinnerdonebetter/backend/pkg/types/fakes"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func checkMealPlanTaskEquality(t *testing.T, expected, actual *types.MealPlanTask) {
	t.Helper()

	assert.NotZero(t, actual.ID)
	assert.Equal(t, expected.CreationExplanation, actual.CreationExplanation, "expected CreationExplanation for meal plan %s to be %v, but it was %v", expected.CreationExplanation, expected.CreationExplanation, actual.CreationExplanation)
	assert.Equal(t, expected.Status, actual.Status, "expected Status for meal plan %s to be %v, but it was %v", expected.Status, expected.Status, actual.Status)
	assert.Equal(t, expected.StatusExplanation, actual.StatusExplanation, "expected StatusExplanation for meal plan %s to be %v, but it was %v", expected.StatusExplanation, expected.StatusExplanation, actual.StatusExplanation)
	assert.Equal(t, expected.AssignedToUser, actual.AssignedToUser, "expected AssignedToUser for meal plan %s to be %v, but it was %v", expected.AssignedToUser, expected.AssignedToUser, actual.AssignedToUser)
	assert.Equal(t, expected.CompletedAt, actual.CompletedAt, "expected CompletedAt for meal plan %s to be %v, but it was %v", expected.CompletedAt, expected.CompletedAt, actual.CompletedAt)

	assert.NotZero(t, actual.CreatedAt)
}

func (s *TestSuite) TestMealPlanTasks_CompleteLifecycle() {
	s.runTest("should CRUD", func(testClients *testClientWrapper) func() {
		return func() {
			t := s.T()

			ctx, span := tracing.StartCustomSpan(s.ctx, t.Name())
			defer span.End()

			createdMealPlan := createMealPlanForTest(ctx, t, nil, testClients.adminClient, testClients.userClient)

			exampleMealPlanTask := fakes.BuildFakeMealPlanTask()
			exampleMealPlanTaskInput := converters.ConvertMealPlanTaskToMealPlanTaskCreationRequestInput(exampleMealPlanTask)

			exampleMealPlanTaskInput.MealPlanOptionID = createdMealPlan.Events[0].Options[0].ID
			exampleMealPlanTaskInput.RecipePrepTaskID = createdMealPlan.Events[0].Options[0].Meal.Components[0].Recipe.PrepTasks[0].ID

			createdMealPlanTask, err := testClients.adminClient.CreateMealPlanTask(ctx, createdMealPlan.ID, exampleMealPlanTaskInput)
			require.NoError(t, err)
			checkMealPlanTaskEquality(t, exampleMealPlanTask, createdMealPlanTask)

			actual, err := testClients.adminClient.GetMealPlanTask(ctx, createdMealPlan.ID, createdMealPlanTask.ID)
			requireNotNilAndNoProblems(t, actual, err)

			// assert meal plan task equality
			checkMealPlanTaskEquality(t, exampleMealPlanTask, actual)
		}
	})
}
