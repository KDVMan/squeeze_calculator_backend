package services_calculator_optimization

import (
	models_calculator_optimization "backend/internal/models/calculator_optimization"
	services_helper "backend/pkg/services/helper"
	"fmt"
)

func loadGrid(request *models_calculator_optimization.CalculatorOptimizationRequestModel) map[string]*models_calculator_optimization.CalculatorOptimizationModel {
	optimizations := make(map[string]*models_calculator_optimization.CalculatorOptimizationModel)
	iteration := 0

	percentInMin, percentInMax, percentInStep, percentInAccuracy := services_helper.GetRangeFloatByInt(
		request.PercentInFrom,
		request.PercentInTo,
		request.PercentInStep,
	)

	percentOutMin, percentOutMax, percentOutStep, percentOutAccuracy := services_helper.GetRangeFloatByInt(
		request.PercentOutFrom,
		request.PercentOutTo,
		request.PercentOutStep,
	)

	stopPercentMin, stopPercentMax, stopPercentStep, stopPercentAccuracy := services_helper.GetRangeFloatByInt(
		request.StopPercentFrom,
		request.StopPercentTo,
		request.StopPercentStep,
	)

	stopTimeOptions := []int64{-1}
	stopPercentOptions := []int64{-1}

	if request.StopTime {
		stopTimeOptions = []int64{}

		for _, value := range services_helper.GenerateRangeByStep(
			request.StopTimeFrom,
			request.StopTimeTo,
			request.StopTimeStep,
		) {
			stopTimeOptions = append(stopTimeOptions, value)
		}
	}

	if request.StopPercent {
		stopPercentOptions = []int64{}

		for _, value := range services_helper.GenerateRangeByStep(stopPercentMin, stopPercentMax, stopPercentStep) {
			stopPercentOptions = append(stopPercentOptions, value)
		}
	}

	for _, bind := range request.Bind {
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

						if iteration >= request.Iterations {
							return optimizations
						}
					}
				}
			}
		}
	}

	return optimizations
}
