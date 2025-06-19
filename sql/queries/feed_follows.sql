-- name: CreateFeedFollow :one
WITH inserted_follow AS (
    INSERT INTO feed_follows(id, created_at, updated_at, user_id, feed_id)
    VALUES($1, $2, $3, $4, $5)
    RETURNING *
)
SELECT 
    f.*, 
    u.name AS user_name,
    fd.name AS feed_name
FROM inserted_follow f
JOIN users u ON f.user_id = u.id
JOIN feeds fd ON f.feed_id = fd.id;

-- name: GetFeedFollowsForUser :many
SELECT feed_follows.*,users.name AS user_name,feeds.name AS feed_name
FROM feed_follows
JOIN users ON feed_follows.user_id = users.id
JOIN feeds ON feed_follows.feed_id = feeds.id
WHERE feed_follows.user_id = $1;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows
WHERE feed_follows.user_id = $1
AND feed_follows.feed_id IN(
    SELECT id FROM feeds WHERE url=$2
);