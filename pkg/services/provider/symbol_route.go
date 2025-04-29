package services_provider

import (
	routes_symbol "backend/internal/routes/symbol"
	routes_interface_symbol "backend/internal/routes/symbol/interface"
)

func (object *ProviderService) SymbolRoute() routes_interface_symbol.SymbolRoute {
	if object.symbolRoute == nil {
		object.symbolRoute = routes_symbol.NewSymbolRoute(
			object.LoggerService,
			object.RequestService,
			object.SymbolService,
			object.ExchangeService,
		)
	}

	return object.symbolRoute
}
