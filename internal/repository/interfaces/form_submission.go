package interfaces

import (
	"context"

	"github.com/alekzandr-s/dynamic-form-engine/internal/domain"
	"github.com/google/uuid"
)

type FormSubmissionRepository interface {
	Create(context.Context, *domain.FormSubmission) error
	ListByTemplateVersion(context.Context, uuid.UUID) ([]domain.FormSubmission, error)
}
