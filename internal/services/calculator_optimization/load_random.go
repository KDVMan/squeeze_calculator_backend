package services_calculator_optimization

import (
	"backend/internal/enums"
	models_calculator_optimization "backend/internal/models/calculator_optimization"
	models_calculator_preset "backend/internal/models/calculator_preset"
	services_helper "backend/pkg/services/helper"
	"fmt"
)

func loadRandom(calculatorModel *models_calculator_preset.CalculatorPresetModel) map[string]*models_calculator_optimization.CalculatorOptimizationModel {
	optimizations := make(map[string]*models_calculator_optimization.CalculatorOptimizationModel)

	for iteration := 0; iteration < calculatorModel.Iterations; iteration++ {
		stopTimeCombinations := []int64{-1}
		stopPercentCombinations := []float64{-1}
		bind := enums.BindRandom(calculatorModel.Bind)

		percentIn := services_helper.GetRandomFloatByInt(
			calculatorModel.PercentInFrom,
			calculatorModel.PercentInTo,
			calculatorModel.PercentInStep,
		)

		percentOut := services_helper.GetRandomFloatByInt(
			calculatorModel.PercentOutFrom,
			calculatorModel.PercentOutTo,
			calculatorModel.PercentOutStep,
		)

		if calculatorModel.StopTime && calculatorModel.StopPercent {
			stopTime := services_helper.GetRandomInt(
				calculatorModel.StopTimeFrom,
				calculatorModel.StopTimeTo,
				calculatorModel.StopTimeStep,
			)

			stopTimeCombinations = append(stopTimeCombinations, stopTime*60*1000)

			stopPercent := services_helper.GetRandomFloatByInt(
				calculatorModel.StopPercentFrom,
				calculatorModel.StopPercentTo,
				calculatorModel.StopPercentStep,
			)

			stopPercentCombinations = append(stopPercentCombinations, stopPercent)
		} else if calculatorModel.StopTime {
			stopTime := services_helper.GetRandomInt(
				calculatorModel.StopTimeFrom,
				calculatorModel.StopTimeTo,
				calculatorModel.StopTimeStep,
			)

			stopTimeCombinations = append(stopTimeCombinations, stopTime*60*1000)
		} else if calculatorModel.StopPercent {
			stopPercent := services_helper.GetRandomFloatByInt(
				calculatorModel.StopPercentFrom,
				calculatorModel.StopPercentTo,
				calculatorModel.StopPercentStep,
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
