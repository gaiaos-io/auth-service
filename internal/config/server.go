package config

import "errors"

type ServerConfig struct {
	Host         string
	Port         int
	ReadTimeout  int
	WriteTimeout int
}

func loadServerConfig(spec *specification) (ServerConfig, error) {
	cfg := ServerConfig{}

	cfg.Host = spec.ServerHost
	cfg.Port = spec.ServerPort
	cfg.ReadTimeout = spec.ServerReadTimeout
	cfg.WriteTimeout = spec.ServerWriteTimeout

	if err := cfg.validate(); err != nil {
		return cfg, err
	}

	return cfg, nil
}

func (config ServerConfig) validate() error {
	if config.Host == "" {
		return errors.New("server host may not be empty")
	}
	if config.Port < 1 || config.Port > 65535 {
		return errors.New("server port must be between 1 and 65535")
	}
	if config.ReadTimeout < 0 || config.WriteTimeout < 0 {
		return errors.New("server timeouts must be non-negative")
	}

	return nil
}
