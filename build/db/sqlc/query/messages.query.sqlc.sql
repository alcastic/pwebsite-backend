-- name: ListMessages :many
SELECT * FROM messages
LIMIT sqlc.arg(page_size) OFFSET sqlc.arg(page_offset);

-- name: CreateMessage :one
INSERT INTO messages (
  remote_addr, content, author_name, author_email
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetLastMessageFromRemoteAddr :one
SELECT * FROM messages
WHERE remote_addr = sqlc.arg(remote_addr)
ORDER BY created_at DESC 
LIMIT 1;