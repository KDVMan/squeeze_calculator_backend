package services_calculator

import (
	"backend/internal/enums"
	models_calculate "backend/internal/models/calculate"
	models_calculator_formula_preset "backend/internal/models/calculator_formula_preset"
	"sort"
)

func (object *calculatorServiceImplementation) LoadResult(symbol string) []*models_calculate.CalculateResultModel {
	var results []*models_calculate.CalculateResultModel
	ranges := make(map[string][2]float64)

	initModel, err := object.initService().Load()
	if err != nil {
		return results
	}

	if symbol == "" {
		symbol = initModel.Symbol
	}

	if symbol != object.calculateSymbol {
		return results
	}

	preset, err := object.calculatorFormulaPresetService().LoadSelected()
	if err != nil {
		return results
	}

	// log.Printf("before, limit: %d, total: %d", initModel.CalculateLimit, len(object.calculateResult))

	for _, calculateResult := range object.calculateResult {
		if models_calculator_formula_preset.ApplyFilters(calculateResult, preset.Filters) {
			results = append(results, calculateResult)
			UpdateValueRanges(calculateResult, ranges)
		}
	}

	// log.Printf("after filter, total: %d", len(results))

	for _, result := range results {
		result.Score = models_calculator_formula_preset.ApplyFormula(result, preset.Formulas, ranges)
	}

	// log.Printf("after formula, total: %d", len(results))

	sort.Slice(results, func(i, j int) bool {
		a := results[i].GetFieldValue(initModel.CalculateSortColumn.String())
		b := results[j].GetFieldValue(initModel.CalculateSortColumn.String())

		if initModel.CalculateSortDirection == enums.SortDirectionDesc {
			return a > b
		}

		return a < b
	})

	// log.Printf("after sort, total: %d", len(results))

	if len(results) > initModel.CalculateLimit {
		results = results[:initModel.CalculateLimit]
	}

	// log.Printf("after slice, total: %d", len(results))

	return results
}
