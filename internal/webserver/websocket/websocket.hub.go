package websocket

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/quantstop/quantstopterminal/internal"
	"github.com/quantstop/quantstopterminal/internal/log"
	"github.com/quantstop/quantstopterminal/internal/webserver/write"
	"github.com/quantstop/quantstopterminal/pkg/exchange"
	"github.com/quantstop/quantstopterminal/pkg/exchange/coinbasepro"
	"golang.org/x/sync/errgroup"
	"net/http"
)

// constant for 3 type actions
const (
	publish     = "publish"
	subscribe   = "subscribe"
	unsubscribe = "unsubscribe"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	db *sql.DB

	// Registered clients.
	clients map[*Client]bool

	// Client subscriptions
	Subscriptions []Subscription

	// Register requests from the clients.
	Register chan *Client

	// Unregister requests from clients.
	Unregister chan *Client
}

type Subscription struct {
	Exchange string
	Clients  []*Client
	SubChan  chan []byte
}

// Message is the type for a valid message from a client
type Message struct {
	Action     string `json:"action"`
	ExchangeID string `json:"exchange_id"`
	Message    string `json:"message"`
}

type MessageResponse struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

func NewHub(eng internal.IEngine) (*Hub, error) {
	db, err := eng.GetSQL()
	if err != nil {
		return nil, err
	}
	return &Hub{
		db:            db,
		clients:       make(map[*Client]bool),
		Subscriptions: make([]Subscription, 0),
		Register:      make(chan *Client),
		Unregister:    make(chan *Client),
	}, nil
}

func (h *Hub) Run(ctx context.Context) error {

	wg, ctx := errgroup.WithContext(ctx)

	wg.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case client := <-h.Register:
				h.clients[client] = true
			case client := <-h.Unregister:
				if _, ok := h.clients[client]; ok {
					h.Unsubscribe(client, "", true)
					h.RemoveClient(client)
					delete(h.clients, client)
					close(client.send)
				}
			default:
				// if there are subscriptions, send any messages on the SubChan, to clients
				if len(h.Subscriptions) != 0 {
					// loop through subscriptions
					for subIndex, sub := range h.Subscriptions {
						// if no clients close subscription, and continue to next subscription
						if len(sub.Clients) == 0 {
							h.Subscriptions = append(h.Subscriptions[:subIndex], h.Subscriptions[subIndex+1:]...)
							continue
						}
						// otherwise, publish any messages on the SubChan to subscribed clients
						select {
						case message := <-sub.SubChan:
							h.Publish(sub.Exchange, message)
						}

					}
				}
			}

		}
	})

	return wg.Wait()

}

func (h *Hub) Send(client *Client, message []byte) {
	client.send <- message
}

func (h *Hub) RemoveClient(client *Client) {
	// Read all subs
	for _, sub := range h.Subscriptions {
		// Read all clients
		for i := 0; i < len(sub.Clients); i++ {
			if client.ID == (sub.Clients)[i].ID {
				// If found, remove client
				if i == len(sub.Clients)-1 {
					// if it's stored as the last element, crop the array length
					sub.Clients = (sub.Clients)[:len(sub.Clients)-1]
				} else {
					// if it's stored in between elements, overwrite the element and reduce iterator to prevent out-of-bound
					sub.Clients = append((sub.Clients)[:i], (sub.Clients)[i+1:]...)
					i--
				}
			}
		}
	}
}

func (h *Hub) ProcessMessage(client *Client, messageType int, payload []byte) *Hub {
	m := Message{}
	if err := json.Unmarshal(payload, &m); err != nil {
		msgRes := MessageResponse{
			Type:    "error",
			Message: "error: invalid payload",
		}
		res, err := json.Marshal(msgRes)
		if err != nil {
			log.Error(log.Webserver, err)
		}
		h.Send(client, res)
		return h
	}
	log.Debugf(log.Webserver, "Client Message: %v", m)

	switch m.Action {
	case publish:
		h.Publish(m.ExchangeID, []byte(m.Message))
		break

	case subscribe:
		h.Subscribe(client, m.ExchangeID, m.Message)
		break

	case unsubscribe:
		h.Unsubscribe(client, m.ExchangeID, false)
		break

	default:
		msgRes := MessageResponse{
			Type:    "error",
			Message: "error: unrecognized action",
		}
		res, err := json.Marshal(msgRes)
		if err != nil {
			log.Error(log.Webserver, err)
		}
		h.Send(client, res)
		break
	}

	return h
}

func (h *Hub) Publish(topic string, message []byte) {
	var clients []*Client

	// get list of clients subscribed to topic
	for _, sub := range h.Subscriptions {
		if sub.Exchange == topic {
			clients = append(clients, sub.Clients...)
		}
	}

	// send to clients
	for _, client := range clients {
		h.Send(client, message)
	}
}

func (h *Hub) Subscribe(client *Client, exchange string, product string) {
	exist := false

	// find existing topics
	for _, sub := range h.Subscriptions {
		// if found, add client
		if sub.Exchange == exchange {
			exist = true
			sub.Clients = append(sub.Clients, client)
		}
	}

	// else, add new topic & add client to that topic
	if !exist {
		newClient := []*Client{
			client,
		}

		newTopic := Subscription{
			Exchange: exchange,
			Clients:  newClient,
			SubChan:  make(chan []byte),
		}

		h.Subscriptions = append(h.Subscriptions, newTopic)

		switch exchange {
		case "coinbasepro":
			h.RunCoinbaseproWebsocket(newTopic.SubChan, product)
		}

	}
}

func (h *Hub) Unsubscribe(client *Client, exch string, unsubAll bool) {
	// Read all topics
	for subIndex, sub := range h.Subscriptions {
		if !unsubAll && sub.Exchange == exch {
			// Read all topics' client
			for i := 0; i < len(sub.Clients); i++ {
				if client.ID == (sub.Clients)[i].ID {
					// If found, remove client
					if i == len(sub.Clients)-1 {
						// if it's stored as the last element, crop the array length
						sub.Clients = (sub.Clients)[:len(sub.Clients)-1]
					} else {
						// if it's stored in between elements, overwrite the element and reduce iterator to prevent out-of-bound
						sub.Clients = append((sub.Clients)[:i], (sub.Clients)[i+1:]...)
						i--
					}
				}
			}
		} else {
			// Read all topics' client
			for i := 0; i < len(sub.Clients); i++ {
				if client.ID == (sub.Clients)[i].ID {
					// If found, remove client
					if i == len(sub.Clients)-1 {
						// if it's stored as the last element, crop the array length
						sub.Clients = (sub.Clients)[:len(sub.Clients)-1]
					} else {
						// if it's stored in between elements, overwrite the element and reduce iterator to prevent out-of-bound
						sub.Clients = append((sub.Clients)[:i], (sub.Clients)[i+1:]...)
						i--
					}
				}
			}
		}

		// if that was last client
		if len(sub.Clients) == 0 {
			h.Subscriptions = append(h.Subscriptions[:subIndex], h.Subscriptions[subIndex+1:]...)
			channelNames := []coinbasepro.ChannelName{
				coinbasepro.ChannelNameHeartbeat,
				coinbasepro.ChannelNameLevel2,
				coinbasepro.ChannelNameMatches,
			}
			unsubReq := coinbasepro.NewUnsubscriptionRequest([]coinbasepro.ProductID{}, channelNames, []coinbasepro.Channel{})
			err := exchange.Coinbasepro.Conn.WriteJSON(unsubReq)
			if err != nil {
				log.Error(log.Webserver, err)
			}
		}
	}
}

func (h *Hub) RunCoinbaseproWebsocket(msgChan chan []byte, product string) {

	ctx := context.TODO()

	// create a new subscription request
	prods := []coinbasepro.ProductID{coinbasepro.ProductID(product)}
	channelNames := []coinbasepro.ChannelName{
		coinbasepro.ChannelNameHeartbeat,
		coinbasepro.ChannelNameLevel2,
	}
	channels := []coinbasepro.Channel{
		coinbasepro.Channel{
			Name:       coinbasepro.ChannelNameMatches,
			ProductIDs: []coinbasepro.ProductID{coinbasepro.ProductID(product)},
		},
	}
	subReq := coinbasepro.NewSubscriptionRequest(prods, channelNames, channels)
	feed := coinbasepro.NewFeed()
	wg, ctx := errgroup.WithContext(ctx)

	// start api client feed
	wg.Go(func() error {
		return exchange.Coinbasepro.Watch(ctx, subReq, feed)
	})

	// testing - Loop on Heartbeat channel
	wg.Go(func() error {
		for message := range feed.Heartbeat {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
				res, err := json.Marshal(message)
				if err != nil {
					log.Error(log.TraderLogger, err)
				}
				msgChan <- res
				log.Debugf(log.TraderLogger, "%s | %s", message.Type, message.Time.String())
			}
		}
		return nil
	})

	// testing - Loop on L2Channel channel
	wg.Go(func() error {
		for message := range feed.Level2 {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
				res, err := json.Marshal(message)
				if err != nil {
					log.Error(log.TraderLogger, err)
				}
				msgChan <- res
				log.Debugf(log.TraderLogger, "%s | %s", message.Type, message.Time.String())
			}
		}
		return nil
	})

	// testing - Loop on L2ChannelSnapshot channel
	wg.Go(func() error {
		for message := range feed.Level2Snap {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
				res, err := json.Marshal(message)
				if err != nil {
					log.Error(log.TraderLogger, err)
				}
				msgChan <- res
				log.Debugf(log.TraderLogger, "%s | %s", message.Type, message.ProductId)
			}
		}
		return nil
	})

	// testing - Loop on Matches channel
	wg.Go(func() error {
		for message := range feed.Matches {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
				res, err := json.Marshal(message)
				if err != nil {
					log.Error(log.TraderLogger, err)
				}
				msgChan <- res
				log.Debugf(log.TraderLogger, "%s | %s", message.Type, message.Time.String())
			}
		}
		return nil
	})

	_ = wg.Wait()
	return
}

var upgrader = websocket.Upgrader{
	//https://github.com/kataras/neffos/issues/11#issuecomment-520689681
	// todo this allows all origins ...
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// ServeWs handles websocket requests from the peer.
func ServeWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error(log.Webserver, err)
		write.Error(err)
		return
	}
	client := &Client{
		ID:   uuid.Must(uuid.NewRandom()).String(),
		hub:  hub,
		conn: conn,
		send: make(chan []byte, 256),
	}
	client.hub.Register <- client

	// greet the new client
	msgRes := MessageResponse{
		Type:    "welcome",
		Message: "QuantstopTerminal Websocket Server: Welcome! Your ID is " + client.ID,
	}
	res, err := json.Marshal(msgRes)
	if err != nil {
		log.Error(log.Webserver, err)
		write.Error(err)
		return
	}
	hub.Send(client, res)

	// Allow collection of memory referenced by the caller by doing all work in new goroutines.
	go client.writePump()
	go client.readPump()
}
