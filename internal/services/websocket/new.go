package services_websocket

import (
	models_websocket "backend/internal/models/websocket"
	services_interface_bot "backend/internal/services/bot/interface"
	services_interface_calculator "backend/internal/services/calculator/interface"
	services_interface_exchange_limit "backend/internal/services/exchange_limit/interface"
	services_interface_symbol "backend/internal/services/symbol/interface"
	services_websocket_interface "backend/internal/services/websocket/interface"
	services_websocket_connection_interface "backend/internal/services/websocket_connection/interface"
	services_interface_logger "backend/pkg/services/logger/interface"
	"sync"
)

type websocketServiceImplementation struct {
	loggerService        func() services_interface_logger.LoggerService
	exchangeLimitService func() services_interface_exchange_limit.ExchangeLimitService
	symbolService        func() services_interface_symbol.SymbolService
	calculatorService    func() services_interface_calculator.CalculatorService
	botService           func() services_interface_bot.BotService
	connections          map[services_websocket_connection_interface.WebsocketConnectionService]bool
	registerChannel      chan services_websocket_connection_interface.WebsocketConnectionService
	unregisterChannel    chan services_websocket_connection_interface.WebsocketConnectionService
	broadcastChannel     chan *models_websocket.BroadcastChannelModel
	progressChannel      chan *models_websocket.ProgressChannelModel
	lock                 sync.Mutex
}

func NewWebsocketService(
	loggerService func() services_interface_logger.LoggerService,
	exchangeLimitService func() services_interface_exchange_limit.ExchangeLimitService,
	symbolService func() services_interface_symbol.SymbolService,
	calculatorService func() services_interface_calculator.CalculatorService,
	botService func() services_interface_bot.BotService,
) services_websocket_interface.WebsocketService {
	return &websocketServiceImplementation{
		loggerService:        loggerService,
		exchangeLimitService: exchangeLimitService,
		symbolService:        symbolService,
		calculatorService:    calculatorService,
		botService:           botService,
		connections:          make(map[services_websocket_connection_interface.WebsocketConnectionService]bool),
		registerChannel:      make(chan services_websocket_connection_interface.WebsocketConnectionService, 1000),
		unregisterChannel:    make(chan services_websocket_connection_interface.WebsocketConnectionService, 1000),
		broadcastChannel:     make(chan *models_websocket.BroadcastChannelModel, 10000),
		progressChannel:      make(chan *models_websocket.ProgressChannelModel, 10000),
	}
}

func (object *websocketServiceImplementation) GetRegisterChannel() chan services_websocket_connection_interface.WebsocketConnectionService {
	return object.registerChannel
}

func (object *websocketServiceImplementation) GetUnregisterChannel() chan services_websocket_connection_interface.WebsocketConnectionService {
	return object.unregisterChannel
}

func (object *websocketServiceImplementation) GetBroadcastChannel() chan *models_websocket.BroadcastChannelModel {
	return object.broadcastChannel
}

func (object *websocketServiceImplementation) GetProgressChannel() chan *models_websocket.ProgressChannelModel {
	return object.progressChannel
}
