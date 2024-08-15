package run

import (
	"boilerplate/config"
	"boilerplate/internal/infrastructure/components"
	srvc "boilerplate/internal/module/service"
	store "boilerplate/internal/module/storage"
	"boilerplate/internal/rpc"
	"go.uber.org/zap"
)

type App struct {
	log    *zap.Logger
	config *config.Config
}

func NewApp(log *zap.Logger, config *config.Config) *App {
	return &App{
		log:    log,
		config: config,
	}
}

func (a *App) Run() {
	component := components.NewComponents(a.config, a.log) // будущие компоненты
	storage := store.NewStorage(component.Config, component.Log)
	service := srvc.NewService(component.Config, component.Log, storage)
	application := rpc.NewRPC(component.Log, component.Config) // инициализация RPC сервера
	application.Start(service)                                 // его запуск
}
