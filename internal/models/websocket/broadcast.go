package models_websocket

import (
	enums_websocket "backend/internal/enums/websocket"
)

type BroadcastChannelModel struct {
	Event enums_websocket.WebsocketEvent `json:"event"`
	Data  interface{}                    `json:"data"`
}
