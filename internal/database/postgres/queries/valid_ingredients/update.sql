UPDATE valid_ingredients SET
	name = $1,
	description = $2,
	warning = $3,
	contains_egg = $4,
	contains_dairy = $5,
	contains_peanut = $6,
	contains_tree_nut = $7,
	contains_soy = $8,
	contains_wheat = $9,
	contains_shellfish = $10,
	contains_sesame = $11,
	contains_fish = $12,
	contains_gluten = $13,
	animal_flesh = $14,
	volumetric = $15,
	is_liquid = $16,
	icon_path = $17,
	animal_derived = $18,
	plural_name = $19,
	restrict_to_preparations = $20,
	minimum_ideal_storage_temperature_in_celsius = $21,
	maximum_ideal_storage_temperature_in_celsius = $22,
	storage_instructions = $23,
    slug = $24,
    contains_alcohol = $25,
    shopping_suggestions = $26,
	last_updated_at = NOW()
WHERE archived_at IS NULL AND id = $27;
