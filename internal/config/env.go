package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

const (
	EnvDevelopment = "development"
	EnvStaging     = "staging"
	EnvProduction  = "production"
)

type EnvVar struct {
	Default  string
	Critical bool
}

var EnvVars = map[string]EnvVar{
	"ENV":            {Default: EnvDevelopment, Critical: true},
	"AUTH_GRPC_ADDR": {Default: ":50051", Critical: true},
}

func LoadEnv() (string, error) {
	env, err := GetEnv("ENV")
	if err != nil {
		return "", err
	}

	switch env {
	case EnvDevelopment, EnvStaging:
		if err := godotenv.Load(".env." + env); err != nil {
			return env, fmt.Errorf("failed to load .env file for %s: %w", env, err)
		}
	case EnvProduction:
	default:
		return "", fmt.Errorf("invalid ENV value: %s", env)
	}

	return env, nil
}

func GetEnv(key string) (string, error) {
	val := os.Getenv(key)

	if val == "" {
		envVar, exists := EnvVars[key]
		if !exists {
			return "", fmt.Errorf("unknown environment variable: %s", key)
		}

		if envVar.Critical && envVar.Default == "" {
			return "", fmt.Errorf("required environment variable missing: %s", key)
		}
		return envVar.Default, nil
	}

	return val, nil
}
