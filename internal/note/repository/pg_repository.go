package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/ithaquaKr/taskManager/internal/models"
	"github.com/ithaquaKr/taskManager/internal/note"
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

func (r *noteRepo) CreateNote(ctx context.Context, note *models.Note) (*models.Note, error) {
	var n models.Note
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
func (r *noteRepo) UpdateNote(ctx context.Context, note *models.Note) (*models.Note, error) {
	var n models.Note
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

func (r *noteRepo) GetNote(ctx context.Context, id uuid.UUID) (*models.Note, error) {
	var n models.Note
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

func (r *noteRepo) AllNotes(ctx context.Context, pq *utils.PaginationQuery) ([]*models.Note, error) {
	// Implement this
	return nil, nil
}
