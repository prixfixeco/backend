// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: service_settings.sql

package generated

import (
	"context"
	"database/sql"
	"time"
)

const archiveServiceSetting = `-- name: ArchiveServiceSetting :execrows

UPDATE service_settings SET archived_at = NOW() WHERE id = $1
`

func (q *Queries) ArchiveServiceSetting(ctx context.Context, db DBTX, id string) (int64, error) {
	result, err := db.ExecContext(ctx, archiveServiceSetting, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

const checkServiceSettingExistence = `-- name: CheckServiceSettingExistence :one

SELECT EXISTS (
	SELECT service_settings.id
	FROM service_settings
	WHERE service_settings.archived_at IS NULL
	AND service_settings.id = $1
)
`

func (q *Queries) CheckServiceSettingExistence(ctx context.Context, db DBTX, id string) (bool, error) {
	row := db.QueryRowContext(ctx, checkServiceSettingExistence, id)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const createServiceSetting = `-- name: CreateServiceSetting :exec

INSERT INTO service_settings (
	id,
	name,
	type,
	description,
	default_value,
	enumeration,
	admins_only
) VALUES (
	$1,
	$2,
	$3,
	$4,
	$5,
	$6,
	$7
)
`

type CreateServiceSettingParams struct {
	ID           string
	Name         string
	Type         SettingType
	Description  string
	Enumeration  string
	DefaultValue sql.NullString
	AdminsOnly   bool
}

func (q *Queries) CreateServiceSetting(ctx context.Context, db DBTX, arg *CreateServiceSettingParams) error {
	_, err := db.ExecContext(ctx, createServiceSetting,
		arg.ID,
		arg.Name,
		arg.Type,
		arg.Description,
		arg.DefaultValue,
		arg.Enumeration,
		arg.AdminsOnly,
	)
	return err
}

const getServiceSetting = `-- name: GetServiceSetting :one

SELECT
	service_settings.id,
	service_settings.name,
	service_settings.type,
	service_settings.description,
	service_settings.default_value,
	service_settings.enumeration,
	service_settings.admins_only,
	service_settings.created_at,
	service_settings.last_updated_at,
	service_settings.archived_at
FROM service_settings
WHERE service_settings.archived_at IS NULL
	AND service_settings.id = $1
`

func (q *Queries) GetServiceSetting(ctx context.Context, db DBTX, id string) (*ServiceSettings, error) {
	row := db.QueryRowContext(ctx, getServiceSetting, id)
	var i ServiceSettings
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Type,
		&i.Description,
		&i.DefaultValue,
		&i.Enumeration,
		&i.AdminsOnly,
		&i.CreatedAt,
		&i.LastUpdatedAt,
		&i.ArchivedAt,
	)
	return &i, err
}

const getServiceSettings = `-- name: GetServiceSettings :many

SELECT
	service_settings.id,
	service_settings.name,
	service_settings.type,
	service_settings.description,
	service_settings.default_value,
	service_settings.enumeration,
	service_settings.admins_only,
	service_settings.created_at,
	service_settings.last_updated_at,
	service_settings.archived_at,
	(
		SELECT COUNT(service_settings.id)
		FROM service_settings
		WHERE service_settings.archived_at IS NULL
			AND service_settings.created_at > COALESCE($1, (SELECT NOW() - '999 years'::INTERVAL))
			AND service_settings.created_at < COALESCE($2, (SELECT NOW() + '999 years'::INTERVAL))
			AND (
				service_settings.last_updated_at IS NULL
				OR service_settings.last_updated_at > COALESCE($3, (SELECT NOW() - '999 years'::INTERVAL))
			)
			AND (
				service_settings.last_updated_at IS NULL
				OR service_settings.last_updated_at < COALESCE($4, (SELECT NOW() + '999 years'::INTERVAL))
			)
	) AS filtered_count,
	(
		SELECT COUNT(service_settings.id)
		FROM service_settings
		WHERE service_settings.archived_at IS NULL
	) AS total_count
FROM service_settings
WHERE service_settings.archived_at IS NULL
	AND service_settings.created_at > COALESCE($1, (SELECT NOW() - '999 years'::INTERVAL))
	AND service_settings.created_at < COALESCE($2, (SELECT NOW() + '999 years'::INTERVAL))
	AND (
		service_settings.last_updated_at IS NULL
		OR service_settings.last_updated_at > COALESCE($4, (SELECT NOW() - '999 years'::INTERVAL))
	)
	AND (
		service_settings.last_updated_at IS NULL
		OR service_settings.last_updated_at < COALESCE($3, (SELECT NOW() + '999 years'::INTERVAL))
	)
LIMIT $6
OFFSET $5
`

type GetServiceSettingsParams struct {
	CreatedAfter  sql.NullTime
	CreatedBefore sql.NullTime
	UpdatedBefore sql.NullTime
	UpdatedAfter  sql.NullTime
	QueryOffset   sql.NullInt32
	QueryLimit    sql.NullInt32
}

type GetServiceSettingsRow struct {
	CreatedAt     time.Time
	LastUpdatedAt sql.NullTime
	ArchivedAt    sql.NullTime
	ID            string
	Name          string
	Type          SettingType
	Description   string
	Enumeration   string
	DefaultValue  sql.NullString
	FilteredCount int64
	TotalCount    int64
	AdminsOnly    bool
}

func (q *Queries) GetServiceSettings(ctx context.Context, db DBTX, arg *GetServiceSettingsParams) ([]*GetServiceSettingsRow, error) {
	rows, err := db.QueryContext(ctx, getServiceSettings,
		arg.CreatedAfter,
		arg.CreatedBefore,
		arg.UpdatedBefore,
		arg.UpdatedAfter,
		arg.QueryOffset,
		arg.QueryLimit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*GetServiceSettingsRow{}
	for rows.Next() {
		var i GetServiceSettingsRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Type,
			&i.Description,
			&i.DefaultValue,
			&i.Enumeration,
			&i.AdminsOnly,
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

const searchForServiceSettings = `-- name: SearchForServiceSettings :many

SELECT
	service_settings.id,
	service_settings.name,
	service_settings.type,
	service_settings.description,
	service_settings.default_value,
	service_settings.enumeration,
	service_settings.admins_only,
	service_settings.created_at,
	service_settings.last_updated_at,
	service_settings.archived_at
FROM service_settings
WHERE service_settings.archived_at IS NULL
	AND service_settings.name ILIKE '%' || $1::text || '%'
LIMIT 50
`

func (q *Queries) SearchForServiceSettings(ctx context.Context, db DBTX, nameQuery string) ([]*ServiceSettings, error) {
	rows, err := db.QueryContext(ctx, searchForServiceSettings, nameQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*ServiceSettings{}
	for rows.Next() {
		var i ServiceSettings
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Type,
			&i.Description,
			&i.DefaultValue,
			&i.Enumeration,
			&i.AdminsOnly,
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
