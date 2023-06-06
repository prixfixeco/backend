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

func checkRecipeRatingEquality(t *testing.T, expected, actual *types.RecipeRating) {
	t.Helper()

	assert.NotZero(t, actual.ID)
	assert.Equal(t, expected.Notes, actual.Notes, "expected Notes for recipe rating %s to be %v, but it was %v", expected.ID, expected.Notes, actual.Notes)
	assert.Equal(t, expected.RecipeID, actual.RecipeID, "expected RecipeID for recipe rating %s to be %v, but it was %v", expected.ID, expected.RecipeID, actual.RecipeID)
	assert.Equal(t, expected.Taste, actual.Taste, "expected Taste for recipe rating %s to be %v, but it was %v", expected.ID, expected.Taste, actual.Taste)
	assert.Equal(t, expected.Instructions, actual.Instructions, "expected Instructions for recipe rating %s to be %v, but it was %v", expected.ID, expected.Instructions, actual.Instructions)
	assert.Equal(t, expected.Overall, actual.Overall, "expected Overall for recipe rating %s to be %v, but it was %v", expected.ID, expected.Overall, actual.Overall)
	assert.Equal(t, expected.Cleanup, actual.Cleanup, "expected Cleanup for recipe rating %s to be %v, but it was %v", expected.ID, expected.Cleanup, actual.Cleanup)
	assert.Equal(t, expected.Difficulty, actual.Difficulty, "expected Difficulty for recipe rating %s to be %v, but it was %v", expected.ID, expected.Difficulty, actual.Difficulty)
	assert.NotZero(t, actual.CreatedAt)
}

func (s *TestSuite) TestRecipeRatings_CompleteLifecycle() {
	s.runForEachClient("should be creatable and readable and updatable and deletable", func(testClients *testClientWrapper) func() {
		return func() {
			t := s.T()

			ctx, span := tracing.StartCustomSpan(s.ctx, t.Name())
			defer span.End()

			_, _, createdRecipe := createRecipeForTest(ctx, t, testClients.admin, testClients.user, nil)

			t.Log("creating recipe rating")
			exampleRecipeRating := fakes.BuildFakeRecipeRating()
			exampleRecipeRating.RecipeID = createdRecipe.ID
			exampleRecipeRatingInput := converters.ConvertRecipeRatingToRecipeRatingCreationRequestInput(exampleRecipeRating)
			createdRecipeRating, err := testClients.user.CreateRecipeRating(ctx, createdRecipe.ID, exampleRecipeRatingInput)
			require.NoError(t, err)
			t.Logf("recipe rating %q created", createdRecipeRating.ID)
			checkRecipeRatingEquality(t, exampleRecipeRating, createdRecipeRating)

			createdRecipeRating, err = testClients.user.GetRecipeRating(ctx, createdRecipe.ID, createdRecipeRating.ID)
			requireNotNilAndNoProblems(t, createdRecipeRating, err)
			checkRecipeRatingEquality(t, exampleRecipeRating, createdRecipeRating)

			t.Log("changing recipe rating")
			newRecipeRating := fakes.BuildFakeRecipeRating()
			newRecipeRating.RecipeID = createdRecipe.ID
			createdRecipeRating.Update(converters.ConvertRecipeRatingToRecipeRatingUpdateRequestInput(newRecipeRating))
			assert.NoError(t, testClients.admin.UpdateRecipeRating(ctx, createdRecipeRating))

			t.Log("fetching changed recipe rating")
			actual, err := testClients.admin.GetRecipeRating(ctx, createdRecipe.ID, createdRecipeRating.ID)
			requireNotNilAndNoProblems(t, actual, err)

			// assert recipe rating equality
			checkRecipeRatingEquality(t, newRecipeRating, actual)
			assert.NotNil(t, actual.LastUpdatedAt)

			t.Log("cleaning up recipe rating")
			assert.NoError(t, testClients.admin.ArchiveRecipeRating(ctx, createdRecipe.ID, createdRecipeRating.ID))
		}
	})
}

func (s *TestSuite) TestRecipeRatings_Listing() {
	s.runForEachClient("should be readable in paginated form", func(testClients *testClientWrapper) func() {
		return func() {
			t := s.T()

			ctx, span := tracing.StartCustomSpan(s.ctx, t.Name())
			defer span.End()

			_, _, createdRecipe := createRecipeForTest(ctx, t, testClients.admin, testClients.user, nil)

			t.Log("creating recipe rating")
			exampleRecipeRating := fakes.BuildFakeRecipeRating()
			exampleRecipeRating.RecipeID = createdRecipe.ID
			exampleRecipeRatingInput := converters.ConvertRecipeRatingToRecipeRatingCreationRequestInput(exampleRecipeRating)
			createdRecipeRating, err := testClients.user.CreateRecipeRating(ctx, createdRecipe.ID, exampleRecipeRatingInput)
			require.NoError(t, err)
			t.Logf("recipe rating %q created", createdRecipeRating.ID)
			checkRecipeRatingEquality(t, exampleRecipeRating, createdRecipeRating)

			// assert recipe rating list equality
			actual, err := testClients.admin.GetRecipeRatings(ctx, createdRecipe.ID, nil)
			requireNotNilAndNoProblems(t, actual, err)
			assert.Equal(t, len(actual.Data), 1, "expected %d to be <= %d", len(actual.Data), 1)

			assert.NoError(t, testClients.admin.ArchiveRecipeRating(ctx, createdRecipe.ID, createdRecipeRating.ID))
		}
	})
}