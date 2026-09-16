package valueshandler

import (
	"fmt"
	"net/http"
	"whoknowsyourdata/models"
	"whoknowsyourdata/server"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// Try to create and persist new Value
func (handler *ValueHandler) CreateValue(w http.ResponseWriter, req *http.Request) error {
	var JSONValue models.JSONValueIn

	if err := handler.DecodeJSON(req, &JSONValue); err != nil {
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

	if err := handler.EncodeJSON(w, &JSONValueOut); err != nil {
		return err
	}
	return nil
}

// Try to remove a Value from the database with an UUID
func (handler *ValueHandler) DeleteValue(w http.ResponseWriter, req *http.Request) error {
	valueUUID, err := uuid.Parse(chi.URLParam(req, server.ParamValueUUID))

	if err != nil {
		err := fmt.Errorf("unable to parse uuid: %w", err)
		return InvalidUUID(err)
	}

	if err := handler.ValueService.DeleteValue(req.Context(), valueUUID); err != nil {
		return err
	}
	return nil
}

// Try to create and persist several new Values
func (handler *ValueHandler) CreateValues(w http.ResponseWriter, req *http.Request) error {
	var JSONValuesIn []models.JSONValueIn

	if err := handler.DecodeJSON(req, &JSONValuesIn); err != nil {
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

	if err := handler.EncodeJSON(w, &JSONValuesOut); err != nil {
		return err
	}
	return nil
}

// Try to get an existing Value from its UUID
func (handler *ValueHandler) GetValue(w http.ResponseWriter, req *http.Request) error {
	valueUUID, err := uuid.Parse(chi.URLParam(req, server.ParamValueUUID))

	if err != nil {
		err := fmt.Errorf("unable to parse uuid: %w", err)
		return InvalidUUID(err)
	}

	value, err := handler.ValueService.GetValue(req.Context(), valueUUID)
	if err != nil {
		return err
	}

	JSONValueOut := handler.NewJSONValueOut(value)

	if err := handler.EncodeJSON(w, &JSONValueOut); err != nil {
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

	JSONValueOut := handler.NewJSONValuesOut(values)

	if err := handler.EncodeJSON(w, &JSONValueOut); err != nil {
		return err
	}
	return nil
}
