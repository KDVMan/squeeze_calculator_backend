package services_bot

import (
	models_bot "backend/internal/models/bot"
	"time"
)

func (object *botServiceImplementation) retryApi() {
	for botID := range object.retryApiChannel {
		var botModel models_bot.BotModel

		if err := object.storageService().DB().First(&botModel, botID).Error; err != nil {
			object.loggerService().Error().Printf("retry failed to load bot %d: %v", botID, err)
			continue
		}

		if botModel.ApiSend {
			continue
		}

		if err := object.sendApi(&botModel); err != nil {
			object.loggerService().Error().Printf("retry failed to send API for bot %d: %v", botID, err)
			object.addRetryApi(botID)
			continue
		}

		botModel.ApiSend = true
		_ = object.storageService().DB().Save(&botModel)

		object.loggerService().Info().Printf("retry successfully sent API for bot %d", botID)
	}
}

func (object *botServiceImplementation) addRetryApi(botID uint) {
	go func() {
		time.Sleep(10 * time.Second)
		object.retryApiChannel <- botID
	}()
}
