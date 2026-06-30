package domain

import (
	"time"

	"github.com/google/uuid"
)

type FormSubmission struct {
	ID                uuid.UUID
	TemplateVersionID uuid.UUID
	Data              map[string]any
	CreatedAt         time.Time
}
