package services_symbol_list

import (
	services_symbol_list_interface "backend/internal/services/symbol_list/interface"
	services_storage_interface "backend/pkg/services/storage/interface"
)

type symbolListServiceImplementation struct {
	storageService func() services_storage_interface.StorageService
}

func NewSymbolListService(
	storageService func() services_storage_interface.StorageService,
) services_symbol_list_interface.SymbolListService {
	return &symbolListServiceImplementation{
		storageService: storageService,
	}
}
