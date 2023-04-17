package types

import (
	"context"
	"encoding/gob"
	"errors"
	"net/http"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

const (
	// MealPlanEventDataType indicates an event is related to a meal plan event.
	MealPlanEventDataType dataType = "meal_plan_event"

	// MealPlanEventCreatedCustomerEventType indicates a meal plan was created.
	MealPlanEventCreatedCustomerEventType CustomerEventType = "meal_plan_event_created"
	// MealPlanEventUpdatedCustomerEventType indicates a meal plan was updated.
	MealPlanEventUpdatedCustomerEventType CustomerEventType = "meal_plan_event_updated"
	// MealPlanEventArchivedCustomerEventType indicates a meal plan was archived.
	MealPlanEventArchivedCustomerEventType CustomerEventType = "meal_plan_event_archived"

	// BreakfastMealName represents breakfast.
	BreakfastMealName = "breakfast"
	// SecondBreakfastMealName represents second breakfast.
	SecondBreakfastMealName = "second_breakfast"
	// BrunchMealName represents brunch.
	BrunchMealName = "brunch"
	// LunchMealName represents lunch.
	LunchMealName = "lunch"
	// SupperMealName represents supper.
	SupperMealName = "supper"
	// DinnerMealName represents dinner.
	DinnerMealName = "dinner"
)

func init() {
	gob.Register(new(MealPlanEvent))
	gob.Register(new(MealPlanEventCreationRequestInput))
	gob.Register(new(MealPlanEventUpdateRequestInput))
}

type (
	// MealPlanEvent represents a meal plan.
	MealPlanEvent struct {
		_ struct{}

		CreatedAt         time.Time         `json:"createdAt"`
		StartsAt          time.Time         `json:"startsAt"`
		EndsAt            time.Time         `json:"endsAt"`
		ArchivedAt        *time.Time        `json:"archivedAt"`
		LastUpdatedAt     *time.Time        `json:"lastUpdatedAt"`
		MealName          string            `json:"mealName"`
		Notes             string            `json:"notes"`
		BelongsToMealPlan string            `json:"belongsToMealPlan"`
		ID                string            `json:"id"`
		Options           []*MealPlanOption `json:"options"`
	}

	// MealPlanEventCreationRequestInput represents what a user could set as input for creating meal plans.
	MealPlanEventCreationRequestInput struct {
		_ struct{}

		EndsAt   time.Time                             `json:"endsAt"`
		StartsAt time.Time                             `json:"startsAt"`
		Notes    string                                `json:"notes"`
		MealName string                                `json:"mealName"`
		Options  []*MealPlanOptionCreationRequestInput `json:"options"`
	}

	// MealPlanEventDatabaseCreationInput represents what a user could set as input for creating meal plans.
	MealPlanEventDatabaseCreationInput struct {
		_ struct{}

		StartsAt          time.Time
		EndsAt            time.Time
		BelongsToMealPlan string
		Notes             string
		MealName          string
		ID                string
		Options           []*MealPlanOptionDatabaseCreationInput
	}

	// MealPlanEventUpdateRequestInput represents what a user could set as input for updating meal plans.
	MealPlanEventUpdateRequestInput struct {
		_ struct{}

		Notes             *string    `json:"notes,omitempty"`
		StartsAt          *time.Time `json:"startsAt,omitempty"`
		MealName          *string    `json:"mealName,omitempty"`
		EndsAt            *time.Time `json:"endsAt,omitempty"`
		BelongsToMealPlan string     `json:"-"`
	}

	// MealPlanEventDataManager describes a structure capable of storing meal plans permanently.
	MealPlanEventDataManager interface {
		MealPlanEventExists(ctx context.Context, mealPlanID, mealPlanEventID string) (bool, error)
		GetMealPlanEvent(ctx context.Context, mealPlanID, mealPlanEventID string) (*MealPlanEvent, error)
		GetMealPlanEvents(ctx context.Context, mealPlanID string, filter *QueryFilter) (*QueryFilteredResult[MealPlanEvent], error)
		MealPlanEventIsEligibleForVoting(ctx context.Context, mealPlanID, mealPlanEventID string) (bool, error)
		CreateMealPlanEvent(ctx context.Context, input *MealPlanEventDatabaseCreationInput) (*MealPlanEvent, error)
		UpdateMealPlanEvent(ctx context.Context, updated *MealPlanEvent) error
		ArchiveMealPlanEvent(ctx context.Context, mealPlanID, mealPlanEventID string) error
	}

	// MealPlanEventDataService describes a structure capable of serving traffic related to meal plans.
	MealPlanEventDataService interface {
		ListHandler(res http.ResponseWriter, req *http.Request)
		CreateHandler(res http.ResponseWriter, req *http.Request)
		ReadHandler(res http.ResponseWriter, req *http.Request)
		UpdateHandler(res http.ResponseWriter, req *http.Request)
		ArchiveHandler(res http.ResponseWriter, req *http.Request)
	}
)

// Update merges an MealPlanEventUpdateRequestInput with a meal plan.
func (x *MealPlanEvent) Update(input *MealPlanEventUpdateRequestInput) {
	if input.Notes != nil && *input.Notes != x.Notes {
		x.Notes = *input.Notes
	}

	if input.StartsAt != nil && *input.StartsAt != x.StartsAt {
		x.StartsAt = *input.StartsAt
	}

	if input.EndsAt != nil && *input.EndsAt != x.EndsAt {
		x.EndsAt = *input.EndsAt
	}

	if input.MealName != nil && *input.MealName != x.MealName {
		x.MealName = *input.MealName
	}
}

var _ validation.ValidatableWithContext = (*MealPlanEventCreationRequestInput)(nil)

var errStartsAfterItEnds = errors.New("invalid start and end dates")

// ValidateWithContext validates a MealPlanEventCreationRequestInput.
func (x *MealPlanEventCreationRequestInput) ValidateWithContext(ctx context.Context) error {
	if x.StartsAt == x.EndsAt || x.StartsAt.After(x.EndsAt) {
		return errStartsAfterItEnds
	}

	mealCounts := map[string]uint{}
	for _, option := range x.Options {
		if _, ok := mealCounts[option.MealID]; !ok {
			mealCounts[option.MealID] = 1
		} else {
			mealCounts[option.MealID]++
			if mealCounts[option.MealID] >= 3 {
				return errTooFewUniqueMeals
			}
		}
	}

	return validation.ValidateStructWithContext(
		ctx,
		x,
		validation.Field(&x.StartsAt, validation.Required),
		validation.Field(&x.EndsAt, validation.Required),
		validation.Field(&x.MealName, validation.Required),
	)
}

var _ validation.ValidatableWithContext = (*MealPlanEventDatabaseCreationInput)(nil)

// ValidateWithContext validates a MealPlanEventDatabaseCreationInput.
func (x *MealPlanEventDatabaseCreationInput) ValidateWithContext(ctx context.Context) error {
	return validation.ValidateStructWithContext(
		ctx,
		x,
		validation.Field(&x.ID, validation.Required),
		validation.Field(&x.StartsAt, validation.Required),
		validation.Field(&x.EndsAt, validation.Required),
		validation.Field(&x.MealName, validation.In(
			BreakfastMealName,
			SecondBreakfastMealName,
			BrunchMealName,
			LunchMealName,
			SupperMealName,
			DinnerMealName,
		)),
		validation.Field(&x.BelongsToMealPlan, validation.Required),
	)
}

var _ validation.ValidatableWithContext = (*MealPlanEventUpdateRequestInput)(nil)

// ValidateWithContext validates a MealPlanEventUpdateRequestInput.
func (x *MealPlanEventUpdateRequestInput) ValidateWithContext(ctx context.Context) error {
	return validation.ValidateStructWithContext(
		ctx,
		x,
		validation.Field(&x.StartsAt, validation.Required),
		validation.Field(&x.EndsAt, validation.Required),
		validation.Field(&x.MealName, validation.In(
			BreakfastMealName,
			SecondBreakfastMealName,
			BrunchMealName,
			LunchMealName,
			SupperMealName,
			DinnerMealName,
		)),
	)
}
