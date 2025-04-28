package services_interface_calculate

import (
	models_calculate "backend/internal/models/calculate"
	models_quote "backend/internal/models/quote"
)

type CalculateService interface {
	Calculate() *models_calculate.CalculateResultModel
	CalculatePriceIn(*models_quote.QuoteModel) float64
}
