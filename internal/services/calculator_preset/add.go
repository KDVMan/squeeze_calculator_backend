package services_calculator_preset

import (
	models_calculator_preset "backend/internal/models/calculator_preset"
	"errors"
	"gorm.io/gorm"
)

func (object *calculatorPresetServiceImplementation) Add(
	request *models_calculator_preset.AddRequestModel,
) ([]*models_calculator_preset.CalculatorPresetModel, error) {
	var presetModel models_calculator_preset.CalculatorPresetModel

	err := object.storageService().DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&models_calculator_preset.CalculatorPresetModel{}).
			Where("selected = ?", true).
			Update("selected", false).Error; err != nil {
			return err
		}

		if err := tx.Where("name = ?", request.Name).
			First(&models_calculator_preset.CalculatorPresetModel{}).Error; err == nil {
			return errors.New("preset with the same name already exists")
		}

		presetModel = models_calculator_preset.CalculatorPresetModel{
			Name:     request.Name,
			Selected: true,
		}

		if err := tx.Create(&presetModel).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return object.Load()
}
