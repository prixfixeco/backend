package postgres

import (
	"context"
	"database/sql"
	_ "embed"
	"log"

	"github.com/dinnerdonebetter/backend/internal/database"
	"github.com/dinnerdonebetter/backend/internal/database/postgres/generated"
	"github.com/dinnerdonebetter/backend/internal/identifiers"
	"github.com/dinnerdonebetter/backend/internal/observability"
	"github.com/dinnerdonebetter/backend/internal/observability/keys"
	"github.com/dinnerdonebetter/backend/internal/observability/tracing"
	"github.com/dinnerdonebetter/backend/pkg/types"

	"github.com/Masterminds/squirrel"
)

var (
	_ types.RecipeDataManager = (*Querier)(nil)

	// recipesTableColumns are the columns for the recipes table.
	recipesTableColumns = []string{
		"recipes.id",
		"recipes.name",
		"recipes.slug",
		"recipes.source",
		"recipes.description",
		"recipes.inspired_by_recipe_id",
		"recipes.min_estimated_portions",
		"recipes.max_estimated_portions",
		"recipes.portion_name",
		"recipes.plural_portion_name",
		"recipes.seal_of_approval",
		"recipes.eligible_for_meals",
		"recipes.yields_component_type",
		"recipes.created_at",
		"recipes.last_updated_at",
		"recipes.archived_at",
		"recipes.created_by_user",
	}
)

// scanRecipe takes a database Scanner (i.e. *sql.Row) and scans the result into a recipe struct.
func (q *Querier) scanRecipe(ctx context.Context, scan database.Scanner, includeCounts bool) (x *types.Recipe, filteredCount, totalCount uint64, err error) {
	_, span := q.tracer.StartSpan(ctx)
	defer span.End()

	x = &types.Recipe{}

	targetVars := []any{
		&x.ID,
		&x.Name,
		&x.Slug,
		&x.Source,
		&x.Description,
		&x.InspiredByRecipeID,
		&x.MinimumEstimatedPortions,
		&x.MaximumEstimatedPortions,
		&x.PortionName,
		&x.PluralPortionName,
		&x.SealOfApproval,
		&x.EligibleForMeals,
		&x.YieldsComponentType,
		&x.CreatedAt,
		&x.LastUpdatedAt,
		&x.ArchivedAt,
		&x.CreatedByUser,
	}

	if includeCounts {
		targetVars = append(targetVars, &filteredCount, &totalCount)
	}

	if err = scan.Scan(targetVars...); err != nil {
		return nil, 0, 0, observability.PrepareError(err, span, "")
	}

	return x, filteredCount, totalCount, nil
}

// scanRecipes takes some database rows and turns them into a slice of recipes.
func (q *Querier) scanRecipes(ctx context.Context, rows database.ResultIterator, includeCounts bool) (recipes []*types.Recipe, filteredCount, totalCount uint64, err error) {
	_, span := q.tracer.StartSpan(ctx)
	defer span.End()

	for rows.Next() {
		x, fc, tc, scanErr := q.scanRecipe(ctx, rows, includeCounts)
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

		recipes = append(recipes, x)
	}

	if err = q.checkRowsForErrorAndClose(ctx, rows); err != nil {
		return nil, 0, 0, observability.PrepareError(err, span, "handling rows")
	}

	return recipes, filteredCount, totalCount, nil
}

//go:embed queries/recipes/exists.sql
var recipeExistenceQuery string

// RecipeExists fetches whether a recipe exists from the database.
func (q *Querier) RecipeExists(ctx context.Context, recipeID string) (exists bool, err error) {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	if recipeID == "" {
		return false, ErrInvalidIDProvided
	}
	tracing.AttachRecipeIDToSpan(span, recipeID)

	args := []any{
		recipeID,
	}

	result, err := q.performBooleanQuery(ctx, q.db, recipeExistenceQuery, args)
	if err != nil {
		return false, observability.PrepareError(err, span, "performing recipe existence check")
	}

	return result, nil
}

//go:embed queries/recipes/get_by_id.sql
var getRecipeByIDQuery string

//go:embed queries/recipes/get_by_id_and_author_id.sql
var getRecipeByIDAndAuthorIDQuery string

// scanRecipeAndStep takes a database Scanner (i.e. *sql.Row) and scans the result into a recipe struct.
func (q *Querier) scanRecipeAndStep(ctx context.Context, scan database.Scanner) (x *types.Recipe, y *types.RecipeStep, filteredCount, totalCount uint64, err error) {
	_, span := q.tracer.StartSpan(ctx)
	defer span.End()

	x = &types.Recipe{}
	y = &types.RecipeStep{}

	targetVars := []any{
		&x.ID,
		&x.Name,
		&x.Slug,
		&x.Source,
		&x.Description,
		&x.InspiredByRecipeID,
		&x.MinimumEstimatedPortions,
		&x.MaximumEstimatedPortions,
		&x.PortionName,
		&x.PluralPortionName,
		&x.SealOfApproval,
		&x.EligibleForMeals,
		&x.YieldsComponentType,
		&x.CreatedAt,
		&x.LastUpdatedAt,
		&x.ArchivedAt,
		&x.CreatedByUser,
		&y.ID,
		&y.Index,
		&y.Preparation.ID,
		&y.Preparation.Name,
		&y.Preparation.Description,
		&y.Preparation.IconPath,
		&y.Preparation.YieldsNothing,
		&y.Preparation.RestrictToIngredients,
		&y.Preparation.MinimumIngredientCount,
		&y.Preparation.MaximumIngredientCount,
		&y.Preparation.MinimumInstrumentCount,
		&y.Preparation.MaximumInstrumentCount,
		&y.Preparation.TemperatureRequired,
		&y.Preparation.TimeEstimateRequired,
		&y.Preparation.ConditionExpressionRequired,
		&y.Preparation.ConsumesVessel,
		&y.Preparation.OnlyForVessels,
		&y.Preparation.MinimumVesselCount,
		&y.Preparation.MaximumVesselCount,
		&y.Preparation.Slug,
		&y.Preparation.PastTense,
		&y.Preparation.CreatedAt,
		&y.Preparation.LastUpdatedAt,
		&y.Preparation.ArchivedAt,
		&y.MinimumEstimatedTimeInSeconds,
		&y.MaximumEstimatedTimeInSeconds,
		&y.MinimumTemperatureInCelsius,
		&y.MaximumTemperatureInCelsius,
		&y.Notes,
		&y.ExplicitInstructions,
		&y.ConditionExpression,
		&y.Optional,
		&y.StartTimerAutomatically,
		&y.CreatedAt,
		&y.LastUpdatedAt,
		&y.ArchivedAt,
		&y.BelongsToRecipe,
	}

	if err = scan.Scan(targetVars...); err != nil {
		return nil, nil, 0, 0, observability.PrepareError(err, span, "")
	}

	return x, y, filteredCount, totalCount, nil
}

// getRecipe fetches a recipe from the database.
func (q *Querier) getRecipe(ctx context.Context, recipeID, userID string) (*types.Recipe, error) {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	if recipeID == "" {
		return nil, ErrInvalidIDProvided
	}
	tracing.AttachRecipeIDToSpan(span, recipeID)

	args := []any{
		recipeID,
	}

	query := getRecipeByIDQuery
	if userID != "" {
		query = getRecipeByIDAndAuthorIDQuery
		args = append(args, userID)
	}

	rows, err := q.getRows(ctx, q.db, "get recipe", query, args)
	if err != nil {
		return nil, observability.PrepareError(err, span, "fetching recipe data")
	}

	var x *types.Recipe
	for rows.Next() {
		recipe, recipeStep, _, _, recipeScanErr := q.scanRecipeAndStep(ctx, rows)
		if recipeScanErr != nil {
			return nil, observability.PrepareError(recipeScanErr, span, "scanning recipe")
		}

		if x == nil {
			x = recipe
		}

		x.Steps = append(x.Steps, recipeStep)
	}

	if x == nil {
		return nil, sql.ErrNoRows
	}

	prepTasks, err := q.getRecipePrepTasksForRecipe(ctx, recipeID)
	if err != nil {
		return nil, observability.PrepareError(err, span, "fetching recipe step prep tasks for recipe")
	}
	x.PrepTasks = prepTasks

	recipeMedia, err := q.getRecipeMediaForRecipe(ctx, recipeID)
	if err != nil {
		return nil, observability.PrepareError(err, span, "fetching recipe step ingredients for recipe")
	}
	x.Media = recipeMedia

	ingredients, err := q.getRecipeStepIngredientsForRecipe(ctx, recipeID)
	if err != nil {
		return nil, observability.PrepareError(err, span, "fetching recipe step ingredients for recipe")
	}

	products, err := q.getRecipeStepProductsForRecipe(ctx, recipeID)
	if err != nil {
		return nil, observability.PrepareError(err, span, "fetching recipe step products for recipe")
	}

	instruments, err := q.getRecipeStepInstrumentsForRecipe(ctx, recipeID)
	if err != nil {
		return nil, observability.PrepareError(err, span, "fetching recipe step instruments for recipe")
	}

	vessels, err := q.getRecipeStepVesselsForRecipe(ctx, recipeID)
	if err != nil {
		return nil, observability.PrepareError(err, span, "fetching recipe step vessels for recipe")
	}

	completionConditions, err := q.getRecipeStepCompletionConditionsForRecipe(ctx, recipeID)
	if err != nil {
		return nil, observability.PrepareError(err, span, "fetching recipe step completion conditions for recipe")
	}

	supportingRecipeInfo := map[string]bool{}
	for i, step := range x.Steps {
		for _, ingredient := range ingredients {
			if ingredient.RecipeStepProductRecipeID != nil {
				supportingRecipeInfo[*ingredient.RecipeStepProductRecipeID] = true
			}

			if ingredient.BelongsToRecipeStep == step.ID {
				x.Steps[i].Ingredients = append(x.Steps[i].Ingredients, ingredient)
			}
		}

		for _, product := range products {
			if product.BelongsToRecipeStep == step.ID {
				x.Steps[i].Products = append(x.Steps[i].Products, product)
			}
		}

		for _, instrument := range instruments {
			if instrument.BelongsToRecipeStep == step.ID {
				x.Steps[i].Instruments = append(x.Steps[i].Instruments, instrument)
			}
		}

		for _, vessel := range vessels {
			if vessel.BelongsToRecipeStep == step.ID {
				x.Steps[i].Vessels = append(x.Steps[i].Vessels, vessel)
			}
		}

		for _, completionCondition := range completionConditions {
			if completionCondition.BelongsToRecipeStep == step.ID {
				x.Steps[i].CompletionConditions = append(x.Steps[i].CompletionConditions, completionCondition)
			}
		}

		recipeMedia, err = q.getRecipeMediaForRecipeStep(ctx, recipeID, step.ID)
		if err != nil {
			return nil, observability.PrepareError(err, span, "fetching recipe step ingredients for recipe")
		}
		x.Steps[i].Media = recipeMedia
	}

	var supportingRecipeIDs []string
	for supportingRecipe := range supportingRecipeInfo {
		supportingRecipeIDs = append(supportingRecipeIDs, supportingRecipe)
	}

	x.SupportingRecipes, err = q.GetRecipesWithIDs(ctx, supportingRecipeIDs)
	if err != nil {
		return nil, observability.PrepareError(err, span, "fetching supporting recipes")
	}

	return x, nil
}

// GetRecipe fetches a recipe from the database.
func (q *Querier) GetRecipe(ctx context.Context, recipeID string) (*types.Recipe, error) {
	return q.getRecipe(ctx, recipeID, "")
}

// GetRecipeByIDAndUser fetches a recipe from the database.
func (q *Querier) GetRecipeByIDAndUser(ctx context.Context, recipeID, userID string) (*types.Recipe, error) {
	return q.getRecipe(ctx, recipeID, userID)
}

// GetRecipes fetches a list of recipes from the database that meet a particular filter.
func (q *Querier) GetRecipes(ctx context.Context, filter *types.QueryFilter) (x *types.QueryFilteredResult[types.Recipe], err error) {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	x = &types.QueryFilteredResult[types.Recipe]{}
	tracing.AttachQueryFilterToSpan(span, filter)

	if filter != nil {
		if filter.Page != nil {
			x.Page = *filter.Page
		}

		if filter.Limit != nil {
			x.Limit = *filter.Limit
		}
	}

	query, args := q.buildListQuery(ctx, "recipes", nil, nil, nil, "", recipesTableColumns, "", false, filter)

	rows, err := q.getRows(ctx, q.db, "recipes", query, args)
	if err != nil {
		return nil, observability.PrepareError(err, span, "executing recipes list retrieval query")
	}

	if x.Data, x.FilteredCount, x.TotalCount, err = q.scanRecipes(ctx, rows, true); err != nil {
		return nil, observability.PrepareError(err, span, "scanning recipes")
	}

	return x, nil
}

// GetRecipesWithIDs fetches a list of recipes from the database that meet a particular filter.
func (q *Querier) GetRecipesWithIDs(ctx context.Context, ids []string) ([]*types.Recipe, error) {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	logger := q.logger.Clone()

	recipes := []*types.Recipe{}
	for _, id := range ids {
		r, err := q.getRecipe(ctx, id, "")
		if err != nil {
			return nil, observability.PrepareAndLogError(err, logger, span, "getting recipe")
		}

		recipes = append(recipes, r)
	}

	return recipes, nil
}

//go:embed queries/recipes/get_needing_indexing.sql
var recipesNeedingIndexingQuery string

// GetRecipeIDsThatNeedSearchIndexing fetches a list of recipe IDs from the database that meet a particular filter.
func (q *Querier) GetRecipeIDsThatNeedSearchIndexing(ctx context.Context) ([]string, error) {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	rows, err := q.getRows(ctx, q.db, "recipes needing indexing", recipesNeedingIndexingQuery, nil)
	if err != nil {
		return nil, observability.PrepareError(err, span, "executing recipes list retrieval query")
	}

	return q.scanIDs(ctx, rows)
}

//go:embed queries/recipes/ids_for_meal.sql
var getRecipesForMealQuery string

// getRecipeIDsForMeal fetches a list of recipe IDs from the database that are associated with a given meal.
func (q *Querier) getRecipeIDsForMeal(ctx context.Context, mealID string) (x []string, err error) {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	if mealID == "" {
		return nil, ErrInvalidIDProvided
	}
	tracing.AttachMealIDToSpan(span, mealID)

	args := []any{
		mealID,
	}

	rows, err := q.getRows(ctx, q.db, "recipes for meal", getRecipesForMealQuery, args)
	if err != nil {
		return nil, observability.PrepareError(err, span, "executing recipes list retrieval query")
	}

	if x, err = q.scanIDs(ctx, rows); err != nil {
		return nil, observability.PrepareError(err, span, "scanning recipes")
	}

	return x, nil
}

// SearchForRecipes fetches a list of recipes from the database that match a query.
func (q *Querier) SearchForRecipes(ctx context.Context, recipeNameQuery string, filter *types.QueryFilter) (x *types.QueryFilteredResult[types.Recipe], err error) {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	x = &types.QueryFilteredResult[types.Recipe]{}
	tracing.AttachQueryFilterToSpan(span, filter)

	if filter != nil {
		if filter.Page != nil {
			x.Page = *filter.Page
		}

		if filter.Limit != nil {
			x.Limit = *filter.Limit
		}
	}

	where := squirrel.ILike{"name": wrapQueryForILIKE(recipeNameQuery)}
	query, args := q.buildListQueryWithILike(ctx, "recipes", nil, nil, where, "", recipesTableColumns, "", false, filter)

	rows, err := q.getRows(ctx, q.db, "recipes search", query, args)
	if err != nil {
		return nil, observability.PrepareError(err, span, "executing recipes search query")
	}

	if x.Data, x.FilteredCount, x.TotalCount, err = q.scanRecipes(ctx, rows, true); err != nil {
		return nil, observability.PrepareError(err, span, "scanning recipes")
	}

	return x, nil
}

//go:embed queries/recipes/create.sql
var recipeCreationQuery string

// CreateRecipe creates a recipe in the database.
func (q *Querier) CreateRecipe(ctx context.Context, input *types.RecipeDatabaseCreationInput) (*types.Recipe, error) {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	if input == nil {
		return nil, ErrNilInputProvided
	}

	logger := q.logger.WithValue(keys.RecipeIDKey, input.ID)

	tx, err := q.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "beginning transaction")
	}

	args := []any{
		input.ID,
		input.Name,
		input.Slug,
		input.Source,
		input.Description,
		input.InspiredByRecipeID,
		input.MinimumEstimatedPortions,
		input.MaximumEstimatedPortions,
		input.PortionName,
		input.PluralPortionName,
		input.SealOfApproval,
		input.EligibleForMeals,
		input.YieldsComponentType,
		input.CreatedByUser,
	}

	// create the recipe.
	if err = q.performWriteQuery(ctx, q.db, "recipe creation", recipeCreationQuery, args); err != nil {
		q.rollbackTransaction(ctx, tx)
		return nil, observability.PrepareAndLogError(err, logger, span, "performing recipe creation query")
	}

	x := &types.Recipe{
		ID:                       input.ID,
		Name:                     input.Name,
		Slug:                     input.Slug,
		Source:                   input.Source,
		Description:              input.Description,
		InspiredByRecipeID:       input.InspiredByRecipeID,
		CreatedByUser:            input.CreatedByUser,
		MinimumEstimatedPortions: input.MinimumEstimatedPortions,
		MaximumEstimatedPortions: input.MaximumEstimatedPortions,
		SealOfApproval:           input.SealOfApproval,
		EligibleForMeals:         input.EligibleForMeals,
		PortionName:              input.PortionName,
		PluralPortionName:        input.PluralPortionName,
		YieldsComponentType:      input.YieldsComponentType,
		CreatedAt:                q.currentTime(),
	}

	findCreatedRecipeStepProductsForIngredients(input)
	findCreatedRecipeStepProductsForInstruments(input)
	findCreatedRecipeStepProductsForVessels(input)

	for i, stepInput := range input.Steps {
		stepInput.Index = uint32(i)
		stepInput.BelongsToRecipe = x.ID

		s, createErr := q.createRecipeStep(ctx, tx, stepInput)
		if createErr != nil {
			q.rollbackTransaction(ctx, tx)
			return nil, observability.PrepareError(createErr, span, "creating recipe step #%d", i+1)
		}

		x.Steps = append(x.Steps, s)
	}

	for i, prepTaskInput := range input.PrepTasks {
		pt, createPrepTaskErr := q.createRecipePrepTask(ctx, tx, prepTaskInput)
		if createPrepTaskErr != nil {
			q.rollbackTransaction(ctx, tx)
			return nil, observability.PrepareError(createPrepTaskErr, span, "creating recipe prep task #%d", i+1)
		}

		x.PrepTasks = append(x.PrepTasks, pt)
	}

	if input.AlsoCreateMeal {
		_, mealCreateErr := q.createMeal(ctx, tx, &types.MealDatabaseCreationInput{
			ID:                       identifiers.New(),
			Name:                     x.Name,
			Description:              x.Description,
			MinimumEstimatedPortions: x.MinimumEstimatedPortions,
			MaximumEstimatedPortions: x.MaximumEstimatedPortions,
			EligibleForMealPlans:     x.EligibleForMeals,
			CreatedByUser:            x.CreatedByUser,
			Components: []*types.MealComponentDatabaseCreationInput{
				{
					RecipeID:      x.ID,
					RecipeScale:   1.0,
					ComponentType: types.MealComponentTypesMain,
				},
			},
		})

		if mealCreateErr != nil {
			q.rollbackTransaction(ctx, tx)
			return nil, observability.PrepareError(mealCreateErr, span, "creating meal from recipe")
		}
	}

	if err = tx.Commit(); err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "committing transaction")
	}

	tracing.AttachRecipeIDToSpan(span, x.ID)
	logger.Info("recipe created")

	return x, nil
}

func findCreatedRecipeStepProductsForIngredients(recipe *types.RecipeDatabaseCreationInput) {
	for _, step := range recipe.Steps {
		for _, ingredient := range step.Ingredients {
			if ingredient.ProductOfRecipeStepIndex != nil && ingredient.ProductOfRecipeStepProductIndex != nil {
				enoughSteps := len(recipe.Steps) > int(*ingredient.ProductOfRecipeStepIndex)
				enoughRecipeStepProducts := len(recipe.Steps[int(*ingredient.ProductOfRecipeStepIndex)].Products) > int(*ingredient.ProductOfRecipeStepProductIndex)
				relevantProductIsIngredient := recipe.Steps[*ingredient.ProductOfRecipeStepIndex].Products[*ingredient.ProductOfRecipeStepProductIndex].Type == types.RecipeStepProductIngredientType
				if enoughSteps && enoughRecipeStepProducts && relevantProductIsIngredient {
					ingredient.RecipeStepProductID = &recipe.Steps[*ingredient.ProductOfRecipeStepIndex].Products[*ingredient.ProductOfRecipeStepProductIndex].ID
				}
			}
		}
	}
}

func findCreatedRecipeStepProductsForInstruments(recipe *types.RecipeDatabaseCreationInput) {
	for _, step := range recipe.Steps {
		for _, instrument := range step.Instruments {
			if instrument.ProductOfRecipeStepIndex != nil && instrument.ProductOfRecipeStepProductIndex != nil {
				enoughSteps := len(recipe.Steps) > int(*instrument.ProductOfRecipeStepIndex)
				enoughRecipeStepProducts := len(recipe.Steps[int(*instrument.ProductOfRecipeStepIndex)].Products) > int(*instrument.ProductOfRecipeStepProductIndex)
				relevantProductIsInstrument := recipe.Steps[*instrument.ProductOfRecipeStepIndex].Products[*instrument.ProductOfRecipeStepProductIndex].Type == types.RecipeStepProductInstrumentType
				if enoughSteps && enoughRecipeStepProducts && relevantProductIsInstrument {
					instrument.RecipeStepProductID = &recipe.Steps[*instrument.ProductOfRecipeStepIndex].Products[*instrument.ProductOfRecipeStepProductIndex].ID
				}
			}
		}
	}
}

func findCreatedRecipeStepProductsForVessels(recipe *types.RecipeDatabaseCreationInput) {
	for _, step := range recipe.Steps {
		for _, vessel := range step.Vessels {
			if vessel.ProductOfRecipeStepIndex != nil && vessel.ProductOfRecipeStepProductIndex != nil {
				enoughSteps := len(recipe.Steps) > int(*vessel.ProductOfRecipeStepIndex)
				enoughRecipeStepProducts := len(recipe.Steps[int(*vessel.ProductOfRecipeStepIndex)].Products) > int(*vessel.ProductOfRecipeStepProductIndex)
				relevantProductIsVessel := recipe.Steps[*vessel.ProductOfRecipeStepIndex].Products[*vessel.ProductOfRecipeStepProductIndex].Type == types.RecipeStepProductVesselType
				if enoughSteps && enoughRecipeStepProducts && relevantProductIsVessel {
					vessel.RecipeStepProductID = &recipe.Steps[*vessel.ProductOfRecipeStepIndex].Products[*vessel.ProductOfRecipeStepProductIndex].ID
				} else {
					log.Printf("for recipe step id %q, vessel ID %q, not enough steps: %t, not enough recipe step products: %t, relevant product is vessel: %t", step.ID, vessel.ID, enoughSteps, enoughRecipeStepProducts, relevantProductIsVessel)
				}
			}
		}
	}
}

//go:embed queries/recipes/update.sql
var updateRecipeQuery string

// UpdateRecipe updates a particular recipe.
func (q *Querier) UpdateRecipe(ctx context.Context, updated *types.Recipe) error {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	if updated == nil {
		return ErrNilInputProvided
	}

	logger := q.logger.WithValue(keys.RecipeIDKey, updated.ID)
	tracing.AttachRecipeIDToSpan(span, updated.ID)
	tracing.AttachUserIDToSpan(span, updated.CreatedByUser)

	args := []any{
		updated.Name,
		updated.Slug,
		updated.Source,
		updated.Description,
		updated.InspiredByRecipeID,
		updated.MinimumEstimatedPortions,
		updated.MaximumEstimatedPortions,
		updated.PortionName,
		updated.PluralPortionName,
		updated.SealOfApproval,
		updated.EligibleForMeals,
		updated.YieldsComponentType,
		updated.CreatedByUser,
		updated.ID,
	}

	if err := q.performWriteQuery(ctx, q.db, "recipe update", updateRecipeQuery, args); err != nil {
		return observability.PrepareAndLogError(err, logger, span, "updating recipe")
	}

	logger.Info("recipe updated")

	return nil
}

//go:embed queries/recipes/update_last_indexed_at.sql
var updateRecipeLastIndexedAtQuery string

// MarkRecipeAsIndexed updates a particular recipe's last_indexed_at value.
func (q *Querier) MarkRecipeAsIndexed(ctx context.Context, recipeID string) error {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	logger := q.logger.Clone()

	if recipeID == "" {
		return ErrInvalidIDProvided
	}
	logger = logger.WithValue(keys.RecipeIDKey, recipeID)
	tracing.AttachRecipeIDToSpan(span, recipeID)

	args := []any{
		recipeID,
	}

	if err := q.performWriteQuery(ctx, q.db, "recipe last_indexed_at", updateRecipeLastIndexedAtQuery, args); err != nil {
		return observability.PrepareAndLogError(err, logger, span, "marking recipe as indexed")
	}

	logger.Info("recipe marked as indexed")

	return nil
}

// ArchiveRecipe archives a recipe from the database by its ID.
func (q *Querier) ArchiveRecipe(ctx context.Context, recipeID, userID string) error {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	if recipeID == "" {
		return ErrInvalidIDProvided
	}
	logger := q.logger.WithValue(keys.RecipeIDKey, recipeID)
	tracing.AttachRecipeIDToSpan(span, recipeID)

	if userID == "" {
		return ErrInvalidIDProvided
	}
	logger = logger.WithValue(keys.UserIDKey, userID)
	tracing.AttachUserIDToSpan(span, userID)

	if err := q.generatedQuerier.ArchiveRecipe(ctx, q.db, &generated.ArchiveRecipeParams{
		CreatedByUser: userID,
		ID:            recipeID,
	}); err != nil {
		return observability.PrepareAndLogError(err, logger, span, "archiving recipe")
	}

	logger.Info("recipe archived")

	return nil
}
