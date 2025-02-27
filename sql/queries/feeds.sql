-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (gen_random_uuid(), NOW(), NOW(), $1, $2, $3)
RETURNING *;
-- name: ListAllFeeds :many
SELECT feeds.name as feed_name,
    feeds.url,
    users.name as user_name
FROM feeds
    JOIN users ON feeds.user_id = users.id
ORDER BY feeds.created_at ASC;
-- name: DeleteAllFeeds :exec
DELETE FROM feeds;
-- name: GetFeedByUrl :one
SELECT *
FROM feeds
WHERE url = $1;
-- name: MarkFeedFetched :one
UPDATE feeds
SET last_fetched_at = NOW(),
    updated_at = NOW()
WHERE id = $1
RETURNING *;
-- name: GetNextFeedToFetch :one
SELECT *
FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT 1;