package integration

import (
	"fmt"
	"github.com/prixfixeco/api_server/pkg/types/converters"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/prixfixeco/api_server/internal/observability/tracing"
	"github.com/prixfixeco/api_server/pkg/types"
	"github.com/prixfixeco/api_server/pkg/types/fakes"
)

func checkValidIngredientEquality(t *testing.T, expected, actual *types.ValidIngredient) {
	t.Helper()

	assert.NotZero(t, actual.ID)
	assert.Equal(t, expected.Name, actual.Name, "expected Name for valid ingredient %s to be %v, but it was %v", expected.ID, expected.Name, actual.Name)
	assert.Equal(t, expected.Description, actual.Description, "expected Description for valid ingredient %s to be %v, but it was %v", expected.ID, expected.Description, actual.Description)
	assert.Equal(t, expected.Warning, actual.Warning, "expected Warning for valid ingredient %s to be %v, but it was %v", expected.ID, expected.Warning, actual.Warning)
	assert.Equal(t, expected.ContainsEgg, actual.ContainsEgg, "expected ContainsEgg for valid ingredient %s to be %v, but it was %v", expected.ID, expected.ContainsEgg, actual.ContainsEgg)
	assert.Equal(t, expected.ContainsDairy, actual.ContainsDairy, "expected ContainsDairy for valid ingredient %s to be %v, but it was %v", expected.ID, expected.ContainsDairy, actual.ContainsDairy)
	assert.Equal(t, expected.ContainsPeanut, actual.ContainsPeanut, "expected ContainsPeanut for valid ingredient %s to be %v, but it was %v", expected.ID, expected.ContainsPeanut, actual.ContainsPeanut)
	assert.Equal(t, expected.ContainsTreeNut, actual.ContainsTreeNut, "expected ContainsTreeNut for valid ingredient %s to be %v, but it was %v", expected.ID, expected.ContainsTreeNut, actual.ContainsTreeNut)
	assert.Equal(t, expected.ContainsSoy, actual.ContainsSoy, "expected ContainsSoy for valid ingredient %s to be %v, but it was %v", expected.ID, expected.ContainsSoy, actual.ContainsSoy)
	assert.Equal(t, expected.ContainsWheat, actual.ContainsWheat, "expected ContainsWheat for valid ingredient %s to be %v, but it was %v", expected.ID, expected.ContainsWheat, actual.ContainsWheat)
	assert.Equal(t, expected.ContainsShellfish, actual.ContainsShellfish, "expected ContainsShellfish for valid ingredient %s to be %v, but it was %v", expected.ID, expected.ContainsShellfish, actual.ContainsShellfish)
	assert.Equal(t, expected.ContainsSesame, actual.ContainsSesame, "expected ContainsSesame for valid ingredient %s to be %v, but it was %v", expected.ID, expected.ContainsSesame, actual.ContainsSesame)
	assert.Equal(t, expected.ContainsFish, actual.ContainsFish, "expected ContainsFish for valid ingredient %s to be %v, but it was %v", expected.ID, expected.ContainsFish, actual.ContainsFish)
	assert.Equal(t, expected.ContainsGluten, actual.ContainsGluten, "expected ContainsGluten for valid ingredient %s to be %v, but it was %v", expected.ID, expected.ContainsGluten, actual.ContainsGluten)
	assert.Equal(t, expected.AnimalFlesh, actual.AnimalFlesh, "expected AnimalFlesh for valid ingredient %s to be %v, but it was %v", expected.ID, expected.AnimalFlesh, actual.AnimalFlesh)
	assert.Equal(t, expected.IsMeasuredVolumetrically, actual.IsMeasuredVolumetrically, "expected IsMeasuredVolumetrically for valid ingredient %s to be %v, but it was %v", expected.ID, expected.IsMeasuredVolumetrically, actual.IsMeasuredVolumetrically)
	assert.Equal(t, expected.IsLiquid, actual.IsLiquid, "expected IsLiquid for valid ingredient %s to be %v, but it was %v", expected.ID, expected.IsLiquid, actual.IsLiquid)
	assert.Equal(t, expected.IconPath, actual.IconPath, "expected IconPath for valid ingredient %s to be %v, but it was %v", expected.ID, expected.IconPath, actual.IconPath)
	assert.Equal(t, expected.PluralName, actual.PluralName, "expected PluralName for valid ingredient %s to be %v, but it was %v", expected.ID, expected.PluralName, actual.PluralName)
	assert.Equal(t, expected.AnimalDerived, actual.AnimalDerived, "expected AnimalDerived for valid ingredient %s to be %v, but it was %v", expected.ID, expected.AnimalDerived, actual.AnimalDerived)
	assert.Equal(t, expected.RestrictToPreparations, actual.RestrictToPreparations, "expected RestrictToPreparations for valid ingredient %s to be %v, but it was %v", expected.ID, expected.RestrictToPreparations, actual.RestrictToPreparations)
	assert.Equal(t, expected.MinimumIdealStorageTemperatureInCelsius, actual.MinimumIdealStorageTemperatureInCelsius, "expected MinimumIdealStorageTemperatureInCelsius for valid ingredient %s to be %v, but it was %v", expected.ID, expected.MinimumIdealStorageTemperatureInCelsius, actual.MinimumIdealStorageTemperatureInCelsius)
	assert.Equal(t, expected.MaximumIdealStorageTemperatureInCelsius, actual.MaximumIdealStorageTemperatureInCelsius, "expected MaximumIdealStorageTemperatureInCelsius for valid ingredient %s to be %v, but it was %v", expected.ID, expected.MaximumIdealStorageTemperatureInCelsius, actual.MaximumIdealStorageTemperatureInCelsius)
	assert.Equal(t, expected.StorageInstructions, actual.StorageInstructions, "expected StorageInstructions for valid ingredient %s to be %v, but it was %v", expected.ID, expected.StorageInstructions, actual.StorageInstructions)
	assert.NotZero(t, actual.CreatedAt)
}

func (s *TestSuite) TestValidIngredients_CompleteLifecycle() {
	s.runForEachClient("should be creatable and readable and updatable and deletable", func(testClients *testClientWrapper) func() {
		return func() {
			t := s.T()

			ctx, span := tracing.StartCustomSpan(s.ctx, t.Name())
			defer span.End()

			t.Log("creating valid ingredient")
			exampleValidIngredient := fakes.BuildFakeValidIngredient()
			exampleValidIngredientInput := converters.ConvertValidIngredientToValidIngredientCreationRequestInput(exampleValidIngredient)
			createdValidIngredient, err := testClients.admin.CreateValidIngredient(ctx, exampleValidIngredientInput)
			require.NoError(t, err)
			t.Logf("valid ingredient %q created", createdValidIngredient.ID)
			checkValidIngredientEquality(t, exampleValidIngredient, createdValidIngredient)

			createdValidIngredient, err = testClients.admin.GetValidIngredient(ctx, createdValidIngredient.ID)
			requireNotNilAndNoProblems(t, createdValidIngredient, err)
			checkValidIngredientEquality(t, exampleValidIngredient, createdValidIngredient)

			t.Log("changing valid ingredient")
			newValidIngredient := fakes.BuildFakeValidIngredient()
			createdValidIngredient.Update(converters.ConvertValidIngredientToValidIngredientUpdateRequestInput(newValidIngredient))
			assert.NoError(t, testClients.admin.UpdateValidIngredient(ctx, createdValidIngredient))

			t.Log("fetching changed valid ingredient")
			actual, err := testClients.admin.GetValidIngredient(ctx, createdValidIngredient.ID)
			requireNotNilAndNoProblems(t, actual, err)

			// assert valid ingredient equality
			checkValidIngredientEquality(t, newValidIngredient, actual)
			assert.NotNil(t, actual.LastUpdatedAt)

			t.Log("cleaning up valid ingredient")
			assert.NoError(t, testClients.admin.ArchiveValidIngredient(ctx, createdValidIngredient.ID))
		}
	})
}

func (s *TestSuite) TestValidIngredients_GetRandom() {
	s.runForEachClient("should be able to get a random valid ingredient", func(testClients *testClientWrapper) func() {
		return func() {
			t := s.T()

			ctx, span := tracing.StartCustomSpan(s.ctx, t.Name())
			defer span.End()

			t.Log("creating valid ingredient")
			exampleValidIngredient := fakes.BuildFakeValidIngredient()
			exampleValidIngredientInput := converters.ConvertValidIngredientToValidIngredientCreationRequestInput(exampleValidIngredient)
			createdValidIngredient, err := testClients.admin.CreateValidIngredient(ctx, exampleValidIngredientInput)
			require.NoError(t, err)
			t.Logf("valid ingredient %q created", createdValidIngredient.ID)
			checkValidIngredientEquality(t, exampleValidIngredient, createdValidIngredient)

			createdValidIngredient, err = testClients.admin.GetRandomValidIngredient(ctx)
			requireNotNilAndNoProblems(t, createdValidIngredient, err)

			t.Log("cleaning up valid ingredient")
			assert.NoError(t, testClients.admin.ArchiveValidIngredient(ctx, createdValidIngredient.ID))
		}
	})
}

func (s *TestSuite) TestValidIngredients_Listing() {
	s.runForEachClient("should be readable in paginated form", func(testClients *testClientWrapper) func() {
		return func() {
			t := s.T()

			ctx, span := tracing.StartCustomSpan(s.ctx, t.Name())
			defer span.End()

			t.Log("creating valid ingredients")
			var expected []*types.ValidIngredient
			for i := 0; i < 5; i++ {
				exampleValidIngredient := fakes.BuildFakeValidIngredient()
				exampleValidIngredientInput := converters.ConvertValidIngredientToValidIngredientCreationRequestInput(exampleValidIngredient)
				createdValidIngredient, createdValidIngredientErr := testClients.admin.CreateValidIngredient(ctx, exampleValidIngredientInput)
				require.NoError(t, createdValidIngredientErr)
				t.Logf("valid ingredient %q created", createdValidIngredient.ID)

				checkValidIngredientEquality(t, exampleValidIngredient, createdValidIngredient)

				expected = append(expected, createdValidIngredient)
			}

			// assert valid ingredient list equality
			actual, err := testClients.admin.GetValidIngredients(ctx, nil)
			requireNotNilAndNoProblems(t, actual, err)
			assert.True(
				t,
				len(expected) <= len(actual.ValidIngredients),
				"expected %d to be <= %d",
				len(expected),
				len(actual.ValidIngredients),
			)

			t.Log("cleaning up")
			for _, createdValidIngredient := range expected {
				assert.NoError(t, testClients.admin.ArchiveValidIngredient(ctx, createdValidIngredient.ID))
			}
		}
	})
}

func (s *TestSuite) TestValidIngredients_Searching() {
	s.runForEachClient("should be able to be search for valid ingredients", func(testClients *testClientWrapper) func() {
		return func() {
			t := s.T()

			ctx, span := tracing.StartCustomSpan(s.ctx, t.Name())
			defer span.End()

			t.Log("creating valid ingredients")
			var expected []*types.ValidIngredient
			exampleValidIngredient := fakes.BuildFakeValidIngredient()
			exampleValidIngredient.Name = fmt.Sprintf("example_%s", testClients.authType)
			searchQuery := exampleValidIngredient.Name
			for i := 0; i < 5; i++ {
				exampleValidIngredient.Name = fmt.Sprintf("%s %d", searchQuery, i)
				exampleValidIngredientInput := converters.ConvertValidIngredientToValidIngredientCreationRequestInput(exampleValidIngredient)
				createdValidIngredient, createdValidIngredientErr := testClients.admin.CreateValidIngredient(ctx, exampleValidIngredientInput)
				require.NoError(t, createdValidIngredientErr)
				checkValidIngredientEquality(t, exampleValidIngredient, createdValidIngredient)
				t.Logf("valid ingredient %q created", createdValidIngredient.ID)

				expected = append(expected, createdValidIngredient)
			}

			exampleLimit := uint8(20)

			// assert valid ingredient list equality
			actual, err := testClients.admin.SearchValidIngredients(ctx, searchQuery, exampleLimit)
			requireNotNilAndNoProblems(t, actual, err)
			assert.True(
				t,
				len(expected) <= len(actual),
				"expected %d to be <= %d",
				len(expected),
				len(actual),
			)

			t.Log("cleaning up")
			for _, createdValidIngredient := range expected {
				assert.NoError(t, testClients.admin.ArchiveValidIngredient(ctx, createdValidIngredient.ID))
			}
		}
	})
}
