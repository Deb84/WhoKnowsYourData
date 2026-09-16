package server

import apperr "whoknowsyourdata/errors"

type HTTPError struct {
	Err        error
	StatusCode int
}

func NewHTTPError(err error, statusCode int) *HTTPError {
	return &HTTPError{
		Err:        err,
		StatusCode: statusCode,
	}
}

func (e *HTTPError) Error() string {
	return e.Err.Error()
}

func (e *HTTPError) Unwrap() error {
	return e.Err
}

// Provide public error for internal errors
var ErrDataInternal = apperr.PublicData{Code: "server_error", Message: "internal server error"}

func Internal(err error) *apperr.PublicError {
	return apperr.Internal(ErrDataInternal, err)
}

var ErrDataInvalidJSON = apperr.PublicData{Code: "invalid_json", Message: "invalid json"}

func InvalidJSON(err error) *apperr.PublicError {
	return apperr.Validation(ErrDataInvalidJSON, err)
}
