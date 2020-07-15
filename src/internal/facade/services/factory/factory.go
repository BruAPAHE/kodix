package factory

import (
	"kodix/src/internal/auto/repositories"
	"kodix/src/internal/auto/services"
	"kodix/src/internal/database"
	"kodix/src/internal/logger"
)

type ServicesFactory struct {
	auto      *services.AutoService
	connector *database.ConnectManager
}

func NewServicesFactory(connector *database.ConnectManager) *ServicesFactory {
	factory := &ServicesFactory{
		connector: connector,
	}
	if err := factory.initServices(); err != nil {
		logger.Logger.Fatal(err)
	}

	return factory
}

func (factory *ServicesFactory) initServices() error {
	return factory.createAutoService()
}

func (factory *ServicesFactory) createAutoService() error {
	connect, err := factory.connector.ConnectMongo("AutoData")
	if err != nil {
		return err
	}

	repo := repositories.NewAutoRepository(connect)
	factory.auto = services.NewAutoService(repo)

	return nil
}

func (factory *ServicesFactory) Auto() *services.AutoService {
	return factory.auto
}
