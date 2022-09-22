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

var (
	_ types.AdvancedPrepStepDataManager = (*Querier)(nil)

	// advancedPrepStepsTableColumns are the columns for the valid_instruments table.
	advancedPrepStepsTableColumns = []string{
		"advanced_prep_steps.id",
		"advanced_prep_steps.belongs_to_meal_plan_option",
		"advanced_prep_steps.satisfies_recipe_step",
		"advanced_prep_steps.status",
		"advanced_prep_steps.status_explanation",
		"advanced_prep_steps.creation_explanation",
		"advanced_prep_steps.cannot_complete_before",
		"advanced_prep_steps.cannot_complete_after",
		"advanced_prep_steps.created_at",
		"advanced_prep_steps.settled_at",
	}
)

// scanAdvancedPrepStep takes a database Scanner (i.e. *sql.Row) and scans the result into a valid instrument struct.
func (q *Querier) scanAdvancedPrepStep(ctx context.Context, scan database.Scanner, includeCounts bool) (x *types.AdvancedPrepStep, filteredCount, totalCount uint64, err error) {
	_, span := q.tracer.StartSpan(ctx)
	defer span.End()

	x = &types.AdvancedPrepStep{}

	targetVars := []interface{}{
		&x.ID,
		&x.MealPlanOption.ID,
		&x.MealPlanOption.AssignedCook,
		&x.MealPlanOption.AssignedDishwasher,
		&x.MealPlanOption.Chosen,
		&x.MealPlanOption.TieBroken,
		&x.MealPlanOption.Meal.ID,
		&x.MealPlanOption.Notes,
		&x.MealPlanOption.PrepStepsCreated,
		&x.MealPlanOption.CreatedAt,
		&x.MealPlanOption.LastUpdatedAt,
		&x.MealPlanOption.ArchivedAt,
		&x.MealPlanOption.BelongsToMealPlanEvent,
		&x.RecipeStep.ID,
		&x.RecipeStep.Index,
		&x.RecipeStep.Preparation.ID,
		&x.RecipeStep.Preparation.Name,
		&x.RecipeStep.Preparation.Description,
		&x.RecipeStep.Preparation.IconPath,
		&x.RecipeStep.Preparation.YieldsNothing,
		&x.RecipeStep.Preparation.RestrictToIngredients,
		&x.RecipeStep.Preparation.ZeroIngredientsAllowable,
		&x.RecipeStep.Preparation.PastTense,
		&x.RecipeStep.Preparation.CreatedAt,
		&x.RecipeStep.Preparation.LastUpdatedAt,
		&x.RecipeStep.Preparation.ArchivedAt,
		&x.RecipeStep.MinimumEstimatedTimeInSeconds,
		&x.RecipeStep.MaximumEstimatedTimeInSeconds,
		&x.RecipeStep.MinimumTemperatureInCelsius,
		&x.RecipeStep.MaximumTemperatureInCelsius,
		&x.RecipeStep.Notes,
		&x.RecipeStep.ExplicitInstructions,
		&x.RecipeStep.Optional,
		&x.RecipeStep.CreatedAt,
		&x.RecipeStep.LastUpdatedAt,
		&x.RecipeStep.ArchivedAt,
		&x.RecipeStep.BelongsToRecipe,
		&x.Status,
		&x.StatusExplanation,
		&x.CreationExplanation,
		&x.CannotCompleteBefore,
		&x.CannotCompleteAfter,
		&x.CreatedAt,
		&x.SettledAt,
	}

	if includeCounts {
		targetVars = append(targetVars, &filteredCount, &totalCount)
	}

	if err = scan.Scan(targetVars...); err != nil {
		return nil, 0, 0, observability.PrepareError(err, span, "")
	}

	return x, filteredCount, totalCount, nil
}

// scanAdvancedPrepSteps takes some database rows and turns them into a slice of advanced prep steps.
func (q *Querier) scanAdvancedPrepSteps(ctx context.Context, rows database.ResultIterator, includeCounts bool) (validInstruments []*types.AdvancedPrepStep, filteredCount, totalCount uint64, err error) {
	_, span := q.tracer.StartSpan(ctx)
	defer span.End()

	for rows.Next() {
		x, fc, tc, scanErr := q.scanAdvancedPrepStep(ctx, rows, includeCounts)
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

		validInstruments = append(validInstruments, x)
	}

	if err = q.checkRowsForErrorAndClose(ctx, rows); err != nil {
		return nil, 0, 0, observability.PrepareError(err, span, "handling rows")
	}

	return validInstruments, filteredCount, totalCount, nil
}

//go:embed queries/advanced_prep_steps/get_one.sql
var getAdvancedPrepStepsQuery string

// GetAdvancedPrepStep fetches an advanced prep step.
func (q *Querier) GetAdvancedPrepStep(ctx context.Context, advancedPrepStepID string) (x *types.AdvancedPrepStep, err error) {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	logger := q.logger.Clone()

	if advancedPrepStepID == "" {
		return nil, ErrInvalidIDProvided
	}
	logger = logger.WithValue(keys.AdvancedPrepStepIDKey, advancedPrepStepID)
	tracing.AttachAdvancedPrepStepIDToSpan(span, advancedPrepStepID)

	args := []interface{}{
		advancedPrepStepID,
	}

	rows, err := q.performReadQuery(ctx, q.db, "advanced prep step", getAdvancedPrepStepsQuery, args)
	if err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "executing advanced prep step retrieval query")
	}

	if x, _, _, err = q.scanAdvancedPrepStep(ctx, rows, true); err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "scanning advanced prep step")
	}

	logger.Info("advanced steps retrieved")

	return x, nil
}

// GetAdvancedPrepSteps fetches a list of advanced prep steps.
func (q *Querier) GetAdvancedPrepSteps(ctx context.Context, filter *types.QueryFilter) (x *types.AdvancedPrepStepList, err error) {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	logger := q.logger.Clone()

	x = &types.AdvancedPrepStepList{}
	logger = filter.AttachToLogger(logger)
	tracing.AttachQueryFilterToSpan(span, filter)

	if filter != nil {
		if filter.Page != nil {
			x.Page = *filter.Page
		}

		if filter.Limit != nil {
			x.Limit = *filter.Limit
		}
	} else {
		filter = types.DefaultQueryFilter()
	}

	query, args := q.buildListQuery(ctx, "advanced_prep_steps", nil, nil, nil, "", advancedPrepStepsTableColumns, "", false, filter)

	rows, err := q.performReadQuery(ctx, q.db, "advanced prep steps", query, args)
	if err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "executing advanced prep steps list retrieval query")
	}

	if x.AdvancedPrepSteps, x.FilteredCount, x.TotalCount, err = q.scanAdvancedPrepSteps(ctx, rows, true); err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "scanning advanced prep steps")
	}

	return x, nil
}

//go:embed queries/advanced_prep_steps/list_all_by_meal_plan_option.sql
var listAdvancedPrepStepsQuery string

// GetAdvancedPrepStepsForMealPlanOptionID lists advanced prep steps for a given meal plan option.
func (q *Querier) GetAdvancedPrepStepsForMealPlanOptionID(ctx context.Context, mealPlanOptionID string, filter *types.QueryFilter) (x *types.AdvancedPrepStepList, err error) {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	x = &types.AdvancedPrepStepList{}
	logger := filter.AttachToLogger(q.logger.Clone())
	tracing.AttachQueryFilterToSpan(span, filter)

	if mealPlanOptionID == "" {
		return nil, ErrInvalidIDProvided
	}
	logger = logger.WithValue(keys.MealPlanOptionIDKey, mealPlanOptionID)
	tracing.AttachMealPlanOptionIDToSpan(span, mealPlanOptionID)

	args := []interface{}{
		mealPlanOptionID,
	}

	rows, err := q.performReadQuery(ctx, q.db, "advanced prep steps list", listAdvancedPrepStepsQuery, args)
	if err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "executing advanced prep steps list retrieval query")
	}

	if x.AdvancedPrepSteps, x.FilteredCount, x.TotalCount, err = q.scanAdvancedPrepSteps(ctx, rows, true); err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "scanning advanced prep steps")
	}

	logger.Info("advanced steps retrieved")

	return x, nil
}

//go:embed queries/advanced_prep_steps/create.sql
var createAdvancedPrepStepQuery string

//go:embed queries/meal_plan_options/mark_as_steps_created.sql
var markMealPlanOptionAsHavingStepsCreatedQuery string

// CreateAdvancedPrepStep creates advanced prep steps.
func (q *Querier) CreateAdvancedPrepStep(ctx context.Context, input *types.AdvancedPrepStepDatabaseCreationInput) (*types.AdvancedPrepStep, error) {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	logger := q.logger.Clone()

	tx, err := q.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "beginning transaction")
	}

	createAdvancedPrepStepArgs := []interface{}{
		input.ID,
		input.MealPlanOptionID,
		input.RecipeStepID,
		input.Status,
		input.StatusExplanation,
		input.CreationExplanation,
		input.CannotCompleteBefore,
		input.CannotCompleteAfter,
	}

	if err = q.performWriteQuery(ctx, tx, "create advanced prep step", createAdvancedPrepStepQuery, createAdvancedPrepStepArgs); err != nil {
		q.rollbackTransaction(ctx, tx)
		return nil, observability.PrepareAndLogError(err, logger, span, "create advanced prep step")
	}

	// mark prep steps as created for step
	markMealPlanOptionAsHavingStepsCreatedArgs := []interface{}{
		input.MealPlanOptionID,
	}

	if err = q.performWriteQuery(ctx, tx, "create advanced prep step", markMealPlanOptionAsHavingStepsCreatedQuery, markMealPlanOptionAsHavingStepsCreatedArgs); err != nil {
		q.rollbackTransaction(ctx, tx)
		return nil, observability.PrepareAndLogError(err, logger, span, "create advanced prep step")
	}

	if commitErr := tx.Commit(); commitErr != nil {
		return nil, observability.PrepareAndLogError(commitErr, logger, span, "committing transaction")
	}

	x := &types.AdvancedPrepStep{
		ID:                   input.ID,
		CannotCompleteBefore: input.CannotCompleteBefore,
		CannotCompleteAfter:  input.CannotCompleteAfter,
		MealPlanOption:       types.MealPlanOption{ID: input.MealPlanOptionID},
		RecipeStep:           types.RecipeStep{ID: input.RecipeStepID},
		CreatedAt:            q.currentTime(),
		Status:               input.Status,
		StatusExplanation:    input.StatusExplanation,
		CreationExplanation:  input.CreationExplanation,
		SettledAt:            input.CompletedAt,
	}

	logger.Info("advanced step created")

	return x, nil
}

//go:embed queries/advanced_prep_steps/complete.sql
var completeAdvancedPrepStepQuery string

// MarkAdvancedPrepStepAsComplete marks an advanced prep step as complete.
func (q *Querier) MarkAdvancedPrepStepAsComplete(ctx context.Context, advancedPrepStepID string) error {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	logger := q.logger.Clone()

	if advancedPrepStepID == "" {
		return ErrInvalidIDProvided
	}
	tracing.AttachAdvancedPrepStepIDToSpan(span, advancedPrepStepID)
	logger = logger.WithValue(keys.AdvancedPrepStepIDKey, advancedPrepStepID)

	args := []interface{}{
		advancedPrepStepID,
	}

	if err := q.performWriteQuery(ctx, q.db, "prep step complete", completeAdvancedPrepStepQuery, args); err != nil {
		return observability.PrepareAndLogError(err, logger, span, "marking prep step as complete")
	}

	logger.Info("prep step marked as complete")

	return nil
}