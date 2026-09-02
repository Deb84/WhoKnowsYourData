package neo4jrepo

import (
	"context"
	"whoknowsyourdata/domain"

	"github.com/neo4j/neo4j-go-driver/v6/neo4j"
)

const (
	// Technical labels
	IndexLabel = "IndexLabel"

	// Values fields
	FLabel  = "label"
	FProps  = "props"
	FUuid   = "uuid"
	FValue  = "value"
	FType   = "type"
	FSource = "source"

	// Relations fields
	FFrom     = "from"
	FTo       = "to"
	FRelation = "relation"

	// Others
	FProp  = "prop"
	FField = "field"

	FIndexLabel = "index_label"
)

type ValueRepository struct {
	log    domain.Logger
	driver neo4j.Driver
}

func NewValueRepository(log domain.Logger, driver neo4j.Driver) *ValueRepository {
	return &ValueRepository{
		log:    log,
		driver: driver,
	}
}

// InitDB execute a initialization query for setup the database
func (vr *ValueRepository) InitDB(ctx context.Context) error {
	query := NewInitDBQuery()
	_, err := vr.executeQuery(ctx, query, nil)
	vr.log.Info("Database init query done")
	return err
}
