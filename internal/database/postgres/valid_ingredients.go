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
	_ types.ValidIngredientDataManager = (*Querier)(nil)

	// validIngredientsTableColumns are the columns for the valid_ingredients table.
	validIngredientsTableColumns = []string{
		"valid_ingredients.id",
		"valid_ingredients.name",
		"valid_ingredients.description",
		"valid_ingredients.warning",
		"valid_ingredients.contains_egg",
		"valid_ingredients.contains_dairy",
		"valid_ingredients.contains_peanut",
		"valid_ingredients.contains_tree_nut",
		"valid_ingredients.contains_soy",
		"valid_ingredients.contains_wheat",
		"valid_ingredients.contains_shellfish",
		"valid_ingredients.contains_sesame",
		"valid_ingredients.contains_fish",
		"valid_ingredients.contains_gluten",
		"valid_ingredients.animal_flesh",
		"valid_ingredients.volumetric",
		"valid_ingredients.is_liquid",
		"valid_ingredients.icon_path",
		"valid_ingredients.animal_derived",
		"valid_ingredients.plural_name",
		"valid_ingredients.restrict_to_preparations",
		"valid_ingredients.minimum_ideal_storage_temperature_in_celsius",
		"valid_ingredients.maximum_ideal_storage_temperature_in_celsius",
		"valid_ingredients.storage_instructions",
		"valid_ingredients.created_at",
		"valid_ingredients.last_updated_at",
		"valid_ingredients.archived_at",
	}
)

// scanValidIngredient takes a database Scanner (i.e. *sql.Row) and scans the result into a valid ingredient struct.
func (q *Querier) scanValidIngredient(ctx context.Context, scan database.Scanner, includeCounts bool) (x *types.ValidIngredient, filteredCount, totalCount uint64, err error) {
	_, span := q.tracer.StartSpan(ctx)
	defer span.End()

	x = &types.ValidIngredient{}

	targetVars := []interface{}{
		&x.ID,
		&x.Name,
		&x.Description,
		&x.Warning,
		&x.ContainsEgg,
		&x.ContainsDairy,
		&x.ContainsPeanut,
		&x.ContainsTreeNut,
		&x.ContainsSoy,
		&x.ContainsWheat,
		&x.ContainsShellfish,
		&x.ContainsSesame,
		&x.ContainsFish,
		&x.ContainsGluten,
		&x.AnimalFlesh,
		&x.IsMeasuredVolumetrically,
		&x.IsLiquid,
		&x.IconPath,
		&x.AnimalDerived,
		&x.PluralName,
		&x.RestrictToPreparations,
		&x.MinimumIdealStorageTemperatureInCelsius,
		&x.MaximumIdealStorageTemperatureInCelsius,
		&x.StorageInstructions,
		&x.CreatedAt,
		&x.LastUpdatedAt,
		&x.ArchivedAt,
	}

	if includeCounts {
		targetVars = append(targetVars, &filteredCount, &totalCount)
	}

	if err = scan.Scan(targetVars...); err != nil {
		return nil, 0, 0, observability.PrepareError(err, span, "")
	}

	return x, filteredCount, totalCount, nil
}

// scanValidIngredients takes some database rows and turns them into a slice of valid ingredients.
func (q *Querier) scanValidIngredients(ctx context.Context, rows database.ResultIterator, includeCounts bool) (validIngredients []*types.ValidIngredient, filteredCount, totalCount uint64, err error) {
	_, span := q.tracer.StartSpan(ctx)
	defer span.End()

	for rows.Next() {
		x, fc, tc, scanErr := q.scanValidIngredient(ctx, rows, includeCounts)
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

		validIngredients = append(validIngredients, x)
	}

	if err = q.checkRowsForErrorAndClose(ctx, rows); err != nil {
		return nil, 0, 0, observability.PrepareError(err, span, "handling rows")
	}

	return validIngredients, filteredCount, totalCount, nil
}

//go:embed queries/valid_ingredients/exists.sql
var validIngredientExistenceQuery string

// ValidIngredientExists fetches whether a valid ingredient exists from the database.
func (q *Querier) ValidIngredientExists(ctx context.Context, validIngredientID string) (exists bool, err error) {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	logger := q.logger.Clone()

	if validIngredientID == "" {
		return false, ErrInvalidIDProvided
	}
	logger = logger.WithValue(keys.ValidIngredientIDKey, validIngredientID)
	tracing.AttachValidIngredientIDToSpan(span, validIngredientID)

	args := []interface{}{
		validIngredientID,
	}

	result, err := q.performBooleanQuery(ctx, q.db, validIngredientExistenceQuery, args)
	if err != nil {
		return false, observability.PrepareAndLogError(err, logger, span, "performing valid ingredient existence check")
	}

	return result, nil
}

//go:embed queries/valid_ingredients/get_one.sql
var getValidIngredientQuery string

// GetValidIngredient fetches a valid ingredient from the database.
func (q *Querier) GetValidIngredient(ctx context.Context, validIngredientID string) (*types.ValidIngredient, error) {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	logger := q.logger.Clone()

	if validIngredientID == "" {
		return nil, ErrInvalidIDProvided
	}
	logger = logger.WithValue(keys.ValidIngredientIDKey, validIngredientID)
	tracing.AttachValidIngredientIDToSpan(span, validIngredientID)

	args := []interface{}{
		validIngredientID,
	}

	row := q.getOneRow(ctx, q.db, "valid ingredient", getValidIngredientQuery, args)

	validIngredient, _, _, err := q.scanValidIngredient(ctx, row, false)
	if err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "scanning valid ingredient")
	}

	return validIngredient, nil
}

//go:embed queries/valid_ingredients/get_random.sql
var getRandomValidIngredientQuery string

// GetRandomValidIngredient fetches a valid ingredient from the database.
func (q *Querier) GetRandomValidIngredient(ctx context.Context) (*types.ValidIngredient, error) {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	args := []interface{}{}

	row := q.getOneRow(ctx, q.db, "valid ingredient", getRandomValidIngredientQuery, args)

	validIngredient, _, _, err := q.scanValidIngredient(ctx, row, false)
	if err != nil {
		return nil, observability.PrepareError(err, span, "scanning valid ingredient")
	}

	return validIngredient, nil
}

//go:embed queries/valid_ingredients/search.sql
var validIngredientSearchQuery string

// SearchForValidIngredients fetches a valid ingredient from the database.
func (q *Querier) SearchForValidIngredients(ctx context.Context, query string) ([]*types.ValidIngredient, error) {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	logger := q.logger.Clone()

	if query == "" {
		return nil, ErrInvalidIDProvided
	}
	logger = logger.WithValue(keys.SearchQueryKey, query)
	tracing.AttachValidIngredientIDToSpan(span, query)

	args := []interface{}{
		wrapQueryForILIKE(query),
	}

	rows, err := q.performReadQuery(ctx, q.db, "valid ingredients", validIngredientSearchQuery, args)
	if err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "executing valid ingredients list retrieval query")
	}

	x, _, _, err := q.scanValidIngredients(ctx, rows, false)
	if err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "scanning valid ingredients")
	}

	return x, nil
}

// SearchForValidIngredientsForPreparation fetches a valid ingredient from the database.
func (q *Querier) SearchForValidIngredientsForPreparation(ctx context.Context, preparationID, query string) ([]*types.ValidIngredient, error) {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	logger := q.logger.Clone()

	if query == "" {
		return nil, ErrInvalidIDProvided
	}
	logger = logger.WithValue(keys.SearchQueryKey, query)
	tracing.AttachValidIngredientIDToSpan(span, query)

	args := []interface{}{
		wrapQueryForILIKE(query),
	}

	// TODO: find some way to restrict by preparationID

	rows, err := q.performReadQuery(ctx, q.db, "valid ingredients", validIngredientSearchQuery, args)
	if err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "executing valid ingredients list retrieval query")
	}

	x, _, _, err := q.scanValidIngredients(ctx, rows, false)
	if err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "scanning valid ingredients")
	}

	return x, nil
}

// GetValidIngredients fetches a list of valid ingredients from the database that meet a particular filter.
func (q *Querier) GetValidIngredients(ctx context.Context, filter *types.QueryFilter) (x *types.ValidIngredientList, err error) {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	logger := q.logger.Clone()

	x = &types.ValidIngredientList{}
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

	query, args := q.buildListQuery(ctx, "valid_ingredients", nil, nil, nil, householdOwnershipColumn, validIngredientsTableColumns, "", false, filter)

	rows, err := q.performReadQuery(ctx, q.db, "validIngredients", query, args)
	if err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "executing valid ingredients list retrieval query")
	}

	if x.ValidIngredients, x.FilteredCount, x.TotalCount, err = q.scanValidIngredients(ctx, rows, true); err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "scanning valid ingredients")
	}

	return x, nil
}

//go:embed queries/valid_ingredients/create.sql
var validIngredientCreationQuery string

// CreateValidIngredient creates a valid ingredient in the database.
func (q *Querier) CreateValidIngredient(ctx context.Context, input *types.ValidIngredientDatabaseCreationInput) (*types.ValidIngredient, error) {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	if input == nil {
		return nil, ErrNilInputProvided
	}

	logger := q.logger.WithValue(keys.ValidIngredientIDKey, input.ID)

	args := []interface{}{
		input.ID,
		input.Name,
		input.Description,
		input.Warning,
		input.ContainsEgg,
		input.ContainsDairy,
		input.ContainsPeanut,
		input.ContainsTreeNut,
		input.ContainsSoy,
		input.ContainsWheat,
		input.ContainsShellfish,
		input.ContainsSesame,
		input.ContainsFish,
		input.ContainsGluten,
		input.AnimalFlesh,
		input.IsMeasuredVolumetrically,
		input.IsLiquid,
		input.IconPath,
		input.AnimalDerived,
		input.PluralName,
		input.RestrictToPreparations,
		input.MinimumIdealStorageTemperatureInCelsius,
		input.MaximumIdealStorageTemperatureInCelsius,
		input.StorageInstructions,
	}

	// create the valid ingredient.
	if err := q.performWriteQuery(ctx, q.db, "valid ingredient creation", validIngredientCreationQuery, args); err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "performing valid ingredient creation query")
	}

	x := &types.ValidIngredient{
		ID:                                      input.ID,
		Name:                                    input.Name,
		Description:                             input.Description,
		Warning:                                 input.Warning,
		ContainsEgg:                             input.ContainsEgg,
		ContainsDairy:                           input.ContainsDairy,
		ContainsPeanut:                          input.ContainsPeanut,
		ContainsTreeNut:                         input.ContainsTreeNut,
		ContainsSoy:                             input.ContainsSoy,
		ContainsWheat:                           input.ContainsWheat,
		ContainsShellfish:                       input.ContainsShellfish,
		ContainsSesame:                          input.ContainsSesame,
		ContainsFish:                            input.ContainsFish,
		ContainsGluten:                          input.ContainsGluten,
		AnimalFlesh:                             input.AnimalFlesh,
		IsMeasuredVolumetrically:                input.IsMeasuredVolumetrically,
		IsLiquid:                                input.IsLiquid,
		IconPath:                                input.IconPath,
		AnimalDerived:                           input.AnimalDerived,
		PluralName:                              input.PluralName,
		RestrictToPreparations:                  input.RestrictToPreparations,
		MinimumIdealStorageTemperatureInCelsius: input.MinimumIdealStorageTemperatureInCelsius,
		MaximumIdealStorageTemperatureInCelsius: input.MaximumIdealStorageTemperatureInCelsius,
		StorageInstructions:                     input.StorageInstructions,
		CreatedAt:                               q.currentTime(),
	}

	tracing.AttachValidIngredientIDToSpan(span, x.ID)
	logger.Info("valid ingredient created")

	return x, nil
}

//go:embed queries/valid_ingredients/update.sql
var updateValidIngredientQuery string

// UpdateValidIngredient updates a particular valid ingredient.
func (q *Querier) UpdateValidIngredient(ctx context.Context, updated *types.ValidIngredient) error {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	if updated == nil {
		return ErrNilInputProvided
	}

	logger := q.logger.WithValue(keys.ValidIngredientIDKey, updated.ID)
	tracing.AttachValidIngredientIDToSpan(span, updated.ID)

	args := []interface{}{
		updated.Name,
		updated.Description,
		updated.Warning,
		updated.ContainsEgg,
		updated.ContainsDairy,
		updated.ContainsPeanut,
		updated.ContainsTreeNut,
		updated.ContainsSoy,
		updated.ContainsWheat,
		updated.ContainsShellfish,
		updated.ContainsSesame,
		updated.ContainsFish,
		updated.ContainsGluten,
		updated.AnimalFlesh,
		updated.IsMeasuredVolumetrically,
		updated.IsLiquid,
		updated.IconPath,
		updated.AnimalDerived,
		updated.PluralName,
		updated.RestrictToPreparations,
		updated.MinimumIdealStorageTemperatureInCelsius,
		updated.MaximumIdealStorageTemperatureInCelsius,
		updated.StorageInstructions,
		updated.ID,
	}

	if err := q.performWriteQuery(ctx, q.db, "valid ingredient update", updateValidIngredientQuery, args); err != nil {
		return observability.PrepareAndLogError(err, logger, span, "updating valid ingredient")
	}

	logger.Info("valid ingredient updated")

	return nil
}

//go:embed queries/valid_ingredients/archive.sql
var archiveValidIngredientQuery string

// ArchiveValidIngredient archives a valid ingredient from the database by its ID.
func (q *Querier) ArchiveValidIngredient(ctx context.Context, validIngredientID string) error {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	logger := q.logger.Clone()

	if validIngredientID == "" {
		return ErrInvalidIDProvided
	}
	logger = logger.WithValue(keys.ValidIngredientIDKey, validIngredientID)
	tracing.AttachValidIngredientIDToSpan(span, validIngredientID)

	args := []interface{}{
		validIngredientID,
	}

	if err := q.performWriteQuery(ctx, q.db, "valid ingredient archive", archiveValidIngredientQuery, args); err != nil {
		return observability.PrepareAndLogError(err, logger, span, "updating valid ingredient")
	}

	logger.Info("valid ingredient archived")

	return nil
}