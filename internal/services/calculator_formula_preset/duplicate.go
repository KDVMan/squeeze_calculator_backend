package services_calculator_formula_preset

import (
	models_calculator_formula_preset "backend/internal/models/calculator_formula_preset"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func (object *calculatorFormulaPresetServiceImplementation) Duplicate(
	request *models_calculator_formula_preset.DuplicateRequestModel,
) ([]*models_calculator_formula_preset.CalculatorFormulaPresetModel, error) {
	var formulaPresetModel models_calculator_formula_preset.CalculatorFormulaPresetModel

	err := object.storageService().DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&models_calculator_formula_preset.CalculatorFormulaPresetModel{}).
			Where("selected = ?", true).
			Update("selected", false).Error; err != nil {
			return err
		}

		if err := tx.First(&formulaPresetModel, request.ID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("preset not found")
			}

			return err
		}

		name, err := object.generateUniqueName(tx, formulaPresetModel.Name+"_clone")
		if err != nil {
			return err
		}

		newFormulaPresetModel := formulaPresetModel
		newFormulaPresetModel.ID = 0
		newFormulaPresetModel.Name = name
		newFormulaPresetModel.Selected = true

		if err = tx.Create(&newFormulaPresetModel).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return object.Load()
}

func (object *calculatorFormulaPresetServiceImplementation) generateUniqueName(tx *gorm.DB, baseName string) (string, error) {
	name := baseName
	counter := 1

	for {
		var count int64
		if err := tx.Model(&models_calculator_formula_preset.CalculatorFormulaPresetModel{}).
			Where("name = ?", name).
			Count(&count).Error; err != nil {
			return "", err
		}

		if count == 0 {
			break
		}

		name = fmt.Sprintf("%s_%d", baseName, counter)
		counter++
	}

	return name, nil
}
