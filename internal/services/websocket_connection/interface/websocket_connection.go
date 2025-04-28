package services_interface_websocket_connection

import (
	models_websocket "backend/internal/models/websocket"
)

type WebsocketConnectionService interface {
	Read()
	Write()
	GetBroadcastChannel() chan *models_websocket.BroadcastChannelModel
}
