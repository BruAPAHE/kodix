package facade

import "kodix/src/internal/facade/services"

func Services() *services.Facade {
	return services.NewFacade()
}
