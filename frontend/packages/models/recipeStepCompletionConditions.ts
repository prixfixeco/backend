// Code generated by gen_typescript. DO NOT EDIT.

import { ValidIngredientState } from './validIngredientStates';

export interface IRecipeStepCompletionCondition {
  createdAt: NonNullable<string>;
  archivedAt?: string;
  lastUpdatedAt?: string;
  ingredientState: NonNullable<ValidIngredientState>;
  id: NonNullable<string>;
  belongsToRecipeStep: NonNullable<string>;
  notes: NonNullable<string>;
  ingredients: NonNullable<Array<RecipeStepCompletionConditionIngredient>>;
  optional: NonNullable<boolean>;
}

export class RecipeStepCompletionCondition implements IRecipeStepCompletionCondition {
  createdAt: NonNullable<string> = '1970-01-01T00:00:00Z';
  archivedAt?: string;
  lastUpdatedAt?: string;
  ingredientState: NonNullable<ValidIngredientState> = new ValidIngredientState();
  id: NonNullable<string> = '';
  belongsToRecipeStep: NonNullable<string> = '';
  notes: NonNullable<string> = '';
  ingredients: NonNullable<Array<RecipeStepCompletionConditionIngredient>> = [];
  optional: NonNullable<boolean> = false;

  constructor(input: Partial<RecipeStepCompletionCondition> = {}) {
    this.createdAt = input.createdAt ?? '1970-01-01T00:00:00Z';
    this.archivedAt = input.archivedAt;
    this.lastUpdatedAt = input.lastUpdatedAt;
    this.ingredientState = input.ingredientState ?? new ValidIngredientState();
    this.id = input.id ?? '';
    this.belongsToRecipeStep = input.belongsToRecipeStep ?? '';
    this.notes = input.notes ?? '';
    this.ingredients = input.ingredients ?? [];
    this.optional = input.optional ?? false;
  }
}

export interface IRecipeStepCompletionConditionIngredient {
  createdAt: NonNullable<string>;
  archivedAt?: string;
  lastUpdatedAt?: string;
  id: NonNullable<string>;
  belongsToRecipeStepCompletionCondition: NonNullable<string>;
  recipeStepIngredient: NonNullable<string>;
}

export class RecipeStepCompletionConditionIngredient implements IRecipeStepCompletionConditionIngredient {
  createdAt: NonNullable<string> = '1970-01-01T00:00:00Z';
  archivedAt?: string;
  lastUpdatedAt?: string;
  id: NonNullable<string> = '';
  belongsToRecipeStepCompletionCondition: NonNullable<string> = '';
  recipeStepIngredient: NonNullable<string> = '';

  constructor(input: Partial<RecipeStepCompletionConditionIngredient> = {}) {
    this.createdAt = input.createdAt ?? '1970-01-01T00:00:00Z';
    this.archivedAt = input.archivedAt;
    this.lastUpdatedAt = input.lastUpdatedAt;
    this.id = input.id ?? '';
    this.belongsToRecipeStepCompletionCondition = input.belongsToRecipeStepCompletionCondition ?? '';
    this.recipeStepIngredient = input.recipeStepIngredient ?? '';
  }
}

export interface IRecipeStepCompletionConditionCreationRequestInput {
  ingredientState: NonNullable<string>;
  belongsToRecipeStep: NonNullable<string>;
  notes: NonNullable<string>;
  ingredients: NonNullable<Array<number>>;
  optional: NonNullable<boolean>;
}

export class RecipeStepCompletionConditionCreationRequestInput
  implements IRecipeStepCompletionConditionCreationRequestInput
{
  ingredientState: NonNullable<string> = '';
  belongsToRecipeStep: NonNullable<string> = '';
  notes: NonNullable<string> = '';
  ingredients: NonNullable<Array<number>> = [];
  optional: NonNullable<boolean> = false;

  constructor(input: Partial<RecipeStepCompletionConditionCreationRequestInput> = {}) {
    this.ingredientState = input.ingredientState ?? '';
    this.belongsToRecipeStep = input.belongsToRecipeStep ?? '';
    this.notes = input.notes ?? '';
    this.ingredients = input.ingredients ?? [];
    this.optional = input.optional ?? false;
  }
}

export interface IRecipeStepCompletionConditionIngredientCreationRequestInput {
  recipeStepIngredient: NonNullable<string>;
}

export class RecipeStepCompletionConditionIngredientCreationRequestInput
  implements IRecipeStepCompletionConditionIngredientCreationRequestInput
{
  recipeStepIngredient: NonNullable<string> = '';

  constructor(input: Partial<RecipeStepCompletionConditionIngredientCreationRequestInput> = {}) {
    this.recipeStepIngredient = input.recipeStepIngredient ?? '';
  }
}

export interface IRecipeStepCompletionConditionUpdateRequestInput {
  ingredientState?: string;
  belongsToRecipeStep?: string;
  notes?: string;
  optional?: boolean;
}

export class RecipeStepCompletionConditionUpdateRequestInput
  implements IRecipeStepCompletionConditionUpdateRequestInput
{
  ingredientState?: string;
  belongsToRecipeStep?: string;
  notes?: string;
  optional?: boolean = false;

  constructor(input: Partial<RecipeStepCompletionConditionUpdateRequestInput> = {}) {
    this.ingredientState = input.ingredientState;
    this.belongsToRecipeStep = input.belongsToRecipeStep;
    this.notes = input.notes;
    this.optional = input.optional ?? false;
  }
}
