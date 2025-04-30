package models_bot

import models_calculate "backend/internal/models/calculate"

type CalculateRequestModel struct {
	CalculatorRequestModel *CalculatorRequestModel
	Result                 *models_calculate.CalculateResultModel
}
