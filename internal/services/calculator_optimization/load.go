package services_calculator_optimization

import (
	"backend/internal/enums"
	models_calculator_optimization "backend/internal/models/calculator_optimization"
	models_calculator_preset "backend/internal/models/calculator_preset"
)

func Load(calculatorModel *models_calculator_preset.CalculatorPresetModel) []*models_calculator_optimization.CalculatorOptimizationModel {
	result := make(map[string]*models_calculator_optimization.CalculatorOptimizationModel)

	if calculatorModel.Algorithm == enums.AlgorithmRandom {
		result = loadRandom(calculatorModel)
	} else if calculatorModel.Algorithm == enums.AlgorithmGrid {
		result = loadGrid(calculatorModel)
	}

	optimizations := make([]*models_calculator_optimization.CalculatorOptimizationModel, 0, len(result))

	for _, optimization := range result {
		optimizations = append(optimizations, optimization)
	}

	return optimizations
}
