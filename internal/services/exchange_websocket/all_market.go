package services_exchange_websocket

import (
	"github.com/adshao/go-binance/v2/futures"
	"time"
)

func (object *exchangeWebsocketServiceImplementation) allMarket() {
	reconnectAttempts := 0
	maxReconnectAttempts := 5

	object.loggerService().Info().Printf("starting all market updates")

	reconnect := func() {
		for reconnectAttempts < maxReconnectAttempts {
			select {
			case <-object.doneChannel:
				object.loggerService().Info().Printf("websocket stopped, stopping reconnection attempts")
				return
			default:
				object.loggerService().Info().Printf("attempting to reconnect")
				time.Sleep(5 * time.Second)
				reconnectAttempts++
				object.allMarket()
			}
		}

		object.loggerService().Error().Printf("maximum reconnect attempts reached")
	}

	handler := func(event futures.WsAllMarketTickerEvent) {
		if err := object.symbolService().UpdateStatistic(event); err != nil {
			object.loggerService().Error().Printf("failed to update statistic: %v", err)
		}
	}

	errorHandler := func(err error) {
		object.loggerService().Error().Printf("websocket error: %v", err)
		reconnect()
	}

	_, stop, err := futures.WsAllMarketTickerServe(handler, errorHandler)
	if err != nil {
		object.loggerService().Error().Printf("failed to start websocket: %v", err)
		return
	}

	select {
	case <-object.doneChannel:
		object.loggerService().Info().Printf("websocket stopped")
		stop <- struct{}{}
		return
	}
}
