-- name: CreateValidPreparation :exec

INSERT INTO valid_preparations (id,"name",description,icon_path,yields_nothing,restrict_to_ingredients,minimum_ingredient_count,maximum_ingredient_count,minimum_instrument_count,maximum_instrument_count,temperature_required,time_estimate_required,condition_expression_required,consumes_vessel,only_for_vessels,minimum_vessel_count,maximum_vessel_count,past_tense,slug) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19);
