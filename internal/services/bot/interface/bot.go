package services_interface_bot

import models_bot "backend/internal/models/bot"

type BotService interface {
	// Init() error
	Start(*models_bot.StartRequestModel) error
	// Calculate()
	// RunDealChannel()
	// GetDealChannel() chan string
}
