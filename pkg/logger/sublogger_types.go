package logger

import "io"

// Global vars related to the logger package
var (
	subLoggers = map[string]*SubLogger{}

	Global           *SubLogger
	SubsystemLogger  *SubLogger
	WebserverLogger  *SubLogger
	BackTester       *SubLogger
	InternetLogger   *SubLogger
	CommunicationMgr *SubLogger
	APIServerMgr     *SubLogger
	ConfigLogger     *SubLogger
	DatabaseLogger   *SubLogger
	DataHistory      *SubLogger
	StrategyLogger   *SubLogger
	OrderMgr         *SubLogger
	PortfolioMgr     *SubLogger
	SyncMgr          *SubLogger
	NTPLogger        *SubLogger
	WebsocketMgr     *SubLogger
	EventMgr         *SubLogger
	DispatchMgr      *SubLogger

	RequestSys  *SubLogger
	ExchangeSys *SubLogger

	RESTSys *SubLogger

	Ticker    *SubLogger
	OrderBook *SubLogger
	Trade     *SubLogger
	Fill      *SubLogger
)

// logFields is used to store data in a non-global and thread-safe manner
// so logs cannot be modified mid-log causing a data-race issue
type logFields struct {
	info   bool
	warn   bool
	debug  bool
	error  bool
	name   string
	output io.Writer
	logger Logger
}
