package main

import (
	_ "embed"
	"flag"
	"fmt"
	"github.com/quantstop/quantstopterminal/internal/config"
	"github.com/quantstop/quantstopterminal/internal/engine"
	qstlog "github.com/quantstop/quantstopterminal/internal/log"
	"github.com/quantstop/quantstopterminal/pkg/system"
	"log"
	"path/filepath"
	"strings"
)

var (
	BuildFlagVersion   string         // Build flag for version
	BuildFlagIsRelease string         // Build flag for setting release blurb
	Engine             *engine.Engine // Global pointer to Engine
	Config             *config.Config // Global pointer to Config
)

func main() {

	var err error

	// Create config
	Config = &config.Config{}

	// Setup config
	if err = Config.SetupConfig(); err != nil {
		log.Fatalf("Error settup up config: %s\n", err)
	}

	// Verify config
	if err = Config.CheckConfig(); err != nil {
		log.Fatalf("Error checking config: %s\n", err)
	}
	// Setup global logger
	if err = qstlog.SetupGlobalLogger(); err != nil {
		log.Fatalf("Error setting up global logger: %s\n", err)
	}

	// Setup all sub loggers
	if err = qstlog.SetupSubLoggers(Config.Logger.SubLoggers); err != nil {
		log.Fatalf("Error setting up subloggers: %s\n", err)
	}

	// Create default Version
	version := engine.CreateDefaultVersion()

	// Set build flags, unfortunately can only be of type string so must convert for IsRelease
	if BuildFlagIsRelease == "true" {
		version.IsRelease = true
	} else {
		version.IsRelease = false
	}
	version.Version = BuildFlagVersion

	// Parse runtime flags into Version
	flag.BoolVar(&version.IsDaemon, "daemon", false, "run as a background service")
	flag.BoolVar(&version.IsDevelopment, "development", false, "set development mode")
	flag.Parse()

	// Inject the website frontend into variable for webserver
	//webserver.Website = &WebFrontend

	// Print banner and version
	qstlog.Infof(qstlog.Global, "\n"+engine.GetRandomBanner()+"\n"+version.GetVersionString(false))

	// Print logger info
	qstlog.Debugln(qstlog.Global, "Logger initialized.")
	qstlog.Debugf(qstlog.Global, "Using config dir: %s\n", Config.ConfigDir)

	// Print full path of log file name
	if strings.Contains(Config.Logger.Output, "file") {
		qstlog.Debugf(qstlog.Global, "Using log file: %s\n",
			filepath.Join(qstlog.LogPath, Config.Logger.LoggerFileConfig.FileName))
	}

	// Create the bot
	if Engine, err = engine.Create(Config, version); err != nil {
		log.Fatalf("Unable to create bot engine. Error: %s\n", err)
	}

	// Initialize the bot
	if err = Engine.Initialize(); err != nil {
		log.Fatalf("Unable to initialize bot engine. Error: %s\n", err)
	}

	// Run the bot
	if err = Engine.Run(); err != nil {
		log.Fatalf("Unable to start bot engine. Error: %s\n", err)
	}

	// Wait for system interrupt to shut down the bot
	interrupt := system.WaitForInterrupt()
	s := fmt.Sprintf("Captured %v, shutdown requested.", interrupt)
	qstlog.Infoln(qstlog.Global, s)
	Engine.Stop()
	qstlog.Infoln(qstlog.Global, "Exiting.")

}
