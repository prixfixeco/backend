SELECT EXISTS ( SELECT meal_ratings.id FROM meal_ratings WHERE meal_ratings.archived_at IS NULL AND meal_ratings.id = $1 );