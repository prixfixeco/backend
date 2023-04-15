package integration

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/prixfixeco/backend/internal/authentication"
	"github.com/prixfixeco/backend/internal/authorization"
	"github.com/prixfixeco/backend/internal/database"
	dbconfig "github.com/prixfixeco/backend/internal/database/config"
	"github.com/prixfixeco/backend/internal/database/postgres"
	"github.com/prixfixeco/backend/internal/identifiers"
	"github.com/prixfixeco/backend/internal/observability/keys"
	logcfg "github.com/prixfixeco/backend/internal/observability/logging/config"
	"github.com/prixfixeco/backend/internal/observability/tracing"
	serverutils "github.com/prixfixeco/backend/internal/server/utils"
	"github.com/prixfixeco/backend/pkg/types"

	_ "github.com/lib/pq"
)

const (
	debug         = true
	nonexistentID = "_NOT_REAL_LOL_"
)

var (
	urlToUse       string
	parsedURLToUse *url.URL
	dbmanager      database.DataManager

	premadeAdminUser = &types.User{
		ID:              identifiers.New(),
		TwoFactorSecret: "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=",
		EmailAddress:    "integration_tests@prixfixe.email",
		Username:        "exampleUser",
		HashedPassword:  "integration-tests-are-cool",
	}
)

func init() {
	ctx, span := tracing.StartSpan(context.Background())
	defer span.End()

	logger, err := (&logcfg.Config{Provider: logcfg.ProviderZerolog}).ProvideLogger()
	if err != nil {
		log.Fatal(err)
	}

	parsedURLToUse = serverutils.DetermineServiceURL()
	urlToUse = parsedURLToUse.String()

	logger.WithValue(keys.URLKey, urlToUse).Info("checking server")
	serverutils.EnsureServerIsUp(ctx, urlToUse)

	dbAddr := os.Getenv("TARGET_DATABASE")
	if dbAddr == "" {
		panic("empty database address provided")
	}

	cfg := &dbconfig.Config{
		ConnectionDetails: database.ConnectionDetails(dbAddr),
		Debug:             false,
		RunMigrations:     false,
		MaxPingAttempts:   50,
	}

	dbm, err := postgres.ProvideDatabaseClient(ctx, logger, cfg, tracing.NewNoopTracerProvider())
	if err != nil {
		panic(err)
	}

	dbmanager = dbm

	hasher := authentication.ProvideArgon2Authenticator(logger, tracing.NewNoopTracerProvider())
	actuallyHashedPass, err := hasher.HashPassword(ctx, premadeAdminUser.HashedPassword)
	if err != nil {
		panic(err)
	}

	if _, err = dbmanager.GetUserByUsername(ctx, premadeAdminUser.Username); err != nil {
		_, creationErr := dbmanager.CreateUser(ctx, &types.UserDatabaseCreationInput{
			ID:              premadeAdminUser.ID,
			Username:        premadeAdminUser.Username,
			EmailAddress:    premadeAdminUser.EmailAddress,
			HashedPassword:  actuallyHashedPass,
			TwoFactorSecret: premadeAdminUser.TwoFactorSecret,
		})
		if creationErr != nil {
			panic(creationErr)
		}
	}

	if err = dbmanager.MarkUserTwoFactorSecretAsVerified(ctx, premadeAdminUser.ID); err != nil && !errors.Is(err, sql.ErrNoRows) {
		panic(err)
	}

	db, err := sql.Open("postgres", dbAddr)
	if err != nil {
		panic(err)
	}

	if _, err = db.Exec(`UPDATE users SET service_role = $1 WHERE id = $2`, authorization.ServiceAdminRole.String(), premadeAdminUser.ID); err != nil {
		panic(err)
	}

	fiftySpaces := strings.Repeat("\n", 50)
	fmt.Printf("%s\tRunning tests%s", fiftySpaces, fiftySpaces)

	time.Sleep(3 * time.Second)
}
