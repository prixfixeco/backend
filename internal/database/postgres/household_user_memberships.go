package postgres

import (
	"context"
	"database/sql"
	_ "embed"
	"errors"

	"github.com/dinnerdonebetter/backend/internal/authorization"
	"github.com/dinnerdonebetter/backend/internal/database"
	"github.com/dinnerdonebetter/backend/internal/database/postgres/generated"
	"github.com/dinnerdonebetter/backend/internal/identifiers"
	"github.com/dinnerdonebetter/backend/internal/observability"
	"github.com/dinnerdonebetter/backend/internal/observability/keys"
	"github.com/dinnerdonebetter/backend/internal/observability/tracing"
	"github.com/dinnerdonebetter/backend/pkg/types"
)

const (
	// householdsUserMembershipTableName is what the households membership table calls itself.
	householdsUserMembershipTableName = "household_user_memberships"
)

var (
	_ types.HouseholdUserMembershipDataManager = (*Querier)(nil)

	// householdsUserMembershipTableColumns are the columns for the household user memberships table.
	householdsUserMembershipTableColumns = []string{
		"household_user_memberships.id",
		"household_user_memberships.belongs_to_user",
		"household_user_memberships.belongs_to_household",
		"household_user_memberships.household_role",
		"household_user_memberships.default_household",
		"household_user_memberships.created_at",
		"household_user_memberships.last_updated_at",
		"household_user_memberships.archived_at",
	}
)

// BuildSessionContextDataForUser queries the database for the memberships of a user and constructs a SessionContextData struct from the results.
func (q *Querier) BuildSessionContextDataForUser(ctx context.Context, userID string) (*types.SessionContextData, error) {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	if userID == "" {
		return nil, ErrInvalidIDProvided
	}

	logger := q.logger.WithValue(keys.UserIDKey, userID)
	tracing.AttachUserIDToSpan(span, userID)

	results, err := q.generatedQuerier.GetHouseholdUserMembershipsForUser(ctx, q.db, userID)
	if err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "fetching user's memberships from database")
	}

	householdRolesMap := map[string]authorization.HouseholdRolePermissionsChecker{}
	defaultHouseholdID := ""
	for _, result := range results {
		householdRolesMap[result.BelongsToHousehold] = authorization.NewHouseholdRolePermissionChecker(result.HouseholdRole)
		if result.DefaultHousehold {
			defaultHouseholdID = result.BelongsToHousehold
		}
	}

	user, err := q.GetUser(ctx, userID)
	if err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "fetching user from database")
	}

	sessionCtxData := &types.SessionContextData{
		Requester: types.RequesterInfo{
			UserID:                   user.ID,
			Username:                 user.Username,
			EmailAddress:             user.EmailAddress,
			AccountStatus:            user.AccountStatus,
			AccountStatusExplanation: user.AccountStatusExplanation,
			ServicePermissions:       authorization.NewServiceRolePermissionChecker(user.ServiceRole),
		},
		HouseholdPermissions: householdRolesMap,
		ActiveHouseholdID:    defaultHouseholdID,
	}

	logger.Debug("fetched session context data for user")

	return sessionCtxData, nil
}

// GetDefaultHouseholdIDForUser retrieves the default household ID for a user.
func (q *Querier) GetDefaultHouseholdIDForUser(ctx context.Context, userID string) (string, error) {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	logger := q.logger.Clone()

	if userID == "" {
		return "", ErrInvalidIDProvided
	}
	tracing.AttachUserIDToSpan(span, userID)
	logger = logger.WithValue(keys.UserIDKey, userID)

	id, err := q.generatedQuerier.GetDefaultHouseholdIDForUser(ctx, q.db, userID)
	if err != nil {
		return "", observability.PrepareAndLogError(err, logger, span, "fetching default household ID for user")
	}

	return id, nil
}

//go:embed queries/household_user_memberships/mark_as_user_default.sql
var markHouseholdAsUserDefaultQuery string

// markHouseholdAsUserDefault does a thing.
func (q *Querier) markHouseholdAsUserDefault(ctx context.Context, querier database.SQLQueryExecutor, userID, householdID string) error {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	if userID == "" || householdID == "" {
		return ErrInvalidIDProvided
	}

	logger := q.logger.WithValues(map[string]any{
		keys.UserIDKey:      userID,
		keys.HouseholdIDKey: householdID,
	})

	tracing.AttachUserIDToSpan(span, userID)
	tracing.AttachHouseholdIDToSpan(span, householdID)

	args := []any{
		userID,
		householdID,
		userID,
	}

	// create the household.
	if err := q.performWriteQuery(ctx, querier, "user default household assignment", markHouseholdAsUserDefaultQuery, args); err != nil {
		return observability.PrepareAndLogError(err, logger, span, "assigning user default household")
	}

	logger.Debug("household marked as default")

	return nil
}

// MarkHouseholdAsUserDefault does a thing.
func (q *Querier) MarkHouseholdAsUserDefault(ctx context.Context, userID, householdID string) error {
	return q.markHouseholdAsUserDefault(ctx, q.db, userID, householdID)
}

// UserIsMemberOfHousehold does a thing.
func (q *Querier) UserIsMemberOfHousehold(ctx context.Context, userID, householdID string) (bool, error) {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	if userID == "" || householdID == "" {
		return false, ErrInvalidIDProvided
	}
	tracing.AttachUserIDToSpan(span, userID)
	tracing.AttachHouseholdIDToSpan(span, householdID)

	result, err := q.generatedQuerier.UserIsHouseholdMember(ctx, q.db, &generated.UserIsHouseholdMemberParams{
		BelongsToHousehold: householdID,
		BelongsToUser:      userID,
	})
	if err != nil {
		return false, observability.PrepareError(err, span, "performing user membership check query")
	}

	return result, nil
}

//go:embed queries/household_user_memberships/modify_user_permissions.sql
var modifyUserPermissionsQuery string

// ModifyUserPermissions does a thing.
func (q *Querier) ModifyUserPermissions(ctx context.Context, householdID, userID string, input *types.ModifyUserPermissionsInput) error {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	if householdID == "" || userID == "" {
		return ErrInvalidIDProvided
	}

	if input == nil {
		return ErrNilInputProvided
	}

	logger := q.logger.WithValues(map[string]any{
		keys.HouseholdIDKey: householdID,
		keys.UserIDKey:      userID,
		"new_roles":         input.NewRole,
	})

	tracing.AttachUserIDToSpan(span, userID)
	tracing.AttachHouseholdIDToSpan(span, householdID)

	args := []any{
		input.NewRole,
		householdID,
		userID,
	}

	// modify the membership.
	if err := q.performWriteQuery(ctx, q.db, "user household permissions modification", modifyUserPermissionsQuery, args); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return observability.PrepareAndLogError(err, logger, span, "modifying user household permissions")
	}

	logger.Info("user permissions modified")

	return nil
}

//go:embed queries/household_user_memberships/transfer_ownership.sql
var transferHouseholdOwnershipQuery string

// TransferHouseholdOwnership does a thing.
func (q *Querier) TransferHouseholdOwnership(ctx context.Context, householdID string, input *types.HouseholdOwnershipTransferInput) error {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	if householdID == "" {
		return ErrInvalidIDProvided
	}

	if input == nil {
		return ErrNilInputProvided
	}

	logger := q.logger.WithValues(map[string]any{
		keys.HouseholdIDKey: householdID,
		"current_owner":     input.CurrentOwner,
		"new_owner":         input.NewOwner,
	})

	tracing.AttachUserIDToSpan(span, input.NewOwner)
	tracing.AttachHouseholdIDToSpan(span, householdID)

	// begin household transfer transaction
	tx, err := q.db.BeginTx(ctx, nil)
	if err != nil {
		return observability.PrepareAndLogError(err, logger, span, "beginning transaction")
	}

	transferHouseholdOwnershipArgs := []any{
		input.NewOwner,
		input.CurrentOwner,
		householdID,
	}

	// create the membership.
	if err = q.performWriteQuery(ctx, tx, "user ownership transfer", transferHouseholdOwnershipQuery, transferHouseholdOwnershipArgs); err != nil {
		q.rollbackTransaction(ctx, tx)
		return observability.PrepareAndLogError(err, logger, span, "transferring household to new owner")
	}

	isMember, err := q.UserIsMemberOfHousehold(ctx, input.NewOwner, householdID)
	if err != nil {
		q.rollbackTransaction(ctx, tx)
		return observability.PrepareAndLogError(err, logger, span, "checking if user is member of household")
	}

	if !isMember {
		if err = q.addUserToHousehold(ctx, tx, &types.HouseholdUserMembershipDatabaseCreationInput{
			ID:            identifiers.New(),
			Reason:        "transferred ownership",
			UserID:        input.NewOwner,
			HouseholdID:   householdID,
			HouseholdRole: "household_admin",
		}); err != nil {
			q.rollbackTransaction(ctx, tx)
			return observability.PrepareAndLogError(err, logger, span, "adding user to household")
		}
	}

	if err = q.removeUserFromHousehold(ctx, tx, input.CurrentOwner, householdID); err != nil {
		return observability.PrepareAndLogError(err, logger, span, "removing user from household")
	}

	if err = tx.Commit(); err != nil {
		return observability.PrepareAndLogError(err, logger, span, "committing transaction")
	}

	logger.Info("household transferred to new owner")

	return nil
}

//go:embed queries/household_user_memberships/add_user_to_household.sql
var addUserToHouseholdQuery string

// addUserToHousehold does a thing.
func (q *Querier) addUserToHousehold(ctx context.Context, querier database.SQLQueryExecutor, input *types.HouseholdUserMembershipDatabaseCreationInput) error {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	if input == nil {
		return ErrNilInputProvided
	}

	logger := q.logger.WithValues(map[string]any{
		keys.UserIDKey:      input.UserID,
		keys.HouseholdIDKey: input.HouseholdID,
	})

	tracing.AttachUserIDToSpan(span, input.UserID)
	tracing.AttachHouseholdIDToSpan(span, input.HouseholdID)

	addUserToHouseholdArgs := []any{
		input.ID,
		input.UserID,
		input.HouseholdID,
		input.HouseholdRole,
	}

	// create the membership.
	if err := q.performWriteQuery(ctx, querier, "user household membership creation", addUserToHouseholdQuery, addUserToHouseholdArgs); err != nil {
		return observability.PrepareAndLogError(err, logger, span, "performing user household membership creation query")
	}

	logger.Info("user added to household")

	return nil
}

//go:embed queries/household_user_memberships/remove_user_from_household.sql
var removeUserFromHouseholdQuery string

// removeUserFromHousehold removes a user's membership to a household.
func (q *Querier) removeUserFromHousehold(ctx context.Context, querier database.SQLQueryExecutorAndTransactionManager, userID, householdID string) error {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	if userID == "" {
		return ErrInvalidIDProvided
	}
	tracing.AttachUserIDToSpan(span, userID)

	if householdID == "" {
		return ErrInvalidIDProvided
	}
	tracing.AttachHouseholdIDToSpan(span, householdID)

	logger := q.logger.WithValues(map[string]any{
		keys.UserIDKey:      userID,
		keys.HouseholdIDKey: householdID,
	})

	logger.Info("creating")

	args := []any{
		householdID,
		userID,
	}

	// remove the membership.
	if err := q.performWriteQuery(ctx, querier, "user membership removal", removeUserFromHouseholdQuery, args); err != nil {
		q.rollbackTransaction(ctx, querier)
		return observability.PrepareAndLogError(err, logger, span, "removing user from household")
	}

	logger.Info("user removed from household")

	remainingHouseholds, fetchRemainingHouseholdsErr := q.getHouseholdsForUser(ctx, querier, userID, false, nil)
	if fetchRemainingHouseholdsErr != nil {
		q.rollbackTransaction(ctx, querier)
		return observability.PrepareError(fetchRemainingHouseholdsErr, span, "fetching remaining households")
	}

	logger = logger.WithValue("count", len(remainingHouseholds.Data))
	logger.Info("remaining households fetched")

	if len(remainingHouseholds.Data) == 0 {
		if err := q.createHouseholdForUser(ctx, querier, false, "", userID); err != nil {
			return observability.PrepareAndLogError(err, logger, span, "creating household for new user")
		}
		return nil
	}

	household := remainingHouseholds.Data[0]

	logger = logger.WithValue(keys.HouseholdIDKey, household.ID)
	logger.Info("about to mark household as default")

	if err := q.markHouseholdAsUserDefault(ctx, querier, userID, household.ID); err != nil {
		q.rollbackTransaction(ctx, querier)
		return observability.PrepareAndLogError(err, logger, span, "marking household as default")
	}

	logger.Info("marked household as default, committing transaction")

	return nil
}

// RemoveUserFromHousehold removes a user's membership to a household.
func (q *Querier) RemoveUserFromHousehold(ctx context.Context, userID, householdID string) error {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	if userID == "" {
		return ErrInvalidIDProvided
	}
	tracing.AttachUserIDToSpan(span, userID)

	if householdID == "" {
		return ErrInvalidIDProvided
	}
	tracing.AttachHouseholdIDToSpan(span, householdID)

	logger := q.logger.WithValues(map[string]any{
		keys.UserIDKey:      userID,
		keys.HouseholdIDKey: householdID,
	})

	logger.Info("creating")

	tx, createTransactionErr := q.db.BeginTx(ctx, nil)
	if createTransactionErr != nil {
		return observability.PrepareError(createTransactionErr, span, "beginning transaction")
	}

	if err := q.removeUserFromHousehold(ctx, tx, userID, householdID); err != nil {
		return observability.PrepareAndLogError(err, logger, span, "removing user from household")
	}

	if err := tx.Commit(); err != nil {
		return observability.PrepareAndLogError(err, logger, span, "committing transaction")
	}

	return nil
}
