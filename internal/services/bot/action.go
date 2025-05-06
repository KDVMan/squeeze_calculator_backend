package services_bot

import (
	enums_bot "backend/internal/enums/bot"
	enums_websocket "backend/internal/enums/websocket"
	models_bot "backend/internal/models/bot"
	models_websocket "backend/internal/models/websocket"
)

func (object *botServiceImplementation) Action(request *models_bot.ActionRequestModel) error {
	switch request.Action {
	case enums_bot.ActionStartAll:
		return object.actionStartAll()
	case enums_bot.ActionStopAll:
		return object.actionStopAll()
	case enums_bot.ActionDeleteAll:
		return object.actionDeleteAll()
	}

	return nil
}

func (object *botServiceImplementation) actionStartAll() error {
	var botsModels []models_bot.BotModel

	if err := object.storageService().DB().
		Model(&models_bot.BotModel{}).
		Where("status NOT IN ?", []enums_bot.Status{enums_bot.StatusNew, enums_bot.StatusStart}).
		Update("status", enums_bot.StatusNew).Error; err != nil {
		return err
	}

	if err := object.storageService().DB().
		Where("status = ?", enums_bot.StatusNew).
		Find(&botsModels).Error; err != nil {
		return err
	}

	for _, botModel := range botsModels {
		object.stopChannels[botModel.ID] = make(chan struct{})
		object.botRepositoryService().Add(&botModel)

		go func(id uint) {
			object.GetRunChannel() <- id
		}(botModel.ID)
	}

	object.websocketService().GetBroadcastChannel() <- &models_websocket.BroadcastChannelModel{
		Event: enums_websocket.WebsocketEventBotList,
		Data:  object.LoadAll(),
	}

	return nil
}

func (object *botServiceImplementation) actionStopAll() error {
	var botsModels []models_bot.BotModel

	if err := object.storageService().DB().
		Model(&models_bot.BotModel{}).
		Where("status IN ?", []enums_bot.Status{enums_bot.StatusNew, enums_bot.StatusStart}).
		Update("status", enums_bot.StatusStop).Error; err != nil {
		return err
	}

	if err := object.storageService().DB().
		Where("status = ?", enums_bot.StatusStop).
		Find(&botsModels).Error; err != nil {
		return err
	}

	for _, botModel := range botsModels {
		object.botRepositoryService().Add(&botModel)
		object.StopBot(&botModel)
	}

	object.websocketService().GetBroadcastChannel() <- &models_websocket.BroadcastChannelModel{
		Event: enums_websocket.WebsocketEventBotList,
		Data:  object.LoadAll(),
	}

	return nil
}

func (object *botServiceImplementation) actionDeleteAll() error {
	var botsModels []models_bot.BotModel

	if err := object.storageService().DB().Where("status = ?", enums_bot.StatusStop).Find(&botsModels).Error; err != nil {
		return err
	}

	for _, botModel := range botsModels {
		if err := object.storageService().DB().Delete(&botModel).Error; err != nil {
			continue
		}

		object.botRepositoryService().Remove(botModel.ID)
	}

	object.websocketService().GetBroadcastChannel() <- &models_websocket.BroadcastChannelModel{
		Event: enums_websocket.WebsocketEventBotList,
		Data:  object.LoadAll(),
	}

	return nil
}
