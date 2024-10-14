// GENERATED CODE, DO NOT EDIT MANUALLY

package apiclient

import (
	"context"
	"net/http"
	"testing"

	"github.com/dinnerdonebetter/backend/pkg/types"
	"github.com/dinnerdonebetter/backend/pkg/types/fakes"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClient_GetValidMeasurementUnitConversionsFromUnit(T *testing.T) {
	T.Parallel()

	const expectedPathFormat = "/api/v1/valid_measurement_conversions/from_unit/%s"

	T.Run("standard", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		validMeasurementUnitID := fakes.BuildFakeID()

		list := fakes.BuildFakeValidMeasurementUnitConversionsList()

		expected := &types.APIResponse[[]*types.ValidMeasurementUnitConversion]{
			Pagination: &list.Pagination,
			Data:       list.Data,
		}

		spec := newRequestSpec(true, http.MethodGet, "limit=50&page=1&sortBy=asc", expectedPathFormat, validMeasurementUnitID)
		c, _ := buildTestClientWithJSONResponse(t, spec, expected)
		actual, err := c.GetValidMeasurementUnitConversionsFromUnit(ctx, validMeasurementUnitID, nil)

		require.NotNil(t, actual)
		assert.NoError(t, err)
		assert.Equal(t, list, actual)
	})

	T.Run("with empty validMeasurementUnit ID", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		c, _ := buildSimpleTestClient(t)
		actual, err := c.GetValidMeasurementUnitConversionsFromUnit(ctx, "", nil)

		require.Nil(t, actual)
		assert.Error(t, err)
	})

	T.Run("with error building request", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		validMeasurementUnitID := fakes.BuildFakeID()

		c := buildTestClientWithInvalidURL(t)
		actual, err := c.GetValidMeasurementUnitConversionsFromUnit(ctx, validMeasurementUnitID, nil)

		require.Nil(t, actual)
		assert.Error(t, err)
	})

	T.Run("with error executing request", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		validMeasurementUnitID := fakes.BuildFakeID()

		spec := newRequestSpec(true, http.MethodGet, "limit=50&page=1&sortBy=asc", expectedPathFormat, validMeasurementUnitID)
		c := buildTestClientWithInvalidResponse(t, spec)
		actual, err := c.GetValidMeasurementUnitConversionsFromUnit(ctx, validMeasurementUnitID, nil)

		require.Nil(t, actual)
		assert.Error(t, err)
	})
}