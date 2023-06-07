# how to use these resources: https://github.com/philippe-vandermoere/terraform-provider-algolia/blob/11d9e162be54c66c92376ae5647f7f3bd675755a/examples/main.tf

locals {
  default_algolia_ranking_criteria = ["typo", "words", "filters", "proximity", "attribute", "exact", "custom"]
}

resource "algolia_index" "recipes_index" {
  name = "recipes"

  ranking        = local.default_algolia_ranking_criteria
  custom_ranking = []

  searchable_attributes = [
    "name",
    "description",
  ]
}

resource "algolia_index" "meals_index" {
  name = "meals"

  ranking        = local.default_algolia_ranking_criteria
  custom_ranking = []

  searchable_attributes = [
    "name",
    "description",
  ]
}

resource "algolia_index" "valid_ingredients_index" {
  name = "valid_ingredients"

  ranking        = local.default_algolia_ranking_criteria
  custom_ranking = []

  searchable_attributes = [
    "name",
    "pluralName",
    "description",
  ]
}

resource "algolia_index" "valid_instruments_index" {
  name = "valid_instruments"

  ranking        = local.default_algolia_ranking_criteria
  custom_ranking = []

  searchable_attributes = [
    "name",
    "pluralName",
    "description",
  ]
}

resource "algolia_index" "valid_measurement_units_index" {
  name = "valid_measurement_units"

  ranking        = local.default_algolia_ranking_criteria
  custom_ranking = []

  searchable_attributes = [
    "name",
    "pluralName",
    "description",
  ]
}

resource "algolia_index" "valid_preparations_index" {
  name = "valid_preparations"

  ranking        = local.default_algolia_ranking_criteria
  custom_ranking = []

  searchable_attributes = [
    "name",
    "pastTense",
    "description",
  ]
}

resource "algolia_index" "valid_ingredient_states_index" {
  name = "valid_ingredient_states"

  ranking        = local.default_algolia_ranking_criteria
  custom_ranking = []

  searchable_attributes = [
    "name",
    "pastTense",
    "description",
  ]
}

resource "algolia_index" "valid_ingredient_measurement_units_index" {
  name = "valid_ingredient_measurement_units"

  ranking        = local.default_algolia_ranking_criteria
  custom_ranking = []

  searchable_attributes = [
    "measurementUnit.name",
    "measurementUnit.pluralName",
    "measurementUnit.description",
    "ingredient.name",
    "ingredient.pluralName",
    "ingredient.description",
  ]
}

resource "algolia_index" "valid_measurement_unit_conversions_index" {
  name = "valid_measurement_unit_conversions"

  ranking        = local.default_algolia_ranking_criteria
  custom_ranking = []

  searchable_attributes = [
    "notes",
    "fromMeasurementUnit.name",
    "toMeasurementUnit.name",
  ]
}

resource "algolia_index" "valid_preparation_instruments_index" {
  name = "valid_preparation_instruments"

  ranking        = local.default_algolia_ranking_criteria
  custom_ranking = []

  searchable_attributes = [
    "notes",
    "instrument.name",
    "preparation.name",
  ]
}

resource "algolia_index" "valid_ingredient_preparations_index" {
  name = "valid_ingredient_preparations"

  ranking        = local.default_algolia_ranking_criteria
  custom_ranking = []

  searchable_attributes = [
    "notes",
    "ingredient.name",
    "preparation.name",
  ]
}
