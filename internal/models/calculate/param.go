package models_calculate

import (
	"backend/internal/enums"
)

type ParamModel struct {
	TradeDirection enums.TradeDirection `json:"tradeDirection"`
	Interval       enums.Interval       `json:"interval"`
	Bind           enums.Bind           `json:"bind"`
	PercentIn      float64              `json:"percentIn"`
	PercentOut     float64              `json:"percentOut"`
	StopTime       int64                `json:"stopTime"`
	StopPercent    float64              `json:"stopPercent"`
	IsCurrent      bool                 `json:"isCurrent"`
}
