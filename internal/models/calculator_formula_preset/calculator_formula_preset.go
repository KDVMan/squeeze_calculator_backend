package models_calculator_formula_preset

import (
	"backend/internal/models"
	"encoding/json"
	"gorm.io/gorm"
)

type CalculatorFormulaPresetModel struct {
	models.DbModelWithID
	Name         string         `gorm:"uniqueIndex:unique_calculator_formula_preset_01;not null" json:"name"`
	Filters      []FilterModel  `gorm:"-" json:"filters"`
	FiltersJson  string         `json:"-"`
	Formulas     []FormulaModel `gorm:"-" json:"formulas"`
	FormulasJson string         `json:"-"`
	Selected     bool           `json:"selected"`
}

func (CalculatorFormulaPresetModel) TableName() string {
	return "calculators_formulas_presets"
}

func (object *CalculatorFormulaPresetModel) BeforeSave(tx *gorm.DB) (err error) {
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

func (object *CalculatorFormulaPresetModel) AfterFind(tx *gorm.DB) (err error) {
	if object.FiltersJson != "" {
		var list []FilterModel

		err = json.Unmarshal([]byte(object.FiltersJson), &list)
		if err != nil {
			return err
		}

		object.Filters = list
	}

	if object.FormulasJson != "" {
		var list []FormulaModel

		err = json.Unmarshal([]byte(object.FormulasJson), &list)
		if err != nil {
			return err
		}

		object.Formulas = list
	}

	return nil
}
