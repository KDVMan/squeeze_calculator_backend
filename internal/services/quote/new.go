package services_quote

import (
	services_interface_calculator "backend/internal/services/calculator/interface"
	services_interface_calculator_preset "backend/internal/services/calculator_preset/interface"
	services_interface_exchange "backend/internal/services/exchange/interface"
	services_interface_exchange_limit "backend/internal/services/exchange_limit/interface"
	services_interface_exchange_websocket "backend/internal/services/exchange_websocket/interface"
	services_interface_quote "backend/internal/services/quote/interface"
	services_interface_websocket "backend/internal/services/websocket/interface"
	services_interface_storage "backend/pkg/services/storage/interface"
)

type quoteServiceImplementation struct {
	storageService           func() services_interface_storage.StorageService
	websocketService         func() services_interface_websocket.WebsocketService
	exchangeService          func() services_interface_exchange.ExchangeService
	exchangeWebsocketService func() services_interface_exchange_websocket.ExchangeWebSocketService
	calculatorService        func() services_interface_calculator.CalculatorService
	calculatorPresetService  func() services_interface_calculator_preset.CalculatorPresetService
	exchangeLimitService     func() services_interface_exchange_limit.ExchangeLimitService
}

func NewQuoteService(
	storageService func() services_interface_storage.StorageService,
	websocketService func() services_interface_websocket.WebsocketService,
	exchangeService func() services_interface_exchange.ExchangeService,
	exchangeWebsocketService func() services_interface_exchange_websocket.ExchangeWebSocketService,
	calculatorService func() services_interface_calculator.CalculatorService,
	calculatorPresetService func() services_interface_calculator_preset.CalculatorPresetService,
	exchangeLimitService func() services_interface_exchange_limit.ExchangeLimitService,
) services_interface_quote.QuoteService {
	return &quoteServiceImplementation{
		storageService:           storageService,
		websocketService:         websocketService,
		exchangeService:          exchangeService,
		exchangeWebsocketService: exchangeWebsocketService,
		calculatorService:        calculatorService,
		calculatorPresetService:  calculatorPresetService,
		exchangeLimitService:     exchangeLimitService,
	}
}
