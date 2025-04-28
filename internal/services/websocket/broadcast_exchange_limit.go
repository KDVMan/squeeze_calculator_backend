package services_websocket

import (
	enums_websocket "backend/internal/enums/websocket"
	models_websocket "backend/internal/models/websocket"
)

func (object *websocketServiceImplementation) broadcastExchangeLimits() {
	limits, err := object.exchangeLimitService().Load()

	if err != nil {
		object.loggerService().Error().Printf("failed to load exchange limits: %v", err)
		return
	}

	object.broadcastChannel <- &models_websocket.BroadcastChannelModel{
		Event: enums_websocket.WebsocketEventExchangeLimits,
		Data:  limits,
	}
}
