package services_symbol

import (
	enums_symbol "backend/internal/enums/symbol"
	models_symbol "backend/internal/models/symbol"
	"errors"
	"github.com/adshao/go-binance/v2/futures"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

func (object *symbolServiceImplementation) Download(symbols []futures.Symbol) error {
	tx := object.storageService().DB().Begin()

	if tx.Error != nil {
		return tx.Error
	}

	for _, data := range symbols {
		status := getStatus(data.Status)

		if status == enums_symbol.SymbolStatusUnknown {
			continue
		}

		symbol := strings.ToUpper(data.Symbol)
		var symbolModel models_symbol.SymbolModel
		err := tx.Where("symbol = ?", symbol).First(&symbolModel).Error

		if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
			symbolModel = models_symbol.SymbolModel{
				Symbol: symbol,
				Group:  strings.ToUpper(data.QuoteAsset),
				Status: status,
				Limit:  getLimit(data),
			}
		} else if err != nil {
			tx.Rollback()

			return err
		} else {
			symbolModel.Status = status
			symbolModel.Limit = getLimit(data)
		}

		if err = tx.Save(&symbolModel).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Error; err != nil {
		tx.Rollback()

		return err
	}

	return tx.Commit().Error
}

func getStatus(status string) enums_symbol.SymbolStatus {
	switch status {
	case "TRADING":
		return enums_symbol.SymbolStatusActive
	case "BREAK", "CLOSE":
		return enums_symbol.SymbolStatusInactive
	default:
		return enums_symbol.SymbolStatusUnknown
	}
}

func getLimit(data futures.Symbol) models_symbol.SymbolLimitModel {
	var limitModel models_symbol.SymbolLimitModel

	lotSizeFilters := data.LotSizeFilter()
	limitModel.LeftMin, _ = strconv.ParseFloat(lotSizeFilters.MinQuantity, 64)
	limitModel.LeftMax, _ = strconv.ParseFloat(lotSizeFilters.MaxQuantity, 64)
	limitModel.LeftStep, _ = strconv.ParseFloat(lotSizeFilters.StepSize, 64)

	step := strconv.FormatFloat(limitModel.LeftStep, 'f', -1, 64)
	decimal := strings.Split(step, ".")

	if len(decimal) > 1 {
		limitModel.LeftPrecision = int(len(decimal[1]))
	}

	notionalFilter := data.MinNotionalFilter()

	limitModel.RightMin, _ = strconv.ParseFloat(notionalFilter.Notional, 64)
	limitModel.RightMax = 0

	limitModel.Precision = data.PricePrecision
	limitModel.TickSize, _ = strconv.ParseFloat(data.PriceFilter().TickSize, 64)

	return limitModel
}
