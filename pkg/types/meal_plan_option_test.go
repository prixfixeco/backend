package types

import (
	"context"
	"testing"
	"time"

	fake "github.com/brianvoe/gofakeit/v5"
	"github.com/stretchr/testify/assert"
)

func TestMealPlanOptionCreationRequestInput_Validate(T *testing.T) {
	T.Parallel()

	T.Run("standard", func(t *testing.T) {
		t.Parallel()

		x := &MealPlanOptionCreationRequestInput{
			AssignedCook:           stringPointer(fake.LoremIpsumSentence(exampleQuantity)),
			AssignedDishwasher:     stringPointer(fake.LoremIpsumSentence(exampleQuantity)),
			BelongsToMealPlanEvent: fake.LoremIpsumSentence(exampleQuantity),
			Day: time.Weekday(fake.RandomInt([]int{
				int(time.Monday),
				int(time.Tuesday),
				int(time.Wednesday),
				int(time.Thursday),
				int(time.Friday),
				int(time.Saturday),
				int(time.Sunday),
			})),
			MealName: fake.RandomString([]string{
				BreakfastMealName,
				SecondBreakfastMealName,
				BrunchMealName,
				LunchMealName,
				SupperMealName,
				DinnerMealName,
			}),
			MealID:           fake.LoremIpsumSentence(exampleQuantity),
			Notes:            fake.LoremIpsumSentence(exampleQuantity),
			PrepStepsCreated: false,
		}

		actual := x.ValidateWithContext(context.Background())
		assert.Nil(t, actual)
	})

	T.Run("with invalid structure", func(t *testing.T) {
		t.Parallel()

		x := &MealPlanOptionCreationRequestInput{}

		actual := x.ValidateWithContext(context.Background())
		assert.Error(t, actual)
	})
}

func TestMealPlanOptionUpdateRequestInput_Validate(T *testing.T) {
	T.Parallel()

	T.Run("standard", func(t *testing.T) {
		t.Parallel()

		day := time.Weekday(fake.RandomInt([]int{
			int(time.Monday),
			int(time.Tuesday),
			int(time.Wednesday),
			int(time.Thursday),
			int(time.Friday),
			int(time.Saturday),
			int(time.Sunday),
		}))

		mealName := fake.RandomString([]string{
			BreakfastMealName,
			SecondBreakfastMealName,
			BrunchMealName,
			LunchMealName,
			SupperMealName,
			DinnerMealName,
		})

		x := &MealPlanOptionUpdateRequestInput{
			AssignedCook:           stringPointer(fake.LoremIpsumSentence(exampleQuantity)),
			AssignedDishwasher:     stringPointer(fake.LoremIpsumSentence(exampleQuantity)),
			BelongsToMealPlanEvent: stringPointer(fake.LoremIpsumSentence(exampleQuantity)),
			Day:                    &day,
			MealName:               &mealName,
			MealID:                 stringPointer(fake.LoremIpsumSentence(exampleQuantity)),
			Notes:                  stringPointer(fake.LoremIpsumSentence(exampleQuantity)),
			PrepStepsCreated:       boolPointer(false),
		}

		actual := x.ValidateWithContext(context.Background())
		assert.Nil(t, actual)
	})

	T.Run("with empty strings", func(t *testing.T) {
		t.Parallel()

		x := &MealPlanOptionUpdateRequestInput{}

		actual := x.ValidateWithContext(context.Background())
		assert.Error(t, actual)
	})
}
