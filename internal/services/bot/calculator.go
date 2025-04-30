package services_bot

import (
	models_bot "backend/internal/models/bot"
	models_calculate "backend/internal/models/calculate"
	models_calculator_formula_preset "backend/internal/models/calculator_formula_preset"
	models_calculator_optimization "backend/internal/models/calculator_optimization"
	services_calculate "backend/internal/services/calculate"
	services_calculator "backend/internal/services/calculator"
	services_calculator_optimization "backend/internal/services/calculator_optimization"
	services_helper "backend/pkg/services/helper"
	"log"
	"sort"
	"sync"
)

func (object *botServiceImplementation) CalculatorChannel() {
	for request := range object.calculatorChannel {
		stopChannel, exists := object.stopChannels[request.BotID]
		if !exists {
			continue
		}

		var results []*models_calculate.CalculateResultModel
		ranges := make(map[string][2]float64)

		quotes := object.quoteRepositoryService().GetBySymbol(request.Symbol, request.TradeDirection)
		if len(quotes) == 0 {
			continue
		}

		optimizations := services_calculator_optimization.Load(&models_calculator_optimization.CalculatorOptimizationRequestModel{
			Bind:            request.Bind,
			PercentInFrom:   request.PercentInFrom,
			PercentInTo:     request.PercentInTo,
			PercentInStep:   request.PercentInStep,
			PercentOutFrom:  request.PercentOutFrom,
			PercentOutTo:    request.PercentOutTo,
			PercentOutStep:  request.PercentOutStep,
			StopTime:        request.StopTime,
			StopTimeFrom:    request.StopTimeFrom,
			StopTimeTo:      request.StopTimeTo,
			StopTimeStep:    request.StopTimeStep,
			StopPercent:     request.StopPercent,
			StopPercentFrom: request.StopPercentFrom,
			StopPercentTo:   request.StopPercentTo,
			StopPercentStep: request.StopPercentStep,
			Algorithm:       request.Algorithm,
			Iterations:      request.Iterations,
		})

		type job struct {
			param *models_calculate.ParamModel
		}

		var wg sync.WaitGroup
		jobs := make(chan job)
		resultsChannel := make(chan *models_calculate.CalculateResultModel)
		threads := services_helper.GetCpu(0)

		wg.Add(threads)

		for i := 0; i < threads; i++ {
			go func() {
				defer wg.Done()
				for j := range jobs {
					select {
					case <-stopChannel:
						return
					default:
					}

					calculateService := services_calculate.NewCalculateService(j.param, quotes, request.TickSize, object.futuresCommission)

					if currentResult := calculateService.Calculate(); currentResult != nil {
						resultsChannel <- currentResult
					}
				}
			}()
		}

		go func() {
			for _, optimization := range optimizations {
				jobs <- job{
					&models_calculate.ParamModel{
						TradeDirection: request.TradeDirection,
						Interval:       request.Interval,
						Bind:           optimization.Bind,
						PercentIn:      optimization.PercentIn,
						PercentOut:     optimization.PercentOut,
						StopTime:       optimization.StopTime,
						StopPercent:    optimization.StopPercent,
						IsCurrent:      false,
					},
				}
			}

			if request.Param.PercentIn > 0 {
				jobs <- job{
					&models_calculate.ParamModel{
						TradeDirection: request.TradeDirection,
						Interval:       request.Interval,
						Bind:           request.Param.Bind,
						PercentIn:      request.Param.PercentIn,
						PercentOut:     request.Param.PercentOut,
						StopTime:       request.Param.StopTime,
						StopPercent:    request.Param.StopPercent,
						IsCurrent:      true,
					},
				}
			}

			close(jobs)
		}()

		go func() {
			wg.Wait()
			close(resultsChannel)
		}()

		for calculateResult := range resultsChannel {
			if models_calculator_formula_preset.ApplyFilters(calculateResult, request.Filters) {
				results = append(results, calculateResult)
				services_calculator.UpdateValueRanges(calculateResult, ranges)
			}
		}

		for _, result := range results {
			result.Score = models_calculator_formula_preset.ApplyFormula(result, request.Formulas, ranges)
		}

		sort.Slice(results, func(i, j int) bool {
			return results[i].Score > results[j].Score
		})

		if len(results) > 0 {
			var controlResult *models_calculate.CalculateResultModel

			for _, r := range results {
				if r.ParamModel.IsCurrent {
					controlResult = r
					break
				}
			}

			best := results[0]

			if controlResult == nil || controlResult.Score < best.Score {
				if controlResult != nil {
					log.Printf(
						"CONTROL, symbol: %s, direction: %s, "+
							"score: %.20f -> %.20f, "+
							"profit: %.2f -> %.2f, "+
							"bind: %v -> %v, "+
							"in: %.2f -> %.2f, "+
							"out: %.2f -> %.2f, "+
							"stopTime: %v -> %v, "+
							"stopPercent: %.2f -> %.2f, "+
							"current: %v -> %v\n\n",
						request.Symbol,
						request.TradeDirection,
						controlResult.Score, best.Score,
						controlResult.TotalCumulativeProfitPercent, best.TotalCumulativeProfitPercent,
						controlResult.ParamModel.Bind, best.ParamModel.Bind,
						controlResult.ParamModel.PercentIn, best.ParamModel.PercentIn,
						controlResult.ParamModel.PercentOut, best.ParamModel.PercentOut,
						controlResult.ParamModel.StopTime, best.ParamModel.StopTime,
						controlResult.ParamModel.StopPercent, best.ParamModel.StopPercent,
						controlResult.ParamModel.IsCurrent, best.ParamModel.IsCurrent,
					)
				}

				object.GetCalculateChannel() <- &models_bot.CalculateRequestModel{
					CalculatorRequestModel: request,
					Result:                 best,
				}
			}
		}

		object.GetCalculatorChannel() <- request
	}
}

// func (object *botServiceImplementation) CalculatorChannel() {
// 	for request := range object.calculatorChannel {
// 		stopChannel, exists := object.stopChannels[request.BotID]
// 		if !exists {
// 			continue
// 		}
//
// 		var calculateResults []*models_calculate.CalculateResultModel
// 		var results []*models_calculate.CalculateResultModel
// 		ranges := make(map[string][2]float64)
//
// 		quotes := object.quoteRepositoryService().GetBySymbol(request.Symbol, request.TradeDirection)
// 		if len(quotes) == 0 {
// 			continue
// 		}
//
// 		optimizations := services_calculator_optimization.Load(&models_calculator_optimization.CalculatorOptimizationRequestModel{
// 			Bind:            request.Bind,
// 			PercentInFrom:   request.PercentInFrom,
// 			PercentInTo:     request.PercentInTo,
// 			PercentInStep:   request.PercentInStep,
// 			PercentOutFrom:  request.PercentOutFrom,
// 			PercentOutTo:    request.PercentOutTo,
// 			PercentOutStep:  request.PercentOutStep,
// 			StopTime:        request.StopTime,
// 			StopTimeFrom:    request.StopTimeFrom,
// 			StopTimeTo:      request.StopTimeTo,
// 			StopTimeStep:    request.StopTimeStep,
// 			StopPercent:     request.StopPercent,
// 			StopPercentFrom: request.StopPercentFrom,
// 			StopPercentTo:   request.StopPercentTo,
// 			StopPercentStep: request.StopPercentStep,
// 			Algorithm:       request.Algorithm,
// 			Iterations:      request.Iterations,
// 		})
//
// 		for _, optimization := range optimizations {
// 			select {
// 			case <-stopChannel:
// 				continue
// 			default:
// 			}
//
// 			paramModel := &models_calculate.ParamModel{
// 				TradeDirection: request.TradeDirection,
// 				Interval:       request.Interval,
// 				Bind:           optimization.Bind,
// 				PercentIn:      optimization.PercentIn,
// 				PercentOut:     optimization.PercentOut,
// 				StopTime:       optimization.StopTime,
// 				StopPercent:    optimization.StopPercent,
// 				IsCurrent:      false,
// 			}
//
// 			calculateService := services_calculate.NewCalculateService(paramModel, quotes, request.TickSize, object.futuresCommission)
// 			currentResult := calculateService.Calculate()
//
// 			if currentResult != nil {
// 				calculateResults = append(calculateResults, currentResult)
// 			}
// 		}
//
// 		if request.Param.PercentIn > 0 {
// 			currentParam := &models_calculate.ParamModel{
// 				TradeDirection: request.TradeDirection,
// 				Interval:       request.Interval,
// 				Bind:           request.Param.Bind,
// 				PercentIn:      request.Param.PercentIn,
// 				PercentOut:     request.Param.PercentOut,
// 				StopTime:       request.Param.StopTime,
// 				StopPercent:    request.Param.StopPercent,
// 				IsCurrent:      true,
// 			}
//
// 			calculateService := services_calculate.NewCalculateService(currentParam, quotes, request.TickSize, object.futuresCommission)
// 			currentResult := calculateService.Calculate()
//
// 			if currentResult != nil {
// 				calculateResults = append(calculateResults, currentResult)
// 			}
// 		}
//
// 		for _, calculateResult := range calculateResults {
// 			if models_calculator_formula_preset.ApplyFilters(calculateResult, request.Filters) {
// 				results = append(results, calculateResult)
// 				services_calculator.UpdateValueRanges(calculateResult, ranges)
// 			}
// 		}
//
// 		for _, result := range results {
// 			result.Score = models_calculator_formula_preset.ApplyFormula(result, request.Formulas, ranges)
// 		}
//
// 		sort.Slice(results, func(i, j int) bool {
// 			return results[i].Score > results[j].Score
// 		})
//
// 		if len(results) > 0 {
// 			var controlResult *models_calculate.CalculateResultModel
//
// 			for _, r := range results {
// 				if r.ParamModel.IsCurrent {
// 					controlResult = r
// 					break
// 				}
// 			}
//
// 			best := results[0]
//
// 			if controlResult == nil || controlResult.Score < best.Score {
// 				if controlResult != nil {
// 					log.Printf(
// 						"CONTROL, "+
// 							"score: %.20f -> %.20f, "+
// 							"profit: %.2f -> %.2f, "+
// 							"bind: %v -> %v, "+
// 							"in: %.2f -> %.2f, "+
// 							"out: %.2f -> %.2f, "+
// 							"stopTime: %v -> %v, "+
// 							"stopPercent: %.2f -> %.2f, "+
// 							"current: %v -> %v\n\n",
// 						controlResult.Score, best.Score,
// 						controlResult.TotalCumulativeProfitPercent, best.TotalCumulativeProfitPercent,
// 						controlResult.ParamModel.Bind, best.ParamModel.Bind,
// 						controlResult.ParamModel.PercentIn, best.ParamModel.PercentIn,
// 						controlResult.ParamModel.PercentOut, best.ParamModel.PercentOut,
// 						controlResult.ParamModel.StopTime, best.ParamModel.StopTime,
// 						controlResult.ParamModel.StopPercent, best.ParamModel.StopPercent,
// 						controlResult.ParamModel.IsCurrent, best.ParamModel.IsCurrent,
// 					)
// 				}
//
// 				object.GetCalculateChannel() <- &models_bot.CalculateRequestModel{
// 					CalculatorRequestModel: request,
// 					Result:                 best,
// 				}
// 			}
// 		}
//
// 		object.GetCalculatorChannel() <- request
// 	}
// }

func (object *botServiceImplementation) GetCalculatorChannel() chan *models_bot.CalculatorRequestModel {
	return object.calculatorChannel
}
