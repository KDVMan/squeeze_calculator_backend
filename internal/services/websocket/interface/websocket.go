package services_interface_websocket

import (
	models_websocket "backend/internal/models/websocket"
	services_websocket_connection_interface "backend/internal/services/websocket_connection/interface"
)

type WebsocketService interface {
	Start()
	Stop()
	GetRegisterChannel() chan services_websocket_connection_interface.WebsocketConnectionService
	GetUnregisterChannel() chan services_websocket_connection_interface.WebsocketConnectionService
	GetBroadcastChannel() chan *models_websocket.BroadcastChannelModel
	GetProgressChannel() chan *models_websocket.ProgressChannelModel
}
