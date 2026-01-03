package config

import (
	"go.uber.org/zap"
)

type Config struct {
	Env    string
	Server ServerConfig
	Crypto CryptoConfig
}

func LoadConfig(logger *zap.Logger) *Config {
	cfg := &Config{}
	var err error

	cfg.Env, err = GetEnv("ENV")
	if err != nil {
		if cfg.Env == "" {
			logger.Fatal("failed to load environment variables", zap.Error(err))
		} else {
			logger.Warn("failed to load environment variables, using defaults", zap.Error(err))
		}
	}

	cfg.Server, err = loadServerConfig()
	if err != nil {
		logger.Fatal("%v", zap.Error(err))
	}

	cfg.Crypto, err = loadCryptoConfig()
	if err != nil {
		logger.Fatal("%v", zap.Error(err))
	}

	return cfg
}
