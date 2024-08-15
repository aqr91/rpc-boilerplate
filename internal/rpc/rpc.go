package rpc

import (
	"boilerplate/config"
	"boilerplate/internal/module/service"
	authRPC "boilerplate/internal/rpc/auth"
	"go.uber.org/zap"
	"net"
	"net/rpc"
)

type App struct {
	log       *zap.Logger
	port      string
	RPCServer *rpc.Server
}

func NewRPC(log *zap.Logger, config *config.Config) *App {
	return &App{
		log:  log,
		port: config.Port,
	}
}

func (r *App) Start(service *service.Service) {
	r.RPCServer = rpc.NewServer()
	authRPC.RegisterAuthServer(r.RPCServer, service)

	l, err := net.Listen("tcp", ":"+r.port)
	if err != nil {
		r.log.Fatal("Failed to listen", zap.Error(err))
	}

	r.log.Info("RPC server started on port " + r.port)
	r.RPCServer.Accept(l)
}
