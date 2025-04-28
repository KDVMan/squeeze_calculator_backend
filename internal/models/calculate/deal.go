package models_calculate

type CalculateDealModel struct {
	TimeIn          int64   `json:"timeIn"`
	TimeOut         int64   `json:"timeOut"`
	PriceIn         float64 `json:"priceIn"`
	PriceOut        float64 `json:"priceOut"`
	MinPrice        float64 `json:"minPrice"`
	IsStopTime      bool    `json:"isStopTime"`
	IsStopPercent   bool    `json:"isStopPercent"`
	ProfitPercent   float64 `json:"profitPercent"`
	DrawdownPercent float64 `json:"drawdownPercent"`
}
