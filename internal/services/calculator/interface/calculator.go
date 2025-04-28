package services_interface_calculator

import (
	models_calculate "backend/internal/models/calculate"
	models_calculator "backend/internal/models/calculator"
	models_calculator_preset "backend/internal/models/calculator_preset"
)

type CalculatorService interface {
	Update(*models_calculator.CalculateRequestModel) (*models_calculator_preset.CalculatorPresetModel, error)
	Calculate(*models_calculator.CalculateRequestModel) error
	LoadResult(string) []*models_calculate.CalculateResultModel
}
