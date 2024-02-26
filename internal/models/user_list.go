package models

import (
	"time"

	"github.com/google/uuid"
)

// UserList base model
type UserList struct {
	ID        uuid.UUID `json:"id" db:"id"`
	UserID    uuid.UUID `json:"user_id" db:"user_id"`
	ListID    uuid.UUID `json:"list_id" db:"list_id"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" db:"updated_at"`
}
