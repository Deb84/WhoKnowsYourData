package valueservice

import apperr "whoknowsyourdata/errors"

// Use Errorf directly here ?
// E.G. : return apperr.Validation(fmt.Errorf("unable to get a value: %w", err))

// Values

func UnallowedValue(err error) *apperr.PublicError {
	return apperr.Validation("unallowed_value", "unallowed value", err)
}

func ValueNotFound(err error) *apperr.PublicError {
	return apperr.NotFound("value_not_found", "value not found", err)
}

func UnableToCreateValue(err error) *apperr.PublicError {
	return apperr.Internal("unable_create_value", "unable to create value", err)
}

// UnableToDeleteValue should be used if we sure that the value exists, otherwise use ValueNotFound
func UnableToDeleteValue(err error) *apperr.PublicError {
	return apperr.Internal("unable_delete_value", "the value exists, but it's unable to be deleted", err)
}

// UnableToGetValue should be used for other case than ValueNotFound
func UnableToGetValue(err error) *apperr.PublicError {
	return apperr.Internal("unable_get_value", "unable to get the value", err)
}

// Relations

func UnallowedRelation(err error) *apperr.PublicError {
	return apperr.Validation("unallowed_relation", "unallowed relation", err)
}

func UnableToCreateRelations(err error) *apperr.PublicError {
	return apperr.Validation("unable_create_relations", "unable to create the relations", err)
}

// Others

func DatabaseIntegrityError(err error) *apperr.PublicError {
	return apperr.Internal("integrity_error", "data base integrity error", err)
}
