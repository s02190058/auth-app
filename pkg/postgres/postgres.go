package postgres

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrBadConfig    = errors.New("bad config")
	ErrBadConnetion = errors.New("can't establish connection to the postgres server")
)

type Config struct {
	User         string
	Password     string
	Host         string
	Port         string
	Database     string
	SSLMode      string
	PoolMaxConns int
	ConnAttempts int
	ConnTimeout  time.Duration
}

// TODO: pass logger
func New(cfg *Config) (*pgxpool.Pool, error) {
	url := connectionURL(cfg)

	poolCfg, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, ErrBadConfig
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), poolCfg)
	if err != nil {
		return nil, ErrBadConfig
	}

	for cfg.ConnAttempts > 0 {
		err = pool.Ping(context.Background())
		if err == nil {
			break
		}

		cfg.ConnAttempts--

		time.Sleep(cfg.ConnTimeout)
	}

	if err != nil {
		return nil, ErrBadConnetion
	}

	return pool, nil
}

func connectionURL(cfg *Config) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s&pool_max_conns=%d",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
		cfg.SSLMode,
		cfg.PoolMaxConns,
	)
}
