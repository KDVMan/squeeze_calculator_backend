package services_calculate

import (
	"backend/internal/enums"
	models_calculate "backend/internal/models/calculate"
	services_quote_builder "backend/internal/services/quote_builder"
)

func (object *calculateServiceImplementation) Calculate() *models_calculate.CalculateResultModel {
	quoteBuilderService := services_quote_builder.NewQuoteBuilderService(object.paramModel.Interval, enums.Interval1m)
	var deals []*models_calculate.CalculateDealModel
	var deal *models_calculate.CalculateDealModel
	nextPriceIn := 0.0

	for i, quote := range object.quotes {
		quoteBuild := quoteBuilderService.Build(quote)
		currentPriceIn := nextPriceIn

		if quoteBuild.IsClosed {
			nextPriceIn = object.CalculatePriceIn(quoteBuild)
		}

		if currentPriceIn == 0 {
			continue
		}

		if deal != nil && deal.TimeOut >= quote.TimeOpen {
			continue
		}

		if object.calculatePriceBind(object.direction.MinKeyName, quote) < currentPriceIn {
			deal = object.calculateDeal(i, currentPriceIn)
			deals = append(deals, deal)
		}
	}

	if len(deals) == 0 {
		return nil
	}

	return object.calculateStatistic(deals)
}
