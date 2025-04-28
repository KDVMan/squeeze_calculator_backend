package services_calculate

import (
	"backend/internal/enums"
	models_quote "backend/internal/models/quote"
	services_helper "backend/pkg/services/helper"
)

func (object *calculateServiceImplementation) CalculatePriceIn(quote *models_quote.QuoteModel) float64 {
	priceBind := object.calculatePriceBind(object.paramModel.Bind, quote)
	priceIn := priceBind * object.priceInFactor

	return services_helper.Floor(priceIn, object.tickSizeFactor)
}

func (object *calculateServiceImplementation) calculatePriceOut(priceIn float64) float64 {
	return services_helper.Floor(priceIn*object.priceOutFactor, object.tickSizeFactor)
}

func (object *calculateServiceImplementation) calculatePriceStop(priceIn float64) float64 {
	if object.priceStopFactor <= 0 {
		return 0
	}

	return services_helper.Floor(priceIn*object.priceStopFactor, object.tickSizeFactor)
}

func (object *calculateServiceImplementation) calculatePriceBind(bind enums.Bind, quote *models_quote.QuoteModel) float64 {
	switch bind {
	case enums.BindLow:
		return quote.PriceLow
	case enums.BindHigh:
		return quote.PriceHigh
	case enums.BindOpen:
		return quote.PriceOpen
	case enums.BindClose:
		return quote.PriceClose
	case enums.BindMhl:
		return (quote.PriceHigh + quote.PriceLow) / 2
	case enums.BindMoc:
		return (quote.PriceOpen + quote.PriceClose) / 2
	default:
		return 0
	}
}
