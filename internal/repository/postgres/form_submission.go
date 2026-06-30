package postgres

import (
	"context"

	"github.com/alekzandr-s/dynamic-form-engine/internal/domain"
	"github.com/alekzandr-s/dynamic-form-engine/internal/repository/interfaces"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type FormSubmissionRepository struct {
	db *pgxpool.Pool
}

var _ interfaces.FormSubmissionRepository = (*FormSubmissionRepository)(nil)

func NewFormSubmissionRepository(db *pgxpool.Pool) *FormSubmissionRepository {
	return &FormSubmissionRepository{
		db: db,
	}
}

func (r *FormSubmissionRepository) Create(ctx context.Context, submission *domain.FormSubmission) error {
	query := `
	INSERT INTO form_submissions
	(
		id,
		template_version_id,
		data,
		created_at
	)
	VALUES (
	$1, $2, $3, $4
	)
	`
	_, err := r.db.Exec(
		ctx,
		query,
		&submission.ID,
		&submission.TemplateVersionID,
		&submission.Data,
		&submission.CreatedAt,
	)

	return err
}

func (r *FormSubmissionRepository) ListByTemplateVersion(ctx context.Context, versionID uuid.UUID) ([]domain.FormSubmission, error) {
	query := `
	SELECT 
		id,
		template_version_id,
		data,
		created_at
	FROM form_submissions
	WHERE 
		template_version_id = $1
	`

	rows, err := r.db.Query(
		ctx,
		query,
		versionID,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var submissions []domain.FormSubmission

	for rows.Next() {
		var submission domain.FormSubmission

		err := rows.Scan(
			&submission.ID,
			&submission.TemplateVersionID,
			&submission.Data,
			&submission.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		submissions = append(submissions, submission)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return submissions, nil
}
