// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: valid_instruments_search.sql

package generated

import (
	"context"
	"database/sql"
	"time"
)

const SearchForValidInstruments = `-- name: SearchForValidInstruments :exec
SELECT
	valid_instruments.id,
	valid_instruments.name,
	valid_instruments.plural_name,
	valid_instruments.description,
	valid_instruments.icon_path,
	valid_instruments.usable_for_storage,
	valid_instruments.created_at,
	valid_instruments.last_updated_at,
	valid_instruments.archived_at
FROM valid_instruments
WHERE valid_instruments.archived_at IS NULL
	AND valid_instruments.name ILIKE $1 LIMIT 50
`

type SearchForValidInstrumentsRow struct {
	CreatedAt        time.Time    `db:"created_at"`
	LastUpdatedAt    sql.NullTime `db:"last_updated_at"`
	ArchivedAt       sql.NullTime `db:"archived_at"`
	Description      string       `db:"description"`
	IconPath         string       `db:"icon_path"`
	ID               string       `db:"id"`
	Name             string       `db:"name"`
	PluralName       string       `db:"plural_name"`
	UsableForStorage bool         `db:"usable_for_storage"`
}

func (q *Queries) SearchForValidInstruments(ctx context.Context, name string) error {
	_, err := q.exec(ctx, q.searchForValidInstrumentsStmt, SearchForValidInstruments, name)
	return err
}
