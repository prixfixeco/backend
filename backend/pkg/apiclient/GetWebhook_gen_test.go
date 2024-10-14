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

func TestClient_GetWebhook(T *testing.T) {
	T.Parallel()

	const expectedPathFormat = "/api/v1/webhooks/%s"

	T.Run("standard", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		webhookID := fakes.BuildFakeID()

		data := fakes.BuildFakeWebhook()
		expected := &types.APIResponse[*types.Webhook]{
			Data: data,
		}

		spec := newRequestSpec(true, http.MethodGet, "", expectedPathFormat, webhookID)
		c, _ := buildTestClientWithJSONResponse(t, spec, expected)
		actual, err := c.GetWebhook(ctx, webhookID)

		require.NotNil(t, actual)
		assert.NoError(t, err)
		assert.Equal(t, data, actual)
	})

	T.Run("with invalid webhook ID", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		c, _ := buildSimpleTestClient(t)
		actual, err := c.GetWebhook(ctx, "")

		require.Nil(t, actual)
		assert.Error(t, err)
	})

	T.Run("with error building request", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		webhookID := fakes.BuildFakeID()

		c := buildTestClientWithInvalidURL(t)
		actual, err := c.GetWebhook(ctx, webhookID)

		require.Nil(t, actual)
		assert.Error(t, err)
	})

	T.Run("with error executing request", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		webhookID := fakes.BuildFakeID()

		spec := newRequestSpec(true, http.MethodGet, "", expectedPathFormat, webhookID)
		c := buildTestClientWithInvalidResponse(t, spec)
		actual, err := c.GetWebhook(ctx, webhookID)

		require.Nil(t, actual)
		assert.Error(t, err)
	})
}
