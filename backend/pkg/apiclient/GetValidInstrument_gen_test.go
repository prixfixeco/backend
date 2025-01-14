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

func TestClient_GetValidInstrument(T *testing.T) {
	T.Parallel()

	const expectedPathFormat = "/api/v1/valid_instruments/%s"

	T.Run("standard", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		validInstrumentID := fakes.BuildFakeID()

		data := fakes.BuildFakeValidInstrument()
		expected := &types.APIResponse[*types.ValidInstrument]{
			Data: data,
		}

		spec := newRequestSpec(true, http.MethodGet, "", expectedPathFormat, validInstrumentID)
		c, _ := buildTestClientWithJSONResponse(t, spec, expected)
		actual, err := c.GetValidInstrument(ctx, validInstrumentID)

		require.NotNil(t, actual)
		assert.NoError(t, err)
		assert.Equal(t, data, actual)
	})

	T.Run("with invalid validInstrument ID", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		c, _ := buildSimpleTestClient(t)
		actual, err := c.GetValidInstrument(ctx, "")

		require.Nil(t, actual)
		assert.Error(t, err)
	})

	T.Run("with error building request", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		validInstrumentID := fakes.BuildFakeID()

		c := buildTestClientWithInvalidURL(t)
		actual, err := c.GetValidInstrument(ctx, validInstrumentID)

		require.Nil(t, actual)
		assert.Error(t, err)
	})

	T.Run("with error executing request", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		validInstrumentID := fakes.BuildFakeID()

		spec := newRequestSpec(true, http.MethodGet, "", expectedPathFormat, validInstrumentID)
		c := buildTestClientWithInvalidResponse(t, spec)
		actual, err := c.GetValidInstrument(ctx, validInstrumentID)

		require.Nil(t, actual)
		assert.Error(t, err)
	})
}
