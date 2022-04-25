package engine

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/quantstop/quantstopterminal/internal/config"
	"github.com/quantstop/quantstopterminal/internal/log"
	"github.com/quantstop/quantstopterminal/internal/webserver"
	"github.com/quantstop/quantstopterminal/pkg/system"
	"github.com/quantstop/quantstopterminal/pkg/system/convert"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// Create creates a new instance of the engine
func Create(config *config.Config, version *Version) (*Engine, error) {

	engineMutex.Lock()
	defer engineMutex.Unlock()

	if config == nil {
		return nil, errors.New("engine: config is nil")
	}

	var bot Engine
	var err error

	// Set the bot config
	bot.Config = config

	// Set the bot version
	bot.Version = version

	// Set the max processors for go
	err = system.AdjustGoMaxProcs(bot.Config.GoMaxProcessors)
	if err != nil {
		return nil, fmt.Errorf("unable to adjust runtime GOMAXPROCS value. Err: %s", err)
	}

	return &bot, nil
}

// Initialize sets up the engine, creating the subsystems, and the subsystem registry.
func (bot *Engine) Initialize() error {

	if bot == nil {
		return errors.New("engine instance is nil")
	}

	engineMutex.Lock()
	defer engineMutex.Unlock()

	// Create new subsystem registry
	bot.SubsystemRegistry = NewServiceRegistry()

	// Initialize database subsystem
	if err := bot.initDatabaseSubsystem(); err != nil {
		return err
	}

	// Initialize ntp checker subsystem
	if err := bot.initNtpMonitorSubsystem(); err != nil {
		return err
	}

	// Initialize strategy subsystem
	if err := bot.initStrategySubsystem(); err != nil {
		return err
	}

	// Initialize internet checker subsystem
	if err := bot.initInternetMonitorSubsystem(); err != nil {
		return err
	}

	// Initialize sentiment analyzer subsystem
	if err := bot.initSentimentAnalyzerSubsystem(); err != nil {
		return err
	}

	return nil
}

func (bot *Engine) initDatabaseSubsystem() error {

	// Create and init database subsystem
	bot.DatabaseSubsystem = &DatabaseSubsystem{Subsystem: Subsystem{}}
	if err := bot.DatabaseSubsystem.init(bot, DatabaseSubsystemName); err != nil {
		log.Errorf(log.Global, "Database subsystem unable to initialize: %v", err)
		return err
	}

	// Register database subsystem
	if err := bot.SubsystemRegistry.RegisterService(bot.DatabaseSubsystem); err != nil {
		log.Errorf(log.Global, "Database subsystem unable to register: %v", err)
		return err
	}

	return nil
}

func (bot *Engine) initNtpMonitorSubsystem() error {
	if bot.Config.NTP.Enabled {

		// Create and init ntp checker subsystem
		bot.NTPCheckerSubsystem = &NTPCheckerSubsystem{Subsystem: Subsystem{}}
		if err := bot.NTPCheckerSubsystem.init(bot, NTPSubsystemName); err != nil {
			log.Errorf(log.Global, "NTP subsystem unable to initialize: %v", err)
			return err
		}

		// Register ntp checker subsystem
		if err := bot.SubsystemRegistry.RegisterService(bot.NTPCheckerSubsystem); err != nil {
			log.Errorf(log.Global, "NTP subsystem unable to register: %v", err)
			return err
		}

	}
	return nil
}

func (bot *Engine) initStrategySubsystem() error {
	/*if bot.Config.Strategy.Enabled {*/

	// Create and init strategy subsystem
	bot.TraderSubsystem = &TraderSubsystem{Subsystem: Subsystem{}}
	if err := bot.TraderSubsystem.init(bot, TraderSubsystemName); err != nil {
		log.Errorf(log.Global, "Trader subsystem unable to initialize: %v", err)
		return err
	}

	// Register strategy subsystem
	if err := bot.SubsystemRegistry.RegisterService(bot.TraderSubsystem); err != nil {
		log.Errorf(log.Global, "Trader subsystem unable to register: %v", err)
		return err
	}

	//}
	return nil
}

func (bot *Engine) initInternetMonitorSubsystem() error {
	if bot.Config.Internet.Enabled {

		// Create and init internet checker subsystem
		bot.InternetSubsystem = &ConnectionMonitor{Subsystem: Subsystem{}}
		if err := bot.InternetSubsystem.init(bot, InternetCheckerName); err != nil {
			log.Errorf(log.Global, "Internet checker subsystem unable to initialize: %v", err)
			return err
		}

		// Register internet checker subsystem
		if err := bot.SubsystemRegistry.RegisterService(bot.InternetSubsystem); err != nil {
			log.Errorf(log.Global, "Internet checker subsystem unable to register: %v", err)
			return err
		}

	}
	return nil
}

func (bot *Engine) initSentimentAnalyzerSubsystem() error {
	/*if bot.Config.Strategy.Enabled {*/

	// Create and init strategy subsystem
	bot.SentimentAnalyzer = &SentimentAnalyzerSubsystem{Subsystem: Subsystem{}}
	if err := bot.SentimentAnalyzer.init(bot, SentimentAnalyzerName); err != nil {
		log.Errorf(log.Global, "Sentiment Analyzer subsystem unable to initialize: %v", err)
		return err
	}

	// Register strategy subsystem
	if err := bot.SubsystemRegistry.RegisterService(bot.SentimentAnalyzer); err != nil {
		log.Errorf(log.Global, "Sentiment Analyzer subsystem unable to register: %v", err)
		return err
	}

	//}
	return nil
}

// Run starts the newly created instance of the engine
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

	// Everything good, create and run webserver
	// This is done here, because the webserver depends upon the instantiated bot and database connection
	var err error
	bot.Webserver, err = webserver.CreateWebserver(bot, bot.Config.Webserver, bot.IsDevelopment)
	if err != nil {
		return err
	}
	// run api server
	go func() {
		err = bot.Webserver.ListenAndServe(true, bot.Config.ConfigDir)
		if err != nil {
			err = fmt.Errorf("unexpected error from ListenAndServe: %w", err)
		}
	}()

	err = bot.TraderSubsystem.run()
	if err != nil {
		return err
	}

	// Print some info
	log.Infof(log.Global, "QuantstopTerminal started.\n")
	log.Infof(log.Global,
		"Using %d out of %d logical processors for runtime performance\n",
		runtime.GOMAXPROCS(-1), runtime.NumCPU())

	return nil
}

// Stop stops the running instance of the engine
func (bot *Engine) Stop() {

	engineMutex.Lock()
	defer engineMutex.Unlock()

	log.Debugln(log.Global, "Engine shutting down..")

	// Stop webserver
	bot.Webserver.Shutdown()

	// Stop all subsystems
	bot.SubsystemRegistry.StopAll()

	// Wait for subsystems to gracefully shutdown
	bot.SubsystemWG.Wait()
	if err := log.CloseLogger(); err != nil {
		fmt.Printf("Failed to close logger. Error: %v\n", err)
	}

}

// GetUptime returns the time since the bot last started
func (bot *Engine) GetUptime() string {
	//return time.Since(bot.Uptime).String()
	return convert.RoundDuration(time.Since(bot.Uptime), 2).String()
}

// GetSubsystemsStatus returns the status of all engine subsystems
func (bot *Engine) GetSubsystemsStatus() map[string]bool {

	status := make(map[string]bool)

	if bot.DatabaseSubsystem == nil {
		status[DatabaseSubsystemName] = false
	} else {
		status[DatabaseSubsystemName] = bot.DatabaseSubsystem.isRunning()
	}

	if bot.NTPCheckerSubsystem == nil {
		status[NTPSubsystemName] = false
	} else {
		status[NTPSubsystemName] = bot.NTPCheckerSubsystem.isRunning()
	}

	if bot.TraderSubsystem == nil {
		status[TraderSubsystemName] = false
	} else {
		status[TraderSubsystemName] = bot.TraderSubsystem.isRunning()
	}

	if bot.InternetSubsystem == nil {
		status[InternetCheckerName] = false
	} else {
		status[InternetCheckerName] = bot.InternetSubsystem.isRunning()
	}

	if bot.DatabaseSubsystem == nil {
		status[DatabaseSubsystemName] = false
	} else {
		status[DatabaseSubsystemName] = bot.DatabaseSubsystem.isRunning()
	}

	return status
}

// SetSubsystem enables or disables an engine subsystem
func (bot *Engine) SetSubsystem(subSystemName string, enable bool) error {
	if bot == nil {
		return errors.New("engine instance is nil")
	}

	if bot.Config == nil {
		return errNilEngine
	}

	engineMutex.Lock()
	defer engineMutex.Unlock()

	var err error
	switch strings.ToLower(subSystemName) {

	case NTPSubsystemName:
		if enable {
			if bot.NTPCheckerSubsystem == nil {
				err = bot.NTPCheckerSubsystem.init(bot, NTPSubsystemName)
				if err != nil {
					return err
				}
			}
			return bot.NTPCheckerSubsystem.start(&bot.SubsystemWG)
		} else {
			return bot.NTPCheckerSubsystem.stop()
		}

	case TraderSubsystemName:
		if enable {
			if bot.TraderSubsystem == nil {
				err = bot.TraderSubsystem.init(bot, TraderSubsystemName)
				if err != nil {
					return err
				}
			}
			return bot.TraderSubsystem.start(&bot.SubsystemWG)
		} else {
			return bot.TraderSubsystem.stop()
		}

	case InternetCheckerName:
		if enable {
			if bot.InternetSubsystem == nil {
				err = bot.InternetSubsystem.init(bot, InternetCheckerName)
				if err != nil {
					return err
				}
			}
			return bot.InternetSubsystem.start(&bot.SubsystemWG)
		} else {
			return bot.InternetSubsystem.stop()
		}

	}
	return fmt.Errorf("%s: %w", subSystemName, ErrSubsystemNotFound)
}

// GetVersion returns a map of the current version, along with other info
func (bot *Engine) GetVersion() map[string]string {
	version := make(map[string]string)

	version["version"] = bot.Version.Version
	version["copyright"] = bot.Version.Copyright
	version["prereleaseblurb"] = bot.Version.PrereleaseBlurb
	version["github"] = bot.Version.GitHub
	version["issues"] = bot.Version.Issues
	if bot.Version.IsDaemon {
		version["isdaemon"] = "true"
	} else {
		version["isdaemon"] = "false"
	}
	if bot.Version.IsRelease {
		version["isrelease"] = "true"
	} else {
		version["isrelease"] = "false"
	}
	if bot.Version.IsDevelopment {
		version["isdevelopment"] = "true"
	} else {
		version["isdevelopment"] = "false"
	}

	return version

}

// GetSQL returns a pointer to the database connection
func (bot *Engine) GetSQL() (*sql.DB, error) {

	if bot.DatabaseSubsystem.dbConn.SQL != nil {
		return bot.DatabaseSubsystem.dbConn.SQL, nil
	}
	log.Errorln(log.Global, "GetSQL, database is nil!")
	return nil, errors.New("engine cannot return nil database")
}

// SetSystemConfig saves system configuration data
func (bot *Engine) SetSystemConfig(apiUrl string, maxProcs string) error {
	intVar, err := strconv.Atoi(maxProcs)
	if err != nil {
		return err
	}
	bot.Config.GoMaxProcessors = intVar //todo: does this only take effect on restart?

	err = bot.Config.SaveConfig()
	if err != nil {
		return err
	}
	return nil
}
