package valueservice

import (
	"context"
	"fmt"
	"whoknowsyourdata/domain"

	"github.com/google/uuid"
)

func (vs *ValueService) NewValue(val, source, _type, label string) (*domain.Value, error) {
	value, err := domain.NewValue(uuid.New(), val, source, _type, label)
	if err != nil {
		return nil, UnallowedValue(err)
	}
	return value, nil
}

func (vs *ValueService) CreateValue(ctx context.Context, value *domain.Value) error {
	err := vs.repository.CreateValue(ctx, value)
	if err != nil {
		return err
	}
	return nil
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
