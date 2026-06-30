package repository

import (
	"context"

	"github.com/alekzandr-s/dynamic-form-engine/internal/domain"
	"github.com/google/uuid"
)

type FormTemplateRepository interface {
	Create(ctx context.Context, template *domain.FormTemplateVersion)

	GetByID(ctx context.Context, id uuid.UUID) (*domain.FormTemplateVersion, error)
}
