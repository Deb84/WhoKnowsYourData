package server

import (
	"encoding/json"
	"errors"
	"net/http"
	"whoknowsyourdata/domain"
	apperr "whoknowsyourdata/errors"
	"whoknowsyourdata/models"
)

type ErrorHandler struct {
	log        domain.Logger
	trustedCtx bool
}

func NewErrorHandler(log domain.Logger, trustedCtx bool) *ErrorHandler {
	return &ErrorHandler{
		log:        log,
		trustedCtx: trustedCtx,
	}
}

// Resolve HTTP Status Code with PublicError kind
func getStatusCode(kind apperr.Kind) int {
	switch kind {
	case apperr.KindNotFound:
		return http.StatusNotFound
	case apperr.KindValidation:
		return http.StatusBadRequest
	case apperr.KindConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

// Wrap an AppHandlerFunc to handle the errors
func (handler *ErrorHandler) Wrap(fn AppHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {

			publicErr, ok := errors.AsType[*apperr.PublicError](err)

			// if error is not public or the app is not in a trusted context and it's an internal error, return an basic internal error
			if !ok || (publicErr.Kind == apperr.KindInternal && !handler.trustedCtx) {
				publicErr = Internal(err)
			}

			status := getStatusCode(publicErr.Kind)

			handler.log.Error("%s", publicErr.Error())

			JSONError := models.JSONError{
				Code:    publicErr.Code,
				Message: publicErr.Msg,
			}

			w.WriteHeader(status)
			if err := json.NewEncoder(w).Encode(&JSONError); err != nil {
				handler.log.Error("error handler: unable to reply with error: %s", err.Error())
			}
		}
	}
}
