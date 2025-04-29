package services_bot

import (
	enums_bot "backend/internal/enums/bot"
	enums_symbol "backend/internal/enums/symbol"
	models_bot "backend/internal/models/bot"
	services_helper "backend/pkg/services/helper"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func (object *botServiceImplementation) Start(request *models_bot.StartRequestModel) error {
	var botModel models_bot.BotModel

	calculatorPresetModel, err := object.calculatorPresetService().LoadSelected()
	if err != nil {
		return err
	}

	calculatorFormulaPresetModel, err := object.calculatorFormulaPresetService().LoadSelected()
	if err != nil {
		return err
	}

	hash := services_helper.MustConvertStringToMd5(fmt.Sprintf(
		"hash | calculatorPresetModel:%d | calculatorFormulaPresetModel:%d | window:%d | direction:%s |  interval:%s | symbol:%s | ",
		calculatorPresetModel.ID,
		calculatorFormulaPresetModel.ID,
		calculatorPresetModel.Window,
		calculatorPresetModel.TradeDirection,
		calculatorPresetModel.Interval,
		request.Symbol,
	))

	err = object.storageService().DB().Where("hash = ?", hash).First(&botModel).Error
	if err == nil {
		return nil
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	symbolModel, err := object.symbolService().Load(request.Symbol, enums_symbol.SymbolStatusActive)
	if err != nil {
		return err
	}

	botModel = models_bot.BotModel{
		CalculatorPresetID:        calculatorPresetModel.ID,
		CalculatorFormulaPresetID: calculatorFormulaPresetModel.ID,
		Hash:                      hash,
		Symbol:                    request.Symbol,
		Window:                    calculatorPresetModel.Window,
		TradeDirection:            calculatorPresetModel.TradeDirection,
		Interval:                  calculatorPresetModel.Interval,
		Bind:                      calculatorPresetModel.Bind,
		PercentInFrom:             calculatorPresetModel.PercentInFrom,
		PercentInTo:               calculatorPresetModel.PercentInTo,
		PercentInStep:             calculatorPresetModel.PercentInStep,
		PercentOutFrom:            calculatorPresetModel.PercentOutFrom,
		PercentOutTo:              calculatorPresetModel.PercentOutTo,
		PercentOutStep:            calculatorPresetModel.PercentOutStep,
		StopTime:                  calculatorPresetModel.StopTime,
		StopTimeFrom:              calculatorPresetModel.StopTimeFrom,
		StopTimeTo:                calculatorPresetModel.StopTimeTo,
		StopTimeStep:              calculatorPresetModel.StopTimeStep,
		StopPercent:               calculatorPresetModel.StopPercent,
		StopPercentFrom:           calculatorPresetModel.StopPercentFrom,
		StopPercentTo:             calculatorPresetModel.StopPercentTo,
		StopPercentStep:           calculatorPresetModel.StopPercentStep,
		Algorithm:                 calculatorPresetModel.Algorithm,
		Status:                    enums_bot.StatusStart,
		Filters:                   calculatorFormulaPresetModel.Filters,
		Formulas:                  calculatorFormulaPresetModel.Formulas,
		TickSize:                  symbolModel.Limit.TickSize,
		Param:                     models_bot.ParamModel{},
	}

	if err = object.storageService().DB().Create(&botModel).Error; err != nil {
		return err
	}

	// currentTime := time.Now().UnixMilli()
	// timeFrom := currentTime - (tradeModel.Window * 60 * 1000)
	// timeTo := currentTime
	// quoteRange := models_quote.GetRange(int64(object.configService().GetConfig().Binance.FuturesLimit), timeFrom, timeTo, enums.IntervalMilliseconds(enums.Interval1m))
	//
	// quotes, err := object.quoteService().LoadRange(request.Symbol, quoteRange, &models_channel.ProgressChannelModel{})
	// if err != nil {
	// 	return err
	// }
	//
	// object.quoteRepositoryService().Add(tradeModel.Symbol, quotes)
	// object.tradeRepository().Add(&tradeModel)
	// object.exchangeWebsocketService().SubscribeTrade(request.Symbol)

	// object.websocketService().GetBroadcastChannel() <- &models_channel.BroadcastChannelModel{
	// 	Event: enums.WebsocketEventExecTrade,
	// 	Data:  object.tradeRepository().GetAll(),
	// }

	return nil
}
