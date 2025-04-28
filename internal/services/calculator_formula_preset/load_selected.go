package services_calculator_formula_preset

import (
	models_calculator_formula_preset "backend/internal/models/calculator_formula_preset"
	"errors"
)

func (object *calculatorFormulaPresetServiceImplementation) LoadSelected() (*models_calculator_formula_preset.CalculatorFormulaPresetModel, error) {
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
