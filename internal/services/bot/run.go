package services_bot

import (
	"backend/internal/enums"
	enums_bot "backend/internal/enums/bot"
	enums_websocket "backend/internal/enums/websocket"
	models_bot "backend/internal/models/bot"
	models_quote "backend/internal/models/quote"
	models_websocket "backend/internal/models/websocket"
	services_helper "backend/pkg/services/helper"
)

func (object *botServiceImplementation) RunChannel() {
	for botModel := range object.runChannel {
		timeFrom, timeTo := models_quote.GetTimeRange(botModel)
		quoteRange := models_quote.GetRange(object.futuresLimit, timeFrom, timeTo, enums.IntervalMilliseconds(enums.Interval1m))

		quotes, err := object.quoteService().LoadRange(botModel.Symbol, quoteRange, &models_websocket.ProgressChannelModel{})
		if err != nil {
			object.loggerService().Error().Printf("failed to load range: %v", err)
			continue
		}

		object.quoteRepositoryService().Add(botModel.Symbol, quotes)
		object.exchangeWebsocketService().SubscribeSymbol(botModel.Symbol)

		botModel.Status = enums_bot.StatusStart

		if err = object.storageService().DB().Save(&botModel).Error; err != nil {
			object.loggerService().Error().Printf("failed to save bot: %v", err)
			continue
		}

		object.websocketService().GetBroadcastChannel() <- &models_websocket.BroadcastChannelModel{
			Event: enums_websocket.WebsocketEventBot,
			Data:  object.LoadByID(botModel.ID),
		}

		object.GetCalculatorChannel() <- &models_bot.CalculatorRequestModel{
			BotID:           botModel.ID,
			Symbol:          botModel.Symbol,
			Window:          botModel.Window,
			TradeDirection:  botModel.TradeDirection,
			Interval:        botModel.Interval,
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
			Iterations:      services_helper.CalculateOptimalIterations(botModel.Window, services_helper.GetCpu(2), 0.05),
			TickSize:        botModel.TickSize,
			MinAmount:       botModel.MinAmount,
			Filters:         botModel.Filters,
			Formulas:        botModel.Formulas,
			Param:           botModel.Param,
			IsFirstRun:      botModel.IsFirstRun,
			IsEmptySend:     false,
		}
	}
}

func (object *botServiceImplementation) GetRunChannel() chan *models_bot.BotModel {
	return object.runChannel
}
