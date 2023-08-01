package postgres

import (
	"context"
	_ "embed"

	"github.com/dinnerdonebetter/backend/internal/database"
	"github.com/dinnerdonebetter/backend/internal/observability"
	"github.com/dinnerdonebetter/backend/internal/observability/keys"
	"github.com/dinnerdonebetter/backend/internal/observability/tracing"
	"github.com/dinnerdonebetter/backend/pkg/types"
	"github.com/dinnerdonebetter/backend/pkg/types/converters"

	"github.com/Masterminds/squirrel"
)

const (
	validVesselsTable = "valid_vessels"
)

var (
	_ types.ValidVesselDataManager = (*Querier)(nil)

	// validVesselsTableColumns are the columns for the valid_vessels table.
	validVesselsTableColumns = []string{
		"valid_vessels.id",
		"valid_vessels.name",
		"valid_vessels.plural_name",
		"valid_vessels.description",
		"valid_vessels.icon_path",
		"valid_vessels.usable_for_storage",
		"valid_vessels.slug",
		"valid_vessels.display_in_summary_lists",
		"valid_vessels.include_in_generated_instructions",
		"valid_vessels.capacity",
		"valid_measurement_units.id",
		"valid_measurement_units.name",
		"valid_measurement_units.description",
		"valid_measurement_units.volumetric",
		"valid_measurement_units.icon_path",
		"valid_measurement_units.universal",
		"valid_measurement_units.metric",
		"valid_measurement_units.imperial",
		"valid_measurement_units.slug",
		"valid_measurement_units.plural_name",
		"valid_measurement_units.created_at",
		"valid_measurement_units.last_updated_at",
		"valid_measurement_units.archived_at",
		"valid_vessels.width_in_millimeters",
		"valid_vessels.length_in_millimeters",
		"valid_vessels.height_in_millimeters",
		"valid_vessels.shape",
		"valid_vessels.created_at",
		"valid_vessels.last_updated_at",
		"valid_vessels.archived_at",
	}
)

// scanValidVessel takes a database Scanner (i.e. *sql.Row) and scans the result into a valid vessel struct.
func (q *Querier) scanValidVessel(ctx context.Context, scan database.Scanner, includeCounts bool) (x *types.ValidVessel, filteredCount, totalCount uint64, err error) {
	_, span := q.tracer.StartSpan(ctx)
	defer span.End()

	z := &types.NullableValidVessel{
		CapacityUnit: &types.NullableValidMeasurementUnit{},
	}

	targetVars := []any{
		&z.ID,
		&z.Name,
		&z.PluralName,
		&z.Description,
		&z.IconPath,
		&z.UsableForStorage,
		&z.Slug,
		&z.DisplayInSummaryLists,
		&z.IncludeInGeneratedInstructions,
		&z.Capacity,
		&z.CapacityUnit.ID,
		&z.CapacityUnit.Name,
		&z.CapacityUnit.Description,
		&z.CapacityUnit.Volumetric,
		&z.CapacityUnit.IconPath,
		&z.CapacityUnit.Universal,
		&z.CapacityUnit.Metric,
		&z.CapacityUnit.Imperial,
		&z.CapacityUnit.Slug,
		&z.CapacityUnit.PluralName,
		&z.CapacityUnit.CreatedAt,
		&z.CapacityUnit.LastUpdatedAt,
		&z.CapacityUnit.ArchivedAt,
		&z.WidthInMillimeters,
		&z.LengthInMillimeters,
		&z.HeightInMillimeters,
		&z.Shape,
		&z.CreatedAt,
		&z.LastUpdatedAt,
		&z.ArchivedAt,
	}

	if includeCounts {
		targetVars = append(targetVars, &filteredCount, &totalCount)
	}

	if err = scan.Scan(targetVars...); err != nil {
		return nil, 0, 0, observability.PrepareError(err, span, "")
	}

	x = converters.ConvertNullableValidVesselToValidVessel(z)

	if x.CapacityUnit != nil && x.CapacityUnit.ID == "" {
		x.CapacityUnit = nil
	}

	return x, filteredCount, totalCount, nil
}

// scanValidVessels takes some database rows and turns them into a slice of valid vessels.
func (q *Querier) scanValidVessels(ctx context.Context, rows database.ResultIterator, includeCounts bool) (validVessels []*types.ValidVessel, filteredCount, totalCount uint64, err error) {
	_, span := q.tracer.StartSpan(ctx)
	defer span.End()

	for rows.Next() {
		x, fc, tc, scanErr := q.scanValidVessel(ctx, rows, includeCounts)
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

		validVessels = append(validVessels, x)
	}

	if err = q.checkRowsForErrorAndClose(ctx, rows); err != nil {
		return nil, 0, 0, observability.PrepareError(err, span, "handling rows")
	}

	return validVessels, filteredCount, totalCount, nil
}

//go:embed queries/valid_vessels/exists.sql
var validVesselExistenceQuery string

// ValidVesselExists fetches whether a valid vessel exists from the database.
func (q *Querier) ValidVesselExists(ctx context.Context, validVesselID string) (exists bool, err error) {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	logger := q.logger.Clone()

	if validVesselID == "" {
		return false, ErrInvalidIDProvided
	}
	logger = logger.WithValue(keys.ValidVesselIDKey, validVesselID)
	tracing.AttachValidVesselIDToSpan(span, validVesselID)

	args := []any{
		validVesselID,
	}

	result, err := q.performBooleanQuery(ctx, q.db, validVesselExistenceQuery, args)
	if err != nil {
		return false, observability.PrepareAndLogError(err, logger, span, "performing valid vessel existence check")
	}

	return result, nil
}

//go:embed queries/valid_vessels/get_one.sql
var getValidVesselQuery string

// GetValidVessel fetches a valid vessel from the database.
func (q *Querier) GetValidVessel(ctx context.Context, validVesselID string) (*types.ValidVessel, error) {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	logger := q.logger.Clone()

	if validVesselID == "" {
		return nil, ErrInvalidIDProvided
	}
	logger = logger.WithValue(keys.ValidVesselIDKey, validVesselID)
	tracing.AttachValidVesselIDToSpan(span, validVesselID)

	args := []any{
		validVesselID,
	}

	row := q.getOneRow(ctx, q.db, "validVessel", getValidVesselQuery, args)

	validVessel, _, _, err := q.scanValidVessel(ctx, row, false)
	if err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "scanning validVessel")
	}

	return validVessel, nil
}

//go:embed queries/valid_vessels/get_random.sql
var getRandomValidVesselQuery string

// GetRandomValidVessel fetches a valid vessel from the database.
func (q *Querier) GetRandomValidVessel(ctx context.Context) (*types.ValidVessel, error) {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	args := []any{}

	row := q.getOneRow(ctx, q.db, "validVessel", getRandomValidVesselQuery, args)

	validVessel, _, _, err := q.scanValidVessel(ctx, row, false)
	if err != nil {
		return nil, observability.PrepareError(err, span, "scanning validVessel")
	}

	return validVessel, nil
}

//go:embed queries/valid_vessels/search.sql
var validVesselSearchQuery string

// SearchForValidVessels fetches a valid vessel from the database.
func (q *Querier) SearchForValidVessels(ctx context.Context, query string) ([]*types.ValidVessel, error) {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	logger := q.logger.Clone()

	if query == "" {
		return nil, ErrEmptyInputProvided
	}
	logger = logger.WithValue(keys.SearchQueryKey, query)
	tracing.AttachValidVesselIDToSpan(span, query)

	args := []any{
		wrapQueryForILIKE(query),
	}

	rows, err := q.getRows(ctx, q.db, "valid ingredients", validVesselSearchQuery, args)
	if err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "executing valid ingredients list retrieval query")
	}

	validVessels, _, _, err := q.scanValidVessels(ctx, rows, false)
	if err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "scanning validVessel")
	}

	return validVessels, nil
}

// SearchForValidVesselsForPreparation fetches a valid vessel from the database.
func (q *Querier) SearchForValidVesselsForPreparation(ctx context.Context, preparationID, query string) ([]*types.ValidVessel, error) {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	logger := q.logger.Clone()

	if query == "" {
		return nil, ErrEmptyInputProvided
	}
	logger = logger.WithValue(keys.SearchQueryKey, query)
	tracing.AttachValidVesselIDToSpan(span, query)

	// TODO: restrict results by preparation ID

	args := []any{
		wrapQueryForILIKE(query),
	}

	rows, err := q.getRows(ctx, q.db, "valid ingredients search", validVesselSearchQuery, args)
	if err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "executing valid ingredients list retrieval query")
	}

	validVessels, _, _, err := q.scanValidVessels(ctx, rows, false)
	if err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "scanning validVessel")
	}

	return validVessels, nil
}

//go:embed queries/valid_vessels/get_many.sql
var getValidVesselsQuery string

// GetValidVessels fetches a list of valid vessels from the database that meet a particular filter.
func (q *Querier) GetValidVessels(ctx context.Context, filter *types.QueryFilter) (x *types.QueryFilteredResult[types.ValidVessel], err error) {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	logger := q.logger.Clone()

	x = &types.QueryFilteredResult[types.ValidVessel]{}
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

	args := []any{
		filter.CreatedBefore,
		filter.CreatedAfter,
		filter.UpdatedBefore,
		filter.UpdatedAfter,
		filter.QueryOffset(),
		filter.Limit,
	}

	rows, err := q.getRows(ctx, q.db, "valid vessels", getValidVesselsQuery, args)
	if err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "executing valid vessels list retrieval query")
	}

	if x.Data, x.FilteredCount, x.TotalCount, err = q.scanValidVessels(ctx, rows, true); err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "scanning valid vessels")
	}

	return x, nil
}

// GetValidVesselsWithIDs fetches a list of valid vessels from the database that meet a particular filter.
func (q *Querier) GetValidVesselsWithIDs(ctx context.Context, ids []string) ([]*types.ValidVessel, error) {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	logger := q.logger.Clone()

	where := squirrel.Eq{"valid_vessels.id": ids}
	joins := []string{
		"valid_measurement_units ON valid_vessels.capacity_unit=valid_measurement_units.id",
	}
	groupBys := []string{
		"valid_vessels.id",
		"valid_measurement_units.id",
	}
	query, args := q.buildListQuery(ctx, validVesselsTable, joins, groupBys, where, householdOwnershipColumn, validVesselsTableColumns, "", false, nil)

	rows, err := q.getRows(ctx, q.db, "valid vessels", query, args)
	if err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "executing valid vessels id list retrieval query")
	}

	instruments, _, _, err := q.scanValidVessels(ctx, rows, true)
	if err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "scanning valid vessels")
	}

	return instruments, nil
}

//go:embed queries/valid_vessels/get_needing_indexing.sql
var validVesselsNeedingIndexingQuery string

// GetValidVesselIDsThatNeedSearchIndexing fetches a list of valid vessels from the database that meet a particular filter.
func (q *Querier) GetValidVesselIDsThatNeedSearchIndexing(ctx context.Context) ([]string, error) {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	rows, err := q.getRows(ctx, q.db, "valid vessels needing indexing", validVesselsNeedingIndexingQuery, nil)
	if err != nil {
		return nil, observability.PrepareError(err, span, "executing valid vessels list retrieval query")
	}

	return q.scanIDs(ctx, rows)
}

//go:embed queries/valid_vessels/create.sql
var validVesselCreationQuery string

// CreateValidVessel creates a valid vessel in the database.
func (q *Querier) CreateValidVessel(ctx context.Context, input *types.ValidVesselDatabaseCreationInput) (*types.ValidVessel, error) {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	if input == nil {
		return nil, ErrNilInputProvided
	}

	logger := q.logger.WithValue(keys.ValidVesselIDKey, input.ID)

	args := []any{
		input.ID,
		input.Name,
		input.PluralName,
		input.Description,
		input.IconPath,
		input.UsableForStorage,
		input.Slug,
		input.DisplayInSummaryLists,
		input.IncludeInGeneratedInstructions,
		input.Capacity,
		input.CapacityUnitID,
		input.WidthInMillimeters,
		input.LengthInMillimeters,
		input.HeightInMillimeters,
		input.Shape,
	}

	// create the valid vessel.
	if err := q.performWriteQuery(ctx, q.db, "valid vessel creation", validVesselCreationQuery, args); err != nil {
		return nil, observability.PrepareAndLogError(err, logger, span, "performing valid vessel creation query")
	}

	x := &types.ValidVessel{
		ID:                             input.ID,
		Name:                           input.Name,
		PluralName:                     input.PluralName,
		Description:                    input.Description,
		IconPath:                       input.IconPath,
		Slug:                           input.Slug,
		Shape:                          input.Shape,
		WidthInMillimeters:             input.WidthInMillimeters,
		LengthInMillimeters:            input.LengthInMillimeters,
		HeightInMillimeters:            input.HeightInMillimeters,
		Capacity:                       input.Capacity,
		IncludeInGeneratedInstructions: input.IncludeInGeneratedInstructions,
		DisplayInSummaryLists:          input.DisplayInSummaryLists,
		UsableForStorage:               input.UsableForStorage,
		CreatedAt:                      q.currentTime(),
	}

	if input.CapacityUnitID != nil {
		x.CapacityUnit = &types.ValidMeasurementUnit{ID: *input.CapacityUnitID}
	}

	tracing.AttachValidVesselIDToSpan(span, x.ID)
	logger.Info("valid vessel created")

	return x, nil
}

//go:embed queries/valid_vessels/update.sql
var updateValidVesselQuery string

// UpdateValidVessel updates a particular valid vessel.
func (q *Querier) UpdateValidVessel(ctx context.Context, updated *types.ValidVessel) error {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	if updated == nil {
		return ErrNilInputProvided
	}

	logger := q.logger.WithValue(keys.ValidVesselIDKey, updated.ID)
	tracing.AttachValidVesselIDToSpan(span, updated.ID)

	args := []any{
		updated.Name,
		updated.PluralName,
		updated.Description,
		updated.IconPath,
		updated.UsableForStorage,
		updated.Slug,
		updated.DisplayInSummaryLists,
		updated.IncludeInGeneratedInstructions,
		updated.Capacity,
		updated.CapacityUnit.ID,
		updated.WidthInMillimeters,
		updated.LengthInMillimeters,
		updated.HeightInMillimeters,
		updated.Shape,
		updated.ID,
	}

	if err := q.performWriteQuery(ctx, q.db, "valid vessel update", updateValidVesselQuery, args); err != nil {
		return observability.PrepareAndLogError(err, logger, span, "updating valid vessel")
	}

	logger.Info("valid vessel updated")

	return nil
}

//go:embed queries/valid_vessels/update_last_indexed_at.sql
var updateValidVesselLastIndexedAtQuery string

// MarkValidVesselAsIndexed updates a particular valid vessel's last_indexed_at value.
func (q *Querier) MarkValidVesselAsIndexed(ctx context.Context, validVesselID string) error {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	logger := q.logger.Clone()

	if validVesselID == "" {
		return ErrInvalidIDProvided
	}
	logger = logger.WithValue(keys.ValidVesselIDKey, validVesselID)
	tracing.AttachValidVesselIDToSpan(span, validVesselID)

	args := []any{
		validVesselID,
	}

	if err := q.performWriteQuery(ctx, q.db, "set valid vessel last_indexed_at", updateValidVesselLastIndexedAtQuery, args); err != nil {
		return observability.PrepareAndLogError(err, logger, span, "marking valid vessel as indexed")
	}

	logger.Info("valid vessel marked as indexed")

	return nil
}

//go:embed queries/valid_vessels/archive.sql
var archiveValidVesselQuery string

// ArchiveValidVessel archives a valid vessel from the database by its ID.
func (q *Querier) ArchiveValidVessel(ctx context.Context, validVesselID string) error {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	logger := q.logger.Clone()

	if validVesselID == "" {
		return ErrInvalidIDProvided
	}
	logger = logger.WithValue(keys.ValidVesselIDKey, validVesselID)
	tracing.AttachValidVesselIDToSpan(span, validVesselID)

	args := []any{
		validVesselID,
	}

	if err := q.performWriteQuery(ctx, q.db, "valid vessel archive", archiveValidVesselQuery, args); err != nil {
		return observability.PrepareAndLogError(err, logger, span, "updating valid vessel")
	}

	logger.Info("valid vessel archived")

	return nil
}