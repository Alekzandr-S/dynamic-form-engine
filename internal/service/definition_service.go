package service

import (
	"context"
	"time"

	"github.com/alekzandr-s/dynamic-form-engine/internal/domain"
	"github.com/alekzandr-s/dynamic-form-engine/internal/repository/interfaces"
	"github.com/google/uuid"
)

type DefinitionService struct {
	repo interfaces.FormDefinitionRepository
}

func NewDefinitionService(repo interfaces.FormDefinitionRepository) *DefinitionService {
	return &DefinitionService{
		repo: repo,
	}
}

func (s *DefinitionService) Create(ctx context.Context, definition *domain.FormDefinition) error {
	definition.ID = uuid.New()
	definition.CreatedAt = time.Now()
	definition.UpdatedAt = time.Now()

	return s.repo.Create(ctx, definition)
}

func (s *DefinitionService) List(ctx context.Context) ([]domain.FormDefinition, error) {
	return s.repo.List(ctx)
}
