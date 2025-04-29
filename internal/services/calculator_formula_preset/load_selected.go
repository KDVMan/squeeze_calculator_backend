package services_calculator_formula_preset

import (
	models_calculator_formula_preset "backend/internal/models/calculator_formula_preset"
	"errors"
	"gorm.io/gorm"
)

func (object *calculatorFormulaPresetServiceImplementation) LoadSelected() (*models_calculator_formula_preset.CalculatorFormulaPresetModel, error) {
	var formulaPresetModel models_calculator_formula_preset.CalculatorFormulaPresetModel

	if err := object.storageService().DB().
		Where("selected = ?", true).
		First(&formulaPresetModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("no selected formula preset found")
		}

		return nil, err
	}

	return &formulaPresetModel, nil
}
