package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/ithaquaKr/taskManager/internal/models"
	"github.com/ithaquaKr/taskManager/internal/task"
	"github.com/ithaquaKr/taskManager/pkg/utils"
	"github.com/jmoiron/sqlx"
)

// Task repository
type taskRepo struct {
	db *sqlx.DB
}

// Task repository constructor
func NewTaskRepo(db *sqlx.DB) task.Repository {
	return &taskRepo{db: db}
}

func (r *taskRepo) CreateTask(ctx context.Context, task *models.Task) (*models.Task, error) {
	var t models.Task
	if err := r.db.QueryRowxContext(
		ctx,
		createTask,
		&task.ListID,
		&task.Name,
		&task.Description,
		&task.Status,
		&task.Tag,
		&task.Priority,
		&task.DueDate,
	).StructScan(&t); err != nil {
		return nil, fmt.Errorf("taskRepo.CreateTask.QueryRowxContext, Error: %w", err)
	}

	return &t, nil
}

func (r *taskRepo) GetTask(ctx context.Context, id uuid.UUID) (*models.Task, error) {
	var t models.Task
	if err := r.db.GetContext(ctx, &t, getTask, id); err != nil {
		return nil, fmt.Errorf("taskRepo.GetTask.GetContext, Error: %w", err)
	}
	return &t, nil
}

func (r *taskRepo) UpdateTask(ctx context.Context, task *models.Task) (*models.Task, error) {
	var t models.Task
	if err := r.db.QueryRowxContext(
		ctx,
		updateTask,
		&task.ListID,
		&task.Name,
		&task.Description,
		&task.Status,
		&task.Tag,
		&task.Priority,
		&task.DueDate,
	).StructScan(&t); err != nil {
		return nil, fmt.Errorf("taskRepo.UpdateTask.QueryRowxContext, Error: %w", err)
	}

	return &t, nil
}

func (r *taskRepo) DeleteTask(ctx context.Context, id uuid.UUID) error {
	result, err := r.db.ExecContext(ctx, deleteTask, id)
	if err != nil {
		return fmt.Errorf("taskRepo.DeleteTask.ExecContext, Error: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("taskRepo.DeleteTask.RowsAffected, Error: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("taskRepo.DeleteTask.RowsAffected, Error: %w", sql.ErrNoRows)
	}

	return nil
}

func (r *taskRepo) AllTasks(ctx context.Context, pq *utils.PaginationQuery) ([]*models.Task, error) {
	// Implement this function
	return nil, nil
}
