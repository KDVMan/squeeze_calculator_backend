package services_websocket

import (
	enums_websocket "backend/internal/enums/websocket"
	models_websocket "backend/internal/models/websocket"
)

func (object *websocketServiceImplementation) Start() {
	object.loggerService().Info().Printf("starting websocket service")

	for {
		select {
		case connection := <-object.registerChannel:
			object.lock.Lock()
			object.connections[connection] = true
			object.lock.Unlock()

			object.loggerService().Info().Printf("registered websocket connection")

			go object.broadcastSymbols()
			go object.broadcastExchangeLimits()

			object.broadcastChannel <- &models_websocket.BroadcastChannelModel{
				Event: enums_websocket.WebsocketEventCalculateResult,
				Data:  object.calculatorService().LoadResult(""),
			}

			object.broadcastChannel <- &models_websocket.BroadcastChannelModel{
				Event: enums_websocket.WebsocketEventBotList,
				Data:  object.botService().LoadAll(),
			}
		case connection := <-object.unregisterChannel:
			object.lock.Lock()

			if _, ok := object.connections[connection]; ok {
				delete(object.connections, connection)
				close(connection.GetBroadcastChannel())

				object.loggerService().Info().Printf("unregistered websocket connection")
			}

			object.lock.Unlock()
		case data := <-object.broadcastChannel:
			object.broadcast(data)
		case data := <-object.progressChannel:
			object.broadcast(&models_websocket.BroadcastChannelModel{
				Event: data.Event,
				Data:  data,
			})
		}
	}
}
