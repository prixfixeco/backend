package httpclient

import (
	"context"

	"github.com/prixfixeco/api_server/internal/observability"
	"github.com/prixfixeco/api_server/internal/observability/keys"
	"github.com/prixfixeco/api_server/internal/observability/tracing"
	"github.com/prixfixeco/api_server/pkg/types"
)

// SwitchActiveHousehold will switch the household on whose behalf requests are made.
func (c *Client) SwitchActiveHousehold(ctx context.Context, householdID string) error {
	ctx, span := c.tracer.StartSpan(ctx)
	defer span.End()

	if householdID == "" {
		return ErrInvalidIDProvided
	}

	logger := c.logger.WithValue(keys.HouseholdIDKey, householdID)
	tracing.AttachHouseholdIDToSpan(span, householdID)

	if c.authMethod == cookieAuthMethod {
		req, err := c.requestBuilder.BuildSwitchActiveHouseholdRequest(ctx, householdID)
		if err != nil {
			return observability.PrepareError(err, logger, span, "building household switch request")
		}

		if err = c.executeAndUnmarshal(ctx, req, c.authedClient, nil); err != nil {
			return observability.PrepareError(err, logger, span, "executing household switch request")
		}
	}

	c.householdID = householdID

	return nil
}

// GetHousehold retrieves a household.
func (c *Client) GetHousehold(ctx context.Context, householdID string) (*types.Household, error) {
	ctx, span := c.tracer.StartSpan(ctx)
	defer span.End()

	if householdID == "" {
		return nil, ErrInvalidIDProvided
	}

	logger := c.logger.WithValue(keys.HouseholdIDKey, householdID)
	tracing.AttachHouseholdIDToSpan(span, householdID)

	req, err := c.requestBuilder.BuildGetHouseholdRequest(ctx, householdID)
	if err != nil {
		return nil, observability.PrepareError(err, logger, span, "building household retrieval request")
	}

	var household *types.Household
	if err = c.fetchAndUnmarshal(ctx, req, &household); err != nil {
		return nil, observability.PrepareError(err, logger, span, "retrieving household")
	}

	return household, nil
}

// GetHouseholds retrieves a list of households.
func (c *Client) GetHouseholds(ctx context.Context, filter *types.QueryFilter) (*types.HouseholdList, error) {
	ctx, span := c.tracer.StartSpan(ctx)
	defer span.End()

	logger := c.loggerWithFilter(filter)

	tracing.AttachQueryFilterToSpan(span, filter)

	req, err := c.requestBuilder.BuildGetHouseholdsRequest(ctx, filter)
	if err != nil {
		return nil, observability.PrepareError(err, logger, span, "building household list request")
	}

	var households *types.HouseholdList
	if err = c.fetchAndUnmarshal(ctx, req, &households); err != nil {
		return nil, observability.PrepareError(err, logger, span, "retrieving households")
	}

	return households, nil
}

// CreateHousehold creates a household.
func (c *Client) CreateHousehold(ctx context.Context, input *types.HouseholdCreationRequestInput) (*types.Household, error) {
	ctx, span := c.tracer.StartSpan(ctx)
	defer span.End()

	if input == nil {
		return nil, ErrNilInputProvided
	}

	logger := c.logger.WithValue("household_name", input.Name)

	if err := input.ValidateWithContext(ctx); err != nil {
		return nil, observability.PrepareError(err, logger, span, "validating input")
	}

	req, err := c.requestBuilder.BuildCreateHouseholdRequest(ctx, input)
	if err != nil {
		return nil, observability.PrepareError(err, logger, span, "building household creation request")
	}

	var household *types.Household
	if err = c.fetchAndUnmarshal(ctx, req, &household); err != nil {
		return nil, observability.PrepareError(err, logger, span, "creating household")
	}

	return household, nil
}

// UpdateHousehold updates a household.
func (c *Client) UpdateHousehold(ctx context.Context, household *types.Household) error {
	ctx, span := c.tracer.StartSpan(ctx)
	defer span.End()

	if household == nil {
		return ErrNilInputProvided
	}

	logger := c.logger.WithValue(keys.HouseholdIDKey, household.ID)
	tracing.AttachHouseholdIDToSpan(span, household.ID)

	req, err := c.requestBuilder.BuildUpdateHouseholdRequest(ctx, household)
	if err != nil {
		return observability.PrepareError(err, logger, span, "building household update request")
	}

	if err = c.fetchAndUnmarshal(ctx, req, &household); err != nil {
		return observability.PrepareError(err, logger, span, "updating household")
	}

	return nil
}

// ArchiveHousehold archives a household.
func (c *Client) ArchiveHousehold(ctx context.Context, householdID string) error {
	ctx, span := c.tracer.StartSpan(ctx)
	defer span.End()

	if householdID == "" {
		return ErrInvalidIDProvided
	}

	logger := c.logger.WithValue(keys.HouseholdIDKey, householdID)
	tracing.AttachHouseholdIDToSpan(span, householdID)

	req, err := c.requestBuilder.BuildArchiveHouseholdRequest(ctx, householdID)
	if err != nil {
		return observability.PrepareError(err, logger, span, "building household archive request")
	}

	if err = c.fetchAndUnmarshal(ctx, req, nil); err != nil {
		return observability.PrepareError(err, logger, span, "archiving household")
	}

	return nil
}

// InviteUserToHousehold adds a user to a household.
func (c *Client) InviteUserToHousehold(ctx context.Context, input *types.HouseholdInvitationCreationRequestInput) (string, error) {
	ctx, span := c.tracer.StartSpan(ctx)
	defer span.End()

	if input == nil {
		return "", ErrNilInputProvided
	}

	logger := c.logger.WithValue(keys.HouseholdIDKey, input.DestinationHousehold)
	tracing.AttachHouseholdIDToSpan(span, input.DestinationHousehold)

	// we don't validate here because it needs to have the user ID

	req, err := c.requestBuilder.BuildInviteUserToHouseholdRequest(ctx, input)
	if err != nil {
		return "", observability.PrepareError(err, logger, span, "building add user to household request")
	}

	var pwr *types.PreWriteResponse
	if err = c.fetchAndUnmarshal(ctx, req, &pwr); err != nil {
		return "", observability.PrepareError(err, logger, span, "adding user to household")
	}

	return pwr.ID, nil
}

// MarkAsDefault marks a given household as the default for a given user.
func (c *Client) MarkAsDefault(ctx context.Context, householdID string) error {
	ctx, span := c.tracer.StartSpan(ctx)
	defer span.End()

	if householdID == "" {
		return ErrInvalidIDProvided
	}

	logger := c.logger.WithValue(keys.HouseholdIDKey, householdID)
	tracing.AttachHouseholdIDToSpan(span, householdID)

	req, err := c.requestBuilder.BuildMarkAsDefaultRequest(ctx, householdID)
	if err != nil {
		return observability.PrepareError(err, logger, span, "building mark household as default request")
	}

	if err = c.fetchAndUnmarshal(ctx, req, nil); err != nil {
		return observability.PrepareError(err, logger, span, "marking household as default")
	}

	return nil
}

// RemoveUserFromHousehold removes a user from a household.
func (c *Client) RemoveUserFromHousehold(ctx context.Context, householdID, userID string) error {
	ctx, span := c.tracer.StartSpan(ctx)
	defer span.End()

	if householdID == "" {
		return ErrInvalidIDProvided
	}

	if userID == "" {
		return ErrInvalidIDProvided
	}

	logger := c.logger.WithValue(keys.HouseholdIDKey, householdID).WithValue(keys.UserIDKey, userID)
	tracing.AttachHouseholdIDToSpan(span, householdID)
	tracing.AttachUserIDToSpan(span, userID)

	req, err := c.requestBuilder.BuildRemoveUserRequest(ctx, householdID, userID, "")
	if err != nil {
		return observability.PrepareError(err, logger, span, "building remove user from household request")
	}

	if err = c.fetchAndUnmarshal(ctx, req, nil); err != nil {
		return observability.PrepareError(err, logger, span, "removing user from household")
	}

	return nil
}

// ModifyMemberPermissions modifies a given user's permissions for a given household.
func (c *Client) ModifyMemberPermissions(ctx context.Context, householdID, userID string, input *types.ModifyUserPermissionsInput) error {
	ctx, span := c.tracer.StartSpan(ctx)
	defer span.End()

	if householdID == "" {
		return ErrInvalidIDProvided
	}

	if userID == "" {
		return ErrInvalidIDProvided
	}

	if input == nil {
		return ErrNilInputProvided
	}

	logger := c.logger.WithValue(keys.HouseholdIDKey, householdID).WithValue(keys.UserIDKey, userID)
	tracing.AttachHouseholdIDToSpan(span, householdID)
	tracing.AttachUserIDToSpan(span, userID)

	if err := input.ValidateWithContext(ctx); err != nil {
		return observability.PrepareError(err, logger, span, "validating input")
	}

	req, err := c.requestBuilder.BuildModifyMemberPermissionsRequest(ctx, householdID, userID, input)
	if err != nil {
		return observability.PrepareError(err, logger, span, "building modify household member permissions request")
	}

	if err = c.fetchAndUnmarshal(ctx, req, nil); err != nil {
		return observability.PrepareError(err, logger, span, "modifying user household permissions")
	}

	return nil
}

// TransferHouseholdOwnership transfers ownership of a household to a given user.
func (c *Client) TransferHouseholdOwnership(ctx context.Context, householdID string, input *types.HouseholdOwnershipTransferInput) error {
	ctx, span := c.tracer.StartSpan(ctx)
	defer span.End()

	if householdID == "" {
		return ErrInvalidIDProvided
	}

	if input == nil {
		return ErrNilInputProvided
	}

	logger := c.logger.WithValue(keys.HouseholdIDKey, householdID).
		WithValue("old_owner", input.CurrentOwner).
		WithValue("new_owner", input.NewOwner)

	tracing.AttachToSpan(span, "old_owner", input.CurrentOwner)
	tracing.AttachToSpan(span, "new_owner", input.NewOwner)

	if err := input.ValidateWithContext(ctx); err != nil {
		return observability.PrepareError(err, logger, span, "validating input")
	}

	req, err := c.requestBuilder.BuildTransferHouseholdOwnershipRequest(ctx, householdID, input)
	if err != nil {
		return observability.PrepareError(err, logger, span, "building transfer household ownership request")
	}

	if err = c.fetchAndUnmarshal(ctx, req, nil); err != nil {
		return observability.PrepareError(err, logger, span, "transferring household to user")
	}

	return nil
}
