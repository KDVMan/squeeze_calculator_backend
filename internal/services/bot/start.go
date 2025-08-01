package services_bot

import (
	enums_bot "backend/internal/enums/bot"
	enums_symbol "backend/internal/enums/symbol"
	enums_websocket "backend/internal/enums/websocket"
	models_bot "backend/internal/models/bot"
	models_calculator_formula_preset "backend/internal/models/calculator_formula_preset"
	models_calculator_preset "backend/internal/models/calculator_preset"
	models_symbol "backend/internal/models/symbol"
	models_websocket "backend/internal/models/websocket"
	services_helper "backend/pkg/services/helper"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strings"
)

func (object *botServiceImplementation) Start(request *models_bot.StartRequestModel) error {
	ignoreSymbols := map[string]struct{}{
		"BTCUSDT":        {},
		"ETHUSDT":        {},
		"BTCUSDT_250926": {},
		"BTCUSDT_250627": {},
	}

	calculatorPresetModel, err := object.calculatorPresetService().LoadSelected()
	if err != nil {
		return err
	}

	calculatorFormulaPresetModel, err := object.calculatorFormulaPresetService().LoadSelected()
	if err != nil {
		return err
	}

	if request.IsMass {
		symbolList, err := object.symbolListService().Load()
		if err != nil {
			return err
		}

		symbolsModels, err := object.symbolService().LoadByVolume(symbolList.Volume, "USDT")
		if err != nil {
			return err
		}

		for _, symbolModel := range symbolsModels {
			if strings.Contains(symbolModel.Symbol, "_") {
				continue
			}

			if _, skip := ignoreSymbols[symbolModel.Symbol]; skip {
				continue
			}

			if err = object.startProcess(calculatorPresetModel, calculatorFormulaPresetModel, symbolModel); err != nil {
				return err
			}
		}
	} else {
		symbolModel, err := object.symbolService().Load(request.Symbol, enums_symbol.SymbolStatusActive)
		if err != nil {
			return err
		}

		if err = object.startProcess(calculatorPresetModel, calculatorFormulaPresetModel, symbolModel); err != nil {
			return err
		}
	}

	object.websocketService().GetBroadcastChannel() <- &models_websocket.BroadcastChannelModel{
		Event: enums_websocket.WebsocketEventBotList,
		Data:  object.LoadAll(),
	}

	return nil
}

func (object *botServiceImplementation) startProcess(
	calculatorPresetModel *models_calculator_preset.CalculatorPresetModel,
	calculatorFormulaPresetModel *models_calculator_formula_preset.CalculatorFormulaPresetModel,
	symbolModel *models_symbol.SymbolModel,
) error {
	var botModel models_bot.BotModel

	if symbolModel.Limit.RightMin >= 20 {
		return nil
	}

	hash := services_helper.MustConvertStringToMd5(fmt.Sprintf(
		"hash | calculatorPresetModel:%d | calculatorFormulaPresetModel:%d | symbol:%s | ",
		calculatorPresetModel.ID,
		calculatorFormulaPresetModel.ID,
		symbolModel.Symbol,
	))

	err := object.storageService().DB().Where("hash = ?", hash).First(&botModel).Error
	if err == nil {
		return nil
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	botModel = models_bot.BotModel{
		CalculatorPresetID:        calculatorPresetModel.ID,
		CalculatorFormulaPresetID: calculatorFormulaPresetModel.ID,
		Hash:                      hash,
		Symbol:                    symbolModel.Symbol,
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
		Status:                    enums_bot.StatusNew,
		Filters:                   calculatorFormulaPresetModel.Filters,
		Formulas:                  calculatorFormulaPresetModel.Formulas,
		TickSize:                  symbolModel.Limit.TickSize,
		MinAmount:                 symbolModel.Limit.RightMin,
		Iterations:                services_helper.CalculateOptimalIterations(calculatorPresetModel.Window, services_helper.GetCpu(2), 0.05),
		ParamOld:                  models_bot.ParamModel{},
		Param:                     models_bot.ParamModel{},
		IsFirstRun:                true,
		IsEmptySend:               true,
	}

	if err = object.storageService().DB().Create(&botModel).Error; err != nil {
		return err
	}

	object.botRepositoryService().Add(&botModel)
	object.stopChannels[botModel.ID] = make(chan struct{})
	object.GetRunChannel() <- botModel.ID

	object.websocketService().GetBroadcastChannel() <- &models_websocket.BroadcastChannelModel{
		Event: enums_websocket.WebsocketEventBot,
		Data:  botModel,
	}

	return nil
}
