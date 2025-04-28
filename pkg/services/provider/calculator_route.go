package services_provider

import (
	routes_calculator "backend/internal/routes/calculator"
	routes_interface_calculator "backend/internal/routes/calculator/interface"
)

func (object *ProviderService) CalculatorRoute() routes_interface_calculator.CalculatorRoute {
	if object.calculatorRoute == nil {
		object.calculatorRoute = routes_calculator.NewCalculatorRoute(
			object.LoggerService,
			object.RequestService,
			object.CalculatorService,
			object.CalculatorPresetService,
		)
	}

	return object.calculatorRoute
}
