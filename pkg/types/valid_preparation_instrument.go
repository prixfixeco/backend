package types

import (
	"context"
	"encoding/gob"
	"net/http"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

const (
	// ValidPreparationInstrumentDataType indicates an event is related to a valid preparation instrument.
	ValidPreparationInstrumentDataType dataType = "valid_preparation_instrument"

	// ValidPreparationInstrumentCreatedCustomerEventType indicates a valid preparation instrument was created.
	ValidPreparationInstrumentCreatedCustomerEventType CustomerEventType = "valid_preparation_instrument_created"
	// ValidPreparationInstrumentUpdatedCustomerEventType indicates a valid preparation instrument was updated.
	ValidPreparationInstrumentUpdatedCustomerEventType CustomerEventType = "valid_preparation_instrument_updated"
	// ValidPreparationInstrumentArchivedCustomerEventType indicates a valid preparation instrument was archived.
	ValidPreparationInstrumentArchivedCustomerEventType CustomerEventType = "valid_preparation_instrument_archived"
)

func init() {
	gob.Register(new(ValidPreparationInstrument))
	gob.Register(new(ValidPreparationInstrumentList))
	gob.Register(new(ValidPreparationInstrumentCreationRequestInput))
	gob.Register(new(ValidPreparationInstrumentUpdateRequestInput))
}

type (
	// ValidPreparationInstrument represents a valid preparation instrument.
	ValidPreparationInstrument struct {
		_             struct{}
		ArchivedAt    *time.Time       `json:"archivedAt"`
		LastUpdatedAt *time.Time       `json:"lastUpdatedAt"`
		Notes         string           `json:"notes"`
		ID            string           `json:"id"`
		Instrument    ValidInstrument  `json:"instrument"`
		Preparation   ValidPreparation `json:"preparation"`
		CreatedAt     time.Time        `json:"createdAt"`
	}

	// ValidPreparationInstrumentList represents a list of valid preparation instruments.
	ValidPreparationInstrumentList struct {
		_                           struct{}
		ValidPreparationInstruments []*ValidPreparationInstrument `json:"data"`
		Pagination
	}

	// ValidPreparationInstrumentCreationRequestInput represents what a user could set as input for creating valid preparation instruments.
	ValidPreparationInstrumentCreationRequestInput struct {
		_                  struct{}
		ID                 string `json:"-"`
		Notes              string `json:"notes"`
		ValidPreparationID string `json:"validPreparationID"`
		ValidInstrumentID  string `json:"validInstrumentID"`
	}

	// ValidPreparationInstrumentDatabaseCreationInput represents what a user could set as input for creating valid preparation instruments.
	ValidPreparationInstrumentDatabaseCreationInput struct {
		_ struct{}

		ID                 string `json:"id"`
		Notes              string `json:"notes"`
		ValidPreparationID string `json:"validPreparationID"`
		ValidInstrumentID  string `json:"validInstrumentID"`
	}

	// ValidPreparationInstrumentUpdateRequestInput represents what a user could set as input for updating valid preparation instruments.
	ValidPreparationInstrumentUpdateRequestInput struct {
		_ struct{}

		Notes              *string `json:"notes"`
		ValidPreparationID *string `json:"validPreparationID"`
		ValidInstrumentID  *string `json:"validInstrumentID"`
	}

	// ValidPreparationInstrumentDataManager describes a structure capable of storing valid preparation instruments permanently.
	ValidPreparationInstrumentDataManager interface {
		ValidPreparationInstrumentExists(ctx context.Context, validPreparationInstrumentID string) (bool, error)
		GetValidPreparationInstrument(ctx context.Context, validPreparationInstrumentID string) (*ValidPreparationInstrument, error)
		GetValidPreparationInstruments(ctx context.Context, filter *QueryFilter) (*ValidPreparationInstrumentList, error)
		GetValidPreparationInstrumentsForPreparation(ctx context.Context, preparationID string, filter *QueryFilter) (*ValidPreparationInstrumentList, error)
		GetValidPreparationInstrumentsForInstrument(ctx context.Context, instrumentID string, filter *QueryFilter) (*ValidPreparationInstrumentList, error)
		CreateValidPreparationInstrument(ctx context.Context, input *ValidPreparationInstrumentDatabaseCreationInput) (*ValidPreparationInstrument, error)
		UpdateValidPreparationInstrument(ctx context.Context, updated *ValidPreparationInstrument) error
		ArchiveValidPreparationInstrument(ctx context.Context, validPreparationInstrumentID string) error
	}

	// ValidPreparationInstrumentDataService describes a structure capable of serving traffic related to valid preparation instruments.
	ValidPreparationInstrumentDataService interface {
		ListHandler(res http.ResponseWriter, req *http.Request)
		CreateHandler(res http.ResponseWriter, req *http.Request)
		ReadHandler(res http.ResponseWriter, req *http.Request)
		UpdateHandler(res http.ResponseWriter, req *http.Request)
		ArchiveHandler(res http.ResponseWriter, req *http.Request)
		SearchByPreparationHandler(res http.ResponseWriter, req *http.Request)
		SearchByInstrumentHandler(res http.ResponseWriter, req *http.Request)
	}
)

// Update merges an ValidPreparationInstrumentUpdateRequestInput with a valid preparation instrument.
func (x *ValidPreparationInstrument) Update(input *ValidPreparationInstrumentUpdateRequestInput) {
	if input.Notes != nil && *input.Notes != x.Notes {
		x.Notes = *input.Notes
	}

	if input.ValidPreparationID != nil && *input.ValidPreparationID != x.Preparation.ID {
		x.Preparation.ID = *input.ValidPreparationID
	}

	if input.ValidInstrumentID != nil && *input.ValidInstrumentID != x.Instrument.ID {
		x.Instrument.ID = *input.ValidInstrumentID
	}
}

var _ validation.ValidatableWithContext = (*ValidPreparationInstrumentCreationRequestInput)(nil)

// ValidateWithContext validates a ValidPreparationInstrumentCreationRequestInput.
func (x *ValidPreparationInstrumentCreationRequestInput) ValidateWithContext(ctx context.Context) error {
	return validation.ValidateStructWithContext(
		ctx,
		x,
		validation.Field(&x.ValidPreparationID, validation.Required),
		validation.Field(&x.ValidInstrumentID, validation.Required),
	)
}

var _ validation.ValidatableWithContext = (*ValidPreparationInstrumentDatabaseCreationInput)(nil)

// ValidateWithContext validates a ValidPreparationInstrumentDatabaseCreationInput.
func (x *ValidPreparationInstrumentDatabaseCreationInput) ValidateWithContext(ctx context.Context) error {
	return validation.ValidateStructWithContext(
		ctx,
		x,
		validation.Field(&x.ID, validation.Required),
		validation.Field(&x.ValidPreparationID, validation.Required),
		validation.Field(&x.ValidInstrumentID, validation.Required),
	)
}

// ValidPreparationInstrumentFromValidPreparationInstrument creates a DatabaseCreationInput from a CreationInput.
func ValidPreparationInstrumentFromValidPreparationInstrument(input *ValidPreparationInstrument) *ValidPreparationInstrumentUpdateRequestInput {
	x := &ValidPreparationInstrumentUpdateRequestInput{
		Notes:              &input.Notes,
		ValidPreparationID: &input.Preparation.ID,
		ValidInstrumentID:  &input.Instrument.ID,
	}

	return x
}

// ValidPreparationInstrumentDatabaseCreationInputFromValidPreparationInstrumentCreationInput creates a DatabaseCreationInput from a CreationInput.
func ValidPreparationInstrumentDatabaseCreationInputFromValidPreparationInstrumentCreationInput(input *ValidPreparationInstrumentCreationRequestInput) *ValidPreparationInstrumentDatabaseCreationInput {
	x := &ValidPreparationInstrumentDatabaseCreationInput{
		Notes:              input.Notes,
		ValidPreparationID: input.ValidPreparationID,
		ValidInstrumentID:  input.ValidInstrumentID,
	}

	return x
}

var _ validation.ValidatableWithContext = (*ValidPreparationInstrumentUpdateRequestInput)(nil)

// ValidateWithContext validates a ValidPreparationInstrumentUpdateRequestInput.
func (x *ValidPreparationInstrumentUpdateRequestInput) ValidateWithContext(ctx context.Context) error {
	return validation.ValidateStructWithContext(
		ctx,
		x,
		validation.Field(&x.ValidPreparationID, validation.Required),
		validation.Field(&x.ValidInstrumentID, validation.Required),
	)
}
