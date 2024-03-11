package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/ithaquaKr/taskManager/internal/note"
	"github.com/ithaquaKr/taskManager/internal/note/entities"
	"github.com/ithaquaKr/taskManager/pkg/utils"
	"github.com/jmoiron/sqlx"
)

// Note repository
type noteRepo struct {
	db *sqlx.DB
}

// Note repository constructor
func NewNoteRepo(db *sqlx.DB) note.Repository {
	return &noteRepo{db: db}
}

func (r *noteRepo) CreateNote(ctx context.Context, note *entities.Note) (*entities.Note, error) {
	var n entities.Note
	if err := r.db.QueryRowxContext(
		ctx,
		createNote,
		&note.ListID,
		&note.Name,
		&note.Content,
	).StructScan(&n); err != nil {
		return nil, fmt.Errorf("noteRepo.CreateNote.QueryRowxContext, Error: %w", err)
	}

	return &n, nil
}

// TODO: fix this func
func (r *noteRepo) UpdateNote(ctx context.Context, note *entities.Note) (*entities.Note, error) {
	var n entities.Note
	if err := r.db.QueryRowxContext(
		ctx,
		updateNote,
		&note.ListID,
		&note.Name,
		&note.Content,
	).StructScan(&n); err != nil {
		return nil, fmt.Errorf("noteRepo.UpdateNote.QueryRowxContext, Error: %w", err)
	}

	return &n, nil
}

func (r *noteRepo) GetNote(ctx context.Context, id uuid.UUID) (*entities.Note, error) {
	var n entities.Note
	if err := r.db.GetContext(ctx, &n, getNote, id); err != nil {
		return nil, fmt.Errorf("noteRepo.GetNote.QueryContext, Error: %w", err)
	}
	return &n, nil
}

func (r *noteRepo) DeleteNote(ctx context.Context, id uuid.UUID) error {
	result, err := r.db.ExecContext(ctx, deleteNote, id)
	if err != nil {
		return fmt.Errorf("noteRepo.DeleteNote.ExecContext, Error: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("noteRepo.DeleteNote.RowsAffected, Error: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("noteRepo.DeleteNote.RowsAffected, Error: %w", sql.ErrNoRows)
	}
	return nil
}

func (r *noteRepo) AllNotes(ctx context.Context, pq *utils.PaginationQuery) (*entities.AllNote, error) {
	var totalCount int
	if err := r.db.GetContext(ctx, &totalCount, getTotal); err != nil {
		return nil, fmt.Errorf("noteRepo.AllNotes.QueryContext, Error: %w", err)
	}
	if totalCount == 0 {
		return &entities.AllNote{
			Paginate: utils.PaginationResponse{
				TotalCount:   totalCount,
				TotalPage:    utils.GetTotalPages(totalCount, pq.GetPageSize()),
				HasMore:      utils.GetHasMore(pq.GetPageNumber(), totalCount, pq.GetPageSize()),
				NextPage:     utils.GetNextPage(pq.GetPageNumber()),
				PreviousPage: utils.GetPreviousPage(pq.GetPageNumber()),
			},
			Result: make([]*entities.Note, 0),
		}, nil
	}

	var notes []*entities.Note
	rows, err := r.db.QueryxContext(ctx, allNotes, pq.GetLimit(), pq.GetOffset())
	if err != nil {
		return nil, fmt.Errorf("noteRepo.AllNotes.QueryxContext, Error: %w", err)
	}

	for rows.Next() {
		var note entities.Note
		if err := rows.StructScan(&note); err != nil {
			return nil, fmt.Errorf("noteRepo.AllNotes.StructScan, Error: %w", err)
		}
		notes = append(notes, &note)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("noteRepo.AllNotes.Rows.Err, Error: %w", err)
	}

	return &entities.AllNote{
		Paginate: utils.PaginationResponse{
			TotalCount:   totalCount,
			TotalPage:    utils.GetTotalPages(totalCount, pq.GetPageSize()),
			HasMore:      utils.GetHasMore(pq.GetPageNumber(), totalCount, pq.GetPageSize()),
			NextPage:     utils.GetNextPage(pq.GetPageNumber()),
			PreviousPage: utils.GetPreviousPage(pq.GetPageNumber()),
		},
		Result: notes,
	}, nil
}
