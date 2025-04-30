package enums_websocket

type WebsocketEvent string

const (
	WebsocketEventExchangeLimits    WebsocketEvent = "exchangeLimits"
	WebsocketEventSymbolList        WebsocketEvent = "symbolList"
	WebsocketEventCurrentPrice      WebsocketEvent = "currentPrice"
	WebsocketEventCalculateProgress WebsocketEvent = "calculateProgress"
	WebsocketEventCalculateResult   WebsocketEvent = "calculateResult"
	WebsocketEventBotList           WebsocketEvent = "botList"
	WebsocketEventBot               WebsocketEvent = "bot"
)
