// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: users_get_by_email.sql

package generated

import (
	"context"
	"database/sql"
	"time"
)

const GetUserByEmailAddress = `-- name: GetUserByEmailAddress :one
SELECT
	users.id,
	users.username,
	users.email_address,
	users.avatar_src,
	users.hashed_password,
	users.requires_password_change,
	users.password_last_changed_at,
	users.two_factor_secret,
	users.two_factor_secret_verified_at,
	users.service_roles,
	users.user_account_status,
	users.user_account_status_explanation,
	users.birth_day,
	users.birth_month,
	users.created_at,
	users.last_updated_at,
	users.archived_at
FROM users
WHERE users.archived_at IS NULL
	AND users.email_address = $1
`

type GetUserByEmailAddressRow struct {
	CreatedAt                    time.Time      `db:"created_at"`
	PasswordLastChangedAt        sql.NullTime   `db:"password_last_changed_at"`
	LastUpdatedAt                sql.NullTime   `db:"last_updated_at"`
	TwoFactorSecretVerifiedAt    sql.NullTime   `db:"two_factor_secret_verified_at"`
	ArchivedAt                   sql.NullTime   `db:"archived_at"`
	UserAccountStatusExplanation string         `db:"user_account_status_explanation"`
	HashedPassword               string         `db:"hashed_password"`
	TwoFactorSecret              string         `db:"two_factor_secret"`
	EmailAddress                 string         `db:"email_address"`
	ServiceRoles                 string         `db:"service_roles"`
	UserAccountStatus            string         `db:"user_account_status"`
	Username                     string         `db:"username"`
	ID                           string         `db:"id"`
	AvatarSrc                    sql.NullString `db:"avatar_src"`
	BirthDay                     sql.NullInt16  `db:"birth_day"`
	BirthMonth                   sql.NullInt16  `db:"birth_month"`
	RequiresPasswordChange       bool           `db:"requires_password_change"`
}

func (q *Queries) GetUserByEmailAddress(ctx context.Context, emailAddress string) (*GetUserByEmailAddressRow, error) {
	row := q.queryRow(ctx, q.getUserByEmailAddressStmt, GetUserByEmailAddress, emailAddress)
	var i GetUserByEmailAddressRow
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.EmailAddress,
		&i.AvatarSrc,
		&i.HashedPassword,
		&i.RequiresPasswordChange,
		&i.PasswordLastChangedAt,
		&i.TwoFactorSecret,
		&i.TwoFactorSecretVerifiedAt,
		&i.ServiceRoles,
		&i.UserAccountStatus,
		&i.UserAccountStatusExplanation,
		&i.BirthDay,
		&i.BirthMonth,
		&i.CreatedAt,
		&i.LastUpdatedAt,
		&i.ArchivedAt,
	)
	return &i, err
}
