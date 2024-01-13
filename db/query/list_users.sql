-- name: GetListUser :one
SELECT * FROM list_users
WHERE id = $1 LIMIT 1;

-- name: ListListUser :many
SELECT * FROM lists
WHERE id = $1 
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: CreateListUser :one
INSERT INTO list_users (
    user_id,
    list_id,
    updated_at
) VALUES (
    $1, $2, $3
) RETURNING *;

