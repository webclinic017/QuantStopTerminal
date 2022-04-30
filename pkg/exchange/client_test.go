package exchange

import (
	"context"
	"encoding/json"
	"fmt"
	sqlite "github.com/quantstop/quantstopterminal/internal/database/drivers/sqlite3"
	"github.com/quantstop/quantstopterminal/internal/database/models"
	"github.com/quantstop/quantstopterminal/internal/log"
	"github.com/quantstop/quantstopterminal/pkg/exchange/coinbasepro"
	"golang.org/x/sync/errgroup"
	"testing"
)

var CBPClient *coinbasepro.Client

func TestNewClient(t *testing.T) {

	var err error
	//dbConfig := *database.GenDefaultSettings()
	dbConn, err := sqlite.Connect("")

	e := models.CryptoExchange{}
	err = e.GetCryptoExchangeByName(dbConn.SQL, "coinbasepro")
	if err != nil {
		log.Error(log.TraderLogger, err)
		return
	}

	//Create a client instance
	CBPClient, err = coinbasepro.NewSandboxClient(
		&coinbasepro.Auth{
			Key:        e.AuthKey,
			Passphrase: e.AuthPassphrase,
			Secret:     e.AuthSecret,
		},
	)

	if err != nil {
		t.Errorf("Error creating new client: %v", err)
	}
}

func TestClient_ListAccounts(t *testing.T) {
	TestNewClient(t)
	accounts, err := CBPClient.ListAccounts(context.TODO())

	if err != nil {
		t.Errorf("Error listing accounts: %v", err)
	}

	print(prettyPrint(accounts))
}

func TestClient_GetAccount(t *testing.T) {
	TestNewClient(t)
	accounts, err := CBPClient.GetAccount(context.TODO(), "test1")

	if err != nil {
		t.Errorf("Error getting account: %v", err)
	}

	print(prettyPrint(accounts))
}

func TestClient_GetOrderbook(t *testing.T) {
	TestNewClient(t)
	accounts, err := CBPClient.GetOrderBook(context.TODO(), "BTC-USD")

	if err != nil {
		t.Errorf("Error getting account: %v", err)
	}

	print(prettyPrint(accounts))
}

func TestClient_GetAggregatedOrderbook(t *testing.T) {
	TestNewClient(t)
	accounts, err := CBPClient.GetAggregatedOrderBook(context.TODO(), "BTC-USD", coinbasepro.BookLevelTop50)

	if err != nil {
		t.Errorf("Error getting account: %v", err)
	}

	print(prettyPrint(accounts))
}

func TestClient_GetHistoricRates(t *testing.T) {
	TestNewClient(t)
	product, err := CBPClient.GetHistoricRates(context.TODO(), "BTC-USD", coinbasepro.HistoricRateFilter{
		Granularity: 60,
	})

	if err != nil {
		t.Errorf("Error getting historic rates: %v", err)
	}

	print(prettyPrint(product))
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func TestClient_WebsocketFeed(t *testing.T) {
	TestNewClient(t)
	ctx := context.TODO()

	// create a new subscription request
	prods := []coinbasepro.ProductID{"BTC-USD"}
	channelNames := []coinbasepro.ChannelName{
		coinbasepro.ChannelNameMatches,
	}
	channels := []coinbasepro.Channel{{
		Name:       "matches",
		ProductIDs: prods,
	}}

	subReq := coinbasepro.NewSubscriptionRequest(prods, channelNames, channels)
	//feed := coinbasepro.NewFeed()

	// dial connection
	wsConn, err := CBPClient.Websocket.Dial()
	if err != nil {
		//logger.Debugf(logger.StrategyLogger, "%v", err)
	}

	// subscription request must be sent within 5 seconds of open or socket will auto-close
	err = wsConn.WriteJSON(subReq)
	if err != nil {
		//logger.Debugf(logger.StrategyLogger, "%v", err)
	}

	// start reading messages
	messages := make(chan interface{})
	wg, ctx := errgroup.WithContext(ctx)
	wg.Go(func() error {
		defer close(messages)
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case messages <- func() interface{} {

				var trade coinbasepro.ProductTrade
				err := wsConn.ReadJSON(&trade)
				if err != nil {
					return err
				}

				tradeID := fmt.Sprintf("TradeID: %d | ", trade.TradeID)
				tradeTime := fmt.Sprintf("Time: %s | ", trade.Time.Time())
				tradePrice := fmt.Sprintf("Price: %f | ", trade.Price)
				tradeSize := fmt.Sprintf("Size: %f | ", trade.Size)
				tradeSide := fmt.Sprintf("Side: %s", trade.Side)

				//logger.Debugln(logger.StrategyLogger, "Trade received - " + tradeID + tradeTime + tradePrice + tradeSize + tradeSide)

				println("Trade received - " + tradeID + tradeTime + tradePrice + tradeSize + tradeSide)

				return nil
			}():
			}
		}
	})

	// todo: not updated for new websocket logic
	/*wg.Go(func() error {
		for message := range messages {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case feed.Messages <- message:
				fmt.Println("publish message on channel")
			default:
			}
		}
		return nil
	})*/
	_ = wg.Wait()

}
