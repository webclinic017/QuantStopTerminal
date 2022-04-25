package trader

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/d5/tengo/v2"
	"github.com/quantstop/quantstopterminal/internal/database/models"
	"github.com/quantstop/quantstopterminal/internal/log"
	"github.com/quantstop/quantstopterminal/internal/webserver"
	"github.com/quantstop/quantstopterminal/pkg/exchange/coinbasepro"
	"golang.org/x/sync/errgroup"
	"time"
)

var script = tengo.NewScript([]byte(
	`each := func(seq, fn) {
    for x in seq { fn(x) }
}

sum := 0
mul := 1
each([a, b, c, d], func(x) {
    sum += x
    mul *= x
})`))

var Coinbasepro *coinbasepro.Client

func Run(db *sql.DB, hub *webserver.Hub) {
	/*log.Debugln(log.TraderLogger, "starting workers ...")
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go Worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}*/

	var err error

	exchange := models.CryptoExchange{}
	err = exchange.GetCryptoExchangeByName(db, "coinbasepro")
	if err != nil {
		log.Error(log.TraderLogger, err)
		return
	}

	//Create a client instance
	Coinbasepro, err = coinbasepro.NewSandboxClient(
		&coinbasepro.Auth{
			Key:        exchange.AuthKey,
			Passphrase: exchange.AuthPassphrase,
			Secret:     exchange.AuthSecret,
		},
	)
	if err != nil {
		log.Error(log.TraderLogger, err)
		return
	}

	go testReadWebsocketTrades(hub)
}

func testReadWebsocketTrades(hub *webserver.Hub) {
	ctx := context.TODO()

	// create a new subscription request
	prods := []coinbasepro.ProductID{"BTC-USD"}
	channelNames := []coinbasepro.ChannelName{
		coinbasepro.ChannelNameHeartbeat,
		coinbasepro.ChannelNameMatches,
		coinbasepro.ChannelNameLevel2,
	}
	channels := []coinbasepro.Channel{
		{
			Name:       coinbasepro.ChannelNameHeartbeat,
			ProductIDs: prods,
		},
		{
			Name:       coinbasepro.ChannelNameMatches,
			ProductIDs: prods,
		},
		{
			Name:       coinbasepro.ChannelNameLevel2,
			ProductIDs: prods,
		},
	}

	subReq := coinbasepro.NewSubscriptionRequest(prods, channelNames, channels)
	feed := coinbasepro.NewFeed()

	wg, ctx := errgroup.WithContext(ctx)
	wg.Go(func() error {
		return Coinbasepro.Watch(ctx, subReq, feed)
	})

	wg.Go(func() error {
		for message := range feed.Messages {

			select {
			case <-ctx.Done():
				return ctx.Err()
			case feed.Messages <- message:

			default:

				var inInterface map[string]interface{}
				inrec, _ := json.Marshal(message)
				err := json.Unmarshal(inrec, &inInterface)
				if err != nil {
					return err
				}

				hub.TradeChannel <- inrec

				switch v := message.(type) {

				case map[string]interface{}:
					/*o.Pair = "BTCEUR"
					o.Timestamp = v["Timestamp"].(string)
					o.Broker = "Bitstamp"*/
					// here again you need to type assert correctly and it may be a map again
					//o.Asks = v["Asks"].(OrderBookItem)
					//o.Bids = v["Bids"].(OrderBookItem)

					switch v["type"] {
					case "l2update":
						tradeID := fmt.Sprintf("ProductID: %s | ", v["product_id"].(string))
						tradeTime := fmt.Sprintf("Time: %s | ", v["time"].(string))

						tradeSide := ""
						tradePrice := ""
						tradeSize := ""

						switch s := v["changes"].(type) {
						case []interface{}:

							for _, x := range s {
								switch z := x.(type) {
								case []interface{}:
									tradeSide = fmt.Sprintf("Side: %s", z[0].(string))
									tradePrice = fmt.Sprintf("Price: %s | ", z[1].(string))
									tradeSize = fmt.Sprintf("Size: %s | ", z[2].(string))
								}
							}
						}

						log.Debugln(log.TraderLogger, "L2Update | "+tradeID+tradeTime+tradePrice+tradeSize+tradeSide)

					case "heartbeat":
						tradeID := fmt.Sprintf("Sequence: %v | ", v["sequence"].(float64))
						tradeTime := fmt.Sprintf("LastTradeId: %v | ", v["last_trade_id"].(float64))
						tradePrice := fmt.Sprintf("ProductId: %s | ", v["product_id"].(string))
						tradeSize := fmt.Sprintf("Time: %s | ", v["time"].(string))

						log.Debugln(log.TraderLogger, "Heartbeat | "+tradeID+tradeTime+tradePrice+tradeSize)

					/*case "matches":
					tradeID := fmt.Sprintf("TradeID: %d | ", v["sequence"].(float64))
					tradeTime := fmt.Sprintf("Time: %s | ", v["sequence"].(float64))
					tradePrice := fmt.Sprintf("Price: %f | ", v["sequence"].(float64))
					tradeSize := fmt.Sprintf("Size: %f | ", v["sequence"].(float64))
					tradeSide := fmt.Sprintf("Side: %s", v["sequence"].(float64))

					if tradeID != 0 {
						log.Debugln(log.TraderLogger, "Matches |" + tradeID + tradeTime + tradePrice + tradeSize + tradeSide)
					}
					*/
					default:
						log.Debugf(log.TraderLogger, "%v", v["type"].(string))

					}

				}

			}
		}
		return nil
	})
	_ = wg.Wait()
	return

}

func Worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		//fmt.Println("worker", id, "started  job", j)
		log.Debugf(log.TraderLogger, "worker %v started job %v", id, j)
		time.Sleep(time.Second)

		//fmt.Println("worker", id, "finished job", j)
		log.Debugf(log.TraderLogger, "worker %v finished job %v", id, j)
		results <- j * 2
	}
}
