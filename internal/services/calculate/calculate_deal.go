package services_calculate

import (
	"backend/internal/enums"
	enums_calculate "backend/internal/enums/calculate"
	models_calculate "backend/internal/models/calculate"
)

func (object *calculateServiceImplementation) calculateDeal(index int, currentPriceIn float64) *models_calculate.CalculateDealModel {
	deal := &models_calculate.CalculateDealModel{
		TimeIn:        object.quotes[index].TimeOpen,
		TimeOut:       0,
		PriceIn:       currentPriceIn,
		PriceOut:      0,
		MinPrice:      currentPriceIn,
		ProfitPercent: 0,
	}

	priceOut := object.calculatePriceOut(currentPriceIn)
	priceStop := object.calculatePriceStop(currentPriceIn)

	check := object.checkDeal(object.quotes[index], deal, priceOut, priceStop, enums.BindClose)
	index++

	for check == false && index < len(object.quotes) {
		check = object.checkDeal(object.quotes[index], deal, priceOut, priceStop, object.direction.MaxKeyName)
		index++
	}

	if check == false {
		object.buildDeal(deal, object.quotes[len(object.quotes)-1].PriceClose, object.quotes[len(object.quotes)-1].TimeClose, enums_calculate.DealTypeOpen)
	}

	return deal
}
