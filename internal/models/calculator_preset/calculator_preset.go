package models_calculator_preset

import (
	"backend/internal/enums"
	"backend/internal/models"
	"encoding/json"
	"gorm.io/gorm"
)

type CalculatorPresetModel struct {
	models.DbModelWithID
	Name            string               `gorm:"uniqueIndex:unique_calculator_preset_01;not null" json:"name"`
	Window          int64                `json:"window"`
	TradeDirection  enums.TradeDirection `json:"tradeDirection"`
	Interval        enums.Interval       `json:"interval"`
	TimeFrom        int64                `json:"timeFrom"`
	TimeTo          int64                `json:"timeTo"`
	Bind            []enums.Bind         `gorm:"-" json:"bind"`
	BindJson        string               `gorm:"type:text" json:"-"`
	PercentInFrom   float64              `json:"percentInFrom"`
	PercentInTo     float64              `json:"percentInTo"`
	PercentInStep   float64              `json:"percentInStep"`
	PercentOutFrom  float64              `json:"percentOutFrom"`
	PercentOutTo    float64              `json:"percentOutTo"`
	PercentOutStep  float64              `json:"percentOutStep"`
	StopTime        bool                 `json:"stopTime"`
	StopTimeFrom    int64                `json:"stopTimeFrom"`
	StopTimeTo      int64                `json:"stopTimeTo"`
	StopTimeStep    int64                `json:"stopTimeStep"`
	StopPercent     bool                 `json:"stopPercent"`
	StopPercentFrom float64              `json:"stopPercentFrom"`
	StopPercentTo   float64              `json:"stopPercentTo"`
	StopPercentStep float64              `json:"stopPercentStep"`
	Algorithm       enums.Algorithm      `json:"algorithm"`
	Iterations      int                  `json:"iterations"`
	Selected        bool                 `json:"selected"`
}

func (CalculatorPresetModel) TableName() string {
	return "calculators_presets"
}

func (object *CalculatorPresetModel) BeforeSave(tx *gorm.DB) (err error) {
	bindData, err := json.Marshal(object.Bind)

	if err != nil {
		return err
	}

	object.BindJson = string(bindData)

	return nil
}

func (object *CalculatorPresetModel) AfterFind(tx *gorm.DB) (err error) {
	if object.BindJson != "" {
		err = json.Unmarshal([]byte(object.BindJson), &object.Bind)

		if err != nil {
			return err
		}
	}

	return nil
}

// func LoadDefault(tradeDirection enums.TradeDirection, interval enums.Interval) *CalculatorModel {
// 	currentTime := time.Now()
// 	start := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location()).AddDate(0, 0, -7)
// 	end := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 23, 59, 59, 0, currentTime.Location())
//
// 	model := &CalculatorModel{
// 		TradeDirection:  tradeDirection,
// 		Interval:        interval,
// 		TimeFrom:        start.UnixMilli(),
// 		TimeTo:          end.UnixMilli(),
// 		Bind:            enums.BindValues(),
// 		PercentInFrom:   1,
// 		PercentInTo:     5,
// 		PercentInStep:   0.1,
// 		PercentOutFrom:  0.1,
// 		PercentOutTo:    3,
// 		PercentOutStep:  0.1,
// 		StopTime:        true,
// 		StopTimeFrom:    5,
// 		StopTimeTo:      92,
// 		StopTimeStep:    1,
// 		StopPercent:     true,
// 		StopPercentFrom: 0.1,
// 		StopPercentTo:   6,
// 		StopPercentStep: 0.1,
// 		Algorithm:       enums.AlgorithmRandom,
// 		Iterations:      100000,
// 	}
//
// 	switch interval {
// 	case enums.Interval3m:
// 		model.PercentInTo = 7
// 		model.PercentOutTo = 4
// 		model.StopTimeFrom = 7
// 		model.Iterations = 50000
// 	case enums.Interval5m:
// 		model.PercentInTo = 9
// 		model.PercentOutTo = 6
// 		model.StopTimeFrom = 10
// 		model.Iterations = 70000
// 	case enums.Interval15m:
// 		model.PercentInTo = 12
// 		model.PercentOutTo = 8
// 		model.StopTimeFrom = 20
// 		model.Iterations = 100000
// 	case enums.Interval30m:
// 		model.PercentInTo = 15
// 		model.PercentOutTo = 10
// 		model.StopTimeFrom = 30
// 		model.StopTimeTo = 184
// 		model.Iterations = 150000
// 	case enums.Interval1h:
// 		model.PercentInTo = 20
// 		model.PercentOutTo = 12
// 		model.StopTimeFrom = 30
// 		model.StopTimeTo = 244
// 		model.Iterations = 200000
// 	}
//
// 	return model
// }
