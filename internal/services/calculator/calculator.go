package services_calculator

import (
	"backend/internal/enums"
	enums_symbol "backend/internal/enums/symbol"
	enums_websocket "backend/internal/enums/websocket"
	models_calculate "backend/internal/models/calculate"
	models_calculator "backend/internal/models/calculator"
	models_calculator_optimization "backend/internal/models/calculator_optimization"
	models_quote "backend/internal/models/quote"
	models_websocket "backend/internal/models/websocket"
	services_calculate "backend/internal/services/calculate"
	services_calculator_optimization "backend/internal/services/calculator_optimization"
	services_helper "backend/pkg/services/helper"
	"fmt"
	"sync"
)

func (object *calculatorServiceImplementation) Calculator(request *models_calculator.CalculatorRequestModel) error {
	var wg sync.WaitGroup
	object.calculateSymbol = request.Symbol
	object.calculateResult = nil
	timeFrom := request.TimeTo - request.Window*60000

	symbolModel, err := object.symbolService().Load(request.Symbol, enums_symbol.SymbolStatusActive)
	if err != nil {
		return fmt.Errorf("failed to load symbolModel: %w", err)
	}

	calculatorModel, err := object.Update(request)
	if err != nil {
		return fmt.Errorf("failed to update calculatorModel: %w", err)
	}

	optimizations := services_calculator_optimization.Load(&models_calculator_optimization.CalculatorOptimizationRequestModel{
		Bind:            calculatorModel.Bind,
		PercentInFrom:   calculatorModel.PercentInFrom,
		PercentInTo:     calculatorModel.PercentInTo,
		PercentInStep:   calculatorModel.PercentInStep,
		PercentOutFrom:  calculatorModel.PercentOutFrom,
		PercentOutTo:    calculatorModel.PercentOutTo,
		PercentOutStep:  calculatorModel.PercentOutStep,
		StopTime:        calculatorModel.StopTime,
		StopTimeFrom:    calculatorModel.StopTimeFrom,
		StopTimeTo:      calculatorModel.StopTimeTo,
		StopTimeStep:    calculatorModel.StopTimeStep,
		StopPercent:     calculatorModel.StopPercent,
		StopPercentFrom: calculatorModel.StopPercentFrom,
		StopPercentTo:   calculatorModel.StopPercentTo,
		StopPercentStep: calculatorModel.StopPercentStep,
		Algorithm:       calculatorModel.Algorithm,
		Iterations:      calculatorModel.Iterations,
	})

	quoteRange := models_quote.GetRange(int64(object.configService().GetConfig().Binance.FuturesLimit), timeFrom, request.TimeTo, enums.IntervalMilliseconds(enums.Interval1m))

	progressModel := &models_websocket.ProgressChannelModel{
		Count:  0,
		Total:  int64(quoteRange.Iterations+len(optimizations)) + 1,
		Status: enums_websocket.WebsocketStatusProgress,
		Event:  enums_websocket.WebsocketEventCalculateProgress,
	}

	object.websocketService().GetProgressChannel() <- progressModel

	quotes, err := object.quoteService().LoadRange(request.Symbol, quoteRange, progressModel)
	if err != nil {
		return err
	}

	if request.TradeDirection == enums.TradeDirectionShort {
		quotes = models_quote.InvertAll(quotes)
	}

	threadsResults := make(chan *models_calculate.CalculateResultModel, len(optimizations))
	threads := min(services_helper.GetCpu(0), len(optimizations))
	chunkSize := (len(optimizations) + threads - 1) / threads

	worker := func(data []*models_calculator_optimization.CalculatorOptimizationModel) {
		defer wg.Done()

		for _, optimization := range data {
			paramModel := &models_calculate.ParamModel{
				TradeDirection: calculatorModel.TradeDirection,
				Interval:       calculatorModel.Interval,
				Bind:           optimization.Bind,
				PercentIn:      optimization.PercentIn,
				PercentOut:     optimization.PercentOut,
				StopTime:       optimization.StopTime,
				StopPercent:    optimization.StopPercent,
			}

			calculateService := services_calculate.NewCalculateService(paramModel, quotes, symbolModel.Limit.TickSize, object.configService().GetConfig().Binance.FuturesCommission)
			result := calculateService.Calculate()

			if result != nil {
				threadsResults <- result
			}
		}
	}

	for i := 0; i < threads; i++ {
		start := i * chunkSize
		end := start + chunkSize

		if start >= len(optimizations) {
			continue
		}

		if end > len(optimizations) {
			end = len(optimizations)
		}

		wg.Add(1)
		go worker(optimizations[start:end])
	}

	go func() {
		wg.Wait()
		close(threadsResults)
	}()

	for result := range threadsResults {
		object.calculateResult = append(object.calculateResult, result)
		progressModel.Count++

		if progressModel.Count%1000 == 0 {
			object.websocketService().GetProgressChannel() <- progressModel
		}
	}

	progressModel.Count++
	progressModel.Status = enums_websocket.WebsocketStatusDone
	object.websocketService().GetProgressChannel() <- progressModel

	object.websocketService().GetBroadcastChannel() <- &models_websocket.BroadcastChannelModel{
		Event: enums_websocket.WebsocketEventCalculateResult,
		Data:  object.LoadResult(request.Symbol),
	}

	// for _, optimization := range optimizations {
	// 	paramModel := &models_calculate.ParamModel{
	// 		TradeDirection: calculatorModel.TradeDirection,
	// 		Interval:       calculatorModel.Interval,
	// 		Bind:           optimization.Bind,
	// 		PercentIn:      optimization.PercentIn,
	// 		PercentOut:     optimization.PercentOut,
	// 		StopTime:       optimization.StopTime,
	// 		StopPercent:    optimization.StopPercent,
	// 	}
	//
	// 	calculateService := services_calculate.NewCalculateService(paramModel, quotes, symbolModel.Limit.TickSize, object.configService().GetConfig().Binance.FuturesCommission)
	// 	result := calculateService.Calculate()
	//
	// 	if result != nil {
	// 		object.calculateResult = append(object.calculateResult, result)
	// 	}
	//
	// 	progressModel.Count++
	//
	// 	if progressModel.Count%1000 == 0 {
	// 		object.websocketService().GetProgressChannel() <- progressModel
	// 	}
	// }
	//
	// progressModel.Count++
	// progressModel.Status = enums.WebsocketStatusDone
	// object.websocketService().GetProgressChannel() <- progressModel
	//
	// object.websocketService().GetBroadcastChannel() <- &models_channel.BroadcastChannelModel{
	// 	Event: enums.WebsocketEventCalculateResult,
	// 	Data:  object.LoadResult(request.Symbol),
	// }

	return nil
}
