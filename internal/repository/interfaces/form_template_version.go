package interfaces

import (
	"context"

	"github.com/alekzandr-s/dynamic-form-engine/internal/domain"
	"github.com/google/uuid"
)

type FormTemplateVersionRepository interface {
	Create(context.Context, *domain.FormTemplateVersion) error
	GetPublished(context.Context, uuid.UUID) (*domain.FormTemplateVersion, error)
	GetDraft(context.Context, uuid.UUID) (*domain.FormTemplateVersion, error)
	Update(context.Context, *domain.FormTemplateVersion) error
	Publish(context.Context, uuid.UUID) error
}
