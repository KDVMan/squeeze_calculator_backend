package services_provider

import (
	services_calculator_formula_preset "backend/internal/services/calculator_formula_preset"
	services_interface_calculator_formula_preset "backend/internal/services/calculator_formula_preset/interface"
)

func (object *ProviderService) CalculatorFormulaPresetService() services_interface_calculator_formula_preset.CalculatorFormulaPresetService {
	if object.calculatorFormulaPresetService == nil {
		object.calculatorFormulaPresetService = services_calculator_formula_preset.NewCalculatorFormulaPresetService(
			object.ConfigService,
			object.StorageService,
			object.WebsocketService,
			object.CalculatorService,
			object.CalculatorPresetService,
		)
	}

	return object.calculatorFormulaPresetService
}
