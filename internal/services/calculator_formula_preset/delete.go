package services_calculator_formula_preset

import (
	models_calculator_formula_preset "backend/internal/models/calculator_formula_preset"
)

func (object *calculatorFormulaPresetServiceImplementation) Delete(
	request *models_calculator_formula_preset.DeleteRequestModel,
) ([]*models_calculator_formula_preset.CalculatorFormulaPresetModel, error) {
	if err := object.storageService().DB().Delete(&models_calculator_formula_preset.CalculatorFormulaPresetModel{}, request.ID).Error; err != nil {
		return nil, err
	}

	return object.Load()
}
