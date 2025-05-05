package services_bot

import (
	enums_bot "backend/internal/enums/bot"
	models_bot "backend/internal/models/bot"
)

func (object *botServiceImplementation) Init() {
	var botsModels []models_bot.BotModel

	if err := object.storageService().DB().
		Where("status = ?", enums_bot.StatusStart).
		Find(&botsModels).Error; err != nil {
		object.loggerService().Error().Printf("failed to load bots on init: %v", err)
		return
	}

	for _, botModel := range botsModels {
		botModel.Status = enums_bot.StatusNew

		if err := object.storageService().DB().Save(&botModel).Error; err != nil {
			object.loggerService().Error().Printf("failed to reset bot status: %v", err)
			continue
		}

		object.botRepositoryService().Add(&botModel)
		object.stopChannels[botModel.ID] = make(chan struct{})
		object.GetRunChannel() <- botModel.ID
	}
}
