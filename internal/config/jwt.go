package config

import "errors"

type JWTConfig struct {
	Issuer        string
	Audience      string
	AccessTTLMin  int
	PrivateKeyPEM string
	PublicKeyPEM  string
}

func loadJWTConfig() (JWTConfig, error) {
	cfg := JWTConfig{}
	var err error

	cfg.Issuer, err = GetEnvStr("JWT_ISSUER")
	if err != nil {
		return cfg, err
	}

	cfg.Audience, err = GetEnvStr("JWT_AUDIENCE")
	if err != nil {
		return cfg, err
	}

	cfg.AccessTTLMin, err = GetEnvInt("JWT_ACCESS_TTL_MIN")
	if err != nil {
		return cfg, err
	}

	cfg.PrivateKeyPEM, err = GetEnvStr("JWT_PRIVATE_KEY_PEM")
	if err != nil {
		return cfg, err
	}

	cfg.PublicKeyPEM, err = GetEnvStr("JWT_PUBLIC_KEY_PEM")
	if err != nil {
		return cfg, err
	}

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
	if config.AccessTTLMin <= 0 {
		return errors.New("jwt access token ttl must be positive")
	}
	if config.PrivateKeyPEM == "" {
		return errors.New("jwt private key must not be empty")
	}
	if config.PublicKeyPEM == "" {
		return errors.New("jwt public key must not be empty")
	}

	return nil
}
