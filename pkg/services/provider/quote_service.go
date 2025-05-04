package services_provider

import (
	services_quote "backend/internal/services/quote"
	services_interface_quote "backend/internal/services/quote/interface"
)

func (object *ProviderService) QuoteService() services_interface_quote.QuoteService {
	if object.quoteService == nil {
		object.quoteService = services_quote.NewQuoteService(
			object.StorageService,
			object.WebsocketService,
			object.ExchangeService,
			object.ExchangeWebsocketService,
			object.CalculatorService,
			object.CalculatorPresetService,
			object.ExchangeLimitService,
		)
	}

	return object.quoteService
}
