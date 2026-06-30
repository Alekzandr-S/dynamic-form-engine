package domain

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type FormTemplateVersion struct {
	ID               uuid.UUID       `json:"id"`
	DefinitionID     uuid.UUID       `json:"definitionId"`
	Version          int             `json:"version"`
	Status           TemplateStatus  `json:"status"`
	UISchema         json.RawMessage `json:"uiSchema"`
	ValidationSchema json.RawMessage `json:"validationSchema"`
	CreatedAt        time.Time       `json:"createAt"`
}
