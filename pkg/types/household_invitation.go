package types

import (
	"context"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

const (
	// HouseholdInvitationDataType indicates an event is related to a household invitation.
	HouseholdInvitationDataType dataType = "household_invitation"

	// PendingHouseholdInvitationStatus indicates a household invitation is pending.
	PendingHouseholdInvitationStatus HouseholdInvitationStatus = "pending"
	// CancelledHouseholdInvitationStatus indicates a household invitation was accepted.
	CancelledHouseholdInvitationStatus HouseholdInvitationStatus = "cancelled"
	// AcceptedHouseholdInvitationStatus indicates a household invitation was accepted.
	AcceptedHouseholdInvitationStatus HouseholdInvitationStatus = "accepted"
	// RejectedHouseholdInvitationStatus indicates a household invitation was rejected.
	RejectedHouseholdInvitationStatus HouseholdInvitationStatus = "rejected"
)

type (
	// HouseholdInvitationStatus is the type to use/compare against when checking invitation status.
	HouseholdInvitationStatus string

	// HouseholdInvitationCreationRequestInput represents what a User could set as input for creating household invitations.
	HouseholdInvitationCreationRequestInput struct {
		_ struct{}

		ID                   string `json:"-"`
		FromUser             string `json:"-"`
		Note                 string `json:"note"`
		ToEmail              string `json:"toEmail"`
		DestinationHousehold string `json:"-"`
	}

	// HouseholdInvitationDatabaseCreationInput represents what a User could set as input for creating household invitations.
	HouseholdInvitationDatabaseCreationInput struct {
		_ struct{}

		ID                   string  `json:"-"`
		FromUser             string  `json:"fromUser"`
		ToUser               *string `json:"toUser"`
		Note                 string  `json:"note"`
		ToEmail              string  `json:"toEmail"`
		Token                string  `json:"token"`
		DestinationHousehold string  `json:"destinationHousehold"`
	}

	// HouseholdInvitation represents a household invitation.
	HouseholdInvitation struct {
		_ struct{}

		LastUpdatedOn        *uint64                   `json:"lastUpdatedOn"`
		ArchivedOn           *uint64                   `json:"archivedOn"`
		FromUser             string                    `json:"fromUser"`
		ToEmail              string                    `json:"toEmail"`
		ToUser               *string                   `json:"toUser"`
		Note                 string                    `json:"note"`
		StatusNote           string                    `json:"statusNote"`
		Token                string                    `json:"token"`
		DestinationHousehold string                    `json:"destinationHousehold"`
		ID                   string                    `json:"id"`
		Status               HouseholdInvitationStatus `json:"status"`
		CreatedOn            uint64                    `json:"createdOn"`
	}

	// HouseholdInvitationList represents a list of households.
	HouseholdInvitationList struct {
		_ struct{}

		HouseholdInvitations []*HouseholdInvitation `json:"householdInvitations"`
		Pagination
	}

	// HouseholdInvitationDataManager describes a structure capable of storing household invitations permanently.
	HouseholdInvitationDataManager interface {
		HouseholdInvitationExists(ctx context.Context, householdInvitationID string) (bool, error)
		GetHouseholdInvitationByHouseholdAndID(ctx context.Context, householdID, householdInvitationID string) (*HouseholdInvitation, error)
		GetHouseholdInvitationByEmailAndToken(ctx context.Context, emailAddress, token string) (*HouseholdInvitation, error)
		GetAllHouseholdInvitationsCount(ctx context.Context) (uint64, error)
		GetPendingHouseholdInvitationsFromUser(ctx context.Context, userID string, filter *QueryFilter) (*HouseholdInvitationList, error)
		GetPendingHouseholdInvitationsForUser(ctx context.Context, userID string, filter *QueryFilter) (*HouseholdInvitationList, error)
		CreateHouseholdInvitation(ctx context.Context, input *HouseholdInvitationDatabaseCreationInput) (*HouseholdInvitation, error)
		CancelHouseholdInvitation(ctx context.Context, householdID, householdInvitationID, note string) error
		AcceptHouseholdInvitation(ctx context.Context, householdID, householdInvitationID, note string) error
		RejectHouseholdInvitation(ctx context.Context, householdID, householdInvitationID, note string) error
	}

	// HouseholdInvitationDataService describes a structure capable of serving traffic related to household invitations.
	HouseholdInvitationDataService interface {
		ReadHandler(res http.ResponseWriter, req *http.Request)
		InboundInvitesHandler(res http.ResponseWriter, req *http.Request)
		OutboundInvitesHandler(res http.ResponseWriter, req *http.Request)
		InviteMemberHandler(res http.ResponseWriter, req *http.Request)
		CancelInviteHandler(res http.ResponseWriter, req *http.Request)
		AcceptInviteHandler(res http.ResponseWriter, req *http.Request)
		RejectInviteHandler(res http.ResponseWriter, req *http.Request)
	}
)

var _ validation.ValidatableWithContext = (*HouseholdInvitationCreationRequestInput)(nil)

// ValidateWithContext validates a HouseholdCreationRequestInput.
func (x *HouseholdInvitationCreationRequestInput) ValidateWithContext(ctx context.Context) error {
	return validation.ValidateStructWithContext(ctx, x,
		validation.Field(&x.ToEmail, validation.Required),
		validation.Field(&x.FromUser, validation.Required),
		validation.Field(&x.DestinationHousehold, validation.Required),
	)
}

// HouseholdInvitationDatabaseCreationInputFromHouseholdInvitationCreationInput creates a DatabaseCreationInput from a CreationInput.
func HouseholdInvitationDatabaseCreationInputFromHouseholdInvitationCreationInput(input *HouseholdInvitationCreationRequestInput) *HouseholdInvitationDatabaseCreationInput {
	x := &HouseholdInvitationDatabaseCreationInput{
		ID:                   input.ID,
		FromUser:             input.FromUser,
		ToEmail:              input.ToEmail,
		DestinationHousehold: input.DestinationHousehold,
	}

	return x
}
