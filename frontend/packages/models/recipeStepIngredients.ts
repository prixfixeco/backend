// Code generated by gen_typescript. DO NOT EDIT.

import { ValidIngredient } from './validIngredients';
import { ValidMeasurementUnit } from './validMeasurementUnits';

export interface IRecipeStepIngredient {
  createdAt: NonNullable<string>;
  recipeStepProductID?: string;
  archivedAt?: string;
  ingredient?: ValidIngredient;
  lastUpdatedAt?: string;
  maximumQuantity?: number;
  vesselIndex?: number;
  productPercentageToUse?: number;
  productOfRecipeID?: string;
  quantityNotes: NonNullable<string>;
  id: NonNullable<string>;
  belongsToRecipeStep: NonNullable<string>;
  ingredientNotes: NonNullable<string>;
  name: NonNullable<string>;
  measurementUnit: NonNullable<ValidMeasurementUnit>;
  minimumQuantity: NonNullable<number>;
  optionIndex: NonNullable<number>;
  optional: NonNullable<boolean>;
  toTaste: NonNullable<boolean>;
}

export class RecipeStepIngredient implements IRecipeStepIngredient {
  createdAt: NonNullable<string> = '1970-01-01T00:00:00Z';
  recipeStepProductID?: string;
  archivedAt?: string;
  ingredient?: ValidIngredient;
  lastUpdatedAt?: string;
  maximumQuantity?: number;
  vesselIndex?: number;
  productPercentageToUse?: number;
  productOfRecipeID?: string;
  quantityNotes: NonNullable<string> = '';
  id: NonNullable<string> = '';
  belongsToRecipeStep: NonNullable<string> = '';
  ingredientNotes: NonNullable<string> = '';
  name: NonNullable<string> = '';
  measurementUnit: NonNullable<ValidMeasurementUnit> = new ValidMeasurementUnit();
  minimumQuantity: NonNullable<number> = 0;
  optionIndex: NonNullable<number> = 0;
  optional: NonNullable<boolean> = false;
  toTaste: NonNullable<boolean> = false;

  constructor(input: Partial<RecipeStepIngredient> = {}) {
    this.createdAt = input.createdAt ?? '1970-01-01T00:00:00Z';
    this.recipeStepProductID = input.recipeStepProductID;
    this.archivedAt = input.archivedAt;
    this.ingredient = input.ingredient;
    this.lastUpdatedAt = input.lastUpdatedAt;
    this.maximumQuantity = input.maximumQuantity;
    this.vesselIndex = input.vesselIndex;
    this.productPercentageToUse = input.productPercentageToUse;
    this.productOfRecipeID = input.productOfRecipeID;
    this.quantityNotes = input.quantityNotes ?? '';
    this.id = input.id ?? '';
    this.belongsToRecipeStep = input.belongsToRecipeStep ?? '';
    this.ingredientNotes = input.ingredientNotes ?? '';
    this.name = input.name ?? '';
    this.measurementUnit = input.measurementUnit ?? new ValidMeasurementUnit();
    this.minimumQuantity = input.minimumQuantity ?? 0;
    this.optionIndex = input.optionIndex ?? 0;
    this.optional = input.optional ?? false;
    this.toTaste = input.toTaste ?? false;
  }
}

export interface IRecipeStepIngredientCreationRequestInput {
  ingredientID?: string;
  productOfRecipeStepIndex?: number;
  productOfRecipeStepProductIndex?: number;
  maximumQuantity?: number;
  vesselIndex?: number;
  productPercentageToUse?: number;
  productOfRecipeID?: string;
  ingredientNotes: NonNullable<string>;
  measurementUnitID: NonNullable<string>;
  name: NonNullable<string>;
  quantityNotes: NonNullable<string>;
  minimumQuantity: NonNullable<number>;
  optionIndex: NonNullable<number>;
  optional: NonNullable<boolean>;
  toTaste: NonNullable<boolean>;
}

export class RecipeStepIngredientCreationRequestInput implements IRecipeStepIngredientCreationRequestInput {
  ingredientID?: string;
  productOfRecipeStepIndex?: number;
  productOfRecipeStepProductIndex?: number;
  maximumQuantity?: number;
  vesselIndex?: number;
  productPercentageToUse?: number;
  productOfRecipeID?: string;
  ingredientNotes: NonNullable<string> = '';
  measurementUnitID: NonNullable<string> = '';
  name: NonNullable<string> = '';
  quantityNotes: NonNullable<string> = '';
  minimumQuantity: NonNullable<number> = 0;
  optionIndex: NonNullable<number> = 0;
  optional: NonNullable<boolean> = false;
  toTaste: NonNullable<boolean> = false;

  constructor(input: Partial<RecipeStepIngredientCreationRequestInput> = {}) {
    this.ingredientID = input.ingredientID;
    this.productOfRecipeStepIndex = input.productOfRecipeStepIndex;
    this.productOfRecipeStepProductIndex = input.productOfRecipeStepProductIndex;
    this.maximumQuantity = input.maximumQuantity;
    this.vesselIndex = input.vesselIndex;
    this.productPercentageToUse = input.productPercentageToUse;
    this.productOfRecipeID = input.productOfRecipeID;
    this.ingredientNotes = input.ingredientNotes ?? '';
    this.measurementUnitID = input.measurementUnitID ?? '';
    this.name = input.name ?? '';
    this.quantityNotes = input.quantityNotes ?? '';
    this.minimumQuantity = input.minimumQuantity ?? 0;
    this.optionIndex = input.optionIndex ?? 0;
    this.optional = input.optional ?? false;
    this.toTaste = input.toTaste ?? false;
  }
}

export interface IRecipeStepIngredientUpdateRequestInput {
  ingredientID?: string;
  recipeStepProductID?: string;
  name?: string;
  optional?: boolean;
  measurementUnitID?: string;
  quantityNotes?: string;
  ingredientNotes?: string;
  belongsToRecipeStep?: string;
  minimumQuantity?: number;
  maximumQuantity?: number;
  optionIndex?: number;
  vesselIndex?: number;
  toTaste?: boolean;
  productPercentageToUse?: number;
  productOfRecipeID?: string;
}

export class RecipeStepIngredientUpdateRequestInput implements IRecipeStepIngredientUpdateRequestInput {
  ingredientID?: string;
  recipeStepProductID?: string;
  name?: string;
  optional?: boolean = false;
  measurementUnitID?: string;
  quantityNotes?: string;
  ingredientNotes?: string;
  belongsToRecipeStep?: string;
  minimumQuantity?: number;
  maximumQuantity?: number;
  optionIndex?: number;
  vesselIndex?: number;
  toTaste?: boolean = false;
  productPercentageToUse?: number;
  productOfRecipeID?: string;

  constructor(input: Partial<RecipeStepIngredientUpdateRequestInput> = {}) {
    this.ingredientID = input.ingredientID;
    this.recipeStepProductID = input.recipeStepProductID;
    this.name = input.name;
    this.optional = input.optional ?? false;
    this.measurementUnitID = input.measurementUnitID;
    this.quantityNotes = input.quantityNotes;
    this.ingredientNotes = input.ingredientNotes;
    this.belongsToRecipeStep = input.belongsToRecipeStep;
    this.minimumQuantity = input.minimumQuantity;
    this.maximumQuantity = input.maximumQuantity;
    this.optionIndex = input.optionIndex;
    this.vesselIndex = input.vesselIndex;
    this.toTaste = input.toTaste ?? false;
    this.productPercentageToUse = input.productPercentageToUse;
    this.productOfRecipeID = input.productOfRecipeID;
  }
}