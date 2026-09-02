package main

import (
	"fmt"
	"log/slog"
	"os"
	"slices"
	"strings"
	"whoknowsyourdata/domain"
	"whoknowsyourdata/models"

	"github.com/joho/godotenv"
)

// Env fields
const (
	ENV = "ENV"

	NEO4J_HOST     = "NEO4J_HOST"
	NEO4J_PORT     = "NEO4J_PORT"
	NEO4J_USER     = "NEO4J_USER"
	NEO4J_PASSWORD = "NEO4J_PASSWORD"
)

// Possibles env values
const (
	// ENV values
	EnvProd      = "prod"
	EnvProdLocal = "local_prod"
	EnvLocal     = "local"
	EnvDev       = "dev"
)

// Env values
var envValues = []string{EnvProd, EnvProdLocal, EnvLocal, EnvDev}

// Trusted context
var trustedByENV = map[string]bool{
	EnvProd:      false,
	EnvProdLocal: true,
	EnvLocal:     true,
	EnvDev:       true,
}

const envFile = "../.env"

func EmptyFieldErr(field string) error {
	return fmt.Errorf("env: %q empty", field)
}

func ValidateNeo4jEnv(log domain.Logger, env models.Neo4jEnv) error {
	if env.HOST == "" {
		return EmptyFieldErr(NEO4J_HOST)
	}
	if env.PORT == "" {
		return EmptyFieldErr(NEO4J_PORT)
	}
	if env.USER == "" {
		return EmptyFieldErr(NEO4J_USER)
	}
	if env.PASSWORD == "" {
		return EmptyFieldErr(NEO4J_PASSWORD)
	}
	return nil
}

func ValidateAppEnv(log domain.Logger, env *models.AppEnv) error {
	if !slices.Contains(envValues, env.ENV) {
		return fmt.Errorf("env: ENV value is incorrect, correct_values=[%s]", strings.Join(envValues, `, `))
	}

	env.TRUSTED_CTX = trustedByENV[env.ENV]

	return nil
}

func ValidateEnv(log domain.Logger, env *models.Env) error {

	if err := ValidateNeo4jEnv(log, env.Neo4j); err != nil {
		return err
	}

	if err := ValidateAppEnv(log, &env.App); err != nil {
		return err
	}

	return nil
}

func GetEnv(log domain.Logger) (*models.Env, error) {
	err := godotenv.Load(envFile)
	if err != nil {
		slog.Info("Unable to load .env, assuming Docker is used")
	}

	neo4jEnv := models.Neo4jEnv{
		HOST:     os.Getenv(NEO4J_HOST),
		PORT:     os.Getenv(NEO4J_PORT),
		USER:     os.Getenv(NEO4J_USER),
		PASSWORD: os.Getenv(NEO4J_PASSWORD),
	}

	appEnv := models.AppEnv{
		ENV: os.Getenv(ENV),
	}

	env := &models.Env{
		Neo4j: neo4jEnv,
		App:   appEnv,
	}

	if err := ValidateEnv(log, env); err != nil {
		return nil, err
	}

	return env, nil
}
