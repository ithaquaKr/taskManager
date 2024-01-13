// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: task.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createTask = `-- name: CreateTask :one
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
) RETURNING id, list_id, title, description, status, tag, priority, due_date, created_at, updated_at
`

type CreateTaskParams struct {
	ListID      uuid.UUID      `json:"list_id"`
	Title       string         `json:"title"`
	Description sql.NullString `json:"description"`
	Status      string         `json:"status"`
	Tag         sql.NullString `json:"tag"`
	Priority    string         `json:"priority"`
	DueDate     sql.NullTime   `json:"due_date"`
	UpdatedAt   sql.NullTime   `json:"updated_at"`
}

func (q *Queries) CreateTask(ctx context.Context, arg CreateTaskParams) (Task, error) {
	row := q.db.QueryRowContext(ctx, createTask,
		arg.ListID,
		arg.Title,
		arg.Description,
		arg.Status,
		arg.Tag,
		arg.Priority,
		arg.DueDate,
		arg.UpdatedAt,
	)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.ListID,
		&i.Title,
		&i.Description,
		&i.Status,
		&i.Tag,
		&i.Priority,
		&i.DueDate,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getTask = `-- name: GetTask :one
SELECT id, list_id, title, description, status, tag, priority, due_date, created_at, updated_at FROM tasks
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetTask(ctx context.Context, id uuid.UUID) (Task, error) {
	row := q.db.QueryRowContext(ctx, getTask, id)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.ListID,
		&i.Title,
		&i.Description,
		&i.Status,
		&i.Tag,
		&i.Priority,
		&i.DueDate,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listTask = `-- name: ListTask :many
SELECT id, list_id, title, description, status, tag, priority, due_date, created_at, updated_at FROM tasks
WHERE id = $1 
ORDER BY id
LIMIT $2
OFFSET $3
`

type ListTaskParams struct {
	ID     uuid.UUID `json:"id"`
	Limit  int32     `json:"limit"`
	Offset int32     `json:"offset"`
}

func (q *Queries) ListTask(ctx context.Context, arg ListTaskParams) ([]Task, error) {
	rows, err := q.db.QueryContext(ctx, listTask, arg.ID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Task{}
	for rows.Next() {
		var i Task
		if err := rows.Scan(
			&i.ID,
			&i.ListID,
			&i.Title,
			&i.Description,
			&i.Status,
			&i.Tag,
			&i.Priority,
			&i.DueDate,
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
