package services_calculator_preset

import (
	models_calculator_preset "backend/internal/models/calculator_preset"
)

func (object *calculatorPresetServiceImplementation) Delete(
	request *models_calculator_preset.DeleteRequestModel,
) ([]*models_calculator_preset.CalculatorPresetModel, error) {
	if err := object.storageService().DB().Delete(&models_calculator_preset.CalculatorPresetModel{}, request.ID).Error; err != nil {
		return nil, err
	}

	return object.LoadAll()
}
