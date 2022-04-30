package coinbasepro

import "github.com/quantstop/quantstopterminal/pkg/exchange/base"

type CoinbasePro struct {
	base.HttpAPI
	base.Websocket
}

func (c *CoinbasePro) GetHistoricCandles() {

}

func (c *CoinbasePro) GetRealtimeFeed() {

}
