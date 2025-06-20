-- name: CreateFeed :one
INSERT INTO feeds(id,created_at,updated_at,name,url,user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetAllFeeds :many
SELECT feeds.*,users.name AS user_name
FROM feeds
JOIN users ON feeds.user_id = users.id;

-- name: GetFeedByUrl :one
SELECT feeds.id,feeds.name
FROM feeds
WHERE url = $1;

-- name: MarkFeedFetched :exec
UPDATE feeds
SET updated_at=NOW(),last_fetched_at=NOW()
WHERE id=$1;

