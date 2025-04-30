package services_calculator_preset

import (
	models_calculator_preset "backend/internal/models/calculator_preset"
	"errors"
	"gorm.io/gorm"
)

func (object *calculatorPresetServiceImplementation) LoadAll() ([]*models_calculator_preset.CalculatorPresetModel, error) {
	var presetsModels []*models_calculator_preset.CalculatorPresetModel

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
