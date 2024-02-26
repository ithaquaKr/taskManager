package models

import (
	"time"

	"github.com/google/uuid"
)

// Note base model
type Note struct {
	ID        uuid.UUID `json:"id" db:"id" validate:"omitempty,uuid"`
	ListID    uuid.UUID `json:"list_id" db:"list_id" validate:"uuid, required"`
	Name      string    `json:"name" db:"name" validate:"required"`
	Content   string    `json:"content" db:"content"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
