package auth

import (
	"boilerplate/internal/module/service"
	"net/rpc"
)

type RPCAuth struct{}

func RegisterAuthServer(aRPC *rpc.Server, service *service.Service) {
	err := aRPC.Register(service)
	if err != nil {
		panic(err)
	}
}
