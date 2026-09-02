// Package app initialize the app
package app

import (
	"context"
	"whoknowsyourdata/domain"
	"whoknowsyourdata/models"
	"whoknowsyourdata/neo4j"
	neo4jrepo "whoknowsyourdata/repository/neo4j"
	"whoknowsyourdata/server"
	"whoknowsyourdata/server/handlers"
	valueservice "whoknowsyourdata/services/value"
)

func ServerBootstrap(log domain.Logger, env models.Env) error {
	errorHandler := server.NewErrorHandler(log, env.App.TRUSTED_CTX)

	router := server.NewRouter(log, errorHandler)
	routes := server.NewRoutes()
	srv := server.New(":8601", router)

	driver, err := neo4j.NewSession(env.Neo4j)
	if err != nil {
		return err
	}

	valueRepository := neo4jrepo.NewValueRepository(log, driver)
	valueService := valueservice.NewValueService(valueRepository)

	err = valueService.PrepareDatabase(context.Background())
	if err != nil {
		return err
	}

	handler := &handlers.Handler{Log: log}
	valueHandler := handlers.NewValueHandler(handler, valueService)

	router.Post(routes.APIValue(), valueHandler.CreateValue)
	router.Post(routes.APIValues(), valueHandler.CreateValues)
	router.Post(routes.APIRelation(), valueHandler.CreateRelation)
	router.Post(routes.APIRelations(), valueHandler.CreateRelations)

	router.Get(routes.APIValuesLabel(), valueHandler.GetValuesFromLabel)
	router.Delete(routes.APIValueUUID(), valueHandler.DeleteValue)

	log.Info("HTTP Server Listening")

	if err = srv.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
