package domain

import (
	"time"

	"github.com/google/uuid"
)

type FormDefinition struct {
	ID          uuid.UUID
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
