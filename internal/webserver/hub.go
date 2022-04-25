package webserver

import (
	"encoding/json"
	"fmt"
	"github.com/quantstop/quantstopterminal/internal/log"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	Broadcast chan []byte

	// Register requests from the clients.
	Register chan *Client

	// Unregister requests from clients.
	Unregister chan *Client

	// Inbound messages from trader
	TradeChannel chan []byte
}

func newHub() *Hub {
	return &Hub{
		Broadcast:    make(chan []byte),
		Register:     make(chan *Client),
		Unregister:   make(chan *Client),
		clients:      make(map[*Client]bool),
		TradeChannel: make(chan []byte),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.Register:
			log.Debugln(log.Webserver, "Register")
			h.clients[client] = true
		case client := <-h.Unregister:
			log.Debugln(log.Webserver, "Unregister")
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.Broadcast:
			log.Debugln(log.Webserver, "Broadcast")
			for client := range h.clients {
				select {
				case client.send <- message:
					msg := string(message)
					switch msg {
					case "numClients":
						numClients := len(h.clients)
						returnString := fmt.Sprintf("hub clients: %v", numClients)
						s, _ := json.Marshal(returnString)
						client.send <- s
					}
					log.Debugln(log.Webserver, "sending client message")

				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		case message := <-h.TradeChannel:
			log.Debugln(log.Webserver, "TradeChannel")
			for client := range h.clients {
				client.send <- message
			}
		}

	}
}
