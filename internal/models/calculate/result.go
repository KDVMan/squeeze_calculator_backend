package models_calculate

type CalculateResultModel struct {
	Total                        int                   `json:"total"`
	TotalStops                   int                   `json:"totalStops"`
	TotalStopsMinus              int                   `json:"totalStopsMinus"`
	TotalStopsPlus               int                   `json:"totalStopsPlus"`
	TotalTakes                   int                   `json:"totalTakes"`
	TotalTakesPlus               int                   `json:"totalTakesPlus"`
	TotalProfitPercent           float64               `json:"totalProfitPercent"`
	TotalCumulativeProfitPercent float64               `json:"totalCumulativeProfitPercent"`
	MaxDrawdownPercent           float64               `json:"maxDrawdownPercent"`
	MaxTimeDeal                  int64                 `json:"maxTimeDeal"`
	InOutRatio                   float64               `json:"inOutRatio"`
	Coefficient                  float64               `json:"coefficient"`
	WinRate                      float64               `json:"winRate"`
	WinRatePlus                  float64               `json:"winRatePlus"`
	Score                        float64               `json:"score"`
	ParamModel                   *ParamModel           `json:"param"`
	Deals                        []*CalculateDealModel `json:"deals"`
}

func (object *CalculateResultModel) GetFieldsValues() map[string]float64 {
	fieldMap := map[string]float64{
		"total":                        float64(object.Total),
		"totalStops":                   float64(object.TotalStops),
		"totalStopsMinus":              float64(object.TotalStopsMinus),
		"totalStopsPlus":               float64(object.TotalStopsPlus),
		"totalTakes":                   float64(object.TotalTakes),
		"totalTakesPlus":               float64(object.TotalTakesPlus),
		"totalProfitPercent":           object.TotalProfitPercent,
		"totalCumulativeProfitPercent": object.TotalCumulativeProfitPercent,
		"maxDrawdownPercent":           object.MaxDrawdownPercent,
		"maxTimeDeal":                  float64(object.MaxTimeDeal),
		"inOutRatio":                   object.InOutRatio,
		"coefficient":                  object.Coefficient,
		"winRate":                      object.WinRate,
		"winRatePlus":                  object.WinRatePlus,
		"score":                        object.Score,
	}

	if object.ParamModel != nil {
		fieldMap["percentIn"] = object.ParamModel.PercentIn
		fieldMap["percentOut"] = object.ParamModel.PercentOut
		fieldMap["stopTime"] = float64(object.ParamModel.StopTime)
		fieldMap["stopPercent"] = object.ParamModel.StopPercent
	}

	return fieldMap
}

func (object *CalculateResultModel) GetFieldValue(fieldName string) float64 {
	if object.ParamModel != nil {
		switch fieldName {
		case "percentIn":
			return object.ParamModel.PercentIn
		case "percentOut":
			return object.ParamModel.PercentOut
		case "stopTime":
			return float64(object.ParamModel.StopTime)
		case "stopPercent":
			return object.ParamModel.StopPercent
		}
	}

	fieldMap := map[string]float64{
		"total":                        float64(object.Total),
		"totalStops":                   float64(object.TotalStops),
		"totalStopsMinus":              float64(object.TotalStopsMinus),
		"totalStopsPlus":               float64(object.TotalStopsPlus),
		"totalTakes":                   float64(object.TotalTakes),
		"totalTakesPlus":               float64(object.TotalTakesPlus),
		"totalProfitPercent":           object.TotalProfitPercent,
		"totalCumulativeProfitPercent": object.TotalCumulativeProfitPercent,
		"maxDrawdownPercent":           object.MaxDrawdownPercent,
		"maxTimeDeal":                  float64(object.MaxTimeDeal),
		"inOutRatio":                   object.InOutRatio,
		"coefficient":                  object.Coefficient,
		"winRate":                      object.WinRate,
		"winRatePlus":                  object.WinRatePlus,
		"score":                        object.Score,
	}

	if value, ok := fieldMap[fieldName]; ok {
		return value
	}

	return 0.0
}
