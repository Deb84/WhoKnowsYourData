package valueshandler

import (
	"net/http"
	"whoknowsyourdata/models"
)

// Try to create and persist a new Relation
func (handler *ValueHandler) CreateRelation(w http.ResponseWriter, r *http.Request) error {
	var JSONRelation models.JSONRelation

	if err := handler.DecodeJSON(r, &JSONRelation); err != nil {
		return err
	}

	relation, err := handler.NewRelation(r.Context(), &JSONRelation)
	if err != nil {
		return err
	}

	if err = handler.ValueService.CreateRelation(r.Context(), relation); err != nil {
		return err
	}
	return nil
}

// Try to create and persist several new Relations
func (handler *ValueHandler) CreateRelations(w http.ResponseWriter, r *http.Request) error {
	var JSONRelations []models.JSONRelation

	if err := handler.DecodeJSON(r, &JSONRelations); err != nil {
		return err
	}

	relations, err := handler.NewRelations(r.Context(), JSONRelations)
	if err != nil {
		return err
	}

	if err := handler.ValueService.CreateRelations(r.Context(), relations); err != nil {
		return err
	}

	return nil
}
