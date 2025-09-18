-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: GetFeedsWithUser :many
SELECT f.*, u.*
FROM feeds f
JOIN users u ON f.user_id = u.id;

-- name: GetFeedByURL :one
SELECT * FROM feeds
WHERE url = $1;

-- name: GetFeed :one
SELECT * FROM feeds
WHERE name = $1;

-- name: DeleteAllFeed :exec
DELETE FROM feeds;