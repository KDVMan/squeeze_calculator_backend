package services_symbol

import (
	services_symbol_interface "backend/internal/services/symbol/interface"
	services_websocket_interface "backend/internal/services/websocket/interface"
	services_storage_interface "backend/pkg/services/storage/interface"
)

type symbolServiceImplementation struct {
	storageService   func() services_storage_interface.StorageService
	websocketService func() services_websocket_interface.WebsocketService
}

func NewSymbolService(
	storageService func() services_storage_interface.StorageService,
	websocketService func() services_websocket_interface.WebsocketService,
) services_symbol_interface.SymbolService {
	return &symbolServiceImplementation{
		storageService:   storageService,
		websocketService: websocketService,
	}
}
