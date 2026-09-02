package valueservice

import (
	"context"
	"fmt"
	"whoknowsyourdata/domain"

	"github.com/google/uuid"
)

func (vs *ValueService) NewRelation(ctx context.Context, relationType string, from, to []string) (*domain.Relation, error) {
	var fromValues []domain.Value

	for _, fromUUID := range from {
		valueUUID, err := uuid.Parse(fromUUID)
		if err != nil {
			return nil, UnallowedRelation(fmt.Errorf("unable to parse %q to an UUID: %w", valueUUID, err))
		}

		value, err := vs.GetValue(ctx, valueUUID)
		if err != nil {
			return nil, UnallowedRelation(fmt.Errorf("unable to get value %q: %w", valueUUID, err))
		}

		fromValues = append(fromValues, *value)
	}

	var toValues []domain.Value

	for _, toUUID := range to {
		valueUUID, err := uuid.Parse(toUUID)
		if err != nil {
			return nil, UnallowedRelation(fmt.Errorf("unable to parse %q to an UUID: %w", valueUUID, err))
		}

		value, err := vs.GetValue(ctx, valueUUID)
		if err != nil {
			return nil, UnallowedRelation(fmt.Errorf("unable to get value %q: %w", valueUUID, err))
		}

		toValues = append(toValues, *value)
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
