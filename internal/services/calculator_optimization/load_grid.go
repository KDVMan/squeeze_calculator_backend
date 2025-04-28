package services_calculator_optimization

import (
	models_calculator_optimization "backend/internal/models/calculator_optimization"
	models_calculator_preset "backend/internal/models/calculator_preset"
	services_helper "backend/pkg/services/helper"
	"fmt"
)

func loadGrid(calculatorModel *models_calculator_preset.CalculatorPresetModel) map[string]*models_calculator_optimization.CalculatorOptimizationModel {
	optimizations := make(map[string]*models_calculator_optimization.CalculatorOptimizationModel)
	iteration := 0

	percentInMin, percentInMax, percentInStep, percentInAccuracy := services_helper.GetRangeFloatByInt(
		calculatorModel.PercentInFrom,
		calculatorModel.PercentInTo,
		calculatorModel.PercentInStep,
	)

	percentOutMin, percentOutMax, percentOutStep, percentOutAccuracy := services_helper.GetRangeFloatByInt(
		calculatorModel.PercentOutFrom,
		calculatorModel.PercentOutTo,
		calculatorModel.PercentOutStep,
	)

	stopPercentMin, stopPercentMax, stopPercentStep, stopPercentAccuracy := services_helper.GetRangeFloatByInt(
		calculatorModel.StopPercentFrom,
		calculatorModel.StopPercentTo,
		calculatorModel.StopPercentStep,
	)

	stopTimeOptions := []int64{-1}
	stopPercentOptions := []int64{-1}

	if calculatorModel.StopTime {
		stopTimeOptions = []int64{}

		for _, value := range services_helper.GenerateRangeByStep(
			calculatorModel.StopTimeFrom,
			calculatorModel.StopTimeTo,
			calculatorModel.StopTimeStep,
		) {
			stopTimeOptions = append(stopTimeOptions, value)
		}
	}

	if calculatorModel.StopPercent {
		stopPercentOptions = []int64{}

		for _, value := range services_helper.GenerateRangeByStep(stopPercentMin, stopPercentMax, stopPercentStep) {
			stopPercentOptions = append(stopPercentOptions, value)
		}
	}

	for _, bind := range calculatorModel.Bind {
		for percentIn := percentInMin; percentIn <= percentInMax; percentIn += percentInStep {
			for percentOut := percentOutMin; percentOut <= percentOutMax; percentOut += percentOutStep {
				for _, stopTime := range stopTimeOptions {
					adjustedStopTime := stopTime

					if adjustedStopTime > 0 {
						adjustedStopTime = adjustedStopTime * 60 * 1000
					}

					for _, stopPercent := range stopPercentOptions {
						key := fmt.Sprintf("%s-%d-%d-%d-%d", bind, percentIn, percentOut, stopTime, stopPercent)

						adjustedStopPercent := float64(stopPercent)

						if adjustedStopPercent > 0 {
							adjustedStopPercent = adjustedStopPercent / float64(stopPercentAccuracy)
						}

						if _, exists := optimizations[key]; !exists {
							optimizations[key] = &models_calculator_optimization.CalculatorOptimizationModel{
								Bind:        bind,
								PercentIn:   float64(percentIn) / float64(percentInAccuracy),
								PercentOut:  float64(percentOut) / float64(percentOutAccuracy),
								StopTime:    adjustedStopTime,
								StopPercent: adjustedStopPercent,
							}
						}

						iteration++

						if iteration >= calculatorModel.Iterations {
							return optimizations
						}
					}
				}
			}
		}
	}

	return optimizations
}
