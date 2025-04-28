package models_init

import (
	"backend/internal/enums"
	enums_calculate "backend/internal/enums/calculate"
	models_quote "backend/internal/models/quote"
	"encoding/json"
	"gorm.io/gorm"
)

type InitModel struct {
	ID                     int                        `json:"-" gorm:"primaryKey"`
	Symbol                 string                     `json:"symbol"`
	IntervalsJson          string                     `json:"-"` // поле для хранения JSON в базе данных
	Intervals              []*models_quote.Interval   `json:"intervals" gorm:"-"`
	QuotesLimit            uint                       `json:"quotesLimit"`
	Precision              int                        `json:"precision"`
	CalculateLimit         int                        `json:"calculateLimit"`
	CalculateSortColumn    enums_calculate.SortColumn `json:"calculateSortColumn"`
	CalculateSortDirection enums.SortDirection        `json:"calculateSortDirection"`
	ExecActive             enums.ExecActive           `json:"execActive"`
}

func (InitModel) TableName() string {
	return "init"
}

func LoadDefault() *InitModel {
	return &InitModel{
		Symbol:                 "BTCUSDT",
		Intervals:              models_quote.IntervalLoadDefault(),
		QuotesLimit:            1500,
		Precision:              2,
		CalculateLimit:         1000,
		CalculateSortColumn:    enums_calculate.SortColumnScore,
		CalculateSortDirection: enums.SortDirectionDesc,
		ExecActive:             enums.ExecActiveBot,
	}
}

func (object *InitModel) BeforeSave(tx *gorm.DB) (err error) {
	if object.Intervals != nil {
		data, err := json.Marshal(object.Intervals)
		if err != nil {
			return err
		}

		object.IntervalsJson = string(data)
	}

	return nil
}

func (object *InitModel) AfterFind(tx *gorm.DB) (err error) {
	if object.IntervalsJson != "" {
		var list []*models_quote.Interval

		err = json.Unmarshal([]byte(object.IntervalsJson), &list)
		if err != nil {
			return err
		}

		object.Intervals = list
	}

	return nil
}
