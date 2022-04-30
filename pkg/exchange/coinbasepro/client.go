package coinbasepro

/*
Coinbasepro developer documentation: https://docs.cloud.coinbase.com/exchange/docs/channels
*/

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/quantstop/quantstopterminal/pkg/exchange/base"
	"golang.org/x/sync/errgroup"
	"log"
	"strings"
)

const (

	// Base URL's
	coinbaseproAPIURL       = "https://api.pro.coinbase.com/"
	coinbaseproWebsocketURL = "wss://ws-feed.exchange.coinbase.com"

	// Base URL's for sandbox environment
	coinbaseproSandboxWebsiteURL   = "https://public.sandbox.exchange.coinbase.com/"
	coinbaseproSandboxRestAPIURL   = "https://api-public.sandbox.exchange.coinbase.com/"
	coinbaseproSandboxWebsocketURL = "wss://ws-feed-public.sandbox.exchange.coinbase.com"
	coinbaseproSandboxFixAPIURL    = "tcp+ssl://fix-public.sandbox.exchange.coinbase.com:4198"

	// Endpoints
	coinbaseproAccounts                = "accounts"
	coinbaseproProducts                = "products"
	coinbaseproOrderbook               = "book"
	coinbaseproTicker                  = "ticker"
	coinbaseproTrades                  = "trades"
	coinbaseproHistory                 = "candles"
	coinbaseproStats                   = "stats"
	coinbaseproCurrencies              = "currencies"
	coinbaseproLedger                  = "ledger"
	coinbaseproHolds                   = "holds"
	coinbaseproOrders                  = "orders"
	coinbaseproFills                   = "fills"
	coinbaseproTransfers               = "transfers"
	coinbaseproReports                 = "reports"
	coinbaseproTime                    = "time"
	coinbaseproMarginTransfer          = "profiles/margin-transfer"
	coinbaseproPosition                = "position"
	coinbaseproPositionClose           = "position/close"
	coinbaseproPaymentMethod           = "payment-methods"
	coinbaseproPaymentMethodDeposit    = "deposits/payment-method"
	coinbaseproDepositCoinbase         = "deposits/coinbase-account"
	coinbaseproWithdrawalPaymentMethod = "withdrawals/payment-method"
	coinbaseproWithdrawalCoinbase      = "withdrawals/coinbase"
	coinbaseproWithdrawalCrypto        = "withdrawals/crypto"
	coinbaseproCoinbaseAccounts        = "coinbase-accounts"
	coinbaseproTrailingVolume          = "users/self/trailing-volume"
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

type Client struct {
	base.HttpAPI
	base.Websocket
	Conn *websocket.Conn
}

func query(params []string) string {
	if len(params) == 0 {
		return ""
	}
	return "?" + strings.Join(params, "&")
}

// NewClient creates a high-level Coinbase Pro API client.
func NewClient(auth *Auth) (*Client, error) {
	apiClient, err := NewAPIClient(auth)
	if err != nil {
		return nil, err
	}
	return &Client{
		apiClient,
		&base.WebsocketDialer{
			URL: apiClient.feedURL.String(),
		},
		&websocket.Conn{},
	}, nil
}

func NewSandboxClient(auth *Auth) (*Client, error) {
	apiClient, err := NewSandboxAPIClient(auth)
	if err != nil {
		return nil, err
	}
	return &Client{
		apiClient,
		&base.WebsocketDialer{
			URL: apiClient.feedURL.String(),
		},
		&websocket.Conn{},
	}, nil
}

func (c *Client) Close() error {
	return c.Close()
}

// Watch provides a feed of real-time market data updates for orders and trades.
func (c *Client) Watch(ctx context.Context, subscriptionRequest SubscriptionRequest, feed Feed) (capture error) {
	var err error
	c.Conn, err = c.Websocket.Dial()
	if err != nil {
		return err
	}
	// subscription request must be sent within 5 seconds of open or socket will auto-close
	err = c.Conn.WriteJSON(subscriptionRequest)

	if err != nil {
		return err
	}
	return c.watch(ctx, feed)
}

type jsonReader interface {
	ReadJSON(v interface{}) error
}

func (c *Client) watch(ctx context.Context, feed Feed) (capture error) {

	messages := make(chan []byte)
	wg, ctx := errgroup.WithContext(ctx)

	// read messages from coinbase
	wg.Go(func() error {
		defer close(messages)
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case messages <- func() []byte {
				_, message, err := c.Conn.ReadMessage()
				if err != nil {
					if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
						log.Printf("coinbase websocket read error: %v", err)
					}
				}
				message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
				return message
			}():
			}
		}
	})

	// testing
	wg.Go(func() error {
		for message := range messages {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
				var messageInterface map[string]interface{}
				if err := json.Unmarshal(message, &messageInterface); err != nil {
					return err
				}

				msgType := MessageType(messageInterface["type"].(string))
				switch msgType {
				case MessageTypeError:
					/*errMsg := Error{}
					if err := json.Unmarshal(message, &errMsg); err != nil {
						return err
					}
					feed.Error <- errMsg*/
				case MessageTypeL2Update:
					l2up := L2UpdateMessage{}
					if err := json.Unmarshal(message, &l2up); err != nil {
						return err
					}
					feed.Level2 <- l2up
				case MessageTypeSnapshot:
					l2snap := L2SnapshotMessage{}
					if err := json.Unmarshal(message, &l2snap); err != nil {
						return err
					}
					feed.Level2Snap <- l2snap
				case MessageTypeHeartbeat:
					hbMsg := HeartbeatMessage{}
					if err := json.Unmarshal(message, &hbMsg); err != nil {
						return err
					}
					feed.Heartbeat <- hbMsg
				case MessageTypeMatch:
					match := MatchMessage{}
					if err := json.Unmarshal(message, &match); err != nil {
						return err
					}
					feed.Matches <- match
				case MessageTypeSubscriptions:
					/*match := SubscriptionResponse{}
					if err := json.Unmarshal(message, &match); err != nil {
						return err
					}
					feed.Matches <- match*/
				case MessageTypeTicker:

				}
			}
		}
		return nil
	})

	return wg.Wait()

}
