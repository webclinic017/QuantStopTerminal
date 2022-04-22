package engine

import (
	"github.com/quantstop/quantstopterminal/internal/config"
	"github.com/quantstop/quantstopterminal/internal/webserver"
	"sync"
	"time"
)

type Engine struct {
	*Version
	Config              *config.Config
	SubsystemRegistry   *ServiceRegistry
	DatabaseSubsystem   *DatabaseSubsystem
	NTPCheckerSubsystem *NTPCheckerSubsystem
	TraderSubsystem     *TraderSubsystem
	InternetSubsystem   *ConnectionMonitor
	Webserver           *webserver.Webserver
	SentimentAnalyzer   *SentimentAnalyzerSubsystem
	SubsystemWG         sync.WaitGroup
	Uptime              time.Time
}

const (
	DatabaseSubsystemName string = "database"
	NTPSubsystemName      string = "ntp_timekeeper"
	TraderSubsystemName   string = "active_trader"
	InternetCheckerName   string = "internet_monitor"
	SentimentAnalyzerName string = "sentiment_analyzer"
)

// engineMutex only locks and unlocks on engine creation functions
// as engine modifies global files, this protects the main bot creation
// functions from interfering with each other
var engineMutex sync.Mutex
