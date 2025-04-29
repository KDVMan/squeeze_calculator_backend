package services_provider

import (
	routes_init "backend/internal/routes/init"
	routes_interface_init "backend/internal/routes/init/interface"
)

func (object *ProviderService) InitRoute() routes_interface_init.InitRoute {
	if object.initRoute == nil {
		object.initRoute = routes_init.NewInitRoute(
			object.LoggerService,
			object.RequestService,
			object.InitService,
		)
	}

	return object.initRoute
}
