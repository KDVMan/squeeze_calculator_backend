package services_provider

import (
	routes_interface_bot "backend/internal/routes/bot/interface"
	routes_interface_calculator "backend/internal/routes/calculator/interface"
	routes_interface_calculator_formula_preset "backend/internal/routes/calculator_formula_preset/interface"
	routes_interface_calculator_preset "backend/internal/routes/calculator_preset/interface"
	routes_interface_chart_settings "backend/internal/routes/chart_settings/interface"
	routes_interface_init "backend/internal/routes/init/interface"
	routes_interface_quote "backend/internal/routes/quote/interface"
	routes_interface_symbol "backend/internal/routes/symbol/interface"
	routes_interface_symbol_list "backend/internal/routes/symbol_list/interface"
	services_interface_bot "backend/internal/services/bot/interface"
	services_interface_bot_repository "backend/internal/services/bot_repository/interface"
	services_interface_calculator "backend/internal/services/calculator/interface"
	services_interface_calculator_formula_preset "backend/internal/services/calculator_formula_preset/interface"
	services_interface_calculator_preset "backend/internal/services/calculator_preset/interface"
	services_interface_chart_settings "backend/internal/services/chart_settings/interface"
	services_interface_exchange "backend/internal/services/exchange/interface"
	services_interface_exchange_limit "backend/internal/services/exchange_limit/interface"
	services_interface_exchange_websocket "backend/internal/services/exchange_websocket/interface"
	services_interface_init "backend/internal/services/init/interface"
	services_interface_quote "backend/internal/services/quote/interface"
	services_interface_quote_repository "backend/internal/services/quote_repository/interface"
	services_interface_symbol "backend/internal/services/symbol/interface"
	services_interface_symbol_list "backend/internal/services/symbol_list/interface"
	services_interface_websocket "backend/internal/services/websocket/interface"
	services_interface_config "backend/pkg/services/config/interface"
	services_interface_dump "backend/pkg/services/dump/interface"
	services_interface_logger "backend/pkg/services/logger/interface"
	services_interface_request "backend/pkg/services/request/interface"
	services_interface_router "backend/pkg/services/router/interface"
	services_interface_storage "backend/pkg/services/storage/interface"
	"context"
)

type ProviderService struct {
	ctx       context.Context
	cancelCtx context.CancelFunc

	// ядро
	configService  services_interface_config.ConfigService
	loggerService  services_interface_logger.LoggerService
	storageService services_interface_storage.StorageService
	dumpService    services_interface_dump.DumpService
	routerService  services_interface_router.RouterService
	requestService services_interface_request.RequestService

	// роуты
	initRoute                    routes_interface_init.InitRoute
	symbolRoute                  routes_interface_symbol.SymbolRoute
	symbolListRoute              routes_interface_symbol_list.SymbolListRoute
	chartSettingsRoute           routes_interface_chart_settings.ChartSettingsRoute
	quoteRoute                   routes_interface_quote.QuoteRoute
	calculatorRoute              routes_interface_calculator.CalculatorRoute
	calculatorPresetRoute        routes_interface_calculator_preset.CalculatorPresetRoute
	calculatorFormulaPresetRoute routes_interface_calculator_formula_preset.CalculatorFormulaPresetRoute
	botRoute                     routes_interface_bot.BotRoute

	// сервисы
	websocketService               services_interface_websocket.WebsocketService
	initService                    services_interface_init.InitService
	symbolService                  services_interface_symbol.SymbolService
	symbolListService              services_interface_symbol_list.SymbolListService
	exchangeService                services_interface_exchange.ExchangeService
	exchangeLimitService           services_interface_exchange_limit.ExchangeLimitService
	exchangeWebsocketService       services_interface_exchange_websocket.ExchangeWebSocketService
	chartSettingsService           services_interface_chart_settings.ChartSettingsService
	quoteService                   services_interface_quote.QuoteService
	calculatorService              services_interface_calculator.CalculatorService
	calculatorPresetService        services_interface_calculator_preset.CalculatorPresetService
	calculatorFormulaPresetService services_interface_calculator_formula_preset.CalculatorFormulaPresetService
	botService                     services_interface_bot.BotService
	quoteRepositoryService         services_interface_quote_repository.QuoteRepositoryService
	botRepositoryService           services_interface_bot_repository.BotRepositoryService
}

func NewProviderService(parentCtx context.Context) *ProviderService {
	ctx, cancelCtx := context.WithCancel(parentCtx)

	return &ProviderService{
		ctx:       ctx,
		cancelCtx: cancelCtx,
	}
}

func (object *ProviderService) Shutdown() {
	object.loggerService.Info().Println("shutting down provider service...")

	if object.exchangeWebsocketService != nil {
		object.exchangeWebsocketService.Stop()
	}

	if object.websocketService != nil {
		object.websocketService.Stop()
	}

	object.cancelCtx()

	object.loggerService.Info().Println("provider service stopped.")
}
