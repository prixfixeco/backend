package requests

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/prixfixeco/api_server/internal/observability"
	"github.com/prixfixeco/api_server/internal/observability/tracing"
	"github.com/prixfixeco/api_server/pkg/types"
	"github.com/prixfixeco/api_server/pkg/types/converters"
)

const (
	recipesBasePath = "recipes"
)

// BuildGetRecipeRequest builds an HTTP request for fetching a recipe.
func (b *Builder) BuildGetRecipeRequest(ctx context.Context, recipeID string) (*http.Request, error) {
	ctx, span := b.tracer.StartSpan(ctx)
	defer span.End()

	if recipeID == "" {
		return nil, ErrInvalidIDProvided
	}
	tracing.AttachRecipeIDToSpan(span, recipeID)

	uri := b.BuildURL(
		ctx,
		nil,
		recipesBasePath,
		recipeID,
	)
	tracing.AttachRequestURIToSpan(span, uri)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, http.NoBody)
	if err != nil {
		return nil, observability.PrepareError(err, span, "building request")
	}

	return req, nil
}

// BuildGetRecipesRequest builds an HTTP request for fetching a list of recipes.
func (b *Builder) BuildGetRecipesRequest(ctx context.Context, filter *types.QueryFilter) (*http.Request, error) {
	ctx, span := b.tracer.StartSpan(ctx)
	defer span.End()

	uri := b.BuildURL(
		ctx,
		filter.ToValues(),
		recipesBasePath,
	)
	tracing.AttachRequestURIToSpan(span, uri)
	tracing.AttachQueryFilterToSpan(span, filter)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, http.NoBody)
	if err != nil {
		return nil, observability.PrepareError(err, span, "building request")
	}

	return req, nil
}

// BuildSearchForRecipesRequest builds an HTTP request for fetching a list of recipes.
func (b *Builder) BuildSearchForRecipesRequest(ctx context.Context, query string, filter *types.QueryFilter) (*http.Request, error) {
	ctx, span := b.tracer.StartSpan(ctx)
	defer span.End()

	tracing.AttachSearchQueryToSpan(span, query)

	queryParams := filter.ToValues()
	queryParams.Set(types.SearchQueryKey, query)

	uri := b.BuildURL(
		ctx,
		queryParams,
		recipesBasePath,
		"search",
	)
	tracing.AttachRequestURIToSpan(span, uri)
	tracing.AttachQueryFilterToSpan(span, filter)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, http.NoBody)
	if err != nil {
		return nil, observability.PrepareError(err, span, "building request")
	}

	return req, nil
}

// BuildCreateRecipeRequest builds an HTTP request for creating a recipe.
func (b *Builder) BuildCreateRecipeRequest(ctx context.Context, input *types.RecipeCreationRequestInput) (*http.Request, error) {
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
		recipesBasePath,
	)
	tracing.AttachRequestURIToSpan(span, uri)

	req, err := b.buildDataRequest(ctx, http.MethodPost, uri, input)
	if err != nil {
		return nil, observability.PrepareError(err, span, "building request")
	}

	return req, nil
}

// BuildUpdateRecipeRequest builds an HTTP request for updating a recipe.
func (b *Builder) BuildUpdateRecipeRequest(ctx context.Context, recipe *types.Recipe) (*http.Request, error) {
	ctx, span := b.tracer.StartSpan(ctx)
	defer span.End()

	if recipe == nil {
		return nil, ErrNilInputProvided
	}
	tracing.AttachRecipeIDToSpan(span, recipe.ID)

	uri := b.BuildURL(
		ctx,
		nil,
		recipesBasePath,
		recipe.ID,
	)
	tracing.AttachRequestURIToSpan(span, uri)

	input := converters.ConvertRecipeToRecipeUpdateRequestInput(recipe)

	req, err := b.buildDataRequest(ctx, http.MethodPut, uri, input)
	if err != nil {
		return nil, observability.PrepareError(err, span, "building request")
	}

	return req, nil
}

// BuildArchiveRecipeRequest builds an HTTP request for archiving a recipe.
func (b *Builder) BuildArchiveRecipeRequest(ctx context.Context, recipeID string) (*http.Request, error) {
	ctx, span := b.tracer.StartSpan(ctx)
	defer span.End()

	if recipeID == "" {
		return nil, ErrInvalidIDProvided
	}
	tracing.AttachRecipeIDToSpan(span, recipeID)

	uri := b.BuildURL(
		ctx,
		nil,
		recipesBasePath,
		recipeID,
	)
	tracing.AttachRequestURIToSpan(span, uri)

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, uri, http.NoBody)
	if err != nil {
		return nil, observability.PrepareError(err, span, "building request")
	}

	return req, nil
}

// BuildGetRecipeDAGRequest builds an HTTP request for fetching a recipe.
func (b *Builder) BuildGetRecipeDAGRequest(ctx context.Context, recipeID string) (*http.Request, error) {
	ctx, span := b.tracer.StartSpan(ctx)
	defer span.End()

	if recipeID == "" {
		return nil, ErrInvalidIDProvided
	}
	tracing.AttachRecipeIDToSpan(span, recipeID)

	uri := b.BuildURL(
		ctx,
		nil,
		recipesBasePath,
		recipeID,
		"dag",
	)
	tracing.AttachRequestURIToSpan(span, uri)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, http.NoBody)
	if err != nil {
		return nil, observability.PrepareError(err, span, "building request")
	}

	return req, nil
}

// BuildGetRecipeMealPlanTasksRequest builds an HTTP request for fetching a recipe.
func (b *Builder) BuildGetRecipeMealPlanTasksRequest(ctx context.Context, recipeID string) (*http.Request, error) {
	ctx, span := b.tracer.StartSpan(ctx)
	defer span.End()

	if recipeID == "" {
		return nil, ErrInvalidIDProvided
	}
	tracing.AttachRecipeIDToSpan(span, recipeID)

	uri := b.BuildURL(
		ctx,
		nil,
		recipesBasePath,
		recipeID,
		"prep_steps",
	)
	tracing.AttachRequestURIToSpan(span, uri)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, http.NoBody)
	if err != nil {
		return nil, observability.PrepareError(err, span, "building request")
	}

	return req, nil
}

// BuildMediaUploadRequest builds an HTTP request that sets a user's avatar to the provided content.
func (b *Builder) BuildMediaUploadRequest(ctx context.Context, media []byte, extension, recipeID string) (*http.Request, error) {
	ctx, span := b.tracer.StartSpan(ctx)
	defer span.End()

	if recipeID == "" {
		return nil, ErrInvalidIDProvided
	}
	tracing.AttachRecipeIDToSpan(span, recipeID)

	if len(media) == 0 {
		return nil, ErrNilInputProvided
	}

	var ct string

	switch strings.ToLower(strings.TrimSpace(extension)) {
	case "jpeg":
		ct = "image/jpeg"
	case "png":
		ct = "image/png"
	case "gif":
		ct = "image/gif"
	default:
		return nil, fmt.Errorf("%s: %w", extension, ErrInvalidPhotoEncodingForUpload)
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("media", fmt.Sprintf("media.%s", extension))
	if err != nil {
		return nil, observability.PrepareError(err, span, "creating form file")
	}

	if _, err = io.Copy(part, bytes.NewReader(media)); err != nil {
		return nil, observability.PrepareError(err, span, "copying file contents to request")
	}

	if err = writer.Close(); err != nil {
		return nil, observability.PrepareError(err, span, "closing media writer")
	}

	uri := b.BuildURL(ctx, nil, recipesBasePath, recipeID, "images")

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, uri, body)
	if err != nil {
		return nil, observability.PrepareError(err, span, "building media upload request")
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("X-Upload-Content-Type", ct)

	return req, nil
}
