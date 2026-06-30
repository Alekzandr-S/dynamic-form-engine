package postgres

import (
	"context"

	"github.com/alekzandr-s/dynamic-form-engine/internal/domain"
	"github.com/alekzandr-s/dynamic-form-engine/internal/repository/interfaces"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

var _ interfaces.FormDefinitionRepository = (*FormDefinitionRepository)(nil)

type FormDefinitionRepository struct {
	db *pgxpool.Pool
}

func NewFormDefinitionRepository(db *pgxpool.Pool) *FormDefinitionRepository {
	return &FormDefinitionRepository{
		db: db,
	}
}

func (r *FormDefinitionRepository) Create(ctx context.Context, definition *domain.FormDefinition) error {
	query := `
	INSERT INTO form_definitions
	(
		id,
		name,
		description,
		created_at,
		updated_at
	)
	VALUES ($1, $2, $3, $4, $5)
	`
	_, err := r.db.Exec(
		ctx,
		query,
		definition.ID,
		definition.Name,
		definition.Description,
		definition.CreatedAt,
		definition.UpdatedAt,
	)

	return err
}

func (r *FormDefinitionRepository) GetByID(ctx context.Context, definitionID uuid.UUID) (*domain.FormDefinition, error) {
	query := `
	SELECT 
		id,
		name,
		description,
		created_at,
		updated_at
	FROM form_definitions
	WHERE id = $1
	`
	definition := &domain.FormDefinition{}
	err := r.db.QueryRow(ctx, query, definitionID).Scan(
		&definition.ID,
		&definition.Name,
		&definition.Description,
		&definition.CreatedAt,
		&definition.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return definition, nil
}
