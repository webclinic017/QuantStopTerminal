package grpcserver

// Config stores the gRPC settings
type Config struct {
	Enabled                bool
	Initialized            bool
	ListenAddress          string
	GRPCProxyEnabled       bool
	GRPCProxyListenAddress string
	TimeInNanoSeconds      bool
	Username               string
	Password               string
}
