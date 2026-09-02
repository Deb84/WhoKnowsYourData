package valueservice

import (
	"context"
	"whoknowsyourdata/domain"

	"github.com/google/uuid"
)

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
