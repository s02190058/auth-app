package app

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/s02190058/auth-app/internal/config"
	"github.com/s02190058/auth-app/internal/transport/http"
	"github.com/s02190058/auth-app/pkg/httpserver"
	"github.com/s02190058/auth-app/pkg/postgres"
)

func Run(cfg *config.Config) {
	postgresCfg := postgres.Config(cfg.Postgres)
	pool, err := postgres.New(&postgresCfg)
	if err != nil {
		//
	}

	_ = pool

	router := http.InitRouter()

	httpServerCfg := httpserver.Config(cfg.HTTP.Server)
	httpServer := httpserver.New(&httpServerCfg, router)

	httpServer.Start()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-quit:
		//
	case <-httpServer.Notify():
		//
	}

	if err = httpServer.Shutdown(); err != nil {
		//
	}
}
