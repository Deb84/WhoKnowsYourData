package neo4jrepo

import (
	"context"
	"fmt"
	"whoknowsyourdata/domain"

	"github.com/neo4j/neo4j-go-driver/v6/neo4j"
)

func (vr *ValueRepository) CreateRelation(ctx context.Context, relation *domain.Relation, expectedRelations int) error {
	NEO4JRelation := NewNEO4JRelation(relation)
	query := NewQuery().CreateRelation(NEO4JRelation)

	_, err := vr.executeQuery(ctx, query, func(tx neo4j.ExplicitTransaction, result neo4j.EagerResult) (bool, error) {
		relationsCreated := result.Summary.Counters().RelationshipsCreated()
		if expectedRelations != relationsCreated {
			return false, fmt.Errorf("unable to create relation: created %d, expected %d", relationsCreated, expectedRelations)
		}
		return true, nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (vr *ValueRepository) CreateRelations(ctx context.Context, relations []domain.Relation, expectedRelations int) error {
	var queries []Query

	for _, relation := range relations {
		NEO4JRelation := NewNEO4JRelation(&relation)
		queries = append(queries, *NewQuery().CreateRelation(NEO4JRelation))
	}

	_, err := vr.executeQueries(ctx, queries, func(tx neo4j.ExplicitTransaction, results []neo4j.EagerResult) (bool, error) {
		var relationsCreated int
		for _, result := range results {
			relationsCreated += result.Summary.Counters().RelationshipsCreated()
		}

		if expectedRelations != relationsCreated {
			return false, fmt.Errorf("unable to create relation: created %d, expected %d", relationsCreated, expectedRelations)
		}

		return true, nil
	})
	if err != nil {
		return err
	}

	return nil
}
