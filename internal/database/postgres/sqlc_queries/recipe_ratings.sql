-- name: ArchiveRecipeRating :execrows

UPDATE recipe_ratings SET archived_at = NOW() WHERE archived_at IS NULL AND id = $1;

-- name: CreateRecipeRating :exec

INSERT INTO recipe_ratings (id,recipe_id,taste,difficulty,cleanup,instructions,overall,notes,by_user) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9);

-- name: CheckRecipeRatingExistence :one

SELECT EXISTS ( SELECT recipe_ratings.id FROM recipe_ratings WHERE recipe_ratings.archived_at IS NULL AND recipe_ratings.id = $1 );

-- name: GetRecipeRatings :many

SELECT
	recipe_ratings.id,
    recipe_ratings.recipe_id,
    recipe_ratings.taste,
    recipe_ratings.difficulty,
    recipe_ratings.cleanup,
    recipe_ratings.instructions,
    recipe_ratings.overall,
    recipe_ratings.notes,
    recipe_ratings.by_user,
    recipe_ratings.created_at,
    recipe_ratings.last_updated_at,
    recipe_ratings.archived_at,
	(
	 SELECT
		COUNT(recipe_ratings.id)
	 FROM
		recipe_ratings
	 WHERE
		recipe_ratings.archived_at IS NULL
	 AND recipe_ratings.created_at > COALESCE(sqlc.narg(created_after), (SELECT NOW() - interval '999 years'))
	 AND recipe_ratings.created_at < COALESCE(sqlc.narg(created_before), (SELECT NOW() + interval '999 years'))
	 AND (recipe_ratings.last_updated_at IS NULL OR recipe_ratings.last_updated_at > COALESCE(sqlc.narg(updated_after), (SELECT NOW() - interval '999 years')))
	 AND (recipe_ratings.last_updated_at IS NULL OR recipe_ratings.last_updated_at < COALESCE(sqlc.narg(updated_before), (SELECT NOW() + interval '999 years')))
	) as filtered_count,
	(
	 SELECT
		COUNT(recipe_ratings.id)
	 FROM
		recipe_ratings
	 WHERE
		recipe_ratings.archived_at IS NULL
	) as total_count
FROM
	recipe_ratings
WHERE
	recipe_ratings.archived_at IS NULL
	AND recipe_ratings.created_at > COALESCE(sqlc.narg(created_after), (SELECT NOW() - interval '999 years'))
	AND recipe_ratings.created_at < COALESCE(sqlc.narg(created_before), (SELECT NOW() + interval '999 years'))
	AND (recipe_ratings.last_updated_at IS NULL OR recipe_ratings.last_updated_at > COALESCE(sqlc.narg(updated_after), (SELECT NOW() - interval '999 years')))
	AND (recipe_ratings.last_updated_at IS NULL OR recipe_ratings.last_updated_at < COALESCE(sqlc.narg(updated_before), (SELECT NOW() + interval '999 years')))
GROUP BY
	recipe_ratings.id
ORDER BY
	recipe_ratings.id
    OFFSET sqlc.narg(query_offset)
	LIMIT sqlc.narg(query_limit);

-- name: GetRecipeRating :one

SELECT
	recipe_ratings.id,
    recipe_ratings.recipe_id,
    recipe_ratings.taste,
    recipe_ratings.difficulty,
    recipe_ratings.cleanup,
    recipe_ratings.instructions,
    recipe_ratings.overall,
    recipe_ratings.notes,
    recipe_ratings.by_user,
    recipe_ratings.created_at,
    recipe_ratings.last_updated_at,
    recipe_ratings.archived_at
FROM recipe_ratings
WHERE recipe_ratings.archived_at IS NULL
	AND recipe_ratings.id = $1;

-- name: UpdateRecipeRating :execrows

UPDATE recipe_ratings
SET
	recipe_id = sqlc.arg(recipe_id),
    taste = sqlc.arg(taste),
    difficulty = sqlc.arg(difficulty),
    cleanup = sqlc.arg(cleanup),
    instructions = sqlc.arg(instructions),
    overall = sqlc.arg(overall),
    notes = sqlc.arg(notes),
	last_updated_at = NOW()
WHERE archived_at IS NULL
	AND id = sqlc.arg(id);
