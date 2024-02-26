package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/ithaquaKr/taskManager/internal/models"
	"github.com/ithaquaKr/taskManager/internal/note"
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
	// Implement this
	return nil, nil
}

func (r *noteRepo) UpdateNote(ctx context.Context, note *models.Note) (*models.Note, error) {
	// Implement this
	return nil, nil
}

func (r *noteRepo) GetNote(ctx context.Context, id uuid.UUID) (*models.Note, error) {
	// Implement this
	return nil, nil
}

func (r *noteRepo) DeleteNote(ctx context.Context, id uuid.UUID) error {
	// Implement this
	return nil
}

func (r *noteRepo) AllNotes(ctx context.Context, offset, limit int) ([]*models.Note, error) {
	// Implement this
	return nil, nil
}
