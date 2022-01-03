package exchange

type IExchange interface {
	GetHistoricCandles()
	GetRealtimeFeed()
}
