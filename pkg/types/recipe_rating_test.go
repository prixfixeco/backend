package types

import (
	"context"
	"testing"

	"github.com/dinnerdonebetter/backend/internal/pkg/pointers"

	fake "github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func TestRecipeRating_Update(T *testing.T) {
	T.Parallel()

	T.Run("standard", func(t *testing.T) {
		t.Parallel()

		x := &RecipeRating{}
		input := &RecipeRatingUpdateRequestInput{}

		fake.Struct(&input)

		x.Update(input)
	})
}

func TestRecipeRatingCreationRequestInput_ValidateWithContext(T *testing.T) {
	T.Parallel()

	T.Run("standard", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		x := &RecipeRatingCreationRequestInput{
			RecipeID:   t.Name(),
			Difficulty: 1.0,
		}

		assert.NoError(t, x.ValidateWithContext(ctx))
	})

	T.Run("with invalid structure", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		x := &RecipeRatingCreationRequestInput{}

		assert.Error(t, x.ValidateWithContext(ctx))
	})
}

func TestRecipeRatingDatabaseCreationInput_ValidateWithContext(T *testing.T) {
	T.Parallel()

	T.Run("standard", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		x := &RecipeRatingDatabaseCreationInput{
			ID:         t.Name(),
			RecipeID:   t.Name(),
			ByUser:     t.Name(),
			Difficulty: 1.0,
		}

		assert.NoError(t, x.ValidateWithContext(ctx))
	})

	T.Run("with invalid structure", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		x := &RecipeRatingDatabaseCreationInput{}

		assert.Error(t, x.ValidateWithContext(ctx))
	})
}

func TestRecipeRatingUpdateRequestInput_ValidateWithContext(T *testing.T) {
	T.Parallel()

	T.Run("standard", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		x := &RecipeRatingUpdateRequestInput{
			ByUser:     pointers.Pointer(t.Name()),
			RecipeID:   pointers.Pointer(t.Name()),
			Difficulty: pointers.Pointer[float32](1.0),
		}

		assert.NoError(t, x.ValidateWithContext(ctx))
	})

	T.Run("with invalid structure", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		x := &RecipeRatingUpdateRequestInput{}

		assert.Error(t, x.ValidateWithContext(ctx))
	})
}
