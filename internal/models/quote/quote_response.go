package models_quote

import models_calculate "backend/internal/models/calculate"

type QuoteResponseModel struct {
	Quotes   []*QuoteModel                          `json:"quotes"`
	TimeFrom int64                                  `json:"timeFrom"`
	TimeTo   int64                                  `json:"timeTo"`
	Deals    []*models_calculate.CalculateDealModel `json:"deals"`
}
