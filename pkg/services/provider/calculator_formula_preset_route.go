package services_provider

import (
	routes_calculator_formula_preset "backend/internal/routes/calculator_formula_preset"
	routes_interface_calculator_formula_preset "backend/internal/routes/calculator_formula_preset/interface"
)

func (object *ProviderService) CalculatorFormulaPresetRoute() routes_interface_calculator_formula_preset.CalculatorFormulaPresetRoute {
	if object.calculatorFormulaPresetRoute == nil {
		object.calculatorFormulaPresetRoute = routes_calculator_formula_preset.NewCalculatorFormulaPresetRoute(
			object.LoggerService,
			object.RequestService,
			object.CalculatorFormulaPresetService,
		)
	}

	return object.calculatorFormulaPresetRoute
}
