package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/ithaquaKr/taskManager/internal/list"
	"github.com/ithaquaKr/taskManager/internal/list/entities"
	"github.com/ithaquaKr/taskManager/pkg/utils"
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

func (r *listRepo) CreateList(ctx context.Context, list *entities.List) (*entities.List, error) {
	var n entities.List
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

func (r *listRepo) GetList(ctx context.Context, id uuid.UUID) (*entities.List, error) {
	var n entities.List
	if err := r.db.GetContext(ctx, &n, getList, id); err != nil {
		return nil, err
	}

	return &n, nil
}

func (r *listRepo) UpdateList(ctx context.Context, list *entities.List) (*entities.List, error) {
	var n entities.List
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

func (r *listRepo) AllLists(ctx context.Context, pq *utils.PaginationQuery) (*entities.AllList, error) {
	var totalCount int
	if err := r.db.GetContext(ctx, &totalCount, getTotal); err != nil {
		return nil, fmt.Errorf("listRepo.AllLists.GetContext, Error: %w", err)
	}

	if totalCount == 0 {
		return &entities.AllList{
			Paginate: utils.PaginationResponse{
				TotalCount:   totalCount,
				TotalPage:    utils.GetTotalPages(totalCount, pq.GetPageSize()),
				HasMore:      utils.GetHasMore(pq.GetPageNumber(), totalCount, pq.GetPageSize()),
				NextPage:     utils.GetNextPage(pq.GetPageNumber()),
				PreviousPage: utils.GetPreviousPage(pq.GetPageNumber()),
			},
			Result: make([]*entities.List, 0),
		}, nil
	}

	var lists = make([]*entities.List, 0, pq.GetPageSize())
	rows, err := r.db.QueryxContext(ctx, allLists, pq.GetLimit(), pq.GetOffset())
	if err != nil {
		return nil, fmt.Errorf("listRepo.AllLists.QueryRowxContext, Error: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var list entities.List
		if err := rows.StructScan(&list); err != nil {
			return nil, fmt.Errorf("listRepo.AllLists.StructScan, Error: %w", err)
		}
		lists = append(lists, &list)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("listRepo.AllLists.Rows.Err, Error: %w", err)
	}

	return &entities.AllList{
		Paginate: utils.PaginationResponse{
			TotalCount:   totalCount,
			TotalPage:    utils.GetTotalPages(totalCount, pq.GetPageSize()),
			HasMore:      utils.GetHasMore(pq.GetPageNumber(), totalCount, pq.GetPageSize()),
			NextPage:     utils.GetNextPage(pq.GetPageNumber()),
			PreviousPage: utils.GetPreviousPage(pq.GetPageNumber()),
		},
		Result: lists,
	}, nil
}
