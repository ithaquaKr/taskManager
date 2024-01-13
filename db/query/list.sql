-- name: GetList :one
SELECT * FROM lists
WHERE id = $1 LIMIT 1;

-- name: ListList :many
SELECT * FROM lists
WHERE id = $1 
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: CreateList :one
INSERT INTO lists (
    name,
    type,
    updated_at
) VALUES (
    $1, $2, $3
) RETURNING *;
