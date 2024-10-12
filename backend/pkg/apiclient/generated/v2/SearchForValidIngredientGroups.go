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

func (c *Client) SearchForValidIngredientGroups(
	ctx context.Context,
	q string,
	filter *types.QueryFilter,
) (*types.QueryFilteredResult[types.ValidIngredientGroup], error) {
	ctx, span := c.tracer.StartSpan(ctx)
	defer span.End()

	logger := c.logger.Clone()

	if filter == nil {
		filter = types.DefaultQueryFilter()
	}
	logger = filter.AttachToLogger(logger)
	tracing.AttachQueryFilterToSpan(span, filter)

	if q == "" {
		return nil, buildInvalidIDError("q")
	}
	logger = logger.WithValue(keys.SearchQueryKey, q)
	tracing.AttachToSpan(span, keys.SearchQueryKey, q)

	values := filter.ToValues()
	values.Set(types.QueryKeySearch, q)

	u := c.BuildURL(ctx, values, fmt.Sprintf("/api/v1/valid_ingredient_groups/search"))
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, http.NoBody)
	if err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "building request to fetch list of ValidIngredientGroup")
	}

	var apiResponse *types.APIResponse[[]*types.ValidIngredientGroup]
	if err = c.fetchAndUnmarshal(ctx, req, &apiResponse); err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "loading response for list of ValidIngredientGroup")
	}

	if err = apiResponse.Error.AsError(); err != nil {
		return nil, err
	}

	result := &types.QueryFilteredResult[types.ValidIngredientGroup]{
		Data:       apiResponse.Data,
		Pagination: *apiResponse.Pagination,
	}

	return result, nil
}