UPDATE recipe_steps SET
    index = $1,
    preparation_id = $2,
    minimum_estimated_time_in_seconds = $3,
    maximum_estimated_time_in_seconds = $4,
    minimum_temperature_in_celsius = $5,
    maximum_temperature_in_celsius = $6,
    notes = $7,
    explicit_instructions = $8,
    optional = $9,
    last_updated_at = NOW()
WHERE archived_at IS NULL
  AND belongs_to_recipe = $10
  AND id = $11;