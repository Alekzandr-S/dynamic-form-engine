package service

import (
	"context"

	"github.com/alekzandr-s/dynamic-form-engine/internal/domain"
	"github.com/alekzandr-s/dynamic-form-engine/internal/repository/interfaces"
	"github.com/google/uuid"
)

type VersionService struct {
	versionRepo interfaces.FormTemplateVersionRepository
}

func NewVersionService(repo interfaces.FormTemplateVersionRepository) *VersionService {
	return &VersionService{
		versionRepo: repo,
	}
}

func (s *VersionService) CreateDraft(ctx context.Context, version *domain.FormTemplateVersion) error {
	version.ID = uuid.New()
	version.Status = domain.DrafTemplate

	return s.versionRepo.Create(ctx, version)

}

func (s *VersionService) Publish(ctx context.Context, definitionID uuid.UUID) error {
	return s.versionRepo.Publish(ctx, definitionID)
}

func (s *VersionService) Archive(ctx context.Context, definitionID uuid.UUID) error {
	version, err := s.versionRepo.GetPublished(ctx, definitionID)

	if err != nil {
		return err
	}

	version.Status = domain.ArchivedTemplate

	return s.versionRepo.Update(ctx, version)
}
