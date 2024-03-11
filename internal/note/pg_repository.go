//go:generate mockgen -source internal/note/pg_repository.go -destination internal/note/mock/pg_repository_mock.go -package mock
package note

import (
	"context"

	"github.com/google/uuid"
	"github.com/ithaquaKr/taskManager/internal/note/entities"
	"github.com/ithaquaKr/taskManager/pkg/utils"
)

// Note Repository
type Repository interface {
	// CreateNote create a new note
	CreateNote(ctx context.Context, note *entities.Note) (*entities.Note, error)
	// GetNote retrieves a note by id
	GetNote(ctx context.Context, id uuid.UUID) (*entities.Note, error)
	// UpdateNote updates a note by id
	UpdateNote(ctx context.Context, note *entities.Note) (*entities.Note, error)
	// DeleteNote deletes a note by id
	DeleteNote(ctx context.Context, id uuid.UUID) error
	// AllNotes retrieves a list of notes
	AllNotes(ctx context.Context, pq *utils.PaginationQuery) (*entities.AllNote, error)
}
