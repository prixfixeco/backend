// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package generated

import (
	"context"
)

type Querier interface {
	AddUserToHouseholdDuringCreation(ctx context.Context, arg *AddUserToHouseholdDuringCreationParams) error
	AddUserToHouseholdQuery(ctx context.Context, arg *AddUserToHouseholdQueryParams) error
	ArchiveAPIClient(ctx context.Context, arg *ArchiveAPIClientParams) error
	ArchiveHousehold(ctx context.Context, arg *ArchiveHouseholdParams) error
	ArchiveMeal(ctx context.Context, arg *ArchiveMealParams) error
	ArchiveMealPlan(ctx context.Context, arg *ArchiveMealPlanParams) error
	ArchiveMealPlanOption(ctx context.Context, arg *ArchiveMealPlanOptionParams) error
	ArchiveMealPlanOptionVote(ctx context.Context, arg *ArchiveMealPlanOptionVoteParams) error
	ArchiveMemberships(ctx context.Context, belongsToUser string) error
	ArchiveRecipe(ctx context.Context, arg *ArchiveRecipeParams) error
	ArchiveRecipeStep(ctx context.Context, arg *ArchiveRecipeStepParams) error
	ArchiveRecipeStepIngredient(ctx context.Context, arg *ArchiveRecipeStepIngredientParams) error
	ArchiveRecipeStepInstrument(ctx context.Context, arg *ArchiveRecipeStepInstrumentParams) error
	ArchiveRecipeStepProduct(ctx context.Context, arg *ArchiveRecipeStepProductParams) error
	ArchiveUser(ctx context.Context, id string) error
	ArchiveValidIngredient(ctx context.Context, id string) error
	ArchiveValidIngredientMeasurementUnit(ctx context.Context, id string) error
	ArchiveValidIngredientPreparation(ctx context.Context, id string) error
	ArchiveValidInstrument(ctx context.Context, id string) error
	ArchiveValidMeasurementUnit(ctx context.Context, id string) error
	ArchiveValidPreparation(ctx context.Context, id string) error
	ArchiveValidPreparationInstrument(ctx context.Context, id string) error
	ArchiveWebhook(ctx context.Context, arg *ArchiveWebhookParams) error
	AttachInvitationsToUserID(ctx context.Context, arg *AttachInvitationsToUserIDParams) error
	CreateAPIClient(ctx context.Context, arg *CreateAPIClientParams) error
	CreateHousehold(ctx context.Context, arg *CreateHouseholdParams) error
	CreateHouseholdInvitation(ctx context.Context, arg *CreateHouseholdInvitationParams) error
	CreateHouseholdMembershipForNewUser(ctx context.Context, arg *CreateHouseholdMembershipForNewUserParams) error
	CreateMeal(ctx context.Context, arg *CreateMealParams) error
	CreateMealPlan(ctx context.Context, arg *CreateMealPlanParams) error
	CreateMealRecipe(ctx context.Context, arg *CreateMealRecipeParams) error
	CreatePasswordResetToken(ctx context.Context, arg *CreatePasswordResetTokenParams) error
	CreateRecipe(ctx context.Context, arg *CreateRecipeParams) error
	CreateRecipeStep(ctx context.Context, arg *CreateRecipeStepParams) error
	CreateRecipeStepIngredient(ctx context.Context, arg *CreateRecipeStepIngredientParams) error
	CreateRecipeStepInstrument(ctx context.Context, arg *CreateRecipeStepInstrumentParams) error
	CreateRecipeStepProduct(ctx context.Context, arg *CreateRecipeStepProductParams) error
	CreateUser(ctx context.Context, arg *CreateUserParams) error
	CreateValidIngredient(ctx context.Context, arg *CreateValidIngredientParams) error
	CreateValidIngredientMeasurementUnit(ctx context.Context, arg *CreateValidIngredientMeasurementUnitParams) error
	CreateValidIngredientPreparation(ctx context.Context, arg *CreateValidIngredientPreparationParams) error
	CreateValidInstrument(ctx context.Context, arg *CreateValidInstrumentParams) error
	CreateValidMeasurementUnit(ctx context.Context, arg *CreateValidMeasurementUnitParams) error
	CreateValidPreparation(ctx context.Context, arg *CreateValidPreparationParams) error
	CreateValidPreparationInstrument(ctx context.Context, arg *CreateValidPreparationInstrumentParams) error
	CreateWebhook(ctx context.Context, arg *CreateWebhookParams) error
	FinalizeMealPlan(ctx context.Context, arg *FinalizeMealPlanParams) error
	FinalizeMealPlanOption(ctx context.Context, arg *FinalizeMealPlanOptionParams) error
	GetAPIClientByClientID(ctx context.Context, clientID string) (*GetAPIClientByClientIDRow, error)
	GetAPIClientByDatabaseID(ctx context.Context, arg *GetAPIClientByDatabaseIDParams) (*GetAPIClientByDatabaseIDRow, error)
	GetAdminUserByUsername(ctx context.Context, username string) (*GetAdminUserByUsernameRow, error)
	GetAllHouseholdInvitationsCount(ctx context.Context) (int64, error)
	GetAllHouseholdsCount(ctx context.Context) (int64, error)
	GetAllUsersCount(ctx context.Context) (int64, error)
	GetAllWebhooksCount(ctx context.Context) (int64, error)
	GetDefaultHouseholdIDForUserQuery(ctx context.Context, arg *GetDefaultHouseholdIDForUserQueryParams) (string, error)
	GetExpiredAndUnresolvedMealPlanIDs(ctx context.Context) ([]*MealPlans, error)
	GetHousehold(ctx context.Context, id string) ([]*GetHouseholdRow, error)
	GetHouseholdByID(ctx context.Context, id string) ([]*GetHouseholdByIDRow, error)
	GetHouseholdInvitationByEmailAndToken(ctx context.Context, arg *GetHouseholdInvitationByEmailAndTokenParams) (*GetHouseholdInvitationByEmailAndTokenRow, error)
	GetHouseholdInvitationByHouseholdAndID(ctx context.Context, arg *GetHouseholdInvitationByHouseholdAndIDParams) ([]*GetHouseholdInvitationByHouseholdAndIDRow, error)
	GetHouseholdInvitationByTokenAndID(ctx context.Context, arg *GetHouseholdInvitationByTokenAndIDParams) (*GetHouseholdInvitationByTokenAndIDRow, error)
	GetHouseholdMembershipsForUserQuery(ctx context.Context, belongsToUser string) ([]*GetHouseholdMembershipsForUserQueryRow, error)
	GetMealByID(ctx context.Context, id string) (*GetMealByIDRow, error)
	GetMealPlan(ctx context.Context, arg *GetMealPlanParams) ([]*GetMealPlanRow, error)
	GetMealPlanOptionQuery(ctx context.Context, arg *GetMealPlanOptionQueryParams) (*GetMealPlanOptionQueryRow, error)
	GetMealPlanOptionVote(ctx context.Context, arg *GetMealPlanOptionVoteParams) (*MealPlanOptionVotes, error)
	GetMealPlanPastVotingDeadline(ctx context.Context, arg *GetMealPlanPastVotingDeadlineParams) ([]*GetMealPlanPastVotingDeadlineRow, error)
	GetPasswordResetToken(ctx context.Context, token string) (*PasswordResetTokens, error)
	GetRandomValidIngredient(ctx context.Context) (*GetRandomValidIngredientRow, error)
	GetRandomValidInstrument(ctx context.Context) (*ValidInstruments, error)
	GetRandomValidMeasurementUnit(ctx context.Context) (*GetRandomValidMeasurementUnitRow, error)
	GetRandomValidPreparation(ctx context.Context) (*ValidPreparations, error)
	GetRecipeByID(ctx context.Context, id string) ([]*GetRecipeByIDRow, error)
	GetRecipeByIDAndAuthorID(ctx context.Context, arg *GetRecipeByIDAndAuthorIDParams) ([]*GetRecipeByIDAndAuthorIDRow, error)
	GetRecipeStep(ctx context.Context, arg *GetRecipeStepParams) (*GetRecipeStepRow, error)
	GetRecipeStepIngredient(ctx context.Context, arg *GetRecipeStepIngredientParams) ([]*GetRecipeStepIngredientRow, error)
	GetRecipeStepInstrument(ctx context.Context, arg *GetRecipeStepInstrumentParams) ([]*GetRecipeStepInstrumentRow, error)
	GetRecipeStepInstrumentsForRecipe(ctx context.Context, arg *GetRecipeStepInstrumentsForRecipeParams) ([]*GetRecipeStepInstrumentsForRecipeRow, error)
	GetRecipeStepProduct(ctx context.Context, arg *GetRecipeStepProductParams) ([]*GetRecipeStepProductRow, error)
	GetTotalAPIClientCount(ctx context.Context) error
	GetTotalMealPlanOptionVotesCount(ctx context.Context) (int64, error)
	GetTotalMealPlanOptionsCount(ctx context.Context) (int64, error)
	GetTotalMealPlansCount(ctx context.Context) (int64, error)
	GetTotalMealsCount(ctx context.Context) (int64, error)
	GetTotalRecipeStepInstrumentCount(ctx context.Context) (int64, error)
	GetTotalRecipeStepsCount(ctx context.Context) (int64, error)
	GetTotalRecipesCount(ctx context.Context) (int64, error)
	GetTotalValidIngredientCount(ctx context.Context) (int64, error)
	GetTotalValidIngredientMeasurementUnitsCount(ctx context.Context) (int64, error)
	GetTotalValidIngredientPreparationsCount(ctx context.Context) (int64, error)
	GetTotalValidInstrumentCount(ctx context.Context) (int64, error)
	GetTotalValidMeasurementUnitCount(ctx context.Context) (int64, error)
	GetTotalValidPreparationInstrumentsCount(ctx context.Context) (int64, error)
	GetTotalValidPreparationsCount(ctx context.Context) (int64, error)
	GetUserByID(ctx context.Context, id string) (*GetUserByIDRow, error)
	GetUserByUsername(ctx context.Context, username string) (*GetUserByUsernameRow, error)
	GetUserIDByEmail(ctx context.Context, emailAddress string) (*GetUserIDByEmailRow, error)
	GetUserWithVerified2FA(ctx context.Context, id string) (*GetUserWithVerified2FARow, error)
	GetValidIngredient(ctx context.Context, id string) (*GetValidIngredientRow, error)
	GetValidIngredientMeasurementUnit(ctx context.Context, id string) (*GetValidIngredientMeasurementUnitRow, error)
	GetValidIngredientPreparation(ctx context.Context, id string) (*GetValidIngredientPreparationRow, error)
	GetValidInstrument(ctx context.Context, id string) (*ValidInstruments, error)
	GetValidMeasurementUnit(ctx context.Context, id string) (*GetValidMeasurementUnitRow, error)
	GetValidPreparation(ctx context.Context, id string) (*ValidPreparations, error)
	GetValidPreparationInstrument(ctx context.Context, id string) (*GetValidPreparationInstrumentRow, error)
	GetWebhook(ctx context.Context, arg *GetWebhookParams) (*Webhooks, error)
	HouseholdInvitationExists(ctx context.Context, id string) (bool, error)
	MarkHouseholdAsUserDefaultQuery(ctx context.Context, arg *MarkHouseholdAsUserDefaultQueryParams) error
	MarkUserTwoFactorSecretAsVerified(ctx context.Context, arg *MarkUserTwoFactorSecretAsVerifiedParams) error
	MealExists(ctx context.Context, id string) (bool, error)
	MealPlanExists(ctx context.Context, id string) (bool, error)
	MealPlanOptionCreation(ctx context.Context, arg *MealPlanOptionCreationParams) error
	MealPlanOptionExists(ctx context.Context, arg *MealPlanOptionExistsParams) (bool, error)
	MealPlanOptionVoteCreation(ctx context.Context, arg *MealPlanOptionVoteCreationParams) error
	MealPlanOptionVoteExists(ctx context.Context, arg *MealPlanOptionVoteExistsParams) (bool, error)
	ModifyUserPermissionsQuery(ctx context.Context, arg *ModifyUserPermissionsQueryParams) error
	RecipeExists(ctx context.Context, id string) (bool, error)
	RecipeStepExists(ctx context.Context, arg *RecipeStepExistsParams) (bool, error)
	RecipeStepIngredientExists(ctx context.Context, arg *RecipeStepIngredientExistsParams) (bool, error)
	RecipeStepInstrumentExists(ctx context.Context, arg *RecipeStepInstrumentExistsParams) (bool, error)
	RecipeStepProductExists(ctx context.Context, arg *RecipeStepProductExistsParams) (bool, error)
	RecipeStepProductsForRecipeQuery(ctx context.Context, arg *RecipeStepProductsForRecipeQueryParams) ([]*RecipeStepProductsForRecipeQueryRow, error)
	RedeemPasswordResetToken(ctx context.Context, id string) error
	RemoveUserFromHouseholdQuery(ctx context.Context, arg *RemoveUserFromHouseholdQueryParams) error
	SearchForUserByUsername(ctx context.Context, username string) ([]*SearchForUserByUsernameRow, error)
	SearchForValidIngredients(ctx context.Context, name string) ([]*SearchForValidIngredientsRow, error)
	SearchForValidInstruments(ctx context.Context, name string) ([]*ValidInstruments, error)
	SearchForValidMeasurementUnits(ctx context.Context, name string) ([]*SearchForValidMeasurementUnitsRow, error)
	SetInvitationStatus(ctx context.Context, arg *SetInvitationStatusParams) error
	SetUserAccountStatus(ctx context.Context, arg *SetUserAccountStatusParams) error
	TotalRecipeStepIngredientCount(ctx context.Context) (int64, error)
	TotalRecipeStepProductCount(ctx context.Context) (int64, error)
	TransferHouseholdMembershipQuery(ctx context.Context, arg *TransferHouseholdMembershipQueryParams) error
	TransferHouseholdOwnershipQuery(ctx context.Context, arg *TransferHouseholdOwnershipQueryParams) error
	UpdateHousehold(ctx context.Context, arg *UpdateHouseholdParams) error
	UpdateMealPlan(ctx context.Context, arg *UpdateMealPlanParams) error
	UpdateMealPlanOption(ctx context.Context, arg *UpdateMealPlanOptionParams) error
	UpdateMealPlanOptionVote(ctx context.Context, arg *UpdateMealPlanOptionVoteParams) error
	UpdateRecipe(ctx context.Context, arg *UpdateRecipeParams) error
	UpdateRecipeStep(ctx context.Context, arg *UpdateRecipeStepParams) error
	UpdateRecipeStepIngredient(ctx context.Context, arg *UpdateRecipeStepIngredientParams) error
	UpdateRecipeStepInstrument(ctx context.Context, arg *UpdateRecipeStepInstrumentParams) error
	UpdateRecipeStepProduct(ctx context.Context, arg *UpdateRecipeStepProductParams) error
	UpdateUser(ctx context.Context, arg *UpdateUserParams) error
	UpdateUserPassword(ctx context.Context, arg *UpdateUserPasswordParams) error
	UpdateUserTwoFactorSecret(ctx context.Context, arg *UpdateUserTwoFactorSecretParams) error
	UpdateValidIngredient(ctx context.Context, arg *UpdateValidIngredientParams) error
	UpdateValidIngredientMeasurementUnit(ctx context.Context, arg *UpdateValidIngredientMeasurementUnitParams) error
	UpdateValidIngredientPreparation(ctx context.Context, arg *UpdateValidIngredientPreparationParams) error
	UpdateValidInstrument(ctx context.Context, arg *UpdateValidInstrumentParams) error
	UpdateValidMeasurementUnit(ctx context.Context, arg *UpdateValidMeasurementUnitParams) error
	UpdateValidPreparation(ctx context.Context, arg *UpdateValidPreparationParams) error
	UpdateValidPreparationInstrument(ctx context.Context, arg *UpdateValidPreparationInstrumentParams) error
	UserHasStatus(ctx context.Context, arg *UserHasStatusParams) (bool, error)
	UserIsMemberOfHouseholdQuery(ctx context.Context, arg *UserIsMemberOfHouseholdQueryParams) (bool, error)
	ValidIngredientExists(ctx context.Context, id string) (bool, error)
	ValidIngredientMeasurementUnitExists(ctx context.Context, id string) (bool, error)
	ValidIngredientPreparationExists(ctx context.Context, id string) (bool, error)
	ValidInstrumentExists(ctx context.Context, id string) (bool, error)
	ValidMeasurementUnitExists(ctx context.Context, id string) (bool, error)
	ValidPreparationExists(ctx context.Context, id string) (bool, error)
	ValidPreparationInstrumentExists(ctx context.Context, id string) (bool, error)
	ValidPreparationsSearch(ctx context.Context, name string) ([]*ValidPreparations, error)
	WebhookExists(ctx context.Context, arg *WebhookExistsParams) (bool, error)
}

var _ Querier = (*Queries)(nil)
