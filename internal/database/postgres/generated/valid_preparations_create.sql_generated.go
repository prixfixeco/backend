// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: valid_preparations_create.sql

package generated

import (
	"context"
)

const CreateValidPreparation = `-- name: CreateValidPreparation :exec
INSERT INTO valid_preparations (id,name,description,icon_path,yields_nothing,restrict_to_ingredients,zero_ingredients_allowable,past_tense) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
`

type CreateValidPreparationParams struct {
	ID                       string `db:"id"`
	Name                     string `db:"name"`
	Description              string `db:"description"`
	IconPath                 string `db:"icon_path"`
	PastTense                string `db:"past_tense"`
	YieldsNothing            bool   `db:"yields_nothing"`
	RestrictToIngredients    bool   `db:"restrict_to_ingredients"`
	ZeroIngredientsAllowable bool   `db:"zero_ingredients_allowable"`
}

func (q *Queries) CreateValidPreparation(ctx context.Context, db DBTX, arg *CreateValidPreparationParams) error {
	_, err := db.ExecContext(ctx, CreateValidPreparation,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.IconPath,
		arg.YieldsNothing,
		arg.RestrictToIngredients,
		arg.ZeroIngredientsAllowable,
		arg.PastTense,
	)
	return err
}
