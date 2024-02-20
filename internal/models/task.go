package models

import (
	"time"

	"github.com/google/uuid"
)

// Task base model
type Task struct {
	ID          uuid.UUID `json:"id"`
	ListID      uuid.UUID `json:"list_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Tag         string    `json:"tag"`
	Priority    string    `json:"priority"`
	DueDate     time.Time `json:"due_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
