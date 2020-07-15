package services

import (
	"sync"

	"kodix/src/internal/auto/services"
	"kodix/src/internal/config"
	"kodix/src/internal/database"
	"kodix/src/internal/facade/services/factory"
	"kodix/src/internal/logger"
)

type Facade struct {
	factory *factory.ServicesFactory
}

var (
	facade *Facade
	once   sync.Once
)

func NewFacade() *Facade {
	once.Do(func() {
		conf := config.NewConfig()
		connector := database.NewConnectManager(conf.Database)
		if err := connector.InitConnections(); err != nil {
			logger.Logger.Fatal(err)
		}

		facade = &Facade{
			factory: factory.NewServicesFactory(connector),
		}
	})

	return facade
}

// Auto service ...
func (h Facade) Auto() *services.AutoService {
	return h.factory.Auto()
}
