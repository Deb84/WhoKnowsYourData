package server

import apperr "whoknowsyourdata/errors"

type PublicData = apperr.PublicData

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
var ErrDataInternal PublicData = PublicData{Code: "server_error", Message: "internal server error"}

func Internal(err error) *apperr.PublicError {
	return apperr.Internal(ErrDataInternal, err)
}

var ErrDataInvalidJSON PublicData = PublicData{Code: "invalid_json", Message: "invalid json"}

func InvalidJSON(err error) *apperr.PublicError {
	return apperr.Validation(ErrDataInvalidJSON, err)
}
