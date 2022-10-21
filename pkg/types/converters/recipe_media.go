package converters

import (
	"github.com/prixfixeco/api_server/pkg/types"
)

// ConvertRecipeMediaToRecipeMediaUpdateRequestInput creates a RecipeMediaUpdateRequestInput from a RecipeMedia.
func ConvertRecipeMediaToRecipeMediaUpdateRequestInput(input *types.RecipeMedia) *types.RecipeMediaUpdateRequestInput {
	x := &types.RecipeMediaUpdateRequestInput{
		BelongsToRecipe:     input.BelongsToRecipe,
		BelongsToRecipeStep: input.BelongsToRecipeStep,
		MimeType:            &input.MimeType,
		InternalPath:        &input.InternalPath,
		ExternalPath:        &input.ExternalPath,
	}

	return x
}

// ConvertRecipeMediaCreationRequestInputToRecipeMediaDatabaseCreationInput creates a RecipeMediaDatabaseCreationInput from a RecipeMediaCreationRequestInput.
func ConvertRecipeMediaCreationRequestInputToRecipeMediaDatabaseCreationInput(input *types.RecipeMediaCreationRequestInput) *types.RecipeMediaDatabaseCreationInput {
	x := &types.RecipeMediaDatabaseCreationInput{
		BelongsToRecipe:     input.BelongsToRecipe,
		BelongsToRecipeStep: input.BelongsToRecipeStep,
		MimeType:            input.MimeType,
		InternalPath:        input.InternalPath,
		ExternalPath:        input.ExternalPath,
	}

	return x
}

// ConvertRecipeMediaToRecipeMediaCreationRequestInput builds a RecipeMediaCreationRequestInput from a RecipeMedia.
func ConvertRecipeMediaToRecipeMediaCreationRequestInput(recipeMedia *types.RecipeMedia) *types.RecipeMediaCreationRequestInput {
	return &types.RecipeMediaCreationRequestInput{
		BelongsToRecipe:     recipeMedia.BelongsToRecipe,
		BelongsToRecipeStep: recipeMedia.BelongsToRecipeStep,
		MimeType:            recipeMedia.MimeType,
		InternalPath:        recipeMedia.InternalPath,
		ExternalPath:        recipeMedia.ExternalPath,
	}
}

// ConvertRecipeMediaToRecipeMediaDatabaseCreationInput builds a RecipeMediaDatabaseCreationInput from a RecipeMedia.
func ConvertRecipeMediaToRecipeMediaDatabaseCreationInput(recipeMedia *types.RecipeMedia) *types.RecipeMediaDatabaseCreationInput {
	return &types.RecipeMediaDatabaseCreationInput{
		ID:                  recipeMedia.ID,
		BelongsToRecipe:     recipeMedia.BelongsToRecipe,
		BelongsToRecipeStep: recipeMedia.BelongsToRecipeStep,
		MimeType:            recipeMedia.MimeType,
		InternalPath:        recipeMedia.InternalPath,
		ExternalPath:        recipeMedia.ExternalPath,
	}
}