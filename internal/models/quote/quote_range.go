package models_quote

import (
	models_bot "backend/internal/models/bot"
	"time"
)

type QuoteRangeModel struct {
	QuotesLimit int64
	TimeFrom    int64
	TimeTo      int64
	TimeStep    int64
	Iterations  int
}

func GetTimeRange(botModel *models_bot.BotModel) (int64, int64) {
	currentTime := time.Now().UnixMilli()
	// timeFrom := currentTime - (botModel.Window * 60 * 1000)
	timeFrom := currentTime - (1440 * 60 * 1000)
	timeTo := currentTime

	return timeFrom, timeTo
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
