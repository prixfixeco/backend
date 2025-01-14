// GENERATED CODE, DO NOT EDIT MANUALLY

import { RecipeStepCompletionConditionIngredientForExistingRecipeCreationRequestInput } from './RecipeStepCompletionConditionIngredientForExistingRecipeCreationRequestInput.gen';

export interface IRecipeStepCompletionConditionForExistingRecipeCreationRequestInput {
  belongsToRecipeStep: string;
  ingredientStateID: string;
  ingredients: RecipeStepCompletionConditionIngredientForExistingRecipeCreationRequestInput[];
  notes: string;
  optional: boolean;
}

export class RecipeStepCompletionConditionForExistingRecipeCreationRequestInput
  implements IRecipeStepCompletionConditionForExistingRecipeCreationRequestInput
{
  belongsToRecipeStep: string;
  ingredientStateID: string;
  ingredients: RecipeStepCompletionConditionIngredientForExistingRecipeCreationRequestInput[];
  notes: string;
  optional: boolean;
  constructor(input: Partial<RecipeStepCompletionConditionForExistingRecipeCreationRequestInput> = {}) {
    this.belongsToRecipeStep = input.belongsToRecipeStep || '';
    this.ingredientStateID = input.ingredientStateID || '';
    this.ingredients = input.ingredients || [];
    this.notes = input.notes || '';
    this.optional = input.optional || false;
  }
}
