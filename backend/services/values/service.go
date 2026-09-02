// Package valueservice
package valueservice

import (
	"context"
	"fmt"
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

func (vs *ValueService) NewValue(val, source, _type, label string) (*domain.Value, error) {
	value, err := domain.NewValue(uuid.New(), val, source, _type, label)
	if err != nil {
		return nil, UnallowedValue(err)
	}
	return value, nil
}

func (vs *ValueService) NewRelation(ctx context.Context, relationType string, from, to []string) (*domain.Relation, error) {
	var fromValues []domain.Value

	for _, fromUUID := range from {
		_uuid, err := uuid.Parse(fromUUID)
		if err != nil {
			return nil, err
		}

		value, err := vs.GetValue(ctx, _uuid)
		if err != nil {
			return nil, err
		}

		fromValues = append(fromValues, *value)
	}

	var toValues []domain.Value

	for _, toUUID := range to {
		_uuid, err := uuid.Parse(toUUID)
		if err != nil {
			return nil, err
		}

		value, err := vs.GetValue(ctx, _uuid)
		if err != nil {
			return nil, err
		}

		toValues = append(toValues, *value)
	}

	relation, err := domain.NewRelation(relationType, fromValues, toValues)
	if err != nil {
		return nil, err
	}

	return relation, nil
}

func (vs *ValueService) PrepareDatabase(ctx context.Context) error {
	err := vs.repository.InitDB(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (vs *ValueService) CreateValue(ctx context.Context, value *domain.Value) error {
	err := vs.repository.CreateValue(ctx, value)
	if err != nil {
		return err
	}
	return nil
}

func (vs *ValueService) CreateValues(ctx context.Context, values []domain.Value) error {
	err := vs.repository.CreateValues(ctx, values)
	if err != nil {
		return err
	}

	return nil
}

func (vs *ValueService) GetValuesFromLabel(ctx context.Context, label string) ([]domain.Value, error) {
	values, err := vs.repository.GetValuesFromLabel(ctx, label)
	if err != nil {
		return nil, err
	}
	return values, nil
}

func (vs *ValueService) GetValue(ctx context.Context, _uuid uuid.UUID) (*domain.Value, error) {
	value, err := vs.repository.GetValue(ctx, _uuid)
	if err != nil {
		return nil, err

	} else if value == nil {
		return nil, ValueNotFound(fmt.Errorf("unable to found a value for uuid: %q", _uuid.String()))
	}
	return value, nil
}

func (vs *ValueService) DeleteValue(ctx context.Context, valueUUID uuid.UUID) error {
	// verify if the value exists
	_, err := vs.GetValue(ctx, valueUUID)
	if err != nil {
		return err
	}

	// Delete the value
	err = vs.repository.DeleteValue(ctx, valueUUID)
	if err != nil {
		return err
	}
	return nil
}

func (vs *ValueService) CreateRelation(ctx context.Context, relation *domain.Relation) error {
	expectedRelations := len(relation.From) * len(relation.To)

	if err := vs.repository.CreateRelation(ctx, relation, expectedRelations); err != nil {
		return err
	}

	return nil
}

func (vs *ValueService) CreateRelations(ctx context.Context, relations []domain.Relation) error {
	var expectedRelations int

	for _, relation := range relations {
		expectedRelations = expectedRelations + len(relation.From)*len(relation.To)
	}

	err := vs.repository.CreateRelations(ctx, relations, expectedRelations)
	if err != nil {
		return err
	}

	return nil
}
