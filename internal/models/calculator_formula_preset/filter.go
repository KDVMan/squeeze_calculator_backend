package models_calculator_formula_preset

import (
	enums_calculator "backend/internal/enums/calculator"
)

type FilterModel struct {
	Name   string                  `json:"name"`
	Filter enums_calculator.Filter `json:"filter"`
	Value  float64                 `json:"value"`
}
