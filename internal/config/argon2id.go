package config

import (
	"errors"
	"fmt"
	"runtime"
)

type Argon2idConfig struct {
	MemoryMiB   uint32
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}

func loadArgon2idConfig(spec *specification) (Argon2idConfig, error) {
	cfg := Argon2idConfig{}

	cfg.MemoryMiB = spec.Argon2idMemoryMib
	cfg.Iterations = spec.Argon2idIterations
	cfg.Parallelism = spec.Argon2idParallelism
	cfg.SaltLength = spec.Argon2idSaltLength
	cfg.KeyLength = spec.Argon2idKeyLength

	if err := cfg.validate(); err != nil {
		return cfg, err
	}

	return cfg, nil
}

func (config Argon2idConfig) validate() error {
	if config.MemoryMiB < 19 || 1024 < config.MemoryMiB {
		return errors.New("argon2id memory must be 19 and 256 MiB")
	}
	if config.Iterations < 1 || 10 < config.Iterations {
		return errors.New("argon2id iterations must be between 1 and 10")
	}
	if config.Parallelism < 1 || runtime.NumCPU() < int(config.Parallelism) {
		return fmt.Errorf("argon2id parallelism must be 1 <= parallelism <= %d", runtime.NumCPU())
	}
	if config.SaltLength < 16 || 32 < config.SaltLength {
		return errors.New("argon2id salt length must be between 16 and 32 bytes")
	}
	if config.KeyLength < 16 || 64 < config.KeyLength {
		return errors.New("argon2id key length must be between 16 and 64 bytes")
	}

	return nil
}
