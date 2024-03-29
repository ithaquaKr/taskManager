//go:generate mockgen -source internal/list/pg_repository.go -destination internal/list/mock/pg_repository_mock.go -package mock
package list

import (
	"context"

	"github.com/google/uuid"
	"github.com/ithaquaKr/taskManager/internal/list/entities"
	"github.com/ithaquaKr/taskManager/pkg/utils"
)

// List Repository
type Repository interface {
	// CreateList creates a new list
	CreateList(ctx context.Context, list *entities.List) (*entities.List, error)
	// GetList retrieves a list by id
	GetList(ctx context.Context, id uuid.UUID) (*entities.List, error)
	// UpdateList updates a list by id
	UpdateList(ctx context.Context, list *entities.List) (*entities.List, error)
	// DeleteList deletes a list by id
	DeleteList(ctx context.Context, id uuid.UUID) error
	// GetLists retrieves a list of lists
	AllLists(ctx context.Context, pq *utils.PaginationQuery) (*entities.AllList, error)
}
