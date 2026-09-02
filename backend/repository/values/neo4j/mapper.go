package neo4jrepo

import (
	"fmt"
	"whoknowsyourdata/domain"
	"whoknowsyourdata/models"

	"github.com/google/uuid"
	"github.com/neo4j/neo4j-go-driver/v6/neo4j"
)

// TODO: a better system to ignore technical label in GetValueFromRecord

func ConvertValueIDToString(t []domain.ValueID) []string {
	var newT []string
	for _, v := range t {
		newT = append(newT, v.String())
	}
	return newT
}

func NewNEO4JRelation(relation *domain.Relation) *models.NEO4JRelation {
	from := ConvertValueIDToString(relation.From)
	to := ConvertValueIDToString(relation.To)

	return &models.NEO4JRelation{
		Relation: string(relation.Relation),
		From:     from,
		To:       to,
	}
}

// GetValueFromRecord return a domain Value from a neo4j Record
func GetValueFromRecord(record *neo4j.Record) (*domain.Value, error) {
	rawNode, ok := record.Get("n")
	if !ok {
		return nil, fmt.Errorf("unable to get node from record")
	}

	node, ok := rawNode.(neo4j.Node)
	if !ok {
		return nil, fmt.Errorf("unable to cast record to neo4j.Node")
	}

	props := node.Props

	RUuid, ok1 := props[FUuid]
	RValue, ok2 := props[FValue]
	RType, ok3 := props[FType]
	RSource, ok4 := props[FSource]
	if !ok1 || !ok2 || !ok3 || !ok4 {
		return nil, fmt.Errorf("unable to get record props: A props doesn't exists UUID=%t, VALUE=%t, TYPE=%t, SOURCE=%t", ok1, ok2, ok3, ok4)
	}

	// S for string
	// Cast string type to props
	SUuid, ok1 := RUuid.(string)
	SValue, ok2 := RValue.(string)
	SType, ok3 := RType.(string)
	SSource, ok4 := RSource.(string)
	if !ok1 || !ok2 || !ok3 || !ok4 {
		return nil, fmt.Errorf("unable to cast record props to string")
	}

	// Ensure uuid cannot be invalid
	parsedUUID, err := uuid.Parse(SUuid)
	if err != nil {
		return nil, fmt.Errorf("unable to parse uuid UUID=%s ERR=%w", SUuid, err)
	}

	// Ensure label cannot be invalid
	// [1] is used to ignore "IndexLabel"
	label, err := domain.NewLabel(node.Labels[1])
	if err != nil {
		return nil, err
	}

	value := &domain.Value{
		UUID:   domain.ValueID{UUID: parsedUUID},
		Value:  SValue,
		Type:   SType,
		Source: SSource,
		Label:  label,
	}

	return value, nil
}
