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

func checkMealPlanEventEquality(t *testing.T, expected, actual *types.MealPlanEvent) {
	t.Helper()

	assert.NotZero(t, actual.ID)
	assert.Equal(t, expected.Notes, actual.Notes, "expected StatusExplanation for meal plan event %s to be %v, but it was %v", expected.ID, expected.Notes, actual.Notes)
	assert.Equal(t, expected.StartsAt, actual.StartsAt, "expected StartsAt for meal plan event %s to be %v, but it was %v", expected.ID, expected.StartsAt, actual.StartsAt)
	assert.Equal(t, expected.EndsAt, actual.EndsAt, "expected EndsAt for meal plan event %s to be %v, but it was %v", expected.ID, expected.EndsAt, actual.EndsAt)
	assert.Equal(t, expected.MealName, actual.MealName, "expected MealName for meal plan event %s to be %v, but it was %v", expected.ID, expected.MealName, actual.MealName)
	assert.Equal(t, expected.BelongsToMealPlan, actual.BelongsToMealPlan, "expected BelongsToMealPlan for meal plan event %s to be %v, but it was %v", expected.ID, expected.BelongsToMealPlan, actual.BelongsToMealPlan)
	assert.NotZero(t, actual.CreatedAt)
}

func (s *TestSuite) TestMealPlanEvents_CompleteLifecycle() {
	s.runTest("should CRUD", func(testClients *testClientWrapper) func() {
		return func() {
			t := s.T()

			ctx, span := tracing.StartCustomSpan(s.ctx, t.Name())
			defer span.End()

			createdMealPlan := createMealPlanForTest(ctx, t, nil, testClients.adminClient, testClients.userClient)

			require.NotNil(t, createdMealPlan)
			require.NotEmpty(t, createdMealPlan.Events)
			require.NotNil(t, createdMealPlan.Events[0])
			createdMealPlanEvent := createdMealPlan.Events[0]

			newMealPlanEvent := fakes.BuildFakeMealPlanEvent()
			newMealPlanEvent.BelongsToMealPlan = createdMealPlan.ID

			updateInput := converters.ConvertMealPlanEventToMealPlanEventUpdateRequestInput(newMealPlanEvent)
			createdMealPlanEvent.Update(updateInput)
			assert.NoError(t, testClients.userClient.UpdateMealPlanEvent(ctx, createdMealPlan.ID, createdMealPlanEvent.ID, updateInput))

			actual, err := testClients.userClient.GetMealPlanEvent(ctx, createdMealPlan.ID, createdMealPlanEvent.ID)
			requireNotNilAndNoProblems(t, actual, err)

			// assert meal plan event equality
			checkMealPlanEventEquality(t, newMealPlanEvent, actual)
			assert.NotNil(t, actual.LastUpdatedAt)

			assert.NoError(t, testClients.userClient.ArchiveMealPlanEvent(ctx, createdMealPlan.ID, createdMealPlanEvent.ID))

			assert.NoError(t, testClients.userClient.ArchiveMealPlan(ctx, createdMealPlan.ID))
		}
	})
}

func (s *TestSuite) TestMealPlanEvents_Listing() {
	s.runTest("should be readable in paginated form", func(testClients *testClientWrapper) func() {
		return func() {
			t := s.T()

			ctx, span := tracing.StartCustomSpan(s.ctx, t.Name())
			defer span.End()

			createdMealPlan := createMealPlanForTest(ctx, t, nil, testClients.adminClient, testClients.userClient)

			var expected []*types.MealPlanEvent
			for i := 0; i < 5; i++ {
				exampleMealPlanEvent := fakes.BuildFakeMealPlanEvent()
				exampleMealPlanEvent.Options = nil
				exampleMealPlanEvent.BelongsToMealPlan = createdMealPlan.ID

				exampleMealPlanEventInput := converters.ConvertMealPlanEventToMealPlanEventCreationRequestInput(exampleMealPlanEvent)
				createdMealPlanEvent, err := testClients.userClient.CreateMealPlanEvent(ctx, createdMealPlan.ID, exampleMealPlanEventInput)
				require.NoError(t, err)

				checkMealPlanEventEquality(t, exampleMealPlanEvent, createdMealPlanEvent)

				createdMealPlanEvent, err = testClients.userClient.GetMealPlanEvent(ctx, createdMealPlan.ID, createdMealPlanEvent.ID)
				requireNotNilAndNoProblems(t, createdMealPlanEvent, err)
				require.Equal(t, createdMealPlan.ID, createdMealPlanEvent.BelongsToMealPlan)

				expected = append(expected, createdMealPlanEvent)
			}

			// assert meal plan event list equality
			actual, err := testClients.userClient.GetMealPlanEvents(ctx, createdMealPlan.ID, nil)
			requireNotNilAndNoProblems(t, actual, err)
			assert.True(
				t,
				len(expected) <= len(actual.Data),
				"expected %d to be <= %d",
				len(expected),
				len(actual.Data),
			)

			for _, createdMealPlanEvent := range expected {
				assert.NoError(t, testClients.userClient.ArchiveMealPlanEvent(ctx, createdMealPlan.ID, createdMealPlanEvent.ID))
			}

			assert.NoError(t, testClients.userClient.ArchiveMealPlan(ctx, createdMealPlan.ID))
		}
	})
}
