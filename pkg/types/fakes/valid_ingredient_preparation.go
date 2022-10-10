package fakes

import (
	fake "github.com/brianvoe/gofakeit/v5"

	"github.com/prixfixeco/api_server/pkg/types"
)

// BuildFakeValidIngredientPreparation builds a faked valid ingredient preparation.
func BuildFakeValidIngredientPreparation() *types.ValidIngredientPreparation {
	return &types.ValidIngredientPreparation{
		ID:          BuildFakeID(),
		Notes:       buildUniqueString(),
		Preparation: *BuildFakeValidPreparation(),
		Ingredient:  *BuildFakeValidIngredient(),
		CreatedAt:   fake.Date(),
	}
}

// BuildFakeValidIngredientPreparationList builds a faked ValidIngredientPreparationList.
func BuildFakeValidIngredientPreparationList() *types.ValidIngredientPreparationList {
	var examples []*types.ValidIngredientPreparation
	for i := 0; i < exampleQuantity; i++ {
		examples = append(examples, BuildFakeValidIngredientPreparation())
	}

	return &types.ValidIngredientPreparationList{
		Pagination: types.Pagination{
			Page:          1,
			Limit:         20,
			FilteredCount: exampleQuantity / 2,
			TotalCount:    exampleQuantity,
		},
		ValidIngredientPreparations: examples,
	}
}

// BuildFakeValidIngredientPreparationUpdateRequestInput builds a faked ValidIngredientPreparationUpdateRequestInput from a valid ingredient preparation.
func BuildFakeValidIngredientPreparationUpdateRequestInput() *types.ValidIngredientPreparationUpdateRequestInput {
	validIngredientPreparation := BuildFakeValidIngredientPreparation()
	return &types.ValidIngredientPreparationUpdateRequestInput{
		Notes:              &validIngredientPreparation.Notes,
		ValidPreparationID: &validIngredientPreparation.Preparation.ID,
		ValidIngredientID:  &validIngredientPreparation.Ingredient.ID,
	}
}

// BuildFakeValidIngredientPreparationUpdateRequestInputFromValidIngredientPreparation builds a faked ValidIngredientPreparationUpdateRequestInput from a valid ingredient preparation.
func BuildFakeValidIngredientPreparationUpdateRequestInputFromValidIngredientPreparation(validIngredientPreparation *types.ValidIngredientPreparation) *types.ValidIngredientPreparationUpdateRequestInput {
	return &types.ValidIngredientPreparationUpdateRequestInput{
		Notes:              &validIngredientPreparation.Notes,
		ValidPreparationID: &validIngredientPreparation.Preparation.ID,
		ValidIngredientID:  &validIngredientPreparation.Ingredient.ID,
	}
}

// BuildFakeValidIngredientPreparationCreationRequestInput builds a faked ValidIngredientPreparationCreationRequestInput.
func BuildFakeValidIngredientPreparationCreationRequestInput() *types.ValidIngredientPreparationCreationRequestInput {
	validIngredientPreparation := BuildFakeValidIngredientPreparation()
	return BuildFakeValidIngredientPreparationCreationRequestInputFromValidIngredientPreparation(validIngredientPreparation)
}

// BuildFakeValidIngredientPreparationCreationRequestInputFromValidIngredientPreparation builds a faked ValidIngredientPreparationCreationRequestInput from a valid ingredient preparation.
func BuildFakeValidIngredientPreparationCreationRequestInputFromValidIngredientPreparation(validIngredientPreparation *types.ValidIngredientPreparation) *types.ValidIngredientPreparationCreationRequestInput {
	return &types.ValidIngredientPreparationCreationRequestInput{
		ID:                 validIngredientPreparation.ID,
		Notes:              validIngredientPreparation.Notes,
		ValidPreparationID: validIngredientPreparation.Preparation.ID,
		ValidIngredientID:  validIngredientPreparation.Ingredient.ID,
	}
}

// BuildFakeValidIngredientPreparationDatabaseCreationInputFromValidIngredientPreparation builds a faked ValidIngredientPreparationDatabaseCreationInput from a valid ingredient preparation.
func BuildFakeValidIngredientPreparationDatabaseCreationInputFromValidIngredientPreparation(validIngredientPreparation *types.ValidIngredientPreparation) *types.ValidIngredientPreparationDatabaseCreationInput {
	return &types.ValidIngredientPreparationDatabaseCreationInput{
		ID:                 validIngredientPreparation.ID,
		Notes:              validIngredientPreparation.Notes,
		ValidPreparationID: validIngredientPreparation.Preparation.ID,
		ValidIngredientID:  validIngredientPreparation.Ingredient.ID,
	}
}
