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
	GetRunChannel() chan uint
	UpdateStatus(*models_bot.UpdateStatusRequestModel) error
	CalculatorChannel()
	GetCalculatorChannel() chan uint
	CalculateChannel()
	GetCalculateChannel() chan *models_bot.CalculateRequestModel
	StopBot(*models_bot.BotModel)
	Update(*models_bot.UpdateRequestModel) error
	SetTest()
}
