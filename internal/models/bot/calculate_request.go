package models_bot

import models_calculate "backend/internal/models/calculate"

type CalculateRequestModel struct {
	BotID  uint
	Result *models_calculate.CalculateResultModel
}
