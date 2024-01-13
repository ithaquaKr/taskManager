// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: list_users.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createListUser = `-- name: CreateListUser :one
INSERT INTO list_users (
    user_id,
    list_id,
    updated_at
) VALUES (
    $1, $2, $3
) RETURNING id, user_id, list_id, created_at, updated_at
`

type CreateListUserParams struct {
	UserID    uuid.NullUUID `json:"user_id"`
	ListID    uuid.NullUUID `json:"list_id"`
	UpdatedAt sql.NullTime  `json:"updated_at"`
}

func (q *Queries) CreateListUser(ctx context.Context, arg CreateListUserParams) (ListUser, error) {
	row := q.db.QueryRowContext(ctx, createListUser, arg.UserID, arg.ListID, arg.UpdatedAt)
	var i ListUser
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ListID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getListUser = `-- name: GetListUser :one
SELECT id, user_id, list_id, created_at, updated_at FROM list_users
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetListUser(ctx context.Context, id uuid.UUID) (ListUser, error) {
	row := q.db.QueryRowContext(ctx, getListUser, id)
	var i ListUser
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ListID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listListUser = `-- name: ListListUser :many
SELECT id, name, type, created_at, updated_at FROM lists
WHERE id = $1 
ORDER BY id
LIMIT $2
OFFSET $3
`

type ListListUserParams struct {
	ID     uuid.UUID `json:"id"`
	Limit  int32     `json:"limit"`
	Offset int32     `json:"offset"`
}

func (q *Queries) ListListUser(ctx context.Context, arg ListListUserParams) ([]List, error) {
	rows, err := q.db.QueryContext(ctx, listListUser, arg.ID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []List{}
	for rows.Next() {
		var i List
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Type,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
