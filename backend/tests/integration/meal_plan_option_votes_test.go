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

func checkMealPlanOptionVoteEquality(t *testing.T, expected, actual *types.MealPlanOptionVote) {
	t.Helper()

	assert.NotZero(t, actual.ID)
	assert.Equal(t, expected.Rank, actual.Rank, "expected Rank for meal plan option vote %s to be %v, but it was %v", expected.ID, expected.Rank, actual.Rank)
	assert.Equal(t, expected.Abstain, actual.Abstain, "expected Abstain for meal plan option vote %s to be %v, but it was %v", expected.ID, expected.Abstain, actual.Abstain)
	assert.Equal(t, expected.Notes, actual.Notes, "expected StatusExplanation for meal plan option vote %s to be %v, but it was %v", expected.ID, expected.Notes, actual.Notes)
	assert.NotZero(t, actual.CreatedAt)
}

func (s *TestSuite) TestMealPlanOptionVotes_CompleteLifecycle() {
	s.runTest("should CRUD", func(testClients *testClientWrapper) func() {
		return func() {
			t := s.T()

			ctx, span := tracing.StartCustomSpan(s.ctx, t.Name())
			defer span.End()
			createdMealPlan := createMealPlanForTest(ctx, t, nil, testClients.adminClient, testClients.userClient)

			require.NotEmpty(t, createdMealPlan.Events)
			require.NotEmpty(t, createdMealPlan.Events[0].Options)

			createdMealPlanEvent := createdMealPlan.Events[0]
			createdMealPlanOption := createdMealPlanEvent.Options[0]
			require.NotNil(t, createdMealPlanOption)

			exampleMealPlanOptionVote := fakes.BuildFakeMealPlanOptionVote()
			exampleMealPlanOptionVote.BelongsToMealPlanOption = createdMealPlanOption.ID
			exampleMealPlanOptionVoteInput := converters.ConvertMealPlanOptionVoteToMealPlanOptionVoteCreationRequestInput(exampleMealPlanOptionVote)
			createdMealPlanOptionVotes, err := testClients.userClient.CreateMealPlanOptionVote(ctx, createdMealPlan.ID, createdMealPlanEvent.ID, exampleMealPlanOptionVoteInput)
			require.NoError(t, err)

			for _, createdMealPlanOptionVote := range createdMealPlanOptionVotes {
				checkMealPlanOptionVoteEquality(t, exampleMealPlanOptionVote, createdMealPlanOptionVote)

				createdMealPlanOptionVote, err = testClients.userClient.GetMealPlanOptionVote(ctx, createdMealPlan.ID, createdMealPlanEvent.ID, createdMealPlanOption.ID, createdMealPlanOptionVote.ID)
				requireNotNilAndNoProblems(t, createdMealPlanOptionVote, err)
				require.Equal(t, createdMealPlanOption.ID, createdMealPlanOptionVote.BelongsToMealPlanOption)

				checkMealPlanOptionVoteEquality(t, exampleMealPlanOptionVote, createdMealPlanOptionVote)

				newMealPlanOptionVote := fakes.BuildFakeMealPlanOptionVote()
				updateInput := converters.ConvertMealPlanOptionVoteToMealPlanOptionVoteUpdateRequestInput(newMealPlanOptionVote)
				createdMealPlanOptionVote.Update(updateInput)
				assert.NoError(t, testClients.userClient.UpdateMealPlanOptionVote(ctx, createdMealPlan.ID, createdMealPlanEvent.ID, createdMealPlanOption.ID, createdMealPlanOptionVote.ID, updateInput))

				actual, err := testClients.userClient.GetMealPlanOptionVote(ctx, createdMealPlan.ID, createdMealPlanEvent.ID, createdMealPlanOption.ID, createdMealPlanOptionVote.ID)
				requireNotNilAndNoProblems(t, actual, err)

				// assert meal plan option vote equality
				checkMealPlanOptionVoteEquality(t, newMealPlanOptionVote, actual)
				assert.NotNil(t, actual.LastUpdatedAt)

				assert.NoError(t, testClients.userClient.ArchiveMealPlanOptionVote(ctx, createdMealPlan.ID, createdMealPlanEvent.ID, createdMealPlanOption.ID, createdMealPlanOptionVote.ID))
			}

			require.NoError(t, testClients.userClient.ArchiveMealPlanOption(ctx, createdMealPlan.ID, createdMealPlanEvent.ID, createdMealPlanOption.ID))

			require.NoError(t, testClients.userClient.ArchiveMealPlanEvent(ctx, createdMealPlan.ID, createdMealPlanEvent.ID))

			require.NoError(t, testClients.userClient.ArchiveMealPlan(ctx, createdMealPlan.ID))
		}
	})
}

func (s *TestSuite) TestMealPlanOptionVotes_Listing() {
	s.runTest("should be readable in paginated form", func(testClients *testClientWrapper) func() {
		return func() {
			t := s.T()

			ctx, span := tracing.StartCustomSpan(s.ctx, t.Name())
			defer span.End()

			createdMealPlan := createMealPlanForTest(ctx, t, nil, testClients.adminClient, testClients.userClient)

			require.NotEmpty(t, createdMealPlan.Events)
			require.NotEmpty(t, createdMealPlan.Events[0].Options)

			createdMealPlanEvent := createdMealPlan.Events[0]
			createdMealPlanOption := createdMealPlanEvent.Options[0]
			require.NotNil(t, createdMealPlanOption)

			exampleMealPlanOptionVote := fakes.BuildFakeMealPlanOptionVote()
			exampleMealPlanOptionVote.BelongsToMealPlanOption = createdMealPlanOption.ID
			exampleMealPlanOptionVoteInput := converters.ConvertMealPlanOptionVoteToMealPlanOptionVoteCreationRequestInput(exampleMealPlanOptionVote)
			createdMealPlanOptionVotes, err := testClients.userClient.CreateMealPlanOptionVote(ctx, createdMealPlan.ID, createdMealPlanEvent.ID, exampleMealPlanOptionVoteInput)
			require.NoError(t, err)

			for _, createdMealPlanOptionVote := range createdMealPlanOptionVotes {
				checkMealPlanOptionVoteEquality(t, exampleMealPlanOptionVote, createdMealPlanOptionVote)

				createdMealPlanOptionVote, err = testClients.userClient.GetMealPlanOptionVote(ctx, createdMealPlan.ID, createdMealPlanEvent.ID, createdMealPlanOption.ID, createdMealPlanOptionVote.ID)
				requireNotNilAndNoProblems(t, createdMealPlanOptionVote, err)
				require.Equal(t, createdMealPlanOption.ID, createdMealPlanOptionVote.BelongsToMealPlanOption)

				checkMealPlanOptionVoteEquality(t, exampleMealPlanOptionVote, createdMealPlanOptionVote)

				// assert meal plan option vote list equality
				actual, err := testClients.userClient.GetMealPlanOptionVotes(ctx, createdMealPlan.ID, createdMealPlanEvent.ID, createdMealPlanOption.ID, nil)
				requireNotNilAndNoProblems(t, actual, err)
				assert.NotEmpty(t, actual.Data)

				assert.NoError(t, testClients.userClient.ArchiveMealPlanOptionVote(ctx, createdMealPlan.ID, createdMealPlanEvent.ID, createdMealPlanOption.ID, createdMealPlanOptionVote.ID))
			}

			assert.NoError(t, testClients.userClient.ArchiveMealPlanOption(ctx, createdMealPlan.ID, createdMealPlanEvent.ID, createdMealPlanOption.ID))

			require.NoError(t, testClients.userClient.ArchiveMealPlanEvent(ctx, createdMealPlan.ID, createdMealPlanEvent.ID))

			assert.NoError(t, testClients.userClient.ArchiveMealPlan(ctx, createdMealPlan.ID))
		}
	})
}
