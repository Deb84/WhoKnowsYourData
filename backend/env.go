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

	WEBUI_PORT = "WEBUI_PORT"
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

func emptyFieldErr(field string) error {
	return fmt.Errorf("env: %q empty", field)
}

func ValidateNeo4jEnv(log domain.Logger, env models.Neo4jEnv) error {
	if env.HOST == "" {
		return emptyFieldErr(NEO4J_HOST)
	}
	if env.PORT == "" {
		return emptyFieldErr(NEO4J_PORT)
	}
	if env.USER == "" {
		return emptyFieldErr(NEO4J_USER)
	}
	if env.PASSWORD == "" {
		return emptyFieldErr(NEO4J_PASSWORD)
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

func ValidateWebEnv(env models.WebEnv) error {
	if env.PORT == "" {
		return emptyFieldErr(WEBUI_PORT)
	}
	return nil
}

func ValidateEnv(log domain.Logger, env *models.Env) error {
	if err := ValidateNeo4jEnv(log, env.Neo4j); err != nil {
		return err
	}
	if err := ValidateAppEnv(log, &env.App); err != nil {
		return err
	}
	if err := ValidateWebEnv(env.Web); err != nil {
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

	webEnv := models.WebEnv{
		PORT: os.Getenv(WEBUI_PORT),
	}

	env := &models.Env{
		Neo4j: neo4jEnv,
		App:   appEnv,
		Web:   webEnv,
	}

	if err := ValidateEnv(log, env); err != nil {
		return nil, err
	}

	return env, nil
}
