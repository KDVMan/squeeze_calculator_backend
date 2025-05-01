package models_bot

import "backend/internal/enums"

type ApiRequestModel struct {
	Deposit        float64              `json:"deposit"`
	IsReal         bool                 `json:"isReal"`
	Symbol         string               `json:"symbol"`
	Window         int64                `json:"window"`
	Interval       enums.Interval       `json:"interval"`
	TradeDirection enums.TradeDirection `json:"tradeDirection"`
	Bind           enums.Bind           `json:"bind"`
	PercentIn      float64              `json:"percentIn"`
	PercentOut     float64              `json:"percentOut"`
	StopTime       int64                `json:"stopTime"`
	StopPercent    float64              `json:"stopPercent"`
	TriggerStart   float64              `json:"triggerStart"`
	LimitQuotes    int64                `json:"limitQuotes"`
	IsCalculator   bool                 `json:"isCalculator"`
}
