package neo4jrepo

import (
	"fmt"
	"maps"
	"whoknowsyourdata/domain"
	"whoknowsyourdata/models"
)

type Params map[string]any

type Query struct {
	Query  *string
	Params map[string]any
}

type MatchQuery Query

var genericParams = map[string]any{
	FIndexLabel: IndexLabel,
}

func NewQuery() *Query {
	return &Query{
		Params: genericParams,
	}
}

func NewInitDBQuery() *Query {
	query := fmt.Sprintf(`
		CREATE CONSTRAINT node_uuid_unique IF NOT EXISTS
		FOR (n:%s)
		REQUIRE n.uuid IS UNIQUE
	`, IndexLabel)

	params := map[string]any{
		FIndexLabel: IndexLabel,
	}

	return &Query{
		Query:  &query,
		Params: params,
	}
}

func (q *Query) updateQuery(q1 string, p1 Params) {
	query := q1
	if q.Query != nil {
		query = fmt.Sprintf("%s %s", *q.Query, q1)
	}

	params := maps.Clone(q.Params)
	maps.Copy(params, p1)

	q.Query = &query
	q.Params = params
}

func (q *Query) CreateValue(value *domain.Value) *Query {
	query := `
		CREATE (n:$($label):$($index_label) $props)
	`
	params := map[string]any{
		FLabel:      value.Label,
		FIndexLabel: IndexLabel,
		FProps: map[string]string{
			FUuid:   value.UUID.String(),
			FValue:  value.Value,
			FType:   value.Type,
			FSource: value.Source,
		},
	}

	q.updateQuery(query, params)
	return q
}

func (q *Query) DeleteNode() *Query {
	query := `
		DETACH DELETE n
	`
	q.updateQuery(query, nil)
	return q
}

func (q *Query) CreateRelation(relation *models.NEO4JRelation) *Query {
	query := `
		MATCH (n1)
		WHERE n1[$field] IN $from

		MATCH (n2)
		WHERE n2[$field] IN $to

		CREATE (n1)-[:$($relation)]->(n2)
	`

	params := map[string]any{
		FField:    FUuid,
		FFrom:     relation.From,
		FTo:       relation.To,
		FRelation: relation.Relation,
	}

	q.updateQuery(query, params)
	return q
}

func (q *Query) DeleteRelation() *Query {
	query := `
		DETACH DELETE r
	`
	q.updateQuery(query, nil)
	return q
}

func (q *Query) MatchFromLabel(label string) *Query {
	query := `
		MATCH (n:$($label))
	`
	params := map[string]any{
		FLabel: label,
	}

	q.updateQuery(query, params)
	return q
}

func (q *Query) MatchFromProps(key string, value string) *Query {
	query := `
		MATCH (n:$($index_label) {%s: $value})
	`

	query = fmt.Sprintf(query, key)

	params := map[string]any{
		FValue: value,
	}

	q.updateQuery(query, params)
	return q
}

func (q *Query) MatchEveryNode() *Query {
	query := `
		MATCH (n:$($index_label))
	`
	q.updateQuery(query, nil)
	return q
}

func (q *Query) ReturnProp(prop string) *Query {
	query := `
		RETURN n[$prop] AS v
	`

	params := map[string]any{
		FProp: prop,
	}

	q.updateQuery(query, params)
	return q
}

func (q *Query) ReturnNode() *Query {
	query := `
		RETURN n
	`
	q.updateQuery(query, nil)
	return q
}
