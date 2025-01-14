package integration

import (
	"testing"

	"github.com/dinnerdonebetter/backend/internal/observability/tracing"
	"github.com/dinnerdonebetter/backend/internal/pkg/pointer"
	"github.com/dinnerdonebetter/backend/pkg/types"
	"github.com/dinnerdonebetter/backend/pkg/types/converters"
	"github.com/dinnerdonebetter/backend/pkg/types/fakes"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func checkRecipeStepVesselEquality(t *testing.T, expected, actual *types.RecipeStepVessel, checkInstrument bool) {
	t.Helper()

	assert.NotZero(t, actual.ID)
	if checkInstrument {
		checkValidVesselEquality(t, expected.Vessel, actual.Vessel)
	} else {
		assert.Equal(t, expected.Vessel.ID, actual.Vessel.ID, "expected Vessel.ID for recipe step vessel %s to be %v, but it was %v", expected.ID, expected.Vessel.ID, actual.Vessel.ID)
	}

	assert.Equal(t, expected.RecipeStepProductID, actual.RecipeStepProductID, "expected RecipeStepProductID for recipe step vessel %s to be %v, but it was %v", expected.ID, expected.RecipeStepProductID, actual.RecipeStepProductID)
	assert.Equal(t, expected.Name, actual.Name, "expected Name for recipe step vessel %s to be %v, but it was %v", expected.ID, expected.Name, actual.Name)
	assert.Equal(t, expected.Notes, actual.Notes, "expected Notes for recipe step vessel %s to be %v, but it was %v", expected.ID, expected.Notes, actual.Notes)
	assert.Equal(t, expected.BelongsToRecipeStep, actual.BelongsToRecipeStep, "expected BelongsToRecipeStep for recipe step vessel %s to be %v, but it was %v", expected.ID, expected.BelongsToRecipeStep, actual.BelongsToRecipeStep)
	assert.Equal(t, expected.VesselPreposition, actual.VesselPreposition, "expected VesselPreposition for recipe step vessel %s to be %v, but it was %v", expected.ID, expected.VesselPreposition, actual.VesselPreposition)
	assert.Equal(t, expected.Quantity, actual.Quantity, "expected Quantity for recipe step vessel %s to be %v, but it was %v", expected.ID, expected.Quantity, actual.Quantity)
	assert.Equal(t, expected.UnavailableAfterStep, actual.UnavailableAfterStep, "expected UnavailableAfterStep for recipe step vessel %s to be %v, but it was %v", expected.ID, expected.UnavailableAfterStep, actual.UnavailableAfterStep)

	assert.NotZero(t, actual.CreatedAt)
}

func (s *TestSuite) TestRecipeStepVessels_CompleteLifecycle() {
	s.runTest("should CRUD", func(testClients *testClientWrapper) func() {
		return func() {
			t := s.T()

			ctx, span := tracing.StartCustomSpan(s.ctx, t.Name())
			defer span.End()

			_, _, createdRecipe := createRecipeForTest(ctx, t, testClients.adminClient, testClients.userClient, nil)

			var createdRecipeStepID string
			for _, step := range createdRecipe.Steps {
				createdRecipeStepID = step.ID
				break
			}

			createdValidVessel := createValidVesselForTest(t, ctx, nil, testClients.adminClient)

			exampleRecipeStepVessel := fakes.BuildFakeRecipeStepVessel()
			exampleRecipeStepVessel.BelongsToRecipeStep = createdRecipeStepID
			exampleRecipeStepVessel.Vessel = &types.ValidVessel{ID: createdValidVessel.ID}
			exampleRecipeStepVesselInput := converters.ConvertRecipeStepVesselToRecipeStepVesselCreationRequestInput(exampleRecipeStepVessel)
			createdRecipeStepVessel, err := testClients.adminClient.CreateRecipeStepVessel(ctx, createdRecipe.ID, createdRecipeStepID, exampleRecipeStepVesselInput)
			require.NoError(t, err)

			checkRecipeStepVesselEquality(t, exampleRecipeStepVessel, createdRecipeStepVessel, false)

			createdRecipeStepVessel, err = testClients.userClient.GetRecipeStepVessel(ctx, createdRecipe.ID, createdRecipeStepID, createdRecipeStepVessel.ID)
			requireNotNilAndNoProblems(t, createdRecipeStepVessel, err)
			require.Equal(t, createdRecipeStepID, createdRecipeStepVessel.BelongsToRecipeStep)
			exampleRecipeStepVessel.Vessel = createdValidVessel
			exampleRecipeStepVessel.Vessel.CreatedAt = createdRecipeStepVessel.Vessel.CreatedAt

			checkRecipeStepVesselEquality(t, exampleRecipeStepVessel, createdRecipeStepVessel, false)

			newExampleValidInstrument := fakes.BuildFakeValidInstrument()
			newExampleValidInstrumentInput := converters.ConvertValidInstrumentToValidInstrumentCreationRequestInput(newExampleValidInstrument)
			newValidInstrument, err := testClients.adminClient.CreateValidInstrument(ctx, newExampleValidInstrumentInput)
			require.NoError(t, err)
			checkValidInstrumentEquality(t, newExampleValidInstrument, newValidInstrument)

			newRecipeStepVessel := fakes.BuildFakeRecipeStepVessel()
			newRecipeStepVessel.BelongsToRecipeStep = createdRecipeStepID
			newRecipeStepVessel.Vessel = createdValidVessel
			updateInput := converters.ConvertRecipeStepVesselToRecipeStepVesselUpdateRequestInput(newRecipeStepVessel)
			createdRecipeStepVessel.Update(updateInput)
			assert.NoError(t, testClients.adminClient.UpdateRecipeStepVessel(ctx, createdRecipe.ID, createdRecipeStepID, createdRecipeStepVessel.ID, updateInput))

			actual, err := testClients.userClient.GetRecipeStepVessel(ctx, createdRecipe.ID, createdRecipeStepID, createdRecipeStepVessel.ID)
			requireNotNilAndNoProblems(t, actual, err)

			// assert recipe step vessel equality
			checkRecipeStepVesselEquality(t, newRecipeStepVessel, actual, false)
			assert.NotNil(t, actual.LastUpdatedAt)

			assert.NoError(t, testClients.userClient.ArchiveRecipeStepVessel(ctx, createdRecipe.ID, createdRecipeStepID, createdRecipeStepVessel.ID))

			assert.NoError(t, testClients.userClient.ArchiveRecipeStep(ctx, createdRecipe.ID, createdRecipeStepID))

			assert.NoError(t, testClients.adminClient.ArchiveRecipe(ctx, createdRecipe.ID))
		}
	})
}

func (s *TestSuite) TestRecipeStepVessels_AsRecipeStepProducts() {
	s.runTest("should be able to use a recipe step vessel that was the product of a prior recipe step", func(testClients *testClientWrapper) func() {
		return func() {
			t := s.T()

			ctx, span := tracing.StartCustomSpan(s.ctx, t.Name())
			defer span.End()

			lineBase := fakes.BuildFakeValidPreparation()
			lineInput := converters.ConvertValidPreparationToValidPreparationCreationRequestInput(lineBase)
			line, err := testClients.adminClient.CreateValidPreparation(ctx, lineInput)
			require.NoError(t, err)

			roastBase := fakes.BuildFakeValidPreparation()
			roastInput := converters.ConvertValidPreparationToValidPreparationCreationRequestInput(roastBase)
			roast, err := testClients.adminClient.CreateValidPreparation(ctx, roastInput)
			require.NoError(t, err)

			bakingSheet := createValidVesselForTest(t, ctx, nil, testClients.adminClient)

			sheetsBase := fakes.BuildFakeValidMeasurementUnit()
			sheetsBaseInput := converters.ConvertValidMeasurementUnitToValidMeasurementUnitCreationRequestInput(sheetsBase)
			sheets, err := testClients.adminClient.CreateValidMeasurementUnit(ctx, sheetsBaseInput)
			require.NoError(t, err)
			checkValidMeasurementUnitEquality(t, sheetsBase, sheets)

			headsBase := fakes.BuildFakeValidMeasurementUnit()
			headsBaseInput := converters.ConvertValidMeasurementUnitToValidMeasurementUnitCreationRequestInput(headsBase)
			head, err := testClients.adminClient.CreateValidMeasurementUnit(ctx, headsBaseInput)
			require.NoError(t, err)
			checkValidMeasurementUnitEquality(t, headsBase, head)

			exampleUnits := fakes.BuildFakeValidMeasurementUnit()
			exampleUnitsInput := converters.ConvertValidMeasurementUnitToValidMeasurementUnitCreationRequestInput(exampleUnits)
			unit, err := testClients.adminClient.CreateValidMeasurementUnit(ctx, exampleUnitsInput)
			require.NoError(t, err)
			checkValidMeasurementUnitEquality(t, exampleUnits, unit)

			aluminumFoilBase := fakes.BuildFakeValidIngredient()
			aluminumFoilInput := converters.ConvertValidIngredientToValidIngredientCreationRequestInput(aluminumFoilBase)
			aluminumFoil, createdValidIngredientErr := testClients.adminClient.CreateValidIngredient(ctx, aluminumFoilInput)
			require.NoError(t, createdValidIngredientErr)

			garlic := fakes.BuildFakeValidIngredient()
			garlicInput := converters.ConvertValidIngredientToValidIngredientCreationRequestInput(garlic)
			garlic, garlicErr := testClients.adminClient.CreateValidIngredient(ctx, garlicInput)
			require.NoError(t, garlicErr)

			linedBakingSheetName := "lined baking sheet"

			expected := &types.Recipe{
				Name:                t.Name(),
				Slug:                "whatever-who-cares-yadda-yadda-vessels",
				YieldsComponentType: types.MealComponentTypesMain,
				PortionName:         t.Name(),
				PluralPortionName:   t.Name(),
				EstimatedPortions: types.Float32RangeWithOptionalMax{
					Max: nil,
					Min: 1,
				},
				Steps: []*types.RecipeStep{
					{
						Products: []*types.RecipeStepProduct{
							{
								Name:            linedBakingSheetName,
								Type:            types.RecipeStepProductVesselType,
								MeasurementUnit: unit,
								QuantityNotes:   "",
								Quantity: types.OptionalFloat32Range{
									Max: nil,
									Min: pointer.To(float32(1)),
								},
							},
						},
						Notes:       "first step",
						Preparation: *line,
						Ingredients: []*types.RecipeStepIngredient{
							{
								RecipeStepProductID: nil,
								Ingredient:          aluminumFoil,
								Name:                "aluminum foil",
								MeasurementUnit:     *sheets,
								Quantity: types.Float32RangeWithOptionalMax{
									Max: nil,
									Min: 3,
								},
							},
						},
						Vessels: []*types.RecipeStepVessel{
							{
								Vessel: bakingSheet,
								Name:   "baking sheet",
								Quantity: types.Uint16RangeWithOptionalMax{
									Max: nil,
									Min: 1,
								},
							},
						},
						Index: 0,
					},
					{
						Preparation: *roast,
						Vessels: []*types.RecipeStepVessel{
							{
								Name:   linedBakingSheetName,
								Vessel: nil,
							},
						},
						Products: []*types.RecipeStepProduct{
							{
								Name:            "roasted garlic",
								Type:            types.RecipeStepProductIngredientType,
								MeasurementUnit: head,
								QuantityNotes:   "",
								Quantity: types.OptionalFloat32Range{
									Max: nil,
									Min: pointer.To(float32(1)),
								},
							},
						},
						Notes: "second step",
						Ingredients: []*types.RecipeStepIngredient{
							{
								Ingredient:      garlic,
								Name:            "garlic",
								MeasurementUnit: *head,
								Quantity: types.Float32RangeWithOptionalMax{
									Max: nil,
									Min: 1,
								},
							},
						},
						Index: 1,
					},
				},
			}

			exampleRecipeInput := converters.ConvertRecipeToRecipeCreationRequestInput(expected)
			exampleRecipeInput.Steps[1].Vessels[0].ProductOfRecipeStepIndex = pointer.To(uint64(0))
			exampleRecipeInput.Steps[1].Vessels[0].ProductOfRecipeStepProductIndex = pointer.To(uint64(0))

			created, err := testClients.adminClient.CreateRecipe(ctx, exampleRecipeInput)
			require.NoError(t, err)
			checkRecipeEquality(t, expected, created)

			created, err = testClients.userClient.GetRecipe(ctx, created.ID)
			requireNotNilAndNoProblems(t, created, err)
			checkRecipeEquality(t, expected, created)

			recipeStepProductIndex := -1
			for i, vessel := range created.Steps[1].Vessels {
				if vessel.RecipeStepProductID != nil {
					recipeStepProductIndex = i
				}
			}

			require.NotEqual(t, -1, recipeStepProductIndex)
			require.NotNil(t, created.Steps[1].Vessels[recipeStepProductIndex].RecipeStepProductID)
			assert.Equal(t, created.Steps[0].Products[0].ID, *created.Steps[1].Vessels[recipeStepProductIndex].RecipeStepProductID)
		}
	})
}

func (s *TestSuite) TestRecipeStepVessels_Listing() {
	s.runTest("should be readable in paginated form", func(testClients *testClientWrapper) func() {
		return func() {
			t := s.T()

			ctx, span := tracing.StartCustomSpan(s.ctx, t.Name())
			defer span.End()

			_, _, createdRecipe := createRecipeForTest(ctx, t, testClients.adminClient, testClients.userClient, nil)

			var createdRecipeStepID string
			for _, step := range createdRecipe.Steps {
				createdRecipeStepID = step.ID
				break
			}

			createdValidVessel := createValidVesselForTest(t, ctx, nil, testClients.adminClient)

			var expected []*types.RecipeStepVessel
			for i := 0; i < 5; i++ {
				exampleRecipeStepVessel := fakes.BuildFakeRecipeStepVessel()
				exampleRecipeStepVessel.BelongsToRecipeStep = createdRecipeStepID
				exampleRecipeStepVessel.Vessel = &types.ValidVessel{ID: createdValidVessel.ID}
				exampleRecipeStepVesselInput := converters.ConvertRecipeStepVesselToRecipeStepVesselCreationRequestInput(exampleRecipeStepVessel)
				createdRecipeStepVessel, createdRecipeStepVesselErr := testClients.adminClient.CreateRecipeStepVessel(ctx, createdRecipe.ID, createdRecipeStepID, exampleRecipeStepVesselInput)
				require.NoError(t, createdRecipeStepVesselErr)
				checkRecipeStepVesselEquality(t, exampleRecipeStepVessel, createdRecipeStepVessel, false)

				createdRecipeStepVessel, createdRecipeStepVesselErr = testClients.userClient.GetRecipeStepVessel(ctx, createdRecipe.ID, createdRecipeStepID, createdRecipeStepVessel.ID)
				requireNotNilAndNoProblems(t, createdRecipeStepVessel, createdRecipeStepVesselErr)
				require.Equal(t, createdRecipeStepID, createdRecipeStepVessel.BelongsToRecipeStep)

				expected = append(expected, createdRecipeStepVessel)
			}

			// assert recipe step vessel list equality
			actual, err := testClients.userClient.GetRecipeStepVessels(ctx, createdRecipe.ID, createdRecipeStepID, nil)
			requireNotNilAndNoProblems(t, actual, err)
			assert.True(
				t,
				len(expected) <= len(actual.Data),
				"expected %d to be <= %d",
				len(expected),
				len(actual.Data),
			)

			for _, createdRecipeStepVessel := range expected {
				assert.NoError(t, testClients.userClient.ArchiveRecipeStepVessel(ctx, createdRecipe.ID, createdRecipeStepID, createdRecipeStepVessel.ID))
			}

			assert.NoError(t, testClients.userClient.ArchiveRecipeStep(ctx, createdRecipe.ID, createdRecipeStepID))

			assert.NoError(t, testClients.adminClient.ArchiveRecipe(ctx, createdRecipe.ID))
		}
	})
}
