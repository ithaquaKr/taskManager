package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/ithaquaKr/taskManager/internal/task"
	"github.com/ithaquaKr/taskManager/internal/task/entities"
	"github.com/ithaquaKr/taskManager/pkg/utils"
	"github.com/jmoiron/sqlx"
)

// Task repository
type taskRepo struct {
	db *sqlx.DB
}

// Task repository constructor
func NewTaskRepo(db *sqlx.DB) task.TaskRepository {
	return &taskRepo{db: db}
}

func (r *taskRepo) CreateTask(ctx context.Context, task *entities.Task) (*entities.Task, error) {
	var t entities.Task
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

func (r *taskRepo) GetTask(ctx context.Context, id uuid.UUID) (*entities.Task, error) {
	var t entities.Task
	if err := r.db.GetContext(ctx, &t, getTask, id); err != nil {
		return nil, fmt.Errorf("taskRepo.GetTask.GetContext, Error: %w", err)
	}
	return &t, nil
}

func (r *taskRepo) UpdateTask(ctx context.Context, task *entities.Task) (*entities.Task, error) {
	var t entities.Task
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

func (r *taskRepo) AllTasks(ctx context.Context, pq *utils.PaginationQuery) (*entities.AllTask, error) {
	var totalCount int
	if err := r.db.GetContext(ctx, &totalCount, getTotal); err != nil {
		return nil, fmt.Errorf("taskRepo.AllTasks.GetContext, Error: %w", err)
	}
	if totalCount == 0 {
		return &entities.AllTask{
			Paginate: utils.PaginationResponse{
				TotalCount:   totalCount,
				TotalPage:    utils.GetTotalPages(totalCount, pq.GetPageSize()),
				HasMore:      utils.GetHasMore(pq.GetPageNumber(), totalCount, pq.GetPageSize()),
				NextPage:     utils.GetNextPage(pq.GetPageNumber()),
				PreviousPage: utils.GetPreviousPage(pq.GetPageNumber()),
			},
			Result: make([]*entities.Task, 0),
		}, nil
	}

	var tasks []*entities.Task
	rows, err := r.db.QueryxContext(ctx, allTasks, pq.GetLimit(), pq.GetOffset())
	if err != nil {
		return nil, fmt.Errorf("taskRepo.AllTasks.QueryxContext, Error: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var task entities.Task
		if err := rows.StructScan(&task); err != nil {
			return nil, fmt.Errorf("taskRepo.AllTasks.StructScan, Error: %w", err)
		}
		tasks = append(tasks, &task)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("taskRepo.AllTasks.Rows.Err, Error: %w", err)
	}

	return &entities.AllTask{
		Paginate: utils.PaginationResponse{
			TotalCount:   totalCount,
			TotalPage:    utils.GetTotalPages(totalCount, pq.GetPageSize()),
			HasMore:      utils.GetHasMore(pq.GetPageNumber(), totalCount, pq.GetPageSize()),
			NextPage:     utils.GetNextPage(pq.GetPageNumber()),
			PreviousPage: utils.GetPreviousPage(pq.GetPageNumber()),
		},
		Result: tasks,
	}, nil
}
