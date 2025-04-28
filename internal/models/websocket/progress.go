package models_websocket

import (
	enums_websocket "backend/internal/enums/websocket"
)

type ProgressChannelModel struct {
	Count  int64                           `json:"count"`
	Total  int64                           `json:"total"`
	Status enums_websocket.WebsocketStatus `json:"status"`
	Event  enums_websocket.WebsocketEvent  `json:"-"`
}
