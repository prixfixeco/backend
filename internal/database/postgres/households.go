package postgres

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"

	"github.com/dinnerdonebetter/backend/internal/authorization"
	"github.com/dinnerdonebetter/backend/internal/database"
	"github.com/dinnerdonebetter/backend/internal/database/postgres/generated"
	"github.com/dinnerdonebetter/backend/internal/identifiers"
	"github.com/dinnerdonebetter/backend/internal/observability"
	"github.com/dinnerdonebetter/backend/internal/observability/keys"
	"github.com/dinnerdonebetter/backend/internal/observability/tracing"
	"github.com/dinnerdonebetter/backend/pkg/types"

	"github.com/Masterminds/squirrel"
)

const (
	// householdsTableName is what the households table calls itself.
	householdsTableName = "households"

	householdUserMembershipsOnHouseholdsJoinClause = "household_user_memberships ON household_user_memberships.belongs_to_household = households.id"
	usersOnHouseholdUserMembershipsJoinClause      = "users ON household_user_memberships.belongs_to_user = users.id"
)

var (
	_ types.HouseholdDataManager = (*Querier)(nil)

	householdsTableColumns = []string{
		"households.id",
		"households.name",
		"households.billing_status",
		"households.contact_phone",
		"households.address_line_1",
		"households.address_line_2",
		"households.city",
		"households.state",
		"households.zip_code",
		"households.country",
		"households.latitude",
		"households.longitude",
		"households.payment_processor_customer_id",
		"households.subscription_plan_id",
		"households.created_at",
		"households.last_updated_at",
		"households.archived_at",
		"households.belongs_to_user",
		"users.id",
		"users.first_name",
		"users.last_name",
		"users.username",
		"users.email_address",
		"users.email_address_verified_at",
		"users.avatar_src",
		"users.requires_password_change",
		"users.password_last_changed_at",
		"users.two_factor_secret_verified_at",
		"users.service_role",
		"users.user_account_status",
		"users.user_account_status_explanation",
		"users.birthday",
		"users.created_at",
		"users.last_updated_at",
		"users.archived_at",
	}
)

// scanHousehold takes a database Scanner (i.e. *sql.Row) and scans the result into a Household struct.
func (q *Querier) scanHousehold(ctx context.Context, scan database.Scanner, includeCounts bool) (household *types.Household, membership *types.HouseholdUserMembershipWithUser, filteredCount, totalCount uint64, err error) {
	_, span := q.tracer.StartSpan(ctx)
	defer span.End()

	household = &types.Household{Members: []*types.HouseholdUserMembershipWithUser{}}
	membership = &types.HouseholdUserMembershipWithUser{BelongsToUser: &types.User{}}

	targetVars := []any{
		&household.ID,
		&household.Name,
		&household.BillingStatus,
		&household.ContactPhone,
		&household.AddressLine1,
		&household.AddressLine2,
		&household.City,
		&household.State,
		&household.ZipCode,
		&household.Country,
		&household.Latitude,
		&household.Longitude,
		&household.PaymentProcessorCustomerID,
		&household.SubscriptionPlanID,
		&household.CreatedAt,
		&household.LastUpdatedAt,
		&household.ArchivedAt,
		&household.BelongsToUser,
		&membership.BelongsToUser.ID,
		&membership.BelongsToUser.FirstName,
		&membership.BelongsToUser.LastName,
		&membership.BelongsToUser.Username,
		&membership.BelongsToUser.EmailAddress,
		&membership.BelongsToUser.EmailAddressVerifiedAt,
		&membership.BelongsToUser.AvatarSrc,
		&membership.BelongsToUser.RequiresPasswordChange,
		&membership.BelongsToUser.PasswordLastChangedAt,
		&membership.BelongsToUser.TwoFactorSecretVerifiedAt,
		&membership.BelongsToUser.ServiceRole,
		&membership.BelongsToUser.AccountStatus,
		&membership.BelongsToUser.AccountStatusExplanation,
		&membership.BelongsToUser.Birthday,
		&membership.BelongsToUser.CreatedAt,
		&membership.BelongsToUser.LastUpdatedAt,
		&membership.BelongsToUser.ArchivedAt,
		&membership.ID,
		&membership.BelongsToUser.ID,
		&membership.BelongsToHousehold,
		&membership.HouseholdRole,
		&membership.DefaultHousehold,
		&membership.CreatedAt,
		&membership.LastUpdatedAt,
		&membership.ArchivedAt,
	}

	if includeCounts {
		targetVars = append(targetVars, &filteredCount, &totalCount)
	}

	if err = scan.Scan(targetVars...); err != nil {
		return nil, nil, 0, 0, observability.PrepareError(err, span, "fetching memberships from database")
	}

	return household, membership, filteredCount, totalCount, nil
}

// scanHouseholds takes some database rows and turns them into a slice of households.
func (q *Querier) scanHouseholds(ctx context.Context, rows database.ResultIterator, includeCounts bool) (households []*types.Household, filteredCount, totalCount uint64, err error) {
	_, span := q.tracer.StartSpan(ctx)
	defer span.End()

	households = []*types.Household{}

	var currentHousehold *types.Household
	for rows.Next() {
		household, membership, fc, tc, scanErr := q.scanHousehold(ctx, rows, includeCounts)
		if scanErr != nil {
			return nil, 0, 0, scanErr
		}

		if currentHousehold == nil {
			currentHousehold = household
		}

		if currentHousehold.ID != household.ID {
			households = append(households, currentHousehold)
			currentHousehold = household
		}

		currentHousehold.Members = append(currentHousehold.Members, membership)

		if includeCounts {
			if filteredCount == 0 {
				filteredCount = fc
			}

			if totalCount == 0 {
				totalCount = tc
			}
		}
	}

	if currentHousehold != nil {
		households = append(households, currentHousehold)
	}

	if err = q.checkRowsForErrorAndClose(ctx, rows); err != nil {
		return nil, 0, 0, observability.PrepareError(err, span, "handling rows")
	}

	return households, filteredCount, totalCount, nil
}

//go:embed queries/households/get_by_id_with_memberships.sql
var getHouseholdAndMembershipsByIDQuery string

// GetHousehold fetches a household from the database.
func (q *Querier) GetHousehold(ctx context.Context, householdID string) (*types.Household, error) {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	if householdID == "" {
		return nil, ErrInvalidIDProvided
	}
	tracing.AttachHouseholdIDToSpan(span, householdID)

	args := []any{
		householdID,
	}

	rows, err := q.getRows(ctx, q.db, "household", getHouseholdAndMembershipsByIDQuery, args)
	if err != nil {
		return nil, observability.PrepareError(err, span, "executing households list retrieval query")
	}

	households, _, _, err := q.scanHouseholds(ctx, rows, false)
	if err != nil {
		return nil, observability.PrepareError(err, span, "beginning transaction")
	}

	var household *types.Household
	if len(households) > 0 {
		household = households[0]
	}

	if household == nil {
		return nil, sql.ErrNoRows
	}

	return household, nil
}

// GetHouseholdByID fetches a household from the database by its ID.
func (q *Querier) GetHouseholdByID(ctx context.Context, householdID string) (*types.Household, error) {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	if householdID == "" {
		return nil, ErrInvalidIDProvided
	}
	tracing.AttachHouseholdIDToSpan(span, householdID)

	args := []any{
		householdID,
	}

	rows, err := q.getRows(ctx, q.db, "household by id", getHouseholdAndMembershipsByIDQuery, args)
	if err != nil {
		return nil, observability.PrepareError(err, span, "executing households list retrieval query")
	}

	households, _, _, err := q.scanHouseholds(ctx, rows, false)
	if err != nil {
		return nil, observability.PrepareError(err, span, "beginning transaction")
	}

	var household *types.Household
	if len(households) > 0 {
		household = households[0]
	}

	if household == nil {
		return nil, sql.ErrNoRows
	}

	return household, nil
}

// buildGetHouseholdsQuery builds a SQL query selecting households that adhere to a given QueryFilter and belong to a given household,
// and returns both the query and the relevant args to pass to the query executor.
func (q *Querier) buildGetHouseholdsQuery(ctx context.Context, userID string, forAdmin bool, filter *types.QueryFilter) (query string, args []any) {
	_, span := q.tracer.StartSpan(ctx)
	defer span.End()

	tracing.AttachUserIDToSpan(span, userID)

	if filter != nil {
		tracing.AttachFilterDataToSpan(span, filter.Page, filter.Limit, filter.SortBy)
	}

	var includeArchived bool
	if filter != nil && filter.IncludeArchived != nil {
		includeArchived = *filter.IncludeArchived
	}

	filteredCountQuery, filteredCountQueryArgs := q.buildFilteredCountQuery(ctx, householdsTableName, nil, nil, userOwnershipColumn, userID, forAdmin, includeArchived, filter)
	totalCountQuery, totalCountQueryArgs := q.buildTotalCountQuery(ctx, householdsTableName, nil, nil, userOwnershipColumn, userID, forAdmin, includeArchived)

	builder := q.sqlBuilder.Select(append(
		append(householdsTableColumns, householdsUserMembershipTableColumns...),
		fmt.Sprintf("(%s) as filtered_count", filteredCountQuery),
		fmt.Sprintf("(%s) as total_count", totalCountQuery),
	)...).
		From(householdsTableName).
		Join(householdUserMembershipsOnHouseholdsJoinClause).
		Join(usersOnHouseholdUserMembershipsJoinClause)

	if !forAdmin {
		where := squirrel.Eq{
			"households.archived_at":                 nil,
			"household_user_memberships.archived_at": nil,
		}

		if userID != "" {
			where["household_user_memberships.belongs_to_user"] = userID
		}

		builder = builder.Where(where)
	}

	builder = builder.GroupBy(fmt.Sprintf(
		"%s.id, users.id, %s.id",
		householdsTableName,
		householdsUserMembershipTableName,
	))

	if filter != nil {
		builder = applyFilterToQueryBuilder(filter, householdsTableName, builder)
	}

	query, selectArgs := q.buildQuery(span, builder)

	return query, append(append(filteredCountQueryArgs, totalCountQueryArgs...), selectArgs...)
}

// getHouseholdsForUser fetches a list of households from the database that meet a particular filter.
func (q *Querier) getHouseholdsForUser(ctx context.Context, querier database.SQLQueryExecutor, userID string, forAdmin bool, filter *types.QueryFilter) (x *types.QueryFilteredResult[types.Household], err error) {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	logger := q.logger.Clone()

	if userID == "" && !forAdmin {
		return nil, ErrInvalidIDProvided
	}
	tracing.AttachUserIDToSpan(span, userID)

	if filter == nil {
		filter = types.DefaultQueryFilter()
	}
	logger = filter.AttachToLogger(logger)
	tracing.AttachQueryFilterToSpan(span, filter)

	x = &types.QueryFilteredResult[types.Household]{
		Pagination: filter.ToPagination(),
	}

	query, args := q.buildGetHouseholdsQuery(ctx, userID, forAdmin, filter)

	rows, err := q.getRows(ctx, querier, "households", query, args)
	if err != nil {
		return nil, observability.PrepareError(err, span, "executing households list retrieval query")
	}

	if x.Data, x.FilteredCount, x.TotalCount, err = q.scanHouseholds(ctx, rows, true); err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "scanning households from database")
	}

	return x, nil
}

// GetHouseholds fetches a list of households from the database that meet a particular filter.
func (q *Querier) GetHouseholds(ctx context.Context, userID string, filter *types.QueryFilter) (x *types.QueryFilteredResult[types.Household], err error) {
	return q.getHouseholdsForUser(ctx, q.db, userID, false, filter)
}

// GetHouseholdsForAdmin fetches a list of households from the database that meet a particular filter for all users.
func (q *Querier) GetHouseholdsForAdmin(ctx context.Context, userID string, filter *types.QueryFilter) (x *types.QueryFilteredResult[types.Household], err error) {
	return q.getHouseholdsForUser(ctx, q.db, userID, true, filter)
}

// CreateHousehold creates a household in the database.
func (q *Querier) CreateHousehold(ctx context.Context, input *types.HouseholdDatabaseCreationInput) (*types.Household, error) {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	if input == nil {
		return nil, ErrNilInputProvided
	}

	logger := q.logger.WithValue(keys.UserIDKey, input.BelongsToUser)

	// begin household creation transaction
	tx, err := q.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "beginning transaction")
	}

	// create the household.
	if writeErr := q.generatedQuerier.CreateHousehold(ctx, tx, &generated.CreateHouseholdParams{
		City:          input.City,
		Name:          input.Name,
		BillingStatus: string(types.UnpaidHouseholdBillingStatus),
		ContactPhone:  input.ContactPhone,
		AddressLine1:  input.AddressLine1,
		AddressLine2:  input.AddressLine2,
		ID:            input.ID,
		State:         input.State,
		ZipCode:       input.ZipCode,
		Country:       input.Country,
		BelongsToUser: input.BelongsToUser,
		Latitude:      nullStringFromFloat64Pointer(input.Latitude),
		Longitude:     nullStringFromFloat64Pointer(input.Longitude),
	}); writeErr != nil {
		q.rollbackTransaction(ctx, tx)
		return nil, observability.PrepareError(writeErr, span, "creating household")
	}

	household := &types.Household{
		ID:            input.ID,
		Name:          input.Name,
		BelongsToUser: input.BelongsToUser,
		BillingStatus: string(types.UnpaidHouseholdBillingStatus),
		ContactPhone:  input.ContactPhone,
		AddressLine1:  input.AddressLine1,
		AddressLine2:  input.AddressLine2,
		City:          input.City,
		State:         input.State,
		ZipCode:       input.ZipCode,
		Country:       input.Country,
		Latitude:      input.Latitude,
		Longitude:     input.Longitude,
		CreatedAt:     q.currentTime(),
	}

	if err = q.generatedQuerier.AddUserToHousehold(ctx, tx, &generated.AddUserToHouseholdParams{
		ID:                 identifiers.New(),
		BelongsToUser:      input.BelongsToUser,
		BelongsToHousehold: household.ID,
		HouseholdRole:      authorization.HouseholdAdminRole.String(),
	}); err != nil {
		q.rollbackTransaction(ctx, tx)
		return nil, observability.PrepareAndLogError(err, logger, span, "performing household membership creation query")
	}

	if err = tx.Commit(); err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "committing transaction")
	}

	tracing.AttachHouseholdIDToSpan(span, household.ID)
	logger.Info("household created")

	return household, nil
}

// UpdateHousehold updates a particular household. Note that UpdateHousehold expects the provided input to have a valid ID.
func (q *Querier) UpdateHousehold(ctx context.Context, updated *types.Household) error {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	if updated == nil {
		return ErrNilInputProvided
	}
	logger := q.logger.WithValue(keys.HouseholdIDKey, updated.ID)
	tracing.AttachHouseholdIDToSpan(span, updated.ID)

	if err := q.generatedQuerier.UpdateHousehold(ctx, q.db, &generated.UpdateHouseholdParams{
		Name:          updated.Name,
		ContactPhone:  updated.ContactPhone,
		AddressLine1:  updated.AddressLine1,
		AddressLine2:  updated.AddressLine2,
		City:          updated.City,
		State:         updated.State,
		ZipCode:       updated.ZipCode,
		Country:       updated.Country,
		BelongsToUser: updated.BelongsToUser,
		ID:            updated.ID,
		Latitude:      nullStringFromFloat64Pointer(updated.Latitude),
		Longitude:     nullStringFromFloat64Pointer(updated.Longitude),
	}); err != nil {
		return observability.PrepareAndLogError(err, logger, span, "updating household")
	}

	logger.Info("household updated")

	return nil
}

// ArchiveHousehold archives a household from the database by its ID.
func (q *Querier) ArchiveHousehold(ctx context.Context, householdID, userID string) error {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	if householdID == "" || userID == "" {
		return ErrInvalidIDProvided
	}
	tracing.AttachUserIDToSpan(span, userID)
	tracing.AttachHouseholdIDToSpan(span, householdID)

	logger := q.logger.WithValues(map[string]any{
		keys.HouseholdIDKey: householdID,
		keys.UserIDKey:      userID,
	})

	if err := q.generatedQuerier.ArchiveHousehold(ctx, q.db, &generated.ArchiveHouseholdParams{
		BelongsToUser: userID,
		ID:            householdID,
	}); err != nil {
		return observability.PrepareAndLogError(err, logger, span, "archiving household")
	}

	logger.Info("household archived")

	return nil
}
