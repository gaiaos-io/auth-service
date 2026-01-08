package config

import (
	"encoding/pem"
	"errors"
	"time"
)

type JWTConfig struct {
	Issuer        string
	Audience      string
	AccessTTL     time.Duration
	PrivateKeyPEM []byte
	PublicKeyPEM  []byte
}

func loadJWTConfig(spec *specification) (JWTConfig, error) {
	cfg := JWTConfig{
		Issuer:        spec.JwtIssuer,
		Audience:      spec.JwtAudience,
		AccessTTL:     spec.JwtAccessTtl * time.Minute,
		PrivateKeyPEM: []byte(spec.JwtPrivateKeyPem),
		PublicKeyPEM:  []byte(spec.JwtPublicKeyPem),
	}

	return cfg, cfg.validate()
}

func (config JWTConfig) validate() error {
	if config.Issuer == "" {
		return errors.New("jwt issuer must not be empty")
	}
	if config.Audience == "" {
		return errors.New("jwt audience must not be empty")
	}
	if config.AccessTTL < 5 || 30 < config.AccessTTL {
		return errors.New("jwt access ttl must be between 5 and 30 minutes")
	}
	if block, _ := pem.Decode([]byte(config.PrivateKeyPEM)); block == nil {
		return errors.New("invalid jwt private key pem")
	}
	if block, _ := pem.Decode([]byte(config.PublicKeyPEM)); block == nil {
		return errors.New("invalid jwt public key pem")
	}

	return nil
}
