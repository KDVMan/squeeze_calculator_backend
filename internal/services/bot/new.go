package services_bot

import (
	services_interface_bot "backend/internal/services/bot/interface"
	services_interface_calculator_formula_preset "backend/internal/services/calculator_formula_preset/interface"
	services_interface_calculator_preset "backend/internal/services/calculator_preset/interface"
	services_interface_symbol "backend/internal/services/symbol/interface"
	services_interface_storage "backend/pkg/services/storage/interface"
)

type botServiceImplementation struct {
	storageService                 func() services_interface_storage.StorageService
	calculatorPresetService        func() services_interface_calculator_preset.CalculatorPresetService
	calculatorFormulaPresetService func() services_interface_calculator_formula_preset.CalculatorFormulaPresetService
	symbolService                  func() services_interface_symbol.SymbolService
	// websocketService               func() services_websocket_interface.WebsocketService
	// exchangeWebsocketService       func() services_exchange_websocket_interface.ExchangeWebSocketService
	// tradeRepository                func() services_trade_repository_interface.TradeRepositoryService
	// quoteService                   func() services_quote_interface.QuoteService
	// quoteRepositoryService         func() services_quote_repository_interface.QuoteRepositoryService
	// dealChannel                    chan string
}

func NewBotService(
	// configService func() services_config_interface.ConfigService,
	storageService func() services_interface_storage.StorageService,
	calculatorPresetService func() services_interface_calculator_preset.CalculatorPresetService,
	calculatorFormulaPresetService func() services_interface_calculator_formula_preset.CalculatorFormulaPresetService,
	symbolService func() services_interface_symbol.SymbolService,
	// websocketService func() services_websocket_interface.WebsocketService,
	// exchangeWebsocketService func() services_exchange_websocket_interface.ExchangeWebSocketService,
	// tradeRepository func() services_trade_repository_interface.TradeRepositoryService,
	// quoteService func() services_quote_interface.QuoteService,
	// quoteRepositoryService func() services_quote_repository_interface.QuoteRepositoryService,
	// symbolService func() services_symbol_interface.SymbolService,
	// calculatorFormulaPresetService func() services_calculator_formula_preset_interface.CalculatorFormulaPresetService,
) services_interface_bot.BotService {
	return &botServiceImplementation{
		// configService:                  configService,
		storageService:                 storageService,
		calculatorPresetService:        calculatorPresetService,
		calculatorFormulaPresetService: calculatorFormulaPresetService,
		symbolService:                  symbolService,
		// websocketService:               websocketService,
		// exchangeWebsocketService:       exchangeWebsocketService,
		// tradeRepository:                tradeRepository,
		// quoteService:                   quoteService,
		// quoteRepositoryService:         quoteRepositoryService,
		// dealChannel:                    make(chan string, 1000000),
	}
}
