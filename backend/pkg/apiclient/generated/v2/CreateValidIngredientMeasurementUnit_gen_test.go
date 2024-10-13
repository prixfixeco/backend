// GENERATED CODE, DO NOT EDIT MANUALLY

package apiclient

import (
	"context"
	"github.com/dinnerdonebetter/backend/pkg/types"
	"github.com/dinnerdonebetter/backend/pkg/types/fakes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestClient_CreateValidIngredientMeasurementUnit(T *testing.T) {
	T.Parallel()

	const expectedPathFormat = "/api/v1/valid_ingredient_measurement_units"

	T.Run("standard", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()

		data := fakes.BuildFakeValidIngredientMeasurementUnit()
		expected := &types.APIResponse[*types.ValidIngredientMeasurementUnit]{
			Data: data,
		}

		exampleInput := fakes.BuildFakeValidIngredientMeasurementUnitCreationRequestInput()

		spec := newRequestSpec(false, http.MethodPost, "", expectedPathFormat)
		c, _ := buildTestClientWithJSONResponse(t, spec, expected)
		actual, err := c.CreateValidIngredientMeasurementUnit(ctx, exampleInput)

		require.NotNil(t, actual)
		assert.NoError(t, err)
		assert.Equal(t, data, actual)
	})

	T.Run("with error building request", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()

		exampleInput := fakes.BuildFakeValidIngredientMeasurementUnitCreationRequestInput()

		c := buildTestClientWithInvalidURL(t)
		actual, err := c.CreateValidIngredientMeasurementUnit(ctx, exampleInput)

		require.Nil(t, actual)
		assert.Error(t, err)
	})

	T.Run("with error executing request", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()

		exampleInput := fakes.BuildFakeValidIngredientMeasurementUnitCreationRequestInput()

		spec := newRequestSpec(false, http.MethodPost, "", expectedPathFormat)
		c := buildTestClientWithInvalidResponse(t, spec)
		actual, err := c.CreateValidIngredientMeasurementUnit(ctx, exampleInput)

		require.Nil(t, actual)
		assert.Error(t, err)
	})
}
