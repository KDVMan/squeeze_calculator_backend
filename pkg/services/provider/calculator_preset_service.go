package services_provider

import (
	services_calculator_preset "backend/internal/services/calculator_preset"
	services_interface_calculator_preset "backend/internal/services/calculator_preset/interface"
)

func (object *ProviderService) CalculatorPresetService() services_interface_calculator_preset.CalculatorPresetService {
	if object.calculatorPresetService == nil {
		object.calculatorPresetService = services_calculator_preset.NewCalculatorPresetService(
			object.ConfigService,
			object.StorageService,
			object.WebsocketService,
			object.InitService,
			object.SymbolService,
			object.QuoteService,
			object.CalculatorFormulaPresetService,
		)
	}

	return object.calculatorPresetService
}
