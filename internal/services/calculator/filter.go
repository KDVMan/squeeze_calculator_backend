package services_calculator

import (
	enums_calculator "backend/internal/enums/calculator"
	models_calculate "backend/internal/models/calculate"
	models_calculator_formula_preset "backend/internal/models/calculator_formula_preset"
	"math"
)

func ApplyFilters(result *models_calculate.CalculateResultModel, filters []models_calculator_formula_preset.FilterModel) bool {
	for _, filter := range filters {
		if filter.Name == "" || filter.Filter == "" {
			continue
		}

		value := result.GetFieldValue(filter.Name)

		switch filter.Filter {
		case enums_calculator.FilterLt:
			if value >= filter.Value {
				return false
			}
		case enums_calculator.FilterLte:
			if value > filter.Value {
				return false
			}
		case enums_calculator.FilterGt:
			if value <= filter.Value {
				return false
			}
		case enums_calculator.FilterGte:
			if value < filter.Value {
				return false
			}
		case enums_calculator.FilterEqual:
			if value != filter.Value {
				return false
			}
		}
	}

	return true
}

func UpdateValueRanges(result *models_calculate.CalculateResultModel, ranges map[string][2]float64) {
	for field, value := range result.GetFieldsValues() {
		if existingRange, exists := ranges[field]; !exists {
			ranges[field] = [2]float64{value, value}
		} else {
			ranges[field] = [2]float64{
				math.Min(existingRange[0], value),
				math.Max(existingRange[1], value),
			}
		}
	}
}

func ApplyFormula(result *models_calculate.CalculateResultModel, formulas []models_calculator_formula_preset.FormulaModel, ranges map[string][2]float64) float64 {
	score := 0.0

	for _, formula := range formulas {
		value := result.GetFieldValue(formula.Name)

		if rangeValues, exists := ranges[formula.Name]; exists {
			minValue, maxValue := rangeValues[0], rangeValues[1]

			if maxValue > minValue {
				value = (value - minValue) / (maxValue - minValue)
			}
		}

		score += value * formula.Multiplier
	}

	return score
}
