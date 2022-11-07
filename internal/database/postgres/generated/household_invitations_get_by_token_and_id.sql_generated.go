// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: household_invitations_get_by_token_and_id.sql

package generated

import (
	"context"
	"database/sql"
	"time"
)

const GetHouseholdInvitationByTokenAndID = `-- name: GetHouseholdInvitationByTokenAndID :one
SELECT
	household_invitations.id,
	households.id,
	households.name,
	households.billing_status,
	households.contact_email,
	households.contact_phone,
	households.payment_processor_customer_id,
	households.subscription_plan_id,
	households.time_zone,
	households.created_at,
	households.last_updated_at,
	households.archived_at,
	households.belongs_to_user,
	household_invitations.to_email,
	household_invitations.to_user,
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
	users.archived_at,
	household_invitations.status,
	household_invitations.note,
	household_invitations.status_note,
	household_invitations.token,
	household_invitations.created_at,
	household_invitations.last_updated_at,
	household_invitations.archived_at
FROM household_invitations
	JOIN households ON household_invitations.destination_household = households.id
	JOIN users ON household_invitations.from_user = users.id
WHERE household_invitations.archived_at IS NULL
	AND household_invitations.token = $1
	AND household_invitations.id = $2
`

type GetHouseholdInvitationByTokenAndIDParams struct {
	Token string `db:"token"`
	ID    string `db:"id"`
}

type GetHouseholdInvitationByTokenAndIDRow struct {
	CreatedAt_3                  time.Time       `db:"created_at_3"`
	CreatedAt_2                  time.Time       `db:"created_at_2"`
	CreatedAt                    time.Time       `db:"created_at"`
	TwoFactorSecretVerifiedAt    sql.NullTime    `db:"two_factor_secret_verified_at"`
	ArchivedAt_3                 sql.NullTime    `db:"archived_at_3"`
	ArchivedAt_2                 sql.NullTime    `db:"archived_at_2"`
	LastUpdatedAt                sql.NullTime    `db:"last_updated_at"`
	LastUpdatedAt_3              sql.NullTime    `db:"last_updated_at_3"`
	ArchivedAt                   sql.NullTime    `db:"archived_at"`
	PasswordLastChangedAt        sql.NullTime    `db:"password_last_changed_at"`
	LastUpdatedAt_2              sql.NullTime    `db:"last_updated_at_2"`
	TimeZone                     TimeZone        `db:"time_zone"`
	BelongsToUser                string          `db:"belongs_to_user"`
	ToEmail                      string          `db:"to_email"`
	PaymentProcessorCustomerID   string          `db:"payment_processor_customer_id"`
	ContactPhone                 string          `db:"contact_phone"`
	Username                     string          `db:"username"`
	EmailAddress                 string          `db:"email_address"`
	Token                        string          `db:"token"`
	ContactEmail                 string          `db:"contact_email"`
	StatusNote                   string          `db:"status_note"`
	BillingStatus                string          `db:"billing_status"`
	TwoFactorSecret              string          `db:"two_factor_secret"`
	Name                         string          `db:"name"`
	ServiceRoles                 string          `db:"service_roles"`
	UserAccountStatus            string          `db:"user_account_status"`
	UserAccountStatusExplanation string          `db:"user_account_status_explanation"`
	Note                         string          `db:"note"`
	Status                       InvitationState `db:"status"`
	ID_2                         string          `db:"id_2"`
	ID_3                         string          `db:"id_3"`
	HashedPassword               string          `db:"hashed_password"`
	ID                           string          `db:"id"`
	AvatarSrc                    sql.NullString  `db:"avatar_src"`
	ToUser                       sql.NullString  `db:"to_user"`
	SubscriptionPlanID           sql.NullString  `db:"subscription_plan_id"`
	BirthDay                     sql.NullInt16   `db:"birth_day"`
	BirthMonth                   sql.NullInt16   `db:"birth_month"`
	RequiresPasswordChange       bool            `db:"requires_password_change"`
}

func (q *Queries) GetHouseholdInvitationByTokenAndID(ctx context.Context, db DBTX, arg *GetHouseholdInvitationByTokenAndIDParams) (*GetHouseholdInvitationByTokenAndIDRow, error) {
	row := db.QueryRowContext(ctx, GetHouseholdInvitationByTokenAndID, arg.Token, arg.ID)
	var i GetHouseholdInvitationByTokenAndIDRow
	err := row.Scan(
		&i.ID,
		&i.ID_2,
		&i.Name,
		&i.BillingStatus,
		&i.ContactEmail,
		&i.ContactPhone,
		&i.PaymentProcessorCustomerID,
		&i.SubscriptionPlanID,
		&i.TimeZone,
		&i.CreatedAt,
		&i.LastUpdatedAt,
		&i.ArchivedAt,
		&i.BelongsToUser,
		&i.ToEmail,
		&i.ToUser,
		&i.ID_3,
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
		&i.CreatedAt_2,
		&i.LastUpdatedAt_2,
		&i.ArchivedAt_2,
		&i.Status,
		&i.Note,
		&i.StatusNote,
		&i.Token,
		&i.CreatedAt_3,
		&i.LastUpdatedAt_3,
		&i.ArchivedAt_3,
	)
	return &i, err
}
