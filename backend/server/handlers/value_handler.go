package handlers

import (
	"fmt"
	"net/http"
	apperr "whoknowsyourdata/errors"
	"whoknowsyourdata/models"
	"whoknowsyourdata/server"
	valueService "whoknowsyourdata/services/values"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type ValueHandler struct {
	*Handler
	ValueService *valueService.ValueService
}

func NewValueHandler(handler *Handler, valueService *valueService.ValueService) *ValueHandler {
	return &ValueHandler{
		Handler:      handler,
		ValueService: valueService,
	}
}

// Try to create and persist new Value
func (handler *ValueHandler) CreateValue(w http.ResponseWriter, req *http.Request) error {
	var JSONValue models.JSONValueIn

	if err := handler.decodeJSON(req, &JSONValue); err != nil {
		return err
	}

	value, err := handler.NewValue(&JSONValue)
	if err != nil {
		return err
	}

	if err := handler.ValueService.CreateValue(req.Context(), value); err != nil {
		return err
	}

	JSONValueOut := handler.NewJSONValueOut(value)

	if err := handler.encodeJSON(w, &JSONValueOut); err != nil {
		return err
	}
	return nil
}

// Try to remove a Value from the database with an UUID
func (handler *ValueHandler) DeleteValue(w http.ResponseWriter, req *http.Request) error {
	valueUUID, err := uuid.Parse(chi.URLParam(req, server.ParamValueUUID))

	if err != nil {
		err := fmt.Errorf("unable to parse uuid: %w", err)
		return apperr.Validation("invalid_uuid", "invalid uuid", err)
	}

	if err := handler.ValueService.DeleteValue(req.Context(), valueUUID); err != nil {
		return err
	}
	return nil
}

// Try to create and persist several new Values
func (handler *ValueHandler) CreateValues(w http.ResponseWriter, req *http.Request) error {
	var JSONValuesIn []models.JSONValueIn

	if err := handler.decodeJSON(req, &JSONValuesIn); err != nil {
		return err
	}

	values, err := handler.NewValues(JSONValuesIn)
	if err != nil {
		return err
	}

	if err := handler.ValueService.CreateValues(req.Context(), values); err != nil {
		return err
	}

	JSONValuesOut := handler.NewJSONValuesOut(values)

	if err := handler.encodeJSON(w, &JSONValuesOut); err != nil {
		return err
	}
	return nil
}

// Try to get an existing Value from a label
func (handler *ValueHandler) GetValuesFromLabel(w http.ResponseWriter, req *http.Request) error {
	label := chi.URLParam(req, server.ParamLabel)

	values, err := handler.ValueService.GetValuesFromLabel(req.Context(), label) //TODO convert []Value to JSONValuesOut
	if err != nil {
		return err
	}

	if err := handler.encodeJSON(w, &values); err != nil {
		return err
	}
	return nil
}

// Try to create and persist a new Relation
func (handler *ValueHandler) CreateRelation(w http.ResponseWriter, r *http.Request) error {
	var JSONRelation models.JSONRelation

	if err := handler.decodeJSON(r, &JSONRelation); err != nil {
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

	if err := handler.decodeJSON(r, &JSONRelations); err != nil {
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
