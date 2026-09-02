package valueservice

import apperr "whoknowsyourdata/errors"

func UnallowedValue(err error) *apperr.PublicError {
	return apperr.Validation("unallowed_value", "unallowed value", err)
}

func UnallowedRelation(err error) *apperr.PublicError {
	return apperr.Validation("unallowed_relation", "unallowed relation", err)
}

func ValueNotFound(err error) *apperr.PublicError {
	return apperr.NotFound("value_not_found", "value not found", err)
}

func DatabaseIntegrityError(err error) *apperr.PublicError {
	return apperr.Internal("integrity_error", "data base integrity error", err)
}
