package entities

import (
	"time"

	"github.com/google/uuid"
)

// Task base model
type Task struct {
	ID          uuid.UUID  `json:"id" db:"id" validate:"omitempty,uuid"`
	ListID      uuid.UUID  `json:"list_id" db:"list_id" validate:"uuid, required"`
	Name        string     `json:"name" db:"name" validate:"required"`
	Description *string    `json:"description" db:"description"`
	Status      string     `json:"status" db:"status" validate:"required"`
	Tag         *string    `json:"tag" db:"tag"`
	Priority    string     `json:"priority" db:"priority" validate:"required"`
	DueDate     *time.Time `json:"due_date" db:"due_date"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
}
