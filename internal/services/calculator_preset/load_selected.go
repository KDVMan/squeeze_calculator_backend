package services_calculator_preset

import (
	models_calculator_preset "backend/internal/models/calculator_preset"
	"errors"
)

func (object *calculatorPresetServiceImplementation) LoadSelected() (*models_calculator_preset.CalculatorPresetModel, error) {
	presets, err := object.Load()
	if err != nil {
		return nil, err
	}

	for _, preset := range presets {
		if preset.Selected {
			return preset, nil
		}
	}

	return nil, errors.New("no selected preset found")
}
