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

func TestClient_CreateHouseholdInvitation(T *testing.T) {
	T.Parallel()

	const expectedPathFormat = "/api/v1/households/%s/invite"

	T.Run("standard", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		householdID := fakes.BuildFakeID()

		data := fakes.BuildFakeHouseholdInvitation()
		data.DestinationHousehold.WebhookEncryptionKey = ""
		data.FromUser.TwoFactorSecret = ""
		expected := &types.APIResponse[*types.HouseholdInvitation]{
			Data: data,
		}

		exampleInput := fakes.BuildFakeHouseholdInvitationCreationRequestInput()

		spec := newRequestSpec(false, http.MethodPost, "", expectedPathFormat, householdID)
		c, _ := buildTestClientWithJSONResponse(t, spec, expected)
		actual, err := c.CreateHouseholdInvitation(ctx, householdID, exampleInput)

		require.NotNil(t, actual)
		assert.NoError(t, err)
		assert.Equal(t, data, actual)
	})

	T.Run("with invalid household ID", func(t *testing.T) {
		t.Parallel()

		exampleInput := fakes.BuildFakeHouseholdInvitationCreationRequestInput()

		ctx := context.Background()
		c, _ := buildSimpleTestClient(t)
		actual, err := c.CreateHouseholdInvitation(ctx, "", exampleInput)

		require.Nil(t, actual)
		assert.Error(t, err)
	})

	T.Run("with error building request", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		householdID := fakes.BuildFakeID()

		exampleInput := fakes.BuildFakeHouseholdInvitationCreationRequestInput()

		c := buildTestClientWithInvalidURL(t)
		actual, err := c.CreateHouseholdInvitation(ctx, householdID, exampleInput)

		require.Nil(t, actual)
		assert.Error(t, err)
	})

	T.Run("with error executing request", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		householdID := fakes.BuildFakeID()

		exampleInput := fakes.BuildFakeHouseholdInvitationCreationRequestInput()

		spec := newRequestSpec(false, http.MethodPost, "", expectedPathFormat, householdID)
		c := buildTestClientWithInvalidResponse(t, spec)
		actual, err := c.CreateHouseholdInvitation(ctx, householdID, exampleInput)

		require.Nil(t, actual)
		assert.Error(t, err)
	})
}