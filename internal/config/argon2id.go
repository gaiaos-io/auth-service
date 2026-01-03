package config

import "errors"

type Argon2idConfig struct {
	MemoryMB    uint32
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}

func loadArgon2idConfig() (Argon2idConfig, error) {
	cfg := Argon2idConfig{}
	var err error

	cfg.MemoryMB, err = GetEnvUint32("ARGON2ID_MEMORY_MB")
	if err != nil {
		return cfg, err
	}

	cfg.Iterations, err = GetEnvUint32("ARGON2ID_ITERATIONS")
	if err != nil {
		return cfg, err
	}

	cfg.Parallelism, err = GetEnvUint8("ARGON2ID_PARALLELISM")
	if err != nil {
		return cfg, err
	}

	cfg.SaltLength, err = GetEnvUint32("ARGON2ID_SALT_LENGTH")
	if err != nil {
		return cfg, err
	}

	cfg.KeyLength, err = GetEnvUint32("ARGON2ID_KEY_LENGTH")
	if err != nil {
		return cfg, err
	}

	err = cfg.validate()
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}

func (config Argon2idConfig) validate() error {
	if config.MemoryMB < 32 {
		return errors.New("argon2id memory must be at least 32 MB")
	}
	if config.Iterations < 2 {
		return errors.New("argon2id iterations must be at least 2")
	}
	if config.Parallelism < 1 {
		return errors.New("argon2id parallelism must be at least 1")
	}
	if config.SaltLength < 16 {
		return errors.New("argon2id salt length must be at least 16 bytes")
	}
	if config.KeyLength < 32 {
		return errors.New("argon2id key length must be at least 32 bytes")
	}

	return nil
}
