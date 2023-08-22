// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: household_invitations.sql

package generated

import (
	"context"
	"database/sql"
	"time"
)

const attachHouseholdInvitationsToUserID = `-- name: AttachHouseholdInvitationsToUserID :exec

UPDATE household_invitations SET
    to_user = $1,
    last_updated_at = NOW()
WHERE archived_at IS NULL
  AND to_email = LOWER($2)
`

type AttachHouseholdInvitationsToUserIDParams struct {
	EmailAddress string
	UserID       sql.NullString
}

func (q *Queries) AttachHouseholdInvitationsToUserID(ctx context.Context, db DBTX, arg *AttachHouseholdInvitationsToUserIDParams) error {
	_, err := db.ExecContext(ctx, attachHouseholdInvitationsToUserID, arg.UserID, arg.EmailAddress)
	return err
}

const checkHouseholdInvitationExistence = `-- name: CheckHouseholdInvitationExistence :one

SELECT EXISTS ( SELECT household_invitations.id FROM household_invitations WHERE household_invitations.archived_at IS NULL AND household_invitations.id = $1 )
`

func (q *Queries) CheckHouseholdInvitationExistence(ctx context.Context, db DBTX, id string) (bool, error) {
	row := db.QueryRowContext(ctx, checkHouseholdInvitationExistence, id)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const createHouseholdInvitation = `-- name: CreateHouseholdInvitation :exec

INSERT INTO household_invitations (id,from_user,to_user,to_name,note,to_email,token,destination_household,expires_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
`

type CreateHouseholdInvitationParams struct {
	ExpiresAt            time.Time
	ID                   string
	FromUser             string
	ToName               string
	Note                 string
	ToEmail              string
	Token                string
	DestinationHousehold string
	ToUser               sql.NullString
}

func (q *Queries) CreateHouseholdInvitation(ctx context.Context, db DBTX, arg *CreateHouseholdInvitationParams) error {
	_, err := db.ExecContext(ctx, createHouseholdInvitation,
		arg.ID,
		arg.FromUser,
		arg.ToUser,
		arg.ToName,
		arg.Note,
		arg.ToEmail,
		arg.Token,
		arg.DestinationHousehold,
		arg.ExpiresAt,
	)
	return err
}

const getHouseholdInvitationByEmailAndToken = `-- name: GetHouseholdInvitationByEmailAndToken :one

SELECT
	household_invitations.id,
	households.id,
	households.name,
	households.billing_status,
	households.contact_phone,
	households.address_line_1,
	households.address_line_2,
	households.city,
	households.state,
	households.zip_code,
	households.country,
	households.latitude,
    households.longitude,
	households.payment_processor_customer_id,
	households.subscription_plan_id,
	households.created_at,
	households.last_updated_at,
	households.archived_at,
	households.belongs_to_user,
	household_invitations.to_email,
	household_invitations.to_user,
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
	users.created_at,
	users.last_updated_at,
	users.archived_at,
	household_invitations.to_name,
	household_invitations.status,
	household_invitations.note,
	household_invitations.status_note,
	household_invitations.token,
	household_invitations.expires_at,
	household_invitations.created_at,
	household_invitations.last_updated_at,
	household_invitations.archived_at
FROM household_invitations
	JOIN households ON household_invitations.destination_household = households.id
	JOIN users ON household_invitations.from_user = users.id
WHERE household_invitations.archived_at IS NULL
	AND household_invitations.expires_at > NOW()
	AND household_invitations.to_email = LOWER($1)
	AND household_invitations.token = $2
`

type GetHouseholdInvitationByEmailAndTokenParams struct {
	Lower string
	Token string
}

type GetHouseholdInvitationByEmailAndTokenRow struct {
	ExpiresAt                    time.Time
	CreatedAt_2                  time.Time
	CreatedAt                    time.Time
	CreatedAt_3                  time.Time
	LastUpdatedAt_2              sql.NullTime
	ArchivedAt_2                 sql.NullTime
	LastUpdatedAt_3              sql.NullTime
	ArchivedAt_3                 sql.NullTime
	Birthday                     sql.NullTime
	TwoFactorSecretVerifiedAt    sql.NullTime
	PasswordLastChangedAt        sql.NullTime
	EmailAddressVerifiedAt       sql.NullTime
	ArchivedAt                   sql.NullTime
	LastUpdatedAt                sql.NullTime
	Username                     string
	TwoFactorSecret              string
	PaymentProcessorCustomerID   string
	ID_2                         string
	BelongsToUser                string
	ToEmail                      string
	Name                         string
	ID_3                         string
	FirstName                    string
	LastName                     string
	ID                           string
	EmailAddress                 string
	BillingStatus                string
	ContactPhone                 string
	HashedPassword               string
	Token                        string
	Country                      string
	StatusNote                   string
	ZipCode                      string
	ServiceRole                  string
	UserAccountStatus            string
	UserAccountStatusExplanation string
	State                        string
	City                         string
	AddressLine2                 string
	AddressLine1                 string
	ToName                       string
	Status                       InvitationState
	Note                         string
	SubscriptionPlanID           sql.NullString
	AvatarSrc                    sql.NullString
	Latitude                     sql.NullString
	ToUser                       sql.NullString
	Longitude                    sql.NullString
	RequiresPasswordChange       bool
}

func (q *Queries) GetHouseholdInvitationByEmailAndToken(ctx context.Context, db DBTX, arg *GetHouseholdInvitationByEmailAndTokenParams) (*GetHouseholdInvitationByEmailAndTokenRow, error) {
	row := db.QueryRowContext(ctx, getHouseholdInvitationByEmailAndToken, arg.Lower, arg.Token)
	var i GetHouseholdInvitationByEmailAndTokenRow
	err := row.Scan(
		&i.ID,
		&i.ID_2,
		&i.Name,
		&i.BillingStatus,
		&i.ContactPhone,
		&i.AddressLine1,
		&i.AddressLine2,
		&i.City,
		&i.State,
		&i.ZipCode,
		&i.Country,
		&i.Latitude,
		&i.Longitude,
		&i.PaymentProcessorCustomerID,
		&i.SubscriptionPlanID,
		&i.CreatedAt,
		&i.LastUpdatedAt,
		&i.ArchivedAt,
		&i.BelongsToUser,
		&i.ToEmail,
		&i.ToUser,
		&i.ID_3,
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
		&i.CreatedAt_2,
		&i.LastUpdatedAt_2,
		&i.ArchivedAt_2,
		&i.ToName,
		&i.Status,
		&i.Note,
		&i.StatusNote,
		&i.Token,
		&i.ExpiresAt,
		&i.CreatedAt_3,
		&i.LastUpdatedAt_3,
		&i.ArchivedAt_3,
	)
	return &i, err
}

const getHouseholdInvitationByHouseholdAndID = `-- name: GetHouseholdInvitationByHouseholdAndID :one

SELECT
	household_invitations.id,
	households.id,
	households.name,
	households.billing_status,
	households.contact_phone,
	households.address_line_1,
	households.address_line_2,
	households.city,
	households.state,
	households.zip_code,
	households.country,
	households.latitude,
    households.longitude,
	households.payment_processor_customer_id,
	households.subscription_plan_id,
	households.created_at,
	households.last_updated_at,
	households.archived_at,
	households.belongs_to_user,
	household_invitations.to_email,
	household_invitations.to_user,
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
	users.created_at,
	users.last_updated_at,
	users.archived_at,
	household_invitations.to_name,
	household_invitations.status,
	household_invitations.note,
	household_invitations.status_note,
	household_invitations.token,
	household_invitations.expires_at,
	household_invitations.created_at,
	household_invitations.last_updated_at,
	household_invitations.archived_at
FROM household_invitations
	LEFT JOIN households ON household_invitations.destination_household = households.id
	LEFT JOIN users ON household_invitations.from_user = users.id
WHERE household_invitations.archived_at IS NULL
      AND household_invitations.expires_at > NOW()
      AND household_invitations.expires_at > NOW()
	AND household_invitations.destination_household = $1
	AND household_invitations.id = $2
`

type GetHouseholdInvitationByHouseholdAndIDParams struct {
	DestinationHousehold string
	ID                   string
}

type GetHouseholdInvitationByHouseholdAndIDRow struct {
	CreatedAt_3                  time.Time
	ExpiresAt                    time.Time
	CreatedAt                    sql.NullTime
	EmailAddressVerifiedAt       sql.NullTime
	ArchivedAt_3                 sql.NullTime
	PasswordLastChangedAt        sql.NullTime
	TwoFactorSecretVerifiedAt    sql.NullTime
	ArchivedAt                   sql.NullTime
	LastUpdatedAt                sql.NullTime
	LastUpdatedAt_3              sql.NullTime
	ArchivedAt_2                 sql.NullTime
	LastUpdatedAt_2              sql.NullTime
	Birthday                     sql.NullTime
	CreatedAt_2                  sql.NullTime
	ID                           string
	ToName                       string
	Status                       InvitationState
	Note                         string
	StatusNote                   string
	ToEmail                      string
	Token                        string
	Longitude                    sql.NullString
	UserAccountStatus            sql.NullString
	LastName                     sql.NullString
	Username                     sql.NullString
	EmailAddress                 sql.NullString
	ID_3                         sql.NullString
	AvatarSrc                    sql.NullString
	HashedPassword               sql.NullString
	ID_2                         sql.NullString
	ToUser                       sql.NullString
	TwoFactorSecret              sql.NullString
	BelongsToUser                sql.NullString
	ServiceRole                  sql.NullString
	FirstName                    sql.NullString
	UserAccountStatusExplanation sql.NullString
	SubscriptionPlanID           sql.NullString
	PaymentProcessorCustomerID   sql.NullString
	Latitude                     sql.NullString
	Country                      sql.NullString
	ZipCode                      sql.NullString
	State                        sql.NullString
	City                         sql.NullString
	AddressLine2                 sql.NullString
	AddressLine1                 sql.NullString
	ContactPhone                 sql.NullString
	BillingStatus                sql.NullString
	Name                         sql.NullString
	RequiresPasswordChange       sql.NullBool
}

func (q *Queries) GetHouseholdInvitationByHouseholdAndID(ctx context.Context, db DBTX, arg *GetHouseholdInvitationByHouseholdAndIDParams) (*GetHouseholdInvitationByHouseholdAndIDRow, error) {
	row := db.QueryRowContext(ctx, getHouseholdInvitationByHouseholdAndID, arg.DestinationHousehold, arg.ID)
	var i GetHouseholdInvitationByHouseholdAndIDRow
	err := row.Scan(
		&i.ID,
		&i.ID_2,
		&i.Name,
		&i.BillingStatus,
		&i.ContactPhone,
		&i.AddressLine1,
		&i.AddressLine2,
		&i.City,
		&i.State,
		&i.ZipCode,
		&i.Country,
		&i.Latitude,
		&i.Longitude,
		&i.PaymentProcessorCustomerID,
		&i.SubscriptionPlanID,
		&i.CreatedAt,
		&i.LastUpdatedAt,
		&i.ArchivedAt,
		&i.BelongsToUser,
		&i.ToEmail,
		&i.ToUser,
		&i.ID_3,
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
		&i.CreatedAt_2,
		&i.LastUpdatedAt_2,
		&i.ArchivedAt_2,
		&i.ToName,
		&i.Status,
		&i.Note,
		&i.StatusNote,
		&i.Token,
		&i.ExpiresAt,
		&i.CreatedAt_3,
		&i.LastUpdatedAt_3,
		&i.ArchivedAt_3,
	)
	return &i, err
}

const getHouseholdInvitationByTokenAndID = `-- name: GetHouseholdInvitationByTokenAndID :one

SELECT
	household_invitations.id,
	households.id,
	households.name,
	households.billing_status,
	households.contact_phone,
	households.address_line_1,
	households.address_line_2,
	households.city,
	households.state,
	households.zip_code,
	households.country,
	households.latitude,
    households.longitude,
	households.payment_processor_customer_id,
	households.subscription_plan_id,
	households.created_at,
	households.last_updated_at,
	households.archived_at,
	households.belongs_to_user,
	household_invitations.to_email,
	household_invitations.to_user,
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
	users.created_at,
	users.last_updated_at,
	users.archived_at,
	household_invitations.to_name,
	household_invitations.status,
	household_invitations.note,
	household_invitations.status_note,
	household_invitations.token,
	household_invitations.expires_at,
	household_invitations.created_at,
	household_invitations.last_updated_at,
	household_invitations.archived_at
FROM household_invitations
	JOIN households ON household_invitations.destination_household = households.id
	JOIN users ON household_invitations.from_user = users.id
WHERE household_invitations.archived_at IS NULL
	AND household_invitations.expires_at > NOW()
	AND household_invitations.token = $1
	AND household_invitations.id = $2
`

type GetHouseholdInvitationByTokenAndIDParams struct {
	Token string
	ID    string
}

type GetHouseholdInvitationByTokenAndIDRow struct {
	ExpiresAt                    time.Time
	CreatedAt_2                  time.Time
	CreatedAt                    time.Time
	CreatedAt_3                  time.Time
	LastUpdatedAt_2              sql.NullTime
	ArchivedAt_2                 sql.NullTime
	LastUpdatedAt_3              sql.NullTime
	ArchivedAt_3                 sql.NullTime
	Birthday                     sql.NullTime
	TwoFactorSecretVerifiedAt    sql.NullTime
	PasswordLastChangedAt        sql.NullTime
	EmailAddressVerifiedAt       sql.NullTime
	ArchivedAt                   sql.NullTime
	LastUpdatedAt                sql.NullTime
	Username                     string
	TwoFactorSecret              string
	PaymentProcessorCustomerID   string
	ID_2                         string
	BelongsToUser                string
	ToEmail                      string
	Name                         string
	ID_3                         string
	FirstName                    string
	LastName                     string
	ID                           string
	EmailAddress                 string
	BillingStatus                string
	ContactPhone                 string
	HashedPassword               string
	Token                        string
	Country                      string
	StatusNote                   string
	ZipCode                      string
	ServiceRole                  string
	UserAccountStatus            string
	UserAccountStatusExplanation string
	State                        string
	City                         string
	AddressLine2                 string
	AddressLine1                 string
	ToName                       string
	Status                       InvitationState
	Note                         string
	SubscriptionPlanID           sql.NullString
	AvatarSrc                    sql.NullString
	Latitude                     sql.NullString
	ToUser                       sql.NullString
	Longitude                    sql.NullString
	RequiresPasswordChange       bool
}

func (q *Queries) GetHouseholdInvitationByTokenAndID(ctx context.Context, db DBTX, arg *GetHouseholdInvitationByTokenAndIDParams) (*GetHouseholdInvitationByTokenAndIDRow, error) {
	row := db.QueryRowContext(ctx, getHouseholdInvitationByTokenAndID, arg.Token, arg.ID)
	var i GetHouseholdInvitationByTokenAndIDRow
	err := row.Scan(
		&i.ID,
		&i.ID_2,
		&i.Name,
		&i.BillingStatus,
		&i.ContactPhone,
		&i.AddressLine1,
		&i.AddressLine2,
		&i.City,
		&i.State,
		&i.ZipCode,
		&i.Country,
		&i.Latitude,
		&i.Longitude,
		&i.PaymentProcessorCustomerID,
		&i.SubscriptionPlanID,
		&i.CreatedAt,
		&i.LastUpdatedAt,
		&i.ArchivedAt,
		&i.BelongsToUser,
		&i.ToEmail,
		&i.ToUser,
		&i.ID_3,
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
		&i.CreatedAt_2,
		&i.LastUpdatedAt_2,
		&i.ArchivedAt_2,
		&i.ToName,
		&i.Status,
		&i.Note,
		&i.StatusNote,
		&i.Token,
		&i.ExpiresAt,
		&i.CreatedAt_3,
		&i.LastUpdatedAt_3,
		&i.ArchivedAt_3,
	)
	return &i, err
}

const setHouseholdInvitationStatus = `-- name: SetHouseholdInvitationStatus :exec

UPDATE household_invitations SET
	status = $1,
	status_note = $2,
	last_updated_at = NOW(),
	archived_at = NOW()
WHERE archived_at IS NULL
	AND id = $3
`

type SetHouseholdInvitationStatusParams struct {
	Status     InvitationState
	StatusNote string
	ID         string
}

func (q *Queries) SetHouseholdInvitationStatus(ctx context.Context, db DBTX, arg *SetHouseholdInvitationStatusParams) error {
	_, err := db.ExecContext(ctx, setHouseholdInvitationStatus, arg.Status, arg.StatusNote, arg.ID)
	return err
}
