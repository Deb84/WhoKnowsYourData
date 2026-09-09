package main

import (
	"testing"
	"whoknowsyourdata/models"
)

func TestValidateNeo4jEnv(t *testing.T) {
	tests := []struct {
		name    string
		env     models.Neo4jEnv
		wantErr bool
	}{
		{
			name: "valid",
			env: models.Neo4jEnv{
				HOST:     "localhost",
				PORT:     "7687",
				USER:     "neo4j",
				PASSWORD: "password",
			},
			wantErr: false,
		},
		{
			name: "missing host",
			env: models.Neo4jEnv{
				PORT:     "7687",
				USER:     "neo4j",
				PASSWORD: "password",
			},
			wantErr: true,
		},
		{
			name: "missing port",
			env: models.Neo4jEnv{
				HOST:     "localhost",
				USER:     "neo4j",
				PASSWORD: "password",
			},
			wantErr: true,
		},
		{
			name: "missing user",
			env: models.Neo4jEnv{
				HOST:     "localhost",
				PORT:     "7687",
				PASSWORD: "password",
			},
			wantErr: true,
		},
		{
			name: "missing password",
			env: models.Neo4jEnv{
				HOST: "localhost",
				PORT: "7687",
				USER: "neo4j",
			},
			wantErr: true,
		},
		{
			name:    "missing everything",
			env:     models.Neo4jEnv{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateNeo4jEnv(tt.env)

			if (err != nil) != tt.wantErr {
				t.Fatalf("ValidateNeo4jEnv() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateAppEnv(t *testing.T) {
	for _, envValue := range envValues {
		t.Run(envValue, func(t *testing.T) {
			env := &models.AppEnv{
				ENV: envValue,
			}

			err := ValidateAppEnv(env)

			if err != nil {
				t.Fatalf("ValidateAppEnv() error = %v, want nil", err)
			}

			if env.TRUSTED_CTX != trustedByENV[envValue] {
				t.Fatalf(
					"ValidateAppEnv() TRUSTED_CTX = %v, want %v",
					env.TRUSTED_CTX,
					trustedByENV[envValue],
				)
			}
		})
	}

	t.Run("invalid environment", func(t *testing.T) {
		env := &models.AppEnv{
			ENV: "invalid",
		}

		err := ValidateAppEnv(env)

		if err == nil {
			t.Fatal("ValidateAppEnv() error = nil, want error")
		}
	})
}
