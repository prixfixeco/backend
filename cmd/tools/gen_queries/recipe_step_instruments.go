package main

const recipeStepInstrumentsTableName = "recipe_step_instruments"

var recipeStepInstrumentsColumns = []string{
	"id",
	"instrument_id",
	"notes",
	"belongs_to_recipe_step",
	"preference_rank",
	"recipe_step_product_id",
	"name",
	"optional",
	"minimum_quantity",
	"maximum_quantity",
	"option_index",
	createdAtColumn,
	lastUpdatedAtColumn,
	archivedAtColumn,
}
