// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: households.sql

package generated

import (
	"context"
	"database/sql"
	"time"
)

const addToHouseholdDuringCreation = `-- name: AddToHouseholdDuringCreation :exec

INSERT INTO household_user_memberships (
    id,
	belongs_to_household,
	belongs_to_user,
	household_role
) VALUES (
    $1,
	$2,
	$3,
	$4
)
`

type AddToHouseholdDuringCreationParams struct {
	ID                 string
	BelongsToHousehold string
	BelongsToUser      string
	HouseholdRole      string
}

func (q *Queries) AddToHouseholdDuringCreation(ctx context.Context, db DBTX, arg *AddToHouseholdDuringCreationParams) error {
	_, err := db.ExecContext(ctx, addToHouseholdDuringCreation,
		arg.ID,
		arg.BelongsToHousehold,
		arg.BelongsToUser,
		arg.HouseholdRole,
	)
	return err
}

const archiveHousehold = `-- name: ArchiveHousehold :execrows

UPDATE households SET
    last_updated_at = NOW(),
    archived_at = NOW()
WHERE archived_at IS NULL
    AND belongs_to_user = $1
    AND id = $2
`

type ArchiveHouseholdParams struct {
	BelongsToUser string
	ID            string
}

func (q *Queries) ArchiveHousehold(ctx context.Context, db DBTX, arg *ArchiveHouseholdParams) (int64, error) {
	result, err := db.ExecContext(ctx, archiveHousehold, arg.BelongsToUser, arg.ID)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

const createHousehold = `-- name: CreateHousehold :exec

INSERT INTO households (
    id,
	name,
	billing_status,
	contact_phone,
	belongs_to_user,
	address_line_1,
	address_line_2,
	city,
	state,
	zip_code,
	country,
	latitude,
	longitude,
	webhook_hmac_secret
) VALUES (
    $1,
	$2,
	$3,
	$4,
	$5,
	$6,
	$7,
	$8,
	$9,
	$10,
	$11,
	$12,
	$13,
	$14
)
`

type CreateHouseholdParams struct {
	AddressLine2      string
	ID                string
	BillingStatus     string
	ContactPhone      string
	BelongsToUser     string
	AddressLine1      string
	Name              string
	City              string
	Country           string
	ZipCode           string
	State             string
	WebhookHmacSecret string
	Longitude         sql.NullString
	Latitude          sql.NullString
}

func (q *Queries) CreateHousehold(ctx context.Context, db DBTX, arg *CreateHouseholdParams) error {
	_, err := db.ExecContext(ctx, createHousehold,
		arg.ID,
		arg.Name,
		arg.BillingStatus,
		arg.ContactPhone,
		arg.BelongsToUser,
		arg.AddressLine1,
		arg.AddressLine2,
		arg.City,
		arg.State,
		arg.ZipCode,
		arg.Country,
		arg.Latitude,
		arg.Longitude,
		arg.WebhookHmacSecret,
	)
	return err
}

const getHouseholdByIDWithMemberships = `-- name: GetHouseholdByIDWithMemberships :many

SELECT
	households.id,
	households.name,
	households.billing_status,
	households.contact_phone,
	households.payment_processor_customer_id,
	households.subscription_plan_id,
	households.belongs_to_user,
	households.time_zone,
	households.address_line_1,
	households.address_line_2,
	households.city,
	households.state,
	households.zip_code,
	households.country,
	households.latitude,
	households.longitude,
	households.last_payment_provider_sync_occurred_at,
	households.webhook_hmac_secret,
	households.created_at,
	households.last_updated_at,
	households.archived_at,
	users.id as user_id,
	users.username as user_username,
	users.avatar_src as user_avatar_src,
	users.email_address as user_email_address,
	users.hashed_password as user_hashed_password,
	users.password_last_changed_at as user_password_last_changed_at,
	users.requires_password_change as user_requires_password_change,
	users.two_factor_secret as user_two_factor_secret,
	users.two_factor_secret_verified_at as user_two_factor_secret_verified_at,
	users.service_role as user_service_role,
	users.user_account_status as user_user_account_status,
	users.user_account_status_explanation as user_user_account_status_explanation,
	users.birthday as user_birthday,
	users.email_address_verification_token as user_email_address_verification_token,
	users.email_address_verified_at as user_email_address_verified_at,
	users.first_name as user_first_name,
	users.last_name as user_last_name,
	users.last_accepted_terms_of_service as user_last_accepted_terms_of_service,
	users.last_accepted_privacy_policy as user_last_accepted_privacy_policy,
	users.last_indexed_at as user_last_indexed_at,
	users.created_at as user_created_at,
	users.last_updated_at as user_last_updated_at,
	users.archived_at as user_archived_at,
	household_user_memberships.id as membership_id,
	household_user_memberships.belongs_to_household as membership_belongs_to_household,
	household_user_memberships.belongs_to_user as membership_belongs_to_user,
	household_user_memberships.default_household as membership_default_household,
	household_user_memberships.household_role as membership_household_role,
	household_user_memberships.created_at as membership_created_at,
	household_user_memberships.last_updated_at as membership_last_updated_at,
	household_user_memberships.archived_at as membership_archived_at
FROM households
	JOIN household_user_memberships ON household_user_memberships.belongs_to_household = households.id
	JOIN users ON household_user_memberships.belongs_to_user = users.id
WHERE households.archived_at IS NULL
	AND household_user_memberships.archived_at IS NULL
	AND households.id = $1
`

type GetHouseholdByIDWithMembershipsRow struct {
	MembershipCreatedAt               time.Time
	CreatedAt                         time.Time
	UserCreatedAt                     time.Time
	UserLastAcceptedPrivacyPolicy     sql.NullTime
	UserEmailAddressVerifiedAt        sql.NullTime
	UserLastUpdatedAt                 sql.NullTime
	MembershipLastUpdatedAt           sql.NullTime
	UserLastIndexedAt                 sql.NullTime
	MembershipArchivedAt              sql.NullTime
	UserLastAcceptedTermsOfService    sql.NullTime
	UserArchivedAt                    sql.NullTime
	UserBirthday                      sql.NullTime
	UserTwoFactorSecretVerifiedAt     sql.NullTime
	UserPasswordLastChangedAt         sql.NullTime
	ArchivedAt                        sql.NullTime
	LastUpdatedAt                     sql.NullTime
	LastPaymentProviderSyncOccurredAt sql.NullTime
	Country                           string
	UserServiceRole                   string
	Name                              string
	BillingStatus                     string
	UserID                            string
	UserUsername                      string
	ContactPhone                      string
	UserEmailAddress                  string
	UserHashedPassword                string
	ID                                string
	MembershipHouseholdRole           string
	UserTwoFactorSecret               string
	ZipCode                           string
	MembershipBelongsToHousehold      string
	UserUserAccountStatus             string
	UserUserAccountStatusExplanation  string
	State                             string
	WebhookHmacSecret                 string
	City                              string
	UserFirstName                     string
	UserLastName                      string
	AddressLine2                      string
	AddressLine1                      string
	TimeZone                          TimeZone
	BelongsToUser                     string
	MembershipBelongsToUser           string
	PaymentProcessorCustomerID        string
	MembershipID                      string
	UserEmailAddressVerificationToken sql.NullString
	SubscriptionPlanID                sql.NullString
	UserAvatarSrc                     sql.NullString
	Latitude                          sql.NullString
	Longitude                         sql.NullString
	MembershipDefaultHousehold        bool
	UserRequiresPasswordChange        bool
}

func (q *Queries) GetHouseholdByIDWithMemberships(ctx context.Context, db DBTX, id string) ([]*GetHouseholdByIDWithMembershipsRow, error) {
	rows, err := db.QueryContext(ctx, getHouseholdByIDWithMemberships, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*GetHouseholdByIDWithMembershipsRow{}
	for rows.Next() {
		var i GetHouseholdByIDWithMembershipsRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.BillingStatus,
			&i.ContactPhone,
			&i.PaymentProcessorCustomerID,
			&i.SubscriptionPlanID,
			&i.BelongsToUser,
			&i.TimeZone,
			&i.AddressLine1,
			&i.AddressLine2,
			&i.City,
			&i.State,
			&i.ZipCode,
			&i.Country,
			&i.Latitude,
			&i.Longitude,
			&i.LastPaymentProviderSyncOccurredAt,
			&i.WebhookHmacSecret,
			&i.CreatedAt,
			&i.LastUpdatedAt,
			&i.ArchivedAt,
			&i.UserID,
			&i.UserUsername,
			&i.UserAvatarSrc,
			&i.UserEmailAddress,
			&i.UserHashedPassword,
			&i.UserPasswordLastChangedAt,
			&i.UserRequiresPasswordChange,
			&i.UserTwoFactorSecret,
			&i.UserTwoFactorSecretVerifiedAt,
			&i.UserServiceRole,
			&i.UserUserAccountStatus,
			&i.UserUserAccountStatusExplanation,
			&i.UserBirthday,
			&i.UserEmailAddressVerificationToken,
			&i.UserEmailAddressVerifiedAt,
			&i.UserFirstName,
			&i.UserLastName,
			&i.UserLastAcceptedTermsOfService,
			&i.UserLastAcceptedPrivacyPolicy,
			&i.UserLastIndexedAt,
			&i.UserCreatedAt,
			&i.UserLastUpdatedAt,
			&i.UserArchivedAt,
			&i.MembershipID,
			&i.MembershipBelongsToHousehold,
			&i.MembershipBelongsToUser,
			&i.MembershipDefaultHousehold,
			&i.MembershipHouseholdRole,
			&i.MembershipCreatedAt,
			&i.MembershipLastUpdatedAt,
			&i.MembershipArchivedAt,
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

const getHouseholdsForUser = `-- name: GetHouseholdsForUser :many

SELECT
    households.id,
	households.name,
	households.billing_status,
	households.contact_phone,
	households.payment_processor_customer_id,
	households.subscription_plan_id,
	households.belongs_to_user,
	households.time_zone,
	households.address_line_1,
	households.address_line_2,
	households.city,
	households.state,
	households.zip_code,
	households.country,
	households.latitude,
	households.longitude,
	households.last_payment_provider_sync_occurred_at,
	households.webhook_hmac_secret,
	households.created_at,
	households.last_updated_at,
	households.archived_at,
    (
        SELECT
            COUNT(households.id)
        FROM
            households
            JOIN household_user_memberships ON household_user_memberships.belongs_to_household = households.id
        WHERE households.archived_at IS NULL
            AND household_user_memberships.belongs_to_user = $1
			AND households.created_at > COALESCE($2, (SELECT NOW() - '999 years'::INTERVAL))
    AND households.created_at < COALESCE($3, (SELECT NOW() + '999 years'::INTERVAL))
	AND (
		households.last_updated_at IS NULL
		OR households.last_updated_at > COALESCE($4, (SELECT NOW() - '999 years'::INTERVAL))
	)
	AND (
		households.last_updated_at IS NULL
		OR households.last_updated_at < COALESCE($5, (SELECT NOW() + '999 years'::INTERVAL))
	)
	) as filtered_count,
    (
        SELECT COUNT(households.id)
        FROM households
        WHERE households.archived_at IS NULL
    ) AS total_count
FROM households
	JOIN household_user_memberships ON household_user_memberships.belongs_to_household = households.id
    JOIN users ON household_user_memberships.belongs_to_user = users.id
WHERE households.archived_at IS NULL
    AND household_user_memberships.archived_at IS NULL
    AND household_user_memberships.belongs_to_user = $1
	AND households.created_at > COALESCE($2, (SELECT NOW() - '999 years'::INTERVAL))
    AND households.created_at < COALESCE($3, (SELECT NOW() + '999 years'::INTERVAL))
	AND (
		households.last_updated_at IS NULL
		OR households.last_updated_at > COALESCE($4, (SELECT NOW() - '999 years'::INTERVAL))
	)
	AND (
		households.last_updated_at IS NULL
		OR households.last_updated_at < COALESCE($5, (SELECT NOW() + '999 years'::INTERVAL))
	)
	LIMIT $7
	OFFSET $6
`

type GetHouseholdsForUserParams struct {
	UserID        string
	CreatedBefore sql.NullTime
	CreatedAfter  sql.NullTime
	UpdatedBefore sql.NullTime
	UpdatedAfter  sql.NullTime
	QueryOffset   sql.NullInt32
	QueryLimit    sql.NullInt32
}

type GetHouseholdsForUserRow struct {
	CreatedAt                         time.Time
	LastPaymentProviderSyncOccurredAt sql.NullTime
	LastUpdatedAt                     sql.NullTime
	ArchivedAt                        sql.NullTime
	AddressLine1                      string
	State                             string
	BelongsToUser                     string
	TimeZone                          TimeZone
	PaymentProcessorCustomerID        string
	AddressLine2                      string
	City                              string
	ID                                string
	ZipCode                           string
	Country                           string
	ContactPhone                      string
	BillingStatus                     string
	Name                              string
	WebhookHmacSecret                 string
	SubscriptionPlanID                sql.NullString
	Longitude                         sql.NullString
	Latitude                          sql.NullString
	FilteredCount                     int64
	TotalCount                        int64
}

func (q *Queries) GetHouseholdsForUser(ctx context.Context, db DBTX, arg *GetHouseholdsForUserParams) ([]*GetHouseholdsForUserRow, error) {
	rows, err := db.QueryContext(ctx, getHouseholdsForUser,
		arg.UserID,
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
	items := []*GetHouseholdsForUserRow{}
	for rows.Next() {
		var i GetHouseholdsForUserRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.BillingStatus,
			&i.ContactPhone,
			&i.PaymentProcessorCustomerID,
			&i.SubscriptionPlanID,
			&i.BelongsToUser,
			&i.TimeZone,
			&i.AddressLine1,
			&i.AddressLine2,
			&i.City,
			&i.State,
			&i.ZipCode,
			&i.Country,
			&i.Latitude,
			&i.Longitude,
			&i.LastPaymentProviderSyncOccurredAt,
			&i.WebhookHmacSecret,
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

const updateHousehold = `-- name: UpdateHousehold :execrows

UPDATE households SET
	name = $1,
	contact_phone = $2,
	address_line_1 = $3,
	address_line_2 = $4,
	city = $5,
	state = $6,
	zip_code = $7,
	country = $8,
	latitude = $9,
	longitude = $10,
	last_updated_at = NOW()
WHERE archived_at IS NULL
	AND belongs_to_user = $11
	AND id = $12
`

type UpdateHouseholdParams struct {
	Name          string
	ContactPhone  string
	AddressLine1  string
	AddressLine2  string
	City          string
	State         string
	ZipCode       string
	Country       string
	BelongsToUser string
	ID            string
	Latitude      sql.NullString
	Longitude     sql.NullString
}

func (q *Queries) UpdateHousehold(ctx context.Context, db DBTX, arg *UpdateHouseholdParams) (int64, error) {
	result, err := db.ExecContext(ctx, updateHousehold,
		arg.Name,
		arg.ContactPhone,
		arg.AddressLine1,
		arg.AddressLine2,
		arg.City,
		arg.State,
		arg.ZipCode,
		arg.Country,
		arg.Latitude,
		arg.Longitude,
		arg.BelongsToUser,
		arg.ID,
	)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

const updateHouseholdWebhookEncryptionKey = `-- name: UpdateHouseholdWebhookEncryptionKey :execrows

UPDATE households
SET
    webhook_hmac_secret = $1,
    last_updated_at = NOW()
WHERE archived_at IS NULL
    AND belongs_to_user = $2
    AND id = $3
`

type UpdateHouseholdWebhookEncryptionKeyParams struct {
	WebhookHmacSecret string
	BelongsToUser     string
	ID                string
}

func (q *Queries) UpdateHouseholdWebhookEncryptionKey(ctx context.Context, db DBTX, arg *UpdateHouseholdWebhookEncryptionKeyParams) (int64, error) {
	result, err := db.ExecContext(ctx, updateHouseholdWebhookEncryptionKey, arg.WebhookHmacSecret, arg.BelongsToUser, arg.ID)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
