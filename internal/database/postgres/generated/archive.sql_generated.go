// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: archive.sql

package generated

import (
	"context"
	"database/sql"
)

const archiveHousehold = `-- name: ArchiveHousehold :exec

UPDATE households SET last_updated_at = NOW(), archived_at = NOW() WHERE archived_at IS NULL AND belongs_to_user = $1 AND id = $2
`

type ArchiveHouseholdParams struct {
	BelongsToUser string
	ID            string
}

func (q *Queries) ArchiveHousehold(ctx context.Context, db DBTX, arg *ArchiveHouseholdParams) error {
	_, err := db.ExecContext(ctx, archiveHousehold, arg.BelongsToUser, arg.ID)
	return err
}

const archiveHouseholdInstrumentOwnership = `-- name: ArchiveHouseholdInstrumentOwnership :exec

UPDATE household_instrument_ownerships SET archived_at = NOW() WHERE archived_at IS NULL AND id = $1 AND belongs_to_household = $2
`

type ArchiveHouseholdInstrumentOwnershipParams struct {
	ID                 string
	BelongsToHousehold string
}

func (q *Queries) ArchiveHouseholdInstrumentOwnership(ctx context.Context, db DBTX, arg *ArchiveHouseholdInstrumentOwnershipParams) error {
	_, err := db.ExecContext(ctx, archiveHouseholdInstrumentOwnership, arg.ID, arg.BelongsToHousehold)
	return err
}

const archiveMeal = `-- name: ArchiveMeal :exec

UPDATE meals SET archived_at = NOW() WHERE archived_at IS NULL AND created_by_user = $1 AND id = $2
`

type ArchiveMealParams struct {
	CreatedByUser string
	ID            string
}

func (q *Queries) ArchiveMeal(ctx context.Context, db DBTX, arg *ArchiveMealParams) error {
	_, err := db.ExecContext(ctx, archiveMeal, arg.CreatedByUser, arg.ID)
	return err
}

const archiveMealPlan = `-- name: ArchiveMealPlan :exec

UPDATE meal_plans SET archived_at = NOW() WHERE archived_at IS NULL AND belongs_to_household = $1 AND id = $2
`

type ArchiveMealPlanParams struct {
	BelongsToHousehold string
	ID                 string
}

func (q *Queries) ArchiveMealPlan(ctx context.Context, db DBTX, arg *ArchiveMealPlanParams) error {
	_, err := db.ExecContext(ctx, archiveMealPlan, arg.BelongsToHousehold, arg.ID)
	return err
}

const archiveMealPlanEvent = `-- name: ArchiveMealPlanEvent :exec

UPDATE meal_plan_events SET archived_at = NOW() WHERE archived_at IS NULL AND id = $1 AND belongs_to_meal_plan = $2
`

type ArchiveMealPlanEventParams struct {
	ID                string
	BelongsToMealPlan string
}

func (q *Queries) ArchiveMealPlanEvent(ctx context.Context, db DBTX, arg *ArchiveMealPlanEventParams) error {
	_, err := db.ExecContext(ctx, archiveMealPlanEvent, arg.ID, arg.BelongsToMealPlan)
	return err
}

const archiveMealPlanGroceryListItem = `-- name: ArchiveMealPlanGroceryListItem :exec

UPDATE meal_plan_grocery_list_items SET archived_at = NOW() WHERE archived_at IS NULL AND id = $1
`

func (q *Queries) ArchiveMealPlanGroceryListItem(ctx context.Context, db DBTX, id string) error {
	_, err := db.ExecContext(ctx, archiveMealPlanGroceryListItem, id)
	return err
}

const archiveMealPlanOption = `-- name: ArchiveMealPlanOption :exec

UPDATE
	meal_plan_options
SET
	archived_at = NOW()
WHERE
	archived_at IS NULL
	AND belongs_to_meal_plan_event = $1
	AND id = $2
`

type ArchiveMealPlanOptionParams struct {
	ID                     string
	BelongsToMealPlanEvent sql.NullString
}

func (q *Queries) ArchiveMealPlanOption(ctx context.Context, db DBTX, arg *ArchiveMealPlanOptionParams) error {
	_, err := db.ExecContext(ctx, archiveMealPlanOption, arg.BelongsToMealPlanEvent, arg.ID)
	return err
}

const archiveMealPlanOptionVote = `-- name: ArchiveMealPlanOptionVote :exec

UPDATE meal_plan_option_votes SET archived_at = NOW() WHERE archived_at IS NULL AND belongs_to_meal_plan_option = $1 AND id = $2
`

type ArchiveMealPlanOptionVoteParams struct {
	BelongsToMealPlanOption string
	ID                      string
}

func (q *Queries) ArchiveMealPlanOptionVote(ctx context.Context, db DBTX, arg *ArchiveMealPlanOptionVoteParams) error {
	_, err := db.ExecContext(ctx, archiveMealPlanOptionVote, arg.BelongsToMealPlanOption, arg.ID)
	return err
}

const archiveOAuth2Client = `-- name: ArchiveOAuth2Client :exec

UPDATE oauth2_clients SET
	archived_at = NOW()
WHERE archived_at IS NULL
	AND id = $1
`

func (q *Queries) ArchiveOAuth2Client(ctx context.Context, db DBTX, id string) error {
	_, err := db.ExecContext(ctx, archiveOAuth2Client, id)
	return err
}

const archiveRecipe = `-- name: ArchiveRecipe :exec

UPDATE recipes SET archived_at = NOW() WHERE archived_at IS NULL AND created_by_user = $1 AND id = $2
`

type ArchiveRecipeParams struct {
	CreatedByUser string
	ID            string
}

func (q *Queries) ArchiveRecipe(ctx context.Context, db DBTX, arg *ArchiveRecipeParams) error {
	_, err := db.ExecContext(ctx, archiveRecipe, arg.CreatedByUser, arg.ID)
	return err
}

const archiveRecipeMedia = `-- name: ArchiveRecipeMedia :exec

UPDATE recipe_media SET archived_at = NOW() WHERE archived_at IS NULL AND id = $1
`

func (q *Queries) ArchiveRecipeMedia(ctx context.Context, db DBTX, id string) error {
	_, err := db.ExecContext(ctx, archiveRecipeMedia, id)
	return err
}

const archiveRecipePrepTask = `-- name: ArchiveRecipePrepTask :exec

UPDATE recipe_prep_tasks SET archived_at = NOW() WHERE archived_at IS NULL AND id = $1
`

func (q *Queries) ArchiveRecipePrepTask(ctx context.Context, db DBTX, id string) error {
	_, err := db.ExecContext(ctx, archiveRecipePrepTask, id)
	return err
}

const archiveRecipeRating = `-- name: ArchiveRecipeRating :exec

UPDATE recipe_ratings SET archived_at = NOW() WHERE archived_at IS NULL AND id = $1
`

func (q *Queries) ArchiveRecipeRating(ctx context.Context, db DBTX, id string) error {
	_, err := db.ExecContext(ctx, archiveRecipeRating, id)
	return err
}

const archiveRecipeStep = `-- name: ArchiveRecipeStep :exec

UPDATE recipe_steps SET archived_at = NOW() WHERE archived_at IS NULL AND belongs_to_recipe = $1 AND id = $2
`

type ArchiveRecipeStepParams struct {
	BelongsToRecipe string
	ID              string
}

func (q *Queries) ArchiveRecipeStep(ctx context.Context, db DBTX, arg *ArchiveRecipeStepParams) error {
	_, err := db.ExecContext(ctx, archiveRecipeStep, arg.BelongsToRecipe, arg.ID)
	return err
}

const archiveRecipeStepCompletionCondition = `-- name: ArchiveRecipeStepCompletionCondition :exec

UPDATE recipe_step_completion_conditions SET archived_at = NOW() WHERE archived_at IS NULL AND belongs_to_recipe_step = $1 AND id = $2
`

type ArchiveRecipeStepCompletionConditionParams struct {
	BelongsToRecipeStep string
	ID                  string
}

func (q *Queries) ArchiveRecipeStepCompletionCondition(ctx context.Context, db DBTX, arg *ArchiveRecipeStepCompletionConditionParams) error {
	_, err := db.ExecContext(ctx, archiveRecipeStepCompletionCondition, arg.BelongsToRecipeStep, arg.ID)
	return err
}

const archiveRecipeStepIngredient = `-- name: ArchiveRecipeStepIngredient :exec

UPDATE recipe_step_ingredients SET archived_at = NOW() WHERE archived_at IS NULL AND belongs_to_recipe_step = $1 AND id = $2
`

type ArchiveRecipeStepIngredientParams struct {
	BelongsToRecipeStep string
	ID                  string
}

func (q *Queries) ArchiveRecipeStepIngredient(ctx context.Context, db DBTX, arg *ArchiveRecipeStepIngredientParams) error {
	_, err := db.ExecContext(ctx, archiveRecipeStepIngredient, arg.BelongsToRecipeStep, arg.ID)
	return err
}

const archiveRecipeStepInstrument = `-- name: ArchiveRecipeStepInstrument :exec

UPDATE recipe_step_instruments SET archived_at = NOW() WHERE archived_at IS NULL AND belongs_to_recipe_step = $1 AND id = $2
`

type ArchiveRecipeStepInstrumentParams struct {
	BelongsToRecipeStep string
	ID                  string
}

func (q *Queries) ArchiveRecipeStepInstrument(ctx context.Context, db DBTX, arg *ArchiveRecipeStepInstrumentParams) error {
	_, err := db.ExecContext(ctx, archiveRecipeStepInstrument, arg.BelongsToRecipeStep, arg.ID)
	return err
}

const archiveRecipeStepProduct = `-- name: ArchiveRecipeStepProduct :exec

UPDATE recipe_step_products SET archived_at = NOW() WHERE archived_at IS NULL AND belongs_to_recipe_step = $1 AND id = $2
`

type ArchiveRecipeStepProductParams struct {
	BelongsToRecipeStep string
	ID                  string
}

func (q *Queries) ArchiveRecipeStepProduct(ctx context.Context, db DBTX, arg *ArchiveRecipeStepProductParams) error {
	_, err := db.ExecContext(ctx, archiveRecipeStepProduct, arg.BelongsToRecipeStep, arg.ID)
	return err
}

const archiveRecipeStepVessel = `-- name: ArchiveRecipeStepVessel :exec

UPDATE recipe_step_vessels SET archived_at = NOW() WHERE archived_at IS NULL AND belongs_to_recipe_step = $1 AND id = $2
`

type ArchiveRecipeStepVesselParams struct {
	BelongsToRecipeStep string
	ID                  string
}

func (q *Queries) ArchiveRecipeStepVessel(ctx context.Context, db DBTX, arg *ArchiveRecipeStepVesselParams) error {
	_, err := db.ExecContext(ctx, archiveRecipeStepVessel, arg.BelongsToRecipeStep, arg.ID)
	return err
}

const archiveServiceSetting = `-- name: ArchiveServiceSetting :exec

UPDATE service_settings
SET archived_at = NOW()
    WHERE id = $1
`

func (q *Queries) ArchiveServiceSetting(ctx context.Context, db DBTX, id string) error {
	_, err := db.ExecContext(ctx, archiveServiceSetting, id)
	return err
}

const archiveServiceSettingConfiguration = `-- name: ArchiveServiceSettingConfiguration :exec

UPDATE service_setting_configurations SET archived_at = NOW() WHERE archived_at IS NULL AND id = $1
`

func (q *Queries) ArchiveServiceSettingConfiguration(ctx context.Context, db DBTX, id string) error {
	_, err := db.ExecContext(ctx, archiveServiceSettingConfiguration, id)
	return err
}

const archiveUser = `-- name: ArchiveUser :exec

UPDATE users SET
	archived_at = NOW()
WHERE archived_at IS NULL
	AND id = $1
`

func (q *Queries) ArchiveUser(ctx context.Context, db DBTX, id string) error {
	_, err := db.ExecContext(ctx, archiveUser, id)
	return err
}

const archiveUserIngredientPreference = `-- name: ArchiveUserIngredientPreference :exec

UPDATE user_ingredient_preferences SET archived_at = NOW() WHERE archived_at IS NULL AND id = $1 AND belongs_to_user = $2
`

type ArchiveUserIngredientPreferenceParams struct {
	ID            string
	BelongsToUser string
}

func (q *Queries) ArchiveUserIngredientPreference(ctx context.Context, db DBTX, arg *ArchiveUserIngredientPreferenceParams) error {
	_, err := db.ExecContext(ctx, archiveUserIngredientPreference, arg.ID, arg.BelongsToUser)
	return err
}

const archiveValidIngredient = `-- name: ArchiveValidIngredient :exec

UPDATE valid_ingredients SET archived_at = NOW() WHERE archived_at IS NULL AND id = $1
`

func (q *Queries) ArchiveValidIngredient(ctx context.Context, db DBTX, id string) error {
	_, err := db.ExecContext(ctx, archiveValidIngredient, id)
	return err
}

const archiveValidIngredientGroup = `-- name: ArchiveValidIngredientGroup :exec

UPDATE valid_ingredient_groups SET archived_at = NOW() WHERE archived_at IS NULL AND id = $1
`

func (q *Queries) ArchiveValidIngredientGroup(ctx context.Context, db DBTX, id string) error {
	_, err := db.ExecContext(ctx, archiveValidIngredientGroup, id)
	return err
}

const archiveValidIngredientMeasurementUnit = `-- name: ArchiveValidIngredientMeasurementUnit :exec

UPDATE valid_ingredient_measurement_units SET archived_at = NOW() WHERE archived_at IS NULL AND id = $1
`

func (q *Queries) ArchiveValidIngredientMeasurementUnit(ctx context.Context, db DBTX, id string) error {
	_, err := db.ExecContext(ctx, archiveValidIngredientMeasurementUnit, id)
	return err
}

const archiveValidIngredientPreparation = `-- name: ArchiveValidIngredientPreparation :exec

UPDATE valid_ingredient_preparations SET archived_at = NOW() WHERE archived_at IS NULL AND id = $1
`

func (q *Queries) ArchiveValidIngredientPreparation(ctx context.Context, db DBTX, id string) error {
	_, err := db.ExecContext(ctx, archiveValidIngredientPreparation, id)
	return err
}

const archiveValidIngredientState = `-- name: ArchiveValidIngredientState :exec

UPDATE valid_ingredient_states SET archived_at = NOW() WHERE archived_at IS NULL AND id = $1
`

func (q *Queries) ArchiveValidIngredientState(ctx context.Context, db DBTX, id string) error {
	_, err := db.ExecContext(ctx, archiveValidIngredientState, id)
	return err
}

const archiveValidIngredientStateIngredient = `-- name: ArchiveValidIngredientStateIngredient :exec

UPDATE valid_ingredient_state_ingredients SET archived_at = NOW() WHERE archived_at IS NULL AND id = $1
`

func (q *Queries) ArchiveValidIngredientStateIngredient(ctx context.Context, db DBTX, id string) error {
	_, err := db.ExecContext(ctx, archiveValidIngredientStateIngredient, id)
	return err
}

const archiveValidInstrument = `-- name: ArchiveValidInstrument :exec

UPDATE valid_instruments SET archived_at = NOW() WHERE archived_at IS NULL AND id = $1
`

func (q *Queries) ArchiveValidInstrument(ctx context.Context, db DBTX, id string) error {
	_, err := db.ExecContext(ctx, archiveValidInstrument, id)
	return err
}

const archiveValidMeasurementUnitConversion = `-- name: ArchiveValidMeasurementUnitConversion :exec

UPDATE valid_measurement_conversions SET archived_at = NOW() WHERE archived_at IS NULL AND id = $1
`

func (q *Queries) ArchiveValidMeasurementUnitConversion(ctx context.Context, db DBTX, id string) error {
	_, err := db.ExecContext(ctx, archiveValidMeasurementUnitConversion, id)
	return err
}

const archiveValidMeasurementUnit = `-- name: ArchiveValidMeasurementUnit :exec

UPDATE valid_measurement_units SET archived_at = NOW() WHERE archived_at IS NULL AND id = $1
`

func (q *Queries) ArchiveValidMeasurementUnit(ctx context.Context, db DBTX, id string) error {
	_, err := db.ExecContext(ctx, archiveValidMeasurementUnit, id)
	return err
}

const archiveValidPreparation = `-- name: ArchiveValidPreparation :exec

UPDATE valid_preparations SET archived_at = NOW() WHERE archived_at IS NULL AND id = $1
`

func (q *Queries) ArchiveValidPreparation(ctx context.Context, db DBTX, id string) error {
	_, err := db.ExecContext(ctx, archiveValidPreparation, id)
	return err
}

const archiveValidPreparationInstrument = `-- name: ArchiveValidPreparationInstrument :exec

UPDATE valid_preparation_instruments SET archived_at = NOW() WHERE archived_at IS NULL AND id = $1
`

func (q *Queries) ArchiveValidPreparationInstrument(ctx context.Context, db DBTX, id string) error {
	_, err := db.ExecContext(ctx, archiveValidPreparationInstrument, id)
	return err
}

const archiveValidPreparationVessel = `-- name: ArchiveValidPreparationVessel :exec

UPDATE valid_preparation_vessels SET archived_at = NOW() WHERE archived_at IS NULL AND id = $1
`

func (q *Queries) ArchiveValidPreparationVessel(ctx context.Context, db DBTX, id string) error {
	_, err := db.ExecContext(ctx, archiveValidPreparationVessel, id)
	return err
}

const archiveValidVessel = `-- name: ArchiveValidVessel :exec

UPDATE valid_vessels SET archived_at = NOW() WHERE archived_at IS NULL AND id = $1
`

func (q *Queries) ArchiveValidVessel(ctx context.Context, db DBTX, id string) error {
	_, err := db.ExecContext(ctx, archiveValidVessel, id)
	return err
}

const archiveWebhook = `-- name: ArchiveWebhook :exec

UPDATE webhooks
SET
	last_updated_at = NOW(),
	archived_at = NOW()
WHERE archived_at IS NULL
	AND belongs_to_household = $1
	AND id = $2
`

type ArchiveWebhookParams struct {
	BelongsToHousehold string
	ID                 string
}

func (q *Queries) ArchiveWebhook(ctx context.Context, db DBTX, arg *ArchiveWebhookParams) error {
	_, err := db.ExecContext(ctx, archiveWebhook, arg.BelongsToHousehold, arg.ID)
	return err
}
