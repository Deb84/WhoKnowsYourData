package repositories

import apperr "whoknowsyourdata/errors"

var ErrDataDatabaseIntegrityError = apperr.PublicData{Code: "integrity_error", Message: "data base integrity error"}

func DatabaseIntegrityError(err error) *apperr.PublicError {
	return apperr.Internal(ErrDataDatabaseIntegrityError, err)
}
