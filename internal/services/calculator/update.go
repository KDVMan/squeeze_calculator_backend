package services_calculator

import (
	models_calculator "backend/internal/models/calculator"
	models_calculator_preset "backend/internal/models/calculator_preset"
	"errors"
	"gorm.io/gorm"
)

func (object *calculatorServiceImplementation) Update(request *models_calculator.CalculateRequestModel) (*models_calculator_preset.CalculatorPresetModel, error) {
	var presetModel models_calculator_preset.CalculatorPresetModel

	if err := object.storageService().DB().First(&presetModel, request.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("preset not found")
		}

		return nil, err
	}

	presetModel.Window = request.Window
	presetModel.TradeDirection = request.TradeDirection
	presetModel.Interval = request.Interval
	presetModel.TimeFrom = request.TimeFrom
	presetModel.TimeTo = request.TimeTo
	presetModel.Bind = request.Bind
	presetModel.PercentInFrom = request.PercentInFrom
	presetModel.PercentInTo = request.PercentInTo
	presetModel.PercentInStep = request.PercentInStep
	presetModel.PercentOutFrom = request.PercentOutFrom
	presetModel.PercentOutTo = request.PercentOutTo
	presetModel.PercentOutStep = request.PercentOutStep
	presetModel.StopTime = request.StopTime
	presetModel.StopTimeFrom = request.StopTimeFrom
	presetModel.StopTimeTo = request.StopTimeTo
	presetModel.StopTimeStep = request.StopTimeStep
	presetModel.StopPercent = request.StopPercent
	presetModel.StopPercentFrom = request.StopPercentFrom
	presetModel.StopPercentTo = request.StopPercentTo
	presetModel.StopPercentStep = request.StopPercentStep
	presetModel.Algorithm = request.Algorithm
	presetModel.Iterations = int(request.Iterations)

	if err := object.storageService().DB().Save(&presetModel).Error; err != nil {
		return nil, err
	}

	return &presetModel, nil
}
