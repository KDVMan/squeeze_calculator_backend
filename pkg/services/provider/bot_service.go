package services_provider

import (
	services_bot "backend/internal/services/bot"
	services_interface_bot "backend/internal/services/bot/interface"
)

func (object *ProviderService) BotService() services_interface_bot.BotService {
	if object.botService == nil {
		object.botService = services_bot.NewBotService(
			object.LoggerService,
			object.ConfigService,
			object.StorageService,
			object.WebsocketService,
			object.DumpService,
			object.CalculatorPresetService,
			object.CalculatorFormulaPresetService,
			object.SymbolService,
			object.QuoteService,
			object.QuoteRepositoryService,
			object.ExchangeWebsocketService,
			object.SymbolListService,
		)
	}

	return object.botService
}
