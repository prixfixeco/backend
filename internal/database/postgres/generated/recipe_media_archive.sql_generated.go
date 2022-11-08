// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: recipe_media_archive.sql

package generated

import (
	"context"
)

const ArchiveRecipeMedia = `-- name: ArchiveRecipeMedia :exec
UPDATE recipe_media SET archived_at = NOW() WHERE archived_at IS NULL AND id = $1
`

func (q *Queries) ArchiveRecipeMedia(ctx context.Context, id string) error {
	_, err := q.exec(ctx, q.archiveRecipeMediaStmt, ArchiveRecipeMedia, id)
	return err
}
