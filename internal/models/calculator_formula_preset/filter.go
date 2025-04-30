package models_calculator_formula_preset

import (
	enums_calculator "backend/internal/enums/calculator"
	models_calculate "backend/internal/models/calculate"
)

type FilterModel struct {
	Name   string                  `json:"name"`
	Filter enums_calculator.Filter `json:"filter"`
	Value  float64                 `json:"value"`
}

func ApplyFilters(result *models_calculate.CalculateResultModel, filters []FilterModel) bool {
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
