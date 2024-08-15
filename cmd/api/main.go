package main

import (
	"boilerplate/config"
	"boilerplate/internal/infrastructure/logger"
	"boilerplate/run"
)

func main() {
	cfg := config.LoadConfig()
	log := logger.NewLogger()
	app := run.NewApp(log, cfg)
	app.Run()
}
