package database

import (
	"context"
	"database/sql"
	"io"

	models "gitlab.com/prixfixe/prixfixe/models/v1"
)

var (
	_ Scanner = (*sql.Row)(nil)
	_ Querier = (*sql.DB)(nil)
	_ Querier = (*sql.Tx)(nil)
)

type (
	// Scanner represents any database response (i.e. sql.Row[s])
	Scanner interface {
		Scan(dest ...interface{}) error
	}

	// ResultIterator represents any iterable database response (i.e. sql.Rows)
	ResultIterator interface {
		Next() bool
		Err() error
		Scanner
		io.Closer
	}

	// Querier is a subset interface for sql.{DB|Tx} objects
	Querier interface {
		ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
		QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
		QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	}

	// ConnectionDetails is a string alias for dependency injection.
	ConnectionDetails string

	// Database describes anything that stores data for our services.
	Database interface {
		Migrate(ctx context.Context, createDummyUser bool) error
		IsReady(ctx context.Context) (ready bool)

		models.ValidInstrumentDataManager
		models.ValidIngredientDataManager
		models.ValidIngredientTagDataManager
		models.IngredientTagMappingDataManager
		models.ValidPreparationDataManager
		models.RequiredPreparationInstrumentDataManager
		models.ValidIngredientPreparationDataManager
		models.RecipeDataManager
		models.RecipeTagDataManager
		models.RecipeStepDataManager
		models.RecipeStepPreparationDataManager
		models.RecipeStepIngredientDataManager
		models.RecipeIterationDataManager
		models.RecipeIterationStepDataManager
		models.IterationMediaDataManager
		models.InvitationDataManager
		models.ReportDataManager
		models.UserDataManager
		models.OAuth2ClientDataManager
		models.WebhookDataManager
	}
)
