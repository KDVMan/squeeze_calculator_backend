package services_calculate

import (
	"backend/internal/enums"
	models_calculate "backend/internal/models/calculate"
	models_quote "backend/internal/models/quote"
	services_interface_calculate "backend/internal/services/calculate/interface"
)

type calculateServiceImplementation struct {
	paramModel      *models_calculate.ParamModel
	quotes          []*models_quote.QuoteModel
	commission      float64
	direction       models_calculate.DirectionModel
	priceInFactor   float64
	priceOutFactor  float64
	priceStopFactor float64
	tickSizeFactor  int
}

func NewCalculateService(
	paramModel *models_calculate.ParamModel,
	quotes []*models_quote.QuoteModel,
	tickSize float64,
	commission float64,
) services_interface_calculate.CalculateService {
	calculatorService := &calculateServiceImplementation{
		paramModel: paramModel,
		quotes:     quotes,
		commission: commission,
	}

	if paramModel.TradeDirection == enums.TradeDirectionShort {
		calculatorService.direction = models_calculate.DirectionModel{Multiplier: -1, MinKeyName: enums.BindHigh, MaxKeyName: enums.BindLow}
	} else {
		calculatorService.direction = models_calculate.DirectionModel{Multiplier: 1, MinKeyName: enums.BindLow, MaxKeyName: enums.BindHigh}
	}

	calculatorService.priceInFactor = (100 - calculatorService.direction.Multiplier*paramModel.PercentIn) / 100
	calculatorService.priceOutFactor = (100 + calculatorService.direction.Multiplier*paramModel.PercentOut) / 100

	if paramModel.StopPercent > 0 {
		calculatorService.priceStopFactor = (100 - calculatorService.direction.Multiplier*paramModel.StopPercent) / 100
	}

	for tickSize < 1 {
		tickSize *= 10
		calculatorService.tickSizeFactor++
	}

	return calculatorService
}
