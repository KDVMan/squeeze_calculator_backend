package models_calculator_optimization

import "backend/internal/enums"

type CalculatorOptimizationRequestModel struct {
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
}
