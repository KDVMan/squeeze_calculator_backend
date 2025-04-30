package services_calculator_optimization

import (
	"backend/internal/enums"
	models_calculator_optimization "backend/internal/models/calculator_optimization"
)

func Load(request *models_calculator_optimization.CalculatorOptimizationRequestModel) []*models_calculator_optimization.CalculatorOptimizationModel {
	result := make(map[string]*models_calculator_optimization.CalculatorOptimizationModel)

	if request.Algorithm == enums.AlgorithmRandom {
		result = loadRandom(request)
	} else if request.Algorithm == enums.AlgorithmGrid {
		result = loadGrid(request)
	}

	optimizations := make([]*models_calculator_optimization.CalculatorOptimizationModel, 0, len(result))

	for _, optimization := range result {
		optimizations = append(optimizations, optimization)
	}

	return optimizations
}
