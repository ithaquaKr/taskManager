package entities

import (
	"time"

	"github.com/google/uuid"
)

type List struct {
	ID        uuid.UUID `json:"id" db:"id" validate:"omitempty,uuid"`
	Title     string    `json:"name" db:"title" validate:"required"`
	Type      string    `json:"type" db:"type" validate:"required"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" db:"updated_at"`
}
