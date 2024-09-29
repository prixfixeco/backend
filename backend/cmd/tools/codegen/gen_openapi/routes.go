package main

import (
	"github.com/dinnerdonebetter/backend/pkg/types"
)

type routeDetails struct {
	ResponseType  any
	InputType     any
	ID            string
	Description   string
	InputTypeName string
	OAuth2Scopes  []string
	Authless      bool
	ListRoute     bool
	SearchRoute   bool
}

const (
	serviceAdmin    = "service_admin"
	householdAdmin  = "household_admin"
	householdMember = "household_member"
)

var routeInfoMap = map[string]routeDetails{
	"GET /_meta_/live": {
		ID:          "CheckForLiveness",
		Description: "checks for service liveness",
		Authless:    true,
	},
	"GET /_meta_/ready": {
		ID:          "CheckForReadiness",
		Description: "checks for service readiness",
		Authless:    true,
	},
	"POST /api/v1/admin/cycle_cookie_secret": {
		ID:           "AdminCycleCookieSecret",
		OAuth2Scopes: []string{serviceAdmin},
	},
	"POST /api/v1/admin/users/status": {
		ID:           "AdminUpdateUserStatus",
		InputType:    &types.UserAccountStatusUpdateInput{},
		ResponseType: &types.UserStatusResponse{},
		OAuth2Scopes: []string{serviceAdmin},
	},
	"GET /api/v1/audit_log_entries/for_household": {
		ID:           "GetAuditLogEntriesForHousehold",
		ResponseType: &types.AuditLogEntry{},
		OAuth2Scopes: []string{householdAdmin},
	},
	"GET /api/v1/audit_log_entries/for_user": {
		ID:           "GetAuditLogEntriesForUser",
		ResponseType: &types.AuditLogEntry{},
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/audit_log_entries/{auditLogEntryID}": {
		ID:           "GetAuditLogEntryByID",
		ResponseType: &types.AuditLogEntry{},
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/household_invitations/received": {
		ID:           "GetReceivedHouseholdInvitations",
		ResponseType: &types.HouseholdInvitation{},
		ListRoute:    true,
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/household_invitations/sent": {
		ID:           "GetSentHouseholdInvitations",
		ResponseType: &types.HouseholdInvitation{},
		ListRoute:    true,
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/household_invitations/{householdInvitationID}/": {
		ID:           "GetHouseholdInvitation",
		ResponseType: &types.HouseholdInvitation{},
		OAuth2Scopes: []string{householdMember},
	},
	"PUT /api/v1/household_invitations/{householdInvitationID}/accept": {
		ID:           "AcceptHouseholdInvitation",
		ResponseType: &types.HouseholdInvitation{},
		InputType:    &types.HouseholdInvitationUpdateRequestInput{},
		OAuth2Scopes: []string{householdMember},
	},
	"PUT /api/v1/household_invitations/{householdInvitationID}/cancel": {
		ID:           "CancelHouseholdInvitation",
		ResponseType: &types.HouseholdInvitation{},
		InputType:    &types.HouseholdInvitationUpdateRequestInput{},
		OAuth2Scopes: []string{householdMember},
	},
	"PUT /api/v1/household_invitations/{householdInvitationID}/reject": {
		ID:           "RejectHouseholdInvitation",
		ResponseType: &types.HouseholdInvitation{},
		InputType:    &types.HouseholdInvitationUpdateRequestInput{},
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/households/": {
		ID:           "GetHouseholds",
		ResponseType: &types.Household{},
		ListRoute:    true,
		OAuth2Scopes: []string{householdMember},
	},
	"POST /api/v1/households/": {
		ID:           "CreateHousehold",
		ResponseType: &types.Household{},
		InputType:    &types.HouseholdCreationRequestInput{},
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/households/current": {
		ID:           "GetActiveHousehold",
		ResponseType: &types.Household{},
		OAuth2Scopes: []string{householdMember},
	},
	"POST /api/v1/households/instruments/": {
		ID:           "CreateHouseholdInstrumentOwnership",
		ResponseType: &types.HouseholdInstrumentOwnership{},
		InputType:    &types.HouseholdInstrumentOwnershipCreationRequestInput{},
		OAuth2Scopes: []string{householdAdmin},
	},
	"GET /api/v1/households/instruments/": {
		ID:           "GetHouseholdInstrumentOwnerships",
		ResponseType: &types.HouseholdInstrumentOwnership{},
		OAuth2Scopes: []string{householdMember},
		ListRoute:    true,
	},
	"GET /api/v1/households/instruments/{householdInstrumentOwnershipID}/": {
		ID:           "GetHouseholdInstrumentOwnership",
		ResponseType: &types.HouseholdInstrumentOwnership{},
		OAuth2Scopes: []string{householdMember},
	},
	"PUT /api/v1/households/instruments/{householdInstrumentOwnershipID}/": {
		ID:           "UpdateHouseholdInstrumentOwnership",
		ResponseType: &types.HouseholdInstrumentOwnership{},
		InputType:    &types.HouseholdInstrumentOwnershipUpdateRequestInput{},
		OAuth2Scopes: []string{householdAdmin},
	},
	"DELETE /api/v1/households/instruments/{householdInstrumentOwnershipID}/": {
		ID:           "ArchiveHouseholdInstrumentOwnership",
		ResponseType: &types.HouseholdInstrumentOwnership{},
		OAuth2Scopes: []string{householdAdmin},
	},
	"PUT /api/v1/households/{householdID}/": {
		ID:           "UpdateHousehold",
		ResponseType: &types.Household{},
		InputType:    &types.HouseholdUpdateRequestInput{},
		OAuth2Scopes: []string{householdAdmin},
	},
	"DELETE /api/v1/households/{householdID}/": {
		ID:           "ArchiveHousehold",
		ResponseType: &types.Household{},
		OAuth2Scopes: []string{householdAdmin},
	},
	"GET /api/v1/households/{householdID}/": {
		ID:           "GetHousehold",
		ResponseType: &types.Household{},
		OAuth2Scopes: []string{householdMember},
	},
	"POST /api/v1/households/{householdID}/default": {
		ID:           "SetDefaultHousehold",
		ResponseType: &types.Household{},
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/households/{householdID}/invitations/{householdInvitationID}/": {
		ID:           "GetHouseholdInvitationByID",
		ResponseType: &types.HouseholdInvitation{},
		OAuth2Scopes: []string{householdMember},
	},
	// these are duplicate routes lol
	"POST /api/v1/households/{householdID}/invitations/": {
		ID:           "CreateHouseholdInvitation",
		ResponseType: &types.HouseholdInvitation{},
		InputType:    &types.HouseholdInvitationCreationRequestInput{},
		OAuth2Scopes: []string{householdAdmin},
	},
	"POST /api/v1/households/{householdID}/invite": {
		ResponseType: &types.HouseholdInvitation{},
		InputType:    &types.HouseholdInvitationCreationRequestInput{},
		OAuth2Scopes: []string{householdAdmin},
	},
	"DELETE /api/v1/households/{householdID}/members/{userID}": {
		ID:           "ArchiveUserMembership",
		ResponseType: &types.HouseholdUserMembership{},
		OAuth2Scopes: []string{householdAdmin},
	},
	"PATCH /api/v1/households/{householdID}/members/{userID}/permissions": {
		ID:           "UpdateHouseholdMemberPermissions",
		ResponseType: &types.UserPermissionsResponse{},
		InputType:    &types.ModifyUserPermissionsInput{},
		OAuth2Scopes: []string{householdAdmin},
	},
	"POST /api/v1/households/{householdID}/transfer": {
		ID:           "TransferHouseholdOwnership",
		ResponseType: &types.Household{},
		InputType:    &types.HouseholdOwnershipTransferInput{},
		OAuth2Scopes: []string{householdAdmin},
	},
	"GET /api/v1/meal_plans/{mealPlanID}/events/{mealPlanEventID}/options/{mealPlanOptionID}/votes/": {
		ID:           "GetMealPlanOptionVotes",
		ResponseType: &types.MealPlanOptionVote{},
		ListRoute:    true,
		OAuth2Scopes: []string{householdMember},
	},
	"PUT /api/v1/meal_plans/{mealPlanID}/events/{mealPlanEventID}/options/{mealPlanOptionID}/votes/{mealPlanOptionVoteID}/": {
		ID:           "UpdateMealPlanOptionVote",
		ResponseType: &types.MealPlanOptionVote{},
		InputType:    &types.MealPlanOptionVoteUpdateRequestInput{},
		OAuth2Scopes: []string{householdMember},
	},
	"DELETE /api/v1/meal_plans/{mealPlanID}/events/{mealPlanEventID}/options/{mealPlanOptionID}/votes/{mealPlanOptionVoteID}/": {
		ID:           "ArchiveMealPlanOptionVote",
		ResponseType: &types.MealPlanOptionVote{},
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/meal_plans/{mealPlanID}/events/{mealPlanEventID}/options/{mealPlanOptionID}/votes/{mealPlanOptionVoteID}/": {
		ID:           "GetMealPlanOptionVote",
		ResponseType: &types.MealPlanOptionVote{},
		OAuth2Scopes: []string{householdMember},
	},
	"POST /api/v1/meal_plans/{mealPlanID}/events/{mealPlanEventID}/options/": {
		ID:           "CreateMealPlanOption",
		ResponseType: &types.MealPlanOption{},
		InputType:    &types.MealPlanOptionCreationRequestInput{},
		OAuth2Scopes: []string{householdAdmin},
	},
	"GET /api/v1/meal_plans/{mealPlanID}/events/{mealPlanEventID}/options/": {
		ID:           "GetMealPlanOptions",
		ResponseType: &types.MealPlanOption{},
		ListRoute:    true,
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/meal_plans/{mealPlanID}/events/{mealPlanEventID}/options/{mealPlanOptionID}/": {
		ID:           "GetMealPlanOption",
		ResponseType: &types.MealPlanOption{},
		OAuth2Scopes: []string{householdMember},
	},
	"PUT /api/v1/meal_plans/{mealPlanID}/events/{mealPlanEventID}/options/{mealPlanOptionID}/": {
		ID:           "UpdateMealPlanOption",
		ResponseType: &types.MealPlanOption{},
		InputType:    &types.MealPlanOptionUpdateRequestInput{},
		OAuth2Scopes: []string{householdAdmin},
	},
	"DELETE /api/v1/meal_plans/{mealPlanID}/events/{mealPlanEventID}/options/{mealPlanOptionID}/": {
		ID:           "ArchiveMealPlanOption",
		ResponseType: &types.MealPlanOption{},
		OAuth2Scopes: []string{householdAdmin},
	},
	"POST /api/v1/meal_plans/{mealPlanID}/events/": {
		ID:           "CreateMealPlanEvent",
		ResponseType: &types.MealPlanEvent{},
		InputType:    &types.MealPlanEventCreationRequestInput{},
		OAuth2Scopes: []string{householdAdmin},
	},
	"GET /api/v1/meal_plans/{mealPlanID}/events/": {
		ID:           "GetMealPlanEvents",
		ResponseType: &types.MealPlanEvent{},
		ListRoute:    true,
		OAuth2Scopes: []string{householdMember},
	},
	"PUT /api/v1/meal_plans/{mealPlanID}/events/{mealPlanEventID}/": {
		ID:           "UpdateMealPlanEvent",
		ResponseType: &types.MealPlanEvent{},
		InputType:    &types.MealPlanEventUpdateRequestInput{},
		OAuth2Scopes: []string{householdAdmin},
	},
	"DELETE /api/v1/meal_plans/{mealPlanID}/events/{mealPlanEventID}/": {
		ID:           "ArchiveMealPlanEvent",
		ResponseType: &types.MealPlanEvent{},
		OAuth2Scopes: []string{householdAdmin},
	},
	"GET /api/v1/meal_plans/{mealPlanID}/events/{mealPlanEventID}/": {
		ID:           "GetMealPlanEvent",
		ResponseType: &types.MealPlanEvent{},
		OAuth2Scopes: []string{householdMember},
	},
	"POST /api/v1/meal_plans/{mealPlanID}/events/{mealPlanEventID}/vote": {
		ID:           "CreateMealPlanOptionVote",
		ResponseType: &types.MealPlanOptionVote{},
		InputType:    &types.MealPlanOptionVoteCreationRequestInput{},
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/meal_plans/{mealPlanID}/grocery_list_items/": {
		ID:           "GetMealPlanGroceryListItemsForMealPlan",
		ResponseType: &types.MealPlanGroceryListItem{},
		ListRoute:    true,
		OAuth2Scopes: []string{householdMember},
	},
	"POST /api/v1/meal_plans/{mealPlanID}/grocery_list_items/": {
		ID:           "CreateMealPlanGroceryListItem",
		ResponseType: &types.MealPlanGroceryListItem{},
		InputType:    &types.MealPlanGroceryListItemCreationRequestInput{},
		OAuth2Scopes: []string{serviceAdmin},
	},
	"GET /api/v1/meal_plans/{mealPlanID}/grocery_list_items/{mealPlanGroceryListItemID}/": {
		ID:           "GetMealPlanGroceryListItem",
		ResponseType: &types.MealPlanGroceryListItem{},
		OAuth2Scopes: []string{householdMember},
	},
	"PUT /api/v1/meal_plans/{mealPlanID}/grocery_list_items/{mealPlanGroceryListItemID}/": {
		ID:           "UpdateMealPlanGroceryListItem",
		ResponseType: &types.MealPlanGroceryListItem{},
		InputType:    &types.MealPlanGroceryListItemUpdateRequestInput{},
		OAuth2Scopes: []string{householdMember},
	},
	"DELETE /api/v1/meal_plans/{mealPlanID}/grocery_list_items/{mealPlanGroceryListItemID}/": {
		ID:           "ArchiveMealPlanGroceryListItem",
		ResponseType: &types.MealPlanGroceryListItem{},
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/meal_plans/{mealPlanID}/tasks/": {
		ID:           "GetMealPlanTasks",
		ResponseType: &types.MealPlanTask{},
		ListRoute:    true,
		OAuth2Scopes: []string{householdMember},
	},
	"POST /api/v1/meal_plans/{mealPlanID}/tasks/": {
		ID:           "CreateMealPlanTask",
		ResponseType: &types.MealPlanTask{},
		InputType:    &types.MealPlanTaskCreationRequestInput{},
		OAuth2Scopes: []string{serviceAdmin},
	},
	"GET /api/v1/meal_plans/{mealPlanID}/tasks/{mealPlanTaskID}/": {
		ID:           "GetMealPlanTask",
		ResponseType: &types.MealPlanTask{},
		OAuth2Scopes: []string{householdMember},
	},
	"PATCH /api/v1/meal_plans/{mealPlanID}/tasks/{mealPlanTaskID}/": {
		ID:           "UpdateMealPlanTaskStatus",
		ResponseType: &types.MealPlanTask{},
		InputType:    &types.MealPlanTaskStatusChangeRequestInput{},
		OAuth2Scopes: []string{householdMember},
	},
	"POST /api/v1/meal_plans/": {
		ID:           "CreateMealPlan",
		ResponseType: &types.MealPlan{},
		InputType:    &types.MealPlanCreationRequestInput{},
		OAuth2Scopes: []string{householdAdmin},
	},
	"GET /api/v1/meal_plans/": {
		ID:           "GetMealPlans",
		ResponseType: &types.MealPlan{},
		ListRoute:    true,
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/meal_plans/{mealPlanID}/": {
		ID:           "GetMealPlan",
		ResponseType: &types.MealPlan{},
		OAuth2Scopes: []string{householdMember},
	},
	"PUT /api/v1/meal_plans/{mealPlanID}/": {
		ID:           "UpdateMealPlan",
		ResponseType: &types.MealPlan{},
		InputType:    &types.MealPlanUpdateRequestInput{},
		OAuth2Scopes: []string{householdAdmin},
	},
	"DELETE /api/v1/meal_plans/{mealPlanID}/": {
		ID:           "ArchiveMealPlan",
		ResponseType: &types.MealPlan{},
		OAuth2Scopes: []string{householdAdmin},
	},
	"POST /api/v1/meal_plans/{mealPlanID}/finalize": {
		ID:           "FinalizeMealPlan",
		ResponseType: &types.FinalizeMealPlansResponse{},
		OAuth2Scopes: []string{householdAdmin},
		// No input type for this route
	},
	"POST /api/v1/meals/": {
		ID:           "CreateMeal",
		ResponseType: &types.Meal{},
		InputType:    &types.MealCreationRequestInput{},
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/meals/": {
		ID:           "GetMeals",
		ResponseType: &types.Meal{},
		ListRoute:    true,
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/meals/search": {
		SearchRoute:  true,
		ID:           "SearchForMeals",
		ResponseType: &types.Meal{},
		OAuth2Scopes: []string{householdMember},
	},
	"DELETE /api/v1/meals/{mealID}/": {
		ID:           "ArchiveMeal",
		ResponseType: &types.Meal{},
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/meals/{mealID}/": {
		ID:           "GetMeal",
		ResponseType: &types.Meal{},
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/oauth2_clients/": {
		ID:           "GetOAuth2Clients",
		ResponseType: &types.OAuth2Client{},
		ListRoute:    true,
		OAuth2Scopes: []string{householdMember},
	},
	"POST /api/v1/oauth2_clients/": {
		ID:           "CreateOAuth2Client",
		ResponseType: &types.OAuth2ClientCreationResponse{},
		InputType:    &types.OAuth2ClientCreationRequestInput{},
		OAuth2Scopes: []string{serviceAdmin},
	},
	"GET /api/v1/oauth2_clients/{oauth2ClientID}/": {
		ID:           "GetOAuth2Client",
		ResponseType: &types.OAuth2Client{},
		OAuth2Scopes: []string{householdMember},
	},
	"DELETE /api/v1/oauth2_clients/{oauth2ClientID}/": {
		ID:           "ArchiveOAuth2Client",
		ResponseType: &types.OAuth2Client{},
		OAuth2Scopes: []string{serviceAdmin},
	},
	"POST /api/v1/recipes/{recipeID}/prep_tasks/": {
		ID:           "CreateRecipePrepTask",
		ResponseType: &types.RecipePrepTask{},
		InputType:    &types.RecipePrepTaskCreationRequestInput{},
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/recipes/{recipeID}/prep_tasks/": {
		ID:           "GetRecipePrepTasks",
		ResponseType: &types.RecipePrepTask{},
		ListRoute:    true,
		OAuth2Scopes: []string{householdMember},
	},
	"DELETE /api/v1/recipes/{recipeID}/prep_tasks/{recipePrepTaskID}/": {
		ID:           "ArchiveRecipePrepTask",
		ResponseType: &types.RecipePrepTask{},
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/recipes/{recipeID}/prep_tasks/{recipePrepTaskID}/": {
		ID:           "GetRecipePrepTask",
		ResponseType: &types.RecipePrepTask{},
		OAuth2Scopes: []string{householdMember},
	},
	"PUT /api/v1/recipes/{recipeID}/prep_tasks/{recipePrepTaskID}/": {
		ID:           "UpdateRecipePrepTask",
		ResponseType: &types.RecipePrepTask{},
		InputType:    &types.RecipePrepTaskUpdateRequestInput{},
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/recipes/{recipeID}/steps/{recipeStepID}/completion_conditions/": {
		ID:           "GetRecipeStepCompletionConditions",
		ResponseType: &types.RecipeStepCompletionCondition{},
		ListRoute:    true,
		OAuth2Scopes: []string{householdMember},
	},
	"POST /api/v1/recipes/{recipeID}/steps/{recipeStepID}/completion_conditions/": {
		ID:           "CreateRecipeStepCompletionCondition",
		ResponseType: &types.RecipeStepCompletionCondition{},
		InputType:    &types.RecipeStepCompletionConditionForExistingRecipeCreationRequestInput{},
		OAuth2Scopes: []string{householdMember},
	},
	"PUT /api/v1/recipes/{recipeID}/steps/{recipeStepID}/completion_conditions/{recipeStepCompletionConditionID}/": {
		ID:           "UpdateRecipeStepCompletionCondition",
		ResponseType: &types.RecipeStepCompletionCondition{},
		InputType:    &types.RecipeStepCompletionConditionUpdateRequestInput{},
		OAuth2Scopes: []string{householdMember},
	},
	"DELETE /api/v1/recipes/{recipeID}/steps/{recipeStepID}/completion_conditions/{recipeStepCompletionConditionID}/": {
		ID:           "ArchiveRecipeStepCompletionCondition",
		ResponseType: &types.RecipeStepCompletionCondition{},
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/recipes/{recipeID}/steps/{recipeStepID}/completion_conditions/{recipeStepCompletionConditionID}/": {
		ID:           "GetRecipeStepCompletionCondition",
		ResponseType: &types.RecipeStepCompletionCondition{},
		OAuth2Scopes: []string{householdMember},
	},
	"POST /api/v1/recipes/{recipeID}/steps/{recipeStepID}/ingredients/": {
		ID:           "CreateRecipeStepIngredient",
		ResponseType: &types.RecipeStepIngredient{},
		InputType:    &types.RecipeStepIngredientCreationRequestInput{},
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/recipes/{recipeID}/steps/{recipeStepID}/ingredients/": {
		ID:           "GetRecipeStepIngredients",
		ResponseType: &types.RecipeStepIngredient{},
		ListRoute:    true,
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/recipes/{recipeID}/steps/{recipeStepID}/ingredients/{recipeStepIngredientID}/": {
		ID:           "GetRecipeStepIngredient",
		ResponseType: &types.RecipeStepIngredient{},
		OAuth2Scopes: []string{householdMember},
	},
	"PUT /api/v1/recipes/{recipeID}/steps/{recipeStepID}/ingredients/{recipeStepIngredientID}/": {
		ID:           "UpdateRecipeStepIngredient",
		ResponseType: &types.RecipeStepIngredient{},
		InputType:    &types.RecipeStepIngredientUpdateRequestInput{},
		OAuth2Scopes: []string{householdMember},
	},
	"DELETE /api/v1/recipes/{recipeID}/steps/{recipeStepID}/ingredients/{recipeStepIngredientID}/": {
		ID:           "ArchiveRecipeStepIngredient",
		ResponseType: &types.RecipeStepIngredient{},
		OAuth2Scopes: []string{householdMember},
	},
	"POST /api/v1/recipes/{recipeID}/steps/{recipeStepID}/instruments/": {
		ID:           "CreateRecipeStepInstrument",
		ResponseType: &types.RecipeStepInstrument{},
		InputType:    &types.RecipeStepInstrumentCreationRequestInput{},
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/recipes/{recipeID}/steps/{recipeStepID}/instruments/": {
		ID:           "GetRecipeStepInstruments",
		ResponseType: &types.RecipeStepInstrument{},
		ListRoute:    true,
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/recipes/{recipeID}/steps/{recipeStepID}/instruments/{recipeStepInstrumentID}/": {
		ID:           "GetRecipeStepInstrument",
		ResponseType: &types.RecipeStepInstrument{},
		OAuth2Scopes: []string{householdMember},
	},
	"PUT /api/v1/recipes/{recipeID}/steps/{recipeStepID}/instruments/{recipeStepInstrumentID}/": {
		ID:           "UpdateRecipeStepInstrument",
		ResponseType: &types.RecipeStepInstrument{},
		InputType:    &types.RecipeStepInstrumentUpdateRequestInput{},
		OAuth2Scopes: []string{householdMember},
	},
	"DELETE /api/v1/recipes/{recipeID}/steps/{recipeStepID}/instruments/{recipeStepInstrumentID}/": {
		ID:           "ArchiveRecipeStepInstrument",
		ResponseType: &types.RecipeStepInstrument{},
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/recipes/{recipeID}/steps/{recipeStepID}/products/": {
		ID:           "GetRecipeStepProducts",
		ResponseType: &types.RecipeStepProduct{},
		ListRoute:    true,
		OAuth2Scopes: []string{householdMember},
	},
	"POST /api/v1/recipes/{recipeID}/steps/{recipeStepID}/products/": {
		ID:           "CreateRecipeStepProduct",
		ResponseType: &types.RecipeStepProduct{},
		InputType:    &types.RecipeStepProductCreationRequestInput{},
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/recipes/{recipeID}/steps/{recipeStepID}/products/{recipeStepProductID}/": {
		ID:           "GetRecipeStepProduct",
		ResponseType: &types.RecipeStepProduct{},
		OAuth2Scopes: []string{householdMember},
	},
	"PUT /api/v1/recipes/{recipeID}/steps/{recipeStepID}/products/{recipeStepProductID}/": {
		ID:           "UpdateRecipeStepProduct",
		ResponseType: &types.RecipeStepProduct{},
		InputType:    &types.RecipeStepProductUpdateRequestInput{},
		OAuth2Scopes: []string{householdMember},
	},
	"DELETE /api/v1/recipes/{recipeID}/steps/{recipeStepID}/products/{recipeStepProductID}/": {
		ID:           "ArchiveRecipeStepProduct",
		ResponseType: &types.RecipeStepProduct{},
		OAuth2Scopes: []string{householdMember},
	},
	"POST /api/v1/recipes/{recipeID}/steps/{recipeStepID}/vessels/": {
		ID:           "CreateRecipeStepVessel",
		ResponseType: &types.RecipeStepVessel{},
		InputType:    &types.RecipeStepVesselCreationRequestInput{},
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/recipes/{recipeID}/steps/{recipeStepID}/vessels/": {
		ID:           "GetRecipeStepVessels",
		ResponseType: &types.RecipeStepVessel{},
		ListRoute:    true,
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/recipes/{recipeID}/steps/{recipeStepID}/vessels/{recipeStepVesselID}/": {
		ID:           "GetRecipeStepVessel",
		ResponseType: &types.RecipeStepVessel{},
		OAuth2Scopes: []string{householdMember},
	},
	"PUT /api/v1/recipes/{recipeID}/steps/{recipeStepID}/vessels/{recipeStepVesselID}/": {
		ID:           "UpdateRecipeStepVessel",
		ResponseType: &types.RecipeStepVessel{},
		InputType:    &types.RecipeStepVesselUpdateRequestInput{},
		OAuth2Scopes: []string{householdMember},
	},
	"DELETE /api/v1/recipes/{recipeID}/steps/{recipeStepID}/vessels/{recipeStepVesselID}/": {
		ID:           "ArchiveRecipeStepVessel",
		ResponseType: &types.RecipeStepVessel{},
		OAuth2Scopes: []string{householdMember},
	},
	"POST /api/v1/recipes/{recipeID}/steps/": {
		ID:           "CreateRecipeStep",
		ResponseType: &types.RecipeStep{},
		InputType:    &types.RecipeStepCreationRequestInput{},
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/recipes/{recipeID}/steps/": {
		ID:           "GetRecipeSteps",
		ResponseType: &types.RecipeStep{},
		ListRoute:    true,
		OAuth2Scopes: []string{householdMember},
	},
	"DELETE /api/v1/recipes/{recipeID}/steps/{recipeStepID}/": {
		ID:           "ArchiveRecipeStep",
		ResponseType: &types.RecipeStep{},
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/recipes/{recipeID}/steps/{recipeStepID}/": {
		ID:           "GetRecipeStep",
		ResponseType: &types.RecipeStep{},
		OAuth2Scopes: []string{householdMember},
	},
	"PUT /api/v1/recipes/{recipeID}/steps/{recipeStepID}/": {
		ID:           "UpdateRecipeStep",
		ResponseType: &types.RecipeStep{},
		InputType:    &types.RecipeStepUpdateRequestInput{},
		OAuth2Scopes: []string{householdMember},
	},
	"POST /api/v1/recipes/{recipeID}/steps/{recipeStepID}/images": {
		ID:           "UploadMediaForRecipeStep",
		OAuth2Scopes: []string{householdMember},
		ResponseType: []*types.RecipeMedia{},
	},
	"POST /api/v1/recipes/": {
		ID:           "CreateRecipe",
		ResponseType: &types.Recipe{},
		InputType:    &types.RecipeCreationRequestInput{},
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/recipes/": {
		ID:           "GetRecipes",
		ResponseType: &types.Recipe{},
		ListRoute:    true,
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/recipes/search": {
		SearchRoute:  true,
		ID:           "SearchForRecipes",
		ResponseType: &types.Recipe{},
		OAuth2Scopes: []string{householdMember},
	},
	"DELETE /api/v1/recipes/{recipeID}/": {
		ID:           "ArchiveRecipe",
		ResponseType: &types.Recipe{},
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/recipes/{recipeID}/": {
		ID:           "GetRecipe",
		ResponseType: &types.Recipe{},
		OAuth2Scopes: []string{householdMember},
	},
	"PUT /api/v1/recipes/{recipeID}/": {
		ID:           "UpdateRecipe",
		ResponseType: &types.Recipe{},
		InputType:    &types.RecipeUpdateRequestInput{},
		OAuth2Scopes: []string{householdMember},
	},
	"POST /api/v1/recipes/{recipeID}/clone": {
		ID:           "CloneRecipe",
		ResponseType: &types.Recipe{},
		// No input type for this route
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/recipes/{recipeID}/dag": {
		ID:           "GetRecipeDAG",
		ResponseType: &types.APIError{},
		OAuth2Scopes: []string{householdMember},
	},
	"POST /api/v1/recipes/{recipeID}/images": {
		ID:           "UploadMediaForRecipe",
		OAuth2Scopes: []string{householdMember},
		ResponseType: []*types.RecipeMedia{},
	},
	"GET /api/v1/recipes/{recipeID}/mermaid": {
		ID: "GetMermaidDiagramForRecipe",
	},
	"GET /api/v1/recipes/{recipeID}/prep_steps": {
		ID:           "GetRecipeMealPlanTasks",
		ResponseType: &types.RecipePrepTaskStep{},
		OAuth2Scopes: []string{householdMember},
	},
	"POST /api/v1/recipes/{recipeID}/ratings/": {
		ID:           "CreateRecipeRating",
		ResponseType: &types.RecipeRating{},
		InputType:    &types.RecipeRatingCreationRequestInput{},
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/recipes/{recipeID}/ratings/": {
		ID:           "GetRecipeRatings",
		ResponseType: &types.RecipeRating{},
		ListRoute:    true,
		OAuth2Scopes: []string{householdMember},
	},
	"PUT /api/v1/recipes/{recipeID}/ratings/{recipeRatingID}/": {
		ID:           "UpdateRecipeRating",
		ResponseType: &types.RecipeRating{},
		InputType:    &types.RecipeRatingUpdateRequestInput{},
		OAuth2Scopes: []string{householdMember},
	},
	"DELETE /api/v1/recipes/{recipeID}/ratings/{recipeRatingID}/": {
		ID:           "ArchiveRecipeRating",
		ResponseType: &types.RecipeRating{},
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/recipes/{recipeID}/ratings/{recipeRatingID}/": {
		ID:           "GetRecipeRating",
		ResponseType: &types.RecipeRating{},
		OAuth2Scopes: []string{householdMember},
	},
	"POST /api/v1/settings/": {
		ID:           "CreateServiceSetting",
		ResponseType: &types.ServiceSetting{},
		InputType:    &types.ServiceSettingCreationRequestInput{},
		OAuth2Scopes: []string{serviceAdmin},
	},
	"GET /api/v1/settings/": {
		ID:           "GetServiceSettings",
		ResponseType: &types.ServiceSetting{},
		ListRoute:    true,
		OAuth2Scopes: []string{householdMember},
	},
	"POST /api/v1/settings/configurations/": {
		ID:           "CreateServiceSettingConfiguration",
		ResponseType: &types.ServiceSettingConfiguration{},
		InputType:    &types.ServiceSettingConfigurationCreationRequestInput{},
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/settings/configurations/household": {
		ID:           "GetServiceSettingConfigurationsForHousehold",
		ListRoute:    true,
		ResponseType: &types.ServiceSettingConfiguration{},
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/settings/configurations/user": {
		ID:           "GetServiceSettingConfigurationsForUser",
		ListRoute:    true,
		ResponseType: &types.ServiceSettingConfiguration{},
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/settings/configurations/user/{serviceSettingConfigurationName}": {
		ID:           "GetServiceSettingConfigurationByName",
		ResponseType: &types.ServiceSettingConfiguration{},
		ListRoute:    true,
		OAuth2Scopes: []string{householdMember},
	},
	"DELETE /api/v1/settings/configurations/{serviceSettingConfigurationID}": {
		ID:           "ArchiveServiceSettingConfiguration",
		ResponseType: &types.ServiceSettingConfiguration{},
		OAuth2Scopes: []string{householdMember},
	},
	"PUT /api/v1/settings/configurations/{serviceSettingConfigurationID}": {
		ID:           "UpdateServiceSettingConfiguration",
		ResponseType: &types.ServiceSettingConfiguration{},
		InputType:    &types.ServiceSettingConfigurationUpdateRequestInput{},
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/settings/search": {
		SearchRoute:  true,
		ID:           "SearchForServiceSettings",
		ResponseType: &types.ServiceSetting{},
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/settings/{serviceSettingID}/": {
		ID:           "GetServiceSetting",
		ResponseType: &types.ServiceSetting{},
		OAuth2Scopes: []string{householdMember},
	},
	"DELETE /api/v1/settings/{serviceSettingID}/": {
		ID:           "ArchiveServiceSetting",
		ResponseType: &types.ServiceSetting{},
		OAuth2Scopes: []string{serviceAdmin},
	},
	"POST /api/v1/user_ingredient_preferences/": {
		ID:           "CreateUserIngredientPreference",
		ResponseType: &types.UserIngredientPreference{},
		InputType:    &types.UserIngredientPreferenceCreationRequestInput{},
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/user_ingredient_preferences/": {
		ID:           "GetUserIngredientPreferences",
		ResponseType: &types.UserIngredientPreference{},
		ListRoute:    true,
		OAuth2Scopes: []string{householdMember},
	},
	"PUT /api/v1/user_ingredient_preferences/{userIngredientPreferenceID}/": {
		ID:           "UpdateUserIngredientPreference",
		ResponseType: &types.UserIngredientPreference{},
		InputType:    &types.UserIngredientPreferenceUpdateRequestInput{},
		OAuth2Scopes: []string{householdMember},
	},
	"DELETE /api/v1/user_ingredient_preferences/{userIngredientPreferenceID}/": {
		ID:           "ArchiveUserIngredientPreference",
		ResponseType: &types.UserIngredientPreference{},
		OAuth2Scopes: []string{householdMember},
	},
	"POST /api/v1/user_notifications/": {
		ID:           "CreateUserNotification",
		ResponseType: &types.UserNotification{},
		InputType:    &types.UserNotificationCreationRequestInput{},
		OAuth2Scopes: []string{serviceAdmin},
	},
	"GET /api/v1/user_notifications/": {
		ID:           "GetUserNotifications",
		ResponseType: &types.UserNotification{},
		ListRoute:    true,
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/user_notifications/{userNotificationID}/": {
		ID:           "GetUserNotification",
		ResponseType: &types.UserNotification{},
		OAuth2Scopes: []string{householdMember},
	},
	"PATCH /api/v1/user_notifications/{userNotificationID}/": {
		ID:           "UpdateUserNotification",
		InputType:    &types.UserNotificationUpdateRequestInput{},
		ResponseType: &types.UserNotification{},
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/users/": {
		ID:           "GetUsers",
		ResponseType: &types.User{},
		ListRoute:    true,
		OAuth2Scopes: []string{householdMember},
	},
	"POST /api/v1/users/avatar/upload": {
		ID:           "UploadUserAvatar",
		InputType:    &types.AvatarUpdateInput{},
		ResponseType: &types.User{},
		OAuth2Scopes: []string{householdMember},
	},
	"PUT /api/v1/users/details": {
		ID:           "UpdateUserDetails",
		ResponseType: &types.User{},
		InputType:    &types.UserDetailsUpdateRequestInput{},
		OAuth2Scopes: []string{householdMember},
	},
	"PUT /api/v1/users/email_address": {
		ID:           "UpdateUserEmailAddress",
		ResponseType: &types.User{},
		InputType:    &types.UserEmailAddressUpdateInput{},
		OAuth2Scopes: []string{householdMember},
	},
	"POST /api/v1/users/email_address_verification": {
		ID:           "VerifyUserEmailAddress",
		ResponseType: &types.User{},
		InputType:    &types.EmailAddressVerificationRequestInput{},
		OAuth2Scopes: []string{householdMember},
	},
	"POST /api/v1/users/household/select": {
		ID:           "ChangeActiveHousehold",
		ResponseType: &types.Household{},
		InputType:    &types.ChangeActiveHouseholdInput{},
		OAuth2Scopes: []string{householdMember},
	},
	"PUT /api/v1/users/password/new": {
		ID:           "UpdatePassword",
		InputType:    &types.PasswordUpdateInput{},
		OAuth2Scopes: []string{householdMember},
	},
	"POST /api/v1/users/permissions/check": {
		ID:           "CheckPermissions",
		ResponseType: &types.UserPermissionsResponse{},
		InputType:    &types.UserPermissionsRequestInput{},
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/users/search": {
		SearchRoute:  true,
		ID:           "SearchForUsers",
		ResponseType: &types.User{},
		ListRoute:    true,
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/users/self": {
		ID:           "GetSelf",
		ResponseType: &types.User{},
		OAuth2Scopes: []string{householdMember},
	},
	"POST /api/v1/users/totp_secret/new": {
		ID:           "RefreshTOTPSecret",
		InputType:    &types.TOTPSecretRefreshInput{},
		ResponseType: &types.TOTPSecretRefreshResponse{},
		OAuth2Scopes: []string{householdMember},
	},
	"PUT /api/v1/users/username": {
		ID:           "UpdateUserUsername",
		ResponseType: &types.User{},
		InputType:    &types.UsernameUpdateInput{},
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/users/{userID}/": {
		ID:           "GetUser",
		ResponseType: &types.User{},
		OAuth2Scopes: []string{householdMember},
	},
	"DELETE /api/v1/users/{userID}/": {
		ID:           "ArchiveUser",
		ResponseType: &types.User{},
		OAuth2Scopes: []string{householdMember},
	},
	"POST /api/v1/valid_ingredient_groups/": {
		ID:           "CreateValidIngredientGroup",
		ResponseType: &types.ValidIngredientGroup{},
		InputType:    &types.ValidIngredientGroupCreationRequestInput{},
		OAuth2Scopes: []string{serviceAdmin},
	},
	"GET /api/v1/valid_ingredient_groups/": {
		ID:           "GetValidIngredientGroups",
		ResponseType: &types.ValidIngredientGroup{},
		ListRoute:    true,
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/valid_ingredient_groups/search": {
		SearchRoute:  true,
		ID:           "SearchForValidIngredientGroups",
		ResponseType: &types.ValidIngredientGroup{},
		OAuth2Scopes: []string{householdMember},
	},
	"DELETE /api/v1/valid_ingredient_groups/{validIngredientGroupID}/": {
		ID:           "ArchiveValidIngredientGroup",
		ResponseType: &types.ValidIngredientGroup{},
		OAuth2Scopes: []string{serviceAdmin},
	},
	"GET /api/v1/valid_ingredient_groups/{validIngredientGroupID}/": {
		ID:           "GetValidIngredientGroup",
		ResponseType: &types.ValidIngredientGroup{},
		OAuth2Scopes: []string{householdMember},
	},
	"PUT /api/v1/valid_ingredient_groups/{validIngredientGroupID}/": {
		ID:           "UpdateValidIngredientGroup",
		ResponseType: &types.ValidIngredientGroup{},
		InputType:    &types.ValidIngredientGroupUpdateRequestInput{},
		OAuth2Scopes: []string{serviceAdmin},
	},
	"POST /api/v1/valid_ingredient_measurement_units/": {
		ID:           "CreateValidIngredientMeasurementUnit",
		ResponseType: &types.ValidIngredientMeasurementUnit{},
		InputType:    &types.ValidIngredientMeasurementUnitCreationRequestInput{},
		OAuth2Scopes: []string{serviceAdmin},
	},
	"GET /api/v1/valid_ingredient_measurement_units/": {
		ID:           "GetValidIngredientMeasurementUnits",
		ResponseType: &types.ValidIngredientMeasurementUnit{},
		ListRoute:    true,
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/valid_ingredient_measurement_units/by_ingredient/{validIngredientID}/": {
		ID:           "GetValidIngredientMeasurementUnitsByIngredient",
		ResponseType: &types.ValidIngredientMeasurementUnit{},
		ListRoute:    true,
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/valid_ingredient_measurement_units/by_measurement_unit/{validMeasurementUnitID}/": {
		ID:           "GetValidIngredientMeasurementUnitsByMeasurementUnit",
		ResponseType: &types.ValidIngredientMeasurementUnit{},
		ListRoute:    true,
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/valid_ingredient_measurement_units/{validIngredientMeasurementUnitID}/": {
		ID:           "GetValidIngredientMeasurementUnit",
		ResponseType: &types.ValidIngredientMeasurementUnit{},
		OAuth2Scopes: []string{householdMember},
	},
	"PUT /api/v1/valid_ingredient_measurement_units/{validIngredientMeasurementUnitID}/": {
		ID:           "UpdateValidIngredientMeasurementUnit",
		ResponseType: &types.ValidIngredientMeasurementUnit{},
		InputType:    &types.ValidIngredientMeasurementUnitUpdateRequestInput{},
		OAuth2Scopes: []string{serviceAdmin},
	},
	"DELETE /api/v1/valid_ingredient_measurement_units/{validIngredientMeasurementUnitID}/": {
		ID:           "ArchiveValidIngredientMeasurementUnit",
		ResponseType: &types.ValidIngredientMeasurementUnit{},
		OAuth2Scopes: []string{serviceAdmin},
	},
	"GET /api/v1/valid_ingredient_preparations/": {
		ID:           "GetValidIngredientPreparations",
		ResponseType: &types.ValidIngredientPreparation{},
		ListRoute:    true,
		OAuth2Scopes: []string{householdMember},
	},
	"POST /api/v1/valid_ingredient_preparations/": {
		ID:           "CreateValidIngredientPreparation",
		ResponseType: &types.ValidIngredientPreparation{},
		InputType:    &types.ValidIngredientPreparationCreationRequestInput{},
		OAuth2Scopes: []string{serviceAdmin},
	},
	"GET /api/v1/valid_ingredient_preparations/by_ingredient/{validIngredientID}/": {
		ID:           "GetValidIngredientPreparationsByIngredient",
		ResponseType: &types.ValidIngredientPreparation{},
		ListRoute:    true,
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/valid_ingredient_preparations/by_preparation/{validPreparationID}/": {
		ID:           "GetValidIngredientPreparationsByPreparation",
		ResponseType: &types.ValidIngredientPreparation{},
		ListRoute:    true,
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/valid_ingredient_preparations/{validIngredientPreparationID}/": {
		ID:           "GetValidIngredientPreparation",
		ResponseType: &types.ValidIngredientPreparation{},
		OAuth2Scopes: []string{householdMember},
	},
	"PUT /api/v1/valid_ingredient_preparations/{validIngredientPreparationID}/": {
		ID:           "UpdateValidIngredientPreparation",
		ResponseType: &types.ValidIngredientPreparation{},
		InputType:    &types.ValidIngredientPreparationUpdateRequestInput{},
		OAuth2Scopes: []string{serviceAdmin},
	},
	"DELETE /api/v1/valid_ingredient_preparations/{validIngredientPreparationID}/": {
		ID:           "ArchiveValidIngredientPreparation",
		ResponseType: &types.ValidIngredientPreparation{},
		OAuth2Scopes: []string{serviceAdmin},
	},
	"POST /api/v1/valid_ingredient_state_ingredients/": {
		ID:           "CreateValidIngredientStateIngredient",
		ResponseType: &types.ValidIngredientStateIngredient{},
		InputType:    &types.ValidIngredientStateIngredientCreationRequestInput{},
		OAuth2Scopes: []string{serviceAdmin},
	},
	"GET /api/v1/valid_ingredient_state_ingredients/": {
		ID:           "GetValidIngredientStateIngredients",
		ResponseType: &types.ValidIngredientStateIngredient{},
		ListRoute:    true,
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/valid_ingredient_state_ingredients/by_ingredient/{validIngredientID}/": {
		ID:           "GetValidIngredientStateIngredientsByIngredient",
		ResponseType: &types.ValidIngredientStateIngredient{},
		ListRoute:    true,
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/valid_ingredient_state_ingredients/by_ingredient_state/{validIngredientStateID}/": {
		ID:           "GetValidIngredientStateIngredientsByIngredientState",
		ResponseType: &types.ValidIngredientStateIngredient{},
		ListRoute:    true,
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/valid_ingredient_state_ingredients/{validIngredientStateIngredientID}/": {
		ID:           "GetValidIngredientStateIngredient",
		ResponseType: &types.ValidIngredientStateIngredient{},
		OAuth2Scopes: []string{householdMember},
	},
	"PUT /api/v1/valid_ingredient_state_ingredients/{validIngredientStateIngredientID}/": {
		ID:           "UpdateValidIngredientStateIngredient",
		ResponseType: &types.ValidIngredientStateIngredient{},
		InputType:    &types.ValidIngredientStateIngredientUpdateRequestInput{},
		OAuth2Scopes: []string{serviceAdmin},
	},
	"DELETE /api/v1/valid_ingredient_state_ingredients/{validIngredientStateIngredientID}/": {
		ID:           "ArchiveValidIngredientStateIngredient",
		ResponseType: &types.ValidIngredientStateIngredient{},
		OAuth2Scopes: []string{serviceAdmin},
	},
	"POST /api/v1/valid_ingredient_states/": {
		ID:           "CreateValidIngredientState",
		ResponseType: &types.ValidIngredientState{},
		InputType:    &types.ValidIngredientStateCreationRequestInput{},
		OAuth2Scopes: []string{serviceAdmin},
	},
	"GET /api/v1/valid_ingredient_states/": {
		ID:           "GetValidIngredientStates",
		ResponseType: &types.ValidIngredientState{},
		ListRoute:    true,
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/valid_ingredient_states/search": {
		SearchRoute:  true,
		ID:           "SearchForValidIngredientStates",
		ResponseType: &types.ValidIngredientState{},
		OAuth2Scopes: []string{householdMember},
	},
	"PUT /api/v1/valid_ingredient_states/{validIngredientStateID}/": {
		ID:           "UpdateValidIngredientState",
		ResponseType: &types.ValidIngredientState{},
		InputType:    &types.ValidIngredientStateUpdateRequestInput{},
		OAuth2Scopes: []string{serviceAdmin},
	},
	"DELETE /api/v1/valid_ingredient_states/{validIngredientStateID}/": {
		ID:           "ArchiveValidIngredientState",
		ResponseType: &types.ValidIngredientState{},
		OAuth2Scopes: []string{serviceAdmin},
	},
	"GET /api/v1/valid_ingredient_states/{validIngredientStateID}/": {
		ID:           "GetValidIngredientState",
		ResponseType: &types.ValidIngredientState{},
		OAuth2Scopes: []string{householdMember},
	},
	"POST /api/v1/valid_ingredients/": {
		ID:           "CreateValidIngredient",
		ResponseType: &types.ValidIngredient{},
		InputType:    &types.ValidIngredientCreationRequestInput{},
		OAuth2Scopes: []string{serviceAdmin},
	},
	"GET /api/v1/valid_ingredients/": {
		ID:           "GetValidIngredients",
		ResponseType: &types.ValidIngredient{},
		ListRoute:    true,
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/valid_ingredients/by_preparation/{validPreparationID}/": {
		ID:           "GetValidIngredientsByPreparation",
		ListRoute:    true,
		SearchRoute:  true,
		ResponseType: &types.ValidIngredient{},
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/valid_ingredients/random": {
		ID:           "GetRandomValidIngredient",
		ResponseType: &types.ValidIngredient{},
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/valid_ingredients/search": {
		SearchRoute:  true,
		ID:           "SearchForValidIngredients",
		ResponseType: &types.ValidIngredient{},
		OAuth2Scopes: []string{householdMember},
	},
	"PUT /api/v1/valid_ingredients/{validIngredientID}/": {
		ID:           "UpdateValidIngredient",
		ResponseType: &types.ValidIngredient{},
		InputType:    &types.ValidIngredientUpdateRequestInput{},
		OAuth2Scopes: []string{serviceAdmin},
	},
	"DELETE /api/v1/valid_ingredients/{validIngredientID}/": {
		ID:           "ArchiveValidIngredient",
		ResponseType: &types.ValidIngredient{},
		OAuth2Scopes: []string{serviceAdmin},
	},
	"GET /api/v1/valid_ingredients/{validIngredientID}/": {
		ID:           "GetValidIngredient",
		ResponseType: &types.ValidIngredient{},
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/valid_instruments/": {
		ID:           "GetValidInstruments",
		ResponseType: &types.ValidInstrument{},
		ListRoute:    true,
		OAuth2Scopes: []string{householdMember},
	},
	"POST /api/v1/valid_instruments/": {
		ID:           "CreateValidInstrument",
		ResponseType: &types.ValidInstrument{},
		InputType:    &types.ValidInstrumentCreationRequestInput{},
		OAuth2Scopes: []string{serviceAdmin},
	},
	"GET /api/v1/valid_instruments/random": {
		ID:           "GetRandomValidInstrument",
		ResponseType: &types.ValidInstrument{},
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/valid_instruments/search": {
		SearchRoute:  true,
		ID:           "SearchForValidInstruments",
		ResponseType: &types.ValidInstrument{},
		OAuth2Scopes: []string{householdMember},
	},
	"DELETE /api/v1/valid_instruments/{validInstrumentID}/": {
		ID:           "ArchiveValidInstrument",
		ResponseType: &types.ValidInstrument{},
		OAuth2Scopes: []string{serviceAdmin},
	},
	"GET /api/v1/valid_instruments/{validInstrumentID}/": {
		ID:           "GetValidInstrument",
		ResponseType: &types.ValidInstrument{},
		OAuth2Scopes: []string{householdMember},
	},
	"PUT /api/v1/valid_instruments/{validInstrumentID}/": {
		ID:           "UpdateValidInstrument",
		ResponseType: &types.ValidInstrument{},
		InputType:    &types.ValidInstrumentUpdateRequestInput{},
		OAuth2Scopes: []string{serviceAdmin},
	},
	"POST /api/v1/valid_measurement_conversions/": {
		ID:           "CreateValidMeasurementUnitConversion",
		ResponseType: &types.ValidMeasurementUnitConversion{},
		InputType:    &types.ValidMeasurementUnitConversionCreationRequestInput{},
		OAuth2Scopes: []string{serviceAdmin},
	},
	"GET /api/v1/valid_measurement_conversions/from_unit/{validMeasurementUnitID}": {
		ID:           "GetValidMeasurementUnitConversionsFromUnit",
		ResponseType: &types.ValidMeasurementUnitConversion{},
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/valid_measurement_conversions/to_unit/{validMeasurementUnitID}": {
		ID:           "ValidMeasurementUnitConversionsToUnit",
		ResponseType: &types.ValidMeasurementUnitConversion{},
		OAuth2Scopes: []string{householdMember},
	},
	"PUT /api/v1/valid_measurement_conversions/{validMeasurementUnitConversionID}/": {
		ID:           "UpdateValidMeasurementUnitConversion",
		ResponseType: &types.ValidMeasurementUnitConversion{},
		InputType:    &types.ValidMeasurementUnitConversionUpdateRequestInput{},
		OAuth2Scopes: []string{serviceAdmin},
	},
	"DELETE /api/v1/valid_measurement_conversions/{validMeasurementUnitConversionID}/": {
		ID:           "ArchiveValidMeasurementUnitConversion",
		ResponseType: &types.ValidMeasurementUnitConversion{},
		OAuth2Scopes: []string{serviceAdmin},
	},
	"GET /api/v1/valid_measurement_conversions/{validMeasurementUnitConversionID}/": {
		ID:           "GetValidMeasurementUnitConversion",
		ResponseType: &types.ValidMeasurementUnitConversion{},
		OAuth2Scopes: []string{householdMember},
	},
	"POST /api/v1/valid_measurement_units/": {
		ID:           "CreateValidMeasurementUnit",
		ResponseType: &types.ValidMeasurementUnit{},
		InputType:    &types.ValidMeasurementUnitCreationRequestInput{},
		OAuth2Scopes: []string{serviceAdmin},
	},
	"GET /api/v1/valid_measurement_units/": {
		ID:           "GetValidMeasurementUnits",
		ResponseType: &types.ValidMeasurementUnit{},
		ListRoute:    true,
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/valid_measurement_units/by_ingredient/{validIngredientID}": {
		ID:           "SearchValidMeasurementUnitsByIngredient",
		ResponseType: &types.ValidMeasurementUnit{},
		ListRoute:    true,
		SearchRoute:  true,
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/valid_measurement_units/search": {
		SearchRoute:  true,
		ID:           "SearchForValidMeasurementUnits",
		ResponseType: &types.ValidMeasurementUnit{},
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/valid_measurement_units/{validMeasurementUnitID}/": {
		ID:           "GetValidMeasurementUnit",
		ResponseType: &types.ValidMeasurementUnit{},
		OAuth2Scopes: []string{householdMember},
	},
	"PUT /api/v1/valid_measurement_units/{validMeasurementUnitID}/": {
		ID:           "UpdateValidMeasurementUnit",
		ResponseType: &types.ValidMeasurementUnit{},
		InputType:    &types.ValidMeasurementUnitUpdateRequestInput{},
		OAuth2Scopes: []string{serviceAdmin},
	},
	"DELETE /api/v1/valid_measurement_units/{validMeasurementUnitID}/": {
		ID:           "ArchiveValidMeasurementUnit",
		ResponseType: &types.ValidMeasurementUnit{},
		OAuth2Scopes: []string{serviceAdmin},
	},
	"GET /api/v1/valid_preparation_instruments/": {
		ID:           "GetValidPreparationInstruments",
		ResponseType: &types.ValidPreparationInstrument{},
		ListRoute:    true,
		OAuth2Scopes: []string{householdMember},
	},
	"POST /api/v1/valid_preparation_instruments/": {
		ID:           "CreateValidPreparationInstrument",
		ResponseType: &types.ValidPreparationInstrument{},
		InputType:    &types.ValidPreparationInstrumentCreationRequestInput{},
		OAuth2Scopes: []string{serviceAdmin},
	},
	"GET /api/v1/valid_preparation_instruments/by_instrument/{validInstrumentID}/": {
		ID:           "GetValidPreparationInstrumentsByInstrument",
		ResponseType: &types.ValidPreparationInstrument{},
		ListRoute:    true,
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/valid_preparation_instruments/by_preparation/{validPreparationID}/": {
		ID:           "GetValidPreparationInstrumentsByPreparation",
		ResponseType: &types.ValidPreparationInstrument{},
		ListRoute:    true,
		OAuth2Scopes: []string{householdMember},
	},
	"DELETE /api/v1/valid_preparation_instruments/{validPreparationVesselID}/": {
		ID:           "ArchiveValidPreparationInstrument",
		ResponseType: &types.ValidPreparationInstrument{},
		OAuth2Scopes: []string{serviceAdmin},
	},
	"GET /api/v1/valid_preparation_instruments/{validPreparationVesselID}/": {
		ID:           "GetValidPreparationInstrument",
		ResponseType: &types.ValidPreparationInstrument{},
		OAuth2Scopes: []string{householdMember},
	},
	"PUT /api/v1/valid_preparation_instruments/{validPreparationVesselID}/": {
		ID:           "UpdateValidPreparationInstrument",
		ResponseType: &types.ValidPreparationInstrument{},
		InputType:    &types.ValidPreparationInstrumentUpdateRequestInput{},
		OAuth2Scopes: []string{serviceAdmin},
	},
	"POST /api/v1/valid_preparation_vessels/": {
		ID:           "CreateValidPreparationVessel",
		ResponseType: &types.ValidPreparationVessel{},
		InputType:    &types.ValidPreparationVesselCreationRequestInput{},
		OAuth2Scopes: []string{serviceAdmin},
	},
	"GET /api/v1/valid_preparation_vessels/": {
		ID:           "GetValidPreparationVessels",
		ResponseType: &types.ValidPreparationVessel{},
		ListRoute:    true,
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/valid_preparation_vessels/by_preparation/{validPreparationID}/": {
		ID:           "GetValidPreparationVesselsByPreparation",
		ResponseType: &types.ValidPreparationVessel{},
		ListRoute:    true,
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/valid_preparation_vessels/by_vessel/{ValidVesselID}/": {
		ID:           "GetValidPreparationVesselsByVessel",
		ResponseType: &types.ValidPreparationVessel{},
		ListRoute:    true,
		OAuth2Scopes: []string{householdMember},
	},
	"PUT /api/v1/valid_preparation_vessels/{validPreparationVesselID}/": {
		ID:           "UpdateValidPreparationVessel",
		ResponseType: &types.ValidPreparationVessel{},
		InputType:    &types.ValidPreparationVesselUpdateRequestInput{},
		OAuth2Scopes: []string{serviceAdmin},
	},
	"DELETE /api/v1/valid_preparation_vessels/{validPreparationVesselID}/": {
		ID:           "ArchiveValidPreparationVessel",
		ResponseType: &types.ValidPreparationVessel{},
		OAuth2Scopes: []string{serviceAdmin},
	},
	"GET /api/v1/valid_preparation_vessels/{validPreparationVesselID}/": {
		ID:           "GetValidPreparationVessel",
		ResponseType: &types.ValidPreparationVessel{},
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/valid_preparations/": {
		ID:           "GetValidPreparations",
		ResponseType: &types.ValidPreparation{},
		ListRoute:    true,
		OAuth2Scopes: []string{householdMember},
	},
	"POST /api/v1/valid_preparations/": {
		ID:           "CreateValidPreparation",
		ResponseType: &types.ValidPreparation{},
		InputType:    &types.ValidPreparationCreationRequestInput{},
		OAuth2Scopes: []string{serviceAdmin},
	},
	"GET /api/v1/valid_preparations/random": {
		ID:           "GetRandomValidPreparation",
		ResponseType: &types.ValidPreparation{},
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/valid_preparations/search": {
		SearchRoute:  true,
		ID:           "SearchForValidPreparations",
		ResponseType: &types.ValidPreparation{},
		OAuth2Scopes: []string{householdMember},
	},
	"PUT /api/v1/valid_preparations/{validPreparationID}/": {
		ID:           "UpdateValidPreparation",
		ResponseType: &types.ValidPreparation{},
		InputType:    &types.ValidPreparationUpdateRequestInput{},
		OAuth2Scopes: []string{serviceAdmin},
	},
	"DELETE /api/v1/valid_preparations/{validPreparationID}/": {
		ID:           "ArchiveValidPreparation",
		ResponseType: &types.ValidPreparation{},
		OAuth2Scopes: []string{serviceAdmin},
	},
	"GET /api/v1/valid_preparations/{validPreparationID}/": {
		ID:           "GetValidPreparation",
		ResponseType: &types.ValidPreparation{},
		OAuth2Scopes: []string{householdMember},
	},
	"POST /api/v1/valid_vessels/": {
		ID:           "CreateValidVessel",
		ResponseType: &types.ValidVessel{},
		InputType:    &types.ValidVesselCreationRequestInput{},
		OAuth2Scopes: []string{serviceAdmin},
	},
	"GET /api/v1/valid_vessels/": {
		ID:           "GetValidVessels",
		ResponseType: &types.ValidVessel{},
		ListRoute:    true,
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/valid_vessels/random": {
		ID:           "GetRandomValidVessel",
		ResponseType: &types.ValidVessel{},
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/valid_vessels/search": {
		SearchRoute:  true,
		ID:           "SearchForValidVessels",
		ResponseType: &types.ValidVessel{},
		OAuth2Scopes: []string{householdMember},
	},
	"GET /api/v1/valid_vessels/{validVesselID}/": {
		ID:           "GetValidVessel",
		ResponseType: &types.ValidVessel{},
		OAuth2Scopes: []string{householdMember},
	},
	"PUT /api/v1/valid_vessels/{validVesselID}/": {
		ID:           "UpdateValidVessel",
		ResponseType: &types.ValidVessel{},
		InputType:    &types.ValidVesselUpdateRequestInput{},
		OAuth2Scopes: []string{serviceAdmin},
	},
	"DELETE /api/v1/valid_vessels/{validVesselID}/": {
		ID:           "ArchiveValidVessel",
		ResponseType: &types.ValidVessel{},
		OAuth2Scopes: []string{serviceAdmin},
	},
	"GET /api/v1/webhooks/": {
		ID:           "GetWebhooks",
		ResponseType: &types.Webhook{},
		ListRoute:    true,
		OAuth2Scopes: []string{householdMember},
	},
	"POST /api/v1/webhooks/": {
		ID:           "CreateWebhook",
		ResponseType: &types.Webhook{},
		InputType:    &types.WebhookCreationRequestInput{},
		OAuth2Scopes: []string{householdAdmin},
	},
	"GET /api/v1/webhooks/{webhookID}/": {
		ID:           "GetWebhook",
		ResponseType: &types.Webhook{},
		OAuth2Scopes: []string{householdMember},
	},
	"DELETE /api/v1/webhooks/{webhookID}/": {
		ID:           "ArchiveWebhook",
		ResponseType: &types.Webhook{},
		OAuth2Scopes: []string{householdAdmin},
	},
	"POST /api/v1/webhooks/{webhookID}/trigger_events": {
		ID:           "CreateWebhookTriggerEvent",
		ResponseType: &types.WebhookTriggerEvent{},
		InputType:    &types.WebhookTriggerEventCreationRequestInput{},
		OAuth2Scopes: []string{householdAdmin},
	},
	"DELETE /api/v1/webhooks/{webhookID}/trigger_events/{webhookTriggerEventID}/": {
		ID:           "ArchiveWebhookTriggerEvent",
		ResponseType: &types.WebhookTriggerEvent{},
		OAuth2Scopes: []string{householdAdmin},
	},
	"POST /api/v1/workers/finalize_meal_plans": {
		ID:           "RunFinalizeMealPlanWorker",
		ResponseType: &types.FinalizeMealPlansRequest{},
		InputType:    &types.FinalizeMealPlansRequest{},
		OAuth2Scopes: []string{householdAdmin},
	},
	"POST /api/v1/workers/meal_plan_grocery_list_init": {
		ID: "RunMealPlanGroceryListInitializerWorker",
		// no input or output types for this route
		OAuth2Scopes: []string{householdAdmin},
	},
	"POST /api/v1/workers/meal_plan_tasks": {
		ID: "RunMealPlanTaskCreatorWorker",
		// no input or output types for this route
		OAuth2Scopes: []string{householdAdmin},
	},
	"GET /auth/status": {
		ID:           "GetAuthStatus",
		ResponseType: &types.UserStatusResponse{},
		OAuth2Scopes: []string{householdMember},
		// no input type for this route
	},
	"GET /auth/{auth_provider}": {
		// we don't really have control over this route
	},
	"GET /auth/{auth_provider}/callback": {
		// we don't really have control over this route
	},
	"GET /oauth2/authorize": {
		// we don't really have control over this route
	},
	"POST /oauth2/token": {
		// we don't really have control over this route
	},
	"POST /users/": {
		ID:           "CreateUser",
		Description:  "Creates a new user",
		ResponseType: &types.UserCreationResponse{},
		InputType:    &types.UserRegistrationInput{},
	},
	"POST /users/email_address/verify": {
		ID:           "VerifyEmailAddress",
		ResponseType: &types.User{},
		InputType:    &types.EmailAddressVerificationRequestInput{},
	},
	"POST /users/login": {
		ID:           "Login",
		ResponseType: &types.UserStatusResponse{},
		InputType:    &types.UserLoginInput{},
	},
	"POST /users/login/admin": {
		ID:           "AdminLogin",
		ResponseType: &types.UserStatusResponse{},
		InputType:    &types.UserLoginInput{},
	},
	"POST /users/login/jwt": {
		ID:           "LoginForJWT",
		ResponseType: &types.JWTResponse{},
		InputType:    &types.UserLoginInput{},
	},
	"POST /users/login/jwt/admin": {
		ID:           "AdminLoginForJWT",
		ResponseType: &types.JWTResponse{},
		InputType:    &types.UserLoginInput{},
	},
	"POST /users/logout": {
		ID:           "Logout",
		ResponseType: &types.UserStatusResponse{},
	},
	"POST /users/password/reset": {
		ID:           "RequestPasswordResetToken",
		ResponseType: &types.PasswordResetToken{},
		InputType:    &types.PasswordResetTokenCreationRequestInput{},
	},
	"POST /users/password/reset/redeem": {
		ID:           "RedeemPasswordResetToken",
		Description:  "Operation for redeeming a password reset token",
		ResponseType: &types.User{},
		InputType:    &types.PasswordResetTokenRedemptionRequestInput{},
	},
	"POST /users/totp_secret/verify": {
		ID:           "VerifyTOTPSecret",
		ResponseType: &types.User{},
		InputType:    &types.TOTPSecretVerificationInput{},
	},
	"POST /users/username/reminder": {
		ID:           "RequestUsernameReminder",
		ResponseType: &types.User{},
		InputType:    &types.UsernameReminderRequestInput{},
	},
}
