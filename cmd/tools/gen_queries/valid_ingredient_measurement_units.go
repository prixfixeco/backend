package main

const validIngredientMeasurementUnitsTableName = "valid_ingredient_measurement_units"

var validIngredientMeasurementUnitsColumns = []string{
	"id",
	"notes",
	"valid_ingredient_id",
	"valid_measurement_unit_id",
	"minimum_allowable_quantity",
	"maximum_allowable_quantity",
	createdAtColumn,
	lastUpdatedAtColumn,
	archivedAtColumn,
}
