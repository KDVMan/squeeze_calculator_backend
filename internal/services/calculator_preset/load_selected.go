package services_calculator_preset

import (
	models_calculator_preset "backend/internal/models/calculator_preset"
	"errors"
	"gorm.io/gorm"
)

func (object *calculatorPresetServiceImplementation) LoadSelected() (*models_calculator_preset.CalculatorPresetModel, error) {
	var presetModel models_calculator_preset.CalculatorPresetModel

	if err := object.storageService().DB().
		Where("selected = ?", true).
		First(&presetModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("no selected preset found")
		}

		return nil, err
	}

	return &presetModel, nil
}
