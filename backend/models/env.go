package models

type Neo4jEnv struct {
	HOST     string
	PORT     string
	USER     string
	PASSWORD string
}

type SqliteEnv struct {
	PATH string
}

type AppEnv struct {
	ENV         string
	TRUSTED_CTX bool
	LOG_LEVEL   string
}

type WebEnv struct {
	PORT string
}

type Env struct {
	Neo4j  Neo4jEnv
	Sqlite SqliteEnv
	App    AppEnv
	Web    WebEnv
}
