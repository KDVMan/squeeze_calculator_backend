package models_calculator

import (
	"backend/internal/enums"
)

type CalculatorRequestModel struct {
	ID              uint                 `json:"id" validate:"required,gt=0"`
	Symbol          string               `json:"symbol" validate:"required,alphanum,uppercase"`
	Window          int64                `json:"window" validate:"gte=0"`
	TradeDirection  enums.TradeDirection `json:"tradeDirection" validate:"required,tradeDirection"`
	Interval        enums.Interval       `json:"interval" validate:"required,interval"`
	TimeTo          int64                `json:"timeTo" validate:"required,gt=0"`
	Bind            []enums.Bind         `json:"bind" validate:"required,bind"`
	PercentInFrom   float64              `json:"percentInFrom" validate:"required,gt=0"`
	PercentInTo     float64              `json:"percentInTo" validate:"required,gtefield=PercentInFrom"`
	PercentInStep   float64              `json:"percentInStep" validate:"required,gt=0"`
	PercentOutFrom  float64              `json:"percentOutFrom" validate:"required,gt=0"`
	PercentOutTo    float64              `json:"percentOutTo" validate:"required,gtefield=PercentOutFrom"`
	PercentOutStep  float64              `json:"percentOutStep" validate:"required,gt=0"`
	StopTime        bool                 `json:"stopTime" validate:"boolean"`
	StopTimeFrom    int64                `json:"stopTimeFrom" validate:"gte=0"`
	StopTimeTo      int64                `json:"stopTimeTo" validate:"gtefield=StopTimeFrom"`
	StopTimeStep    int64                `json:"stopTimeStep" validate:"gte=0"`
	StopPercent     bool                 `json:"stopPercent" validate:"boolean"`
	StopPercentFrom float64              `json:"stopPercentFrom" validate:"gte=0"`
	StopPercentTo   float64              `json:"stopPercentTo" validate:"gtefield=StopPercentFrom"`
	StopPercentStep float64              `json:"stopPercentStep" validate:"gte=0"`
	Algorithm       enums.Algorithm      `json:"algorithm" validate:"required,algorithm"`
	Iterations      int64                `json:"iterations" validate:"required,gt=0"`
}
