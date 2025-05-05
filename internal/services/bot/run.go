package services_bot

import (
	"backend/internal/enums"
	enums_bot "backend/internal/enums/bot"
	enums_websocket "backend/internal/enums/websocket"
	models_quote "backend/internal/models/quote"
	models_websocket "backend/internal/models/websocket"
)

func (object *botServiceImplementation) RunChannel() {
	for botID := range object.runChannel {
		botModel, exists := object.botRepositoryService().Get(botID)
		if !exists {
			continue
		}

		existingQuotes := object.quoteRepositoryService().GetBySymbol(botModel.Symbol, botModel.TradeDirection)

		if len(existingQuotes) < int(botModel.Window) {
			timeFrom, timeTo := models_quote.GetTimeRange(botModel)
			quoteRange := models_quote.GetRange(object.futuresLimit, timeFrom, timeTo, enums.IntervalMilliseconds(enums.Interval1m))

			quotes, err := object.quoteService().LoadRange(botModel.Symbol, quoteRange, &models_websocket.ProgressChannelModel{})
			if err != nil {
				object.loggerService().Error().
					Printf("failed to load range for %s: %v", botModel.Symbol, err)
				continue
			}

			object.quoteRepositoryService().Add(botModel.Symbol, quotes)
		}

		object.exchangeWebsocketService().SubscribeSymbol(botModel.Symbol)

		botModel.Status = enums_bot.StatusStart

		if err := object.storageService().DB().Save(&botModel).Error; err != nil {
			object.loggerService().Error().Printf("failed to save bot: %v", err)
			continue
		}

		object.websocketService().GetBroadcastChannel() <- &models_websocket.BroadcastChannelModel{
			Event: enums_websocket.WebsocketEventBot,
			Data:  object.LoadByID(botModel.ID),
		}

		object.GetCalculatorChannel() <- botID
	}
}

func (object *botServiceImplementation) GetRunChannel() chan uint {
	return object.runChannel
}
