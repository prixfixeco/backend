package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/prixfixeco/api_server/internal/database"
	"github.com/prixfixeco/api_server/internal/observability"
	"github.com/prixfixeco/api_server/internal/observability/keys"
	"github.com/prixfixeco/api_server/internal/observability/tracing"
	"github.com/prixfixeco/api_server/pkg/types"
)

var (
	_ types.APIClientDataManager = (*Querier)(nil)

	apiClientsTableColumns = []string{
		"api_clients.id",
		"api_clients.name",
		"api_clients.client_id",
		"api_clients.secret_key",
		"api_clients.created_at",
		"api_clients.last_updated_at",
		"api_clients.archived_at",
		"api_clients.belongs_to_user",
	}
)

// scanAPIClient takes a Scanner (i.e. *sql.Row) and scans its results into an APIClient struct.
func (q *Querier) scanAPIClient(ctx context.Context, scan database.Scanner, includeCounts bool) (client *types.APIClient, filteredCount, totalCount uint64, err error) {
	_, span := q.tracer.StartSpan(ctx)
	defer span.End()

	logger := q.logger.WithValue("include_counts", includeCounts)

	client = &types.APIClient{}

	targetVars := []interface{}{
		&client.ID,
		&client.Name,
		&client.ClientID,
		&client.ClientSecret,
		&client.CreatedAt,
		&client.LastUpdatedAt,
		&client.ArchivedAt,
		&client.BelongsToUser,
	}

	if includeCounts {
		targetVars = append(targetVars, &filteredCount, &totalCount)
	}

	if err = scan.Scan(targetVars...); err != nil {
		return nil, 0, 0, observability.PrepareError(err, logger, span, "scanning API client database result")
	}

	return client, filteredCount, totalCount, nil
}

// scanAPIClients takes sql rows and turns them into a slice of API Clients.
func (q *Querier) scanAPIClients(ctx context.Context, rows database.ResultIterator, includeCounts bool) (clients []*types.APIClient, filteredCount, totalCount uint64, err error) {
	_, span := q.tracer.StartSpan(ctx)
	defer span.End()

	logger := q.logger.WithValue("include_counts", includeCounts)

	for rows.Next() {
		client, fc, tc, scanErr := q.scanAPIClient(ctx, rows, includeCounts)
		if scanErr != nil {
			return nil, 0, 0, observability.PrepareError(scanErr, logger, span, "scanning API client")
		}

		if includeCounts {
			if filteredCount == 0 {
				filteredCount = fc
			}

			if totalCount == 0 {
				totalCount = tc
			}
		}

		clients = append(clients, client)
	}

	if err = q.checkRowsForErrorAndClose(ctx, rows); err != nil {
		return nil, 0, 0, observability.PrepareError(err, logger, span, "handling rows")
	}

	return clients, filteredCount, totalCount, nil
}

const getAPIClientByClientIDQuery = `
	SELECT
		api_clients.id,
		api_clients.name,
		api_clients.client_id,
		api_clients.secret_key,
		api_clients.created_at,
		api_clients.last_updated_at,
		api_clients.archived_at,
		api_clients.belongs_to_user
	FROM api_clients
	WHERE api_clients.archived_at IS NULL
	AND api_clients.client_id = $1
`

// GetAPIClientByClientID gets an API client from the database.
func (q *Querier) GetAPIClientByClientID(ctx context.Context, clientID string) (*types.APIClient, error) {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	if clientID == "" {
		return nil, ErrEmptyInputProvided
	}

	tracing.AttachAPIClientClientIDToSpan(span, clientID)
	logger := q.logger.WithValue(keys.APIClientClientIDKey, clientID)

	args := []interface{}{clientID}

	row := q.getOneRow(ctx, q.db, "API client", getAPIClientByClientIDQuery, args)

	client, _, _, err := q.scanAPIClient(ctx, row, false)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}

		return nil, observability.PrepareError(err, logger, span, "querying for API client")
	}

	return client, nil
}

const getAPIClientByDatabaseIDQuery = `
	SELECT
		api_clients.id,
		api_clients.name,
		api_clients.client_id,
		api_clients.secret_key,
		api_clients.created_at,
		api_clients.last_updated_at,
		api_clients.archived_at,
		api_clients.belongs_to_user
	FROM api_clients
	WHERE api_clients.archived_at IS NULL
	AND api_clients.belongs_to_user = $1
	AND api_clients.id = $2
`

// GetAPIClientByDatabaseID gets an API client from the database.
func (q *Querier) GetAPIClientByDatabaseID(ctx context.Context, clientID, userID string) (*types.APIClient, error) {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	if clientID == "" || userID == "" {
		return nil, ErrInvalidIDProvided
	}

	tracing.AttachAPIClientDatabaseIDToSpan(span, clientID)
	tracing.AttachUserIDToSpan(span, userID)

	logger := q.logger.WithValues(map[string]interface{}{
		keys.APIClientDatabaseIDKey: clientID,
		keys.UserIDKey:              userID,
	})

	args := []interface{}{userID, clientID}

	row := q.getOneRow(ctx, q.db, "API client", getAPIClientByDatabaseIDQuery, args)

	client, _, _, err := q.scanAPIClient(ctx, row, false)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}

		return nil, observability.PrepareError(err, logger, span, "querying for API client")
	}

	return client, nil
}

// GetAPIClients gets a list of API clients.
func (q *Querier) GetAPIClients(ctx context.Context, userID string, filter *types.QueryFilter) (x *types.APIClientList, err error) {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	if userID == "" {
		return nil, ErrInvalidIDProvided
	}

	logger := filter.AttachToLogger(q.logger).WithValue(keys.UserIDKey, userID)
	tracing.AttachUserIDToSpan(span, userID)
	tracing.AttachQueryFilterToSpan(span, filter)

	x = &types.APIClientList{}
	if filter != nil {
		if filter.Page != nil {
			x.Page = *filter.Page
		}

		if filter.Limit != nil {
			x.Limit = *filter.Limit
		}
	}

	query, args := q.buildListQuery(ctx, "api_clients", nil, nil, nil, userOwnershipColumn, apiClientsTableColumns, userID, false, filter, true)

	rows, err := q.performReadQuery(ctx, q.db, "API clients", query, args)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}

		return nil, observability.PrepareError(err, logger, span, "querying for API clients")
	}

	if x.Clients, x.FilteredCount, x.TotalCount, err = q.scanAPIClients(ctx, rows, true); err != nil {
		return nil, observability.PrepareError(err, logger, span, "scanning response from database")
	}

	return x, nil
}

const createAPIClientQuery = `
	INSERT INTO api_clients (id,name,client_id,secret_key,belongs_to_user) VALUES ($1,$2,$3,$4,$5)
`

// CreateAPIClient creates an API client.
func (q *Querier) CreateAPIClient(ctx context.Context, input *types.APIClientCreationRequestInput) (*types.APIClient, error) {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	if input == nil {
		return nil, ErrNilInputProvided
	}

	logger := q.logger.WithValues(map[string]interface{}{
		keys.APIClientClientIDKey: input.ClientID,
		keys.UserIDKey:            input.BelongsToUser,
	})

	args := []interface{}{
		input.ID,
		input.Name,
		input.ClientID,
		input.ClientSecret,
		input.BelongsToUser,
	}

	if writeErr := q.performWriteQuery(ctx, q.db, "API client creation", createAPIClientQuery, args); writeErr != nil {
		return nil, observability.PrepareError(writeErr, logger, span, "creating API client")
	}

	tracing.AttachAPIClientDatabaseIDToSpan(span, input.ID)

	client := &types.APIClient{
		ID:            input.ID,
		Name:          input.Name,
		ClientID:      input.ClientID,
		ClientSecret:  input.ClientSecret,
		BelongsToUser: input.BelongsToUser,
		CreatedAt:     q.currentTime(),
	}

	logger.Info("API client created")

	return client, nil
}

const archiveAPIClientQuery = `
	UPDATE api_clients SET
		last_updated_at = NOW(),
		archived_at = NOW()
	WHERE archived_at IS NULL
	AND belongs_to_user = $1 AND id = $2
`

// ArchiveAPIClient archives an API client.
func (q *Querier) ArchiveAPIClient(ctx context.Context, clientID, userID string) error {
	ctx, span := q.tracer.StartSpan(ctx)
	defer span.End()

	if clientID == "" || userID == "" {
		return ErrNilInputProvided
	}

	tracing.AttachHouseholdIDToSpan(span, userID)
	tracing.AttachAPIClientDatabaseIDToSpan(span, clientID)

	logger := q.logger.WithValues(map[string]interface{}{
		keys.APIClientDatabaseIDKey: clientID,
		keys.UserIDKey:              userID,
	})

	args := []interface{}{userID, clientID}

	if err := q.performWriteQuery(ctx, q.db, "API client archive", archiveAPIClientQuery, args); err != nil {
		return observability.PrepareError(err, logger, span, "archiving API client")
	}

	logger.Info("API client archived")

	return nil
}
