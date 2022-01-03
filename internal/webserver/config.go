package webserver

type Config struct {
	Enabled             bool
	HttpListenAddr      string
	WebsocketListenAddr string
}
