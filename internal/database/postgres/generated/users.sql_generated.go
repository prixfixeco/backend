// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: users.sql

package generated

import (
	"context"
	"database/sql"
	"time"
)

const acceptPrivacyPolicyForUser = `-- name: AcceptPrivacyPolicyForUser :exec

UPDATE users SET
	last_accepted_privacy_policy = NOW()
WHERE archived_at IS NULL
	AND id = $1
`

func (q *Queries) AcceptPrivacyPolicyForUser(ctx context.Context, db DBTX, id string) error {
	_, err := db.ExecContext(ctx, acceptPrivacyPolicyForUser, id)
	return err
}

const acceptTermsOfServiceForUser = `-- name: AcceptTermsOfServiceForUser :exec

UPDATE users SET
	last_accepted_terms_of_service = NOW()
WHERE archived_at IS NULL
	AND id = $1
`

func (q *Queries) AcceptTermsOfServiceForUser(ctx context.Context, db DBTX, id string) error {
	_, err := db.ExecContext(ctx, acceptTermsOfServiceForUser, id)
	return err
}

const archiveUser = `-- name: ArchiveUser :exec

UPDATE users SET
	archived_at = NOW()
WHERE archived_at IS NULL
	AND id = $1
`

func (q *Queries) ArchiveUser(ctx context.Context, db DBTX, id string) error {
	_, err := db.ExecContext(ctx, archiveUser, id)
	return err
}

const archiveUserMemberships = `-- name: ArchiveUserMemberships :exec

UPDATE household_user_memberships SET
	archived_at = NOW()
WHERE archived_at IS NULL
	AND belongs_to_user = $1
`

func (q *Queries) ArchiveUserMemberships(ctx context.Context, db DBTX, belongsToUser string) error {
	_, err := db.ExecContext(ctx, archiveUserMemberships, belongsToUser)
	return err
}

const createUser = `-- name: CreateUser :exec

INSERT INTO users (id,first_name,last_name,username,email_address,hashed_password,two_factor_secret,avatar_src,user_account_status,birthday,service_role,email_address_verification_token) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)
`

type CreateUserParams struct {
	ID                            string
	FirstName                     string
	LastName                      string
	Username                      string
	EmailAddress                  string
	HashedPassword                string
	TwoFactorSecret               string
	AvatarSrc                     sql.NullString
	UserAccountStatus             string
	Birthday                      sql.NullTime
	ServiceRole                   string
	EmailAddressVerificationToken sql.NullString
}

func (q *Queries) CreateUser(ctx context.Context, db DBTX, arg *CreateUserParams) error {
	_, err := db.ExecContext(ctx, createUser,
		arg.ID,
		arg.FirstName,
		arg.LastName,
		arg.Username,
		arg.EmailAddress,
		arg.HashedPassword,
		arg.TwoFactorSecret,
		arg.AvatarSrc,
		arg.UserAccountStatus,
		arg.Birthday,
		arg.ServiceRole,
		arg.EmailAddressVerificationToken,
	)
	return err
}

const getAdminUserByUsername = `-- name: GetAdminUserByUsername :one

SELECT
	users.id,
	users.first_name,
	users.last_name,
	users.username,
	users.email_address,
	users.email_address_verified_at,
	users.avatar_src,
	users.hashed_password,
	users.requires_password_change,
	users.password_last_changed_at,
	users.two_factor_secret,
	users.two_factor_secret_verified_at,
	users.service_role,
	users.user_account_status,
	users.user_account_status_explanation,
	users.birthday,
	users.last_accepted_terms_of_service,
    users.last_accepted_privacy_policy,
	users.created_at,
	users.last_updated_at,
	users.archived_at
FROM users
WHERE users.archived_at IS NULL
	AND users.service_role = 'service_admin'
	AND users.username = $1
	AND users.two_factor_secret_verified_at IS NOT NULL
`

type GetAdminUserByUsernameRow struct {
	CreatedAt                    time.Time
	Birthday                     sql.NullTime
	ArchivedAt                   sql.NullTime
	PasswordLastChangedAt        sql.NullTime
	LastUpdatedAt                sql.NullTime
	EmailAddressVerifiedAt       sql.NullTime
	LastAcceptedPrivacyPolicy    sql.NullTime
	LastAcceptedTermsOfService   sql.NullTime
	TwoFactorSecretVerifiedAt    sql.NullTime
	UserAccountStatusExplanation string
	FirstName                    string
	ServiceRole                  string
	UserAccountStatus            string
	LastName                     string
	ID                           string
	HashedPassword               string
	TwoFactorSecret              string
	EmailAddress                 string
	Username                     string
	AvatarSrc                    sql.NullString
	RequiresPasswordChange       bool
}

func (q *Queries) GetAdminUserByUsername(ctx context.Context, db DBTX, username string) (*GetAdminUserByUsernameRow, error) {
	row := db.QueryRowContext(ctx, getAdminUserByUsername, username)
	var i GetAdminUserByUsernameRow
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Username,
		&i.EmailAddress,
		&i.EmailAddressVerifiedAt,
		&i.AvatarSrc,
		&i.HashedPassword,
		&i.RequiresPasswordChange,
		&i.PasswordLastChangedAt,
		&i.TwoFactorSecret,
		&i.TwoFactorSecretVerifiedAt,
		&i.ServiceRole,
		&i.UserAccountStatus,
		&i.UserAccountStatusExplanation,
		&i.Birthday,
		&i.LastAcceptedTermsOfService,
		&i.LastAcceptedPrivacyPolicy,
		&i.CreatedAt,
		&i.LastUpdatedAt,
		&i.ArchivedAt,
	)
	return &i, err
}

const getEmailVerificationTokenByUserID = `-- name: GetEmailVerificationTokenByUserID :one

SELECT
	users.email_address_verification_token
FROM users
WHERE users.archived_at IS NULL
    AND users.email_address_verified_at IS NULL
	AND users.id = $1
`

func (q *Queries) GetEmailVerificationTokenByUserID(ctx context.Context, db DBTX, id string) (sql.NullString, error) {
	row := db.QueryRowContext(ctx, getEmailVerificationTokenByUserID, id)
	var email_address_verification_token sql.NullString
	err := row.Scan(&email_address_verification_token)
	return email_address_verification_token, err
}

const getUserByEmail = `-- name: GetUserByEmail :one

SELECT
	users.id,
	users.first_name,
	users.last_name,
	users.username,
	users.email_address,
	users.email_address_verified_at,
	users.avatar_src,
	users.hashed_password,
	users.requires_password_change,
	users.password_last_changed_at,
	users.two_factor_secret,
	users.two_factor_secret_verified_at,
	users.service_role,
	users.user_account_status,
	users.user_account_status_explanation,
	users.birthday,
	users.last_accepted_terms_of_service,
    users.last_accepted_privacy_policy,
	users.created_at,
	users.last_updated_at,
	users.archived_at
FROM users
WHERE users.archived_at IS NULL
	AND users.email_address = $1
`

type GetUserByEmailRow struct {
	CreatedAt                    time.Time
	Birthday                     sql.NullTime
	ArchivedAt                   sql.NullTime
	PasswordLastChangedAt        sql.NullTime
	LastUpdatedAt                sql.NullTime
	EmailAddressVerifiedAt       sql.NullTime
	LastAcceptedPrivacyPolicy    sql.NullTime
	LastAcceptedTermsOfService   sql.NullTime
	TwoFactorSecretVerifiedAt    sql.NullTime
	UserAccountStatusExplanation string
	FirstName                    string
	ServiceRole                  string
	UserAccountStatus            string
	LastName                     string
	ID                           string
	HashedPassword               string
	TwoFactorSecret              string
	EmailAddress                 string
	Username                     string
	AvatarSrc                    sql.NullString
	RequiresPasswordChange       bool
}

func (q *Queries) GetUserByEmail(ctx context.Context, db DBTX, emailAddress string) (*GetUserByEmailRow, error) {
	row := db.QueryRowContext(ctx, getUserByEmail, emailAddress)
	var i GetUserByEmailRow
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Username,
		&i.EmailAddress,
		&i.EmailAddressVerifiedAt,
		&i.AvatarSrc,
		&i.HashedPassword,
		&i.RequiresPasswordChange,
		&i.PasswordLastChangedAt,
		&i.TwoFactorSecret,
		&i.TwoFactorSecretVerifiedAt,
		&i.ServiceRole,
		&i.UserAccountStatus,
		&i.UserAccountStatusExplanation,
		&i.Birthday,
		&i.LastAcceptedTermsOfService,
		&i.LastAcceptedPrivacyPolicy,
		&i.CreatedAt,
		&i.LastUpdatedAt,
		&i.ArchivedAt,
	)
	return &i, err
}

const getUserByEmailAddressVerificationToken = `-- name: GetUserByEmailAddressVerificationToken :one

SELECT
	users.id,
	users.first_name,
	users.last_name,
	users.username,
	users.email_address,
	users.email_address_verified_at,
	users.avatar_src,
	users.hashed_password,
	users.requires_password_change,
	users.password_last_changed_at,
	users.two_factor_secret,
	users.two_factor_secret_verified_at,
	users.service_role,
	users.user_account_status,
	users.user_account_status_explanation,
	users.birthday,
	users.last_accepted_terms_of_service,
    users.last_accepted_privacy_policy,
	users.created_at,
	users.last_updated_at,
	users.archived_at
FROM users
WHERE users.archived_at IS NULL
	AND users.email_address_verification_token = $1
`

type GetUserByEmailAddressVerificationTokenRow struct {
	CreatedAt                    time.Time
	Birthday                     sql.NullTime
	ArchivedAt                   sql.NullTime
	PasswordLastChangedAt        sql.NullTime
	LastUpdatedAt                sql.NullTime
	EmailAddressVerifiedAt       sql.NullTime
	LastAcceptedPrivacyPolicy    sql.NullTime
	LastAcceptedTermsOfService   sql.NullTime
	TwoFactorSecretVerifiedAt    sql.NullTime
	UserAccountStatusExplanation string
	FirstName                    string
	ServiceRole                  string
	UserAccountStatus            string
	LastName                     string
	ID                           string
	HashedPassword               string
	TwoFactorSecret              string
	EmailAddress                 string
	Username                     string
	AvatarSrc                    sql.NullString
	RequiresPasswordChange       bool
}

func (q *Queries) GetUserByEmailAddressVerificationToken(ctx context.Context, db DBTX, emailAddressVerificationToken sql.NullString) (*GetUserByEmailAddressVerificationTokenRow, error) {
	row := db.QueryRowContext(ctx, getUserByEmailAddressVerificationToken, emailAddressVerificationToken)
	var i GetUserByEmailAddressVerificationTokenRow
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Username,
		&i.EmailAddress,
		&i.EmailAddressVerifiedAt,
		&i.AvatarSrc,
		&i.HashedPassword,
		&i.RequiresPasswordChange,
		&i.PasswordLastChangedAt,
		&i.TwoFactorSecret,
		&i.TwoFactorSecretVerifiedAt,
		&i.ServiceRole,
		&i.UserAccountStatus,
		&i.UserAccountStatusExplanation,
		&i.Birthday,
		&i.LastAcceptedTermsOfService,
		&i.LastAcceptedPrivacyPolicy,
		&i.CreatedAt,
		&i.LastUpdatedAt,
		&i.ArchivedAt,
	)
	return &i, err
}

const getUserByID = `-- name: GetUserByID :one

SELECT
	users.id,
	users.first_name,
	users.last_name,
	users.username,
	users.email_address,
	users.email_address_verified_at,
	users.avatar_src,
	users.hashed_password,
	users.requires_password_change,
	users.password_last_changed_at,
	users.two_factor_secret,
	users.two_factor_secret_verified_at,
	users.service_role,
	users.user_account_status,
	users.user_account_status_explanation,
	users.birthday,
	users.last_accepted_terms_of_service,
    users.last_accepted_privacy_policy,
	users.created_at,
	users.last_updated_at,
	users.archived_at
FROM users
WHERE users.archived_at IS NULL
	AND users.id = $1
`

type GetUserByIDRow struct {
	CreatedAt                    time.Time
	Birthday                     sql.NullTime
	ArchivedAt                   sql.NullTime
	PasswordLastChangedAt        sql.NullTime
	LastUpdatedAt                sql.NullTime
	EmailAddressVerifiedAt       sql.NullTime
	LastAcceptedPrivacyPolicy    sql.NullTime
	LastAcceptedTermsOfService   sql.NullTime
	TwoFactorSecretVerifiedAt    sql.NullTime
	UserAccountStatusExplanation string
	FirstName                    string
	ServiceRole                  string
	UserAccountStatus            string
	LastName                     string
	ID                           string
	HashedPassword               string
	TwoFactorSecret              string
	EmailAddress                 string
	Username                     string
	AvatarSrc                    sql.NullString
	RequiresPasswordChange       bool
}

func (q *Queries) GetUserByID(ctx context.Context, db DBTX, id string) (*GetUserByIDRow, error) {
	row := db.QueryRowContext(ctx, getUserByID, id)
	var i GetUserByIDRow
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Username,
		&i.EmailAddress,
		&i.EmailAddressVerifiedAt,
		&i.AvatarSrc,
		&i.HashedPassword,
		&i.RequiresPasswordChange,
		&i.PasswordLastChangedAt,
		&i.TwoFactorSecret,
		&i.TwoFactorSecretVerifiedAt,
		&i.ServiceRole,
		&i.UserAccountStatus,
		&i.UserAccountStatusExplanation,
		&i.Birthday,
		&i.LastAcceptedTermsOfService,
		&i.LastAcceptedPrivacyPolicy,
		&i.CreatedAt,
		&i.LastUpdatedAt,
		&i.ArchivedAt,
	)
	return &i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one

SELECT
	users.id,
	users.first_name,
	users.last_name,
	users.username,
	users.email_address,
	users.email_address_verified_at,
	users.avatar_src,
	users.hashed_password,
	users.requires_password_change,
	users.password_last_changed_at,
	users.two_factor_secret,
	users.two_factor_secret_verified_at,
	users.service_role,
	users.user_account_status,
	users.user_account_status_explanation,
	users.birthday,
	users.last_accepted_terms_of_service,
    users.last_accepted_privacy_policy,
	users.created_at,
	users.last_updated_at,
	users.archived_at
FROM users
WHERE users.archived_at IS NULL
	AND users.username = $1
`

type GetUserByUsernameRow struct {
	CreatedAt                    time.Time
	Birthday                     sql.NullTime
	ArchivedAt                   sql.NullTime
	PasswordLastChangedAt        sql.NullTime
	LastUpdatedAt                sql.NullTime
	EmailAddressVerifiedAt       sql.NullTime
	LastAcceptedPrivacyPolicy    sql.NullTime
	LastAcceptedTermsOfService   sql.NullTime
	TwoFactorSecretVerifiedAt    sql.NullTime
	UserAccountStatusExplanation string
	FirstName                    string
	ServiceRole                  string
	UserAccountStatus            string
	LastName                     string
	ID                           string
	HashedPassword               string
	TwoFactorSecret              string
	EmailAddress                 string
	Username                     string
	AvatarSrc                    sql.NullString
	RequiresPasswordChange       bool
}

func (q *Queries) GetUserByUsername(ctx context.Context, db DBTX, username string) (*GetUserByUsernameRow, error) {
	row := db.QueryRowContext(ctx, getUserByUsername, username)
	var i GetUserByUsernameRow
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Username,
		&i.EmailAddress,
		&i.EmailAddressVerifiedAt,
		&i.AvatarSrc,
		&i.HashedPassword,
		&i.RequiresPasswordChange,
		&i.PasswordLastChangedAt,
		&i.TwoFactorSecret,
		&i.TwoFactorSecretVerifiedAt,
		&i.ServiceRole,
		&i.UserAccountStatus,
		&i.UserAccountStatusExplanation,
		&i.Birthday,
		&i.LastAcceptedTermsOfService,
		&i.LastAcceptedPrivacyPolicy,
		&i.CreatedAt,
		&i.LastUpdatedAt,
		&i.ArchivedAt,
	)
	return &i, err
}

const getUserIDsNeedingIndexing = `-- name: GetUserIDsNeedingIndexing :many

SELECT users.id
  FROM users
 WHERE (users.archived_at IS NULL)
       AND (
			(
				users.last_indexed_at IS NULL
			)
			OR users.last_indexed_at
				< now() - '24 hours'::INTERVAL
		)
`

func (q *Queries) GetUserIDsNeedingIndexing(ctx context.Context, db DBTX) ([]string, error) {
	rows, err := db.QueryContext(ctx, getUserIDsNeedingIndexing)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []string{}
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		items = append(items, id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserWithUnverifiedTwoFactor = `-- name: GetUserWithUnverifiedTwoFactor :one

SELECT
	users.id,
	users.first_name,
	users.last_name,
	users.username,
	users.email_address,
	users.email_address_verified_at,
	users.avatar_src,
	users.hashed_password,
	users.requires_password_change,
	users.password_last_changed_at,
	users.two_factor_secret,
	users.two_factor_secret_verified_at,
	users.service_role,
	users.user_account_status,
	users.user_account_status_explanation,
	users.birthday,
	users.last_accepted_terms_of_service,
    users.last_accepted_privacy_policy,
	users.created_at,
	users.last_updated_at,
	users.archived_at
FROM users
WHERE users.archived_at IS NULL
	AND users.id = $1
	AND users.two_factor_secret_verified_at IS NULL
`

type GetUserWithUnverifiedTwoFactorRow struct {
	CreatedAt                    time.Time
	Birthday                     sql.NullTime
	ArchivedAt                   sql.NullTime
	PasswordLastChangedAt        sql.NullTime
	LastUpdatedAt                sql.NullTime
	EmailAddressVerifiedAt       sql.NullTime
	LastAcceptedPrivacyPolicy    sql.NullTime
	LastAcceptedTermsOfService   sql.NullTime
	TwoFactorSecretVerifiedAt    sql.NullTime
	UserAccountStatusExplanation string
	FirstName                    string
	ServiceRole                  string
	UserAccountStatus            string
	LastName                     string
	ID                           string
	HashedPassword               string
	TwoFactorSecret              string
	EmailAddress                 string
	Username                     string
	AvatarSrc                    sql.NullString
	RequiresPasswordChange       bool
}

func (q *Queries) GetUserWithUnverifiedTwoFactor(ctx context.Context, db DBTX, id string) (*GetUserWithUnverifiedTwoFactorRow, error) {
	row := db.QueryRowContext(ctx, getUserWithUnverifiedTwoFactor, id)
	var i GetUserWithUnverifiedTwoFactorRow
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Username,
		&i.EmailAddress,
		&i.EmailAddressVerifiedAt,
		&i.AvatarSrc,
		&i.HashedPassword,
		&i.RequiresPasswordChange,
		&i.PasswordLastChangedAt,
		&i.TwoFactorSecret,
		&i.TwoFactorSecretVerifiedAt,
		&i.ServiceRole,
		&i.UserAccountStatus,
		&i.UserAccountStatusExplanation,
		&i.Birthday,
		&i.LastAcceptedTermsOfService,
		&i.LastAcceptedPrivacyPolicy,
		&i.CreatedAt,
		&i.LastUpdatedAt,
		&i.ArchivedAt,
	)
	return &i, err
}

const getUserWithVerifiedTwoFactor = `-- name: GetUserWithVerifiedTwoFactor :one

SELECT
	users.id,
	users.first_name,
	users.last_name,
	users.username,
	users.email_address,
	users.email_address_verified_at,
	users.avatar_src,
	users.hashed_password,
	users.requires_password_change,
	users.password_last_changed_at,
	users.two_factor_secret,
	users.two_factor_secret_verified_at,
	users.service_role,
	users.user_account_status,
	users.user_account_status_explanation,
	users.birthday,
	users.last_accepted_terms_of_service,
    users.last_accepted_privacy_policy,
	users.created_at,
	users.last_updated_at,
	users.archived_at
FROM users
WHERE users.archived_at IS NULL
	AND users.id = $1
	AND users.two_factor_secret_verified_at IS NOT NULL
`

type GetUserWithVerifiedTwoFactorRow struct {
	CreatedAt                    time.Time
	Birthday                     sql.NullTime
	ArchivedAt                   sql.NullTime
	PasswordLastChangedAt        sql.NullTime
	LastUpdatedAt                sql.NullTime
	EmailAddressVerifiedAt       sql.NullTime
	LastAcceptedPrivacyPolicy    sql.NullTime
	LastAcceptedTermsOfService   sql.NullTime
	TwoFactorSecretVerifiedAt    sql.NullTime
	UserAccountStatusExplanation string
	FirstName                    string
	ServiceRole                  string
	UserAccountStatus            string
	LastName                     string
	ID                           string
	HashedPassword               string
	TwoFactorSecret              string
	EmailAddress                 string
	Username                     string
	AvatarSrc                    sql.NullString
	RequiresPasswordChange       bool
}

func (q *Queries) GetUserWithVerifiedTwoFactor(ctx context.Context, db DBTX, id string) (*GetUserWithVerifiedTwoFactorRow, error) {
	row := db.QueryRowContext(ctx, getUserWithVerifiedTwoFactor, id)
	var i GetUserWithVerifiedTwoFactorRow
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Username,
		&i.EmailAddress,
		&i.EmailAddressVerifiedAt,
		&i.AvatarSrc,
		&i.HashedPassword,
		&i.RequiresPasswordChange,
		&i.PasswordLastChangedAt,
		&i.TwoFactorSecret,
		&i.TwoFactorSecretVerifiedAt,
		&i.ServiceRole,
		&i.UserAccountStatus,
		&i.UserAccountStatusExplanation,
		&i.Birthday,
		&i.LastAcceptedTermsOfService,
		&i.LastAcceptedPrivacyPolicy,
		&i.CreatedAt,
		&i.LastUpdatedAt,
		&i.ArchivedAt,
	)
	return &i, err
}

const getUsers = `-- name: GetUsers :many

SELECT
	users.id,
	users.first_name,
	users.last_name,
	users.username,
	users.email_address,
	users.email_address_verified_at,
	users.avatar_src,
	users.hashed_password,
	users.requires_password_change,
	users.password_last_changed_at,
	users.two_factor_secret,
	users.two_factor_secret_verified_at,
	users.service_role,
	users.user_account_status,
	users.user_account_status_explanation,
	users.birthday,
	users.last_accepted_terms_of_service,
    users.last_accepted_privacy_policy,
	users.created_at,
	users.last_updated_at,
	users.archived_at,
    (
        SELECT
            COUNT(users.id)
        FROM
            users
        WHERE
            users.archived_at IS NULL
          AND users.created_at > COALESCE($1, (SELECT NOW() - interval '999 years'))
          AND users.created_at < COALESCE($2, (SELECT NOW() + interval '999 years'))
          AND (
                users.last_updated_at IS NULL
                OR users.last_updated_at > COALESCE($3, (SELECT NOW() - interval '999 years'))
            )
          AND (
                users.last_updated_at IS NULL
                OR users.last_updated_at < COALESCE($4, (SELECT NOW() + interval '999 years'))
            )
        OFFSET $5
    ) as filtered_count,
    (
        SELECT
            COUNT(users.id)
        FROM
            users
        WHERE
            users.archived_at IS NULL
    ) as total_count
FROM users
WHERE
    users.archived_at IS NULL
  AND users.created_at > COALESCE($1, (SELECT NOW() - interval '999 years'))
  AND users.created_at < COALESCE($2, (SELECT NOW() + interval '999 years'))
  AND (
        users.last_updated_at IS NULL
        OR users.last_updated_at > COALESCE($3, (SELECT NOW() - interval '999 years'))
    )
  AND (
        users.last_updated_at IS NULL
        OR users.last_updated_at < COALESCE($4, (SELECT NOW() + interval '999 years'))
    )
OFFSET $5
LIMIT $6
`

type GetUsersParams struct {
	CreatedBefore sql.NullTime
	CreatedAfter  sql.NullTime
	UpdatedBefore sql.NullTime
	UpdatedAfter  sql.NullTime
	QueryOffset   sql.NullInt32
	QueryLimit    sql.NullInt32
}

type GetUsersRow struct {
	CreatedAt                    time.Time
	Birthday                     sql.NullTime
	PasswordLastChangedAt        sql.NullTime
	TwoFactorSecretVerifiedAt    sql.NullTime
	ArchivedAt                   sql.NullTime
	EmailAddressVerifiedAt       sql.NullTime
	LastUpdatedAt                sql.NullTime
	LastAcceptedPrivacyPolicy    sql.NullTime
	LastAcceptedTermsOfService   sql.NullTime
	HashedPassword               string
	TwoFactorSecret              string
	ServiceRole                  string
	Username                     string
	UserAccountStatus            string
	UserAccountStatusExplanation string
	ID                           string
	EmailAddress                 string
	LastName                     string
	FirstName                    string
	AvatarSrc                    sql.NullString
	FilteredCount                int64
	TotalCount                   int64
	RequiresPasswordChange       bool
}

func (q *Queries) GetUsers(ctx context.Context, db DBTX, arg *GetUsersParams) ([]*GetUsersRow, error) {
	rows, err := db.QueryContext(ctx, getUsers,
		arg.CreatedBefore,
		arg.CreatedAfter,
		arg.UpdatedBefore,
		arg.UpdatedAfter,
		arg.QueryOffset,
		arg.QueryLimit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*GetUsersRow{}
	for rows.Next() {
		var i GetUsersRow
		if err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.Username,
			&i.EmailAddress,
			&i.EmailAddressVerifiedAt,
			&i.AvatarSrc,
			&i.HashedPassword,
			&i.RequiresPasswordChange,
			&i.PasswordLastChangedAt,
			&i.TwoFactorSecret,
			&i.TwoFactorSecretVerifiedAt,
			&i.ServiceRole,
			&i.UserAccountStatus,
			&i.UserAccountStatusExplanation,
			&i.Birthday,
			&i.LastAcceptedTermsOfService,
			&i.LastAcceptedPrivacyPolicy,
			&i.CreatedAt,
			&i.LastUpdatedAt,
			&i.ArchivedAt,
			&i.FilteredCount,
			&i.TotalCount,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const markEmailAddressAsVerified = `-- name: MarkEmailAddressAsVerified :exec

UPDATE users SET
	email_address_verified_at = NOW(),
	last_updated_at = NOW()
WHERE email_address_verified_at IS NULL
	AND id = $1
	AND email_address_verification_token = $2
`

type MarkEmailAddressAsVerifiedParams struct {
	ID                            string
	EmailAddressVerificationToken sql.NullString
}

func (q *Queries) MarkEmailAddressAsVerified(ctx context.Context, db DBTX, arg *MarkEmailAddressAsVerifiedParams) error {
	_, err := db.ExecContext(ctx, markEmailAddressAsVerified, arg.ID, arg.EmailAddressVerificationToken)
	return err
}

const markTwoFactorSecretAsUnverified = `-- name: MarkTwoFactorSecretAsUnverified :exec

UPDATE users SET
	two_factor_secret_verified_at = NULL,
	two_factor_secret = $1,
	last_updated_at = NOW()
WHERE archived_at IS NULL
	AND id = $2
`

type MarkTwoFactorSecretAsUnverifiedParams struct {
	TwoFactorSecret string
	ID              string
}

func (q *Queries) MarkTwoFactorSecretAsUnverified(ctx context.Context, db DBTX, arg *MarkTwoFactorSecretAsUnverifiedParams) error {
	_, err := db.ExecContext(ctx, markTwoFactorSecretAsUnverified, arg.TwoFactorSecret, arg.ID)
	return err
}

const markTwoFactorSecretAsVerified = `-- name: MarkTwoFactorSecretAsVerified :exec

UPDATE users SET
	two_factor_secret_verified_at = NOW(),
	last_updated_at = NOW()
WHERE archived_at IS NULL
	AND id = $1
`

func (q *Queries) MarkTwoFactorSecretAsVerified(ctx context.Context, db DBTX, id string) error {
	_, err := db.ExecContext(ctx, markTwoFactorSecretAsVerified, id)
	return err
}

const searchUsersByUsername = `-- name: SearchUsersByUsername :many

SELECT
	users.id,
	users.first_name,
	users.last_name,
	users.username,
	users.email_address,
	users.email_address_verified_at,
	users.avatar_src,
	users.hashed_password,
	users.requires_password_change,
	users.password_last_changed_at,
	users.two_factor_secret,
	users.two_factor_secret_verified_at,
	users.service_role,
	users.user_account_status,
	users.user_account_status_explanation,
	users.birthday,
	users.last_accepted_terms_of_service,
    users.last_accepted_privacy_policy,
	users.created_at,
	users.last_updated_at,
	users.archived_at
FROM users
WHERE users.username ILIKE '%' || $1::text || '%'
AND users.archived_at IS NULL
`

type SearchUsersByUsernameRow struct {
	CreatedAt                    time.Time
	Birthday                     sql.NullTime
	ArchivedAt                   sql.NullTime
	PasswordLastChangedAt        sql.NullTime
	LastUpdatedAt                sql.NullTime
	EmailAddressVerifiedAt       sql.NullTime
	LastAcceptedPrivacyPolicy    sql.NullTime
	LastAcceptedTermsOfService   sql.NullTime
	TwoFactorSecretVerifiedAt    sql.NullTime
	UserAccountStatusExplanation string
	FirstName                    string
	ServiceRole                  string
	UserAccountStatus            string
	LastName                     string
	ID                           string
	HashedPassword               string
	TwoFactorSecret              string
	EmailAddress                 string
	Username                     string
	AvatarSrc                    sql.NullString
	RequiresPasswordChange       bool
}

func (q *Queries) SearchUsersByUsername(ctx context.Context, db DBTX, username string) ([]*SearchUsersByUsernameRow, error) {
	rows, err := db.QueryContext(ctx, searchUsersByUsername, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*SearchUsersByUsernameRow{}
	for rows.Next() {
		var i SearchUsersByUsernameRow
		if err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.Username,
			&i.EmailAddress,
			&i.EmailAddressVerifiedAt,
			&i.AvatarSrc,
			&i.HashedPassword,
			&i.RequiresPasswordChange,
			&i.PasswordLastChangedAt,
			&i.TwoFactorSecret,
			&i.TwoFactorSecretVerifiedAt,
			&i.ServiceRole,
			&i.UserAccountStatus,
			&i.UserAccountStatusExplanation,
			&i.Birthday,
			&i.LastAcceptedTermsOfService,
			&i.LastAcceptedPrivacyPolicy,
			&i.CreatedAt,
			&i.LastUpdatedAt,
			&i.ArchivedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUser = `-- name: UpdateUser :exec

UPDATE users SET
	username = $1,
	first_name = $2,
	last_name = $3,
	hashed_password = $4,
	avatar_src = $5,
	birthday = $6,
	last_updated_at = NOW()
WHERE archived_at IS NULL
	AND id = $7
`

type UpdateUserParams struct {
	Birthday       sql.NullTime
	Username       string
	FirstName      string
	LastName       string
	HashedPassword string
	ID             string
	AvatarSrc      sql.NullString
}

func (q *Queries) UpdateUser(ctx context.Context, db DBTX, arg *UpdateUserParams) error {
	_, err := db.ExecContext(ctx, updateUser,
		arg.Username,
		arg.FirstName,
		arg.LastName,
		arg.HashedPassword,
		arg.AvatarSrc,
		arg.Birthday,
		arg.ID,
	)
	return err
}

const updateUserAvatarSrc = `-- name: UpdateUserAvatarSrc :exec

UPDATE users SET
	avatar_src = $1,
	last_updated_at = NOW()
WHERE archived_at IS NULL
	AND id = $2
`

type UpdateUserAvatarSrcParams struct {
	ID        string
	AvatarSrc sql.NullString
}

func (q *Queries) UpdateUserAvatarSrc(ctx context.Context, db DBTX, arg *UpdateUserAvatarSrcParams) error {
	_, err := db.ExecContext(ctx, updateUserAvatarSrc, arg.AvatarSrc, arg.ID)
	return err
}

const updateUserDetails = `-- name: UpdateUserDetails :exec

UPDATE users SET
	first_name = $1,
	last_name = $2,
	birthday = $3,
	last_updated_at = NOW()
WHERE archived_at IS NULL
	AND id = $4
`

type UpdateUserDetailsParams struct {
	FirstName string
	LastName  string
	Birthday  sql.NullTime
	ID        string
}

func (q *Queries) UpdateUserDetails(ctx context.Context, db DBTX, arg *UpdateUserDetailsParams) error {
	_, err := db.ExecContext(ctx, updateUserDetails,
		arg.FirstName,
		arg.LastName,
		arg.Birthday,
		arg.ID,
	)
	return err
}

const updateUserEmailAddress = `-- name: UpdateUserEmailAddress :exec

UPDATE users SET
	email_address = $1,
	email_address_verified_at = NULL,
	last_updated_at = NOW()
WHERE archived_at IS NULL
	AND id = $2
`

type UpdateUserEmailAddressParams struct {
	EmailAddress string
	ID           string
}

func (q *Queries) UpdateUserEmailAddress(ctx context.Context, db DBTX, arg *UpdateUserEmailAddressParams) error {
	_, err := db.ExecContext(ctx, updateUserEmailAddress, arg.EmailAddress, arg.ID)
	return err
}

const updateUserLastIndexedAt = `-- name: UpdateUserLastIndexedAt :exec

UPDATE users SET last_indexed_at = NOW() WHERE id = $1 AND archived_at IS NULL
`

func (q *Queries) UpdateUserLastIndexedAt(ctx context.Context, db DBTX, id string) error {
	_, err := db.ExecContext(ctx, updateUserLastIndexedAt, id)
	return err
}

const updateUserPassword = `-- name: UpdateUserPassword :exec

UPDATE users SET
	hashed_password = $1,
	password_last_changed_at = NOW(),
	last_updated_at = NOW()
WHERE archived_at IS NULL
	AND id = $2
`

type UpdateUserPasswordParams struct {
	HashedPassword string
	ID             string
}

func (q *Queries) UpdateUserPassword(ctx context.Context, db DBTX, arg *UpdateUserPasswordParams) error {
	_, err := db.ExecContext(ctx, updateUserPassword, arg.HashedPassword, arg.ID)
	return err
}

const updateUserTwoFactorSecret = `-- name: UpdateUserTwoFactorSecret :exec

UPDATE users SET
	two_factor_secret_verified_at = NULL,
	two_factor_secret = $1,
	last_updated_at = NOW()
WHERE archived_at IS NULL
	AND id = $2
`

type UpdateUserTwoFactorSecretParams struct {
	TwoFactorSecret string
	ID              string
}

func (q *Queries) UpdateUserTwoFactorSecret(ctx context.Context, db DBTX, arg *UpdateUserTwoFactorSecretParams) error {
	_, err := db.ExecContext(ctx, updateUserTwoFactorSecret, arg.TwoFactorSecret, arg.ID)
	return err
}

const updateUserUsername = `-- name: UpdateUserUsername :exec

UPDATE users SET
	username = $1,
	last_updated_at = NOW()
WHERE archived_at IS NULL
	AND id = $2
`

type UpdateUserUsernameParams struct {
	Username string
	ID       string
}

func (q *Queries) UpdateUserUsername(ctx context.Context, db DBTX, arg *UpdateUserUsernameParams) error {
	_, err := db.ExecContext(ctx, updateUserUsername, arg.Username, arg.ID)
	return err
}
