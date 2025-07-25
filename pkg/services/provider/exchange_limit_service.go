package services_provider

import (
	services_exchange_limit "backend/internal/services/exchange_limit"
	services_interface_exchange_limit "backend/internal/services/exchange_limit/interface"
)

func (object *ProviderService) ExchangeLimitService() services_interface_exchange_limit.ExchangeLimitService {
	if object.exchangeLimitService == nil {
		object.exchangeLimitService = services_exchange_limit.NewExchangeLimitService(
			object.StorageService,
			object.WebsocketService,
		)
	}

	return object.exchangeLimitService
}
