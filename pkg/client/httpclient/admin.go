package httpclient

import (
	"context"

	"github.com/prixfixeco/api_server/internal/observability"
	"github.com/prixfixeco/api_server/internal/observability/keys"
	"github.com/prixfixeco/api_server/internal/observability/tracing"
	"github.com/prixfixeco/api_server/pkg/types"
)

// UpdateUserReputation updates a user's reputation.
func (c *Client) UpdateUserReputation(ctx context.Context, input *types.UserReputationUpdateInput) error {
	ctx, span := c.tracer.StartSpan(ctx)
	defer span.End()

	if input == nil {
		return ErrNilInputProvided
	}

	logger := c.logger.WithValue(keys.HouseholdIDKey, input.TargetUserID)
	tracing.AttachUserIDToSpan(span, input.TargetUserID)

	if err := input.ValidateWithContext(ctx); err != nil {
		return observability.PrepareError(err, logger, span, "validating input")
	}

	req, err := c.requestBuilder.BuildUserReputationUpdateInputRequest(ctx, input)
	if err != nil {
		return observability.PrepareError(err, logger, span, "building user reputation update request")
	}

	res, err := c.fetchResponseToRequest(ctx, c.authedClient, req)
	if err != nil {
		return observability.PrepareError(err, logger, span, "updating user reputation")
	}

	c.closeResponseBody(ctx, res)

	if err = errorFromResponse(res); err != nil {
		return observability.PrepareError(err, logger, span, "invalid response status")
	}

	return nil
}
