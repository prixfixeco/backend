package main

import (
	"github.com/cristalhq/builq"
)

const mealPlanOptionsTableName = "meal_plan_options"

var mealPlanOptionsColumns = []string{
	idColumn,
	"meal_id",
	"notes",
	"chosen",
	"tiebroken",
	"assigned_cook",
	"assigned_dishwasher",
	"belongs_to_meal_plan_event",
	"meal_scale",
	createdAtColumn,
	lastUpdatedAtColumn,
	archivedAtColumn,
}

func buildMealPlanOptionsQueries() []*Query {
	return []*Query{
		{
			Annotation: QueryAnnotation{
				Name: "",
				Type: ExecType,
			},
			Content: buildRawQuery((&builq.Builder{}).Addf(``)),
		},
	}
}