package services_calculator

import (
	models_calculate "backend/internal/models/calculate"
	"math"
)

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
