package neo4jrepo

import (
	"context"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v6/neo4j"
)

type ExeQueryCallback func(tx neo4j.ExplicitTransaction, result neo4j.EagerResult) (bool, error)
type ExeQueriesCallback func(tx neo4j.ExplicitTransaction, results []neo4j.EagerResult) (bool, error)

// txRollback provides a better readable way to return the error and rollback the transaction
func txRollback(ctx context.Context, tx neo4j.ExplicitTransaction, err error) (*neo4j.EagerResult, error) {
	_ = tx.Rollback(ctx) // err is more important than rollback error
	return nil, err
}

func executeTransaction(
	ctx context.Context,
	tx neo4j.ExplicitTransaction,
	query *Query,
) (*neo4j.EagerResult, error) {
	transformer := neo4j.EagerResultTransformer()

	cursor, err := tx.Run(ctx, *query.Query, query.Params)
	if err != nil {
		return txRollback(ctx, tx, err)
	}
	// fetch the records
	for cursor.Next(ctx) {
		if err := transformer.Accept(cursor.Record()); err != nil {
			return txRollback(ctx, tx, err)
		}
	}
	if err = cursor.Err(); err != nil {
		return txRollback(ctx, tx, err)
	}
	keys, err := cursor.Keys()
	if err != nil {
		return txRollback(ctx, tx, err)
	}
	summary, err := cursor.Consume(ctx)
	if err != nil {
		return txRollback(ctx, tx, err)
	}

	return transformer.Complete(keys, summary)
}

func (vr *ValueRepository) executeQuery(ctx context.Context, query *Query, cb ExeQueryCallback) (*neo4j.EagerResult, error) {
	errTemplate := func(err error) error {
		return fmt.Errorf("unable to execute query %q params=%v err=%w", *query.Query, query.Params, err)
	}

	session := vr.driver.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx) //nolint:errcheck, more readable than using a lambda to handle err

	tx, err := session.BeginTransaction(ctx)
	if err != nil {
		return nil, errTemplate(err)
	}

	result, err := executeTransaction(ctx, tx, query)
	if err != nil {
		return nil, errTemplate(err)
	}

	if cb != nil {
		if ok, err := cb(tx, *result); !ok || err != nil {
			_ = tx.Rollback(ctx)
			return nil, err
		}
	}

	if err = tx.Commit(ctx); err != nil {
		return nil, err
	}

	return result, nil
}

func (vr *ValueRepository) executeQueries(ctx context.Context, queries []Query, cb ExeQueriesCallback) ([]neo4j.EagerResult, error) {
	session := vr.driver.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx) //nolint:errcheck, more readable than using a lambda to handle the error

	var results []neo4j.EagerResult

	tx, err := session.BeginTransaction(ctx)
	if err != nil {
		return nil, err
	}

	for _, query := range queries {

		result, err := executeTransaction(ctx, tx, &query)
		if err != nil {
			return nil, fmt.Errorf("unable to execute query %q params=%v err=%w", *query.Query, query.Params, err)
		}

		results = append(results, *result)
	}

	if cb != nil {
		if ok, err := cb(tx, results); !ok || err != nil {
			_ = tx.Rollback(ctx)
			return nil, err
		}
	}

	if err = tx.Commit(ctx); err != nil {
		return nil, err
	}

	return results, nil
}
