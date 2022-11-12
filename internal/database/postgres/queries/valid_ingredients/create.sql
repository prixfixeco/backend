INSERT INTO valid_ingredients
(
	id,
	name,
	description,
	warning,
	contains_egg,
	contains_dairy,
	contains_peanut,
	contains_tree_nut,
	contains_soy,
	contains_wheat,
	contains_shellfish,
	contains_sesame,
	contains_fish,
	contains_gluten,
	animal_flesh,
	volumetric,
	is_liquid,
	icon_path,
	animal_derived,
	plural_name,
	restrict_to_preparations,
	minimum_ideal_storage_temperature_in_celsius,
	maximum_ideal_storage_temperature_in_celsius,
	storage_instructions,
    slug,
    contains_alcohol,
    shopping_suggestions
) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22,$23,$24,$25,$26,$27);
