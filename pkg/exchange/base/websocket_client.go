package base

import (
	"github.com/gorilla/websocket"
	"io/ioutil"
)

type Websocket interface {
	Dial() (*websocket.Conn, error)
}

type WebsocketDialer struct {
	URL string
}

// Dial returns a connection to the FeedURL websocket.
func (w *WebsocketDialer) Dial() (*websocket.Conn, error) {
	var wsDialer websocket.Dialer
	wsConn, resp, err := wsDialer.Dial(w.URL, nil)
	if err != nil {
		return nil, err
	}
	_, _ = ioutil.ReadAll(resp.Body)
	_ = resp.Body.Close()
	return wsConn, nil
}
