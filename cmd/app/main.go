package main

import (
	"flag"
	"log"

	"github.com/s02190058/auth-app/internal/app"
	"github.com/s02190058/auth-app/internal/config"
)

var configPath = flag.String("config-path", "./configs/main.yml", "Path to config file")

func main() {
	flag.Parse()

	cfg, err := config.New(*configPath)
	if err != nil {
		log.Fatalf("unable to read config: %v", err)
	}

	app.Run(cfg)
}
