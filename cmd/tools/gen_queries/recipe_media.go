package main

const recipeMediaTableName = "recipe_media"

var recipeMediaColumns = []string{
	idColumn,
	"belongs_to_recipe",
	"belongs_to_recipe_step",
	"mime_type",
	"internal_path",
	"external_path",
	"index",
	createdAtColumn,
	lastUpdatedAtColumn,
	archivedAtColumn,
}

func buildRecipeMediaQueries() []*Query {
	return []*Query{
		//
	}
}
