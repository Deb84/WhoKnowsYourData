package domain

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

const (
	LabelAccount Label = "Account"
	LabelCompany Label = "Company"
	LabelValue   Label = "Value"
	LabelPerson  Label = "Person"
)

type Label string

var allowedLabelsMap = map[Label]Label{
	LabelAccount: LabelAccount,
	LabelCompany: LabelCompany,
	LabelValue:   LabelValue,
	LabelPerson:  LabelPerson,
}

var ErrInvalidLabel = errors.New("invalid label")

func NewLabel(label string) (Label, error) {
	if _, ok := allowedLabelsMap[Label(label)]; !ok {
		return "", fmt.Errorf("%w: %q is not an available label", ErrInvalidLabel, label)
	}
	return Label(label), nil
}

type ValueID struct {
	uuid.UUID
}

type Value struct {
	UUID   ValueID
	Value  string
	Source string
	Type   string
	Label  Label
}

func NewValue(_uuid uuid.UUID, value, source, _type, label string) (*Value, error) {
	newLabel, err := NewLabel(label)
	if err != nil {
		return nil, err
	}

	return &Value{
		UUID:   ValueID{UUID: _uuid},
		Value:  value,
		Source: source,
		Type:   _type,
		Label:  newLabel,
	}, nil
}
