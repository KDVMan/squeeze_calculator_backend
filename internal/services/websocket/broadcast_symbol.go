package services_websocket

import (
	enums_websocket "backend/internal/enums/websocket"
	models_websocket "backend/internal/models/websocket"
)

func (object *websocketServiceImplementation) broadcastSymbols() {
	symbols, err := object.symbolService().LoadAll()

	if err != nil {
		object.loggerService().Error().Printf("failed to load symbols: %v", err)
		return
	}

	object.broadcastChannel <- &models_websocket.BroadcastChannelModel{
		Event: enums_websocket.WebsocketEventSymbolList,
		Data:  symbols,
	}
}
