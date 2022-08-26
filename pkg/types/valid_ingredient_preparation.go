package types

import (
	"context"
	"encoding/gob"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

const (
	// ValidIngredientPreparationDataType indicates an event is related to a valid ingredient preparation.
	ValidIngredientPreparationDataType dataType = "valid_ingredient_preparation"

	// ValidIngredientPreparationCreatedCustomerEventType indicates a valid ingredient preparation was created.
	ValidIngredientPreparationCreatedCustomerEventType CustomerEventType = "valid_ingredient_preparation_created"
	// ValidIngredientPreparationUpdatedCustomerEventType indicates a valid ingredient preparation was updated.
	ValidIngredientPreparationUpdatedCustomerEventType CustomerEventType = "valid_ingredient_preparation_updated"
	// ValidIngredientPreparationArchivedCustomerEventType indicates a valid ingredient preparation was archived.
	ValidIngredientPreparationArchivedCustomerEventType CustomerEventType = "valid_ingredient_preparation_archived"
)

func init() {
	gob.Register(new(ValidIngredientPreparation))
	gob.Register(new(ValidIngredientPreparationList))
	gob.Register(new(ValidIngredientPreparationCreationRequestInput))
	gob.Register(new(ValidIngredientPreparationUpdateRequestInput))
}

type (
	// ValidIngredientPreparation represents a valid ingredient preparation.
	ValidIngredientPreparation struct {
		_             struct{}
		ArchivedAt    *uint64          `json:"archivedAt"`
		LastUpdatedAt *uint64          `json:"lastUpdatedAt"`
		Notes         string           `json:"notes"`
		ID            string           `json:"id"`
		Preparation   ValidPreparation `json:"preparation"`
		Ingredient    ValidIngredient  `json:"ingredient"`
		CreatedAt     uint64           `json:"createdAt"`
	}

	// ValidIngredientPreparationList represents a list of valid ingredient preparations.
	ValidIngredientPreparationList struct {
		_                           struct{}
		ValidIngredientPreparations []*ValidIngredientPreparation `json:"data"`
		Pagination
	}

	// ValidIngredientPreparationCreationRequestInput represents what a user could set as input for creating valid ingredient preparations.
	ValidIngredientPreparationCreationRequestInput struct {
		_                  struct{}
		ID                 string `json:"-"`
		Notes              string `json:"notes"`
		ValidPreparationID string `json:"validPreparationID"`
		ValidIngredientID  string `json:"validIngredientID"`
	}

	// ValidIngredientPreparationDatabaseCreationInput represents what a user could set as input for creating valid ingredient preparations.
	ValidIngredientPreparationDatabaseCreationInput struct {
		_ struct{}

		ID                 string `json:"id"`
		Notes              string `json:"notes"`
		ValidPreparationID string `json:"validPreparationID"`
		ValidIngredientID  string `json:"validIngredientID"`
	}

	// ValidIngredientPreparationUpdateRequestInput represents what a user could set as input for updating valid ingredient preparations.
	ValidIngredientPreparationUpdateRequestInput struct {
		_ struct{}

		Notes              *string `json:"notes"`
		ValidPreparationID *string `json:"validPreparationID"`
		ValidIngredientID  *string `json:"validIngredientID"`
	}

	// ValidIngredientPreparationDataManager describes a structure capable of storing valid ingredient preparations permanently.
	ValidIngredientPreparationDataManager interface {
		ValidIngredientPreparationExists(ctx context.Context, validIngredientPreparationID string) (bool, error)
		GetValidIngredientPreparation(ctx context.Context, validIngredientPreparationID string) (*ValidIngredientPreparation, error)
		GetTotalValidIngredientPreparationCount(ctx context.Context) (uint64, error)
		GetValidIngredientPreparations(ctx context.Context, filter *QueryFilter) (*ValidIngredientPreparationList, error)
		GetValidIngredientPreparationsForIngredient(ctx context.Context, ingredientID string, filter *QueryFilter) (*ValidIngredientPreparationList, error)
		GetValidIngredientPreparationsForPreparation(ctx context.Context, preparationID string, filter *QueryFilter) (*ValidIngredientPreparationList, error)
		GetValidIngredientPreparationsWithIDs(ctx context.Context, limit uint8, ids []string) ([]*ValidIngredientPreparation, error)
		CreateValidIngredientPreparation(ctx context.Context, input *ValidIngredientPreparationDatabaseCreationInput) (*ValidIngredientPreparation, error)
		UpdateValidIngredientPreparation(ctx context.Context, updated *ValidIngredientPreparation) error
		ArchiveValidIngredientPreparation(ctx context.Context, validIngredientPreparationID string) error
	}

	// ValidIngredientPreparationDataService describes a structure capable of serving traffic related to valid ingredient preparations.
	ValidIngredientPreparationDataService interface {
		ListHandler(res http.ResponseWriter, req *http.Request)
		CreateHandler(res http.ResponseWriter, req *http.Request)
		ReadHandler(res http.ResponseWriter, req *http.Request)
		UpdateHandler(res http.ResponseWriter, req *http.Request)
		ArchiveHandler(res http.ResponseWriter, req *http.Request)
		SearchByIngredientHandler(res http.ResponseWriter, req *http.Request)
		SearchByPreparationHandler(res http.ResponseWriter, req *http.Request)
	}
)

// Update merges an ValidIngredientPreparationUpdateRequestInput with a valid ingredient preparation.
func (x *ValidIngredientPreparation) Update(input *ValidIngredientPreparationUpdateRequestInput) {
	if input.Notes != nil && *input.Notes != x.Notes {
		x.Notes = *input.Notes
	}

	if input.ValidPreparationID != nil && *input.ValidPreparationID != x.Preparation.ID {
		x.Preparation.ID = *input.ValidPreparationID
	}

	if input.ValidIngredientID != nil && *input.ValidIngredientID != x.Ingredient.ID {
		x.Ingredient.ID = *input.ValidIngredientID
	}
}

var _ validation.ValidatableWithContext = (*ValidIngredientPreparationCreationRequestInput)(nil)

// ValidateWithContext validates a ValidIngredientPreparationCreationRequestInput.
func (x *ValidIngredientPreparationCreationRequestInput) ValidateWithContext(ctx context.Context) error {
	return validation.ValidateStructWithContext(
		ctx,
		x,
		validation.Field(&x.ValidPreparationID, validation.Required),
		validation.Field(&x.ValidIngredientID, validation.Required),
	)
}

var _ validation.ValidatableWithContext = (*ValidIngredientPreparationDatabaseCreationInput)(nil)

// ValidateWithContext validates a ValidIngredientPreparationDatabaseCreationInput.
func (x *ValidIngredientPreparationDatabaseCreationInput) ValidateWithContext(ctx context.Context) error {
	return validation.ValidateStructWithContext(
		ctx,
		x,
		validation.Field(&x.ID, validation.Required),
		validation.Field(&x.ValidPreparationID, validation.Required),
		validation.Field(&x.ValidIngredientID, validation.Required),
	)
}

// ValidIngredientPreparationFromValidIngredientPreparation creates a DatabaseCreationInput from a CreationInput.
func ValidIngredientPreparationFromValidIngredientPreparation(input *ValidIngredientPreparation) *ValidIngredientPreparationUpdateRequestInput {
	x := &ValidIngredientPreparationUpdateRequestInput{
		Notes:              &input.Notes,
		ValidPreparationID: &input.Preparation.ID,
		ValidIngredientID:  &input.Ingredient.ID,
	}

	return x
}

// ValidIngredientPreparationDatabaseCreationInputFromValidIngredientPreparationCreationInput creates a DatabaseCreationInput from a CreationInput.
func ValidIngredientPreparationDatabaseCreationInputFromValidIngredientPreparationCreationInput(input *ValidIngredientPreparationCreationRequestInput) *ValidIngredientPreparationDatabaseCreationInput {
	x := &ValidIngredientPreparationDatabaseCreationInput{
		Notes:              input.Notes,
		ValidPreparationID: input.ValidPreparationID,
		ValidIngredientID:  input.ValidIngredientID,
	}

	return x
}

var _ validation.ValidatableWithContext = (*ValidIngredientPreparationUpdateRequestInput)(nil)

// ValidateWithContext validates a ValidIngredientPreparationUpdateRequestInput.
func (x *ValidIngredientPreparationUpdateRequestInput) ValidateWithContext(ctx context.Context) error {
	return validation.ValidateStructWithContext(
		ctx,
		x,
		validation.Field(&x.ValidPreparationID, validation.Required),
		validation.Field(&x.ValidIngredientID, validation.Required),
	)
}
