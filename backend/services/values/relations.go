package valueservice

import (
	"context"
	"fmt"
	"whoknowsyourdata/domain"

	"github.com/google/uuid"
)

func (vs *ValueService) validateRelationValues(ctx context.Context, UUIDs []string) ([]domain.Value, error) {
	var values []domain.Value

	for _, valueUUID := range UUIDs {
		parsedUUID, err := uuid.Parse(valueUUID)
		if err != nil {
			return nil, UnallowedRelation(fmt.Errorf("unable to parse %q to an UUID: %w", valueUUID, err))
		}

		value, err := vs.GetValue(ctx, parsedUUID)
		if err != nil {
			return nil, UnallowedRelation(fmt.Errorf("unable to get value %q: %w", valueUUID, err))
		}

		values = append(values, *value)
	}

	return values, nil
}

func (vs *ValueService) NewRelation(ctx context.Context, relationType string, from, to []string) (*domain.Relation, error) {
	fromValues, err := vs.validateRelationValues(ctx, from)
	if err != nil {
		return nil, err
	}

	toValues, err := vs.validateRelationValues(ctx, to)
	if err != nil {
		return nil, err
	}

	relation, err := domain.NewRelation(relationType, fromValues, toValues)
	if err != nil {
		return nil, UnallowedRelation(fmt.Errorf("unable to create the relation: %w", err))
	}

	return relation, nil
}

func (vs *ValueService) CreateRelation(ctx context.Context, relation *domain.Relation) error {
	expectedRelations := len(relation.From) * len(relation.To)

	if err := vs.repository.CreateRelation(ctx, relation, expectedRelations); err != nil {
		return UnableToCreateRelations(fmt.Errorf("unable to create a relation: %w", err))
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
		return UnableToCreateRelations(fmt.Errorf("unable to create some relations: %w", err))
	}

	return nil
}
