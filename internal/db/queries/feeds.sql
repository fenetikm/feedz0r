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

