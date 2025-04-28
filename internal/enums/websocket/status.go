package enums_websocket

type WebsocketStatus string

const (
	WebsocketStatusProgress WebsocketStatus = "progress"
	WebsocketStatusDone     WebsocketStatus = "done"
	WebsocketStatusStop     WebsocketStatus = "stop"
	WebsocketStatusError    WebsocketStatus = "error"
)
