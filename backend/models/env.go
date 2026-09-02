package models

type Neo4jEnv struct {
	HOST     string
	PORT     string
	USER     string
	PASSWORD string
}

type AppEnv struct {
	ENV         string
	TRUSTED_CTX bool
}

type Env struct {
	Neo4j Neo4jEnv
	App   AppEnv
}
