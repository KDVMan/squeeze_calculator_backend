package services_interface_calculator_preset

import (
	models_calculator_preset "backend/internal/models/calculator_preset"
)

type CalculatorPresetService interface {
	Load() ([]*models_calculator_preset.CalculatorPresetModel, error)
	LoadSelected() (*models_calculator_preset.CalculatorPresetModel, error)
	Add(*models_calculator_preset.AddRequestModel) ([]*models_calculator_preset.CalculatorPresetModel, error)
	Edit(*models_calculator_preset.EditRequestModel) ([]*models_calculator_preset.CalculatorPresetModel, error)
	Duplicate(*models_calculator_preset.DuplicateRequestModel) ([]*models_calculator_preset.CalculatorPresetModel, error)
	Delete(*models_calculator_preset.DeleteRequestModel) ([]*models_calculator_preset.CalculatorPresetModel, error)
}
