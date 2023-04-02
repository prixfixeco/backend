package grocerylistpreparation

import (
	"context"
	"testing"

	"github.com/prixfixeco/backend/internal/observability/logging"
	"github.com/prixfixeco/backend/internal/observability/tracing"
	"github.com/prixfixeco/backend/internal/pointers"
	"github.com/prixfixeco/backend/pkg/types"
	"github.com/prixfixeco/backend/pkg/types/fakes"

	"github.com/stretchr/testify/assert"
)

func Test_groceryListCreator_GenerateGroceryListInputs(T *testing.T) {
	T.Parallel()

	T.Run("standard", func(t *testing.T) {
		t.Parallel()

		listGenerator := &groceryListCreator{
			logger: logging.NewNoopLogger(),
			tracer: tracing.NewTracer(tracing.NewNoopTracerProvider().Tracer(t.Name())),
		}

		onion := fakes.BuildFakeValidIngredient()
		carrot := fakes.BuildFakeValidIngredient()
		celery := fakes.BuildFakeValidIngredient()
		salt := fakes.BuildFakeValidIngredient()

		grams := fakes.BuildFakeValidMeasurementUnit()

		expectedMealPlan := &types.MealPlan{
			ID: fakes.BuildFakeID(),
			Events: []*types.MealPlanEvent{
				{
					Options: []*types.MealPlanOption{
						{
							Chosen: true,
							Meal: types.Meal{
								Components: []*types.MealComponent{
									{
										Recipe: types.Recipe{
											Steps: []*types.RecipeStep{
												{
													Ingredients: []*types.RecipeStepIngredient{
														{
															Ingredient:      onion,
															MinimumQuantity: 100,
															MaximumQuantity: pointers.Float32(100),
															MeasurementUnit: *grams,
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
				{
					Options: []*types.MealPlanOption{
						{
							Chosen: true,
							Meal: types.Meal{
								Components: []*types.MealComponent{
									{
										Recipe: types.Recipe{
											Steps: []*types.RecipeStep{
												{
													Ingredients: []*types.RecipeStepIngredient{
														{
															Ingredient:      carrot,
															MinimumQuantity: 100,
															MaximumQuantity: pointers.Float32(100),
															MeasurementUnit: *grams,
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
				{
					Options: []*types.MealPlanOption{
						{
							Chosen: true,
							Meal: types.Meal{
								Components: []*types.MealComponent{
									{
										Recipe: types.Recipe{
											Steps: []*types.RecipeStep{
												{
													Ingredients: []*types.RecipeStepIngredient{
														{
															Ingredient:      celery,
															MinimumQuantity: 100,
															MaximumQuantity: pointers.Float32(100),
															MeasurementUnit: *grams,
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
				{
					Options: []*types.MealPlanOption{
						{
							Chosen: true,
							Meal: types.Meal{
								Components: []*types.MealComponent{
									{
										Recipe: types.Recipe{
											Steps: []*types.RecipeStep{
												{
													Ingredients: []*types.RecipeStepIngredient{
														{
															Ingredient:      salt,
															MinimumQuantity: 100,
															MaximumQuantity: pointers.Float32(100),
															MeasurementUnit: *grams,
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
				{
					Options: []*types.MealPlanOption{
						{
							Chosen: true,
							Meal: types.Meal{
								Components: []*types.MealComponent{
									{
										Recipe: types.Recipe{
											Steps: []*types.RecipeStep{
												{
													Ingredients: []*types.RecipeStepIngredient{
														{
															Ingredient:      onion,
															MinimumQuantity: 100,
															MaximumQuantity: pointers.Float32(100),
															MeasurementUnit: *grams,
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		}

		ctx := context.Background()

		expectedGroceryListItemInputs := map[string]*types.MealPlanGroceryListItemDatabaseCreationInput{
			onion.ID: {
				Status:                 types.MealPlanGroceryListItemStatusUnknown,
				ValidMeasurementUnitID: grams.ID,
				ValidIngredientID:      onion.ID,
				BelongsToMealPlan:      expectedMealPlan.ID,
				MinimumQuantityNeeded:  200,
				MaximumQuantityNeeded:  pointers.Float32(200),
			},
			carrot.ID: {
				Status:                 types.MealPlanGroceryListItemStatusUnknown,
				ValidMeasurementUnitID: grams.ID,
				ValidIngredientID:      carrot.ID,
				BelongsToMealPlan:      expectedMealPlan.ID,
				MinimumQuantityNeeded:  100,
				MaximumQuantityNeeded:  pointers.Float32(100),
			},
			celery.ID: {
				Status:                 types.MealPlanGroceryListItemStatusUnknown,
				ValidMeasurementUnitID: grams.ID,
				ValidIngredientID:      celery.ID,
				BelongsToMealPlan:      expectedMealPlan.ID,
				MinimumQuantityNeeded:  100,
				MaximumQuantityNeeded:  pointers.Float32(100),
			},
			salt.ID: {
				Status:                 types.MealPlanGroceryListItemStatusUnknown,
				ValidMeasurementUnitID: grams.ID,
				ValidIngredientID:      salt.ID,
				BelongsToMealPlan:      expectedMealPlan.ID,
				MinimumQuantityNeeded:  100,
				MaximumQuantityNeeded:  pointers.Float32(100),
			},
		}

		actual, err := listGenerator.GenerateGroceryListInputs(ctx, expectedMealPlan)
		assert.NoError(t, err)

		actualMap := map[string]*types.MealPlanGroceryListItemDatabaseCreationInput{}
		for i := range actual {
			actualMap[actual[i].ValidIngredientID] = actual[i]
			expectedGroceryListItemInputs[actual[i].ValidIngredientID].ID = actual[i].ID
		}

		assert.Equal(t, expectedGroceryListItemInputs, actualMap)
	})
}
