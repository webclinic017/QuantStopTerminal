package coinbasepro

import (
	"time"
)

func NewFeed() Feed {
	return Feed{
		Subscriptions: make(chan SubscriptionRequest, 1),
		Messages:      make(chan []byte),
		Heartbeat:     make(chan HeartbeatMessage),
		Status:        make(chan StatusMessage),
		Ticker:        make(chan TickerMessage),
		TickerBatch:   make(chan TickerMessage),
		Level2:        make(chan L2UpdateMessage),
		Level2Snap:    make(chan L2SnapshotMessage),
		//Level2Batch:   make(chan L2Message),
		//User:	       make(chan HeartbeatMessage),
		Matches: make(chan MatchMessage),
		//Full:	       make(chan HeartbeatMessage),
		//Auction:	   make(chan HeartbeatMessage),
	}
}

type Feed struct {
	Subscriptions chan SubscriptionRequest
	Messages      chan []byte

	Heartbeat   chan HeartbeatMessage
	Status      chan StatusMessage
	Ticker      chan TickerMessage
	TickerBatch chan TickerMessage
	Level2      chan L2UpdateMessage
	Level2Snap  chan L2SnapshotMessage
	//Level2Batch   chan L2Message
	//User          chan User
	Matches chan MatchMessage
	//Full          chan Full
	//Auction       chan Auction
}

type HeartbeatMessage struct {
	Type        string    `json:"type"`
	Sequence    int       `json:"sequence"`
	LastTradeId int       `json:"last_trade_id"`
	ProductId   string    `json:"product_id"`
	Time        time.Time `json:"time"`
}

type StatusMessage struct {
	Type     string `json:"type"`
	Products []struct {
		Id             string      `json:"id"`
		BaseCurrency   string      `json:"base_currency"`
		QuoteCurrency  string      `json:"quote_currency"`
		BaseMinSize    string      `json:"base_min_size"`
		BaseMaxSize    string      `json:"base_max_size"`
		BaseIncrement  string      `json:"base_increment"`
		QuoteIncrement string      `json:"quote_increment"`
		DisplayName    string      `json:"display_name"`
		Status         string      `json:"status"`
		StatusMessage  interface{} `json:"status_message"`
		MinMarketFunds string      `json:"min_market_funds"`
		MaxMarketFunds string      `json:"max_market_funds"`
		PostOnly       bool        `json:"post_only"`
		LimitOnly      bool        `json:"limit_only"`
		CancelOnly     bool        `json:"cancel_only"`
		FxStablecoin   bool        `json:"fx_stablecoin"`
	} `json:"products"`
	Currencies []struct {
		Id            string      `json:"id"`
		Name          string      `json:"name"`
		MinSize       string      `json:"min_size"`
		Status        string      `json:"status"`
		StatusMessage interface{} `json:"status_message"`
		MaxPrecision  string      `json:"max_precision"`
		ConvertibleTo []string    `json:"convertible_to"`
		Details       struct {
		} `json:"details"`
	} `json:"currencies"`
}

type TickerMessage struct {
	Type      string    `json:"type"`
	Sequence  int       `json:"sequence"`
	ProductId string    `json:"product_id"`
	Price     string    `json:"price"`
	Open24H   string    `json:"open_24h"`
	Volume24H string    `json:"volume_24h"`
	Low24H    string    `json:"low_24h"`
	High24H   string    `json:"high_24h"`
	Volume30D string    `json:"volume_30d"`
	BestBid   string    `json:"best_bid"`
	BestAsk   string    `json:"best_ask"`
	Side      string    `json:"side"`
	Time      time.Time `json:"time"`
	TradeId   int       `json:"trade_id"`
	LastSize  string    `json:"last_size"`
}

type L2SnapshotMessage struct {
	Type      string     `json:"type"`
	ProductId string     `json:"product_id"`
	Bids      [][]string `json:"bids"`
	Asks      [][]string `json:"asks"`
}

type L2UpdateMessage struct {
	Type      string     `json:"type"`
	ProductId string     `json:"product_id"`
	Time      time.Time  `json:"time"`
	Changes   [][]string `json:"changes"`
}

type MatchMessage struct {
	Type         string    `json:"type"`
	TradeId      int       `json:"trade_id"`
	Sequence     int       `json:"sequence"`
	MakerOrderId string    `json:"maker_order_id"`
	TakerOrderId string    `json:"taker_order_id"`
	Time         time.Time `json:"time"`
	ProductId    string    `json:"product_id"`
	Size         string    `json:"size"`
	Price        string    `json:"price"`
	Side         string    `json:"side"`
}

type ChangeMessage struct {
	Type      string    `json:"type"`
	Time      time.Time `json:"time"`
	Sequence  int       `json:"sequence"`
	OrderId   string    `json:"order_id"`
	ProductId string    `json:"product_id"`
	NewSize   string    `json:"new_size"`
	OldSize   string    `json:"old_size"`
	Price     string    `json:"price"`
	Side      string    `json:"side"`
}
