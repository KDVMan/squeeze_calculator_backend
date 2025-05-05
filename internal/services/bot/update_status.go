package services_bot

import (
	enums_bot "backend/internal/enums/bot"
	enums_websocket "backend/internal/enums/websocket"
	models_bot "backend/internal/models/bot"
	models_websocket "backend/internal/models/websocket"
)

func (object *botServiceImplementation) UpdateStatus(request *models_bot.UpdateStatusRequestModel) error {
	var botModel models_bot.BotModel

	err := object.storageService().DB().First(&botModel, request.ID).Error
	if err != nil {
		return err
	}

	if request.Status == enums_bot.StatusNew {
		object.stopChannels[botModel.ID] = make(chan struct{})
		object.botRepositoryService().Add(&botModel)
		object.GetRunChannel() <- botModel.ID
	} else if request.Status == enums_bot.StatusStop {
		botModel.Status = enums_bot.StatusStop

		if err = object.storageService().DB().Save(&botModel).Error; err != nil {
			return err
		}

		object.botRepositoryService().Add(&botModel)
		object.StopBot(&botModel)

		object.websocketService().GetBroadcastChannel() <- &models_websocket.BroadcastChannelModel{
			Event: enums_websocket.WebsocketEventBot,
			Data:  botModel,
		}
	} else if request.Status == enums_bot.StatusDelete {
		if err = object.storageService().DB().Delete(&botModel).Error; err != nil {
			return err
		}

		object.botRepositoryService().Remove(botModel.ID)

		object.websocketService().GetBroadcastChannel() <- &models_websocket.BroadcastChannelModel{
			Event: enums_websocket.WebsocketEventBotList,
			Data:  object.LoadAll(),
		}
	}

	return nil
}
