package services_provider

import (
	services_exchange "backend/internal/services/exchange"
	services_interface_exchange "backend/internal/services/exchange/interface"
)

func (object *ProviderService) ExchangeService() services_interface_exchange.ExchangeService {
	if object.exchangeService == nil {
		object.exchangeService = services_exchange.NewExchangeService(
			object.StorageService,
			object.ExchangeLimitService,
		)
	}

	return object.exchangeService
}
