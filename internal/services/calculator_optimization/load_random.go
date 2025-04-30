package services_calculator_optimization

import (
	"backend/internal/enums"
	models_calculator_optimization "backend/internal/models/calculator_optimization"
	services_helper "backend/pkg/services/helper"
	"fmt"
)

func loadRandom(request *models_calculator_optimization.CalculatorOptimizationRequestModel) map[string]*models_calculator_optimization.CalculatorOptimizationModel {
	optimizations := make(map[string]*models_calculator_optimization.CalculatorOptimizationModel)

	for iteration := 0; iteration < request.Iterations; iteration++ {
		stopTimeCombinations := []int64{-1}
		stopPercentCombinations := []float64{-1}
		bind := enums.BindRandom(request.Bind)

		percentIn := services_helper.GetRandomFloatByInt(
			request.PercentInFrom,
			request.PercentInTo,
			request.PercentInStep,
		)

		percentOut := services_helper.GetRandomFloatByInt(
			request.PercentOutFrom,
			request.PercentOutTo,
			request.PercentOutStep,
		)

		if request.StopTime && request.StopPercent {
			stopTime := services_helper.GetRandomInt(
				request.StopTimeFrom,
				request.StopTimeTo,
				request.StopTimeStep,
			)

			stopTimeCombinations = append(stopTimeCombinations, stopTime*60*1000)

			stopPercent := services_helper.GetRandomFloatByInt(
				request.StopPercentFrom,
				request.StopPercentTo,
				request.StopPercentStep,
			)

			stopPercentCombinations = append(stopPercentCombinations, stopPercent)
		} else if request.StopTime {
			stopTime := services_helper.GetRandomInt(
				request.StopTimeFrom,
				request.StopTimeTo,
				request.StopTimeStep,
			)

			stopTimeCombinations = append(stopTimeCombinations, stopTime*60*1000)
		} else if request.StopPercent {
			stopPercent := services_helper.GetRandomFloatByInt(
				request.StopPercentFrom,
				request.StopPercentTo,
				request.StopPercentStep,
			)

			stopPercentCombinations = append(stopPercentCombinations, stopPercent)
		}

		for _, stopTime := range stopTimeCombinations {
			for _, stopPercent := range stopPercentCombinations {
				if stopTime == -1 && stopPercent == -1 {
					continue
				}

				key := fmt.Sprintf("%s-%g-%g-%d-%g", bind, percentIn, percentOut, stopTime, stopPercent)

				if _, exists := optimizations[key]; !exists {
					optimizations[key] = &models_calculator_optimization.CalculatorOptimizationModel{
						Bind:        bind,
						PercentIn:   percentIn,
						PercentOut:  percentOut,
						StopTime:    stopTime,
						StopPercent: stopPercent,
					}
				}
			}
		}
	}

	return optimizations
}
