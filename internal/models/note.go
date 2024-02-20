package models

import (
	"time"

	"github.com/google/uuid"
)

// Note base model
type Note struct {
	ID        uuid.UUID `json:"id"`
	ListID    uuid.UUID `json:"list_id"`
	Name      string    `json:"name"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
