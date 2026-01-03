package config

import "errors"

type ServerConfig struct {
	Host         string
	Port         int
	ReadTimeout  int
	WriteTimeout int
}

func loadServerConfig() (ServerConfig, error) {
	cfg := ServerConfig{}
	var err error

	cfg.Host, err = GetEnvStr("SERVER_HOST")
	if err != nil {
		return cfg, err
	}

	cfg.Port, err = GetEnvInt("SERVER_PORT")
	if err != nil {
		return cfg, err
	}

	cfg.ReadTimeout, err = GetEnvInt("SERVER_READ_TIMEOUT")
	if err != nil {
		return cfg, err
	}

	cfg.WriteTimeout, err = GetEnvInt("SERVER_WRITE_TIMEOUT")
	if err != nil {
		return cfg, err
	}

	if err := cfg.validate(); err != nil {
		return cfg, err
	}

	return cfg, nil
}

func (config ServerConfig) validate() error {
	if config.Host == "" {
		return errors.New("server host may not be empty")
	}
	if config.Port <= 0 || config.Port > 65535 {
		return errors.New("server port must be between 1 and 65535")
	}
	if config.ReadTimeout <= 0 || config.WriteTimeout <= 0 {
		return errors.New("server timeouts must be be positive")
	}

	return nil
}
