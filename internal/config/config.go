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
		User         string        `env-required:"true" yaml:"user" env:"USER"`
		Password     string        `env-required:"true" yaml:"password" env:"PASSWORD"`
		Host         string        `env-required:"true" yaml:"host" env:"HOST"`
		Port         string        `env-required:"true" yaml:"port" env:"PORT"`
		Database     string        `env-required:"true" yaml:"database" env:"DATABASE"`
		SSLMode      string        `env-required:"true" yaml:"sslmode" env:"SSLMODE"`
		PoolMaxConns int           `env-required:"true" yaml:"pool_max_conns" env:"POOL_MAX_CONNS"`
		ConnAttempts int           `env-required:"true" yaml:"conn_attempts" env:"CONN_ATTEMPTS"`
		ConnTimeout  time.Duration `env-required:"true" yaml:"conn_timeout" env:"CONN_TIMEOUT"`
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
