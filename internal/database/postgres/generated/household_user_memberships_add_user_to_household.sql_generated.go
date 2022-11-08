// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: household_user_memberships_add_user_to_household.sql

package generated

import (
	"context"
)

const CreateHouseholdUserMembership = `-- name: CreateHouseholdUserMembership :exec
INSERT INTO household_user_memberships (id,belongs_to_user,belongs_to_household,household_roles)
VALUES ($1,$2,$3,$4)
`

type CreateHouseholdUserMembershipParams struct {
	ID                 string `db:"id"`
	BelongsToUser      string `db:"belongs_to_user"`
	BelongsToHousehold string `db:"belongs_to_household"`
	HouseholdRoles     string `db:"household_roles"`
}

func (q *Queries) CreateHouseholdUserMembership(ctx context.Context, arg *CreateHouseholdUserMembershipParams) error {
	_, err := q.exec(ctx, q.createHouseholdUserMembershipStmt, CreateHouseholdUserMembership,
		arg.ID,
		arg.BelongsToUser,
		arg.BelongsToHousehold,
		arg.HouseholdRoles,
	)
	return err
}
