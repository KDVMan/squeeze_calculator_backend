package services_calculate

import (
	models_calculate "backend/internal/models/calculate"
)

func (object *calculateServiceImplementation) calculateStatistic(deals []*models_calculate.CalculateDealModel) *models_calculate.CalculateResultModel {
	result := &models_calculate.CalculateResultModel{
		Total:                        len(deals),
		TotalCumulativeProfitPercent: 100,
		InOutRatio:                   object.paramModel.PercentOut / object.paramModel.PercentIn,
		ParamModel:                   object.paramModel,
		Deals:                        deals,
	}

	var sumTakes float64 = 0
	var sumStops float64 = 0

	for _, deal := range deals {
		result.TotalProfitPercent += deal.ProfitPercent
		result.TotalCumulativeProfitPercent *= 1 + deal.ProfitPercent/100

		if deal.IsStopPercent || deal.IsStopTime {
			result.TotalStops++
		} else {
			result.TotalTakes++
		}

		if deal.ProfitPercent >= 0 {
			sumTakes += deal.ProfitPercent

			if deal.IsStopPercent || deal.IsStopTime {
				result.TotalStopsPlus++
			} else {
				result.TotalTakesPlus++
			}
		} else {
			sumStops += deal.ProfitPercent

			if deal.IsStopPercent || deal.IsStopTime {
				result.TotalStopsMinus++
			}
		}

		if deal.DrawdownPercent > result.MaxDrawdownPercent {
			result.MaxDrawdownPercent = deal.DrawdownPercent
		}

		timeDeal := (deal.TimeOut - deal.TimeIn) / 60000

		if timeDeal > result.MaxTimeDeal {
			result.MaxTimeDeal = timeDeal
		}
	}

	if sumStops != 0 {
		result.Coefficient = -sumTakes / sumStops
	}

	if result.Total > 0 {
		result.WinRate = float64(result.Total-result.TotalStops) / float64(result.Total)
		result.WinRatePlus = float64(result.Total-result.TotalStopsMinus) / float64(result.Total)
	}

	result.TotalCumulativeProfitPercent -= 100

	return result
}
