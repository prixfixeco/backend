// GENERATED CODE, DO NOT EDIT MANUALLY

package apiclient

import (
	"context"
	"net/http"

	"fmt"
	"github.com/dinnerdonebetter/backend/internal/observability"
	"github.com/dinnerdonebetter/backend/internal/observability/keys"
	"github.com/dinnerdonebetter/backend/internal/observability/tracing"
	"github.com/dinnerdonebetter/backend/pkg/types"
)

func (c *Client) UpdateHousehold(
	ctx context.Context,
	householdID string,
	input *types.HouseholdUpdateRequestInput,
) error {
	ctx, span := c.tracer.StartSpan(ctx)
	defer span.End()

	logger := c.logger.Clone()

	if householdID == "" {
		return ErrInvalidIDProvided
	}
	logger = logger.WithValue(keys.HouseholdIDKey, householdID)
	tracing.AttachToSpan(span, keys.HouseholdIDKey, householdID)

	u := c.BuildURL(ctx, nil, fmt.Sprintf("/api/v1/households/%s", householdID))
	req, err := c.buildDataRequest(ctx, http.MethodPut, u, input)
	if err != nil {
		return observability.PrepareAndLogError(err, logger, span, "building request to create a Household")
	}

	var apiResponse *types.APIResponse[*types.Household]
	if err = c.fetchAndUnmarshal(ctx, req, &apiResponse); err != nil {
		return observability.PrepareAndLogError(err, logger, span, "loading Household creation response")
	}

	if err = apiResponse.Error.AsError(); err != nil {
		return err
	}

	return nil
}