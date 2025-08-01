package services_bot

import (
	enums_bot "backend/internal/enums/bot"
	models_bot "backend/internal/models/bot"
	services_interface_bot "backend/internal/services/bot/interface"
	services_interface_bot_repository "backend/internal/services/bot_repository/interface"
	services_interface_calculator_formula_preset "backend/internal/services/calculator_formula_preset/interface"
	services_interface_calculator_preset "backend/internal/services/calculator_preset/interface"
	services_interface_exchange_websocket "backend/internal/services/exchange_websocket/interface"
	services_interface_quote "backend/internal/services/quote/interface"
	services_interface_quote_repository "backend/internal/services/quote_repository/interface"
	services_interface_symbol "backend/internal/services/symbol/interface"
	services_interface_symbol_list "backend/internal/services/symbol_list/interface"
	services_interface_websocket "backend/internal/services/websocket/interface"
	services_interface_config "backend/pkg/services/config/interface"
	services_interface_dump "backend/pkg/services/dump/interface"
	services_interface_logger "backend/pkg/services/logger/interface"
	services_interface_storage "backend/pkg/services/storage/interface"
	"log"
	"sync"
	"sync/atomic"
)

type botServiceImplementation struct {
	loggerService                  func() services_interface_logger.LoggerService
	configService                  func() services_interface_config.ConfigService
	storageService                 func() services_interface_storage.StorageService
	websocketService               func() services_interface_websocket.WebsocketService
	dumpService                    func() services_interface_dump.DumpService
	calculatorPresetService        func() services_interface_calculator_preset.CalculatorPresetService
	calculatorFormulaPresetService func() services_interface_calculator_formula_preset.CalculatorFormulaPresetService
	symbolService                  func() services_interface_symbol.SymbolService
	quoteService                   func() services_interface_quote.QuoteService
	quoteRepositoryService         func() services_interface_quote_repository.QuoteRepositoryService
	exchangeWebsocketService       func() services_interface_exchange_websocket.ExchangeWebSocketService
	symbolListService              func() services_interface_symbol_list.SymbolListService
	botRepositoryService           func() services_interface_bot_repository.BotRepositoryService
	futuresLimit                   int64
	futuresCommission              float64
	runChannel                     chan uint
	calculatorChannel              chan uint
	calculateChannel               chan *models_bot.CalculateRequestModel
	stopChannels                   map[uint]chan struct{}
	calculatorStop                 atomic.Bool
	retryApiChannel                chan uint
	delayedCoins                   sync.Map
	test                           bool
}

func NewBotService(
	loggerService func() services_interface_logger.LoggerService,
	configService func() services_interface_config.ConfigService,
	storageService func() services_interface_storage.StorageService,
	websocketService func() services_interface_websocket.WebsocketService,
	dumpService func() services_interface_dump.DumpService,
	calculatorPresetService func() services_interface_calculator_preset.CalculatorPresetService,
	calculatorFormulaPresetService func() services_interface_calculator_formula_preset.CalculatorFormulaPresetService,
	symbolService func() services_interface_symbol.SymbolService,
	quoteService func() services_interface_quote.QuoteService,
	quoteRepositoryService func() services_interface_quote_repository.QuoteRepositoryService,
	exchangeWebsocketService func() services_interface_exchange_websocket.ExchangeWebSocketService,
	symbolListService func() services_interface_symbol_list.SymbolListService,
	botRepositoryService func() services_interface_bot_repository.BotRepositoryService,
) services_interface_bot.BotService {
	service := &botServiceImplementation{
		loggerService:                  loggerService,
		configService:                  configService,
		storageService:                 storageService,
		websocketService:               websocketService,
		dumpService:                    dumpService,
		calculatorPresetService:        calculatorPresetService,
		calculatorFormulaPresetService: calculatorFormulaPresetService,
		symbolService:                  symbolService,
		quoteService:                   quoteService,
		quoteRepositoryService:         quoteRepositoryService,
		exchangeWebsocketService:       exchangeWebsocketService,
		symbolListService:              symbolListService,
		botRepositoryService:           botRepositoryService,
		futuresLimit:                   int64(configService().GetConfig().Binance.FuturesLimit),
		futuresCommission:              configService().GetConfig().Binance.FuturesCommission,
		runChannel:                     make(chan uint, 10000),
		calculatorChannel:              make(chan uint, 10000),
		calculateChannel:               make(chan *models_bot.CalculateRequestModel, 10000),
		stopChannels:                   make(map[uint]chan struct{}),
		retryApiChannel:                make(chan uint, 1000),
		test:                           false,
	}

	go service.retryApi()

	return service
}

func (object *botServiceImplementation) SetTest() {
	log.Println("test 1", object.test)
	object.test = !object.test
	log.Println("test 2", object.test)
}

func (object *botServiceImplementation) StopBot(botModel *models_bot.BotModel) {
	var count int64

	if ch, ok := object.stopChannels[botModel.ID]; ok {
		close(ch)
		delete(object.stopChannels, botModel.ID)
	}

	object.storageService().DB().
		Model(&models_bot.BotModel{}).
		Where("symbol = ? AND status != ?", botModel.Symbol, enums_bot.StatusStop).
		Count(&count)

	if count == 0 {
		object.quoteRepositoryService().Remove(botModel.Symbol)
		object.exchangeWebsocketService().UnsubscribeSymbol(botModel.Symbol)
	}
}
