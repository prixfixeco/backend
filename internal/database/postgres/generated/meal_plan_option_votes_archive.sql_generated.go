// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: meal_plan_option_votes_archive.sql

package generated

import (
	"context"
)

const ArchiveMealPlanOptionVote = `-- name: ArchiveMealPlanOptionVote :exec
UPDATE meal_plan_option_votes SET archived_at = NOW() WHERE archived_at IS NULL AND belongs_to_meal_plan_option = $1 AND id = $2
`

type ArchiveMealPlanOptionVoteParams struct {
	BelongsToMealPlanOption string `db:"belongs_to_meal_plan_option"`
	ID                      string `db:"id"`
}

func (q *Queries) ArchiveMealPlanOptionVote(ctx context.Context, arg *ArchiveMealPlanOptionVoteParams) error {
	_, err := q.exec(ctx, q.archiveMealPlanOptionVoteStmt, ArchiveMealPlanOptionVote, arg.BelongsToMealPlanOption, arg.ID)
	return err
}
