package services_bot

import (
	enums_websocket "backend/internal/enums/websocket"
	models_bot "backend/internal/models/bot"
	models_websocket "backend/internal/models/websocket"
	"log"
)

func (object *botServiceImplementation) CalculateChannel() {
	for request := range object.calculateChannel {
		newParam := request.Result.ParamModel
		oldParam := request.CalculatorRequestModel.Param

		if newParam.Bind != oldParam.Bind ||
			newParam.PercentIn != oldParam.PercentIn ||
			newParam.PercentOut != oldParam.PercentOut ||
			newParam.StopTime != oldParam.StopTime ||
			newParam.StopPercent != oldParam.StopPercent {
			oldScore := request.CalculatorRequestModel.Param.Score
			oldProfit := request.CalculatorRequestModel.Param.Profit

			request.CalculatorRequestModel.Param = models_bot.ParamModel{
				Bind:        newParam.Bind,
				PercentIn:   newParam.PercentIn,
				PercentOut:  newParam.PercentOut,
				StopTime:    newParam.StopTime,
				StopPercent: newParam.StopPercent,
				Score:       request.Result.Score,
				Profit:      request.Result.TotalCumulativeProfitPercent,
			}

			var botModel models_bot.BotModel

			if err := object.storageService().DB().First(&botModel, request.CalculatorRequestModel.BotID).Error; err != nil {
				object.loggerService().Error().Printf("failed to load bot: %v", err)
				continue
			}

			botModel.Param = request.CalculatorRequestModel.Param

			if err := object.storageService().DB().Save(&botModel).Error; err != nil {
				object.loggerService().Error().Printf("failed to save bot params: %v", err)
				continue
			}

			// object.sendToApi(bot)

			object.websocketService().GetBroadcastChannel() <- &models_websocket.BroadcastChannelModel{
				Event: enums_websocket.WebsocketEventBot,
				Data:  botModel,
			}

			log.Printf(
				"symbol: %s, direction: %s, score: %.4f -> %.4f, "+
					"profit: %.2f -> %.2f, "+
					"bind: %v -> %v, "+
					"in: %.2f -> %.2f, "+
					"out: %.2f -> %.2f, "+
					"stopTime: %v -> %v, "+
					"stopPercent: %.2f -> %.2f\n\n",
				request.CalculatorRequestModel.Symbol,
				request.CalculatorRequestModel.TradeDirection,
				oldScore,
				request.Result.Score,
				oldProfit,
				request.Result.TotalCumulativeProfitPercent,
				oldParam.Bind,
				newParam.Bind,
				oldParam.PercentIn, newParam.PercentIn,
				oldParam.PercentOut, newParam.PercentOut,
				oldParam.StopTime, newParam.StopTime,
				oldParam.StopPercent, newParam.StopPercent,
			)
		}
	}
}

func (object *botServiceImplementation) GetCalculateChannel() chan *models_bot.CalculateRequestModel {
	return object.calculateChannel
}
