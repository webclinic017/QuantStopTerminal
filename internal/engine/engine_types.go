package engine

import (
	"github.com/quantstop/quantstopterminal/internal/config"
	"sync"
	"time"
)

type Engine struct {
	*Version
	Config              *config.Config
	SubsystemRegistry   *ServiceRegistry
	DatabaseSubsystem   *DatabaseSubsystem
	NTPCheckerSubsystem *NTPCheckerSubsystem
	StrategySubsystem   *StrategySubsystem
	InternetSubsystem   *ConnectionMonitor
	WebserverSubsystem  *WebserverSubsystem
	SubsystemWG         sync.WaitGroup
	Uptime              time.Time
}

const (
	DatabaseSubsystemName string = "database"
	NTPSubsystemName      string = "ntp_timekeeper"
	StrategySubsystemName string = "strategy"
	WebserverName         string = "webserver"
	GRPCServerName        string = "grpc"
	GRPCProxyServerName   string = "grpc_proxy"
	InternetCheckerName   string = "internet_monitor"
)

// engineMutex only locks and unlocks on engine creation functions
// as engine modifies global files, this protects the main bot creation
// functions from interfering with each other
var engineMutex sync.Mutex
