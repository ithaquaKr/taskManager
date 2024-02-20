package models

import (
	"time"

	"github.com/google/uuid"
)

// UserList base model
type UserList struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	ListID    uuid.UUID `json:"list_id"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
