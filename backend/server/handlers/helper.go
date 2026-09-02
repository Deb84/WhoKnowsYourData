package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"whoknowsyourdata/server"
)

// Decode the JSON request body to a struct
func (handler *Handler) decodeJSON(req *http.Request, toDecode any) error {
	err := json.NewDecoder(req.Body).Decode(toDecode)
	if err != nil {
		return server.InvalidJSON(fmt.Errorf("unable to decode the json body: %w", err))
	}

	return nil
}

// Encode a struct into the request response
func (handler *Handler) encodeJSON(w http.ResponseWriter, toEncode any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(toEncode); err != nil {
		return server.Internal(fmt.Errorf("unable to encode json response"))
	}
	return nil
}
