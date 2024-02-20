package models

import (
	"time"

	"github.com/google/uuid"
)

// List base model
type List struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"name"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
