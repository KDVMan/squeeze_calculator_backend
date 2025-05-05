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
	"time"
)

func (object *botServiceImplementation) CalculatorChannel() {
	for botID := range object.calculatorChannel {
		object.calculatorStop.Store(false)

		botModel, exists := object.botRepositoryService().Get(botID)
		if !exists {
			continue
		}

		now := time.Now().UnixMilli()

		if tsRaw, ok := object.delayedCoins.Load(botID); ok {
			if ts, ok := tsRaw.(int64); ok && now < ts {
				continue
			}

			object.delayedCoins.Delete(botID)
		}

		stopChannel, exists := object.stopChannels[botID]
		if !exists {
			continue
		}

		quotes := object.quoteRepositoryService().GetWindowBySymbol(botModel.Symbol, botModel.TradeDirection, int(botModel.Window))
		if len(quotes) == 0 {
			continue
		}

		optimizations := services_calculator_optimization.Load(&models_calculator_optimization.CalculatorOptimizationRequestModel{
			Bind:            botModel.Bind,
			PercentInFrom:   botModel.PercentInFrom,
			PercentInTo:     botModel.PercentInTo,
			PercentInStep:   botModel.PercentInStep,
			PercentOutFrom:  botModel.PercentOutFrom,
			PercentOutTo:    botModel.PercentOutTo,
			PercentOutStep:  botModel.PercentOutStep,
			StopTime:        botModel.StopTime,
			StopTimeFrom:    botModel.StopTimeFrom,
			StopTimeTo:      botModel.StopTimeTo,
			StopTimeStep:    botModel.StopTimeStep,
			StopPercent:     botModel.StopPercent,
			StopPercentFrom: botModel.StopPercentFrom,
			StopPercentTo:   botModel.StopPercentTo,
			StopPercentStep: botModel.StopPercentStep,
			Algorithm:       botModel.Algorithm,
			Iterations:      botModel.Iterations,
		})

		type job struct {
			param *models_calculate.ParamModel
		}

		var wg sync.WaitGroup
		jobs := make(chan job)
		resultsChannel := make(chan *models_calculate.CalculateResultModel)
		threads := services_helper.GetCpu(2)

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

					calculateService := services_calculate.NewCalculateService(j.param, quotes, botModel.TickSize, object.futuresCommission)

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
						TradeDirection: botModel.TradeDirection,
						Interval:       botModel.Interval,
						Bind:           optimization.Bind,
						PercentIn:      optimization.PercentIn,
						PercentOut:     optimization.PercentOut,
						StopTime:       optimization.StopTime,
						StopPercent:    optimization.StopPercent,
						IsCurrent:      false,
					},
				}
			}

			if botModel.Param.PercentIn > 0 {
				select {
				case <-stopChannel:
					object.calculatorStop.Store(true)
					close(jobs)
					return
				default:
				}

				jobs <- job{
					&models_calculate.ParamModel{
						TradeDirection: botModel.TradeDirection,
						Interval:       botModel.Interval,
						Bind:           botModel.Param.Bind,
						PercentIn:      botModel.Param.PercentIn,
						PercentOut:     botModel.Param.PercentOut,
						StopTime:       botModel.Param.StopTime,
						StopPercent:    botModel.Param.StopPercent,
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

		var results []*models_calculate.CalculateResultModel
		ranges := make(map[string][2]float64)

		for calculateResult := range resultsChannel {
			if object.calculatorStop.Load() {
				break
			}

			if models_calculator_formula_preset.ApplyFilters(calculateResult, botModel.Filters) {
				results = append(results, calculateResult)
				services_calculator.UpdateValueRanges(calculateResult, ranges)
			}
		}

		for _, result := range results {
			if object.calculatorStop.Load() {
				break
			}

			result.Score = models_calculator_formula_preset.ApplyFormula(result, botModel.Formulas, ranges)
		}

		if object.calculatorStop.Load() {
			continue
		}

		// if object.test {
		// 	results = nil
		// 	log.Println()
		// }

		if len(results) > 0 {
			sort.Slice(results, func(i, j int) bool {
				return results[i].Score > results[j].Score
			})

			bestResult := results[0]
			botModel.IsFirstRun = false
			botModel.IsEmptySend = false

			if bestResult.Score > botModel.Param.Score {
				object.GetCalculateChannel() <- &models_bot.CalculateRequestModel{
					BotID:  botID,
					Result: bestResult,
				}
			}
		} else if botModel.IsEmptySend == false {
			botModel.IsFirstRun = false
			botModel.IsEmptySend = true

			object.GetCalculateChannel() <- &models_bot.CalculateRequestModel{
				BotID:  botID,
				Result: nil,
			}
		}

		if len(results) == 0 {
			delay := 20 * time.Second
			object.delayedCoins.Store(botID, time.Now().Add(delay).UnixMilli())

			go func() {
				time.Sleep(delay)
				object.GetCalculatorChannel() <- botID
			}()
		} else {
			object.GetCalculatorChannel() <- botID
		}
	}
}

func (object *botServiceImplementation) GetCalculatorChannel() chan uint {
	return object.calculatorChannel
}
