package requests

import (
	"context"
	"net/http"

	"github.com/prixfixeco/api_server/internal/observability"
	"github.com/prixfixeco/api_server/internal/observability/tracing"
	"github.com/prixfixeco/api_server/pkg/types"
	"github.com/prixfixeco/api_server/pkg/types/converters"
)

const (
	mealPlanEventsBasePath = "events"
)

// BuildGetMealPlanEventRequest builds an HTTP request for fetching a meal plan option.
func (b *Builder) BuildGetMealPlanEventRequest(ctx context.Context, mealPlanID, mealPlanEventID string) (*http.Request, error) {
	ctx, span := b.tracer.StartSpan(ctx)
	defer span.End()

	if mealPlanID == "" {
		return nil, ErrInvalidIDProvided
	}
	tracing.AttachMealPlanIDToSpan(span, mealPlanID)

	if mealPlanEventID == "" {
		return nil, ErrInvalidIDProvided
	}
	tracing.AttachMealPlanEventIDToSpan(span, mealPlanEventID)

	uri := b.BuildURL(
		ctx,
		nil,
		mealPlansBasePath,
		mealPlanID,
		mealPlanEventsBasePath,
		mealPlanEventID,
	)
	tracing.AttachRequestURIToSpan(span, uri)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, http.NoBody)
	if err != nil {
		return nil, observability.PrepareError(err, span, "building request")
	}

	return req, nil
}

// BuildGetMealPlanEventsRequest builds an HTTP request for fetching a list of meal plan options.
func (b *Builder) BuildGetMealPlanEventsRequest(ctx context.Context, mealPlanID string, filter *types.QueryFilter) (*http.Request, error) {
	ctx, span := b.tracer.StartSpan(ctx)
	defer span.End()

	if mealPlanID == "" {
		return nil, ErrInvalidIDProvided
	}
	tracing.AttachMealPlanIDToSpan(span, mealPlanID)

	uri := b.BuildURL(
		ctx,
		filter.ToValues(),
		mealPlansBasePath,
		mealPlanID,
		mealPlanEventsBasePath,
	)
	tracing.AttachRequestURIToSpan(span, uri)
	tracing.AttachQueryFilterToSpan(span, filter)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, http.NoBody)
	if err != nil {
		return nil, observability.PrepareError(err, span, "building request")
	}

	return req, nil
}

// BuildCreateMealPlanEventRequest builds an HTTP request for creating a meal plan option.
func (b *Builder) BuildCreateMealPlanEventRequest(ctx context.Context, mealPlanID string, input *types.MealPlanEventCreationRequestInput) (*http.Request, error) {
	ctx, span := b.tracer.StartSpan(ctx)
	defer span.End()

	if input == nil {
		return nil, ErrNilInputProvided
	}

	if err := input.ValidateWithContext(ctx); err != nil {
		return nil, observability.PrepareError(err, span, "validating input")
	}

	uri := b.BuildURL(
		ctx,
		nil,
		mealPlansBasePath,
		mealPlanID,
		mealPlanEventsBasePath,
	)
	tracing.AttachRequestURIToSpan(span, uri)

	req, err := b.buildDataRequest(ctx, http.MethodPost, uri, input)
	if err != nil {
		return nil, observability.PrepareError(err, span, "building request")
	}

	return req, nil
}

// BuildUpdateMealPlanEventRequest builds an HTTP request for updating a meal plan option.
func (b *Builder) BuildUpdateMealPlanEventRequest(ctx context.Context, mealPlanEvent *types.MealPlanEvent) (*http.Request, error) {
	ctx, span := b.tracer.StartSpan(ctx)
	defer span.End()

	if mealPlanEvent == nil {
		return nil, ErrNilInputProvided
	}
	tracing.AttachMealPlanEventIDToSpan(span, mealPlanEvent.ID)

	uri := b.BuildURL(
		ctx,
		nil,
		mealPlansBasePath,
		mealPlanEvent.BelongsToMealPlan,
		mealPlanEventsBasePath,
		mealPlanEvent.ID,
	)
	tracing.AttachRequestURIToSpan(span, uri)

	input := converters.ConvertMealPlanEventToMealPlanEventUpdateRequestInput(mealPlanEvent)

	req, err := b.buildDataRequest(ctx, http.MethodPut, uri, input)
	if err != nil {
		return nil, observability.PrepareError(err, span, "building request")
	}

	return req, nil
}

// BuildArchiveMealPlanEventRequest builds an HTTP request for archiving a meal plan option.
func (b *Builder) BuildArchiveMealPlanEventRequest(ctx context.Context, mealPlanID, mealPlanEventID string) (*http.Request, error) {
	ctx, span := b.tracer.StartSpan(ctx)
	defer span.End()

	if mealPlanID == "" {
		return nil, ErrInvalidIDProvided
	}
	tracing.AttachMealPlanIDToSpan(span, mealPlanID)

	if mealPlanEventID == "" {
		return nil, ErrInvalidIDProvided
	}
	tracing.AttachMealPlanEventIDToSpan(span, mealPlanEventID)

	uri := b.BuildURL(
		ctx,
		nil,
		mealPlansBasePath,
		mealPlanID,
		mealPlanEventsBasePath,
		mealPlanEventID,
	)
	tracing.AttachRequestURIToSpan(span, uri)

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, uri, http.NoBody)
	if err != nil {
		return nil, observability.PrepareError(err, span, "building request")
	}

	return req, nil
}