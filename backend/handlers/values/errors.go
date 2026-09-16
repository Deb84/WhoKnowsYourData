package valueshandler

import apperr "whoknowsyourdata/errors"

type PublicData = apperr.PublicData

var ErrDataInvalidUUID PublicData = PublicData{Code: "invalid_uuid", Message: "invalid uuid"}

func InvalidUUID(err error) *apperr.PublicError {
	return apperr.Validation(ErrDataInvalidUUID, err)
}
