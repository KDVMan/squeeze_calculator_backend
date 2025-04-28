package services_calculate

import (
	"backend/internal/enums"
	enums_calculate "backend/internal/enums/calculate"
	models_calculate "backend/internal/models/calculate"
	models_quote "backend/internal/models/quote"
)

func (object *calculateServiceImplementation) checkDeal(
	quote *models_quote.QuoteModel,
	deal *models_calculate.CalculateDealModel,
	priceOut float64,
	priceStop float64,
	bind enums.Bind,
) bool {
	minPrice := object.calculatePriceBind(object.direction.MinKeyName, quote)

	if minPrice < deal.MinPrice {
		deal.MinPrice = minPrice
	}

	if priceStop != 0 {
		if object.calculatePriceBind(object.direction.MinKeyName, quote) <= priceStop {
			return object.buildDeal(deal, priceStop, quote.TimeClose, enums_calculate.DealTypeStopPercent)
		}
	}

	if object.paramModel.StopTime > 0 && (quote.TimeClose-deal.TimeIn > object.paramModel.StopTime) {
		return object.buildDeal(deal, minPrice, quote.TimeClose, enums_calculate.DealTypeStopTime)
	}

	if priceOut < object.calculatePriceBind(bind, quote) {
		return object.buildDeal(deal, priceOut, quote.TimeClose, enums_calculate.DealTypeProfit)
	}

	return false
}

func (object *calculateServiceImplementation) buildDeal(
	deal *models_calculate.CalculateDealModel,
	price float64,
	time int64,
	dealType enums_calculate.DealType,
) bool {
	var drawdownPercent float64 = 0

	if dealType == enums_calculate.DealTypeStopPercent {
		drawdownPercent = object.paramModel.StopPercent
	} else {
		drawdownPercent = object.direction.Multiplier * (deal.PriceIn - deal.MinPrice) / deal.PriceIn * 100
	}

	deal.PriceIn *= object.direction.Multiplier
	deal.PriceOut = object.direction.Multiplier * price
	deal.TimeOut = time
	deal.DrawdownPercent = drawdownPercent

	if dealType == enums_calculate.DealTypeStopPercent {
		deal.IsStopPercent = true
	} else if dealType == enums_calculate.DealTypeStopTime {
		deal.IsStopTime = true
	}

	if object.paramModel.TradeDirection == enums.TradeDirectionShort {
		deal.ProfitPercent = 100/deal.PriceOut*deal.PriceIn - 100 - (100+100/deal.PriceOut*deal.PriceIn)*object.commission/100
	} else {
		deal.ProfitPercent = 100/deal.PriceIn*deal.PriceOut - 100 - (100+100/deal.PriceIn*deal.PriceOut)*object.commission/100
	}

	return true
}
