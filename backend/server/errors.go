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

func Internal(err error) *apperr.PublicError {
	return apperr.Internal("server_error", "internal server error", err)
}

func InvalidJSON(err error) *apperr.PublicError {
	return apperr.Validation("invalid_json", "invalid json", err)
}

func InvalidUUID(err error) *apperr.PublicError {
	return apperr.Validation("invalid_uuid", "invalid uuid", err)
}
