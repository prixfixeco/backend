package types

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHouseholdInvitationCreationInput_ValidateWithContext(T *testing.T) {
	T.Parallel()

	T.Run("standard", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		x := &HouseholdInvitationCreationRequestInput{
			ToEmail:              t.Name(),
			FromUser:             t.Name(),
			DestinationHousehold: t.Name(),
		}

		assert.NoError(t, x.ValidateWithContext(ctx))
	})
}
