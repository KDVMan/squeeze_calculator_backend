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
	"sort"
	"sync"
)

func (object *botServiceImplementation) CalculatorChannel() {
	for request := range object.calculatorChannel {
		object.calculatorStop.Store(false)

		stopChannel, exists := object.stopChannels[request.BotID]
		if !exists {
			continue
		}

		var results []*models_calculate.CalculateResultModel
		ranges := make(map[string][2]float64)

		quotes := object.quoteRepositoryService().GetWindowBySymbol(request.Symbol, request.TradeDirection, int(request.Window))
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
						object.calculatorStop.Store(true)
						return
					default:
					}

					calculateService := services_calculate.NewCalculateService(j.param, quotes, request.TickSize, object.futuresCommission)

					if currentResult := calculateService.Calculate(); currentResult != nil {
						select {
						case <-stopChannel:
							object.calculatorStop.Store(true)
							return
						default:
							resultsChannel <- currentResult
						}
					}
				}
			}()
		}

		go func() {
			for _, optimization := range optimizations {
				select {
				case <-stopChannel:
					object.calculatorStop.Store(true)
					close(jobs)
					return
				default:
				}

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
				select {
				case <-stopChannel:
					object.calculatorStop.Store(true)
					close(jobs)
					return
				default:
				}

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
			if object.calculatorStop.Load() {
				break
			}

			if models_calculator_formula_preset.ApplyFilters(calculateResult, request.Filters) {
				results = append(results, calculateResult)
				services_calculator.UpdateValueRanges(calculateResult, ranges)
			}
		}

		for _, result := range results {
			if object.calculatorStop.Load() {
				break
			}

			result.Score = models_calculator_formula_preset.ApplyFormula(result, request.Formulas, ranges)
		}

		if object.calculatorStop.Load() == false {
			if len(results) > 0 {
				sort.Slice(results, func(i, j int) bool {
					return results[i].Score > results[j].Score
				})

				best := results[0]
				var controlResult *models_calculate.CalculateResultModel

				for _, r := range results {
					if r.ParamModel.IsCurrent {
						controlResult = r
						break
					}
				}

				if controlResult == nil || controlResult.Score < best.Score {
					request.CanSendParam = true

					object.GetCalculateChannel() <- &models_bot.CalculateRequestModel{
						CalculatorRequestModel: request,
						Result:                 best,
					}
				}
			} else if request.CanSendParam {
				request.CanSendParam = false

				object.GetCalculateChannel() <- &models_bot.CalculateRequestModel{
					CalculatorRequestModel: request,
					Result:                 nil,
				}
			}

			object.GetCalculatorChannel() <- request
		}
	}
}

func (object *botServiceImplementation) GetCalculatorChannel() chan *models_bot.CalculatorRequestModel {
	return object.calculatorChannel
}
