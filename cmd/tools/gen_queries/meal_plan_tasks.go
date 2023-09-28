package main

import (
	"github.com/cristalhq/builq"
)

const (
	mealPlanTasksTableName = "meal_plan_tasks"
)

var mealPlanTasksColumns = []string{
	idColumn,
	"belongs_to_meal_plan_option",
	"belongs_to_recipe_prep_task",
	"creation_explanation",
	"status_explanation",
	"status",
	"assigned_to_user",
	"completed_at",
	createdAtColumn,
	lastUpdatedAtColumn,
}

func buildMealPlanTasksQueries() []*Query {
	return []*Query{
		{
			Annotation: QueryAnnotation{
				Name: "",
				Type: "",
			},
			Content: buildRawQuery((&builq.Builder{}).Addf(``)),
		},
	}
}
