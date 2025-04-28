package services_exchange_websocket

import (
	"backend/internal/enums"
	services_interface_exchange "backend/internal/services/exchange/interface"
	services_interface_exchange_websocket "backend/internal/services/exchange_websocket/interface"
	services_interface_quote "backend/internal/services/quote/interface"
	services_interface_symbol "backend/internal/services/symbol/interface"
	services_interface_logger "backend/pkg/services/logger/interface"
	"sync"
)

type exchangeWebsocketServiceImplementation struct {
	loggerService   func() services_interface_logger.LoggerService
	symbolService   func() services_interface_symbol.SymbolService
	exchangeService func() services_interface_exchange.ExchangeService
	quoteService    func() services_interface_quote.QuoteService
	// quoteRepositoryService  func() services_quote_repository_interface.QuoteRepositoryService
	// userService             func() services_user_interface.UserService
	// tradeService            func() services_trade_interface.TradeService
	currentPriceSymbol      string
	currentPriceInterval    enums.Interval
	currentPriceStopChannel chan struct{}
	currentPriceMutex       sync.Mutex
	doneChannel             chan struct{}
	// tradesSubscriptions     map[string]chan struct{}
	// tradeMutex              sync.Mutex
}

func NewExchangeWebsocketService(
	loggerService func() services_interface_logger.LoggerService,
	symbolService func() services_interface_symbol.SymbolService,
	exchangeService func() services_interface_exchange.ExchangeService,
	quoteService func() services_interface_quote.QuoteService,
	// quoteRepositoryService func() services_quote_repository_interface.QuoteRepositoryService,
	// userService func() services_user_interface.UserService,
	// tradeService func() services_trade_interface.TradeService,
) services_interface_exchange_websocket.ExchangeWebSocketService {
	return &exchangeWebsocketServiceImplementation{
		loggerService:   loggerService,
		symbolService:   symbolService,
		exchangeService: exchangeService,
		quoteService:    quoteService,
		// quoteRepositoryService:  quoteRepositoryService,
		// userService:             userService,
		// tradeService:            tradeService,
		currentPriceStopChannel: nil,
		currentPriceMutex:       sync.Mutex{},
		doneChannel:             make(chan struct{}),
		// tradesSubscriptions:     make(map[string]chan struct{}),
		// tradeMutex:              sync.Mutex{},
	}
}

func (object *exchangeWebsocketServiceImplementation) Start() {
	object.loggerService().Info().Printf("starting exchange websocket service")

	go object.allMarket()
}

func (object *exchangeWebsocketServiceImplementation) Stop() {
	object.loggerService().Info().Printf("stopping exchange websocket service")

	if object.currentPriceStopChannel != nil {
		close(object.currentPriceStopChannel)
		object.currentPriceStopChannel = nil
	}

	// object.tradeMutex.Lock()

	// for symbol, stopChannel := range object.tradesSubscriptions {
	// 	close(stopChannel)
	// 	delete(object.tradesSubscriptions, symbol)
	// }

	// object.tradeMutex.Unlock()

	// if err := object.exchangeService().DeleteListenKey(); err != nil {
	// 	object.loggerService().Error().Printf("failed to delete listen key: %v", err)
	// }

	close(object.doneChannel)
}
