package models_init

import (
	"backend/internal/enums"
	enums_calculate "backend/internal/enums/calculate"
	models_quote "backend/internal/models/quote"
)

type UpdateRequestModel struct {
	Symbol                 string                     `json:"symbol" validate:"required,symbolFormat,uppercase"`
	Intervals              []*models_quote.Interval   `json:"intervals,omitempty"`
	CalculateSortColumn    enums_calculate.SortColumn `json:"calculateSortColumn" validate:"required,calculateSortColumn"`
	CalculateSortDirection enums.SortDirection        `json:"calculateSortDirection" validate:"required,sortDirection"`
	ExecActive             enums.ExecActive           `json:"execActive" validate:"required,execActive"`
}
