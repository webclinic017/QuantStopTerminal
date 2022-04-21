package coinbasepro

/*
Coinbasepro developer documentation: https://docs.cloud.coinbase.com/exchange/docs/channels
*/

import (
	"context"
	"github.com/quantstop/quantstopterminal/pkg/exchange/base"
	"golang.org/x/sync/errgroup"
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

type Client struct {
	base.HttpAPI
	base.Websocket
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
	}, nil
}

func (c *Client) Close() error {
	return c.Close()
}

// Watch provides a feed of real-time market data updates for orders and trades.
func (c *Client) Watch(ctx context.Context, subscriptionRequest SubscriptionRequest, feed Feed) (capture error) {
	wsConn, err := c.Websocket.Dial()
	if err != nil {
		return err
	}
	// subscription request must be sent within 5 seconds of open or socket will auto-close
	err = wsConn.WriteJSON(subscriptionRequest)

	if err != nil {
		return err
	}
	return c.watch(ctx, wsConn, feed)
}

type jsonReader interface {
	ReadJSON(v interface{}) error
}

func (c *Client) watch(ctx context.Context, r jsonReader, feed Feed) (capture error) {
	messages := make(chan interface{})
	wg, ctx := errgroup.WithContext(ctx)
	wg.Go(func() error {
		defer close(messages)
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
				// TODO: Does this prevent blocking?
			case messages <- func() interface{} {
				// TODO: Does message have a real structure
				var message interface{}
				err := r.ReadJSON(&message)
				if err != nil {
					return err
				}
				//log.Debugf(log.TraderLogger, "receive message on socket. %v", message)
				return message
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
				//logrus.Debug("publish message on channel")
				//log.Debugln(log.TraderLogger, "publish message on channel")
				//log.Debugf(log.TraderLogger, "publishing message on channel. %v", message)
			default:
			}
		}
		return nil
	})
	return wg.Wait()
}
