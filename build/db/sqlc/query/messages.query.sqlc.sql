-- name: GetMessage :one
SELECT * FROM messages
WHERE id = $1 LIMIT 1;

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

-- name: DeleteMessage :one
DELETE FROM messages
WHERE id = $1
RETURNING *;

-- name: UpdateMessage :one
UPDATE messages SET
  author_name = $2,
  author_email = $3
WHERE id = $1 
RETURNING *;