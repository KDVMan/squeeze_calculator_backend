package services_calculator_formula_preset

import (
	services_interface_calculator "backend/internal/services/calculator/interface"
	services_interface_calculator_formula_preset "backend/internal/services/calculator_formula_preset/interface"
	services_interface_calculator_preset "backend/internal/services/calculator_preset/interface"
	services_interface_websocket "backend/internal/services/websocket/interface"
	services_interface_config "backend/pkg/services/config/interface"
	services_interface_storage "backend/pkg/services/storage/interface"
)

type calculatorFormulaPresetServiceImplementation struct {
	configService           func() services_interface_config.ConfigService
	storageService          func() services_interface_storage.StorageService
	websocketService        func() services_interface_websocket.WebsocketService
	calculatorService       func() services_interface_calculator.CalculatorService
	calculatorPresetService func() services_interface_calculator_preset.CalculatorPresetService
}

func NewCalculatorFormulaPresetService(
	configService func() services_interface_config.ConfigService,
	storageService func() services_interface_storage.StorageService,
	websocketService func() services_interface_websocket.WebsocketService,
	calculatorService func() services_interface_calculator.CalculatorService,
	calculatorPresetService func() services_interface_calculator_preset.CalculatorPresetService,
) services_interface_calculator_formula_preset.CalculatorFormulaPresetService {
	return &calculatorFormulaPresetServiceImplementation{
		configService:           configService,
		storageService:          storageService,
		websocketService:        websocketService,
		calculatorService:       calculatorService,
		calculatorPresetService: calculatorPresetService,
	}
}
