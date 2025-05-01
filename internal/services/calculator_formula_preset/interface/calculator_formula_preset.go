package services_interface_calculator_formula_preset

import (
	models_calculator_formula_preset "backend/internal/models/calculator_formula_preset"
)

type CalculatorFormulaPresetService interface {
	Load() ([]*models_calculator_formula_preset.CalculatorFormulaPresetModel, error)
	LoadSelected() (*models_calculator_formula_preset.CalculatorFormulaPresetModel, error)
	Add(*models_calculator_formula_preset.AddRequestModel) ([]*models_calculator_formula_preset.CalculatorFormulaPresetModel, error)
	Edit(*models_calculator_formula_preset.EditRequestModel) ([]*models_calculator_formula_preset.CalculatorFormulaPresetModel, error)
	Delete(*models_calculator_formula_preset.DeleteRequestModel) ([]*models_calculator_formula_preset.CalculatorFormulaPresetModel, error)
	Update(*models_calculator_formula_preset.UpdateRequestModel) ([]*models_calculator_formula_preset.CalculatorFormulaPresetModel, error)
	Duplicate(*models_calculator_formula_preset.DuplicateRequestModel) ([]*models_calculator_formula_preset.CalculatorFormulaPresetModel, error)
	LoadByID(uint) *models_calculator_formula_preset.CalculatorFormulaPresetModel
}
