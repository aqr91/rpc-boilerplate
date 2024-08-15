package service

import (
	"boilerplate/config"
	"boilerplate/internal/models/auth"
	"boilerplate/internal/module/storage"
	"go.uber.org/zap"
)

type Service struct {
	config  *config.Config
	log     *zap.Logger
	storage *storage.Storage
}

func NewService(config *config.Config, logger *zap.Logger, storage *storage.Storage) *Service {
	return &Service{config: config, log: logger, storage: storage}
}

func (s *Service) Login(args *auth.LoginArgs, reply *auth.LoginReply) error {
	// TODO: storage logic
	reply.Token = "some bearer token"
	reply.Message = "ok"
	return nil
}

func (s *Service) Register(args *auth.RegisterArgs, reply *auth.RegisterReply) error {
	// TODO: storage logic
	reply.Message = "ok"
	return nil
}
