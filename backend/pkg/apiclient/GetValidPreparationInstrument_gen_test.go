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

func TestClient_GetValidPreparationInstrument(T *testing.T) {
	T.Parallel()

	const expectedPathFormat = "/api/v1/valid_preparation_instruments/%s"

	T.Run("standard", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		validPreparationInstrumentID := fakes.BuildFakeID()

		data := fakes.BuildFakeValidPreparationInstrument()
		expected := &types.APIResponse[*types.ValidPreparationInstrument]{
			Data: data,
		}

		spec := newRequestSpec(true, http.MethodGet, "", expectedPathFormat, validPreparationInstrumentID)
		c, _ := buildTestClientWithJSONResponse(t, spec, expected)
		actual, err := c.GetValidPreparationInstrument(ctx, validPreparationInstrumentID)

		require.NotNil(t, actual)
		assert.NoError(t, err)
		assert.Equal(t, data, actual)
	})

	T.Run("with invalid validPreparationInstrument ID", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		c, _ := buildSimpleTestClient(t)
		actual, err := c.GetValidPreparationInstrument(ctx, "")

		require.Nil(t, actual)
		assert.Error(t, err)
	})

	T.Run("with error building request", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		validPreparationInstrumentID := fakes.BuildFakeID()

		c := buildTestClientWithInvalidURL(t)
		actual, err := c.GetValidPreparationInstrument(ctx, validPreparationInstrumentID)

		require.Nil(t, actual)
		assert.Error(t, err)
	})

	T.Run("with error executing request", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		validPreparationInstrumentID := fakes.BuildFakeID()

		spec := newRequestSpec(true, http.MethodGet, "", expectedPathFormat, validPreparationInstrumentID)
		c := buildTestClientWithInvalidResponse(t, spec)
		actual, err := c.GetValidPreparationInstrument(ctx, validPreparationInstrumentID)

		require.Nil(t, actual)
		assert.Error(t, err)
	})
}
