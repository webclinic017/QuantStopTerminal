package coinbasepro

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/sync/errgroup"
	"testing"
)

var Coinbasepro *Client

func TestNewClient(t *testing.T) {

	var err error

	//Create a client instance
	Coinbasepro, err = NewSandboxClient(
		&Auth{
			Key:        "9e0699f3e3819c8a750eb68f080ae563",
			Passphrase: "umjjjkor5jb",
			Secret:     "52jLKk2/N4GHXWDFtcQ5nldhWfYnHo07jb5xhAJ7V/oPAYHZDsgOE6/lI4eFkQw1peWjZHdyelmtPIcUDu77hA=="},
	)

	if err != nil {
		t.Errorf("Error creating new client: %v", err)
	}
}

func TestClient_ListAccounts(t *testing.T) {
	TestNewClient(t)
	accounts, err := Coinbasepro.ListAccounts(context.TODO())

	if err != nil {
		t.Errorf("Error listing accounts: %v", err)
	}

	print(prettyPrint(accounts))
}

func TestClient_GetAccount(t *testing.T) {
	TestNewClient(t)
	accounts, err := Coinbasepro.GetAccount(context.TODO(), "Default Portfolio")

	if err != nil {
		t.Errorf("Error getting account: %v", err)
	}

	print(prettyPrint(accounts))
}

func TestClient_GetHistoricRates(t *testing.T) {
	TestNewClient(t)
	product, err := Coinbasepro.GetHistoricRates(context.TODO(), "BTC-USD", HistoricRateFilter{
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
	prods := []ProductID{"BTC-USD"}
	channelNames := []ChannelName{
		ChannelNameMatches,
	}
	channels := []Channel{{
		Name:       "matches",
		ProductIDs: prods,
	}}

	subReq := NewSubscriptionRequest(prods, channelNames, channels)
	feed := NewFeed()

	// dial connection
	wsConn, err := Coinbasepro.Websocket.Dial()
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

				var trade ProductTrade
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

	wg.Go(func() error {
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
	})
	_ = wg.Wait()

}
