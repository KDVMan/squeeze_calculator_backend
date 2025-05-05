package services_bot

import (
	enums_websocket "backend/internal/enums/websocket"
	models_bot "backend/internal/models/bot"
	models_websocket "backend/internal/models/websocket"
	"time"
)

func (object *botServiceImplementation) CalculateChannel() {
	const checkTime = 3000

	for request := range object.calculateChannel {
		botModel, exists := object.botRepositoryService().Get(request.BotID)
		if !exists {
			continue
		}

		if request.Result == nil {
			botModel.ParamOld = models_bot.ParamModel{
				HasData: botModel.Param.HasData,
			}

			botModel.Param = models_bot.ParamModel{
				LastUpdate: time.Now().UnixMilli(),
				HasData:    true,
			}

			botModel.ApiSend = true

			if err := object.storageService().DB().Save(botModel).Error; err != nil {
				object.loggerService().Error().Printf("failed to save empty params: %v", err)
				continue
			}

			if err := object.sendApi(botModel); err != nil {
				object.loggerService().Error().Printf("failed to send API for bot %d (empty): %v", botModel.ID, err)
				botModel.ApiSend = false
				_ = object.storageService().DB().Save(&botModel)
				object.addRetryApi(botModel.ID)
			}

			object.websocketService().GetBroadcastChannel() <- &models_websocket.BroadcastChannelModel{
				Event: enums_websocket.WebsocketEventBot,
				Data:  botModel,
			}

			continue
		}

		if botModel.Param.LastUpdate+checkTime > time.Now().UnixMilli() {
			continue
		}

		if request.Result.ParamModel.Bind != botModel.Param.Bind ||
			request.Result.ParamModel.PercentIn != botModel.Param.PercentIn ||
			request.Result.ParamModel.PercentOut != botModel.Param.PercentOut ||
			request.Result.ParamModel.StopTime != botModel.Param.StopTime ||
			request.Result.ParamModel.StopPercent != botModel.Param.StopPercent {
			botModel.ParamOld = models_bot.ParamModel{
				Bind:        botModel.Param.Bind,
				PercentIn:   botModel.Param.PercentIn,
				PercentOut:  botModel.Param.PercentOut,
				StopTime:    botModel.Param.StopTime,
				StopPercent: botModel.Param.StopPercent,
				Score:       botModel.Param.Score,
				LastUpdate:  botModel.Param.LastUpdate,
				HasData:     botModel.Param.HasData,
			}

			botModel.Param = models_bot.ParamModel{
				Bind:        request.Result.ParamModel.Bind,
				PercentIn:   request.Result.ParamModel.PercentIn,
				PercentOut:  request.Result.ParamModel.PercentOut,
				StopTime:    request.Result.ParamModel.StopTime,
				StopPercent: request.Result.ParamModel.StopPercent,
				Score:       request.Result.Score,
				LastUpdate:  time.Now().UnixMilli(),
				HasData:     true,
			}

			botModel.ApiSend = true

			if err := object.storageService().DB().Save(botModel).Error; err != nil {
				object.loggerService().Error().Printf("failed to save bot params: %v", err)
				continue
			}

			if err := object.sendApi(botModel); err != nil {
				object.loggerService().Error().Printf("failed to send API for bot %d: %v", botModel.ID, err)

				botModel.ApiSend = false

				_ = object.storageService().DB().Save(&botModel)

				object.addRetryApi(botModel.ID)
			}

			object.websocketService().GetBroadcastChannel() <- &models_websocket.BroadcastChannelModel{
				Event: enums_websocket.WebsocketEventBot,
				Data:  botModel,
			}
		}
	}
}

func (object *botServiceImplementation) GetCalculateChannel() chan *models_bot.CalculateRequestModel {
	return object.calculateChannel
}
