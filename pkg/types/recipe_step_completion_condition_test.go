package types

import (
	"context"
	"testing"

	"github.com/dinnerdonebetter/backend/internal/pkg/pointers"

	fake "github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func TestRecipeStepCompletionCondition_Update(T *testing.T) {
	T.Parallel()

	T.Run("standard", func(t *testing.T) {
		t.Parallel()

		x := &RecipeStepCompletionCondition{}
		input := &RecipeStepCompletionConditionUpdateRequestInput{}

		fake.Struct(&input)
		input.Optional = pointers.Pointer(true)

		x.Update(input)
	})
}

func TestRecipeStepCompletionConditionCreationRequestInput_Validate(T *testing.T) {
	T.Parallel()

	T.Run("standard", func(t *testing.T) {
		t.Parallel()

		x := &RecipeStepCompletionConditionCreationRequestInput{
			IngredientStateID:   t.Name(),
			BelongsToRecipeStep: t.Name(),
			Optional:            fake.Bool(),
			Ingredients:         []uint64{123},
		}

		actual := x.ValidateWithContext(context.Background())
		assert.Nil(t, actual)
	})

	T.Run("with invalid structure", func(t *testing.T) {
		t.Parallel()

		x := &RecipeStepCompletionConditionCreationRequestInput{}

		actual := x.ValidateWithContext(context.Background())
		assert.Error(t, actual)
	})
}

func TestRecipeStepCompletionConditionUpdateRequestInput_Validate(T *testing.T) {
	T.Parallel()

	T.Run("standard", func(t *testing.T) {
		t.Parallel()

		x := &RecipeStepCompletionConditionUpdateRequestInput{
			IngredientStateID:   pointers.Pointer(t.Name()),
			BelongsToRecipeStep: pointers.Pointer(t.Name()),
			Optional:            pointers.Pointer(fake.Bool()),
		}

		actual := x.ValidateWithContext(context.Background())
		assert.Nil(t, actual)
	})

	T.Run("with empty strings", func(t *testing.T) {
		t.Parallel()

		x := &RecipeStepCompletionConditionUpdateRequestInput{}

		actual := x.ValidateWithContext(context.Background())
		assert.Error(t, actual)
	})
}
