package neo4jrepo

import (
	"fmt"
	"slices"
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

	// r for record
	rUUID, ok1 := props[FUuid]
	rValue, ok2 := props[FValue]
	rType, ok3 := props[FType]
	rSource, ok4 := props[FSource]
	if !ok1 || !ok2 || !ok3 || !ok4 {
		return nil, fmt.Errorf("unable to get record props: a props doesn't exists UUID=%t, VALUE=%t, TYPE=%t, SOURCE=%t", ok1, ok2, ok3, ok4)
	}

	// s for string
	// Cast string type to props
	sUUID, ok1 := rUUID.(string)
	sValue, ok2 := rValue.(string)
	sType, ok3 := rType.(string)
	sSource, ok4 := rSource.(string)
	if !ok1 || !ok2 || !ok3 || !ok4 {
		return nil, fmt.Errorf("unable to cast record props to string")
	}

	// Ensure uuid cannot be invalid
	parsedUUID, err := uuid.Parse(sUUID)
	if err != nil {
		return nil, fmt.Errorf("unable to parse uuid UUID=%s ERR=%w", sUUID, err)
	}

	// Ensure label cannot be invalid

	// Check if there is label
	if len(node.Labels) == 0 {
		return nil, fmt.Errorf("node has no label")
	}

	var labelStr string

	// ignore technical labels
	for _, label := range node.Labels {
		if slices.Contains(domain.TechnicalLabels, domain.TechnicalLabel(label)) {
			continue
		}

		labelStr = label
		break
	}

	// in the case of there is only technical label, return an error
	if labelStr == "" {
		return nil, fmt.Errorf("node.label don't contains non-technical label")
	}

	return domain.NewValue(parsedUUID, sValue, sType, sSource, labelStr)
}
