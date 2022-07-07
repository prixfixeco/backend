package keys

const (
	// HouseholdSubscriptionPlanIDKey is the standard key for referring to a household subscription plan ID.
	HouseholdSubscriptionPlanIDKey = "household_subscription_plan.id"
	// RequesterIDKey is the standard key for referring to a requesting user's ID.
	RequesterIDKey = "request.made_by"
	// HouseholdIDKey is the standard key for referring to a household ID.
	HouseholdIDKey = "household.id"
	// HouseholdInvitationIDKey is the standard key for referring to a household ID.
	HouseholdInvitationIDKey = "household_invitation.id"
	// HouseholdInvitationTokenKey is the standard key for referring to a household ID.
	HouseholdInvitationTokenKey = "household_invitation.token"
	// ActiveHouseholdIDKey is the standard key for referring to an active household ID.
	ActiveHouseholdIDKey = "active_household.id"
	// UserIDKey is the standard key for referring to a user ID.
	UserIDKey = "user.id"
	// UserEmailAddressKey is the standard key for referring to a user's email address.
	UserEmailAddressKey = "user.email_address"
	// UserIsServiceAdminKey is the standard key for referring to a user's admin status.
	UserIsServiceAdminKey = "user.is_admin"
	// UsernameKey is the standard key for referring to a username.
	UsernameKey = "user.username"
	// ServiceRoleKey is the standard key for referring to a username.
	ServiceRoleKey = "user.service_role"
	// NameKey is the standard key for referring to a name.
	NameKey = "name"
	// SpanIDKey is the standard key for referring to a span ID.
	SpanIDKey = "span.id"
	// TraceIDKey is the standard key for referring to a trace ID.
	TraceIDKey = "trace.id"
	// FilterCreatedAfterKey is the standard key for referring to a types.QueryFilter's CreatedAfter field.
	FilterCreatedAfterKey = "query_filter.created_after"
	// FilterCreatedBeforeKey is the standard key for referring to a types.QueryFilter's CreatedBefore field.
	FilterCreatedBeforeKey = "query_filter.created_before"
	// FilterUpdatedAfterKey is the standard key for referring to a types.QueryFilter's UpdatedAfter field.
	FilterUpdatedAfterKey = "query_filter.updated_after"
	// FilterUpdatedBeforeKey is the standard key for referring to a types.QueryFilter's UpdatedAfter field.
	FilterUpdatedBeforeKey = "query_filter.updated_before"
	// FilterSortByKey is the standard key for referring to a types.QueryFilter's SortBy field.
	FilterSortByKey = "query_filter.sort_by"
	// FilterPageKey is the standard key for referring to a types.QueryFilter's page.
	FilterPageKey = "query_filter.page"
	// FilterLimitKey is the standard key for referring to a types.QueryFilter's limit.
	FilterLimitKey = "query_filter.limit"
	// FilterIsNilKey is the standard key for referring to a types.QueryFilter's null status.
	FilterIsNilKey = "query_filter.is_nil"
	// APIClientClientIDKey is the standard key for referring to an API client's database ID.
	APIClientClientIDKey = "api_client.client.id"
	// APIClientDatabaseIDKey is the standard key for referring to an API client's database ID.
	APIClientDatabaseIDKey = "api_client.id"
	// WebhookIDKey is the standard key for referring to a webhook's ID.
	WebhookIDKey = "webhook.id"
	// URLKey is the standard key for referring to a url.
	URLKey = "url"
	// RequestHeadersKey is the standard key for referring to an http.Request's Headers.
	RequestHeadersKey = "request.headers"
	// RequestMethodKey is the standard key for referring to an http.Request's Method.
	RequestMethodKey = "request.method"
	// RequestURIKey is the standard key for referring to an http.Request's URI.
	RequestURIKey = "request.uri"
	// RequestURIPathKey is the standard key for referring to an http.Request's URI.
	RequestURIPathKey = "request.uri.path"
	// RequestURIQueryKey is the standard key for referring to an http.Request's URI.
	RequestURIQueryKey = "request.uri.query"
	// ResponseStatusKey is the standard key for referring to an http.Request's URI.
	ResponseStatusKey = "response.status"
	// ResponseHeadersKey is the standard key for referring to an http.Response's Headers.
	ResponseHeadersKey = "response.headers"
	// ReasonKey is the standard key for referring to a reason.
	ReasonKey = "reason"
	// DatabaseQueryKey is the standard key for referring to a database query.
	DatabaseQueryKey = "database_query"
	// URLQueryKey is the standard key for referring to a url query.
	URLQueryKey = "url.query"
	// SearchQueryKey is the standard key for referring to a search query parameter value.
	SearchQueryKey = "search_query"
	// UserAgentOSKey is the standard key for referring to a search query parameter value.
	UserAgentOSKey = "os"
	// UserAgentBotKey is the standard key for referring to a search query parameter value.
	UserAgentBotKey = "is_bot"
	// UserAgentMobileKey is the standard key for referring to a search query parameter value.
	UserAgentMobileKey = "is_mobile"
	// QueryErrorKey is the standard key for referring to an error building a query.
	QueryErrorKey = "QUERY_ERROR"
	// ValidationErrorKey is the standard key for referring to a struct validation error.
	ValidationErrorKey = "validation_error"

	// ValidInstrumentIDKey is the standard key for referring to a valid instrument ID.
	ValidInstrumentIDKey = "valid_instrument.id"

	// ValidIngredientIDKey is the standard key for referring to a valid ingredient ID.
	ValidIngredientIDKey = "valid_ingredient.id"

	// ValidPreparationIDKey is the standard key for referring to a valid preparation ID.
	ValidPreparationIDKey = "valid_preparation.id"

	// ValidIngredientPreparationIDKey is the standard key for referring to a valid ingredient preparation ID.
	ValidIngredientPreparationIDKey = "valid_ingredient_preparation.id"

	// MealIDKey is the standard key for referring to a meal ID.
	MealIDKey = "meal.id"

	// RecipeIDKey is the standard key for referring to a recipe ID.
	RecipeIDKey = "recipe.id"

	// RecipeStepIDKey is the standard key for referring to a recipe step ID.
	RecipeStepIDKey = "recipe_step.id"

	// RecipeStepInstrumentIDKey is the standard key for referring to a recipe step instrument ID.
	RecipeStepInstrumentIDKey = "recipe_step_instrument.id"

	// RecipeStepIngredientIDKey is the standard key for referring to a recipe step ingredient ID.
	RecipeStepIngredientIDKey = "recipe_step_ingredient.id"

	// RecipeStepProductIDKey is the standard key for referring to a recipe step product ID.
	RecipeStepProductIDKey = "recipe_step_product.id"

	// MealPlanIDKey is the standard key for referring to a meal plan ID.
	MealPlanIDKey = "meal_plan.id"

	// MealPlanOptionIDKey is the standard key for referring to a meal plan option ID.
	MealPlanOptionIDKey = "meal_plan_option.id"

	// MealPlanOptionVoteIDKey is the standard key for referring to a meal plan option vote ID.
	MealPlanOptionVoteIDKey = "meal_plan_option_vote.id"
)
