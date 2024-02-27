package task

import (
	"context"

	"github.com/google/uuid"
	"github.com/ithaquaKr/taskManager/internal/models"
)

// Task Repository
type Repository interface {
	// CreateTask create a new task
	CreateTask(ctx context.Context, task *models.Task) (*models.Task, error)
	// GetTask retrieves a task by id
	GetTask(ctx context.Context, id uuid.UUID) (*models.Task, error)
	// UpdateTask updates a task by id
	UpdateTask(ctx context.Context, task *models.Task) (*models.Task, error)
	// DeleteTask deletes a task by id
	DeleteTask(ctx context.Context, id uuid.UUID) error
	// AllTasks retrieves a list of tasks
	AllTasks(ctx context.Context, offset, limit int) ([]*models.Task, error)
}
