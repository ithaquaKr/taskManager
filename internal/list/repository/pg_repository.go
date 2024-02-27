package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/ithaquaKr/taskManager/internal/list"
	"github.com/ithaquaKr/taskManager/internal/models"
	"github.com/jmoiron/sqlx"
)

// List Repository
type listRepo struct {
	db *sqlx.DB
}

// List Repository constructor
func NewListRepo(db *sqlx.DB) list.Repository {
	return &listRepo{db: db}
}

func (r *listRepo) CreateList(ctx context.Context, list *models.List) (*models.List, error) {
	var n models.List
	if err := r.db.QueryRowxContext(
		ctx,
		createList,
		&list.Title,
		&list.Type,
	).StructScan(&n); err != nil {
		return nil, fmt.Errorf("listRepo.CreateList.QueryRowxContext, Error: %w", err)
	}

	return &n, nil
}

func (r *listRepo) GetList(ctx context.Context, id uuid.UUID) (*models.List, error) {
	var n models.List
	if err := r.db.GetContext(ctx, &n, getList, id); err != nil {
		return nil, err
	}

	return &n, nil
}

func (r *listRepo) UpdateList(ctx context.Context, list *models.List) (*models.List, error) {
	var n models.List
	if err := r.db.QueryRowxContext(
		ctx,
		updateList,
		&list.Title,
		&list.Type,
		&list.ID,
	).StructScan(&n); err != nil {
		return nil, fmt.Errorf("listRepo.UpdateList.QueryRowxContext, Error: %w", err)
	}
	return &n, nil
}

func (r *listRepo) DeleteList(ctx context.Context, id uuid.UUID) error {
	result, err := r.db.ExecContext(ctx, deleteList, id)
	if err != nil {
		return fmt.Errorf("listRepo.DeleteList.ExecContext, Error: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("listRepo.DeleteList.RowsAffected, Error: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("listRepo.DeleteList.RowsAffected, Error: %w", sql.ErrNoRows)
	}

	return nil
}

func (r *listRepo) AllLists(ctx context.Context, offset, limit int) ([]*models.List, error) {
	// Implement this
	return nil, nil
}
