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
	validPreparationsOnRecipeStepsJoinClause = "valid_preparations ON recipe_steps.preparation_id=valid_preparations.id"
)

var (
	_ types.ValidPreparationDataManager = (*Querier)(nil)

	// validPreparationsTableColumns are the columns for the valid_preparations table.
	validPreparationsTableColumns = []string{
		"valid_preparations.id",
		"valid_preparations.name",
		"valid_preparations.description",
		"valid_preparations.icon_path",
		"valid_preparations.yields_nothing",
		"valid_preparations.restrict_to_ingredients",
		"valid_preparations.zero_ingredients_allowable",
		"valid_preparations.past_tense",
		"valid_preparations.created_at",
		"valid_preparations.last_updated_at",
		"valid_preparations.archived_at",
	}
)

// scanValidPreparation takes a database Scanner (i.e. *sql.Row) and scans the result into a valid preparation struct.
func (q *Querier) scanValidPreparation(ctx context.Context, scan database.Scanner, includeCounts bool) (x *types.ValidPreparation, filteredCount, totalCount uint64, err error) {
	_, span := q.tracer.StartSpan(ctx)
	defer span.End()

	x = &types.ValidPreparation{}

	targetVars := []interface{}{
		&x.ID,
		&x.Name,
		&x.Description,
		&x.IconPath,
		&x.YieldsNothing,
		&x.RestrictToIngredients,
		&x.ZeroIngredientsAllowable,
		&x.PastTense,
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

// scanValidPreparations takes some database rows and turns them into a slice of valid preparations.
func (q *Querier) scanValidPreparations(ctx context.Context, rows database.ResultIterator, includeCounts bool) (validPreparations []*types.ValidPreparation, filteredCount, totalCount uint64, err error) {
	_, span := q.tracer.StartSpan(ctx)
	defer span.End()

	for rows.Next() {
		x, fc, tc, scanErr := q.scanValidPreparation(ctx, rows, includeCounts)
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

		validPreparations = append(validPreparations, x)
	}

	if err = q.checkRowsForErrorAndClose(ctx, rows); err != nil {
		return nil, 0, 0, observability.PrepareError(err, span, "handling rows")
	}

	return validPreparations, filteredCount, totalCount, nil
}

//go:embed queries/valid_preparations/exists.sql
var validPreparationExistenceQuery string

// ValidPreparationExists fetches whether a valid preparation exists from the database.
func (q *Querier) ValidPreparationExists(ctx context.Context, validPreparationID string) (exists bool, err error) {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	logger := q.logger.Clone()

	if validPreparationID == "" {
		return false, ErrInvalidIDProvided
	}
	logger = logger.WithValue(keys.ValidPreparationIDKey, validPreparationID)
	tracing.AttachValidPreparationIDToSpan(span, validPreparationID)

	args := []interface{}{
		validPreparationID,
	}

	result, err := q.performBooleanQuery(ctx, q.db, validPreparationExistenceQuery, args)
	if err != nil {
		return false, observability.PrepareAndLogError(err, logger, span, "performing valid preparation existence check")
	}

	return result, nil
}

//go:embed queries/valid_preparations/get_one.sql
var getValidPreparationQuery string

// GetValidPreparation fetches a valid preparation from the database.
func (q *Querier) GetValidPreparation(ctx context.Context, validPreparationID string) (*types.ValidPreparation, error) {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	logger := q.logger.Clone()

	if validPreparationID == "" {
		return nil, ErrInvalidIDProvided
	}
	logger = logger.WithValue(keys.ValidPreparationIDKey, validPreparationID)
	tracing.AttachValidPreparationIDToSpan(span, validPreparationID)

	args := []interface{}{
		validPreparationID,
	}

	row := q.getOneRow(ctx, q.db, "validPreparation", getValidPreparationQuery, args)

	validPreparation, _, _, err := q.scanValidPreparation(ctx, row, false)
	if err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "scanning validPreparation")
	}

	return validPreparation, nil
}

//go:embed queries/valid_preparations/get_random.sql
var getRandomValidPreparationQuery string

// GetRandomValidPreparation fetches a valid preparation from the database.
func (q *Querier) GetRandomValidPreparation(ctx context.Context) (*types.ValidPreparation, error) {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	args := []interface{}{}

	row := q.getOneRow(ctx, q.db, "validPreparation", getRandomValidPreparationQuery, args)

	validPreparation, _, _, err := q.scanValidPreparation(ctx, row, false)
	if err != nil {
		return nil, observability.PrepareError(err, span, "scanning validPreparation")
	}

	return validPreparation, nil
}

//go:embed queries/valid_preparations/search.sql
var validPreparationSearchQuery string

// SearchForValidPreparations fetches a valid preparation from the database.
func (q *Querier) SearchForValidPreparations(ctx context.Context, query string) ([]*types.ValidPreparation, error) {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	logger := q.logger.Clone()

	if query == "" {
		return nil, ErrInvalidIDProvided
	}
	logger = logger.WithValue(keys.SearchQueryKey, query)
	tracing.AttachValidPreparationIDToSpan(span, query)

	args := []interface{}{
		wrapQueryForILIKE(query),
	}

	rows, err := q.performReadQuery(ctx, q.db, "valid preparations", validPreparationSearchQuery, args)
	if err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "executing valid preparations list retrieval query")
	}

	x, _, _, err := q.scanValidPreparations(ctx, rows, false)
	if err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "scanning valid preparations")
	}

	return x, nil
}

// GetValidPreparations fetches a list of valid preparations from the database that meet a particular filter.
func (q *Querier) GetValidPreparations(ctx context.Context, filter *types.QueryFilter) (x *types.ValidPreparationList, err error) {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	logger := q.logger.Clone()

	x = &types.ValidPreparationList{}
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

	query, args := q.buildListQuery(ctx, "valid_preparations", nil, nil, nil, householdOwnershipColumn, validPreparationsTableColumns, "", false, filter)

	rows, err := q.performReadQuery(ctx, q.db, "validPreparations", query, args)
	if err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "executing valid preparations list retrieval query")
	}

	if x.ValidPreparations, x.FilteredCount, x.TotalCount, err = q.scanValidPreparations(ctx, rows, true); err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "scanning valid preparations")
	}

	return x, nil
}

//go:embed queries/valid_preparations/create.sql
var validPreparationCreationQuery string

// CreateValidPreparation creates a valid preparation in the database.
func (q *Querier) CreateValidPreparation(ctx context.Context, input *types.ValidPreparationDatabaseCreationInput) (*types.ValidPreparation, error) {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	if input == nil {
		return nil, ErrNilInputProvided
	}

	logger := q.logger.WithValue(keys.ValidPreparationIDKey, input.ID)

	args := []interface{}{
		input.ID,
		input.Name,
		input.Description,
		input.IconPath,
		input.YieldsNothing,
		input.RestrictToIngredients,
		input.ZeroIngredientsAllowable,
		input.PastTense,
	}

	// create the valid preparation.
	if err := q.performWriteQuery(ctx, q.db, "valid preparation creation", validPreparationCreationQuery, args); err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "performing valid preparation creation query")
	}

	x := &types.ValidPreparation{
		ID:                       input.ID,
		Name:                     input.Name,
		Description:              input.Description,
		IconPath:                 input.IconPath,
		YieldsNothing:            input.YieldsNothing,
		RestrictToIngredients:    input.RestrictToIngredients,
		ZeroIngredientsAllowable: input.ZeroIngredientsAllowable,
		PastTense:                input.PastTense,
		CreatedAt:                q.currentTime(),
	}

	tracing.AttachValidPreparationIDToSpan(span, x.ID)
	logger.Info("valid preparation created")

	return x, nil
}

//go:embed queries/valid_preparations/update.sql
var updateValidPreparationQuery string

// UpdateValidPreparation updates a particular valid preparation.
func (q *Querier) UpdateValidPreparation(ctx context.Context, updated *types.ValidPreparation) error {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	if updated == nil {
		return ErrNilInputProvided
	}

	logger := q.logger.WithValue(keys.ValidPreparationIDKey, updated.ID)
	tracing.AttachValidPreparationIDToSpan(span, updated.ID)

	args := []interface{}{
		updated.Name,
		updated.Description,
		updated.IconPath,
		updated.YieldsNothing,
		updated.RestrictToIngredients,
		updated.ZeroIngredientsAllowable,
		updated.PastTense,
		updated.ID,
	}

	if err := q.performWriteQuery(ctx, q.db, "valid preparation update", updateValidPreparationQuery, args); err != nil {
		return observability.PrepareAndLogError(err, logger, span, "updating valid preparation")
	}

	logger.Info("valid preparation updated")

	return nil
}

//go:embed queries/valid_preparations/archive.sql
var archiveValidPreparationQuery string

// ArchiveValidPreparation archives a valid preparation from the database by its ID.
func (q *Querier) ArchiveValidPreparation(ctx context.Context, validPreparationID string) error {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	logger := q.logger.Clone()

	if validPreparationID == "" {
		return ErrInvalidIDProvided
	}
	logger = logger.WithValue(keys.ValidPreparationIDKey, validPreparationID)
	tracing.AttachValidPreparationIDToSpan(span, validPreparationID)

	args := []interface{}{
		validPreparationID,
	}

	if err := q.performWriteQuery(ctx, q.db, "valid preparation archive", archiveValidPreparationQuery, args); err != nil {
		return observability.PrepareAndLogError(err, logger, span, "updating valid preparation")
	}

	logger.Info("valid preparation archived")

	return nil
}