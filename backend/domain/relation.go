package domain

import (
	"errors"
	"fmt"
)

type RelationType string

const (
	RelationOwnsAccount RelationType = "OWNS_ACCOUNT"
	RelationOwnsCompany RelationType = "OWNS_COMPANY"
	RelationIsPartner   RelationType = "IS_PARTNER"
	RelationHasValue    RelationType = "HAS_VALUE"
)

var allowedRelationsMap = map[RelationType]struct{}{
	RelationOwnsAccount: {},
	RelationOwnsCompany: {},
	RelationIsPartner:   {},
	RelationHasValue:    {},
}

var ErrInvalidRelation = errors.New("invalid relation type")

func NewRelationType(relationType string) (RelationType, error) {
	if _, ok := allowedRelationsMap[RelationType(relationType)]; !ok {
		return "", fmt.Errorf("%w: %q is not an available relation", ErrInvalidRelation, relationType)
	}
	return RelationType(relationType), nil
}

type AllowedMap map[Label]map[RelationType]struct{}

var allowedRelationsFrom = AllowedMap{
	LabelAccount: {
		RelationHasValue: {},
	},
	LabelCompany: {
		RelationOwnsCompany: {},
		RelationIsPartner:   {},
	},
	LabelPerson: {
		RelationOwnsAccount: {},
	},
}

var allowedRelationsTo = AllowedMap{
	LabelAccount: {
		RelationOwnsAccount: {},
	},
	LabelCompany: {
		RelationOwnsCompany: {},
		RelationIsPartner:   {},
	},
	LabelValue: {
		RelationHasValue: {},
	},
}

type Relation struct {
	Relation RelationType
	From     []ValueID
	To       []ValueID
}

var ErrUnallowedRelation = errors.New("unallowed relation")

func getValueIDs(relationType RelationType, values []Value, allowedMap AllowedMap, errString string) ([]ValueID, error) {
	var valueIDs []ValueID

	for _, value := range values {
		if _, ok := allowedMap[value.Label][relationType]; !ok {
			return nil, fmt.Errorf(errString, ErrUnallowedRelation, relationType, value.Label)
		}
		valueIDs = append(valueIDs, value.UUID)
	}

	return valueIDs, nil
}

func NewRelation(relationType string, from, to []Value) (*Relation, error) {
	newRelationType, err := NewRelationType(relationType)
	if err != nil {
		return nil, err
	}

	fromValueID, err := getValueIDs(newRelationType, from, allowedRelationsFrom, "%w: %q is not a valid relation from a %q node")
	if err != nil {
		return nil, err
	}

	toValueID, err := getValueIDs(newRelationType, to, allowedRelationsTo, "%w: %q is not a valid relation to a %q node")
	if err != nil {
		return nil, err
	}

	return &Relation{
		Relation: newRelationType,
		From:     fromValueID,
		To:       toValueID,
	}, nil
}
