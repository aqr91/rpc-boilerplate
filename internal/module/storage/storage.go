package storage

import (
	"boilerplate/config"
	"go.uber.org/zap"
)

type Storage struct {
	config *config.Config
	log    *zap.Logger
}

func NewStorage(config *config.Config, logger *zap.Logger) *Storage {
	return &Storage{config: config, log: logger}
}
