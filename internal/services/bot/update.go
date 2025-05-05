package services_bot

import (
	enums_bot "backend/internal/enums/bot"
	enums_websocket "backend/internal/enums/websocket"
	models_bot "backend/internal/models/bot"
	models_websocket "backend/internal/models/websocket"
)

func (object *botServiceImplementation) Update(request *models_bot.UpdateRequestModel) error {
	var botsModels []models_bot.BotModel
	calculatorPresetModel := object.calculatorPresetService().LoadByID(request.CalculatorPresetID)
	calculatorFormulaPresetModel := object.calculatorFormulaPresetService().LoadByID(request.CalculatorFormulaPresetID)

	if err := object.storageService().DB().
		Where("calculator_preset_id = ? OR calculator_formula_preset_id = ?", request.CalculatorPresetID, request.CalculatorFormulaPresetID).
		Find(&botsModels).Error; err != nil {
		return err
	}

	for _, botModel := range botsModels {
		shouldUpdate := false

		if botModel.CalculatorPresetID == request.CalculatorPresetID {
			botModel.Window = calculatorPresetModel.Window
			botModel.TradeDirection = calculatorPresetModel.TradeDirection
			botModel.Interval = calculatorPresetModel.Interval
			botModel.Bind = calculatorPresetModel.Bind
			botModel.PercentInFrom = calculatorPresetModel.PercentInFrom
			botModel.PercentInTo = calculatorPresetModel.PercentInTo
			botModel.PercentInStep = calculatorPresetModel.PercentInStep
			botModel.PercentOutFrom = calculatorPresetModel.PercentOutFrom
			botModel.PercentOutTo = calculatorPresetModel.PercentOutTo
			botModel.PercentOutStep = calculatorPresetModel.PercentOutStep
			botModel.StopTime = calculatorPresetModel.StopTime
			botModel.StopTimeFrom = calculatorPresetModel.StopTimeFrom
			botModel.StopTimeTo = calculatorPresetModel.StopTimeTo
			botModel.StopTimeStep = calculatorPresetModel.StopTimeStep
			botModel.StopPercent = calculatorPresetModel.StopPercent
			botModel.StopPercentFrom = calculatorPresetModel.StopPercentFrom
			botModel.StopPercentTo = calculatorPresetModel.StopPercentTo
			botModel.StopPercentStep = calculatorPresetModel.StopPercentStep
			botModel.Algorithm = calculatorPresetModel.Algorithm
			botModel.ParamOld = models_bot.ParamModel{}
			botModel.Param = models_bot.ParamModel{}
			botModel.IsFirstRun = false
			botModel.IsEmptySend = false

			shouldUpdate = true
		}

		if botModel.CalculatorFormulaPresetID == request.CalculatorFormulaPresetID {
			botModel.Filters = calculatorFormulaPresetModel.Filters
			botModel.Formulas = calculatorFormulaPresetModel.Formulas
			botModel.ParamOld = models_bot.ParamModel{}
			botModel.Param = models_bot.ParamModel{}
			botModel.IsFirstRun = false
			botModel.IsEmptySend = false

			shouldUpdate = true
		}

		if shouldUpdate {
			if err := object.storageService().DB().Save(&botModel).Error; err != nil {
				object.loggerService().Error().Printf("failed to update bot %d: %v", botModel.ID, err)
				continue
			}

			object.botRepositoryService().Add(&botModel)

			if botModel.Status == enums_bot.StatusStart {
				if ch, ok := object.stopChannels[botModel.ID]; ok {
					close(ch)
				}

				object.stopChannels[botModel.ID] = make(chan struct{})
				object.GetRunChannel() <- botModel.ID
			}
		}
	}

	object.websocketService().GetBroadcastChannel() <- &models_websocket.BroadcastChannelModel{
		Event: enums_websocket.WebsocketEventBotList,
		Data:  object.LoadAll(),
	}

	return nil
}
