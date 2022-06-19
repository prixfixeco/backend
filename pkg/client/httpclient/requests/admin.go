package requests

import (
	"context"
	"net/http"

	"github.com/prixfixeco/api_server/internal/observability"
	"github.com/prixfixeco/api_server/internal/observability/keys"
	"github.com/prixfixeco/api_server/pkg/types"
)

const (
	adminBasePath = "admin"
)

// BuildUserAccountStatusUpdateInputRequest builds a request to change a user's account status.
func (b *Builder) BuildUserAccountStatusUpdateInputRequest(ctx context.Context, input *types.UserAccountStatusUpdateInput) (*http.Request, error) {
	ctx, span := b.tracer.StartSpan(ctx)
	defer span.End()

	if input == nil {
		return nil, ErrNilInputProvided
	}

	logger := b.logger.WithValue(keys.UserIDKey, input.TargetUserID)

	if err := input.ValidateWithContext(ctx); err != nil {
		return nil, observability.PrepareError(err, logger, span, "validating input")
	}

	uri := b.BuildURL(ctx, nil, adminBasePath, usersBasePath, "status")

	return b.buildDataRequest(ctx, http.MethodPost, uri, input)
}
