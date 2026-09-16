package valueshandler

import apperr "whoknowsyourdata/errors"

var ErrDataInvalidUUID = apperr.PublicData{Code: "invalid_uuid", Message: "invalid uuid"}

func InvalidUUID(err error) *apperr.PublicError {
	return apperr.Validation(ErrDataInvalidUUID, err)
}
