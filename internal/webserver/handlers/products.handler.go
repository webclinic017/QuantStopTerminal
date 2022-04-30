package handlers

import (
	"context"
	"github.com/quantstop/quantstopterminal/internal"
	"github.com/quantstop/quantstopterminal/internal/database/models"
	"github.com/quantstop/quantstopterminal/internal/log"
	"github.com/quantstop/quantstopterminal/internal/webserver/errors"
	"github.com/quantstop/quantstopterminal/internal/webserver/router"
	"github.com/quantstop/quantstopterminal/internal/webserver/write"
	"github.com/quantstop/quantstopterminal/pkg/exchange"
	"github.com/quantstop/quantstopterminal/pkg/exchange/coinbasepro"
	"net/http"
	"strconv"
)

type getExchangesResponse struct {
	Type      string              `json:"type"`
	Exchanges []SupportedExchange `json:"exchanges"`
}

// todo: move this to an exchange manager
type SupportedExchange struct {
	ID string `json:"id"`
}

var SupportedExchanges = []SupportedExchange{
	{"coinbasepro"},
	{"yfinance"},
}

func GetExchanges(bot internal.IEngine, user *models.User, w http.ResponseWriter, r *http.Request) http.HandlerFunc {

	res := getExchangesResponse{
		Type:      "getExchanges",
		Exchanges: SupportedExchanges,
	}

	return write.JSON(res)
}

type getCandleResponse struct {
	Type                      string `json:"type"`
	coinbasepro.HistoricRates `json:"candles"`
}

// GetCandles
/* Historic rates for a product.
   Rates are returned in grouped buckets.
   Candle schema is of the form [timestamp, price_low, price_high, price_open, price_close]
   Request: GET("/api/exchanges/([^/]+)/products/([^/]+)/candles")
   Params:
	- granularity (string, required) {60, 300, 900, 3600, 21600, 86400}
	- start (string, optional)
	- end (string, optional)
*/
func GetCandles(bot internal.IEngine, user *models.User, w http.ResponseWriter, r *http.Request) http.HandlerFunc {

	exchangeName := router.GetField(r, 0)
	productName := router.GetField(r, 1)

	granularity, _ := strconv.ParseInt(r.URL.Query().Get("granularity"), 0, 0)
	if granularity != 0 {
		// ... process it, will be the first (only) if multiple were given
		// note: if they pass in like ?param1=&param2= param1 will also be "" :|
		// todo: process params
	}

	switch exchangeName {
	case "coinbasepro":
		// todo: maybe something like this, where each exchange client implements the interface?
		//bot.GetExchange(coinbase).GetProduct()
		if exchange.Coinbasepro == nil {
			log.Debugln(log.Webserver, "error: products.handler coinbase client is nil")
		}
		candles, err := exchange.Coinbasepro.GetHistoricRates(
			context.TODO(),
			coinbasepro.ProductID(productName),
			coinbasepro.HistoricRateFilter{
				Granularity: coinbasepro.Timeslice(granularity),
			},
		)
		if err != nil {
			log.Debugf(log.Webserver, "error getting candles: %v", err)
		}
		resp := getCandleResponse{
			Type:          "getProductCandles",
			HistoricRates: candles,
		}
		return write.JSON(resp)
	}

	return write.Error(errors.InternalError)
}

type getProductRequest struct {
	ExchangeID string `json:"exchange_id"`
}

type getProductResponse struct {
	Type     string                `json:"type"`
	Products []coinbasepro.Product `json:"products"`
}

// GetProducts handles GET requests for "/api/exchanges/([^/]+)/products/" example: GET "/api/exchanges/coinbase/products
func GetProducts(bot internal.IEngine, user *models.User, w http.ResponseWriter, r *http.Request) http.HandlerFunc {

	/*decoder := json.NewDecoder(r.Body)
	req := getProductRequest{}
	err := decoder.Decode(&req)
	if err != nil || &req == nil {
		return write.Error(errors.NoJSONBody)
	}*/

	//id, _ := strconv.Atoi(router.GetField(r, 0))
	slug := router.GetField(r, 0)
	name := ""

	for _, e := range SupportedExchanges {
		if e.ID == slug {
			name = e.ID
			break
		}
	}

	switch name {
	case "coinbasepro":
		// todo: maybe something like this, where each exchange client implements the interface?
		//bot.GetExchange(coinbase).GetProduct()
		if exchange.Coinbasepro == nil {
			log.Debugln(log.Webserver, "error: products.handler coinbase client is nil")
		}
		products, err := exchange.Coinbasepro.ListProducts(context.TODO())
		if err != nil {
			log.Debugf(log.Webserver, "error getting candles: %v", err)
		}
		resp := getProductResponse{
			Type:     "getProducts",
			Products: products,
		}
		return write.JSON(resp)
	}

	return write.Error(errors.InternalError)
}
