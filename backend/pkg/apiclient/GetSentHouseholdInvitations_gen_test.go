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

func TestClient_GetSentHouseholdInvitations(T *testing.T) {
	T.Parallel()

	const expectedPathFormat = "/api/v1/household_invitations/sent"

	T.Run("standard", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()

		list := fakes.BuildFakeHouseholdInvitationsList()
		for i := range list.Data {
			list.Data[i].DestinationHousehold.WebhookEncryptionKey = ""
			list.Data[i].FromUser.TwoFactorSecret = ""
		}

		expected := &types.APIResponse[[]*types.HouseholdInvitation]{
			Pagination: &list.Pagination,
			Data:       list.Data,
		}

		spec := newRequestSpec(true, http.MethodGet, "limit=50&page=1&sortBy=asc", expectedPathFormat)
		c, _ := buildTestClientWithJSONResponse(t, spec, expected)
		actual, err := c.GetSentHouseholdInvitations(ctx, nil)

		require.NotNil(t, actual)
		assert.NoError(t, err)
		assert.Equal(t, list, actual)
	})

	T.Run("with error building request", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()

		c := buildTestClientWithInvalidURL(t)
		actual, err := c.GetSentHouseholdInvitations(ctx, nil)

		require.Nil(t, actual)
		assert.Error(t, err)
	})

	T.Run("with error executing request", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()

		spec := newRequestSpec(true, http.MethodGet, "limit=50&page=1&sortBy=asc", expectedPathFormat)
		c := buildTestClientWithInvalidResponse(t, spec)
		actual, err := c.GetSentHouseholdInvitations(ctx, nil)

		require.Nil(t, actual)
		assert.Error(t, err)
	})
}
