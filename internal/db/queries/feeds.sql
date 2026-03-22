-- name: CreateFeed :one
INSERT INTO feeds (created_at, updated_at, name, url)
VALUES (
    ?,
    ?,
    ?,
    ?
)
RETURNING *;

-- name: GetNextFeedToFetch :one
SELECT *
FROM feeds
ORDER BY last_fetched_at NULLS FIRST;

-- name: MarkFeedFetched :exec
UPDATE feeds
SET last_fetched_at = unixepoch(), updated_at = unixepoch()
WHERE id = ?;

-- name: GetFeeds :many
SELECT *
FROM feeds
ORDER BY id
LIMIT ?;

