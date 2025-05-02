package models_bot

import (
	"backend/internal/enums"
)

type ParamModel struct {
	Bind        enums.Bind `json:"bind"`
	PercentIn   float64    `json:"percentIn"`
	PercentOut  float64    `json:"percentOut"`
	StopTime    int64      `json:"stopTime"`
	StopPercent float64    `json:"stopPercent"`
	Score       float64    `json:"score"`  // что бы видеть какой был
	Profit      float64    `json:"profit"` // что бы видеть какой был
	LastUpdate  int64      `json:"lastUpdate"`
}
