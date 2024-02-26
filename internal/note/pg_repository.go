//go:generate mockgen -source internal/note/pg_repository.go -destination internal/note/mock/pg_repository_mock.go -package mock
package note

import (
	"context"

	"github.com/google/uuid"
	"github.com/ithaquaKr/taskManager/internal/models"
)

// Note Repository
type Repository interface {
	// CreateNote create a new note
	CreateNote(ctx context.Context, note *models.Note) (*models.Note, error)
	// GetNote retrieves a note by id
	GetNote(ctx context.Context, id uuid.UUID) (*models.Note, error)
	// UpdateNote updates a note by id
	UpdateNote(ctx context.Context, note *models.Note) (*models.Note, error)
	// DeleteNote deletes a note by id
	DeleteNote(ctx context.Context, id uuid.UUID) error
	// AllNotes retrieves a list of notes
	AllNotes(ctx context.Context, offset, limit int) ([]*models.Note, error)
}
