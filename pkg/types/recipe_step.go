package types

import (
	"context"
	"encoding/gob"
	"errors"
	"net/http"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/hashicorp/go-multierror"
)

const (
	maxIngredientsPerStep = 100
	maxProductsPerStep    = 100

	// RecipeStepDataType indicates an event is related to a recipe step.
	RecipeStepDataType dataType = "recipe_step"

	// RecipeStepCreatedCustomerEventType indicates a recipe step was created.
	RecipeStepCreatedCustomerEventType CustomerEventType = "recipe_step_created"
	// RecipeStepUpdatedCustomerEventType indicates a recipe step was updated.
	RecipeStepUpdatedCustomerEventType CustomerEventType = "recipe_step_updated"
	// RecipeStepArchivedCustomerEventType indicates a recipe step was archived.
	RecipeStepArchivedCustomerEventType CustomerEventType = "recipe_step_archived"
)

func init() {
	gob.Register(new(RecipeStep))
	gob.Register(new(RecipeStepCreationRequestInput))
	gob.Register(new(RecipeStepUpdateRequestInput))
}

type (
	// RecipeStep represents a recipe step.
	RecipeStep struct {
		_ struct{}

		CreatedAt                     time.Time                        `json:"createdAt"`
		MinimumEstimatedTimeInSeconds *uint32                          `json:"minimumEstimatedTimeInSeconds"`
		ArchivedAt                    *time.Time                       `json:"archivedAt"`
		LastUpdatedAt                 *time.Time                       `json:"lastUpdatedAt"`
		MinimumTemperatureInCelsius   *float32                         `json:"minimumTemperatureInCelsius"`
		MaximumTemperatureInCelsius   *float32                         `json:"maximumTemperatureInCelsius"`
		MaximumEstimatedTimeInSeconds *uint32                          `json:"maximumEstimatedTimeInSeconds"`
		ExplicitInstructions          string                           `json:"explicitInstructions"`
		Notes                         string                           `json:"notes"`
		ID                            string                           `json:"id"`
		BelongsToRecipe               string                           `json:"belongsToRecipe"`
		ConditionExpression           string                           `json:"conditionExpression"`
		Products                      []*RecipeStepProduct             `json:"products"`
		Media                         []*RecipeMedia                   `json:"media"`
		Instruments                   []*RecipeStepInstrument          `json:"instruments"`
		Vessels                       []*RecipeStepVessel              `json:"vessels"`
		CompletionConditions          []*RecipeStepCompletionCondition `json:"completionConditions"`
		Ingredients                   []*RecipeStepIngredient          `json:"ingredients"`
		Preparation                   ValidPreparation                 `json:"preparation"`
		Index                         uint32                           `json:"index"`
		Optional                      bool                             `json:"optional"`
	}

	// RecipeStepList represents a list of recipe steps.
	RecipeStepList []*RecipeStep

	// RecipeStepCreationRequestInput represents what a user could set as input for creating recipe steps.
	RecipeStepCreationRequestInput struct {
		_ struct{}

		MaximumTemperatureInCelsius   *float32                                             `json:"maximumTemperatureInCelsius"`
		MinimumTemperatureInCelsius   *float32                                             `json:"minimumTemperatureInCelsius"`
		MaximumEstimatedTimeInSeconds *uint32                                              `json:"maximumEstimatedTimeInSeconds"`
		MinimumEstimatedTimeInSeconds *uint32                                              `json:"minimumEstimatedTimeInSeconds"`
		ConditionExpression           string                                               `json:"conditionExpression"`
		Notes                         string                                               `json:"notes"`
		PreparationID                 string                                               `json:"preparationID"`
		ExplicitInstructions          string                                               `json:"explicitInstructions"`
		Instruments                   []*RecipeStepInstrumentCreationRequestInput          `json:"instruments"`
		Vessels                       []*RecipeStepVesselCreationRequestInput              `json:"vessels"`
		Products                      []*RecipeStepProductCreationRequestInput             `json:"products"`
		Ingredients                   []*RecipeStepIngredientCreationRequestInput          `json:"ingredients"`
		CompletionConditions          []*RecipeStepCompletionConditionCreationRequestInput `json:"completionConditions"`
		Index                         uint32                                               `json:"index"`
		Optional                      bool                                                 `json:"optional"`
	}

	// RecipeStepDatabaseCreationInput represents what a user could set as input for creating recipe steps.
	RecipeStepDatabaseCreationInput struct {
		_ struct{}

		MinimumEstimatedTimeInSeconds *uint32
		MinimumTemperatureInCelsius   *float32
		MaximumEstimatedTimeInSeconds *uint32
		MaximumTemperatureInCelsius   *float32
		ConditionExpression           string
		PreparationID                 string
		ID                            string
		Notes                         string
		ExplicitInstructions          string
		BelongsToRecipe               string
		Ingredients                   []*RecipeStepIngredientDatabaseCreationInput
		Instruments                   []*RecipeStepInstrumentDatabaseCreationInput
		Vessels                       []*RecipeStepVesselDatabaseCreationInput
		Products                      []*RecipeStepProductDatabaseCreationInput
		CompletionConditions          []*RecipeStepCompletionConditionDatabaseCreationInput
		Index                         uint32
		Optional                      bool
	}

	// RecipeStepUpdateRequestInput represents what a user could set as input for updating recipe steps.
	RecipeStepUpdateRequestInput struct {
		_ struct{}

		MinimumTemperatureInCelsius   *float32          `json:"minimumTemperatureInCelsius,omitempty"`
		MaximumTemperatureInCelsius   *float32          `json:"maximumTemperatureInCelsius,omitempty"`
		Notes                         *string           `json:"notes,omitempty"`
		Preparation                   *ValidPreparation `json:"preparation,omitempty"`
		Index                         *uint32           `json:"index,omitempty"`
		MinimumEstimatedTimeInSeconds *uint32           `json:"minimumEstimatedTimeInSeconds,omitempty"`
		MaximumEstimatedTimeInSeconds *uint32           `json:"maximumEstimatedTimeInSeconds,omitempty"`
		Optional                      *bool             `json:"optional,omitempty"`
		ExplicitInstructions          *string           `json:"explicitInstructions,omitempty"`
		ConditionExpression           *string           `json:"conditionExpression,omitempty"`
		BelongsToRecipe               string            `json:"belongsToRecipe"`
	}

	// RecipeStepDataManager describes a structure capable of storing recipe steps permanently.
	RecipeStepDataManager interface {
		RecipeStepExists(ctx context.Context, recipeID, recipeStepID string) (bool, error)
		GetRecipeStep(ctx context.Context, recipeID, recipeStepID string) (*RecipeStep, error)
		GetRecipeSteps(ctx context.Context, recipeID string, filter *QueryFilter) (*QueryFilteredResult[RecipeStep], error)
		CreateRecipeStep(ctx context.Context, input *RecipeStepDatabaseCreationInput) (*RecipeStep, error)
		UpdateRecipeStep(ctx context.Context, updated *RecipeStep) error
		ArchiveRecipeStep(ctx context.Context, recipeID, recipeStepID string) error
	}

	// RecipeStepDataService describes a structure capable of serving traffic related to recipe steps.
	RecipeStepDataService interface {
		ListHandler(res http.ResponseWriter, req *http.Request)
		CreateHandler(res http.ResponseWriter, req *http.Request)
		ReadHandler(res http.ResponseWriter, req *http.Request)
		UpdateHandler(res http.ResponseWriter, req *http.Request)
		ArchiveHandler(res http.ResponseWriter, req *http.Request)
		ImageUploadHandler(res http.ResponseWriter, req *http.Request)
	}
)

// Update merges an RecipeStepUpdateRequestInput with a recipe step.
func (x *RecipeStep) Update(input *RecipeStepUpdateRequestInput) {
	if input.Index != nil && *input.Index != x.Index {
		x.Index = *input.Index
	}

	if input.Preparation.Name != "" && input.Preparation.Name != x.Preparation.Name {
		x.Preparation.Name = input.Preparation.Name
	}

	if input.Preparation.Description != "" && input.Preparation.Description != x.Preparation.Description {
		x.Preparation.Description = input.Preparation.Description
	}

	if input.Preparation.IconPath != "" && input.Preparation.IconPath != x.Preparation.IconPath {
		x.Preparation.IconPath = input.Preparation.IconPath
	}

	if input.MinimumEstimatedTimeInSeconds != nil && input.MinimumEstimatedTimeInSeconds != x.MinimumEstimatedTimeInSeconds {
		x.MinimumEstimatedTimeInSeconds = input.MinimumEstimatedTimeInSeconds
	}

	if input.MaximumEstimatedTimeInSeconds != nil && input.MaximumEstimatedTimeInSeconds != x.MaximumEstimatedTimeInSeconds {
		x.MaximumEstimatedTimeInSeconds = input.MaximumEstimatedTimeInSeconds
	}

	if input.MinimumTemperatureInCelsius != nil && (x.MinimumTemperatureInCelsius == nil || (*input.MinimumTemperatureInCelsius != 0 && *input.MinimumTemperatureInCelsius != *x.MinimumTemperatureInCelsius)) {
		x.MinimumTemperatureInCelsius = input.MinimumTemperatureInCelsius
	}

	if input.MaximumTemperatureInCelsius != nil && (x.MaximumTemperatureInCelsius == nil || (*input.MaximumTemperatureInCelsius != 0 && *input.MaximumTemperatureInCelsius != *x.MaximumTemperatureInCelsius)) {
		x.MaximumTemperatureInCelsius = input.MaximumTemperatureInCelsius
	}

	if input.Notes != nil && *input.Notes != x.Notes {
		x.Notes = *input.Notes
	}

	// TODO: do something about products here

	if input.Optional != nil && *input.Optional != x.Optional {
		x.Optional = *input.Optional
	}

	if input.MinimumTemperatureInCelsius != nil && (x.MinimumTemperatureInCelsius == nil || (*input.MinimumTemperatureInCelsius != 0 && *input.MinimumTemperatureInCelsius != *x.MinimumTemperatureInCelsius)) {
		x.MinimumTemperatureInCelsius = input.MinimumTemperatureInCelsius
	}

	if input.ExplicitInstructions != nil && *input.ExplicitInstructions != x.ExplicitInstructions {
		x.ExplicitInstructions = *input.ExplicitInstructions
	}

	if input.ConditionExpression != nil && *input.ConditionExpression != x.ConditionExpression {
		x.ConditionExpression = *input.ConditionExpression
	}
}

var _ validation.ValidatableWithContext = (*RecipeStepCreationRequestInput)(nil)

// ValidateWithContext validates a RecipeStepCreationRequestInput.
func (x *RecipeStepCreationRequestInput) ValidateWithContext(ctx context.Context) error {
	var err *multierror.Error

	if len(x.Instruments) == 0 && len(x.Vessels) == 0 {
		err = multierror.Append(err, errors.New("at least one instrument or vessel is required"))
	}

	validationErr := validation.ValidateStructWithContext(
		ctx,
		x,
		validation.Field(&x.PreparationID, validation.Required),
		validation.Field(&x.Products, validation.Required),
		validation.Field(&x.Instruments, validation.Required),
	)

	if validationErr != nil {
		err = multierror.Append(err, validationErr)
	}

	if err == nil {
		return nil
	}

	return err
}

var _ validation.ValidatableWithContext = (*RecipeStepDatabaseCreationInput)(nil)

// ValidateWithContext validates a RecipeStepDatabaseCreationInput.
func (x *RecipeStepDatabaseCreationInput) ValidateWithContext(ctx context.Context) error {
	return validation.ValidateStructWithContext(
		ctx,
		x,
		validation.Field(&x.ID, validation.Required),
		validation.Field(&x.Products, validation.Required),
		validation.Field(&x.PreparationID, validation.Required),
	)
}

var _ validation.ValidatableWithContext = (*RecipeStepUpdateRequestInput)(nil)

// ValidateWithContext validates a RecipeStepUpdateRequestInput.
func (x *RecipeStepUpdateRequestInput) ValidateWithContext(ctx context.Context) error {
	return validation.ValidateStructWithContext(
		ctx,
		x,
		validation.Field(&x.Index, validation.Required),
		validation.Field(&x.Preparation, validation.Required),
		validation.Field(&x.MinimumEstimatedTimeInSeconds, validation.Required),
		validation.Field(&x.MaximumEstimatedTimeInSeconds, validation.Required),
		validation.Field(&x.MinimumTemperatureInCelsius, validation.Required),
		validation.Field(&x.Notes, validation.Required),
	)
}
