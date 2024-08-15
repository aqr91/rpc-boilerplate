package components

import (
	"boilerplate/config"
	"go.uber.org/zap"
)

type Components struct {
	Config *config.Config
	Log    *zap.Logger
}

func NewComponents(config *config.Config, log *zap.Logger) *Components {
	return &Components{Config: config, Log: log}
}
