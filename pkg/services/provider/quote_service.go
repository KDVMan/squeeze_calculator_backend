package services_provider

import (
	services_quote "backend/internal/services/quote"
	services_quote_interface "backend/internal/services/quote/interface"
)

func (object *ProviderService) QuoteService() services_quote_interface.QuoteService {
	if object.quoteService == nil {
		object.quoteService = services_quote.NewQuoteService(
			object.StorageService,
			object.WebsocketService,
			object.ExchangeService,
			object.ExchangeWebsocketService,
			object.CalculatorService,
			object.CalculatorPresetService,
		)
	}

	return object.quoteService
}
