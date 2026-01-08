package config

import (
	"fmt"
	"strings"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Env string

const (
	EnvDevelopment Env = "development"
	EnvStaging     Env = "staging"
	EnvProduction  Env = "production"
)

func (env *Env) Set(value string) error {
	lowerValue := strings.ToLower(value)
	switch lowerValue {
	case "development":
		*env = EnvDevelopment
	case "staging":
		*env = EnvStaging
	case "production":
		*env = EnvProduction
	default:
		return fmt.Errorf("invalid env: %s. Must be one of development, staging, production", value)
	}
	return nil
}

type specification struct {
	Env                 Env           `default:"development" required:"true"`
	ServerHost          string        `default:"0.0.0.0" required:"true" split_words:"true"`
	ServerPort          int           `default:"50051" required:"true" split_words:"true"`
	ServerReadTimeout   int           `split_words:"true"`
	ServerWriteTimeout  int           `split_words:"true"`
	JwtIssuer           string        `default:"gaiaos-auth-service" required:"true" split_words:"true"`
	JwtAudience         string        `default:"gaiaos-auth-service-clients" required:"true" split_words:"true"`
	JwtAccessTtl        time.Duration `default:"15" required:"true" split_words:"true"`
	JwtPrivateKeyPem    string        `required:"true" split_words:"true"`
	JwtPublicKeyPem     string        `required:"true" split_words:"true"`
	Argon2idMemoryMib   uint32        `default:"64" required:"true" split_words:"true"`
	Argon2idIterations  uint32        `default:"3" required:"true" split_words:"true"`
	Argon2idParallelism uint8         `default:"1" required:"true" split_words:"true"`
	Argon2idSaltLength  uint32        `default:"16" required:"true" split_words:"true"`
	Argon2idKeyLength   uint32        `default:"32" required:"true" split_words:"true"`
}

func LoadEnvVars() (*specification, error) {
	var s specification
	err := envconfig.Process("auth_service", &s)
	if err != nil {
		return nil, err
	}

	return &s, nil
}
