// GENERATED CODE, DO NOT EDIT MANUALLY

package apiclient

import (
	"context"
	"net/http"

	"fmt"
	"github.com/dinnerdonebetter/backend/internal/observability"
	"github.com/dinnerdonebetter/backend/pkg/types"
)

func (c *Client) RunFinalizeMealPlanWorker(
	ctx context.Context,
	input *types.FinalizeMealPlansRequest,
) (*types.FinalizeMealPlansResponse, error) {
	ctx, span := c.tracer.StartSpan(ctx)
	defer span.End()

	logger := c.logger.Clone()

	if input == nil {
		return nil, ErrNilInputProvided
	}

	if err := input.ValidateWithContext(ctx); err != nil {
		return nil, observability.PrepareError(err, span, "validating input")
	}

	u := c.BuildURL(ctx, nil, fmt.Sprintf("/api/v1/workers/finalize_meal_plans"))
	req, err := c.buildDataRequest(ctx, http.MethodPost, u, input)
	if err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "building request to create a FinalizeMealPlansResponse")
	}

	var apiResponse *types.APIResponse[*types.FinalizeMealPlansResponse]
	if err = c.fetchAndUnmarshal(ctx, req, &apiResponse); err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "loading FinalizeMealPlansResponse creation response")
	}

	if err = apiResponse.Error.AsError(); err != nil {
		return nil, err
	}

	return apiResponse.Data, nil
}
