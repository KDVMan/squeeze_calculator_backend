package models_bot

import (
	"backend/internal/enums"
	enums_bot "backend/internal/enums/bot"
	"backend/internal/models"
	models_calculator_formula_preset "backend/internal/models/calculator_formula_preset"
	"encoding/json"
	"gorm.io/gorm"
)

type BotModel struct {
	models.DbModelWithID
	CalculatorPresetID        uint                                            `gorm:"not null" json:"calculatorPresetId"`
	CalculatorFormulaPresetID uint                                            `gorm:"not null" json:"calculatorFormulaPresetId"`
	Hash                      string                                          `gorm:"not null;uniqueIndex:unique_bot_01" json:"hash"`
	Symbol                    string                                          `gorm:"not null" json:"symbol"`
	Window                    int64                                           `json:"window"`
	TradeDirection            enums.TradeDirection                            `json:"tradeDirection"`
	Interval                  enums.Interval                                  `json:"interval"`
	Bind                      []enums.Bind                                    `gorm:"-" json:"bind"`
	BindJson                  string                                          `gorm:"type:text" json:"-"`
	PercentInFrom             float64                                         `json:"percentInFrom"`
	PercentInTo               float64                                         `json:"percentInTo"`
	PercentInStep             float64                                         `json:"percentInStep"`
	PercentOutFrom            float64                                         `json:"percentOutFrom"`
	PercentOutTo              float64                                         `json:"percentOutTo"`
	PercentOutStep            float64                                         `json:"percentOutStep"`
	StopTime                  bool                                            `json:"stopTime"`
	StopTimeFrom              int64                                           `json:"stopTimeFrom"`
	StopTimeTo                int64                                           `json:"stopTimeTo"`
	StopTimeStep              int64                                           `json:"stopTimeStep"`
	StopPercent               bool                                            `json:"stopPercent"`
	StopPercentFrom           float64                                         `json:"stopPercentFrom"`
	StopPercentTo             float64                                         `json:"stopPercentTo"`
	StopPercentStep           float64                                         `json:"stopPercentStep"`
	Algorithm                 enums.Algorithm                                 `json:"algorithm"`
	Status                    enums_bot.Status                                `gorm:"not null" json:"status"`
	Filters                   []models_calculator_formula_preset.FilterModel  `gorm:"-" json:"filters"`
	FiltersJson               string                                          `json:"-"`
	Formulas                  []models_calculator_formula_preset.FormulaModel `gorm:"-" json:"formulas"`
	FormulasJson              string                                          `json:"-"`
	TickSize                  float64                                         `json:"tickSize"`
	MinAmount                 float64                                         `json:"minAmount"`
	Param                     ParamModel                                      `gorm:"embedded;embeddedPrefix:param_" json:"param"`
	ApiSend                   bool                                            `json:"apiSend"`
}

func (BotModel) TableName() string {
	return "bots"
}

func (object *BotModel) BeforeSave(tx *gorm.DB) (err error) {
	bindData, err := json.Marshal(object.Bind)

	if err != nil {
		return err
	}

	object.BindJson = string(bindData)

	if object.Filters != nil {
		data, err := json.Marshal(object.Filters)
		if err != nil {
			return err
		}

		object.FiltersJson = string(data)
	}

	if object.Formulas != nil {
		data, err := json.Marshal(object.Formulas)
		if err != nil {
			return err
		}

		object.FormulasJson = string(data)
	}

	return nil
}

func (object *BotModel) AfterFind(tx *gorm.DB) (err error) {
	if object.BindJson != "" {
		err = json.Unmarshal([]byte(object.BindJson), &object.Bind)

		if err != nil {
			return err
		}
	}

	if object.FiltersJson != "" {
		var list []models_calculator_formula_preset.FilterModel

		err = json.Unmarshal([]byte(object.FiltersJson), &list)
		if err != nil {
			return err
		}

		object.Filters = list
	}

	if object.FormulasJson != "" {
		var list []models_calculator_formula_preset.FormulaModel

		err = json.Unmarshal([]byte(object.FormulasJson), &list)
		if err != nil {
			return err
		}

		object.Formulas = list
	}

	return nil
}
