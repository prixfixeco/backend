package main

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/pquerna/otp/totp"
	flag "github.com/spf13/pflag"

	"github.com/prixfixeco/api_server/internal/observability/keys"
	"github.com/prixfixeco/api_server/internal/observability/logging"
	"github.com/prixfixeco/api_server/pkg/client/httpclient"
	"github.com/prixfixeco/api_server/pkg/types"
	"github.com/prixfixeco/api_server/pkg/types/fakes"
	testutils "github.com/prixfixeco/api_server/tests/utils"
)

var (
	uri            string
	userCount      uint16
	dataCount      uint16
	debug          bool
	singleUserMode bool

	singleUser *types.User

	quitter = fatalQuitter{}
)

func init() {
	flag.StringVarP(&uri, "url", "u", "", "where the target instance is hosted")
	flag.Uint16VarP(&userCount, "user-count", "c", 0, "how many users to create")
	flag.Uint16VarP(&dataCount, "data-count", "d", 0, "how many households/api clients/etc per user to create")
	flag.BoolVarP(&debug, "debug", "z", false, "whether debug mode is enabled")
	flag.BoolVarP(&singleUserMode, "single-user-mode", "s", false, "whether single user mode is enabled")
}

func clearTheScreen() {
	fmt.Println("\x1b[2J")
	fmt.Printf("\x1b[0;0H")
}

func buildTOTPTokenForSecret(secret string) string {
	secret = strings.ToUpper(secret)
	code, err := totp.GenerateCode(secret, time.Now().UTC())
	if err != nil {
		panic(err)
	}

	if !totp.Validate(code, secret) {
		panic("this shouldn't happen")
	}

	return code
}

func main() {
	flag.Parse()

	ctx := context.Background()
	logger := logging.ProvideLogger(logging.Config{Provider: logging.ProviderZerolog})

	if debug {
		logger.SetLevel(logging.DebugLevel)
	}

	if dataCount <= 0 {
		logger.Debug("exiting early because the requested amount is already satisfied")
		quitter.Quit(0)
	}

	if dataCount == 1 && !singleUserMode {
		singleUserMode = true
	}

	if uri == "" {
		quitter.ComplainAndQuit("uri must be valid")
	}

	parsedURI, uriParseErr := url.Parse(uri)
	if uriParseErr != nil {
		quitter.ComplainAndQuit(fmt.Errorf("parsing provided url: %w", uriParseErr))
	}
	if parsedURI.Scheme == "" {
		quitter.ComplainAndQuit("provided URI missing scheme")
	}

	wg := &sync.WaitGroup{}

	for i := 0; i < int(userCount); i++ {
		wg.Add(1)
		go func(x int, wg *sync.WaitGroup) {
			createdUser, userCreationErr := testutils.CreateServiceUser(ctx, uri, "")
			if userCreationErr != nil {
				quitter.ComplainAndQuit(fmt.Errorf("creating user %q: %w", x, userCreationErr))
			}

			if x == 0 && singleUserMode {
				singleUser = createdUser
			}

			userLogger := logger.
				WithValue("username", createdUser.Username).
				WithValue("password", createdUser.HashedPassword).
				WithValue("totp_secret", createdUser.TwoFactorSecret).
				WithValue("user_id", createdUser.ID).
				WithValue("user_number", x)

			userLogger.Debug("created user")

			cookie, cookieErr := testutils.GetLoginCookie(ctx, uri, createdUser)
			if cookieErr != nil {
				quitter.ComplainAndQuit(fmt.Errorf("getting cookie: %v", cookieErr))
			}

			userClient, err := httpclient.NewClient(parsedURI, httpclient.UsingLogger(userLogger), httpclient.UsingCookie(cookie))
			if err != nil {
				quitter.ComplainAndQuit(fmt.Errorf("initializing client: %w", err))
			}

			userLogger.Debug("assigned user API client")

			wg.Add(1)
			go func(wg *sync.WaitGroup) {
				for j := 0; j < int(dataCount); j++ {
					iterationLogger := userLogger.WithValue("creating", "households").WithValue("iteration", j)

					createdHousehold, householdCreationError := userClient.CreateHousehold(ctx, fakes.BuildFakeHouseholdCreationInput())
					if householdCreationError != nil {
						quitter.ComplainAndQuit(fmt.Errorf("creating household %s: %w", j, householdCreationError))
					}

					iterationLogger.WithValue(keys.HouseholdIDKey, createdHousehold.ID).Debug("created household")
				}
				wg.Done()
			}(wg)

			wg.Add(1)
			go func(wg *sync.WaitGroup) {
				for j := 0; j < int(dataCount); j++ {
					iterationLogger := userLogger.WithValue("creating", "api_clients").WithValue("iteration", j)

					code, codeErr := totp.GenerateCode(strings.ToUpper(createdUser.TwoFactorSecret), time.Now().UTC())
					if codeErr != nil {
						quitter.ComplainAndQuit(fmt.Errorf("creating API Client %s: %w", j, codeErr))
					}

					fakeInput := fakes.BuildFakeAPIClientCreationInput()

					createdAPIClient, apiClientCreationErr := userClient.CreateAPIClient(ctx, cookie, &types.APIClientCreationInput{
						UserLoginInput: types.UserLoginInput{
							Username:  createdUser.Username,
							Password:  createdUser.HashedPassword,
							TOTPToken: code,
						},
						Name: fakeInput.Name,
					})
					if apiClientCreationErr != nil {
						quitter.ComplainAndQuit(fmt.Errorf("API Client webhook %s: %w", j, apiClientCreationErr))
					}

					iterationLogger.WithValue(keys.APIClientDatabaseIDKey, createdAPIClient.ID).Debug("created API Client")
				}
				wg.Done()
			}(wg)

			wg.Add(1)
			go func(wg *sync.WaitGroup) {
				for j := 0; j < int(dataCount); j++ {
					iterationLogger := userLogger.WithValue("creating", "webhooks").WithValue("iteration", j)

					createdWebhookID, webhookCreationErr := userClient.CreateWebhook(ctx, fakes.BuildFakeWebhookCreationInput())
					if webhookCreationErr != nil {
						quitter.ComplainAndQuit(fmt.Errorf("creating webhook %s: %w", j, webhookCreationErr))
					}

					iterationLogger.WithValue(keys.WebhookIDKey, createdWebhookID).Debug("created webhook")
				}
				wg.Done()
			}(wg)

			wg.Add(1)
			go func(wg *sync.WaitGroup) {
				for j := 0; j < int(dataCount); j++ {
					iterationLogger := userLogger.WithValue("creating", "valid instruments").WithValue("iteration", j)

					// create valid instrument
					createdValidInstrumentID, validInstrumentCreationErr := userClient.CreateValidInstrument(ctx, fakes.BuildFakeValidInstrumentCreationRequestInput())
					if validInstrumentCreationErr != nil {
						quitter.ComplainAndQuit(fmt.Errorf("creating valid instrument #%d: %w", j, validInstrumentCreationErr))
					}

					iterationLogger.WithValue(keys.ValidInstrumentIDKey, createdValidInstrumentID).Debug("created valid instrument")
				}
				wg.Done()
			}(wg)

			wg.Add(1)
			go func(wg *sync.WaitGroup) {
				for j := 0; j < int(dataCount); j++ {
					iterationLogger := userLogger.WithValue("creating", "valid ingredients").WithValue("iteration", j)

					// create valid ingredient
					createdValidIngredientID, validIngredientCreationErr := userClient.CreateValidIngredient(ctx, fakes.BuildFakeValidIngredientCreationRequestInput())
					if validIngredientCreationErr != nil {
						quitter.ComplainAndQuit(fmt.Errorf("creating valid ingredient #%d: %w", j, validIngredientCreationErr))
					}

					iterationLogger.WithValue(keys.ValidIngredientIDKey, createdValidIngredientID).Debug("created valid ingredient")
				}
				wg.Done()
			}(wg)

			wg.Add(1)
			go func(wg *sync.WaitGroup) {
				for j := 0; j < int(dataCount); j++ {
					iterationLogger := userLogger.WithValue("creating", "valid preparations").WithValue("iteration", j)

					// create valid preparation
					createdValidPreparationID, validPreparationCreationErr := userClient.CreateValidPreparation(ctx, fakes.BuildFakeValidPreparationCreationRequestInput())
					if validPreparationCreationErr != nil {
						quitter.ComplainAndQuit(fmt.Errorf("creating valid preparation #%d: %w", j, validPreparationCreationErr))
					}

					iterationLogger.WithValue(keys.ValidPreparationIDKey, createdValidPreparationID).Debug("created valid preparation")
				}
				wg.Done()
			}(wg)

			wg.Add(1)
			go func(wg *sync.WaitGroup) {
				for j := 0; j < int(dataCount); j++ {
					iterationLogger := userLogger.WithValue("creating", "valid ingredient preparations").WithValue("iteration", j)

					// create valid ingredient preparation
					createdValidIngredientPreparationID, validIngredientPreparationCreationErr := userClient.CreateValidIngredientPreparation(ctx, fakes.BuildFakeValidIngredientPreparationCreationRequestInput())
					if validIngredientPreparationCreationErr != nil {
						quitter.ComplainAndQuit(fmt.Errorf("creating valid ingredient preparation #%d: %w", j, validIngredientPreparationCreationErr))
					}

					iterationLogger.WithValue(keys.ValidIngredientPreparationIDKey, createdValidIngredientPreparationID).Debug("created valid ingredient preparation")
				}
				wg.Done()
			}(wg)

			wg.Add(1)
			go func(wg *sync.WaitGroup) {
				for j := 0; j < int(dataCount); j++ {
					iterationLogger := userLogger.WithValue("creating", "recipes").WithValue("iteration", j)

					// create recipe
					createdRecipeID, recipeCreationErr := userClient.CreateRecipe(ctx, fakes.BuildFakeRecipeCreationRequestInput())
					if recipeCreationErr != nil {
						quitter.ComplainAndQuit(fmt.Errorf("creating recipe #%d: %w", j, recipeCreationErr))
					}

					iterationLogger.WithValue(keys.RecipeIDKey, createdRecipeID).Debug("created recipe")
				}
				wg.Done()
			}(wg)

			wg.Done()
		}(i, wg)
	}

	wg.Wait()

	if singleUserMode && singleUser != nil {
		logger.Debug("engage single user mode!")

		for range time.Tick(1 * time.Second) {
			clearTheScreen()
			fmt.Printf(`

username:  %s
passwords:  %s
2FA token: %s

`, singleUser.Username, singleUser.HashedPassword, buildTOTPTokenForSecret(singleUser.TwoFactorSecret))
		}
	}
}
