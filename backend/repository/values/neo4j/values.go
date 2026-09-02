package neo4jrepo

import (
	"context"
	"fmt"
	"whoknowsyourdata/domain"
	valuesrepo "whoknowsyourdata/repository/values"

	"github.com/google/uuid"
)

func (vr *ValueRepository) CreateValue(ctx context.Context, value *domain.Value) error {
	query := NewQuery().CreateValue(value)
	_, err := vr.executeQuery(ctx, query, nil)
	return err
}

func (vr *ValueRepository) DeleteValue(ctx context.Context, valueUUID uuid.UUID) error {

	query := NewQuery().MatchFromProps(FUuid, valueUUID.String()).DeleteNode()

	_, err := vr.executeQuery(ctx, query, nil)
	return err
}

func (vr *ValueRepository) CreateValues(ctx context.Context, values []domain.Value) error {
	var queries []Query

	for _, value := range values {
		queries = append(queries, *NewQuery().CreateValue(&value))
	}

	_, err := vr.executeQueries(ctx, queries, nil)
	if err != nil {
		return err
	}
	return nil
}

func (vr *ValueRepository) GetValuesFromLabel(ctx context.Context, label string) ([]domain.Value, error) {
	query := NewQuery().MatchFromLabel(label).ReturnNode()
	result, err := vr.executeQuery(ctx, query, nil)
	if err != nil {
		return nil, err
	}

	values := []domain.Value{}

	for _, record := range result.Records {
		value, err := GetValueFromRecord(record)
		if err != nil {
			vr.log.Error(err.Error())
			continue
		}

		values = append(values, *value)
	}

	return values, nil
}

func (vr *ValueRepository) GetValue(ctx context.Context, _uuid uuid.UUID) (*domain.Value, error) {
	query := NewQuery().MatchFromProps(FUuid, _uuid.String()).ReturnNode()

	result, err := vr.executeQuery(ctx, query, nil)
	if err != nil {
		return nil, err
	}

	if len(result.Records) == 0 {
		return nil, nil

	} else if len(result.Records) > 1 {
		return nil, valuesrepo.DatabaseIntegrityError(fmt.Errorf("value appearing twice in the database, uuid: %q", _uuid.String()))
	}

	value, err := GetValueFromRecord(result.Records[0])
	if err != nil {
		return nil, err
	}

	return value, nil
}
