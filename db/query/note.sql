-- name: GetNote :one
SELECT * FROM notes
WHERE id = $1 LIMIT 1;

-- name: ListNote :many
SELECT * FROM notes
WHERE id = $1 
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: CreateNote :one
INSERT INTO notes (
    list_id,
    title,
    content,
    reminder,
    updated_at
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

