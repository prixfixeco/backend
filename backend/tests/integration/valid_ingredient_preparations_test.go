package integration

import (
	"fmt"
	"testing"

	"github.com/dinnerdonebetter/backend/internal/observability/tracing"
	"github.com/dinnerdonebetter/backend/pkg/types"
	"github.com/dinnerdonebetter/backend/pkg/types/converters"
	"github.com/dinnerdonebetter/backend/pkg/types/fakes"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func checkValidIngredientPreparationEquality(t *testing.T, expected, actual *types.ValidIngredientPreparation) {
	t.Helper()

	assert.NotZero(t, actual.ID)
	assert.Equal(t, expected.Notes, actual.Notes, "expected StatusExplanation for valid ingredient preparation %s to be %v, but it was %v", expected.ID, expected.Notes, actual.Notes)
	assert.Equal(t, expected.Preparation.ID, actual.Preparation.ID, "expected Preparation for valid ingredient preparation %s to be %v, but it was %v", expected.ID, expected.Preparation.ID, actual.Preparation.ID)
	assert.Equal(t, expected.Ingredient.ID, actual.Ingredient.ID, "expected Ingredient for valid ingredient preparation %s to be %v, but it was %v", expected.ID, expected.Ingredient.ID, actual.Ingredient.ID)
	assert.NotZero(t, actual.CreatedAt)
}

func (s *TestSuite) TestValidIngredientPreparations_CompleteLifecycle() {
	s.runTest("should CRUD", func(testClients *testClientWrapper) func() {
		return func() {
			t := s.T()

			ctx, span := tracing.StartCustomSpan(s.ctx, t.Name())
			defer span.End()

			exampleValidPreparation := fakes.BuildFakeValidPreparation()
			exampleValidPreparationInput := converters.ConvertValidPreparationToValidPreparationCreationRequestInput(exampleValidPreparation)
			createdValidPreparation, err := testClients.adminClient.CreateValidPreparation(ctx, exampleValidPreparationInput)
			require.NoError(t, err)

			checkValidPreparationEquality(t, exampleValidPreparation, createdValidPreparation)

			createdValidPreparation, err = testClients.userClient.GetValidPreparation(ctx, createdValidPreparation.ID)
			requireNotNilAndNoProblems(t, createdValidPreparation, err)
			checkValidPreparationEquality(t, exampleValidPreparation, createdValidPreparation)

			exampleValidIngredient := fakes.BuildFakeValidIngredient()
			exampleValidIngredientInput := converters.ConvertValidIngredientToValidIngredientCreationRequestInput(exampleValidIngredient)
			createdValidIngredient, err := testClients.adminClient.CreateValidIngredient(ctx, exampleValidIngredientInput)
			require.NoError(t, err)

			checkValidIngredientEquality(t, exampleValidIngredient, createdValidIngredient)

			createdValidIngredient, err = testClients.userClient.GetValidIngredient(ctx, createdValidIngredient.ID)
			requireNotNilAndNoProblems(t, createdValidIngredient, err)
			checkValidIngredientEquality(t, exampleValidIngredient, createdValidIngredient)

			exampleValidIngredientPreparation := fakes.BuildFakeValidIngredientPreparation()
			exampleValidIngredientPreparation.Ingredient = *createdValidIngredient
			exampleValidIngredientPreparation.Preparation = *createdValidPreparation
			exampleValidIngredientPreparationInput := converters.ConvertValidIngredientPreparationToValidIngredientPreparationCreationRequestInput(exampleValidIngredientPreparation)
			createdValidIngredientPreparation, err := testClients.adminClient.CreateValidIngredientPreparation(ctx, exampleValidIngredientPreparationInput)
			require.NoError(t, err)

			checkValidIngredientPreparationEquality(t, exampleValidIngredientPreparation, createdValidIngredientPreparation)

			createdValidIngredientPreparation, err = testClients.userClient.GetValidIngredientPreparation(ctx, createdValidIngredientPreparation.ID)
			requireNotNilAndNoProblems(t, createdValidIngredientPreparation, err)

			checkValidIngredientPreparationEquality(t, exampleValidIngredientPreparation, createdValidIngredientPreparation)

			newValidIngredientPreparation := fakes.BuildFakeValidIngredientPreparation()
			newValidIngredientPreparation.Ingredient = *createdValidIngredient
			newValidIngredientPreparation.Preparation = *createdValidPreparation

			updateInput := converters.ConvertValidIngredientPreparationToValidIngredientPreparationUpdateRequestInput(newValidIngredientPreparation)
			createdValidIngredientPreparation.Update(updateInput)
			assert.NoError(t, testClients.adminClient.UpdateValidIngredientPreparation(ctx, createdValidIngredientPreparation.ID, updateInput))

			actual, err := testClients.userClient.GetValidIngredientPreparation(ctx, createdValidIngredientPreparation.ID)
			requireNotNilAndNoProblems(t, actual, err)

			// assert valid ingredient preparation equality
			checkValidIngredientPreparationEquality(t, newValidIngredientPreparation, actual)
			assert.NotNil(t, actual.LastUpdatedAt)

			assert.NoError(t, testClients.adminClient.ArchiveValidIngredientPreparation(ctx, createdValidIngredientPreparation.ID))

			assert.NoError(t, testClients.adminClient.ArchiveValidIngredient(ctx, createdValidIngredient.ID))

			assert.NoError(t, testClients.adminClient.ArchiveValidPreparation(ctx, createdValidPreparation.ID))
		}
	})
}

func (s *TestSuite) TestValidIngredientPreparations_Listing() {
	s.runTest("should be readable in paginated form", func(testClients *testClientWrapper) func() {
		return func() {
			t := s.T()

			ctx, span := tracing.StartCustomSpan(s.ctx, t.Name())
			defer span.End()

			var expected []*types.ValidIngredientPreparation
			for i := 0; i < 5; i++ {
				exampleValidPreparation := fakes.BuildFakeValidPreparation()
				exampleValidPreparationInput := converters.ConvertValidPreparationToValidPreparationCreationRequestInput(exampleValidPreparation)
				createdValidPreparation, err := testClients.adminClient.CreateValidPreparation(ctx, exampleValidPreparationInput)
				require.NoError(t, err)

				checkValidPreparationEquality(t, exampleValidPreparation, createdValidPreparation)

				createdValidPreparation, err = testClients.userClient.GetValidPreparation(ctx, createdValidPreparation.ID)
				requireNotNilAndNoProblems(t, createdValidPreparation, err)
				checkValidPreparationEquality(t, exampleValidPreparation, createdValidPreparation)

				exampleValidIngredient := fakes.BuildFakeValidIngredient()
				exampleValidIngredientInput := converters.ConvertValidIngredientToValidIngredientCreationRequestInput(exampleValidIngredient)
				createdValidIngredient, err := testClients.adminClient.CreateValidIngredient(ctx, exampleValidIngredientInput)
				require.NoError(t, err)

				checkValidIngredientEquality(t, exampleValidIngredient, createdValidIngredient)

				createdValidIngredient, err = testClients.userClient.GetValidIngredient(ctx, createdValidIngredient.ID)
				requireNotNilAndNoProblems(t, createdValidIngredient, err)
				checkValidIngredientEquality(t, exampleValidIngredient, createdValidIngredient)

				exampleValidIngredientPreparation := fakes.BuildFakeValidIngredientPreparation()
				exampleValidIngredientPreparation.Ingredient = *createdValidIngredient
				exampleValidIngredientPreparation.Preparation = *createdValidPreparation
				exampleValidIngredientPreparationInput := converters.ConvertValidIngredientPreparationToValidIngredientPreparationCreationRequestInput(exampleValidIngredientPreparation)
				createdValidIngredientPreparation, createdValidIngredientPreparationErr := testClients.adminClient.CreateValidIngredientPreparation(ctx, exampleValidIngredientPreparationInput)
				require.NoError(t, createdValidIngredientPreparationErr)

				exampleValidIngredientPreparation.Ingredient = types.ValidIngredient{ID: createdValidIngredient.ID}
				exampleValidIngredientPreparation.Preparation = types.ValidPreparation{ID: createdValidPreparation.ID}

				checkValidIngredientPreparationEquality(t, exampleValidIngredientPreparation, createdValidIngredientPreparation)

				expected = append(expected, createdValidIngredientPreparation)
			}

			// assert valid ingredient preparation list equality
			actual, err := testClients.userClient.GetValidIngredientPreparations(ctx, nil)
			requireNotNilAndNoProblems(t, actual, err)
			assert.True(
				t,
				len(expected) <= len(actual.Data),
				"expected %d to be <= %d",
				len(expected),
				len(actual.Data),
			)

			for _, createdValidIngredientPreparation := range expected {
				assert.NoError(t, testClients.adminClient.ArchiveValidIngredientPreparation(ctx, createdValidIngredientPreparation.ID))
			}
		}
	})
}

func (s *TestSuite) TestValidIngredientPreparations_Listing_ByValues() {
	s.runTest("should be findable via either member of the bridge type", func(testClients *testClientWrapper) func() {
		return func() {
			t := s.T()

			ctx, span := tracing.StartCustomSpan(s.ctx, t.Name())
			defer span.End()

			exampleValidPreparation := fakes.BuildFakeValidPreparation()
			exampleValidPreparationInput := converters.ConvertValidPreparationToValidPreparationCreationRequestInput(exampleValidPreparation)
			createdValidPreparation, err := testClients.adminClient.CreateValidPreparation(ctx, exampleValidPreparationInput)
			require.NoError(t, err)

			checkValidPreparationEquality(t, exampleValidPreparation, createdValidPreparation)

			createdValidPreparation, err = testClients.userClient.GetValidPreparation(ctx, createdValidPreparation.ID)
			requireNotNilAndNoProblems(t, createdValidPreparation, err)
			checkValidPreparationEquality(t, exampleValidPreparation, createdValidPreparation)

			exampleValidIngredient := fakes.BuildFakeValidIngredient()
			exampleValidIngredientInput := converters.ConvertValidIngredientToValidIngredientCreationRequestInput(exampleValidIngredient)
			createdValidIngredient, err := testClients.adminClient.CreateValidIngredient(ctx, exampleValidIngredientInput)
			require.NoError(t, err)

			checkValidIngredientEquality(t, exampleValidIngredient, createdValidIngredient)

			createdValidIngredient, err = testClients.userClient.GetValidIngredient(ctx, createdValidIngredient.ID)
			requireNotNilAndNoProblems(t, createdValidIngredient, err)
			checkValidIngredientEquality(t, exampleValidIngredient, createdValidIngredient)

			exampleValidIngredientPreparation := fakes.BuildFakeValidIngredientPreparation()
			exampleValidIngredientPreparation.Ingredient = *createdValidIngredient
			exampleValidIngredientPreparation.Preparation = *createdValidPreparation
			exampleValidIngredientPreparationInput := converters.ConvertValidIngredientPreparationToValidIngredientPreparationCreationRequestInput(exampleValidIngredientPreparation)
			createdValidIngredientPreparation, err := testClients.adminClient.CreateValidIngredientPreparation(ctx, exampleValidIngredientPreparationInput)
			requireNotNilAndNoProblems(t, createdValidIngredientPreparation, err)

			createdValidIngredientPreparation, err = testClients.adminClient.GetValidIngredientPreparation(ctx, createdValidIngredientPreparation.ID)
			requireNotNilAndNoProblems(t, createdValidIngredientPreparation, err)

			checkValidIngredientPreparationEquality(t, exampleValidIngredientPreparation, createdValidIngredientPreparation)

			validIngredientMeasurementUnitsForValidIngredient, err := testClients.userClient.GetValidIngredientPreparationsByIngredient(ctx, createdValidIngredient.ID, nil)
			requireNotNilAndNoProblems(t, validIngredientMeasurementUnitsForValidIngredient, err)

			require.Len(t, validIngredientMeasurementUnitsForValidIngredient.Data, 1)
			assert.Equal(t, validIngredientMeasurementUnitsForValidIngredient.Data[0].ID, createdValidIngredientPreparation.ID)

			validIngredientMeasurementUnitsForValidMeasurementUnit, err := testClients.userClient.GetValidIngredientPreparationsByPreparation(ctx, createdValidPreparation.ID, nil)
			requireNotNilAndNoProblems(t, validIngredientMeasurementUnitsForValidMeasurementUnit, err)

			require.Len(t, validIngredientMeasurementUnitsForValidMeasurementUnit.Data, 1)
			assert.Equal(t, validIngredientMeasurementUnitsForValidMeasurementUnit.Data[0].ID, createdValidIngredientPreparation.ID)

			assert.NoError(t, testClients.adminClient.ArchiveValidIngredientPreparation(ctx, createdValidIngredientPreparation.ID))

			assert.NoError(t, testClients.adminClient.ArchiveValidIngredient(ctx, createdValidIngredient.ID))

			assert.NoError(t, testClients.adminClient.ArchiveValidPreparation(ctx, createdValidPreparation.ID))
		}
	})

	s.runTest("should be searchable via preparation ID and ingredient name", func(testClients *testClientWrapper) func() {
		return func() {
			t := s.T()

			ctx, span := tracing.StartCustomSpan(s.ctx, t.Name())
			defer span.End()

			exampleValidPreparation := fakes.BuildFakeValidPreparation()
			exampleValidPreparationInput := converters.ConvertValidPreparationToValidPreparationCreationRequestInput(exampleValidPreparation)
			createdValidPreparation, err := testClients.adminClient.CreateValidPreparation(ctx, exampleValidPreparationInput)
			require.NoError(t, err)

			checkValidPreparationEquality(t, exampleValidPreparation, createdValidPreparation)

			createdValidPreparation, err = testClients.userClient.GetValidPreparation(ctx, createdValidPreparation.ID)
			requireNotNilAndNoProblems(t, createdValidPreparation, err)
			checkValidPreparationEquality(t, exampleValidPreparation, createdValidPreparation)

			createdValidIngredients := []*types.ValidIngredient{}
			createdValidIngredientPreparations := []*types.ValidIngredientPreparation{}

			exampleValidIngredient := fakes.BuildFakeValidIngredient()
			originalName := exampleValidIngredient.Name
			for i := 0; i < 5; i++ {
				exampleValidIngredient.Name = fmt.Sprintf("%s #%d", originalName, i)
				exampleValidIngredientInput := converters.ConvertValidIngredientToValidIngredientCreationRequestInput(exampleValidIngredient)
				createdValidIngredient, err := testClients.adminClient.CreateValidIngredient(ctx, exampleValidIngredientInput)
				require.NoError(t, err)

				checkValidIngredientEquality(t, exampleValidIngredient, createdValidIngredient)

				createdValidIngredient, err = testClients.userClient.GetValidIngredient(ctx, createdValidIngredient.ID)
				requireNotNilAndNoProblems(t, createdValidIngredient, err)
				checkValidIngredientEquality(t, exampleValidIngredient, createdValidIngredient)

				exampleValidIngredientPreparation := fakes.BuildFakeValidIngredientPreparation()
				exampleValidIngredientPreparation.Ingredient = *createdValidIngredient
				exampleValidIngredientPreparation.Preparation = *createdValidPreparation
				exampleValidIngredientPreparationInput := converters.ConvertValidIngredientPreparationToValidIngredientPreparationCreationRequestInput(exampleValidIngredientPreparation)
				createdValidIngredientPreparation, err := testClients.adminClient.CreateValidIngredientPreparation(ctx, exampleValidIngredientPreparationInput)
				require.NoError(t, err)

				checkValidIngredientPreparationEquality(t, exampleValidIngredientPreparation, createdValidIngredientPreparation)

				createdValidIngredients = append(createdValidIngredients, createdValidIngredient)
				createdValidIngredientPreparations = append(createdValidIngredientPreparations, createdValidIngredientPreparation)
			}

			//searchQuery := createdValidIngredients[0].Name[0:3]
			validIngredientMeasurementUnitsForValidIngredient, err := testClients.userClient.GetValidIngredientPreparationsByPreparation(ctx, createdValidPreparation.ID, nil)
			requireNotNilAndNoProblems(t, validIngredientMeasurementUnitsForValidIngredient, err)

			assert.Equal(t, len(validIngredientMeasurementUnitsForValidIngredient.Data), len(createdValidIngredients))

			for _, createdValidIngredientsPreparation := range createdValidIngredientPreparations {
				assert.NoError(t, testClients.adminClient.ArchiveValidIngredientPreparation(ctx, createdValidIngredientsPreparation.ID))
			}

			for _, createdValidIngredient := range createdValidIngredients {
				assert.NoError(t, testClients.adminClient.ArchiveValidIngredient(ctx, createdValidIngredient.ID))
			}

			assert.NoError(t, testClients.adminClient.ArchiveValidPreparation(ctx, createdValidPreparation.ID))
		}
	})
}
