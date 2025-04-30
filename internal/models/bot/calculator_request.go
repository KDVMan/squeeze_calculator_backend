package models_bot

import (
	"backend/internal/enums"
	models_calculator_formula_preset "backend/internal/models/calculator_formula_preset"
)

type CalculatorRequestModel struct {
	BotID           uint
	Symbol          string
	TradeDirection  enums.TradeDirection
	Interval        enums.Interval
	Bind            []enums.Bind
	PercentInFrom   float64
	PercentInTo     float64
	PercentInStep   float64
	PercentOutFrom  float64
	PercentOutTo    float64
	PercentOutStep  float64
	StopTime        bool
	StopTimeFrom    int64
	StopTimeTo      int64
	StopTimeStep    int64
	StopPercent     bool
	StopPercentFrom float64
	StopPercentTo   float64
	StopPercentStep float64
	Algorithm       enums.Algorithm
	Iterations      int
	TickSize        float64
	Filters         []models_calculator_formula_preset.FilterModel
	Formulas        []models_calculator_formula_preset.FormulaModel `gorm:"-" json:"formulas"`
	Param           ParamModel
}
