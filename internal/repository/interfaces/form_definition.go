package interfaces

import (
	"context"

	"github.com/alekzandr-s/dynamic-form-engine/internal/domain"
	"github.com/google/uuid"
)

type FormDefinitionRepository interface {
	Create(context.Context, *domain.FormDefinition) error
	GetByID(context.Context, uuid.UUID) (*domain.FormDefinition, error)
	List(ctx context.Context) ([]domain.FormDefinition, error)
}
