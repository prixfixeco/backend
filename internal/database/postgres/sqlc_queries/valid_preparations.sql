-- name: ArchiveValidPreparation :exec

UPDATE valid_preparations SET archived_at = NOW() WHERE archived_at IS NULL AND id = $1;

-- name: CreateValidPreparation :exec

INSERT INTO valid_preparations (id,"name",description,icon_path,yields_nothing,restrict_to_ingredients,minimum_ingredient_count,maximum_ingredient_count,minimum_instrument_count,maximum_instrument_count,temperature_required,time_estimate_required,condition_expression_required,consumes_vessel,only_for_vessels,minimum_vessel_count,maximum_vessel_count,past_tense,slug) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19);

-- name: CheckValidPreparationExistence :one

SELECT EXISTS ( SELECT valid_preparations.id FROM valid_preparations WHERE valid_preparations.archived_at IS NULL AND valid_preparations.id = $1 );

-- name: GetValidPreparationByID :one

SELECT valid_preparations.id,
       valid_preparations.name,
       valid_preparations.description,
       valid_preparations.icon_path,
       valid_preparations.yields_nothing,
       valid_preparations.restrict_to_ingredients,
       valid_preparations.minimum_ingredient_count,
       valid_preparations.maximum_ingredient_count,
       valid_preparations.minimum_instrument_count,
       valid_preparations.maximum_instrument_count,
       valid_preparations.temperature_required,
       valid_preparations.time_estimate_required,
       valid_preparations.condition_expression_required,
       valid_preparations.consumes_vessel,
       valid_preparations.only_for_vessels,
       valid_preparations.minimum_vessel_count,
       valid_preparations.maximum_vessel_count,
       valid_preparations.slug,
       valid_preparations.past_tense,
       valid_preparations.created_at,
       valid_preparations.last_updated_at,
       valid_preparations.archived_at
  FROM valid_preparations
 WHERE valid_preparations.archived_at IS NULL;

-- name: GetValidPreparations :many

SELECT
	valid_preparations.id,
	valid_preparations.name,
	valid_preparations.description,
	valid_preparations.icon_path,
	valid_preparations.yields_nothing,
	valid_preparations.restrict_to_ingredients,
	valid_preparations.minimum_ingredient_count,
	valid_preparations.maximum_ingredient_count,
	valid_preparations.minimum_instrument_count,
	valid_preparations.maximum_instrument_count,
	valid_preparations.temperature_required,
	valid_preparations.time_estimate_required,
    valid_preparations.condition_expression_required,
    valid_preparations.consumes_vessel,
    valid_preparations.only_for_vessels,
    valid_preparations.minimum_vessel_count,
    valid_preparations.maximum_vessel_count,
	valid_preparations.slug,
	valid_preparations.past_tense,
	valid_preparations.created_at,
	valid_preparations.last_updated_at,
	valid_preparations.archived_at,
    (
        SELECT
            COUNT(valid_preparations.id)
        FROM
            valid_preparations
        WHERE
            valid_preparations.archived_at IS NULL
          AND valid_preparations.created_at > COALESCE(sqlc.narg(created_after), (SELECT NOW() - interval '999 years'))
          AND valid_preparations.created_at < COALESCE(sqlc.narg(created_before), (SELECT NOW() + interval '999 years'))
          AND (
                valid_preparations.last_updated_at IS NULL
                OR valid_preparations.last_updated_at > COALESCE(sqlc.narg(updated_after), (SELECT NOW() - interval '999 years'))
            )
          AND (
                valid_preparations.last_updated_at IS NULL
                OR valid_preparations.last_updated_at < COALESCE(sqlc.narg(updated_before), (SELECT NOW() + interval '999 years'))
            )
    ) as filtered_count,
    (
        SELECT
            COUNT(valid_preparations.id)
        FROM
            valid_preparations
        WHERE
            valid_preparations.archived_at IS NULL
    ) as total_count
FROM valid_preparations
WHERE
    valid_preparations.archived_at IS NULL
  AND valid_preparations.created_at > (COALESCE(sqlc.narg(created_after), (SELECT NOW() - interval '999 years')))
  AND valid_preparations.created_at < (COALESCE(sqlc.narg(created_before), (SELECT NOW() + interval '999 years')))
  AND (
        valid_preparations.last_updated_at IS NULL
        OR valid_preparations.last_updated_at > COALESCE(sqlc.narg(updated_after), (SELECT NOW() - interval '999 years'))
    )
  AND (
        valid_preparations.last_updated_at IS NULL
        OR valid_preparations.last_updated_at < COALESCE(sqlc.narg(updated_before), (SELECT NOW() + interval '999 years'))
    )
GROUP BY
    valid_preparations.id
ORDER BY
    valid_preparations.id
OFFSET
    sqlc.narg(query_offset)
    LIMIT
    sqlc.narg(query_limit);

-- name: GetValidPreparationsNeedingIndexing :many

SELECT valid_preparations.id
  FROM valid_preparations
 WHERE (valid_preparations.archived_at IS NULL)
       AND (
			(valid_preparations.last_indexed_at IS NULL)
			OR valid_preparations.last_indexed_at
				< now() - '24 hours'::INTERVAL
		);

-- name: GetValidPreparation :one

SELECT
	valid_preparations.id,
	valid_preparations.name,
	valid_preparations.description,
	valid_preparations.icon_path,
	valid_preparations.yields_nothing,
	valid_preparations.restrict_to_ingredients,
	valid_preparations.minimum_ingredient_count,
	valid_preparations.maximum_ingredient_count,
	valid_preparations.minimum_instrument_count,
	valid_preparations.maximum_instrument_count,
	valid_preparations.temperature_required,
	valid_preparations.time_estimate_required,
    valid_preparations.condition_expression_required,
    valid_preparations.consumes_vessel,
    valid_preparations.only_for_vessels,
    valid_preparations.minimum_vessel_count,
    valid_preparations.maximum_vessel_count,
	valid_preparations.slug,
	valid_preparations.past_tense,
	valid_preparations.created_at,
	valid_preparations.last_updated_at,
	valid_preparations.archived_at
FROM valid_preparations
WHERE valid_preparations.archived_at IS NULL
	AND valid_preparations.id = $1;

-- name: GetRandomValidPreparation :one

SELECT
	valid_preparations.id,
	valid_preparations.name,
	valid_preparations.description,
	valid_preparations.icon_path,
	valid_preparations.yields_nothing,
	valid_preparations.restrict_to_ingredients,
	valid_preparations.minimum_ingredient_count,
	valid_preparations.maximum_ingredient_count,
	valid_preparations.minimum_instrument_count,
	valid_preparations.maximum_instrument_count,
	valid_preparations.temperature_required,
	valid_preparations.time_estimate_required,
	valid_preparations.condition_expression_required,
    valid_preparations.consumes_vessel,
    valid_preparations.only_for_vessels,
    valid_preparations.minimum_vessel_count,
    valid_preparations.maximum_vessel_count,
	valid_preparations.slug,
	valid_preparations.past_tense,
	valid_preparations.created_at,
	valid_preparations.last_updated_at,
	valid_preparations.archived_at
FROM valid_preparations
WHERE valid_preparations.archived_at IS NULL
	ORDER BY random() LIMIT 1;

-- name: GetValidPreparationsWithIDs :many

SELECT
	valid_preparations.id,
	valid_preparations.name,
	valid_preparations.description,
	valid_preparations.icon_path,
	valid_preparations.yields_nothing,
	valid_preparations.restrict_to_ingredients,
	valid_preparations.minimum_ingredient_count,
	valid_preparations.maximum_ingredient_count,
	valid_preparations.minimum_instrument_count,
	valid_preparations.maximum_instrument_count,
	valid_preparations.temperature_required,
	valid_preparations.time_estimate_required,
    valid_preparations.condition_expression_required,
    valid_preparations.consumes_vessel,
    valid_preparations.only_for_vessels,
    valid_preparations.minimum_vessel_count,
    valid_preparations.maximum_vessel_count,
	valid_preparations.slug,
	valid_preparations.past_tense,
	valid_preparations.created_at,
	valid_preparations.last_updated_at,
	valid_preparations.archived_at
FROM valid_preparations
WHERE valid_preparations.archived_at IS NULL
	AND valid_preparations.id = ANY($1::text[]);

-- name: SearchForValidPreparations :many

SELECT
	valid_preparations.id,
	valid_preparations.name,
	valid_preparations.description,
	valid_preparations.icon_path,
	valid_preparations.yields_nothing,
	valid_preparations.restrict_to_ingredients,
	valid_preparations.minimum_ingredient_count,
	valid_preparations.maximum_ingredient_count,
	valid_preparations.minimum_instrument_count,
	valid_preparations.maximum_instrument_count,
	valid_preparations.temperature_required,
	valid_preparations.time_estimate_required,
	valid_preparations.condition_expression_required,
    valid_preparations.consumes_vessel,
    valid_preparations.only_for_vessels,
    valid_preparations.minimum_vessel_count,
    valid_preparations.maximum_vessel_count,
	valid_preparations.slug,
	valid_preparations.past_tense,
	valid_preparations.created_at,
	valid_preparations.last_updated_at,
	valid_preparations.archived_at
FROM valid_preparations
WHERE valid_preparations.archived_at IS NULL
	AND valid_preparations.name ILIKE '%' || sqlc.arg(query)::text || '%'
LIMIT 50;

-- name: UpdateValidPreparation :exec

UPDATE valid_preparations
SET
	name = sqlc.arg(name),
	description = sqlc.arg(description),
	icon_path = sqlc.arg(icon_path),
	yields_nothing = sqlc.arg(yields_nothing),
	restrict_to_ingredients = sqlc.arg(restrict_to_ingredients),
	minimum_ingredient_count = sqlc.arg(minimum_ingredient_count),
	maximum_ingredient_count = sqlc.arg(maximum_ingredient_count),
	minimum_instrument_count = sqlc.arg(minimum_instrument_count),
	maximum_instrument_count = sqlc.arg(maximum_instrument_count),
	temperature_required = sqlc.arg(temperature_required),
	time_estimate_required = sqlc.arg(time_estimate_required),
	condition_expression_required = sqlc.arg(condition_expression_required),
    consumes_vessel = sqlc.arg(consumes_vessel),
    only_for_vessels = sqlc.arg(only_for_vessels),
    minimum_vessel_count = sqlc.arg(minimum_vessel_count),
    maximum_vessel_count = sqlc.arg(maximum_vessel_count),
	slug = sqlc.arg(slug),
	past_tense = sqlc.arg(past_tense),
	last_updated_at = NOW()
WHERE archived_at IS NULL
	AND id = sqlc.arg(id);

-- name: UpdateValidPreparationLastIndexedAt :exec

UPDATE valid_preparations SET last_indexed_at = NOW() WHERE id = $1 AND archived_at IS NULL;
