package types

import (
	"context"
	"testing"

	fake "github.com/brianvoe/gofakeit/v5"
	"github.com/stretchr/testify/assert"
)

func TestValidIngredientMeasurementUnitCreationRequestInput_Validate(T *testing.T) {
	T.Parallel()

	T.Run("standard", func(t *testing.T) {
		t.Parallel()

		x := &ValidIngredientMeasurementUnitCreationRequestInput{
			Notes:                  fake.LoremIpsumSentence(exampleQuantity),
			ValidMeasurementUnitID: fake.LoremIpsumSentence(exampleQuantity),
			ValidIngredientID:      fake.LoremIpsumSentence(exampleQuantity),
		}

		actual := x.ValidateWithContext(context.Background())
		assert.Nil(t, actual)
	})

	T.Run("with invalid structure", func(t *testing.T) {
		t.Parallel()

		x := &ValidIngredientMeasurementUnitCreationRequestInput{}

		actual := x.ValidateWithContext(context.Background())
		assert.Error(t, actual)
	})
}

func TestValidIngredientMeasurementUnitUpdateRequestInput_Validate(T *testing.T) {
	T.Parallel()

	T.Run("standard", func(t *testing.T) {
		t.Parallel()

		x := &ValidIngredientMeasurementUnitUpdateRequestInput{
			Notes:                  stringPointer(fake.LoremIpsumSentence(exampleQuantity)),
			ValidMeasurementUnitID: stringPointer(fake.LoremIpsumSentence(exampleQuantity)),
			ValidIngredientID:      stringPointer(fake.LoremIpsumSentence(exampleQuantity)),
		}

		actual := x.ValidateWithContext(context.Background())
		assert.Nil(t, actual)
	})

	T.Run("with empty strings", func(t *testing.T) {
		t.Parallel()

		x := &ValidIngredientMeasurementUnitUpdateRequestInput{}

		actual := x.ValidateWithContext(context.Background())
		assert.Error(t, actual)
	})
}
