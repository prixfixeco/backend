// GENERATED CODE, DO NOT EDIT MANUALLY

package apiclient

import (
	"context"
	"github.com/dinnerdonebetter/backend/pkg/types"
	"github.com/dinnerdonebetter/backend/pkg/types/fakes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestClient_RejectHouseholdInvitation(T *testing.T) {
	T.Parallel()

	const expectedPathFormat = "/api/v1/household_invitations/%s/reject"

	T.Run("standard", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		householdInvitationID := fakes.BuildFakeID()

		data := fakes.BuildFakeHouseholdInvitation()
		data.DestinationHousehold.WebhookEncryptionKey = ""
		data.FromUser.TwoFactorSecret = ""
		expected := &types.APIResponse[*types.HouseholdInvitation]{
			Data: data,
		}

		exampleInput := fakes.BuildFakeHouseholdInvitationUpdateRequestInput()

		spec := newRequestSpec(false, http.MethodPut, "", expectedPathFormat, householdInvitationID)
		c, _ := buildTestClientWithJSONResponse(t, spec, expected)
		err := c.RejectHouseholdInvitation(ctx, householdInvitationID, exampleInput)

		assert.NoError(t, err)

	})

	T.Run("with invalid householdInvitation ID", func(t *testing.T) {
		t.Parallel()

		exampleInput := fakes.BuildFakeHouseholdInvitationUpdateRequestInput()

		ctx := context.Background()
		c, _ := buildSimpleTestClient(t)
		err := c.RejectHouseholdInvitation(ctx, "", exampleInput)

		assert.Error(t, err)
	})

	T.Run("with error building request", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		householdInvitationID := fakes.BuildFakeID()

		exampleInput := fakes.BuildFakeHouseholdInvitationUpdateRequestInput()

		c := buildTestClientWithInvalidURL(t)
		err := c.RejectHouseholdInvitation(ctx, householdInvitationID, exampleInput)

		assert.Error(t, err)
	})

	T.Run("with error executing request", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		householdInvitationID := fakes.BuildFakeID()

		exampleInput := fakes.BuildFakeHouseholdInvitationUpdateRequestInput()

		spec := newRequestSpec(false, http.MethodPut, "", expectedPathFormat, householdInvitationID)
		c := buildTestClientWithInvalidResponse(t, spec)
		err := c.RejectHouseholdInvitation(ctx, householdInvitationID, exampleInput)

		assert.Error(t, err)
	})
}
