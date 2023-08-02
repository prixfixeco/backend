// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: archive_by_code.sql

package generated

import (
	"context"
)

const ArchiveOAuth2ClientTokenByCode = `-- name: ArchiveOAuth2ClientTokenByCode :exec

DELETE FROM oauth2_client_tokens WHERE code = $1
`

func (q *Queries) ArchiveOAuth2ClientTokenByCode(ctx context.Context, db DBTX, code string) error {
	_, err := db.ExecContext(ctx, ArchiveOAuth2ClientTokenByCode, code)
	return err
}