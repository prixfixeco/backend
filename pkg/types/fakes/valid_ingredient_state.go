package fakes

import (
	"github.com/prixfixeco/backend/pkg/types"
	"github.com/prixfixeco/backend/pkg/types/converters"
)

// BuildFakeValidIngredientState builds a faked valid preparation.
func BuildFakeValidIngredientState() *types.ValidIngredientState {
	return &types.ValidIngredientState{
		ID:          BuildFakeID(),
		Name:        buildUniqueString(),
		Description: buildUniqueString(),
		IconPath:    buildUniqueString(),
		Slug:        buildUniqueString(),
		PastTense:   buildUniqueString(),
		CreatedAt:   BuildFakeTime(),
	}
}

// BuildFakeValidIngredientStateList builds a faked ValidIngredientStateList.
func BuildFakeValidIngredientStateList() *types.ValidIngredientStateList {
	var examples []*types.ValidIngredientState
	for i := 0; i < exampleQuantity; i++ {
		examples = append(examples, BuildFakeValidIngredientState())
	}

	return &types.ValidIngredientStateList{
		Pagination: types.Pagination{
			Page:          1,
			Limit:         20,
			FilteredCount: exampleQuantity / 2,
			TotalCount:    exampleQuantity,
		},
		ValidIngredientStates: examples,
	}
}

// BuildFakeValidIngredientStateUpdateRequestInput builds a faked ValidIngredientStateUpdateRequestInput from a valid preparation.
func BuildFakeValidIngredientStateUpdateRequestInput() *types.ValidIngredientStateUpdateRequestInput {
	validIngredientState := BuildFakeValidIngredientState()
	return converters.ConvertValidIngredientStateToValidIngredientStateUpdateRequestInput(validIngredientState)
}

// BuildFakeValidIngredientStateCreationRequestInput builds a faked ValidIngredientStateCreationRequestInput.
func BuildFakeValidIngredientStateCreationRequestInput() *types.ValidIngredientStateCreationRequestInput {
	validIngredientState := BuildFakeValidIngredientState()
	return converters.ConvertValidIngredientStateToValidIngredientStateCreationRequestInput(validIngredientState)
}