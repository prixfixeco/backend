// GENERATED CODE, DO NOT EDIT MANUALLY

package apiclient

import (
	"context"
	"net/http"

	"github.com/dinnerdonebetter/backend/internal/observability"
	"github.com/dinnerdonebetter/backend/pkg/types"
)

func (c *Client) RequestPasswordResetToken(
	ctx context.Context,
	input *types.PasswordResetTokenCreationRequestInput,
	reqMods ...RequestModifier,
) (*types.PasswordResetToken, error) {
	ctx, span := c.tracer.StartSpan(ctx)
	defer span.End()

	logger := c.logger.Clone()

	if input == nil {
		return nil, ErrNilInputProvided
	}

	if err := input.ValidateWithContext(ctx); err != nil {
		return nil, observability.PrepareError(err, span, "validating input")
	}

	u := c.BuildURL(ctx, nil, "/users/password/reset")
	req, err := c.buildDataRequest(ctx, http.MethodPost, u, input)
	if err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "building request to create a PasswordResetToken")
	}

	for _, mod := range reqMods {
		mod(req)
	}

	var apiResponse *types.APIResponse[*types.PasswordResetToken]
	if err = c.fetchAndUnmarshal(ctx, req, &apiResponse); err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "loading PasswordResetToken creation response")
	}

	if err = apiResponse.Error.AsError(); err != nil {
		return nil, err
	}

	return apiResponse.Data, nil
}
