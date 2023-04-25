package config

import (
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		HTTP           HTTP           `yaml:"http" env-prefix:"HTTP_"`
		Postgres       Postgres       `yaml:"postgres" env-prefix:"PG_"`
		Redis          Redis          `yaml:"redis" env-prefix:"REDIS_"`
		Authentication Authentication `yaml:"authentication" env-prefix:"AUTH_"`
	}

	HTTP struct {
		Server Server `yaml:"server" env-prefix:"SRV_"`
	}

	Server struct {
		Port            string        `env-required:"true" yaml:"port" env:"PORT"`
		ReadTimeout     time.Duration `env-required:"true" yaml:"read_timeout" env:"READ_TIMEOUT"`
		WriteTimeout    time.Duration `env-required:"true" yaml:"write_timeout" env:"WRITE_TIMEOUT"`
		ShutdownTimeout time.Duration `env-required:"true" yaml:"shutdown_timeout" env:"SHUTDOWN_TIMEOUT"`
	}

	Postgres struct {
	}

	Redis struct {
	}

	Authentication struct {
		JWT             JWT           `yaml:"jwt" env-prefix:"JWT_"`
		RefreshTokenTTL time.Duration `env-required:"true" env:"REFRESH_TOKEN_TTL"`
	}

	JWT struct {
		Secret   string        `env-required:"true" env:"SECRET"`
		TokenTTL time.Duration `env-required:"true" yaml:"token_ttl" env:"TOKEN_TTL"`
	}
)

func New(path string) (*Config, error) {
	cfg := new(Config)
	if err := cleanenv.ReadConfig(path, cfg); err != nil {
		return nil, err
	}

	if err := cleanenv.ReadEnv(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
