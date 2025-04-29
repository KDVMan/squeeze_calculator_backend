package services_provider

import (
	routes_quote "backend/internal/routes/quote"
	routes_interface_quote "backend/internal/routes/quote/interface"
)

func (object *ProviderService) QuoteRoute() routes_interface_quote.QuoteRoute {
	if object.quoteRoute == nil {
		object.quoteRoute = routes_quote.NewQuoteRoute(
			object.LoggerService,
			object.RequestService,
			object.QuoteService,
		)
	}

	return object.quoteRoute
}
