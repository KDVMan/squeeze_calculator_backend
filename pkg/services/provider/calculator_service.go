package services_provider

import (
	services_calculator "backend/internal/services/calculator"
	services_interface_calculator "backend/internal/services/calculator/interface"
)

func (object *ProviderService) CalculatorService() services_interface_calculator.CalculatorService {
	if object.calculatorService == nil {
		object.calculatorService = services_calculator.NewCalculatorService(
			object.ConfigService,
			object.StorageService,
			object.WebsocketService,
			object.InitService,
			object.SymbolService,
			object.QuoteService,
			object.CalculatorFormulaPresetService,
		)
	}

	return object.calculatorService
}
