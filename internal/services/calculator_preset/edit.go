package services_calculator_preset

import (
	models_calculator_preset "backend/internal/models/calculator_preset"
	"errors"
	"gorm.io/gorm"
)

func (object *calculatorPresetServiceImplementation) Edit(
	request *models_calculator_preset.EditRequestModel,
) ([]*models_calculator_preset.CalculatorPresetModel, error) {
	var presetModel models_calculator_preset.CalculatorPresetModel

	err := object.storageService().DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&models_calculator_preset.CalculatorPresetModel{}).
			Where("selected = ?", true).
			Update("selected", false).Error; err != nil {
			return err
		}

		if err := tx.First(&presetModel, request.ID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("preset not found")
			}

			return err
		}

		presetModel.Name = request.Name
		presetModel.Selected = true

		if err := tx.Save(&presetModel).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return object.Load()
}
