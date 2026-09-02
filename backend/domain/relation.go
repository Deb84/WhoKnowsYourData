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

var allowedRelationsMap = map[RelationType]RelationType{
	RelationOwnsAccount: RelationOwnsAccount,
	RelationOwnsCompany: RelationOwnsCompany,
	RelationIsPartner:   RelationIsPartner,
	RelationHasValue:    RelationHasValue,
}

var ErrInvalidRelation = errors.New("invalid relation type")

func NewRelationType(relationType string) (RelationType, error) {
	if _, ok := allowedRelationsMap[RelationType(relationType)]; !ok {
		return "", fmt.Errorf("%w: %q is not an available relation", ErrInvalidRelation, relationType)
	}
	return RelationType(relationType), nil
}

var allowedRelationsTo = map[RelationType]map[Label]struct{}{
	RelationOwnsAccount: {
		LabelAccount: {},
	},
	RelationOwnsCompany: {
		LabelCompany: {},
	},
	RelationIsPartner: {
		LabelCompany: {},
	},
	RelationHasValue: {
		LabelValue: {},
	},
}

var allowedRelationsFrom = map[Label]map[RelationType]struct{}{ // Not implemented yet
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

type Relation struct {
	Relation RelationType
	From     []ValueID
	To       []ValueID
}

var ErrUnallowedRelation = errors.New("unallowed relation")

func NewRelation(relationType string, from, to []Value) (*Relation, error) {
	newRelationType, err := NewRelationType(relationType)
	if err != nil {
		return nil, err
	}

	var fromValueID []ValueID

	for _, valueFrom := range from {
		if _, ok := allowedRelationsFrom[valueFrom.Label][newRelationType]; !ok {
			return nil, fmt.Errorf("%w: %q is not a valid relation from a %q node", ErrUnallowedRelation, newRelationType, valueFrom.Label)
		}
		fromValueID = append(fromValueID, valueFrom.UUID)
	}

	var toValueID []ValueID

	for _, valueTo := range to {
		if _, ok := allowedRelationsTo[newRelationType][valueTo.Label]; !ok {
			return nil, fmt.Errorf("%w: %q is not a valid relation to a %q node", ErrUnallowedRelation, newRelationType, valueTo.Label)
		}
		toValueID = append(toValueID, valueTo.UUID)
	}

	return &Relation{
		Relation: newRelationType,
		From:     fromValueID,
		To:       toValueID,
	}, nil
}
