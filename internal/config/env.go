package config

import (
	"fmt"
	"os"
	"strconv"

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

func GetEnvStr(key string) (string, error) {
	envVar, err := GetEnv(key)
	if err != nil {
		return "", err
	}
	return envVar, nil
}

func GetEnvInt(key string) (int, error) {
	envVar, err := GetEnv(key)
	if err != nil {
		return 0, nil
	}

	intEnvVar, err := strconv.Atoi(envVar)
	if err != nil {
		return 0, fmt.Errorf("error converting %s env var to int: %v", key, err)
	}

	return intEnvVar, nil
}

func GetEnvUint8(key string) (uint8, error) {
	envVar, err := GetEnv(key)
	if err != nil {
		return 0, err
	}

	uint8EnvVar, err := strconv.ParseUint(envVar, 10, 8)
	if err != nil {
		return 0, fmt.Errorf("error converting %s env var to uint8: %v", key, err)
	}

	return uint8(uint8EnvVar), nil
}

func GetEnvUint32(key string) (uint32, error) {
	envVar, err := GetEnv(key)
	if err != nil {
		return 0, err
	}

	uint32EnvVar, err := strconv.ParseUint(envVar, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("error converting %s env var to uint32: %v", key, err)
	}

	return uint32(uint32EnvVar), err
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
