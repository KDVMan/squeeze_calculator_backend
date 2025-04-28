package services_provider

import (
	routes_calculator_preset "backend/internal/routes/calculator_preset"
	routes_interface_calculator_preset "backend/internal/routes/calculator_preset/interface"
)

func (object *ProviderService) CalculatorPresetRoute() routes_interface_calculator_preset.CalculatorPresetRoute {
	if object.calculatorPresetRoute == nil {
		object.calculatorPresetRoute = routes_calculator_preset.NewCalculatorPresetRoute(
			object.LoggerService,
			object.RequestService,
			object.CalculatorPresetService,
		)
	}

	return object.calculatorPresetRoute
}
