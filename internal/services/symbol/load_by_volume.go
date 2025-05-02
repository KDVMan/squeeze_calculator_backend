package services_symbol

import (
	enums_symbol "backend/internal/enums/symbol"
	models_symbol "backend/internal/models/symbol"
)

func (object *symbolServiceImplementation) LoadByVolume(volume int, group string) ([]*models_symbol.SymbolModel, error) {
	var symbols []*models_symbol.SymbolModel

	if err := object.storageService().DB().
		Model(&models_symbol.SymbolModel{}).
		Where("status = ?", enums_symbol.SymbolStatusActive).
		Where("statistic_volume > ?", volume).
		Where(`"group" = ?`, group).
		Order("symbol").
		Find(&symbols).
		Error; err != nil {
		return nil, err
	}

	return symbols, nil
}
