package postgres

import (
	"context"
	"database/sql"
	_ "embed"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/prixfixeco/backend/internal/authorization"
	"github.com/prixfixeco/backend/internal/database"
	"github.com/prixfixeco/backend/internal/identifiers"
	"github.com/prixfixeco/backend/internal/observability"
	"github.com/prixfixeco/backend/internal/observability/keys"
	"github.com/prixfixeco/backend/internal/observability/tracing"
	"github.com/prixfixeco/backend/pkg/types"
)

const (
	householdOnHouseholdInvitationsJoin = "households ON household_invitations.destination_household = households.id"
	usersOnHouseholdInvitationsJoin     = "users ON household_invitations.from_user = users.id"
)

var (
	_ types.HouseholdInvitationDataManager = (*Querier)(nil)

	// householdInvitationsTableColumns are the columns for the household invitations table.
	householdInvitationsTableColumns = []string{
		"household_invitations.id",
		"households.id",
		"households.name",
		"households.billing_status",
		"households.contact_email",
		"households.contact_phone",
		"households.payment_processor_customer_id",
		"households.subscription_plan_id",
		"households.time_zone",
		"households.created_at",
		"households.last_updated_at",
		"households.archived_at",
		"households.belongs_to_user",
		"household_invitations.to_email",
		"household_invitations.to_user",
		"users.id",
		"users.username",
		"users.email_address",
		"users.avatar_src",
		"users.hashed_password",
		"users.requires_password_change",
		"users.password_last_changed_at",
		"users.two_factor_secret",
		"users.two_factor_secret_verified_at",
		"users.service_role",
		"users.user_account_status",
		"users.user_account_status_explanation",
		"users.birthday",
		"users.created_at",
		"users.last_updated_at",
		"users.archived_at",
		"household_invitations.status",
		"household_invitations.note",
		"household_invitations.status_note",
		"household_invitations.token",
		"household_invitations.expires_at",
		"household_invitations.created_at",
		"household_invitations.last_updated_at",
		"household_invitations.archived_at",
	}
)

// scanHouseholdInvitation is a consistent way to turn a *sql.Row into an invitation struct.
func (q *Querier) scanHouseholdInvitation(ctx context.Context, scan database.Scanner, includeCounts bool) (householdInvitation *types.HouseholdInvitation, filteredCount, totalCount uint64, err error) {
	_, span := q.tracer.StartSpan(ctx)
	defer span.End()

	householdInvitation = &types.HouseholdInvitation{
		DestinationHousehold: types.Household{},
	}

	targetVars := []any{
		&householdInvitation.ID,
		&householdInvitation.DestinationHousehold.ID,
		&householdInvitation.DestinationHousehold.Name,
		&householdInvitation.DestinationHousehold.BillingStatus,
		&householdInvitation.DestinationHousehold.ContactEmail,
		&householdInvitation.DestinationHousehold.ContactPhone,
		&householdInvitation.DestinationHousehold.PaymentProcessorCustomerID,
		&householdInvitation.DestinationHousehold.SubscriptionPlanID,
		&householdInvitation.DestinationHousehold.TimeZone,
		&householdInvitation.DestinationHousehold.CreatedAt,
		&householdInvitation.DestinationHousehold.LastUpdatedAt,
		&householdInvitation.DestinationHousehold.ArchivedAt,
		&householdInvitation.DestinationHousehold.BelongsToUser,
		&householdInvitation.ToEmail,
		&householdInvitation.ToUser,
		&householdInvitation.FromUser.ID,
		&householdInvitation.FromUser.Username,
		&householdInvitation.FromUser.EmailAddress,
		&householdInvitation.FromUser.AvatarSrc,
		&householdInvitation.FromUser.HashedPassword,
		&householdInvitation.FromUser.RequiresPasswordChange,
		&householdInvitation.FromUser.PasswordLastChangedAt,
		&householdInvitation.FromUser.TwoFactorSecret,
		&householdInvitation.FromUser.TwoFactorSecretVerifiedAt,
		&householdInvitation.FromUser.ServiceRole,
		&householdInvitation.FromUser.AccountStatus,
		&householdInvitation.FromUser.AccountStatusExplanation,
		&householdInvitation.FromUser.Birthday,
		&householdInvitation.FromUser.CreatedAt,
		&householdInvitation.FromUser.LastUpdatedAt,
		&householdInvitation.FromUser.ArchivedAt,
		&householdInvitation.Status,
		&householdInvitation.Note,
		&householdInvitation.StatusNote,
		&householdInvitation.Token,
		&householdInvitation.ExpiresAt,
		&householdInvitation.CreatedAt,
		&householdInvitation.LastUpdatedAt,
		&householdInvitation.ArchivedAt,
	}

	if includeCounts {
		targetVars = append(targetVars, &filteredCount, &totalCount)
	}

	if err = scan.Scan(targetVars...); err != nil {
		return nil, 0, 0, observability.PrepareError(err, span, "scanning household invitation")
	}

	return householdInvitation, filteredCount, totalCount, nil
}

// scanHouseholdInvitations provides a consistent way to turn sql rows into a slice of household_invitations.
func (q *Querier) scanHouseholdInvitations(ctx context.Context, rows database.ResultIterator, includeCounts bool) (householdInvitations []*types.HouseholdInvitation, filteredCount, totalCount uint64, err error) {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	for rows.Next() {
		householdInvitation, fc, tc, scanErr := q.scanHouseholdInvitation(ctx, rows, includeCounts)
		if scanErr != nil {
			return nil, 0, 0, scanErr
		}

		if includeCounts {
			if filteredCount == 0 {
				filteredCount = fc
			}

			if totalCount == 0 {
				totalCount = tc
			}
		}

		householdInvitations = append(householdInvitations, householdInvitation)
	}

	if err = rows.Err(); err != nil {
		return nil, 0, 0, observability.PrepareError(err, span, "fetching household invitation from database")
	}

	if err = rows.Close(); err != nil {
		return nil, 0, 0, observability.PrepareError(err, span, "fetching household invitation from database")
	}

	return householdInvitations, filteredCount, totalCount, nil
}

//go:embed queries/household_invitations/exists.sql
var householdInvitationExistenceQuery string

// HouseholdInvitationExists fetches whether a household invitation exists from the database.
func (q *Querier) HouseholdInvitationExists(ctx context.Context, householdInvitationID string) (bool, error) {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	logger := q.logger.Clone()

	if householdInvitationID == "" {
		return false, ErrInvalidIDProvided
	}
	logger = logger.WithValue(keys.HouseholdInvitationIDKey, householdInvitationID)
	tracing.AttachHouseholdInvitationIDToSpan(span, householdInvitationID)

	args := []any{
		householdInvitationID,
	}

	result, err := q.performBooleanQuery(ctx, q.db, householdInvitationExistenceQuery, args)
	if err != nil {
		return false, observability.PrepareAndLogError(err, logger, span, "performing household invitation existence check")
	}

	return result, nil
}

//go:embed queries/household_invitations/get_by_household_and_id.sql
var getHouseholdInvitationByHouseholdAndIDQuery string

// GetHouseholdInvitationByHouseholdAndID fetches an invitation from the database.
func (q *Querier) GetHouseholdInvitationByHouseholdAndID(ctx context.Context, householdID, householdInvitationID string) (*types.HouseholdInvitation, error) {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	logger := q.logger.Clone()

	if householdID == "" {
		return nil, ErrInvalidIDProvided
	}
	logger = logger.WithValue(keys.HouseholdIDKey, householdID)
	tracing.AttachHouseholdIDToSpan(span, householdID)

	if householdInvitationID == "" {
		return nil, ErrInvalidIDProvided
	}
	logger = logger.WithValue(keys.HouseholdInvitationIDKey, householdInvitationID)
	tracing.AttachHouseholdInvitationIDToSpan(span, householdInvitationID)

	args := []any{
		householdID,
		householdInvitationID,
	}

	row := q.getOneRow(ctx, q.db, "household invitation", getHouseholdInvitationByHouseholdAndIDQuery, args)

	householdInvitation, _, _, err := q.scanHouseholdInvitation(ctx, row, false)
	if err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "scanning household invitation")
	}

	return householdInvitation, nil
}

/* #nosec G101 */
//go:embed queries/household_invitations/get_by_token_and_id.sql
var getHouseholdInvitationByTokenAndIDQuery string

// GetHouseholdInvitationByTokenAndID fetches an invitation from the database.
func (q *Querier) GetHouseholdInvitationByTokenAndID(ctx context.Context, token, invitationID string) (*types.HouseholdInvitation, error) {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	logger := q.logger.Clone()

	if token == "" {
		return nil, ErrInvalidIDProvided
	}

	if invitationID == "" {
		return nil, ErrInvalidIDProvided
	}
	logger = logger.WithValue(keys.HouseholdInvitationIDKey, invitationID)
	tracing.AttachHouseholdInvitationIDToSpan(span, invitationID)

	logger.Debug("fetching household invitation")

	args := []any{
		token,
		invitationID,
	}

	row := q.getOneRow(ctx, q.db, "household invitation", getHouseholdInvitationByTokenAndIDQuery, args)

	householdInvitation, _, _, err := q.scanHouseholdInvitation(ctx, row, false)
	if err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "scanning household invitation")
	}

	return householdInvitation, nil
}

/* #nosec G101 */
//go:embed queries/household_invitations/get_by_email_and_token.sql
var getHouseholdInvitationByEmailAndTokenQuery string

// GetHouseholdInvitationByEmailAndToken fetches an invitation from the database.
func (q *Querier) GetHouseholdInvitationByEmailAndToken(ctx context.Context, emailAddress, token string) (*types.HouseholdInvitation, error) {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	logger := q.logger.Clone()

	if emailAddress == "" {
		return nil, ErrInvalidIDProvided
	}
	logger = logger.WithValue(keys.UserEmailAddressKey, emailAddress)
	tracing.AttachEmailAddressToSpan(span, emailAddress)

	if token == "" {
		return nil, ErrInvalidIDProvided
	}
	logger = logger.WithValue(keys.HouseholdInvitationTokenKey, token)
	tracing.AttachHouseholdInvitationTokenToSpan(span, token)

	args := []any{
		emailAddress,
		token,
	}

	row := q.getOneRow(ctx, q.db, "household invitation", getHouseholdInvitationByEmailAndTokenQuery, args)

	invitation, _, _, err := q.scanHouseholdInvitation(ctx, row, false)
	if err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "scanning invitation")
	}

	return invitation, nil
}

//go:embed queries/household_invitations/create.sql
var createHouseholdInvitationQuery string

// CreateHouseholdInvitation creates an invitation in a database.
func (q *Querier) CreateHouseholdInvitation(ctx context.Context, input *types.HouseholdInvitationDatabaseCreationInput) (*types.HouseholdInvitation, error) {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	if input == nil {
		return nil, ErrNilInputProvided
	}

	logger := q.logger.WithValue(keys.HouseholdInvitationIDKey, input.ID)

	args := []any{
		input.ID,
		input.FromUser,
		input.ToUser,
		input.Note,
		input.ToEmail,
		input.Token,
		input.DestinationHouseholdID,
		input.ExpiresAt,
	}

	if err := q.performWriteQuery(ctx, q.db, "household invitation creation", createHouseholdInvitationQuery, args); err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "performing household invitation creation query")
	}

	x := &types.HouseholdInvitation{
		ID:                   input.ID,
		FromUser:             types.User{ID: input.FromUser},
		ToUser:               input.ToUser,
		Note:                 input.Note,
		ToEmail:              input.ToEmail,
		Token:                input.Token,
		StatusNote:           "",
		Status:               types.PendingHouseholdInvitationStatus,
		DestinationHousehold: types.Household{ID: input.DestinationHouseholdID},
		ExpiresAt:            input.ExpiresAt,
		CreatedAt:            q.currentTime(),
	}

	tracing.AttachHouseholdInvitationIDToSpan(span, x.ID)
	logger = logger.WithValue(keys.HouseholdInvitationIDKey, x.ID)

	logger.Info("household invitation created")

	return x, nil
}

// BuildGetPendingHouseholdInvitationsFromUserQuery builds a query for fetching pending household invitations sent by a given user.
func (q *Querier) BuildGetPendingHouseholdInvitationsFromUserQuery(ctx context.Context, userID string, filter *types.QueryFilter) (query string, args []any) {
	_, span := q.tracer.StartSpan(ctx)
	defer span.End()

	where := squirrel.Eq{
		"household_invitations.from_user":   userID,
		"household_invitations.archived_at": nil,
		"household_invitations.status":      types.PendingHouseholdInvitationStatus,
	}

	joins := []string{householdOnHouseholdInvitationsJoin, usersOnHouseholdInvitationsJoin}

	filteredCountQuery, filteredCountQueryArgs := q.buildFilteredCountQuery(ctx, "household_invitations", joins, where, "", "", false, false, filter)
	totalCountQuery, totalCountQueryArgs := q.buildTotalCountQuery(ctx, "household_invitations", joins, where, "", "", false, false)

	queryBuilder := q.sqlBuilder.Select(
		append(
			householdInvitationsTableColumns,
			fmt.Sprintf("(%s) as filtered_count", filteredCountQuery),
			fmt.Sprintf("(%s) as total_count", totalCountQuery),
		)...,
	).
		From("household_invitations").
		Join(householdOnHouseholdInvitationsJoin).
		Join(usersOnHouseholdInvitationsJoin).
		Where(where)

	queryBuilder = applyFilterToQueryBuilder(filter, "household_invitations", queryBuilder)

	query, args, err := queryBuilder.ToSql()
	q.logQueryBuildingError(span, err)

	return query, append(append(filteredCountQueryArgs, totalCountQueryArgs...), args...)
}

// GetPendingHouseholdInvitationsFromUser fetches pending household invitations sent from a given user.
func (q *Querier) GetPendingHouseholdInvitationsFromUser(ctx context.Context, userID string, filter *types.QueryFilter) (*types.HouseholdInvitationList, error) {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	logger := q.logger.WithValue(keys.UserIDKey, userID)
	filter.AttachToLogger(logger)

	query, args := q.BuildGetPendingHouseholdInvitationsFromUserQuery(ctx, userID, filter)

	rows, err := q.getRows(ctx, q.db, "household invitations from user", query, args)
	if err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "reading household invitations from user")
	}

	householdInvitations, fc, tc, err := q.scanHouseholdInvitations(ctx, rows, true)
	if err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "reading household invitations from user")
	}

	returnList := &types.HouseholdInvitationList{
		Pagination: types.Pagination{
			FilteredCount: fc,
			TotalCount:    tc,
		},
		HouseholdInvitations: householdInvitations,
	}

	if filter != nil {
		if filter.Page != nil {
			returnList.Page = *filter.Page
		}

		if filter.Limit != nil {
			returnList.Limit = *filter.Limit
		}
	}

	return returnList, nil
}

// BuildGetPendingHouseholdInvitationsForUserQuery builds a query for fetching pending household invitations sent to a given user.
func (q *Querier) BuildGetPendingHouseholdInvitationsForUserQuery(ctx context.Context, userID string, filter *types.QueryFilter) (query string, args []any) {
	_, span := q.tracer.StartSpan(ctx)
	defer span.End()

	where := squirrel.Eq{
		"household_invitations.to_user":     userID,
		"household_invitations.archived_at": nil,
		"household_invitations.status":      types.PendingHouseholdInvitationStatus,
	}

	joins := []string{householdOnHouseholdInvitationsJoin, usersOnHouseholdInvitationsJoin}

	filteredCountQuery, filteredCountQueryArgs := q.buildFilteredCountQuery(ctx, "household_invitations", joins, where, "", "", false, false, filter)
	totalCountQuery, totalCountQueryArgs := q.buildTotalCountQuery(ctx, "household_invitations", joins, where, "", "", false, false)

	queryBuilder := q.sqlBuilder.Select(
		append(
			householdInvitationsTableColumns,
			fmt.Sprintf("(%s) as filtered_count", filteredCountQuery),
			fmt.Sprintf("(%s) as total_count", totalCountQuery),
		)...,
	).
		From("household_invitations").
		Join(householdOnHouseholdInvitationsJoin).
		Join(usersOnHouseholdInvitationsJoin).
		Where(where)

	queryBuilder = applyFilterToQueryBuilder(filter, "household_invitations", queryBuilder)

	query, args, err := queryBuilder.ToSql()
	q.logQueryBuildingError(span, err)

	return query, append(append(filteredCountQueryArgs, totalCountQueryArgs...), args...)
}

// GetPendingHouseholdInvitationsForUser fetches pending household invitations sent to a given user.
func (q *Querier) GetPendingHouseholdInvitationsForUser(ctx context.Context, userID string, filter *types.QueryFilter) (*types.HouseholdInvitationList, error) {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	logger := q.logger.WithValue(keys.UserIDKey, userID)
	filter.AttachToLogger(logger)

	query, args := q.BuildGetPendingHouseholdInvitationsForUserQuery(ctx, userID, filter)

	rows, err := q.getRows(ctx, q.db, "household invitations from user", query, args)
	if err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "reading household invitations from user")
	}

	householdInvitations, fc, tc, err := q.scanHouseholdInvitations(ctx, rows, true)
	if err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "reading household invitations from user")
	}

	returnList := &types.HouseholdInvitationList{
		Pagination: types.Pagination{
			FilteredCount: fc,
			TotalCount:    tc,
		},
		HouseholdInvitations: householdInvitations,
	}

	if filter != nil {
		if filter.Page != nil {
			returnList.Page = *filter.Page
		}

		if filter.Limit != nil {
			returnList.Limit = *filter.Limit
		}
	}

	return returnList, nil
}

//go:embed queries/household_invitations/set_status.sql
var setInvitationStatusQuery string

func (q *Querier) setInvitationStatus(ctx context.Context, querier database.SQLQueryExecutor, householdInvitationID, note string, status types.HouseholdInvitationStatus) error {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	logger := q.logger.WithValue("new_status", status)

	if householdInvitationID == "" {
		return ErrInvalidIDProvided
	}
	logger = logger.WithValue(keys.HouseholdInvitationIDKey, householdInvitationID)
	tracing.AttachHouseholdInvitationIDToSpan(span, householdInvitationID)

	args := []any{
		status,
		note,
		householdInvitationID,
	}

	if err := q.performWriteQuery(ctx, querier, "household invitation status change", setInvitationStatusQuery, args); err != nil {
		return observability.PrepareAndLogError(err, logger, span, "changing household invitation status")
	}

	logger.Debug("household invitation updated")

	return nil
}

// CancelHouseholdInvitation cancels a household invitation by its ID with a note.
func (q *Querier) CancelHouseholdInvitation(ctx context.Context, householdInvitationID, note string) error {
	return q.setInvitationStatus(ctx, q.db, householdInvitationID, note, types.CancelledHouseholdInvitationStatus)
}

// AcceptHouseholdInvitation accepts a household invitation by its ID with a note.
func (q *Querier) AcceptHouseholdInvitation(ctx context.Context, householdInvitationID, token, note string) error {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	logger := q.logger.Clone()

	if householdInvitationID == "" {
		return ErrInvalidIDProvided
	}
	logger = logger.WithValue(keys.HouseholdInvitationIDKey, householdInvitationID)
	tracing.AttachHouseholdInvitationIDToSpan(span, householdInvitationID)

	tx, err := q.db.BeginTx(ctx, nil)
	if err != nil {
		return observability.PrepareAndLogError(err, logger, span, "beginning transaction")
	}

	if err = q.setInvitationStatus(ctx, tx, householdInvitationID, note, types.AcceptedHouseholdInvitationStatus); err != nil {
		q.rollbackTransaction(ctx, tx)
		return observability.PrepareAndLogError(err, logger, span, "accepting household invitation")
	}

	invitation, err := q.GetHouseholdInvitationByTokenAndID(ctx, token, householdInvitationID)
	if err != nil {
		q.rollbackTransaction(ctx, tx)
		return observability.PrepareAndLogError(err, logger, span, "fetching household invitation")
	}

	addUserInput := &types.HouseholdUserMembershipDatabaseCreationInput{
		ID:            identifiers.New(),
		Reason:        fmt.Sprintf("accepted household invitation %q", householdInvitationID),
		HouseholdID:   invitation.DestinationHousehold.ID,
		HouseholdRole: "household_member",
	}
	if invitation.ToUser != nil {
		addUserInput.UserID = *invitation.ToUser
	}

	if err = q.addUserToHousehold(ctx, tx, addUserInput); err != nil {
		q.rollbackTransaction(ctx, tx)
		return observability.PrepareAndLogError(err, logger, span, "adding user to household")
	}

	if err = tx.Commit(); err != nil {
		return observability.PrepareAndLogError(err, logger, span, "committing transaction")
	}

	return nil
}

// RejectHouseholdInvitation rejects a household invitation by its ID with a note.
func (q *Querier) RejectHouseholdInvitation(ctx context.Context, householdInvitationID, note string) error {
	return q.setInvitationStatus(ctx, q.db, householdInvitationID, note, types.RejectedHouseholdInvitationStatus)
}

//go:embed queries/household_invitations/attach_invitations_to_user_id.sql
var attachInvitationsToUserIDQuery string

func (q *Querier) attachInvitationsToUser(ctx context.Context, querier database.SQLQueryExecutor, userEmail, userID string) error {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	logger := q.logger

	if userEmail == "" {
		return ErrInvalidIDProvided
	}
	logger = logger.WithValue(keys.UserEmailAddressKey, userEmail)
	tracing.AttachHouseholdIDToSpan(span, userEmail)

	if userID == "" {
		return ErrInvalidIDProvided
	}
	logger = logger.WithValue(keys.UserIDKey, userID)
	tracing.AttachHouseholdInvitationIDToSpan(span, userID)

	args := []any{userID, userEmail}

	if err := q.performWriteQuery(ctx, querier, "invitation attachment", attachInvitationsToUserIDQuery, args); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return observability.PrepareAndLogError(err, logger, span, "attaching invitations to user")
	}

	logger.Info("invitations associated with user")

	return nil
}

func (q *Querier) acceptInvitationForUser(ctx context.Context, querier database.SQLQueryExecutorAndTransactionManager, input *types.UserDatabaseCreationInput) error {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	logger := q.logger.WithValue(keys.UsernameKey, input.Username).WithValue(keys.UserEmailAddressKey, input.EmailAddress)

	invitation, tokenCheckErr := q.GetHouseholdInvitationByEmailAndToken(ctx, input.EmailAddress, input.InvitationToken)
	if tokenCheckErr != nil {
		q.rollbackTransaction(ctx, querier)
		return observability.PrepareError(tokenCheckErr, span, "fetching household invitation")
	}

	logger.Debug("fetched invitation to accept for user")

	createHouseholdMembershipForNewUserArgs := []any{
		identifiers.New(),
		input.ID,
		input.DestinationHouseholdID,
		true,
		authorization.HouseholdMemberRole.String(),
	}

	if err := q.performWriteQuery(ctx, querier, "household user membership creation", createHouseholdMembershipForNewUserQuery, createHouseholdMembershipForNewUserArgs); err != nil {
		q.rollbackTransaction(ctx, querier)
		return observability.PrepareAndLogError(err, logger, span, "writing destination household membership")
	}

	logger.Debug("created membership via invitation")

	if err := q.setInvitationStatus(ctx, querier, invitation.ID, "", types.AcceptedHouseholdInvitationStatus); err != nil {
		q.rollbackTransaction(ctx, querier)
		return observability.PrepareAndLogError(err, logger, span, "accepting household invitation")
	}

	logger.Debug("marked invitation as accepted")

	return nil
}
