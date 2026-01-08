package config

import (
	"encoding/pem"
	"errors"
)

type JWTConfig struct {
	Issuer        string
	Audience      string
	AccessTTLMin  int
	PrivateKeyPEM string
	PublicKeyPEM  string
}

func loadJWTConfig(spec *specification) (JWTConfig, error) {
	cfg := JWTConfig{}

	cfg.Issuer = spec.JwtIssuer
	cfg.Audience = spec.JwtAudience
	cfg.AccessTTLMin = spec.JwtAccessTtlMin
	cfg.PrivateKeyPEM = spec.JwtPrivateKeyPem
	cfg.PublicKeyPEM = spec.JwtPublicKeyPem

	if err := cfg.validate(); err != nil {
		return cfg, err
	}

	return cfg, nil
}

func (config JWTConfig) validate() error {
	if config.Issuer == "" {
		return errors.New("jwt issuer must not be empty")
	}
	if config.Audience == "" {
		return errors.New("jwt audience must not be empty")
	}
	if config.AccessTTLMin < 5 || 30 < config.AccessTTLMin {
		return errors.New("jwt access token ttl must be between 5 and 30 minutes")
	}
	if config.PrivateKeyPEM == "" {
		return errors.New("jwt private key must not be empty")
	}
	if block, _ := pem.Decode([]byte(config.PrivateKeyPEM)); block == nil {
		return errors.New("jwt private key is not valid PEM")
	}
	if config.PublicKeyPEM == "" {
		return errors.New("jwt public key must not be empty")
	}
	if block, _ := pem.Decode([]byte(config.PublicKeyPEM)); block == nil {
		return errors.New("jwt public key is not valid PEM")
	}

	return nil
}
