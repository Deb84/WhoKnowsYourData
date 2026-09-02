// Package valueservice
package valueservice

import (
	"context"
	"whoknowsyourdata/domain"

	"github.com/google/uuid"
)

type ValueRepository interface {
	InitDB(ctx context.Context) error
	CreateValue(ctx context.Context, value *domain.Value) error
	DeleteValue(ctx context.Context, valueUUID uuid.UUID) error
	CreateValues(ctx context.Context, values []domain.Value) error
	GetValue(ctx context.Context, _uuid uuid.UUID) (*domain.Value, error)
	GetValuesFromLabel(ctx context.Context, label string) ([]domain.Value, error)
	CreateRelation(ctx context.Context, relation *domain.Relation, expectedRelations int) error
	CreateRelations(ctx context.Context, relations []domain.Relation, expectedRelations int) error
}

type ValueService struct {
	log        domain.Logger
	repository ValueRepository
}

func NewValueService(log domain.Logger, repository ValueRepository) *ValueService {

	return &ValueService{
		log:        log,
		repository: repository,
	}
}

func (vs *ValueService) PrepareDatabase(ctx context.Context) error {
	err := vs.repository.InitDB(ctx)
	if err != nil {
		return err
	}
	return nil
}
