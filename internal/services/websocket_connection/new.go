package services_websocket_connection

import (
	models_websocket "backend/internal/models/websocket"
	services_websocket_interface "backend/internal/services/websocket/interface"
	services_websocket_connection_interface "backend/internal/services/websocket_connection/interface"
	services_logger_interface "backend/pkg/services/logger/interface"
	"github.com/gorilla/websocket"
)

type websocketConnectionServiceImplementation struct {
	loggerService    func() services_logger_interface.LoggerService
	websocketService func() services_websocket_interface.WebsocketService
	websocket        *websocket.Conn
	broadcastChannel chan *models_websocket.BroadcastChannelModel
}

func NewWebsocketConnectionService(
	loggerService func() services_logger_interface.LoggerService,
	websocketService func() services_websocket_interface.WebsocketService,
	websocket *websocket.Conn,
) services_websocket_connection_interface.WebsocketConnectionService {
	return &websocketConnectionServiceImplementation{
		loggerService:    loggerService,
		websocketService: websocketService,
		websocket:        websocket,
		broadcastChannel: make(chan *models_websocket.BroadcastChannelModel, 1000000),
	}
}

func (object *websocketConnectionServiceImplementation) GetBroadcastChannel() chan *models_websocket.BroadcastChannelModel {
	return object.broadcastChannel
}
