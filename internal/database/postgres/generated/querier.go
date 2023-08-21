// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package generated

import (
	"context"
	"database/sql"
)

type Querier interface {
	AcceptPrivacyPolicyForUser(ctx context.Context, db DBTX, id string) error
	AcceptTermsOfServiceForUser(ctx context.Context, db DBTX, id string) error
	AddToHouseholdDuringCreation(ctx context.Context, db DBTX, arg *AddToHouseholdDuringCreationParams) error
	AddUserToHousehold(ctx context.Context, db DBTX, arg *AddUserToHouseholdParams) error
	ArchiveHousehold(ctx context.Context, db DBTX, arg *ArchiveHouseholdParams) error
	ArchiveHouseholdInstrumentOwnership(ctx context.Context, db DBTX, arg *ArchiveHouseholdInstrumentOwnershipParams) error
	ArchiveMeal(ctx context.Context, db DBTX, arg *ArchiveMealParams) error
	ArchiveMealPlan(ctx context.Context, db DBTX, arg *ArchiveMealPlanParams) error
	ArchiveMealPlanEvent(ctx context.Context, db DBTX, arg *ArchiveMealPlanEventParams) error
	ArchiveMealPlanGroceryListItem(ctx context.Context, db DBTX, id string) error
	ArchiveMealPlanOption(ctx context.Context, db DBTX, arg *ArchiveMealPlanOptionParams) error
	ArchiveMealPlanOptionVote(ctx context.Context, db DBTX, arg *ArchiveMealPlanOptionVoteParams) error
	ArchiveOAuth2Client(ctx context.Context, db DBTX, id string) error
	ArchiveOAuth2ClientTokenByAccess(ctx context.Context, db DBTX, access string) error
	ArchiveOAuth2ClientTokenByCode(ctx context.Context, db DBTX, code string) error
	ArchiveOAuth2ClientTokenByRefresh(ctx context.Context, db DBTX, refresh string) error
	ArchiveRecipe(ctx context.Context, db DBTX, arg *ArchiveRecipeParams) error
	ArchiveRecipeMedia(ctx context.Context, db DBTX, id string) error
	ArchiveRecipePrepTask(ctx context.Context, db DBTX, id string) error
	ArchiveRecipeRating(ctx context.Context, db DBTX, id string) error
	ArchiveRecipeStep(ctx context.Context, db DBTX, arg *ArchiveRecipeStepParams) error
	ArchiveRecipeStepCompletionCondition(ctx context.Context, db DBTX, arg *ArchiveRecipeStepCompletionConditionParams) error
	ArchiveRecipeStepIngredient(ctx context.Context, db DBTX, arg *ArchiveRecipeStepIngredientParams) error
	ArchiveRecipeStepInstrument(ctx context.Context, db DBTX, arg *ArchiveRecipeStepInstrumentParams) error
	ArchiveRecipeStepProduct(ctx context.Context, db DBTX, arg *ArchiveRecipeStepProductParams) error
	ArchiveRecipeStepVessel(ctx context.Context, db DBTX, arg *ArchiveRecipeStepVesselParams) error
	ArchiveServiceSetting(ctx context.Context, db DBTX, id string) error
	ArchiveServiceSettingConfiguration(ctx context.Context, db DBTX, id string) error
	ArchiveUser(ctx context.Context, db DBTX, id string) error
	ArchiveUserIngredientPreference(ctx context.Context, db DBTX, arg *ArchiveUserIngredientPreferenceParams) error
	ArchiveUserMemberships(ctx context.Context, db DBTX, belongsToUser string) error
	ArchiveValidIngredient(ctx context.Context, db DBTX, id string) error
	ArchiveValidIngredientGroup(ctx context.Context, db DBTX, id string) error
	ArchiveValidIngredientGroupMember(ctx context.Context, db DBTX, arg *ArchiveValidIngredientGroupMemberParams) error
	ArchiveValidIngredientMeasurementUnit(ctx context.Context, db DBTX, id string) error
	ArchiveValidIngredientPreparation(ctx context.Context, db DBTX, id string) error
	ArchiveValidIngredientState(ctx context.Context, db DBTX, id string) error
	ArchiveValidIngredientStateIngredient(ctx context.Context, db DBTX, id string) error
	ArchiveValidInstrument(ctx context.Context, db DBTX, id string) error
	ArchiveValidMeasurementUnit(ctx context.Context, db DBTX, id string) error
	ArchiveValidMeasurementUnitConversion(ctx context.Context, db DBTX, id string) error
	ArchiveValidPreparation(ctx context.Context, db DBTX, id string) error
	ArchiveValidPreparationInstrument(ctx context.Context, db DBTX, id string) error
	ArchiveValidPreparationVessel(ctx context.Context, db DBTX, id string) error
	ArchiveValidVessel(ctx context.Context, db DBTX, id string) error
	ArchiveWebhook(ctx context.Context, db DBTX, arg *ArchiveWebhookParams) error
	AttachHouseholdInvitationsToUserID(ctx context.Context, db DBTX, arg *AttachHouseholdInvitationsToUserIDParams) error
	ChangeMealPlanTaskStatus(ctx context.Context, db DBTX, arg *ChangeMealPlanTaskStatusParams) error
	CheckHouseholdInstrumentOwnershipExistence(ctx context.Context, db DBTX, arg *CheckHouseholdInstrumentOwnershipExistenceParams) (bool, error)
	CheckHouseholdInvitationExistence(ctx context.Context, db DBTX, id string) (bool, error)
	CheckMealExistence(ctx context.Context, db DBTX, id string) (bool, error)
	CheckMealPlanEventExistence(ctx context.Context, db DBTX, arg *CheckMealPlanEventExistenceParams) (bool, error)
	CheckMealPlanExistence(ctx context.Context, db DBTX, arg *CheckMealPlanExistenceParams) (bool, error)
	CheckMealPlanGroceryListItemExistence(ctx context.Context, db DBTX, arg *CheckMealPlanGroceryListItemExistenceParams) (bool, error)
	CheckMealPlanOptionExistence(ctx context.Context, db DBTX, arg *CheckMealPlanOptionExistenceParams) (bool, error)
	CheckMealPlanOptionVoteExistence(ctx context.Context, db DBTX, arg *CheckMealPlanOptionVoteExistenceParams) (bool, error)
	CheckMealPlanTaskExistence(ctx context.Context, db DBTX, arg *CheckMealPlanTaskExistenceParams) (bool, error)
	CheckOAuth2ClientTokenExistence(ctx context.Context, db DBTX, id string) (bool, error)
	CheckRecipeExistence(ctx context.Context, db DBTX, id string) (bool, error)
	CheckRecipeMediaExistence(ctx context.Context, db DBTX, id string) (bool, error)
	CheckRecipePrepTaskExistence(ctx context.Context, db DBTX, arg *CheckRecipePrepTaskExistenceParams) (bool, error)
	CheckRecipeRatingExistence(ctx context.Context, db DBTX, id string) (bool, error)
	CheckRecipeStepCompletionConditionExistence(ctx context.Context, db DBTX, arg *CheckRecipeStepCompletionConditionExistenceParams) (bool, error)
	CheckRecipeStepExistence(ctx context.Context, db DBTX, arg *CheckRecipeStepExistenceParams) (bool, error)
	CheckRecipeStepIngredientExistence(ctx context.Context, db DBTX, arg *CheckRecipeStepIngredientExistenceParams) (bool, error)
	CheckRecipeStepInstrumentExistence(ctx context.Context, db DBTX, arg *CheckRecipeStepInstrumentExistenceParams) (bool, error)
	CheckRecipeStepProductExistence(ctx context.Context, db DBTX, arg *CheckRecipeStepProductExistenceParams) (bool, error)
	CheckRecipeStepVesselExistence(ctx context.Context, db DBTX, arg *CheckRecipeStepVesselExistenceParams) (bool, error)
	CheckServiceSettingConfigurationExistence(ctx context.Context, db DBTX, id string) (bool, error)
	CheckServiceSettingExistence(ctx context.Context, db DBTX, id string) (bool, error)
	CheckUserIngredientPreferenceExistence(ctx context.Context, db DBTX, arg *CheckUserIngredientPreferenceExistenceParams) (bool, error)
	CheckValidIngredientExistence(ctx context.Context, db DBTX, id string) (bool, error)
	CheckValidIngredientGroupExistence(ctx context.Context, db DBTX, id string) (bool, error)
	CheckValidIngredientMeasurementUnitExistence(ctx context.Context, db DBTX, id string) (bool, error)
	CheckValidIngredientPreparationExistence(ctx context.Context, db DBTX, id string) (bool, error)
	CheckValidIngredientStateExistence(ctx context.Context, db DBTX, id string) (bool, error)
	CheckValidIngredientStateIngredientExistence(ctx context.Context, db DBTX, id string) (bool, error)
	CheckValidInstrumentExistence(ctx context.Context, db DBTX, id string) (bool, error)
	CheckValidMeasurementUnitConversionExistence(ctx context.Context, db DBTX, id string) (bool, error)
	CheckValidMeasurementUnitExistence(ctx context.Context, db DBTX, id string) (bool, error)
	CheckValidPreparationExistence(ctx context.Context, db DBTX, id string) (bool, error)
	CheckValidPreparationInstrumentExistence(ctx context.Context, db DBTX, id string) (bool, error)
	CheckValidPreparationVesselExistence(ctx context.Context, db DBTX, id string) (bool, error)
	CheckValidVesselExistence(ctx context.Context, db DBTX, id string) (bool, error)
	CheckValidityOfValidIngredientStateIngredientPair(ctx context.Context, db DBTX, arg *CheckValidityOfValidIngredientStateIngredientPairParams) (bool, error)
	CheckWebhookExistence(ctx context.Context, db DBTX, arg *CheckWebhookExistenceParams) (bool, error)
	CreateHousehold(ctx context.Context, db DBTX, arg *CreateHouseholdParams) error
	CreateHouseholdInstrumentOwnership(ctx context.Context, db DBTX, arg *CreateHouseholdInstrumentOwnershipParams) error
	CreateHouseholdInvitation(ctx context.Context, db DBTX, arg *CreateHouseholdInvitationParams) error
	CreateHouseholdUserMembershipForNewUser(ctx context.Context, db DBTX, arg *CreateHouseholdUserMembershipForNewUserParams) error
	CreateMeal(ctx context.Context, db DBTX, arg *CreateMealParams) error
	CreateMealComponent(ctx context.Context, db DBTX, arg *CreateMealComponentParams) error
	CreateMealPlan(ctx context.Context, db DBTX, arg *CreateMealPlanParams) error
	CreateMealPlanEvent(ctx context.Context, db DBTX, arg *CreateMealPlanEventParams) error
	CreateMealPlanGroceryListItem(ctx context.Context, db DBTX, arg *CreateMealPlanGroceryListItemParams) error
	CreateMealPlanOption(ctx context.Context, db DBTX, arg *CreateMealPlanOptionParams) error
	CreateMealPlanOptionVote(ctx context.Context, db DBTX, arg *CreateMealPlanOptionVoteParams) error
	CreateMealPlanTask(ctx context.Context, db DBTX, arg *CreateMealPlanTaskParams) error
	CreateOAuth2Client(ctx context.Context, db DBTX, arg *CreateOAuth2ClientParams) error
	CreateOAuth2ClientToken(ctx context.Context, db DBTX, arg *CreateOAuth2ClientTokenParams) error
	CreatePasswordResetToken(ctx context.Context, db DBTX, arg *CreatePasswordResetTokenParams) error
	CreateRecipe(ctx context.Context, db DBTX, arg *CreateRecipeParams) error
	CreateRecipeMedia(ctx context.Context, db DBTX, arg *CreateRecipeMediaParams) error
	CreateRecipePrepTask(ctx context.Context, db DBTX, arg *CreateRecipePrepTaskParams) error
	CreateRecipePrepTaskStep(ctx context.Context, db DBTX, arg *CreateRecipePrepTaskStepParams) error
	CreateRecipeRating(ctx context.Context, db DBTX, arg *CreateRecipeRatingParams) error
	CreateRecipeStep(ctx context.Context, db DBTX, arg *CreateRecipeStepParams) error
	CreateRecipeStepCompletionCondition(ctx context.Context, db DBTX, arg *CreateRecipeStepCompletionConditionParams) error
	CreateRecipeStepCompletionConditionIngredient(ctx context.Context, db DBTX, arg *CreateRecipeStepCompletionConditionIngredientParams) error
	CreateRecipeStepIngredient(ctx context.Context, db DBTX, arg *CreateRecipeStepIngredientParams) error
	CreateRecipeStepInstrument(ctx context.Context, db DBTX, arg *CreateRecipeStepInstrumentParams) error
	CreateRecipeStepProduct(ctx context.Context, db DBTX, arg *CreateRecipeStepProductParams) error
	CreateRecipeStepVessel(ctx context.Context, db DBTX, arg *CreateRecipeStepVesselParams) error
	CreateServiceSetting(ctx context.Context, db DBTX, arg *CreateServiceSettingParams) error
	CreateServiceSettingConfiguration(ctx context.Context, db DBTX, arg *CreateServiceSettingConfigurationParams) error
	CreateUser(ctx context.Context, db DBTX, arg *CreateUserParams) error
	CreateUserIngredientPreference(ctx context.Context, db DBTX, arg *CreateUserIngredientPreferenceParams) error
	CreateValidIngredient(ctx context.Context, db DBTX, arg *CreateValidIngredientParams) error
	CreateValidIngredientGroup(ctx context.Context, db DBTX, arg *CreateValidIngredientGroupParams) error
	CreateValidIngredientGroupMember(ctx context.Context, db DBTX, arg *CreateValidIngredientGroupMemberParams) error
	CreateValidIngredientMeasurementUnit(ctx context.Context, db DBTX, arg *CreateValidIngredientMeasurementUnitParams) error
	CreateValidIngredientPreparation(ctx context.Context, db DBTX, arg *CreateValidIngredientPreparationParams) error
	CreateValidIngredientState(ctx context.Context, db DBTX, arg *CreateValidIngredientStateParams) error
	CreateValidIngredientStateIngredient(ctx context.Context, db DBTX, arg *CreateValidIngredientStateIngredientParams) error
	CreateValidInstrument(ctx context.Context, db DBTX, arg *CreateValidInstrumentParams) error
	CreateValidMeasurementUnit(ctx context.Context, db DBTX, arg *CreateValidMeasurementUnitParams) error
	CreateValidMeasurementUnitConversion(ctx context.Context, db DBTX, arg *CreateValidMeasurementUnitConversionParams) error
	CreateValidPreparation(ctx context.Context, db DBTX, arg *CreateValidPreparationParams) error
	CreateValidPreparationInstrument(ctx context.Context, db DBTX, arg *CreateValidPreparationInstrumentParams) error
	CreateValidPreparationVessel(ctx context.Context, db DBTX, arg *CreateValidPreparationVesselParams) error
	CreateValidVessel(ctx context.Context, db DBTX, arg *CreateValidVesselParams) error
	CreateWebhook(ctx context.Context, db DBTX, arg *CreateWebhookParams) error
	CreateWebhookTriggerEvent(ctx context.Context, db DBTX, arg *CreateWebhookTriggerEventParams) error
	FinalizeMealPlan(ctx context.Context, db DBTX, arg *FinalizeMealPlanParams) error
	FinalizeMealPlanOption(ctx context.Context, db DBTX, arg *FinalizeMealPlanOptionParams) error
	GetAdminUserByUsername(ctx context.Context, db DBTX, username string) (*GetAdminUserByUsernameRow, error)
	GetAllRecipeStepCompletionConditionsForRecipe(ctx context.Context, db DBTX, id string) ([]*GetAllRecipeStepCompletionConditionsForRecipeRow, error)
	GetAllValidMeasurementUnitConversionsFromMeasurementUnit(ctx context.Context, db DBTX, fromUnit string) ([]*GetAllValidMeasurementUnitConversionsFromMeasurementUnitRow, error)
	GetAllValidMeasurementUnitConversionsToMeasurementUnit(ctx context.Context, db DBTX, id string) ([]*GetAllValidMeasurementUnitConversionsToMeasurementUnitRow, error)
	GetDefaultHouseholdIDForUser(ctx context.Context, db DBTX, belongsToUser string) (string, error)
	GetEmailVerificationTokenByUserID(ctx context.Context, db DBTX, id string) (sql.NullString, error)
	GetExpiredAndUnresolvedMealPlans(ctx context.Context, db DBTX) ([]*GetExpiredAndUnresolvedMealPlansRow, error)
	GetFinalizedMealPlansForPlanning(ctx context.Context, db DBTX) ([]*GetFinalizedMealPlansForPlanningRow, error)
	GetFinalizedMealPlansWithoutGroceryListInit(ctx context.Context, db DBTX) ([]*GetFinalizedMealPlansWithoutGroceryListInitRow, error)
	GetHouseholdByIDWithMemberships(ctx context.Context, db DBTX, id string) ([]*GetHouseholdByIDWithMembershipsRow, error)
	GetHouseholdInstrumentOwnership(ctx context.Context, db DBTX, arg *GetHouseholdInstrumentOwnershipParams) (*GetHouseholdInstrumentOwnershipRow, error)
	GetHouseholdInstrumentOwnerships(ctx context.Context, db DBTX, arg *GetHouseholdInstrumentOwnershipsParams) ([]*GetHouseholdInstrumentOwnershipsRow, error)
	GetHouseholdInvitationByEmailAndToken(ctx context.Context, db DBTX, arg *GetHouseholdInvitationByEmailAndTokenParams) (*GetHouseholdInvitationByEmailAndTokenRow, error)
	GetHouseholdInvitationByHouseholdAndID(ctx context.Context, db DBTX, arg *GetHouseholdInvitationByHouseholdAndIDParams) (*GetHouseholdInvitationByHouseholdAndIDRow, error)
	GetHouseholdInvitationByTokenAndID(ctx context.Context, db DBTX, arg *GetHouseholdInvitationByTokenAndIDParams) (*GetHouseholdInvitationByTokenAndIDRow, error)
	GetHouseholdUserMembershipsForUser(ctx context.Context, db DBTX, belongsToUser string) ([]*GetHouseholdUserMembershipsForUserRow, error)
	GetMeal(ctx context.Context, db DBTX, id string) (*GetMealRow, error)
	GetMealPlan(ctx context.Context, db DBTX, arg *GetMealPlanParams) (*GetMealPlanRow, error)
	GetMealPlanEvent(ctx context.Context, db DBTX, arg *GetMealPlanEventParams) (*MealPlanEvents, error)
	GetMealPlanEventsForMealPlan(ctx context.Context, db DBTX, belongsToMealPlan string) ([]*MealPlanEvents, error)
	GetMealPlanGroceryListItem(ctx context.Context, db DBTX, arg *GetMealPlanGroceryListItemParams) (*GetMealPlanGroceryListItemRow, error)
	GetMealPlanGroceryListItemsForMealPlan(ctx context.Context, db DBTX, belongsToMealPlan string) ([]*GetMealPlanGroceryListItemsForMealPlanRow, error)
	GetMealPlanOption(ctx context.Context, db DBTX, arg *GetMealPlanOptionParams) (*GetMealPlanOptionRow, error)
	GetMealPlanOptionByID(ctx context.Context, db DBTX, id string) (*GetMealPlanOptionByIDRow, error)
	GetMealPlanOptionVote(ctx context.Context, db DBTX, arg *GetMealPlanOptionVoteParams) (*MealPlanOptionVotes, error)
	GetMealPlanOptionVotesForMealPlanOption(ctx context.Context, db DBTX, arg *GetMealPlanOptionVotesForMealPlanOptionParams) ([]*MealPlanOptionVotes, error)
	GetMealPlanOptionsForMealPlanEvent(ctx context.Context, db DBTX, arg *GetMealPlanOptionsForMealPlanEventParams) ([]*GetMealPlanOptionsForMealPlanEventRow, error)
	GetMealPlanTask(ctx context.Context, db DBTX, id string) (*GetMealPlanTaskRow, error)
	GetMealsNeedingIndexing(ctx context.Context, db DBTX) ([]string, error)
	GetOAuth2ClientByClientID(ctx context.Context, db DBTX, clientID string) (*GetOAuth2ClientByClientIDRow, error)
	GetOAuth2ClientByDatabaseID(ctx context.Context, db DBTX, id string) (*GetOAuth2ClientByDatabaseIDRow, error)
	GetOAuth2ClientTokenByAccess(ctx context.Context, db DBTX, access string) (*Oauth2ClientTokens, error)
	GetOAuth2ClientTokenByCode(ctx context.Context, db DBTX, code string) (*Oauth2ClientTokens, error)
	GetOAuth2ClientTokenByRefresh(ctx context.Context, db DBTX, refresh string) (*Oauth2ClientTokens, error)
	GetOnePastVotingDeadline(ctx context.Context, db DBTX, arg *GetOnePastVotingDeadlineParams) (*GetOnePastVotingDeadlineRow, error)
	GetPasswordResetToken(ctx context.Context, db DBTX, token string) (*PasswordResetTokens, error)
	GetRandomValidIngredient(ctx context.Context, db DBTX) (*GetRandomValidIngredientRow, error)
	GetRandomValidInstrument(ctx context.Context, db DBTX) (*GetRandomValidInstrumentRow, error)
	GetRandomValidMeasurementUnit(ctx context.Context, db DBTX) (*GetRandomValidMeasurementUnitRow, error)
	GetRandomValidPreparation(ctx context.Context, db DBTX) (*GetRandomValidPreparationRow, error)
	GetRandomValidVessel(ctx context.Context, db DBTX) (*GetRandomValidVesselRow, error)
	GetRecipeByID(ctx context.Context, db DBTX, id string) (*GetRecipeByIDRow, error)
	GetRecipeByIDAndAuthorID(ctx context.Context, db DBTX, arg *GetRecipeByIDAndAuthorIDParams) ([]*GetRecipeByIDAndAuthorIDRow, error)
	GetRecipeIDsForMeal(ctx context.Context, db DBTX, id string) ([]string, error)
	GetRecipeMedia(ctx context.Context, db DBTX, id string) (*GetRecipeMediaRow, error)
	GetRecipeMediaForRecipe(ctx context.Context, db DBTX, belongsToRecipe sql.NullString) ([]*GetRecipeMediaForRecipeRow, error)
	GetRecipeMediaForRecipeStep(ctx context.Context, db DBTX, arg *GetRecipeMediaForRecipeStepParams) ([]*GetRecipeMediaForRecipeStepRow, error)
	GetRecipePrepTask(ctx context.Context, db DBTX, id string) (*GetRecipePrepTaskRow, error)
	GetRecipeRating(ctx context.Context, db DBTX, id string) (*RecipeRatings, error)
	GetRecipeRatings(ctx context.Context, db DBTX, arg *GetRecipeRatingsParams) ([]*GetRecipeRatingsRow, error)
	GetRecipeStep(ctx context.Context, db DBTX, arg *GetRecipeStepParams) (*GetRecipeStepRow, error)
	GetRecipeStepByRecipeID(ctx context.Context, db DBTX, id string) (*GetRecipeStepByRecipeIDRow, error)
	GetRecipeStepCompletionCondition(ctx context.Context, db DBTX, arg *GetRecipeStepCompletionConditionParams) (*GetRecipeStepCompletionConditionRow, error)
	GetRecipeStepCompletionConditions(ctx context.Context, db DBTX, arg *GetRecipeStepCompletionConditionsParams) ([]*GetRecipeStepCompletionConditionsRow, error)
	GetRecipeStepIngredient(ctx context.Context, db DBTX, arg *GetRecipeStepIngredientParams) (*GetRecipeStepIngredientRow, error)
	GetRecipeStepIngredientsForRecipe(ctx context.Context, db DBTX, id string) ([]*GetRecipeStepIngredientsForRecipeRow, error)
	GetRecipeStepInstrument(ctx context.Context, db DBTX, arg *GetRecipeStepInstrumentParams) (*GetRecipeStepInstrumentRow, error)
	GetRecipeStepInstrumentsForRecipe(ctx context.Context, db DBTX, belongsToRecipe string) ([]*GetRecipeStepInstrumentsForRecipeRow, error)
	GetRecipeStepProduct(ctx context.Context, db DBTX, arg *GetRecipeStepProductParams) (*GetRecipeStepProductRow, error)
	GetRecipeStepProductsForRecipe(ctx context.Context, db DBTX, belongsToRecipe string) ([]*GetRecipeStepProductsForRecipeRow, error)
	GetRecipeStepVessel(ctx context.Context, db DBTX, arg *GetRecipeStepVesselParams) (*GetRecipeStepVesselRow, error)
	GetRecipeStepVesselsForRecipe(ctx context.Context, db DBTX, belongsToRecipe string) ([]*GetRecipeStepVesselsForRecipeRow, error)
	GetRecipesNeedingIndexing(ctx context.Context, db DBTX) ([]string, error)
	GetServiceSetting(ctx context.Context, db DBTX, id string) (*GetServiceSettingRow, error)
	GetServiceSettingConfigurationByID(ctx context.Context, db DBTX, id string) (*GetServiceSettingConfigurationByIDRow, error)
	GetServiceSettingConfigurationsForHousehold(ctx context.Context, db DBTX, belongsToHousehold string) ([]*GetServiceSettingConfigurationsForHouseholdRow, error)
	GetServiceSettingConfigurationsForHouseholdBySettingName(ctx context.Context, db DBTX, arg *GetServiceSettingConfigurationsForHouseholdBySettingNameParams) ([]*GetServiceSettingConfigurationsForHouseholdBySettingNameRow, error)
	GetServiceSettingConfigurationsForUser(ctx context.Context, db DBTX, belongsToUser string) ([]*GetServiceSettingConfigurationsForUserRow, error)
	GetServiceSettingConfigurationsForUserBySettingName(ctx context.Context, db DBTX, arg *GetServiceSettingConfigurationsForUserBySettingNameParams) ([]*GetServiceSettingConfigurationsForUserBySettingNameRow, error)
	GetServiceSettings(ctx context.Context, db DBTX, arg *GetServiceSettingsParams) ([]*GetServiceSettingsRow, error)
	GetUserByEmail(ctx context.Context, db DBTX, emailAddress string) (*GetUserByEmailRow, error)
	GetUserByEmailAddressVerificationToken(ctx context.Context, db DBTX, emailAddressVerificationToken sql.NullString) (*GetUserByEmailAddressVerificationTokenRow, error)
	GetUserByID(ctx context.Context, db DBTX, id string) (*GetUserByIDRow, error)
	GetUserByUsername(ctx context.Context, db DBTX, username string) (*GetUserByUsernameRow, error)
	GetUserIDsNeedingIndexing(ctx context.Context, db DBTX) ([]string, error)
	GetUserIngredientPreference(ctx context.Context, db DBTX, arg *GetUserIngredientPreferenceParams) (*GetUserIngredientPreferenceRow, error)
	GetUserIngredientPreferencesForUser(ctx context.Context, db DBTX, arg *GetUserIngredientPreferencesForUserParams) ([]*GetUserIngredientPreferencesForUserRow, error)
	GetUserWithUnverifiedTwoFactor(ctx context.Context, db DBTX, id string) (*GetUserWithUnverifiedTwoFactorRow, error)
	GetUserWithVerifiedTwoFactor(ctx context.Context, db DBTX, id string) (*GetUserWithVerifiedTwoFactorRow, error)
	GetUsers(ctx context.Context, db DBTX, arg *GetUsersParams) ([]*GetUsersRow, error)
	GetValidIngredient(ctx context.Context, db DBTX, id string) (*GetValidIngredientRow, error)
	GetValidIngredientByID(ctx context.Context, db DBTX) (*GetValidIngredientByIDRow, error)
	GetValidIngredientGroup(ctx context.Context, db DBTX, id string) (*GetValidIngredientGroupRow, error)
	GetValidIngredientGroupMembers(ctx context.Context, db DBTX, belongsToGroup string) ([]*GetValidIngredientGroupMembersRow, error)
	GetValidIngredientGroups(ctx context.Context, db DBTX, arg *GetValidIngredientGroupsParams) ([]*GetValidIngredientGroupsRow, error)
	GetValidIngredientMeasurementUnit(ctx context.Context, db DBTX, id string) (*GetValidIngredientMeasurementUnitRow, error)
	GetValidIngredientMeasurementUnits(ctx context.Context, db DBTX, arg *GetValidIngredientMeasurementUnitsParams) ([]*GetValidIngredientMeasurementUnitsRow, error)
	GetValidIngredientMeasurementUnitsForIngredient(ctx context.Context, db DBTX, arg *GetValidIngredientMeasurementUnitsForIngredientParams) ([]*GetValidIngredientMeasurementUnitsForIngredientRow, error)
	GetValidIngredientMeasurementUnitsForMeasurementUnit(ctx context.Context, db DBTX, arg *GetValidIngredientMeasurementUnitsForMeasurementUnitParams) ([]*GetValidIngredientMeasurementUnitsForMeasurementUnitRow, error)
	GetValidIngredientPreparation(ctx context.Context, db DBTX, id string) (*GetValidIngredientPreparationRow, error)
	GetValidIngredientPreparations(ctx context.Context, db DBTX, arg *GetValidIngredientPreparationsParams) ([]*GetValidIngredientPreparationsRow, error)
	GetValidIngredientPreparationsForIngredient(ctx context.Context, db DBTX, arg *GetValidIngredientPreparationsForIngredientParams) ([]*GetValidIngredientPreparationsForIngredientRow, error)
	GetValidIngredientPreparationsForPreparation(ctx context.Context, db DBTX, arg *GetValidIngredientPreparationsForPreparationParams) ([]*GetValidIngredientPreparationsForPreparationRow, error)
	GetValidIngredientState(ctx context.Context, db DBTX, id string) (*GetValidIngredientStateRow, error)
	GetValidIngredientStateIngredient(ctx context.Context, db DBTX, id string) (*GetValidIngredientStateIngredientRow, error)
	GetValidIngredientStateIngredients(ctx context.Context, db DBTX, arg *GetValidIngredientStateIngredientsParams) ([]*GetValidIngredientStateIngredientsRow, error)
	GetValidIngredientStateIngredientsForIngredient(ctx context.Context, db DBTX, arg *GetValidIngredientStateIngredientsForIngredientParams) ([]*GetValidIngredientStateIngredientsForIngredientRow, error)
	GetValidIngredientStateIngredientsForIngredientState(ctx context.Context, db DBTX, arg *GetValidIngredientStateIngredientsForIngredientStateParams) ([]*GetValidIngredientStateIngredientsForIngredientStateRow, error)
	GetValidIngredientStateIngredientsWithIDs(ctx context.Context, db DBTX, ids []string) ([]*GetValidIngredientStateIngredientsWithIDsRow, error)
	GetValidIngredientStates(ctx context.Context, db DBTX, arg *GetValidIngredientStatesParams) ([]*GetValidIngredientStatesRow, error)
	GetValidIngredientStatesNeedingIndexing(ctx context.Context, db DBTX) ([]string, error)
	GetValidIngredientStatesWithIDs(ctx context.Context, db DBTX, dollar_1 []string) ([]*GetValidIngredientStatesWithIDsRow, error)
	GetValidIngredients(ctx context.Context, db DBTX, arg *GetValidIngredientsParams) ([]*GetValidIngredientsRow, error)
	GetValidIngredientsNeedingIndexing(ctx context.Context, db DBTX) ([]string, error)
	GetValidIngredientsWithIDs(ctx context.Context, db DBTX, ids []string) ([]*GetValidIngredientsWithIDsRow, error)
	GetValidInstrument(ctx context.Context, db DBTX, id string) (*GetValidInstrumentRow, error)
	GetValidInstrumentWithIDs(ctx context.Context, db DBTX, ids []string) ([]*GetValidInstrumentWithIDsRow, error)
	GetValidInstruments(ctx context.Context, db DBTX, arg *GetValidInstrumentsParams) ([]*GetValidInstrumentsRow, error)
	GetValidInstrumentsNeedingIndexing(ctx context.Context, db DBTX) ([]string, error)
	GetValidMeasurementUnit(ctx context.Context, db DBTX, id string) (*GetValidMeasurementUnitRow, error)
	GetValidMeasurementUnitByID(ctx context.Context, db DBTX) (*GetValidMeasurementUnitByIDRow, error)
	GetValidMeasurementUnitConversion(ctx context.Context, db DBTX, id string) (*GetValidMeasurementUnitConversionRow, error)
	GetValidMeasurementUnits(ctx context.Context, db DBTX, arg *GetValidMeasurementUnitsParams) ([]*GetValidMeasurementUnitsRow, error)
	GetValidMeasurementUnitsNeedingIndexing(ctx context.Context, db DBTX) ([]string, error)
	GetValidMeasurementUnitsWithIDs(ctx context.Context, db DBTX, ids []string) ([]*GetValidMeasurementUnitsWithIDsRow, error)
	GetValidPreparation(ctx context.Context, db DBTX, id string) (*GetValidPreparationRow, error)
	GetValidPreparationByID(ctx context.Context, db DBTX) (*GetValidPreparationByIDRow, error)
	GetValidPreparationInstrument(ctx context.Context, db DBTX, id string) (*GetValidPreparationInstrumentRow, error)
	GetValidPreparationInstruments(ctx context.Context, db DBTX, arg *GetValidPreparationInstrumentsParams) ([]*GetValidPreparationInstrumentsRow, error)
	GetValidPreparationInstrumentsForInstrument(ctx context.Context, db DBTX, arg *GetValidPreparationInstrumentsForInstrumentParams) ([]*GetValidPreparationInstrumentsForInstrumentRow, error)
	GetValidPreparationInstrumentsForPreparation(ctx context.Context, db DBTX, arg *GetValidPreparationInstrumentsForPreparationParams) ([]*GetValidPreparationInstrumentsForPreparationRow, error)
	GetValidPreparationVessel(ctx context.Context, db DBTX, id string) (*GetValidPreparationVesselRow, error)
	GetValidPreparationVessels(ctx context.Context, db DBTX, arg *GetValidPreparationVesselsParams) ([]*GetValidPreparationVesselsRow, error)
	GetValidPreparationVesselsForPreparation(ctx context.Context, db DBTX, arg *GetValidPreparationVesselsForPreparationParams) ([]*GetValidPreparationVesselsForPreparationRow, error)
	GetValidPreparationVesselsForVessel(ctx context.Context, db DBTX, arg *GetValidPreparationVesselsForVesselParams) ([]*GetValidPreparationVesselsForVesselRow, error)
	GetValidPreparations(ctx context.Context, db DBTX, arg *GetValidPreparationsParams) ([]*GetValidPreparationsRow, error)
	GetValidPreparationsNeedingIndexing(ctx context.Context, db DBTX) ([]string, error)
	GetValidPreparationsWithIDs(ctx context.Context, db DBTX, dollar_1 []string) ([]*GetValidPreparationsWithIDsRow, error)
	GetValidVessel(ctx context.Context, db DBTX, id string) (*GetValidVesselRow, error)
	GetValidVesselIDsNeedingIndexing(ctx context.Context, db DBTX) ([]string, error)
	GetValidVessels(ctx context.Context, db DBTX, arg *GetValidVesselsParams) ([]*GetValidVesselsRow, error)
	GetValidVesselsWithIDs(ctx context.Context, db DBTX, ids []string) ([]*GetValidVesselsWithIDsRow, error)
	GetWebhook(ctx context.Context, db DBTX, arg *GetWebhookParams) ([]*GetWebhookRow, error)
	GetWebhooks(ctx context.Context, db DBTX, arg *GetWebhooksParams) ([]*GetWebhooksRow, error)
	GetWebhooksForUser(ctx context.Context, db DBTX, arg *GetWebhooksForUserParams) ([]*GetWebhooksForUserRow, error)
	ListAllMealPlanTasksByMealPlan(ctx context.Context, db DBTX, id string) ([]*ListAllMealPlanTasksByMealPlanRow, error)
	ListAllRecipePrepTasksByRecipe(ctx context.Context, db DBTX, id string) ([]*ListAllRecipePrepTasksByRecipeRow, error)
	ListIncompleteMealPlanTasksByMealPlanOption(ctx context.Context, db DBTX, belongsToMealPlanOption string) ([]*ListIncompleteMealPlanTasksByMealPlanOptionRow, error)
	MarkEmailAddressAsVerified(ctx context.Context, db DBTX, arg *MarkEmailAddressAsVerifiedParams) error
	MarkHouseholdUserMembershipAsUserDefault(ctx context.Context, db DBTX, arg *MarkHouseholdUserMembershipAsUserDefaultParams) error
	MarkMealPlanAsGroceryListInitialized(ctx context.Context, db DBTX, id string) error
	MarkMealPlanAsPrepTasksCreated(ctx context.Context, db DBTX, id string) error
	MarkTwoFactorSecretAsUnverified(ctx context.Context, db DBTX, arg *MarkTwoFactorSecretAsUnverifiedParams) error
	MarkTwoFactorSecretAsVerified(ctx context.Context, db DBTX, id string) error
	MealPlanEventIsEligibleForVoting(ctx context.Context, db DBTX, arg *MealPlanEventIsEligibleForVotingParams) (bool, error)
	ModifyHouseholdUserPermissions(ctx context.Context, db DBTX, arg *ModifyHouseholdUserPermissionsParams) error
	RedeemPasswordResetToken(ctx context.Context, db DBTX, id string) error
	RemoveUserFromHousehold(ctx context.Context, db DBTX, arg *RemoveUserFromHouseholdParams) error
	SearchForServiceSettings(ctx context.Context, db DBTX, name string) ([]*SearchForServiceSettingsRow, error)
	SearchForValidIngredientGroups(ctx context.Context, db DBTX, arg *SearchForValidIngredientGroupsParams) ([]*SearchForValidIngredientGroupsRow, error)
	SearchForValidIngredientStates(ctx context.Context, db DBTX, query string) ([]*SearchForValidIngredientStatesRow, error)
	SearchForValidIngredients(ctx context.Context, db DBTX, dollar_1 string) ([]*SearchForValidIngredientsRow, error)
	SearchForValidInstruments(ctx context.Context, db DBTX, query string) ([]*SearchForValidInstrumentsRow, error)
	SearchForValidMeasurementUnits(ctx context.Context, db DBTX, query string) ([]*SearchForValidMeasurementUnitsRow, error)
	SearchForValidPreparations(ctx context.Context, db DBTX, query string) ([]*SearchForValidPreparationsRow, error)
	SearchForValidVessels(ctx context.Context, db DBTX, query string) ([]*SearchForValidVesselsRow, error)
	SearchUsersByUsername(ctx context.Context, db DBTX, username string) ([]*SearchUsersByUsernameRow, error)
	SearchValidIngredientPreparationsByPreparationAndIngredientName(ctx context.Context, db DBTX, arg *SearchValidIngredientPreparationsByPreparationAndIngredientNameParams) ([]*SearchValidIngredientPreparationsByPreparationAndIngredientNameRow, error)
	SearchValidIngredientsByPreparationAndIngredientName(ctx context.Context, db DBTX, arg *SearchValidIngredientsByPreparationAndIngredientNameParams) ([]*SearchValidIngredientsByPreparationAndIngredientNameRow, error)
	SearchValidMeasurementUnitsByIngredientID(ctx context.Context, db DBTX, arg *SearchValidMeasurementUnitsByIngredientIDParams) ([]*SearchValidMeasurementUnitsByIngredientIDRow, error)
	SetHouseholdInvitationStatus(ctx context.Context, db DBTX, arg *SetHouseholdInvitationStatusParams) error
	SetUserAccountStatus(ctx context.Context, db DBTX, arg *SetUserAccountStatusParams) (sql.Result, error)
	TransferHouseholdMembership(ctx context.Context, db DBTX, arg *TransferHouseholdMembershipParams) error
	TransferHouseholdOwnership(ctx context.Context, db DBTX, arg *TransferHouseholdOwnershipParams) error
	UpdateHousehold(ctx context.Context, db DBTX, arg *UpdateHouseholdParams) error
	UpdateHouseholdInstrumentOwnership(ctx context.Context, db DBTX, arg *UpdateHouseholdInstrumentOwnershipParams) error
	UpdateMealLastIndexedAt(ctx context.Context, db DBTX, id string) error
	UpdateMealPlan(ctx context.Context, db DBTX, arg *UpdateMealPlanParams) error
	UpdateMealPlanEvent(ctx context.Context, db DBTX, arg *UpdateMealPlanEventParams) error
	UpdateMealPlanGroceryListItem(ctx context.Context, db DBTX, arg *UpdateMealPlanGroceryListItemParams) error
	UpdateMealPlanOption(ctx context.Context, db DBTX, arg *UpdateMealPlanOptionParams) error
	UpdateMealPlanOptionVote(ctx context.Context, db DBTX, arg *UpdateMealPlanOptionVoteParams) error
	UpdateRecipe(ctx context.Context, db DBTX, arg *UpdateRecipeParams) error
	UpdateRecipeLastIndexedAt(ctx context.Context, db DBTX, id string) error
	UpdateRecipeMedia(ctx context.Context, db DBTX, arg *UpdateRecipeMediaParams) error
	UpdateRecipePrepTask(ctx context.Context, db DBTX, arg *UpdateRecipePrepTaskParams) error
	UpdateRecipeRating(ctx context.Context, db DBTX, arg *UpdateRecipeRatingParams) error
	UpdateRecipeStep(ctx context.Context, db DBTX, arg *UpdateRecipeStepParams) error
	UpdateRecipeStepCompletionCondition(ctx context.Context, db DBTX, arg *UpdateRecipeStepCompletionConditionParams) error
	UpdateRecipeStepIngredient(ctx context.Context, db DBTX, arg *UpdateRecipeStepIngredientParams) error
	UpdateRecipeStepInstrument(ctx context.Context, db DBTX, arg *UpdateRecipeStepInstrumentParams) error
	UpdateRecipeStepProduct(ctx context.Context, db DBTX, arg *UpdateRecipeStepProductParams) error
	UpdateRecipeStepVessel(ctx context.Context, db DBTX, arg *UpdateRecipeStepVesselParams) error
	UpdateServiceSettingConfiguration(ctx context.Context, db DBTX, arg *UpdateServiceSettingConfigurationParams) error
	UpdateUser(ctx context.Context, db DBTX, arg *UpdateUserParams) error
	UpdateUserAvatarSrc(ctx context.Context, db DBTX, arg *UpdateUserAvatarSrcParams) error
	UpdateUserDetails(ctx context.Context, db DBTX, arg *UpdateUserDetailsParams) error
	UpdateUserEmailAddress(ctx context.Context, db DBTX, arg *UpdateUserEmailAddressParams) error
	UpdateUserIngredientPreference(ctx context.Context, db DBTX, arg *UpdateUserIngredientPreferenceParams) error
	UpdateUserLastIndexedAt(ctx context.Context, db DBTX, id string) error
	UpdateUserPassword(ctx context.Context, db DBTX, arg *UpdateUserPasswordParams) error
	UpdateUserTwoFactorSecret(ctx context.Context, db DBTX, arg *UpdateUserTwoFactorSecretParams) error
	UpdateUserUsername(ctx context.Context, db DBTX, arg *UpdateUserUsernameParams) error
	UpdateValidIngredient(ctx context.Context, db DBTX, arg *UpdateValidIngredientParams) error
	UpdateValidIngredientGroup(ctx context.Context, db DBTX, arg *UpdateValidIngredientGroupParams) error
	UpdateValidIngredientLastIndexedAt(ctx context.Context, db DBTX, id string) error
	UpdateValidIngredientMeasurementUnit(ctx context.Context, db DBTX, arg *UpdateValidIngredientMeasurementUnitParams) error
	UpdateValidIngredientPreparation(ctx context.Context, db DBTX, arg *UpdateValidIngredientPreparationParams) error
	UpdateValidIngredientState(ctx context.Context, db DBTX, arg *UpdateValidIngredientStateParams) error
	UpdateValidIngredientStateIngredient(ctx context.Context, db DBTX, arg *UpdateValidIngredientStateIngredientParams) error
	UpdateValidIngredientStateLastIndexedAt(ctx context.Context, db DBTX, id string) error
	UpdateValidInstrument(ctx context.Context, db DBTX, arg *UpdateValidInstrumentParams) error
	UpdateValidInstrumentLastIndexedAt(ctx context.Context, db DBTX, id string) error
	UpdateValidMeasurementUnit(ctx context.Context, db DBTX, arg *UpdateValidMeasurementUnitParams) error
	UpdateValidMeasurementUnitConversion(ctx context.Context, db DBTX, arg *UpdateValidMeasurementUnitConversionParams) error
	UpdateValidMeasurementUnitLastIndexedAt(ctx context.Context, db DBTX, id string) error
	UpdateValidPreparation(ctx context.Context, db DBTX, arg *UpdateValidPreparationParams) error
	UpdateValidPreparationInstrument(ctx context.Context, db DBTX, arg *UpdateValidPreparationInstrumentParams) error
	UpdateValidPreparationLastIndexedAt(ctx context.Context, db DBTX, id string) error
	UpdateValidPreparationVessel(ctx context.Context, db DBTX, arg *UpdateValidPreparationVesselParams) error
	UpdateValidVessel(ctx context.Context, db DBTX, arg *UpdateValidVesselParams) error
	UpdateValidVesselLastIndexedAt(ctx context.Context, db DBTX, id string) error
	UserIsHouseholdMember(ctx context.Context, db DBTX, arg *UserIsHouseholdMemberParams) (bool, error)
	ValidIngredientMeasurementUnitPairIsValid(ctx context.Context, db DBTX, arg *ValidIngredientMeasurementUnitPairIsValidParams) (bool, error)
	ValidIngredientPreparationPairIsValid(ctx context.Context, db DBTX, arg *ValidIngredientPreparationPairIsValidParams) (bool, error)
	ValidPreparationInstrumentPairIsValid(ctx context.Context, db DBTX, arg *ValidPreparationInstrumentPairIsValidParams) (bool, error)
	ValidPreparationVesselPairIsValid(ctx context.Context, db DBTX, arg *ValidPreparationVesselPairIsValidParams) (bool, error)
}

var _ Querier = (*Queries)(nil)
