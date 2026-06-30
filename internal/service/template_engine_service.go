package service

import (
	"context"
	"time"

	"github.com/alekzandr-s/dynamic-form-engine/internal/domain"
	"github.com/alekzandr-s/dynamic-form-engine/internal/repository/interfaces"
	"github.com/alekzandr-s/dynamic-form-engine/internal/validator"
	"github.com/google/uuid"
)

type TemplateEngineService struct {
	versionRepo    interfaces.FormTemplateVersionRepository
	submissionRepo interfaces.FormSubmissionRepository
	validator      validator.Validator
}

func NewTemplateEngineService(
	versionRepo interfaces.FormTemplateVersionRepository,
	submissionRepo interfaces.FormSubmissionRepository,
	validator validator.Validator,
) *TemplateEngineService {
	return &TemplateEngineService{
		versionRepo:    versionRepo,
		submissionRepo: submissionRepo,
		validator:      validator,
	}
}

func (s *TemplateEngineService) GetPublishedForm(ctx context.Context, definitionID uuid.UUID) (*domain.FormTemplateVersion, error) {
	return s.versionRepo.GetPublished(ctx, definitionID)
}

func (s *TemplateEngineService) Submit(ctx context.Context, definitionID uuid.UUID, data map[string]any) error {
	version, err := s.versionRepo.GetPublished(ctx, definitionID)
	if err != nil {
		return err
	}

	if err := s.validator.Validate(version.ValidationSchema, data); err != nil {
		return err
	}

	submission := &domain.FormSubmission{
		ID:                uuid.New(),
		TemplateVersionID: version.ID,
		Data:              data,
		CreatedAt:         time.Now(),
	}

	return s.submissionRepo.Create(ctx, submission)
}
