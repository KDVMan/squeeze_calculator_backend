package services_calculator_formula_preset

import (
	enums_websocket "backend/internal/enums/websocket"
	models_calculator_formula_preset "backend/internal/models/calculator_formula_preset"
	models_websocket "backend/internal/models/websocket"
	"errors"
	"gorm.io/gorm"
)

func (object *calculatorFormulaPresetServiceImplementation) Update(
	request *models_calculator_formula_preset.UpdateRequestModel,
) ([]*models_calculator_formula_preset.CalculatorFormulaPresetModel, error) {
	var presetModel models_calculator_formula_preset.CalculatorFormulaPresetModel

	if err := object.storageService().DB().First(&presetModel, request.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("preset not found")
		}

		return nil, err
	}

	presetModel.Name = request.Name
	presetModel.Filters = request.Filters
	presetModel.Formulas = request.Formulas

	if err := object.storageService().DB().Save(&presetModel).Error; err != nil {
		return nil, err
	}

	object.websocketService().GetBroadcastChannel() <- &models_websocket.BroadcastChannelModel{
		Event: enums_websocket.WebsocketEventCalculateResult,
		Data:  object.calculatorService().LoadResult(""),
	}

	return object.Load()
}
