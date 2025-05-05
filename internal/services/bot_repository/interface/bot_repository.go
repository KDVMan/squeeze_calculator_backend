package services_interface_bot_repository

import (
	models_bot "backend/internal/models/bot"
)

type BotRepositoryService interface {
	Add(*models_bot.BotModel)
	Get(uint) (*models_bot.BotModel, bool)
	Remove(uint)
}
