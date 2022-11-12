package fakes

import (
	"github.com/prixfixeco/backend/pkg/types"
	"github.com/prixfixeco/backend/pkg/types/converters"
)

// BuildFakeMeal builds a faked meal.
func BuildFakeMeal() *types.Meal {
	recipes := []*types.MealComponent{}
	for i := 0; i < exampleQuantity; i++ {
		recipes = append(recipes, BuildFakeMealComponent())
	}

	return &types.Meal{
		ID:            BuildFakeID(),
		Name:          buildUniqueString(),
		Description:   buildUniqueString(),
		CreatedAt:     BuildFakeTime(),
		CreatedByUser: BuildFakeID(),
		Components:    recipes,
	}
}

// BuildFakeMealComponent builds a faked meal component.
func BuildFakeMealComponent() *types.MealComponent {
	return &types.MealComponent{
		Recipe:        *BuildFakeRecipe(),
		ComponentType: types.MealComponentTypesAmuseBouche,
	}
}

// BuildFakeMealList builds a faked MealList.
func BuildFakeMealList() *types.MealList {
	var examples []*types.Meal
	for i := 0; i < exampleQuantity; i++ {
		examples = append(examples, BuildFakeMeal())
	}

	return &types.MealList{
		Pagination: types.Pagination{
			Page:          1,
			Limit:         20,
			FilteredCount: exampleQuantity / 2,
			TotalCount:    exampleQuantity,
		},
		Meals: examples,
	}
}

// BuildFakeMealCreationRequestInput builds a faked MealCreationRequestInput.
func BuildFakeMealCreationRequestInput() *types.MealCreationRequestInput {
	recipe := BuildFakeMeal()
	return converters.ConvertMealToMealCreationRequestInput(recipe)
}
