package models_quote

type LoadRequest struct {
	Symbol   string   `json:"symbol" validate:"required,alphanum,uppercase"`
	Interval Interval `json:"interval" validate:"required"`
	TimeEnd  int64    `json:"timeEnd" validate:"required,gte=0"`
}
