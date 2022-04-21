package coinbasepro

import "time"

func NewFeed() Feed {
	return Feed{
		Subscriptions: make(chan SubscriptionRequest, 1),
		Messages:      make(chan interface{}),
	}
}

type Feed struct {
	Subscriptions chan SubscriptionRequest
	Messages      chan interface{}
}

type L2Update struct {
	Type      string     `json:"type"`
	ProductId string     `json:"product_id"`
	Time      time.Time  `json:"time"`
	Changes   [][]string `json:"changes"`
}

type Heartbeat struct {
	Type        string `json:"type"`
	Sequence    string `json:"sequence"`
	LastTradeId string `json:"last_trade_id"`
	ProductId   string `json:"product_id"`
	Time        string `json:"time"`
}
