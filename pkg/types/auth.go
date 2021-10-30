package types

import (
	"bytes"
	"context"
	"encoding/gob"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/prixfixeco/api_server/internal/authorization"
	"github.com/prixfixeco/api_server/internal/observability/keys"
	"github.com/prixfixeco/api_server/internal/observability/logging"
)

const (
	// SessionContextDataKey is the non-string type we use for referencing SessionContextData structs.
	SessionContextDataKey ContextKey = "session_context_data"
	// UserIDContextKey is the non-string type we use for referencing SessionContextData structs.
	UserIDContextKey ContextKey = "user_id"
	// HouseholdIDContextKey is the non-string type we use for referencing SessionContextData structs.
	HouseholdIDContextKey ContextKey = "household_id"
	// UserLoginInputContextKey is the non-string type we use for referencing SessionContextData structs.
	UserLoginInputContextKey ContextKey = "user_login_input"
	// UserRegistrationInputContextKey is the non-string type we use for referencing SessionContextData structs.
	UserRegistrationInputContextKey ContextKey = "user_registration_input"
)

func init() {
	gob.Register(&SessionContextData{})
}

type (
	// UserHouseholdMembershipInfo represents key information about an household membership.
	UserHouseholdMembershipInfo struct {
		_ struct{}

		HouseholdName  string   `json:"name"`
		HouseholdID    string   `json:"householdID"`
		HouseholdRoles []string `json:"-"`
	}

	// SessionContextData represents what we encode in our passwords cookies.
	SessionContextData struct {
		_ struct{}

		HouseholdPermissions map[string]authorization.HouseholdRolePermissionsChecker `json:"-"`
		Requester            RequesterInfo                                            `json:"-"`
		ActiveHouseholdID    string                                                   `json:"-"`
	}

	// RequesterInfo contains data relevant to the user making a request.
	RequesterInfo struct {
		_ struct{}

		ServicePermissions    authorization.ServiceRolePermissionChecker `json:"-"`
		Reputation            householdStatus                            `json:"-"`
		ReputationExplanation string                                     `json:"-"`
		UserID                string                                     `json:"-"`
	}

	// UserStatusResponse is what we encode when the frontend wants to check auth status.
	UserStatusResponse struct {
		_ struct{}

		UserReputation            householdStatus `json:"householdStatus,omitempty"`
		UserReputationExplanation string          `json:"reputationExplanation"`
		ActiveHousehold           string          `json:"activeHousehold,omitempty"`
		UserIsAuthenticated       bool            `json:"isAuthenticated"`
	}

	// ChangeActiveHouseholdInput represents what a User could set as input for switching households.
	ChangeActiveHouseholdInput struct {
		_ struct{}

		HouseholdID string `json:"householdID"`
	}

	// PASETOCreationInput is used to create a PASETO.
	PASETOCreationInput struct {
		_ struct{}

		ClientID          string `json:"clientID"`
		HouseholdID       string `json:"householdID"`
		RequestTime       int64  `json:"requestTime"`
		RequestedLifetime uint64 `json:"requestedLifetime,omitempty"`
	}

	// PASETOResponse is used to respond to a PASETO request.
	PASETOResponse struct {
		_ struct{}

		Token     string `json:"token"`
		ExpiresAt string `json:"expiresAt"`
	}

	// AuthService describes a structure capable of handling passwords and authorization requests.
	AuthService interface {
		StatusHandler(res http.ResponseWriter, req *http.Request)
		BeginSessionHandler(res http.ResponseWriter, req *http.Request)
		EndSessionHandler(res http.ResponseWriter, req *http.Request)
		CycleCookieSecretHandler(res http.ResponseWriter, req *http.Request)
		PASETOHandler(res http.ResponseWriter, req *http.Request)
		ChangeActiveHouseholdHandler(res http.ResponseWriter, req *http.Request)

		PermissionFilterMiddleware(permissions ...authorization.Permission) func(next http.Handler) http.Handler
		CookieRequirementMiddleware(next http.Handler) http.Handler
		UserAttributionMiddleware(next http.Handler) http.Handler
		AuthorizationMiddleware(next http.Handler) http.Handler
		ServiceAdminMiddleware(next http.Handler) http.Handler

		AuthenticateUser(ctx context.Context, loginData *UserLoginInput) (*User, *http.Cookie, error)
		LogoutUser(ctx context.Context, req *http.Request, res http.ResponseWriter) error
	}
)

var _ validation.ValidatableWithContext = (*ChangeActiveHouseholdInput)(nil)

// ValidateWithContext validates a ChangeActiveHouseholdInput.
func (x *ChangeActiveHouseholdInput) ValidateWithContext(ctx context.Context) error {
	return validation.ValidateStructWithContext(ctx, x,
		validation.Field(&x.HouseholdID, validation.Required),
	)
}

var _ validation.ValidatableWithContext = (*PASETOCreationInput)(nil)

// ValidateWithContext ensures our  provided UserLoginInput meets expectations.
func (i *PASETOCreationInput) ValidateWithContext(ctx context.Context) error {
	return validation.ValidateStructWithContext(ctx, i,
		validation.Field(&i.ClientID, validation.Required),
		validation.Field(&i.RequestTime, validation.Required),
	)
}

// HouseholdRolePermissionsChecker returns the relevant HouseholdRolePermissionsChecker.
func (x *SessionContextData) HouseholdRolePermissionsChecker() authorization.HouseholdRolePermissionsChecker {
	return x.HouseholdPermissions[x.ActiveHouseholdID]
}

// ServiceRolePermissionChecker returns the relevant ServiceRolePermissionChecker.
func (x *SessionContextData) ServiceRolePermissionChecker() authorization.ServiceRolePermissionChecker {
	return x.Requester.ServicePermissions
}

// ToBytes returns the gob encoded session info.
func (x *SessionContextData) ToBytes() []byte {
	var b bytes.Buffer

	if err := gob.NewEncoder(&b).Encode(x); err != nil {
		panic(err)
	}

	return b.Bytes()
}

// AttachToLogger provides a consistent way to attach a SessionContextData object to a logger.
func (x *SessionContextData) AttachToLogger(logger logging.Logger) logging.Logger {
	if x != nil {
		logger = logger.WithValue(keys.RequesterIDKey, x.Requester.UserID).
			WithValue(keys.ActiveHouseholdIDKey, x.ActiveHouseholdID)

		if x.Requester.ServicePermissions != nil {
			logger = logger.WithValue(keys.ServiceRoleKey, x.Requester.ServicePermissions.IsServiceAdmin())
		}
	}

	return logger
}
