package postgres

import (
	"context"

	"github.com/alekzandr-s/dynamic-form-engine/internal/domain"
	"github.com/alekzandr-s/dynamic-form-engine/internal/repository/interfaces"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

var _ interfaces.FormTemplateVersionRepository = (*FormTemplateVersionRepository)(nil)

type FormTemplateVersionRepository struct {
	db *pgxpool.Pool
}

func NewFormTemplateVersionRepository(db *pgxpool.Pool) *FormTemplateVersionRepository {
	return &FormTemplateVersionRepository{
		db: db,
	}
}

func (r *FormTemplateVersionRepository) Create(ctx context.Context, version *domain.FormTemplateVersion) error {
	query := `INSERT INTO form_template_versions
		(
			id,
			definition_id,
			version,
			status,
			ui_schema,
			validation_schema,
			created_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		`
	_, err := r.db.Exec(
		ctx,
		query,
		version.ID,
		version.DefinitionID,
		version.Version,
		version.Status,
		version.UISchema,
		version.ValidationSchema,
		version.CreatedAt,
	)

	return err
}

func (r *FormTemplateVersionRepository) GetPublished(ctx context.Context, definitionID uuid.UUID) (*domain.FormTemplateVersion, error) {
	query := `
	SELECT
		id,
		definition_id,
		version,
		status,
		ui_schema,
		validation_schema,
		created_at
	FROM form_template_versions
	WHERE definition_id = $1
	AND status = 'PUBLISHED'
	LIMIT 1
	`
	version := &domain.FormTemplateVersion{}

	err := r.db.QueryRow(ctx, query, definitionID).Scan(
		&version.ID,
		&version.DefinitionID,
		&version.Version,
		&version.Status,
		&version.UISchema,
		&version.ValidationSchema,
		&version.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return version, nil
}

func (r *FormTemplateVersionRepository) Update(ctx context.Context, version *domain.FormTemplateVersion) error {
	query := `
	UPDATE form_template_versions
	SET status = $1
	WHERE id = $2
	`
	_, err := r.db.Exec(
		ctx,
		query,
		version.Status,
		version.ID,
	)

	return err
}

func (r *FormTemplateVersionRepository) GetDraft(ctx context.Context, definitionID uuid.UUID) (*domain.FormTemplateVersion, error) {
	query := `
	SELECT
		id,
		definition_id,
		version,
		status,
		ui_schema,
		validation_schema,
		created_at
	FROM form_template_versions
	WHERE definition_id = $1
	AND status = 'DRAFT'
	LIMIT 1
	`
	version := &domain.FormTemplateVersion{}

	err := r.db.QueryRow(ctx, query, definitionID).Scan(
		&version.ID,
		&version.DefinitionID,
		&version.Version,
		&version.Status,
		&version.UISchema,
		&version.ValidationSchema,
		&version.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return version, nil
}

func (r *FormTemplateVersionRepository) Publish(ctx context.Context, definitionID uuid.UUID) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	//Get Draft
	var draft domain.FormTemplateVersion

	query := `
	SELECT
		id,
		definition_id,
		version,
		status,
		ui_schema,
		validation_schema,
		created_at
	FROM form_template_versions
	WHERE definition_id = $1
	AND status = $2
	LIMIT 1
	`

	err = tx.QueryRow(ctx, query, definitionID, domain.DrafTemplate).Scan(
		&draft.ID,
		&draft.DefinitionID,
		&draft.Version,
		&draft.Status,
		&draft.UISchema,
		&draft.ValidationSchema,
		&draft.CreatedAt,
	)
	if err != nil {
		return err
	}

	//Archive currentl published version (if one exists)
	updateQuery := `
	UPDATE form_template_versions
	SET status = $1
	WHERE definition_id = $2
	AND status = $3
	`
	_, err = tx.Exec(ctx, updateQuery, domain.ArchivedTemplate, definitionID, domain.PublishedTemplate)
	if err != nil {
		return err
	}

	publishQuery := `
	UPDATE form_template_versions
	SET status = $1
	WHERE id = $2
	`
	_, err = tx.Exec(ctx, publishQuery, domain.PublishedTemplate, draft.ID)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)

}
