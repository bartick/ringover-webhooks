-- name: GetWebhookEventAll :many
SELECT *
FROM webhook_events
ORDER BY created_at
;

-- name: GetWebhookEventAllByResourceEvent :many
SELECT *
FROM webhook_events
WHERE
  resource = ? AND
  event = ?
ORDER BY created_at
;

-- name: GetWebhookEventOneById :one
SELECT *
FROM webhook_events
WHERE id = ?
LIMIT 1
;

-- name: CreateWebhookEvent :exec
INSERT INTO webhook_events (
  resource, event, timestamp, data
) VALUES (
  ?, ?, ?, ?
)
;

-- name: DeleteWebhookEvent :exec
DELETE FROM webhook_events
WHERE id = ?
;