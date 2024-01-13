-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (
    username,
    hash_password,
    email,
    first_name,
    last_name,
    updated_at
) VALUES (
    $1, $2, $3, $4, $5, $6 
) RETURNING *;
