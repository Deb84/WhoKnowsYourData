package neo4j

import (
	"context"
	"fmt"
	"whoknowsyourdata/models"

	"github.com/neo4j/neo4j-go-driver/v6/neo4j"
)

func NewSession(env models.Neo4jEnv) (neo4j.Driver, error) {
	ctx := context.Background()
	dsn := fmt.Sprintf("neo4j://%s:%s", env.HOST, env.PORT)

	driver, err := neo4j.NewDriver(dsn, neo4j.BasicAuth(env.USER, env.PASSWORD, ""))
	if err != nil {
		return nil, fmt.Errorf("unable to create the neo4j driver HOST=%s PORT=%s USER=%s", env.HOST, env.PORT, env.USER)
	}

	err = driver.VerifyConnectivity(ctx)
	if err != nil {
		_ = driver.Close(ctx)
		return nil, fmt.Errorf("unable to verify connectivity to neo4j db HOST=%s PORT=%s USER=%s", env.HOST, env.PORT, env.USER)
	}

	return driver, nil
}
