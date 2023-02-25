package types

import (
	"context"
	"encoding/gob"
	"net/http"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

const (
	// ValidMeasurementUnitDataType indicates an event is related to a valid measurement unit.
	ValidMeasurementUnitDataType dataType = "valid_measurement_unit"

	// ValidMeasurementUnitCreatedCustomerEventType indicates a valid measurement unit was created.
	ValidMeasurementUnitCreatedCustomerEventType CustomerEventType = "valid_measurement_unit_created"
	// ValidMeasurementUnitUpdatedCustomerEventType indicates a valid measurement unit was updated.
	ValidMeasurementUnitUpdatedCustomerEventType CustomerEventType = "valid_measurement_unit_updated"
	// ValidMeasurementUnitArchivedCustomerEventType indicates a valid measurement unit was archived.
	ValidMeasurementUnitArchivedCustomerEventType CustomerEventType = "valid_measurement_unit_archived"
)

func init() {
	gob.Register(new(ValidMeasurementUnit))
	gob.Register(new(ValidMeasurementUnitCreationRequestInput))
	gob.Register(new(ValidMeasurementUnitUpdateRequestInput))
}

type (
	// ValidMeasurementUnit represents a valid measurement unit.
	ValidMeasurementUnit struct {
		_ struct{}

		CreatedAt     time.Time  `json:"createdAt"`
		LastUpdatedAt *time.Time `json:"lastUpdatedAt"`
		ArchivedAt    *time.Time `json:"archivedAt"`
		Name          string     `json:"name"`
		IconPath      string     `json:"iconPath"`
		ID            string     `json:"id"`
		Description   string     `json:"description"`
		PluralName    string     `json:"pluralName"`
		Slug          string     `json:"slug"`
		Volumetric    bool       `json:"volumetric"`
		Universal     bool       `json:"universal"`
		Metric        bool       `json:"metric"`
		Imperial      bool       `json:"imperial"`
	}

	// ValidMeasurementUnitCreationRequestInput represents what a user could set as input for creating valid measurement units.
	ValidMeasurementUnitCreationRequestInput struct {
		_ struct{}

		CreatedAt   time.Time `json:"createdAt"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		IconPath    string    `json:"iconPath"`
		PluralName  string    `json:"pluralName"`
		Slug        string    `json:"slug"`
		Volumetric  bool      `json:"volumetric"`
		Universal   bool      `json:"universal"`
		Metric      bool      `json:"metric"`
		Imperial    bool      `json:"imperial"`
	}

	// ValidMeasurementUnitDatabaseCreationInput represents what a user could set as input for creating valid measurement units.
	ValidMeasurementUnitDatabaseCreationInput struct {
		_ struct{}

		CreatedAt   time.Time
		Name        string
		Description string
		ID          string
		IconPath    string
		PluralName  string
		Slug        string
		Volumetric  bool
		Universal   bool
		Metric      bool
		Imperial    bool
	}

	// ValidMeasurementUnitUpdateRequestInput represents what a user could set as input for updating valid measurement units.
	ValidMeasurementUnitUpdateRequestInput struct {
		_ struct{}

		Name        *string    `json:"name"`
		Description *string    `json:"description"`
		IconPath    *string    `json:"iconPath"`
		CreatedAt   *time.Time `json:"createdAt"`
		Volumetric  *bool      `json:"volumetric"`
		Universal   *bool      `json:"universal"`
		Metric      *bool      `json:"metric"`
		Imperial    *bool      `json:"imperial"`
		PluralName  *string    `json:"pluralName"`
		Slug        *string    `json:"slug"`
	}

	// ValidMeasurementUnitDataManager describes a structure capable of storing valid measurement units permanently.
	ValidMeasurementUnitDataManager interface {
		ValidMeasurementUnitExists(ctx context.Context, validMeasurementUnitID string) (bool, error)
		GetValidMeasurementUnit(ctx context.Context, validMeasurementUnitID string) (*ValidMeasurementUnit, error)
		GetValidMeasurementUnits(ctx context.Context, filter *QueryFilter) (*QueryFilteredResult[ValidMeasurementUnit], error)
		SearchForValidMeasurementUnitsByName(ctx context.Context, query string) ([]*ValidMeasurementUnit, error)
		ValidMeasurementUnitsForIngredientID(ctx context.Context, validIngredientID string, filter *QueryFilter) (*QueryFilteredResult[ValidMeasurementUnit], error)
		CreateValidMeasurementUnit(ctx context.Context, input *ValidMeasurementUnitDatabaseCreationInput) (*ValidMeasurementUnit, error)
		UpdateValidMeasurementUnit(ctx context.Context, updated *ValidMeasurementUnit) error
		ArchiveValidMeasurementUnit(ctx context.Context, validMeasurementUnitID string) error
	}

	// ValidMeasurementUnitDataService describes a structure capable of serving traffic related to valid measurement units.
	ValidMeasurementUnitDataService interface {
		SearchHandler(res http.ResponseWriter, req *http.Request)
		SearchByIngredientIDHandler(res http.ResponseWriter, req *http.Request)
		ListHandler(res http.ResponseWriter, req *http.Request)
		CreateHandler(res http.ResponseWriter, req *http.Request)
		ReadHandler(res http.ResponseWriter, req *http.Request)
		UpdateHandler(res http.ResponseWriter, req *http.Request)
		ArchiveHandler(res http.ResponseWriter, req *http.Request)
	}
)

// Update merges an ValidMeasurementUnitUpdateRequestInput with a valid measurement unit.
func (x *ValidMeasurementUnit) Update(input *ValidMeasurementUnitUpdateRequestInput) {
	if input.Name != nil && *input.Name != x.Name {
		x.Name = *input.Name
	}

	if input.Description != nil && *input.Description != x.Description {
		x.Description = *input.Description
	}

	if input.Volumetric != nil && *input.Volumetric != x.Volumetric {
		x.Volumetric = *input.Volumetric
	}

	if input.IconPath != nil && *input.IconPath != x.IconPath {
		x.IconPath = *input.IconPath
	}

	if input.Universal != nil && *input.Universal != x.Universal {
		x.Universal = *input.Universal
	}

	if input.Metric != nil && *input.Metric != x.Metric {
		x.Metric = *input.Metric
	}

	if input.Imperial != nil && *input.Imperial != x.Imperial {
		x.Imperial = *input.Imperial
	}

	if input.PluralName != nil && *input.PluralName != x.PluralName {
		x.PluralName = *input.PluralName
	}

	if input.Slug != nil && *input.Slug != x.Slug {
		x.Slug = *input.Slug
	}
}

var _ validation.ValidatableWithContext = (*ValidMeasurementUnitCreationRequestInput)(nil)

// ValidateWithContext validates a ValidMeasurementUnitCreationRequestInput.
func (x *ValidMeasurementUnitCreationRequestInput) ValidateWithContext(ctx context.Context) error {
	return validation.ValidateStructWithContext(
		ctx,
		x,
		validation.Field(&x.Name, validation.Required),
	)
}

var _ validation.ValidatableWithContext = (*ValidMeasurementUnitDatabaseCreationInput)(nil)

// ValidateWithContext validates a ValidMeasurementUnitDatabaseCreationInput.
func (x *ValidMeasurementUnitDatabaseCreationInput) ValidateWithContext(ctx context.Context) error {
	return validation.ValidateStructWithContext(
		ctx,
		x,
		validation.Field(&x.ID, validation.Required),
		validation.Field(&x.Name, validation.Required),
	)
}

var _ validation.ValidatableWithContext = (*ValidMeasurementUnitUpdateRequestInput)(nil)

// ValidateWithContext validates a ValidMeasurementUnitUpdateRequestInput.
func (x *ValidMeasurementUnitUpdateRequestInput) ValidateWithContext(ctx context.Context) error {
	return validation.ValidateStructWithContext(
		ctx,
		x,
		validation.Field(&x.Name, validation.Required),
	)
}
