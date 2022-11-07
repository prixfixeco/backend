// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: users_exists_with_status.sql

package generated

import (
	"context"
)

const UserExistsWithStatus = `-- name: UserExistsWithStatus :one
SELECT EXISTS ( SELECT users.id FROM users WHERE users.archived_at IS NULL AND users.id = $1 AND (users.user_account_status = $2 OR users.user_account_status = $3) )
`

type UserExistsWithStatusParams struct {
	ID                  string `db:"id"`
	UserAccountStatus   string `db:"user_account_status"`
	UserAccountStatus_2 string `db:"user_account_status_2"`
}

func (q *Queries) UserExistsWithStatus(ctx context.Context, db DBTX, arg *UserExistsWithStatusParams) (bool, error) {
	row := db.QueryRowContext(ctx, UserExistsWithStatus, arg.ID, arg.UserAccountStatus, arg.UserAccountStatus_2)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}
