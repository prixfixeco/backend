package types

import (
	"context"
	"encoding/gob"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

const (
	// RecipeStepProductDataType indicates an event is related to a recipe step product.
	RecipeStepProductDataType dataType = "recipe_step_product"

	// RecipeStepProductIngredientType represents one of the valid recipe step product type values.
	RecipeStepProductIngredientType = "ingredient"
	// RecipeStepProductInstrumentType represents one of the valid recipe step product type values.
	RecipeStepProductInstrumentType = "instrument"
	// RecipeStepProductVesselType represents one of the valid recipe step product type values.
	RecipeStepProductVesselType = "vessel"

	// RecipeStepProductCreatedCustomerEventType indicates a recipe step product was created.
	RecipeStepProductCreatedCustomerEventType CustomerEventType = "recipe_step_product_created"
	// RecipeStepProductUpdatedCustomerEventType indicates a recipe step product was updated.
	RecipeStepProductUpdatedCustomerEventType CustomerEventType = "recipe_step_product_updated"
	// RecipeStepProductArchivedCustomerEventType indicates a recipe step product was archived.
	RecipeStepProductArchivedCustomerEventType CustomerEventType = "recipe_step_product_archived"
)

func init() {
	gob.Register(new(RecipeStepProduct))
	gob.Register(new(RecipeStepProductCreationRequestInput))
	gob.Register(new(RecipeStepProductUpdateRequestInput))
}

type (
	// RecipeStepProduct represents a recipe step product.
	RecipeStepProduct struct {
		_ struct{}

		CreatedAt                          time.Time            `json:"createdAt"`
		MaximumStorageDurationInSeconds    *uint32              `json:"maximumStorageDurationInSeconds"`
		MinimumStorageTemperatureInCelsius *float32             `json:"minimumStorageTemperatureInCelsius"`
		MaximumStorageTemperatureInCelsius *float32             `json:"maximumStorageTemperatureInCelsius"`
		ArchivedAt                         *time.Time           `json:"archivedAt"`
		LastUpdatedAt                      *time.Time           `json:"lastUpdatedAt"`
		ID                                 string               `json:"id"`
		Type                               string               `json:"type"`
		Name                               string               `json:"name"`
		StorageInstructions                string               `json:"storageInstructions"`
		QuantityNotes                      string               `json:"quantityNotes"`
		BelongsToRecipeStep                string               `json:"belongsToRecipeStep"`
		MeasurementUnit                    ValidMeasurementUnit `json:"measurementUnit"`
		MaximumQuantity                    float32              `json:"maximumQuantity"`
		MinimumQuantity                    float32              `json:"minimumQuantity"`
		IsWaste                            bool                 `json:"isWaste"`
		IsLiquid                           bool                 `json:"isLiquid"`
		Index                              uint16               `json:"index"`
		ContainedInVesselIndex             uint16               `json:"containedInVesselIndex"`
		Compostable                        bool                 `json:"compostable"`
	}

	// RecipeStepProductCreationRequestInput represents what a user could set as input for creating recipe step products.
	RecipeStepProductCreationRequestInput struct {
		_ struct{}

		MinimumStorageTemperatureInCelsius *float32 `json:"minimumStorageTemperatureInCelsius"`
		MaximumStorageTemperatureInCelsius *float32 `json:"maximumStorageTemperatureInCelsius"`
		MaximumStorageDurationInSeconds    *uint32  `json:"maximumStorageDurationInSeconds"`
		Type                               string   `json:"type"`
		MeasurementUnitID                  string   `json:"measurementUnitID"`
		QuantityNotes                      string   `json:"quantityNotes"`
		Name                               string   `json:"name"`
		StorageInstructions                string   `json:"storageInstructions"`
		MaximumQuantity                    float32  `json:"maximumQuantity"`
		MinimumQuantity                    float32  `json:"minimumQuantity"`
		Compostable                        bool     `json:"compostable"`
		IsLiquid                           bool     `json:"isLiquid"`
		IsWaste                            bool     `json:"isWaste"`
		Index                              uint16   `json:"index"`
		ContainedInVesselIndex             uint16   `json:"containedInVesselIndex"`
	}

	// RecipeStepProductDatabaseCreationInput represents what a user could set as input for creating recipe step products.
	RecipeStepProductDatabaseCreationInput struct {
		_                                  struct{}
		MinimumStorageTemperatureInCelsius *float32
		MaximumStorageTemperatureInCelsius *float32
		MaximumStorageDurationInSeconds    *uint32
		Type                               string
		QuantityNotes                      string
		BelongsToRecipeStep                string
		MeasurementUnitID                  string
		Name                               string
		ID                                 string
		StorageInstructions                string
		MaximumQuantity                    float32
		MinimumQuantity                    float32
		Compostable                        bool
		IsLiquid                           bool
		IsWaste                            bool
		Index                              uint16
		ContainedInVesselIndex             uint16
	}

	// RecipeStepProductUpdateRequestInput represents what a user could set as input for updating recipe step products.
	RecipeStepProductUpdateRequestInput struct {
		_                                  struct{}
		Name                               *string  `json:"name"`
		Type                               *string  `json:"type"`
		MeasurementUnitID                  *string  `json:"measurementUnitID"`
		QuantityNotes                      *string  `json:"quantityNotes"`
		BelongsToRecipeStep                *string  `json:"belongsToRecipeStep"`
		MinimumQuantity                    *float32 `json:"minimumQuantity"`
		MaximumQuantity                    *float32 `json:"maximumQuantity"`
		Compostable                        *bool    `json:"compostable"`
		MaximumStorageDurationInSeconds    *uint32  `json:"maximumStorageDurationInSeconds"`
		MinimumStorageTemperatureInCelsius *float32 `json:"minimumStorageTemperatureInCelsius"`
		MaximumStorageTemperatureInCelsius *float32 `json:"maximumStorageTemperatureInCelsius"`
		StorageInstructions                *string  `json:"storageInstructions"`
		IsLiquid                           *bool    `json:"isLiquid"`
		IsWaste                            *bool    `json:"isWaste"`
		Index                              *uint16  `json:"index"`
		ContainedInVesselIndex             *uint16  `json:"containedInVesselIndex"`
	}

	// RecipeStepProductDataManager describes a structure capable of storing recipe step products permanently.
	RecipeStepProductDataManager interface {
		RecipeStepProductExists(ctx context.Context, recipeID, recipeStepID, recipeStepProductID string) (bool, error)
		GetRecipeStepProduct(ctx context.Context, recipeID, recipeStepID, recipeStepProductID string) (*RecipeStepProduct, error)
		GetRecipeStepProducts(ctx context.Context, recipeID, recipeStepID string, filter *QueryFilter) (*QueryFilteredResult[RecipeStepProduct], error)
		CreateRecipeStepProduct(ctx context.Context, input *RecipeStepProductDatabaseCreationInput) (*RecipeStepProduct, error)
		UpdateRecipeStepProduct(ctx context.Context, updated *RecipeStepProduct) error
		ArchiveRecipeStepProduct(ctx context.Context, recipeStepID, recipeStepProductID string) error
	}
)

// Update merges an RecipeStepProductUpdateRequestInput with a recipe step product.
func (x *RecipeStepProduct) Update(input *RecipeStepProductUpdateRequestInput) {
	if input.Name != nil && *input.Name != x.Name {
		x.Name = *input.Name
	}

	if input.Type != nil && *input.Type != x.Type {
		x.Type = *input.Type
	}

	if input.MeasurementUnitID != nil && *input.MeasurementUnitID != x.MeasurementUnit.ID {
		x.MeasurementUnit = ValidMeasurementUnit{ID: *input.MeasurementUnitID}
	}

	if input.MinimumQuantity != nil && *input.MinimumQuantity != x.MinimumQuantity {
		x.MinimumQuantity = *input.MinimumQuantity
	}

	if input.MaximumQuantity != nil && *input.MaximumQuantity != x.MaximumQuantity {
		x.MaximumQuantity = *input.MaximumQuantity
	}

	if input.QuantityNotes != nil && *input.QuantityNotes != x.QuantityNotes {
		x.QuantityNotes = *input.QuantityNotes
	}

	if input.Compostable != nil && *input.Compostable != x.Compostable {
		x.Compostable = *input.Compostable
	}

	if input.MaximumStorageDurationInSeconds != nil && input.MaximumStorageDurationInSeconds != x.MaximumStorageDurationInSeconds {
		x.MaximumStorageDurationInSeconds = input.MaximumStorageDurationInSeconds
	}

	if input.MinimumStorageTemperatureInCelsius != nil && input.MinimumStorageTemperatureInCelsius != x.MinimumStorageTemperatureInCelsius {
		x.MinimumStorageTemperatureInCelsius = input.MinimumStorageTemperatureInCelsius
	}

	if input.MaximumStorageTemperatureInCelsius != nil && input.MaximumStorageTemperatureInCelsius != x.MaximumStorageTemperatureInCelsius {
		x.MaximumStorageTemperatureInCelsius = input.MaximumStorageTemperatureInCelsius
	}

	if input.StorageInstructions != nil && *input.StorageInstructions != x.StorageInstructions {
		x.StorageInstructions = *input.StorageInstructions
	}

	if input.IsLiquid != nil && *input.IsLiquid != x.IsLiquid {
		x.IsLiquid = *input.IsLiquid
	}

	if input.IsWaste != nil && *input.IsWaste != x.IsWaste {
		x.IsWaste = *input.IsWaste
	}

	if input.Index != nil && *input.Index != x.Index {
		x.Index = *input.Index
	}

	if input.ContainedInVesselIndex != nil && *input.ContainedInVesselIndex != x.ContainedInVesselIndex {
		x.ContainedInVesselIndex = *input.ContainedInVesselIndex
	}
}

var _ validation.ValidatableWithContext = (*RecipeStepProductCreationRequestInput)(nil)

// ValidateWithContext validates a RecipeStepProductCreationRequestInput.
func (x *RecipeStepProductCreationRequestInput) ValidateWithContext(ctx context.Context) error {
	return validation.ValidateStructWithContext(
		ctx,
		x,
		validation.Field(&x.Name, validation.Required),
		validation.Field(&x.Type, validation.In(RecipeStepProductIngredientType, RecipeStepProductInstrumentType, RecipeStepProductVesselType)),
		validation.Field(&x.MinimumQuantity, validation.Required),
	)
}

var _ validation.ValidatableWithContext = (*RecipeStepProductDatabaseCreationInput)(nil)

// ValidateWithContext validates a RecipeStepProductDatabaseCreationInput.
func (x *RecipeStepProductDatabaseCreationInput) ValidateWithContext(ctx context.Context) error {
	return validation.ValidateStructWithContext(
		ctx,
		x,
		validation.Field(&x.ID, validation.Required),
		validation.Field(&x.Name, validation.Required),
		validation.Field(&x.Type, validation.In(RecipeStepProductIngredientType, RecipeStepProductInstrumentType, RecipeStepProductVesselType)),
		validation.Field(&x.MinimumQuantity, validation.Required),
		validation.Field(&x.MeasurementUnitID, validation.Required),
	)
}

var _ validation.ValidatableWithContext = (*RecipeStepProductUpdateRequestInput)(nil)

// ValidateWithContext validates a RecipeStepProductUpdateRequestInput.
func (x *RecipeStepProductUpdateRequestInput) ValidateWithContext(ctx context.Context) error {
	return validation.ValidateStructWithContext(
		ctx,
		x,
		validation.Field(&x.Name, validation.Required),
		validation.Field(&x.Type, validation.In(RecipeStepProductIngredientType, RecipeStepProductInstrumentType, RecipeStepProductVesselType)),
		validation.Field(&x.MeasurementUnitID, validation.Required),
		validation.Field(&x.MinimumQuantity, validation.Required),
		validation.Field(&x.MaximumQuantity, validation.Required),
	)
}
