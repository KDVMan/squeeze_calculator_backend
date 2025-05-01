package services_calculator_formula_preset

import (
	models_calculator_formula_preset "backend/internal/models/calculator_formula_preset"
	"errors"
	"gorm.io/gorm"
)

func (object *calculatorFormulaPresetServiceImplementation) LoadByID(ID uint) *models_calculator_formula_preset.CalculatorFormulaPresetModel {
	var calculatorFormulaPresetModel models_calculator_formula_preset.CalculatorFormulaPresetModel

	if err := object.storageService().DB().
		Where("id = ?", ID).
		First(&calculatorFormulaPresetModel).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			object.loggerService().Error().Printf("failed to load calculator formula preset by id : %v", err)
		}

		return nil
	}

	return &calculatorFormulaPresetModel
}
