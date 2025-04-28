package models_quote

type QuoteRangeModel struct {
	QuotesLimit int64
	TimeFrom    int64
	TimeTo      int64
	TimeStep    int64
	Iterations  int
}

func GetRange(limit int64, timeFrom int64, timeTo int64, milliseconds int64) *QuoteRangeModel {
	timeRange := (timeTo - timeFrom) + 1000
	total := timeRange / milliseconds

	if limit > total {
		total = limit
	}

	return &QuoteRangeModel{
		QuotesLimit: limit,
		TimeFrom:    timeFrom,
		TimeTo:      timeTo,
		TimeStep:    milliseconds * limit,
		Iterations:  int((total + limit - 1) / limit),
	}
}
