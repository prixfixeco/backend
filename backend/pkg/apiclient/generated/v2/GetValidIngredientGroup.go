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

func (c *Client) GetValidIngredientGroup(
	ctx context.Context,
	validIngredientGroupID string,
) (*types.ValidIngredientGroup, error) {
	ctx, span := c.tracer.StartSpan(ctx)
	defer span.End()

	logger := c.logger.Clone()

	if validIngredientGroupID == "" {
		return nil, buildInvalidIDError("validIngredientGroup")
	}
	logger = logger.WithValue(keys.ValidIngredientGroupIDKey, validIngredientGroupID)
	tracing.AttachToSpan(span, keys.ValidIngredientGroupIDKey, validIngredientGroupID)

	u := c.BuildURL(ctx, nil, fmt.Sprintf("/api/v1/valid_ingredient_groups/%s", validIngredientGroupID))
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, http.NoBody)
	if err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "building request to fetch a ValidIngredientGroup")
	}

	var apiResponse *types.APIResponse[*types.ValidIngredientGroup]
	if err = c.fetchAndUnmarshal(ctx, req, &apiResponse); err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "loading ValidIngredientGroup response")
	}

	if err = apiResponse.Error.AsError(); err != nil {
		return nil, err
	}

	return apiResponse.Data, nil
}
