package valueservice

import apperr "whoknowsyourdata/errors"

// Use Errorf directly here ?
// E.G. : return apperr.Validation(fmt.Errorf("unable to get a value: %w", err))

type PublicData = apperr.PublicData

// Values

var ErrDataUnallowedValue PublicData = PublicData{Code: "unallowed_value", Message: "unallowed value"}

func UnallowedValue(err error) *apperr.PublicError {
	return apperr.Validation(ErrDataUnallowedValue, err)
}

var ErrDataValueNotFound PublicData = apperr.PublicData{Code: "value_not_found", Message: "value not found"}

func ValueNotFound(err error) *apperr.PublicError {
	return apperr.NotFound(ErrDataValueNotFound, err)
}

var ErrDataUnableToCreateValue PublicData = apperr.PublicData{Code: "unable_create_value", Message: "unable to create value"}

func UnableToCreateValue(err error) *apperr.PublicError {
	return apperr.Internal(ErrDataUnableToCreateValue, err)
}

// UnableToDeleteValue should be used if we sure that the value exists, otherwise use ValueNotFound
var ErrDataUnableToDeleteValue PublicData = apperr.PublicData{Code: "unable_delete_value", Message: "the value exists, but it's unable to be deleted"}

func UnableToDeleteValue(err error) *apperr.PublicError {
	return apperr.Internal(ErrDataUnableToDeleteValue, err)
}

// UnableToGetValue should be used for other case than ValueNotFound
var ErrDataUnableToGetValue PublicData = apperr.PublicData{Code: "unable_get_value", Message: "unable to get the value"}

func UnableToGetValue(err error) *apperr.PublicError {
	return apperr.Internal(ErrDataUnableToGetValue, err)
}

// Relations

var ErrDataUnallowedRelation PublicData = apperr.PublicData{Code: "unallowed_relation", Message: "unallowed relation"}

func UnallowedRelation(err error) *apperr.PublicError {
	return apperr.Validation(ErrDataUnallowedRelation, err)
}

var ErrDataUnableToCreateRelations PublicData = apperr.PublicData{Code: "unable_create_relations", Message: "unable to create the relations"}

func UnableToCreateRelations(err error) *apperr.PublicError {
	return apperr.Validation(ErrDataUnableToCreateRelations, err)
}
