package services_calculator_formula_preset

import (
	models_calculator_formula_preset "backend/internal/models/calculator_formula_preset"
	"errors"
	"gorm.io/gorm"
)

func (object *calculatorFormulaPresetServiceImplementation) Load() ([]*models_calculator_formula_preset.CalculatorFormulaPresetModel, error) {
	var presetsModels []*models_calculator_formula_preset.CalculatorFormulaPresetModel

	if err := object.storageService().DB().
		Order("name ASC").
		Find(&presetsModels).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return presetsModels, nil
}
