package services_calculator_preset

import (
	models_calculator_preset "backend/internal/models/calculator_preset"
	"errors"
	"gorm.io/gorm"
)

func (object *calculatorPresetServiceImplementation) LoadByID(ID uint) *models_calculator_preset.CalculatorPresetModel {
	var calculatorPresetModel models_calculator_preset.CalculatorPresetModel

	if err := object.storageService().DB().
		Where("id = ?", ID).
		First(&calculatorPresetModel).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			object.loggerService().Error().Printf("failed to load calculator preset by id : %v", err)
		}

		return nil
	}

	return &calculatorPresetModel
}
