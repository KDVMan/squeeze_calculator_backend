package services_exchange_websocket

import (
	"backend/internal/enums"
	services_interface_exchange "backend/internal/services/exchange/interface"
	services_interface_exchange_websocket "backend/internal/services/exchange_websocket/interface"
	services_interface_quote "backend/internal/services/quote/interface"
	services_interface_quote_repository "backend/internal/services/quote_repository/interface"
	services_interface_symbol "backend/internal/services/symbol/interface"
	services_interface_logger "backend/pkg/services/logger/interface"
	"sync"
	"time"
)

type exchangeWebsocketServiceImplementation struct {
	loggerService           func() services_interface_logger.LoggerService
	symbolService           func() services_interface_symbol.SymbolService
	exchangeService         func() services_interface_exchange.ExchangeService
	quoteService            func() services_interface_quote.QuoteService
	quoteRepositoryService  func() services_interface_quote_repository.QuoteRepositoryService
	currentPriceSymbol      string
	currentPriceInterval    enums.Interval
	currentPriceStopChannel chan struct{}
	symbolsSubscriptions    map[string]chan struct{}
	symbolMutex             sync.Mutex
	doneChannel             chan struct{}
	reconnectDelay          time.Duration
}

func NewExchangeWebsocketService(
	loggerService func() services_interface_logger.LoggerService,
	symbolService func() services_interface_symbol.SymbolService,
	exchangeService func() services_interface_exchange.ExchangeService,
	quoteService func() services_interface_quote.QuoteService,
	quoteRepositoryService func() services_interface_quote_repository.QuoteRepositoryService,
) services_interface_exchange_websocket.ExchangeWebSocketService {
	return &exchangeWebsocketServiceImplementation{
		loggerService:           loggerService,
		symbolService:           symbolService,
		exchangeService:         exchangeService,
		quoteService:            quoteService,
		quoteRepositoryService:  quoteRepositoryService,
		currentPriceStopChannel: nil,
		symbolsSubscriptions:    make(map[string]chan struct{}),
		symbolMutex:             sync.Mutex{},
		doneChannel:             make(chan struct{}),
		reconnectDelay:          5 * time.Second,
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

	object.symbolMutex.Lock()

	for symbol, stopChannel := range object.symbolsSubscriptions {
		object.loggerService().Info().Printf("unsubscribing from %s", symbol)
		close(stopChannel)
		delete(object.symbolsSubscriptions, symbol)
	}

	object.symbolMutex.Unlock()

	// if err := object.exchangeService().DeleteListenKey(); err != nil {
	// 	object.loggerService().Error().Printf("failed to delete listen key: %v", err)
	// }

	close(object.doneChannel)
}
