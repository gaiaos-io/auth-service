package config

type CryptoConfig struct {
	Argon2id Argon2idConfig
	JWT      JWTConfig
}

func loadCryptoConfig(spec *specification) (CryptoConfig, error) {
	cfg := CryptoConfig{}
	var err error

	cfg.Argon2id, err = loadArgon2idConfig(spec)
	if err != nil {
		return cfg, err
	}

	cfg.JWT, err = loadJWTConfig(spec)
	if err != nil {
		return cfg, err
	}

	return cfg, err
}
