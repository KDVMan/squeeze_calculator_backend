package services_interface_bot

import (
	models_bot "backend/internal/models/bot"
)

type BotService interface {
	Init()
	Start(*models_bot.StartRequestModel) error
	LoadAll() []*models_bot.BotModel
	LoadByID(uint) *models_bot.BotModel
	RunChannel()
	GetRunChannel() chan *models_bot.BotModel
	UpdateStatus(*models_bot.UpdateStatusRequestModel) error
	CalculatorChannel()
	GetCalculatorChannel() chan *models_bot.CalculatorRequestModel
	CalculateChannel()
	GetCalculateChannel() chan *models_bot.CalculateRequestModel
	StopBot(*models_bot.BotModel)
}
