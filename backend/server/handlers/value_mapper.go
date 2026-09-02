package handlers

import (
	"context"
	"whoknowsyourdata/domain"
	"whoknowsyourdata/models"
)

// NewValue create a new domain Value from a JSONValue
func (handler *ValueHandler) NewValue(JSONValue *models.JSONValueIn) (*domain.Value, error) {
	return handler.ValueService.NewValue(JSONValue.Value, JSONValue.Source, JSONValue.Type, JSONValue.Label)
}

func (handler *ValueHandler) NewValues(JSONValues *models.JSONValuesIn) ([]domain.Value, error) {
	var values []domain.Value

	for _, JSONValue := range *JSONValues {
		value, err := handler.NewValue(&JSONValue)
		if err != nil {
			return nil, err
		}

		values = append(values, *value)
	}

	return values, nil
}

// NewRelation create a new domain Relation from a JSONRelation
func (handler *ValueHandler) NewRelation(ctx context.Context, JSONRelation *models.JSONRelation) (*domain.Relation, error) {
	relation, err := handler.ValueService.NewRelation(ctx, JSONRelation.Relation, JSONRelation.From, JSONRelation.To)
	if err != nil {
		return nil, err
	}

	return relation, nil
}

func (handler *ValueHandler) NewRelations(ctx context.Context, JSONRelations []models.JSONRelation) ([]domain.Relation, error) {
	var relations []domain.Relation

	for _, JSONRelation := range JSONRelations {
		relation, err := handler.NewRelation(ctx, &JSONRelation)
		if err != nil {
			return nil, err
		}

		relations = append(relations, *relation)
	}

	return relations, nil
}

func (handler *ValueHandler) NewJSONValueOut(value *domain.Value) *models.JSONValueOut {
	return &models.JSONValueOut{
		UUID:   value.UUID.String(),
		Value:  value.Value,
		Source: value.Source,
		Type:   value.Type,
		Label:  string(value.Label),
	}
}

func (handler *ValueHandler) NewJSONValuesOut(values []domain.Value) []models.JSONValueOut {
	var JSONValues []models.JSONValueOut
	for _, value := range values {
		JSONValues = append(JSONValues, *handler.NewJSONValueOut(&value))
	}
	return JSONValues
}
