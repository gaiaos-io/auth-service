package config

type Config struct {
	Env    string
	Server ServerConfig
	Crypto CryptoConfig
}

func LoadConfig(spec *specification) (*Config, error) {
	cfg := &Config{}
	var err error

	cfg.Env = string(spec.Env)

	cfg.Server, err = loadServerConfig(spec)
	if err != nil {
		return nil, err
	}

	cfg.Crypto, err = loadCryptoConfig(spec)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
