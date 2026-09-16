package valuesrepo

import apperr "whoknowsyourdata/errors"

func DatabaseIntegrityError(err error) *apperr.PublicError {
	return apperr.Internal("integrity_error", "data base integrity error", err)
}
