-- name: GetTask :one
SELECT * FROM tasks
WHERE id = $1 LIMIT 1;

-- name: ListTask :many
SELECT * FROM tasks
WHERE id = $1 
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: CreateTask :one
INSERT INTO tasks (
    list_id,
    title,
    description,
    status,
    tag,
    priority,
    due_date,
    updated_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING *;
