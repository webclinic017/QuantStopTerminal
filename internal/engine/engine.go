package engine

import (
	"errors"
	"fmt"
	"github.com/quantstop/quantstopterminal/internal/config"
	"github.com/quantstop/quantstopterminal/internal/qstlog"
	"github.com/quantstop/quantstopterminal/internal/webserver"
	"github.com/quantstop/quantstopterminal/pkg/system"
	"runtime"
	"strings"
	"time"
)

// Create creates a new instance of the engine
func Create(config *config.Config) (*Engine, error) {

	engineMutex.Lock()
	defer engineMutex.Unlock()

	if config == nil {
		return nil, errors.New("engine: config is nil")
	}

	var bot Engine
	var err error

	// Set the bot config
	bot.Config = config

	// Set the max processors for go
	err = system.AdjustGoMaxProcs(bot.Config.GoMaxProcessors)
	if err != nil {
		return nil, fmt.Errorf("unable to adjust runtime GOMAXPROCS value. Err: %s", err)
	}

	return &bot, nil
}

func (bot *Engine) Initialize() error {

	if bot == nil {
		return errors.New("engine instance is nil")
	}

	engineMutex.Lock()
	defer engineMutex.Unlock()

	//var err error

	// Create new subsystem registry
	bot.SubsystemRegistry = NewServiceRegistry()

	// Initialize database subsystem
	/*if bot.Config.Database.Enabled {

		// Create and init database subsystem
		bot.DatabaseSubsystem = &DatabaseSubsystem{Subsystem: Subsystem{}}
		if err = InitSubsystem(bot.DatabaseSubsystem, DatabaseSubsystemName, bot.Config); err != nil {
			logger.Errorf(logger.Global, "Database subsystem unable to initialize: %v", err)
		}

		// Register database subsystem
		if err = bot.SubsystemRegistry.RegisterService(bot.DatabaseSubsystem); err != nil {
			logger.Errorf(logger.Global, "Database subsystem unable to register: %v", err)
		}

	}*/

	// Initialize webserver subsystem
	//initWebserverSubsystem(bot)

	// Initialize ntp checker subsystem
	/*if bot.Config.NTP.Enabled {

		// Create and init ntp checker subsystem
		bot.NTPCheckerSubsystem = &NTPCheckerSubsystem{Subsystem: Subsystem{}}
		if err = InitSubsystem(bot.NTPCheckerSubsystem, NTPSubsystemName, bot.Config); err != nil {
			logger.Errorf(logger.Global, "NTP subsystem unable to initialize: %v", err)
		}

		// Register ntp checker subsystem
		if err = bot.SubsystemRegistry.RegisterService(bot.NTPCheckerSubsystem); err != nil {
			logger.Errorf(logger.Global, "NTP subsystem unable to register: %v", err)
		}

	}*/

	// Initialize strategy subsystem
	/*if bot.Config.Strategy.Enabled {

		// Create and init strategy subsystem
		bot.StrategySubsystem = &StrategySubsystem{Subsystem: Subsystem{}}
		if err = InitSubsystem(bot.StrategySubsystem, StrategySubsystemName, bot.Config); err != nil {
			logger.Errorf(logger.Global, "Strategy subsystem unable to initialize: %v", err)
		}

		// Register strategy subsystem
		if err = bot.SubsystemRegistry.RegisterService(bot.StrategySubsystem); err != nil {
			logger.Errorf(logger.Global, "Strategy subsystem unable to register: %v", err)
		}

	}*/

	// Initialize internet checker subsystem
	/*if bot.Config.Internet.Enabled {

		// Create and init internet checker subsystem
		bot.InternetSubsystem = &ConnectionMonitor{Subsystem: Subsystem{}}
		if err = InitSubsystem(bot.InternetSubsystem, InternetCheckerName, bot.Config); err != nil {
			logger.Errorf(logger.Global, "Internet checker subsystem unable to initialize: %v", err)
		}

		// Register internet checker subsystem
		if err = bot.SubsystemRegistry.RegisterService(bot.InternetSubsystem); err != nil {
			logger.Errorf(logger.Global, "Internet checker subsystem unable to register: %v", err)
		}

	}*/

	return nil
}

func initWebserverSubsystem(bot *Engine) {
	// Create and init webserver subsystem
	bot.WebserverSubsystem = &WebserverSubsystem{Subsystem: Subsystem{}}
	if err := InitSubsystem(bot.WebserverSubsystem, WebserverName, bot.Config); err != nil {
		qstlog.Errorf(qstlog.Global, "Webserver subsystem unable to initialize: %v", err)
	}

	// Register webserver subsystem
	if err := bot.SubsystemRegistry.RegisterService(bot.WebserverSubsystem); err != nil {
		qstlog.Errorf(qstlog.Global, "Webserver subsystem unable to register: %v", err)
	}
}

// Run start the newly created instance of the engine
func (bot *Engine) Run() error {

	if bot == nil {
		return errors.New("engine instance is nil")
	}

	engineMutex.Lock()
	defer engineMutex.Unlock()

	// Set the current uptime to now
	bot.Uptime = time.Now()

	// Start all subsystems
	bot.SubsystemRegistry.StartAll(&bot.SubsystemWG)

	// start web server
	webserver.StartHttpServer(bot.Config)

	// Print some info
	qstlog.Debugf(qstlog.Global, "QuantStopTerminal started.\n")
	qstlog.Debugf(qstlog.Global,
		"Using %d out of %d logical processors for runtime performance\n",
		runtime.GOMAXPROCS(-1), runtime.NumCPU())

	return nil
}

// Stop stops the running instance of the engine
func (bot *Engine) Stop() {

	engineMutex.Lock()
	defer engineMutex.Unlock()

	qstlog.Debugln(qstlog.Global, "Engine shutting down..")

	// Stop all subsystems
	bot.SubsystemRegistry.StopAll()

	// Wait for subsystems to gracefully shutdown
	bot.SubsystemWG.Wait()
	if err := qstlog.CloseLogger(); err != nil {
		fmt.Printf("Failed to close logger. Error: %v\n", err)
	}

}

// GetSubsystemsStatus returns the status of all engine subsystems
func (bot *Engine) GetSubsystemsStatus() map[string]bool {
	return map[string]bool{
		DatabaseSubsystemName: bot.DatabaseSubsystem.isRunning(),
		NTPSubsystemName:      bot.NTPCheckerSubsystem.isRunning(),
		StrategySubsystemName: bot.StrategySubsystem.isRunning(),
		InternetCheckerName:   bot.InternetSubsystem.isRunning(),
	}
}

// SetSubsystem enables or disables an engine subsystem
func (bot *Engine) SetSubsystem(subSystemName string, enable bool) error {
	if bot == nil {
		return errors.New("engine instance is nil")
	}

	if bot.Config == nil {
		return errNilConfig
	}

	var err error
	switch strings.ToLower(subSystemName) {

	case DatabaseSubsystemName:
		if enable {
			if bot.DatabaseSubsystem == nil {
				err = InitSubsystem(bot.DatabaseSubsystem, DatabaseSubsystemName, bot.Config)
				if err != nil {
					return err
				}
			}
			return StartSubsystem(bot.DatabaseSubsystem, &bot.SubsystemWG)
		} else {
			return StopSubsystem(bot.DatabaseSubsystem)
		}

	case NTPSubsystemName:
		if enable {
			if bot.NTPCheckerSubsystem == nil {
				err = InitSubsystem(bot.NTPCheckerSubsystem, NTPSubsystemName, bot.Config)
				if err != nil {
					return err
				}
			}
			return StartSubsystem(bot.NTPCheckerSubsystem, &bot.SubsystemWG)
		} else {
			return StopSubsystem(bot.NTPCheckerSubsystem)
		}

	case StrategySubsystemName:
		if enable {
			if bot.StrategySubsystem == nil {
				err = InitSubsystem(bot.StrategySubsystem, StrategySubsystemName, bot.Config)
				if err != nil {
					return err
				}
			}
			return StartSubsystem(bot.StrategySubsystem, &bot.SubsystemWG)
		} else {
			return StopSubsystem(bot.StrategySubsystem)
		}

	case InternetCheckerName:
		if enable {
			if bot.InternetSubsystem == nil {
				err = InitSubsystem(bot.InternetSubsystem, InternetCheckerName, bot.Config)
				if err != nil {
					return err
				}
			}
			return StartSubsystem(bot.InternetSubsystem, &bot.SubsystemWG)
		} else {
			return StopSubsystem(bot.InternetSubsystem)
		}

	}
	return fmt.Errorf("%s: %w", subSystemName, ErrSubsystemNotFound)
}
