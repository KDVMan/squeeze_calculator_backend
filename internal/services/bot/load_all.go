package services_bot

import (
	models_bot "backend/internal/models/bot"
	"errors"
	"gorm.io/gorm"
)

func (object *botServiceImplementation) LoadAll() []*models_bot.BotModel {
	var botsModels []*models_bot.BotModel

	if err := object.storageService().DB().Find(&botsModels).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			object.loggerService().Error().Printf("failed to load bots: %v", err)
		}

		return []*models_bot.BotModel{}
	}

	// repoBots := object.botRepositoryService().GetAll()
	//
	// for i, dbBot := range botsModels {
	// 	for _, repoBot := range repoBots {
	// 		if dbBot.ID == repoBot.ID {
	// 			botsModels[i] = repoBot
	// 			break
	// 		}
	// 	}
	// }

	return botsModels
}
