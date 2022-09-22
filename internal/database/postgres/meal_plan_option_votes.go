package postgres

import (
	"context"
	_ "embed"

	"github.com/prixfixeco/api_server/internal/database"
	"github.com/prixfixeco/api_server/internal/observability"
	"github.com/prixfixeco/api_server/internal/observability/keys"
	"github.com/prixfixeco/api_server/internal/observability/tracing"
	"github.com/prixfixeco/api_server/pkg/types"
)

const (
	mealPlanOptionsOnMealPlanOptionVotesJoinClause = "meal_plan_options ON meal_plan_option_votes.belongs_to_meal_plan_option=meal_plan_options.id"
)

var (
	_ types.MealPlanOptionVoteDataManager = (*Querier)(nil)

	// mealPlanOptionVotesTableColumns are the columns for the meal_plan_option_votes table.
	mealPlanOptionVotesTableColumns = []string{
		"meal_plan_option_votes.id",
		"meal_plan_option_votes.rank",
		"meal_plan_option_votes.abstain",
		"meal_plan_option_votes.notes",
		"meal_plan_option_votes.by_user",
		"meal_plan_option_votes.created_at",
		"meal_plan_option_votes.last_updated_at",
		"meal_plan_option_votes.archived_at",
		"meal_plan_option_votes.belongs_to_meal_plan_option",
	}

	getMealPlanOptionVotesJoins = []string{
		mealPlanOptionsOnMealPlanOptionVotesJoinClause,
		mealPlanEventsOnMealPlanOptionsJoinClause,
		mealPlansOnMealPlanEventsJoinClause,
	}
)

// scanMealPlanOptionVote takes a database Scanner (i.e. *sql.Row) and scans the result into a meal plan option vote struct.
func (q *Querier) scanMealPlanOptionVote(ctx context.Context, scan database.Scanner, includeCounts bool) (x *types.MealPlanOptionVote, filteredCount, totalCount uint64, err error) {
	_, span := q.tracer.StartSpan(ctx)
	defer span.End()

	x = &types.MealPlanOptionVote{}

	targetVars := []interface{}{
		&x.ID,
		&x.Rank,
		&x.Abstain,
		&x.Notes,
		&x.ByUser,
		&x.CreatedAt,
		&x.LastUpdatedAt,
		&x.ArchivedAt,
		&x.BelongsToMealPlanOption,
	}

	if includeCounts {
		targetVars = append(targetVars, &filteredCount, &totalCount)
	}

	if err = scan.Scan(targetVars...); err != nil {
		return nil, 0, 0, observability.PrepareError(err, span, "")
	}

	return x, filteredCount, totalCount, nil
}

// scanMealPlanOptionVotes takes some database rows and turns them into a slice of meal plan option votes.
func (q *Querier) scanMealPlanOptionVotes(ctx context.Context, rows database.ResultIterator, includeCounts bool) (mealPlanOptionVotes []*types.MealPlanOptionVote, filteredCount, totalCount uint64, err error) {
	_, span := q.tracer.StartSpan(ctx)
	defer span.End()

	for rows.Next() {
		x, fc, tc, scanErr := q.scanMealPlanOptionVote(ctx, rows, includeCounts)
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

		mealPlanOptionVotes = append(mealPlanOptionVotes, x)
	}

	if err = q.checkRowsForErrorAndClose(ctx, rows); err != nil {
		return nil, 0, 0, observability.PrepareError(err, span, "handling rows")
	}

	return mealPlanOptionVotes, filteredCount, totalCount, nil
}

//go:embed queries/meal_plan_option_votes/exists.sql
var mealPlanOptionVoteExistenceQuery string

// MealPlanOptionVoteExists fetches whether a meal plan option vote exists from the database.
func (q *Querier) MealPlanOptionVoteExists(ctx context.Context, mealPlanID, mealPlanEventID, mealPlanOptionID, mealPlanOptionVoteID string) (exists bool, err error) {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	logger := q.logger.Clone()

	if mealPlanID == "" {
		return false, ErrInvalidIDProvided
	}
	logger = logger.WithValue(keys.MealPlanIDKey, mealPlanID)
	tracing.AttachMealPlanIDToSpan(span, mealPlanID)

	if mealPlanEventID == "" {
		return false, ErrInvalidIDProvided
	}
	logger = logger.WithValue(keys.MealPlanEventIDKey, mealPlanEventID)
	tracing.AttachMealPlanEventIDToSpan(span, mealPlanEventID)

	if mealPlanOptionID == "" {
		return false, ErrInvalidIDProvided
	}
	logger = logger.WithValue(keys.MealPlanOptionIDKey, mealPlanOptionID)
	tracing.AttachMealPlanOptionIDToSpan(span, mealPlanOptionID)

	if mealPlanOptionVoteID == "" {
		return false, ErrInvalidIDProvided
	}
	logger = logger.WithValue(keys.MealPlanOptionVoteIDKey, mealPlanOptionVoteID)
	tracing.AttachMealPlanOptionVoteIDToSpan(span, mealPlanOptionVoteID)

	args := []interface{}{
		mealPlanOptionID,
		mealPlanOptionVoteID,
		mealPlanEventID,
		mealPlanID,
	}

	result, err := q.performBooleanQuery(ctx, q.db, mealPlanOptionVoteExistenceQuery, args)
	if err != nil {
		return false, observability.PrepareAndLogError(err, logger, span, "performing meal plan option vote existence check")
	}

	return result, nil
}

//go:embed queries/meal_plan_option_votes/get_one.sql
var getMealPlanOptionVoteQuery string

// GetMealPlanOptionVote fetches a meal plan option vote from the database.
func (q *Querier) GetMealPlanOptionVote(ctx context.Context, mealPlanID, mealPlanEventID, mealPlanOptionID, mealPlanOptionVoteID string) (*types.MealPlanOptionVote, error) {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	logger := q.logger.Clone()

	if mealPlanID == "" {
		return nil, ErrInvalidIDProvided
	}
	logger = logger.WithValue(keys.MealPlanIDKey, mealPlanID)
	tracing.AttachMealPlanIDToSpan(span, mealPlanID)

	if mealPlanOptionID == "" {
		return nil, ErrInvalidIDProvided
	}
	logger = logger.WithValue(keys.MealPlanOptionIDKey, mealPlanOptionID)
	tracing.AttachMealPlanOptionIDToSpan(span, mealPlanOptionID)

	if mealPlanOptionVoteID == "" {
		return nil, ErrInvalidIDProvided
	}
	logger = logger.WithValue(keys.MealPlanOptionVoteIDKey, mealPlanOptionVoteID)
	tracing.AttachMealPlanOptionVoteIDToSpan(span, mealPlanOptionVoteID)

	args := []interface{}{
		mealPlanOptionID,
		mealPlanOptionVoteID,
		mealPlanEventID,
		mealPlanID,
	}

	row := q.getOneRow(ctx, q.db, "meal plan option vote", getMealPlanOptionVoteQuery, args)

	mealPlanOptionVote, _, _, err := q.scanMealPlanOptionVote(ctx, row, false)
	if err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "scanning mealPlanOptionVote")
	}

	return mealPlanOptionVote, nil
}

//go:embed queries/meal_plan_option_votes/get_for_meal_plan_option.sql
var getMealPlanOptionVotesForMealPlanOptionQuery string

// GetMealPlanOptionVotesForMealPlanOption fetches a list of meal plan option votes from the database that meet a particular filter.
func (q *Querier) GetMealPlanOptionVotesForMealPlanOption(ctx context.Context, mealPlanID, mealPlanEventID, mealPlanOptionID string) (x []*types.MealPlanOptionVote, err error) {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	logger := q.logger.Clone()

	if mealPlanID == "" {
		return nil, ErrInvalidIDProvided
	}
	logger = logger.WithValue(keys.MealPlanIDKey, mealPlanID)
	tracing.AttachMealPlanIDToSpan(span, mealPlanID)

	if mealPlanEventID == "" {
		return nil, ErrInvalidIDProvided
	}
	logger = logger.WithValue(keys.MealPlanEventIDKey, mealPlanEventID)
	tracing.AttachMealPlanEventIDToSpan(span, mealPlanEventID)

	if mealPlanOptionID == "" {
		return nil, ErrInvalidIDProvided
	}
	logger = logger.WithValue(keys.MealPlanOptionIDKey, mealPlanOptionID)
	tracing.AttachMealPlanOptionIDToSpan(span, mealPlanOptionID)

	args := []interface{}{
		mealPlanID,
		mealPlanEventID,
		mealPlanOptionID,
	}

	rows, err := q.performReadQuery(ctx, q.db, "meal plan option votes for meal plan option", getMealPlanOptionVotesForMealPlanOptionQuery, args)
	if err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "executing meal plan option votes for meal plan option list retrieval query")
	}

	x, _, _, err = q.scanMealPlanOptionVotes(ctx, rows, false)
	if err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "scanning meal plan option votes")
	}

	return x, nil
}

// GetMealPlanOptionVotes fetches a list of meal plan option votes from the database that meet a particular filter.
func (q *Querier) GetMealPlanOptionVotes(ctx context.Context, mealPlanID, mealPlanEventID, mealPlanOptionID string, filter *types.QueryFilter) (x *types.MealPlanOptionVoteList, err error) {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	logger := q.logger.Clone()

	if mealPlanID == "" {
		return nil, ErrInvalidIDProvided
	}
	logger = logger.WithValue(keys.MealPlanIDKey, mealPlanID)
	tracing.AttachMealPlanIDToSpan(span, mealPlanID)

	if mealPlanEventID == "" {
		return nil, ErrInvalidIDProvided
	}
	logger = logger.WithValue(keys.MealPlanEventIDKey, mealPlanEventID)
	tracing.AttachMealPlanEventIDToSpan(span, mealPlanEventID)

	if mealPlanOptionID == "" {
		return nil, ErrInvalidIDProvided
	}
	logger = logger.WithValue(keys.MealPlanOptionIDKey, mealPlanOptionID)
	tracing.AttachMealPlanOptionIDToSpan(span, mealPlanOptionID)

	x = &types.MealPlanOptionVoteList{}
	logger = filter.AttachToLogger(logger)
	tracing.AttachQueryFilterToSpan(span, filter)

	if filter != nil {
		if filter.Page != nil {
			x.Page = *filter.Page
		}

		if filter.Limit != nil {
			x.Limit = *filter.Limit
		}
	}

	query, args := q.buildListQuery(ctx, "meal_plan_option_votes", getMealPlanOptionVotesJoins, nil, nil, householdOwnershipColumn, mealPlanOptionVotesTableColumns, "", false, filter)

	rows, err := q.performReadQuery(ctx, q.db, "meal plan option votes", query, args)
	if err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "executing meal plan option votes list retrieval query")
	}

	if x.MealPlanOptionVotes, x.FilteredCount, x.TotalCount, err = q.scanMealPlanOptionVotes(ctx, rows, true); err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "scanning meal plan option votes")
	}

	return x, nil
}

//go:embed queries/meal_plan_option_votes/create.sql
var mealPlanOptionVoteCreationQuery string

// CreateMealPlanOptionVote creates a meal plan option vote in the database.
func (q *Querier) CreateMealPlanOptionVote(ctx context.Context, input *types.MealPlanOptionVoteDatabaseCreationInput) ([]*types.MealPlanOptionVote, error) {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	if input == nil {
		return nil, ErrNilInputProvided
	}

	logger := q.logger.WithValue("vote_count", len(input.Votes)).WithValue(keys.UserIDKey, input.ByUser)

	// begin transaction
	tx, err := q.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "beginning transaction")
	}

	votes := []*types.MealPlanOptionVote{}
	for _, vote := range input.Votes {
		args := []interface{}{
			vote.ID,
			vote.Rank,
			vote.Abstain,
			vote.Notes,
			vote.ByUser,
			vote.BelongsToMealPlanOption,
		}

		// create the meal plan option vote.
		if err = q.performWriteQuery(ctx, tx, "meal plan option vote creation", mealPlanOptionVoteCreationQuery, args); err != nil {
			q.rollbackTransaction(ctx, tx)
			return nil, observability.PrepareAndLogError(err, logger, span, "creating meal plan option vote")
		}

		x := &types.MealPlanOptionVote{
			ID:                      vote.ID,
			Rank:                    vote.Rank,
			Abstain:                 vote.Abstain,
			Notes:                   vote.Notes,
			ByUser:                  vote.ByUser,
			BelongsToMealPlanOption: vote.BelongsToMealPlanOption,
			CreatedAt:               q.currentTime(),
		}

		tracing.AttachMealPlanOptionVoteIDToSpan(span, x.ID)
		logger.Info("meal plan option vote created")

		votes = append(votes, x)
	}

	if err = tx.Commit(); err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "committing transaction")
	}

	return votes, nil
}

//go:embed queries/meal_plan_option_votes/update.sql
var updateMealPlanOptionVoteQuery string

// UpdateMealPlanOptionVote updates a particular meal plan option vote.
func (q *Querier) UpdateMealPlanOptionVote(ctx context.Context, updated *types.MealPlanOptionVote) error {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	if updated == nil {
		return ErrNilInputProvided
	}

	logger := q.logger.WithValue(keys.MealPlanOptionVoteIDKey, updated.ID)
	tracing.AttachMealPlanOptionVoteIDToSpan(span, updated.ID)

	args := []interface{}{
		updated.Rank,
		updated.Abstain,
		updated.Notes,
		updated.ByUser,
		updated.BelongsToMealPlanOption,
		updated.ID,
	}

	if err := q.performWriteQuery(ctx, q.db, "meal plan option vote update", updateMealPlanOptionVoteQuery, args); err != nil {
		return observability.PrepareAndLogError(err, logger, span, "updating meal plan option vote")
	}

	logger.Info("meal plan option vote updated")

	return nil
}

//go:embed queries/meal_plan_option_votes/archive.sql
var archiveMealPlanOptionVoteQuery string

// ArchiveMealPlanOptionVote archives a meal plan option vote from the database by its ID.
func (q *Querier) ArchiveMealPlanOptionVote(ctx context.Context, mealPlanID, mealPlanEventID, mealPlanOptionID, mealPlanOptionVoteID string) error {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	logger := q.logger.Clone()

	if mealPlanID == "" {
		return ErrInvalidIDProvided
	}
	logger = logger.WithValue(keys.MealPlanIDKey, mealPlanID)
	tracing.AttachMealPlanIDToSpan(span, mealPlanID)

	if mealPlanEventID == "" {
		return ErrInvalidIDProvided
	}
	logger = logger.WithValue(keys.MealPlanEventIDKey, mealPlanEventID)
	tracing.AttachMealPlanEventIDToSpan(span, mealPlanEventID)

	if mealPlanOptionID == "" {
		return ErrInvalidIDProvided
	}
	logger = logger.WithValue(keys.MealPlanOptionIDKey, mealPlanOptionID)
	tracing.AttachMealPlanOptionIDToSpan(span, mealPlanOptionID)

	if mealPlanOptionVoteID == "" {
		return ErrInvalidIDProvided
	}
	logger = logger.WithValue(keys.MealPlanOptionVoteIDKey, mealPlanOptionVoteID)
	tracing.AttachMealPlanOptionVoteIDToSpan(span, mealPlanOptionVoteID)

	args := []interface{}{
		mealPlanOptionID,
		mealPlanOptionVoteID,
	}

	if err := q.performWriteQuery(ctx, q.db, "meal plan option vote archive", archiveMealPlanOptionVoteQuery, args); err != nil {
		return observability.PrepareAndLogError(err, logger, span, "updating meal plan option vote")
	}

	logger.Info("meal plan option vote archived")

	return nil
}