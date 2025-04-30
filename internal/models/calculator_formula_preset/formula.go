package models_calculator_formula_preset

import models_calculate "backend/internal/models/calculate"

type FormulaModel struct {
	Name       string  `json:"name"`
	Multiplier float64 `json:"multiplier"`
}

func ApplyFormula(result *models_calculate.CalculateResultModel, formulas []FormulaModel, ranges map[string][2]float64) float64 {
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
