package postgres

import (
	"context"
	"database/sql"
	_ "embed"
	"math/rand"

	"github.com/dinnerdonebetter/backend/internal/database"
	"github.com/dinnerdonebetter/backend/internal/database/postgres/generated"
	"github.com/dinnerdonebetter/backend/internal/observability"
	"github.com/dinnerdonebetter/backend/internal/observability/keys"
	"github.com/dinnerdonebetter/backend/internal/observability/tracing"
	"github.com/dinnerdonebetter/backend/pkg/types"

	"resenje.org/schulze"
)

const (
	mealsOnMealPlanOptionsJoinClause = "meals ON meal_plan_options.meal_id=meals.id"
)

var (
	_ types.MealPlanOptionDataManager = (*Querier)(nil)

	// mealPlanOptionsTableColumns are the columns for the meal_plan_options table.
	mealPlanOptionsTableColumns = []string{
		"meal_plan_options.id",
		"meal_plan_options.assigned_cook",
		"meal_plan_options.assigned_dishwasher",
		"meal_plan_options.chosen",
		"meal_plan_options.tiebroken",
		"meal_plan_options.meal_scale",
		"meal_plan_options.meal_id",
		"meal_plan_options.notes",
		"meal_plan_options.created_at",
		"meal_plan_options.last_updated_at",
		"meal_plan_options.archived_at",
		"meal_plan_options.belongs_to_meal_plan_event",
		"meals.id",
		"meals.name",
		"meals.description",
		"meals.min_estimated_portions",
		"meals.max_estimated_portions",
		"meals.eligible_for_meal_plans",
		"meals.created_at",
		"meals.last_updated_at",
		"meals.archived_at",
		"meals.created_by_user",
	}

	getMealPlanOptionsJoins = []string{
		mealPlanEventsOnMealPlanOptionsJoinClause,
		mealsOnMealPlanOptionsJoinClause,
		mealPlansOnMealPlanEventsJoinClause,
	}
)

// scanMealPlanOption takes a database Scanner (i.e. *sql.Row) and scans the result into a meal plan option struct.
func (q *Querier) scanMealPlanOption(ctx context.Context, scan database.Scanner, includeCounts bool) (x *types.MealPlanOption, filteredCount, totalCount uint64, err error) {
	_, span := q.tracer.StartSpan(ctx)
	defer span.End()

	x = &types.MealPlanOption{
		Votes: []*types.MealPlanOptionVote{},
	}

	targetVars := []any{
		&x.ID,
		&x.AssignedCook,
		&x.AssignedDishwasher,
		&x.Chosen,
		&x.TieBroken,
		&x.MealScale,
		&x.Meal.ID,
		&x.Notes,
		&x.CreatedAt,
		&x.LastUpdatedAt,
		&x.ArchivedAt,
		&x.BelongsToMealPlanEvent,
		&x.Meal.ID,
		&x.Meal.Name,
		&x.Meal.Description,
		&x.Meal.MinimumEstimatedPortions,
		&x.Meal.MaximumEstimatedPortions,
		&x.Meal.EligibleForMealPlans,
		&x.Meal.CreatedAt,
		&x.Meal.LastUpdatedAt,
		&x.Meal.ArchivedAt,
		&x.Meal.CreatedByUser,
	}

	if includeCounts {
		targetVars = append(targetVars, &filteredCount, &totalCount)
	}

	if err = scan.Scan(targetVars...); err != nil {
		return nil, 0, 0, observability.PrepareError(err, span, "")
	}

	return x, filteredCount, totalCount, nil
}

// scanMealPlanOptions takes some database rows and turns them into a slice of meal plan options.
func (q *Querier) scanMealPlanOptions(ctx context.Context, rows database.ResultIterator, includeCounts bool) (mealPlanOptions []*types.MealPlanOption, filteredCount, totalCount uint64, err error) {
	_, span := q.tracer.StartSpan(ctx)
	defer span.End()

	for rows.Next() {
		x, fc, tc, scanErr := q.scanMealPlanOption(ctx, rows, includeCounts)
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

		mealPlanOptions = append(mealPlanOptions, x)
	}

	if err = q.checkRowsForErrorAndClose(ctx, rows); err != nil {
		return nil, 0, 0, observability.PrepareError(err, span, "handling rows")
	}

	return mealPlanOptions, filteredCount, totalCount, nil
}

// MealPlanOptionExists fetches whether a meal plan option exists from the database.
func (q *Querier) MealPlanOptionExists(ctx context.Context, mealPlanID, mealPlanEventID, mealPlanOptionID string) (exists bool, err error) {
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

	result, err := q.generatedQuerier.CheckMealPlanOptionExistence(ctx, q.db, &generated.CheckMealPlanOptionExistenceParams{
		MealPlanEventID:  nullStringFromString(mealPlanEventID),
		MealPlanOptionID: mealPlanOptionID,
		MealPlanID:       mealPlanID,
	})
	if err != nil {
		logger.Error(err, "performing meal plan option existence check")
		return false, observability.PrepareAndLogError(err, logger, span, "performing meal plan option existence check")
	}

	return result, nil
}

//go:embed queries/meal_plan_options/get_one.sql
var getMealPlanOptionQuery string

// GetMealPlanOption fetches a meal plan option from the database.
func (q *Querier) GetMealPlanOption(ctx context.Context, mealPlanID, mealPlanEventID, mealPlanOptionID string) (*types.MealPlanOption, error) {
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

	args := []any{
		mealPlanID,
		mealPlanEventID,
		mealPlanOptionID,
	}

	row := q.getOneRow(ctx, q.db, "meal plan option", getMealPlanOptionQuery, args)

	mealPlanOption, _, _, err := q.scanMealPlanOption(ctx, row, false)
	if err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "scanning meal plan option")
	}

	return mealPlanOption, nil
}

//go:embed queries/meal_plan_options/get_one_by_id.sql
var getMealPlanOptionByIDQuery string

// getMealPlanOptionByID fetches a meal plan option from the database by its ID.
func (q *Querier) getMealPlanOptionByID(ctx context.Context, mealPlanOptionID string) (*types.MealPlanOption, error) {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	logger := q.logger.Clone()

	if mealPlanOptionID == "" {
		return nil, ErrInvalidIDProvided
	}
	logger = logger.WithValue(keys.MealPlanOptionIDKey, mealPlanOptionID)
	tracing.AttachMealPlanOptionIDToSpan(span, mealPlanOptionID)

	args := []any{
		mealPlanOptionID,
	}

	row := q.getOneRow(ctx, q.db, "meal plan option by ID", getMealPlanOptionByIDQuery, args)

	mealPlanOption, _, _, err := q.scanMealPlanOption(ctx, row, false)
	if err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "scanning meal plan option")
	}

	return mealPlanOption, nil
}

//go:embed queries/meal_plan_options/get_for_meal_plan_event.sql
var getMealPlanOptionsForMealPlanEventQuery string

// getMealPlanOptionsForMealPlanEvent fetches a list of meal plan options from the database that meet a particular filter.
func (q *Querier) getMealPlanOptionsForMealPlanEvent(ctx context.Context, mealPlanID, mealPlanEventID string) (x []*types.MealPlanOption, err error) {
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
	tracing.AttachMealPlanIDToSpan(span, mealPlanEventID)

	args := []any{
		mealPlanEventID,
		mealPlanID,
	}

	rows, err := q.getRows(ctx, q.db, "meal plan options for meal plan event", getMealPlanOptionsForMealPlanEventQuery, args)
	if err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "executing meal plan options for meal plan event list retrieval query")
	}

	x, _, _, err = q.scanMealPlanOptions(ctx, rows, false)
	if err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "scanning meal plan options for meal plan event")
	}

	for i, opt := range x {
		votes, voteFetchErr := q.GetMealPlanOptionVotesForMealPlanOption(ctx, mealPlanID, mealPlanEventID, opt.ID)
		if voteFetchErr != nil {
			return nil, observability.PrepareError(voteFetchErr, span, "fetching meal plan option votes for meal plan option")
		}
		x[i].Votes = votes

		m, mealFetchErr := q.GetMeal(ctx, opt.Meal.ID)
		if mealFetchErr != nil {
			return nil, observability.PrepareAndLogError(mealFetchErr, logger, span, "scanning meal plan options for meal plan event")
		}
		x[i].Meal = *m
	}

	logger.WithValue("quantity", len(x)).Info("fetched meal plan options for meal plan event")

	return x, nil
}

// GetMealPlanOptions fetches a list of meal plan options from the database that meet a particular filter.
func (q *Querier) GetMealPlanOptions(ctx context.Context, mealPlanID, mealPlanEventID string, filter *types.QueryFilter) (x *types.QueryFilteredResult[types.MealPlanOption], err error) {
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

	if filter == nil {
		filter = types.DefaultQueryFilter()
	}
	logger = filter.AttachToLogger(logger)
	tracing.AttachQueryFilterToSpan(span, filter)

	x = &types.QueryFilteredResult[types.MealPlanOption]{
		Pagination: filter.ToPagination(),
	}

	groupBys := []string{"meal_plan_options.id", "meals.id"}
	query, args := q.buildListQuery(ctx, "meal_plan_options", getMealPlanOptionsJoins, groupBys, nil, householdOwnershipColumn, mealPlanOptionsTableColumns, "", false, filter)

	rows, err := q.getRows(ctx, q.db, "meal plan options", query, args)
	if err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "executing meal plan options list retrieval query")
	}

	if x.Data, x.FilteredCount, x.TotalCount, err = q.scanMealPlanOptions(ctx, rows, true); err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "scanning meal plan options")
	}

	return x, nil
}

// createMealPlanOption creates a meal plan option in the database.
func (q *Querier) createMealPlanOption(ctx context.Context, db database.SQLQueryExecutor, input *types.MealPlanOptionDatabaseCreationInput, markAsChosen bool) (*types.MealPlanOption, error) {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	if input == nil {
		return nil, ErrNilInputProvided
	}
	tracing.AttachMealPlanOptionIDToSpan(span, input.ID)
	logger := q.logger.WithValue(keys.MealPlanOptionIDKey, input.ID)

	// create the meal plan option.
	if err := q.generatedQuerier.CreateMealPlanOption(ctx, db, &generated.CreateMealPlanOptionParams{
		ID:                     input.ID,
		AssignedCook:           nullStringFromStringPointer(input.AssignedCook),
		AssignedDishwasher:     nullStringFromStringPointer(input.AssignedDishwasher),
		MealID:                 input.MealID,
		Notes:                  input.Notes,
		MealScale:              stringFromFloat32(input.MealScale),
		BelongsToMealPlanEvent: nullStringFromString(input.BelongsToMealPlanEvent),
		Chosen:                 markAsChosen,
	}); err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "creating meal plan option")
	}

	x := &types.MealPlanOption{
		ID:                     input.ID,
		AssignedCook:           input.AssignedCook,
		Meal:                   types.Meal{ID: input.MealID},
		Notes:                  input.Notes,
		BelongsToMealPlanEvent: input.BelongsToMealPlanEvent,
		CreatedAt:              q.currentTime(),
		MealScale:              input.MealScale,
		Votes:                  []*types.MealPlanOptionVote{},
	}

	logger.Info("meal plan option created")

	return x, nil
}

// CreateMealPlanOption creates a meal plan option in the database.
func (q *Querier) CreateMealPlanOption(ctx context.Context, input *types.MealPlanOptionDatabaseCreationInput) (*types.MealPlanOption, error) {
	return q.createMealPlanOption(ctx, q.db, input, false)
}

// UpdateMealPlanOption updates a particular meal plan option.
func (q *Querier) UpdateMealPlanOption(ctx context.Context, updated *types.MealPlanOption) error {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	if updated == nil {
		return ErrNilInputProvided
	}
	logger := q.logger.WithValue(keys.MealPlanOptionIDKey, updated.ID)
	tracing.AttachMealPlanOptionIDToSpan(span, updated.ID)

	if err := q.generatedQuerier.UpdateMealPlanOption(ctx, q.db, &generated.UpdateMealPlanOptionParams{
		MealID:                 updated.Meal.ID,
		Notes:                  updated.Notes,
		MealScale:              stringFromFloat32(updated.MealScale),
		ID:                     updated.ID,
		AssignedCook:           nullStringFromStringPointer(updated.AssignedCook),
		AssignedDishwasher:     nullStringFromStringPointer(updated.AssignedDishwasher),
		BelongsToMealPlanEvent: nullStringFromString(updated.BelongsToMealPlanEvent),
	}); err != nil {
		return observability.PrepareAndLogError(err, logger, span, "updating meal plan option")
	}

	logger.Info("meal plan option updated")

	return nil
}

// ArchiveMealPlanOption archives a meal plan option from the database by its ID.
func (q *Querier) ArchiveMealPlanOption(ctx context.Context, mealPlanID, mealPlanEventID, mealPlanOptionID string) error {
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

	if err := q.generatedQuerier.ArchiveMealPlanOption(ctx, q.db, &generated.ArchiveMealPlanOptionParams{
		ID:                     mealPlanOptionID,
		BelongsToMealPlanEvent: sql.NullString{String: mealPlanEventID, Valid: true},
	}); err != nil {
		return observability.PrepareAndLogError(err, logger, span, "archiving meal plan option")
	}

	logger.Info("meal plan option archived")

	return nil
}

func (q *Querier) determineWinner(winners []schulze.Result[string]) string {
	var (
		highestScore int
		scoreWinners []string
	)

	for _, winner := range winners {
		if winner.Wins == highestScore {
			scoreWinners = append(scoreWinners, winner.Choice)
		} else if winner.Wins > highestScore {
			highestScore = winner.Wins
			scoreWinners = []string{winner.Choice}
		}
	}

	/* #nosec: G404 */
	return scoreWinners[rand.Intn(len(scoreWinners))]
}

func (q *Querier) decideOptionWinner(ctx context.Context, options []*types.MealPlanOption) (_ string, _, _ bool) {
	_, span := q.tracer.StartSpan(ctx)
	defer span.End()

	candidateMap := map[string]struct{}{}
	votesByUser := map[string]schulze.Ballot[string]{}

	logger := q.logger.WithValue("options.count", len(options))

	for _, option := range options {
		for _, v := range option.Votes {
			if votesByUser[v.ByUser] == nil {
				votesByUser[v.ByUser] = schulze.Ballot[string]{}
			}

			if !v.Abstain {
				votesByUser[v.ByUser][v.BelongsToMealPlanOption] = int(v.Rank)
			}

			candidateMap[v.BelongsToMealPlanOption] = struct{}{}
		}
	}

	candidates := []string{}
	for c := range candidateMap {
		candidates = append(candidates, c)
	}

	e := schulze.NewVoting[string](candidates)
	for _, vote := range votesByUser {
		if _, err := e.Vote(vote); err != nil {
			// this actually can never happen because we use uints for ranks, lol
			observability.AcknowledgeError(err, logger, span, "an invalid vote was received")
		}
	}

	winners, tie := e.Compute()
	if tie {
		return q.determineWinner(winners), true, true
	}

	if len(winners) > 0 {
		return winners[0].Choice, false, true
	}

	return "", false, false
}

// FinalizeMealPlanOption archives a meal plan option vote from the database by its ID.
func (q *Querier) FinalizeMealPlanOption(ctx context.Context, mealPlanID, mealPlanEventID, mealPlanOptionID, householdID string) (changed bool, err error) {
	_, span := q.tracer.StartSpan(ctx)
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

	if householdID == "" {
		return false, ErrInvalidIDProvided
	}
	logger = logger.WithValue(keys.HouseholdIDKey, householdID)
	tracing.AttachHouseholdIDToSpan(span, householdID)

	mealPlan, err := q.GetMealPlan(ctx, mealPlanID, householdID)
	if err != nil {
		return false, observability.PrepareAndLogError(err, logger, span, "fetching meal plan")
	}

	var (
		mealPlanEvent  *types.MealPlanEvent
		mealPlanOption *types.MealPlanOption
	)
	for _, event := range mealPlan.Events {
		if event.ID == mealPlanEventID {
			mealPlanEvent = event
			for _, option := range event.Options {
				if option.ID == mealPlanOptionID {
					mealPlanOption = option
					break
				}
			}
		}
	}

	// fetch household data
	household, err := q.GetHouseholdByID(ctx, mealPlan.BelongsToHousehold)
	if err != nil {
		return false, observability.PrepareAndLogError(err, logger, span, "fetching household")
	}

	// go through all the votes for this meal plan option and determine if they're all there
	for _, member := range household.Members {
		memberVoteFound := false
		for _, vote := range mealPlanOption.Votes {
			if vote.ByUser == member.BelongsToUser.ID {
				memberVoteFound = true
				break
			}
		}

		if !memberVoteFound {
			return false, nil
		}
	}

	winner, tiebroken, chosen := q.decideOptionWinner(ctx, mealPlanEvent.Options)
	if chosen {
		if err = q.generatedQuerier.FinalizeMealPlanOption(ctx, q.db, &generated.FinalizeMealPlanOptionParams{
			BelongsToMealPlanEvent: nullStringFromString(mealPlanEventID),
			ID:                     winner,
			Tiebroken:              tiebroken,
		}); err != nil {
			return false, observability.PrepareAndLogError(err, logger, span, "finalizing meal plan option")
		}

		logger.Debug("finalized meal plan option")
	}

	return chosen, nil
}
