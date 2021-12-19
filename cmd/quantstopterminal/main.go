package main

import (
	"fmt"
	"github.com/quantstop/quantstopterminal/internal/config"
	"github.com/quantstop/quantstopterminal/internal/engine"
	"github.com/quantstop/quantstopterminal/internal/qstlog"
	"github.com/quantstop/quantstopterminal/pkg/system"
	"log"
	"path/filepath"
	"strings"
)

var (
	Bot           *engine.Engine // Global pointer to Engine
	Config        *config.Config // Global pointer to Config
	MajorVersion  string         // Flag variable for MajorVersion
	MinorVersion  string         // Flag variable for MinorVersion
	IsRelease     string         // Flag variable for IsRelease
	IsDevelopment string         // Flag variable for IsDevelopment
)

func main() {

	var err error

	// Convert flag vars for IsRelease and IsDevelopment to booleans
	isRelease := false
	isDevelopment := false
	if IsRelease == "true" {
		isRelease = true
	}
	if IsDevelopment == "true" {
		isDevelopment = true
	}

	// Setup config
	Config = &config.Config{}
	if err = Config.SetupConfig(MajorVersion, MinorVersion, isRelease, isDevelopment); err != nil {
		log.Fatalf("Error settup up config: %s\n", err)
	}

	// Verify config
	if err = Config.CheckConfig(); err != nil {
		log.Fatalf("Error checking config: %s\n", err)
	}

	// Setup logger
	if err := qstlog.SetupGlobalLogger(); err != nil {
		log.Fatalf("Error setting up global logger: %s\n", err)
	}
	if err := qstlog.SetupSubLoggers(Config.Logger.SubLoggers); err != nil {
		log.Fatalf("Error setting up subloggers: %s\n", err)
	}

	// Print banner and version
	qstlog.Infof(qstlog.Global, "\n"+engine.Banner+"\n"+Config.GetVersion(false))

	// Print logger info
	qstlog.Debugln(qstlog.Global, "Logger initialized.")
	qstlog.Debugf(qstlog.Global, "Using config dir: %s\n", Config.ConfigDir)

	// Print full path of log file name
	if strings.Contains(Config.Logger.Output, "file") {
		qstlog.Debugf(qstlog.Global, "Using log file: %s\n",
			filepath.Join(qstlog.LogPath, Config.Logger.LoggerFileConfig.FileName))
	}

	// Create the bot
	if Bot, err = engine.Create(Config); err != nil {
		log.Fatalf("Unable to create bot engine. Error: %s\n", err)
	}

	// Initialize the bot
	if err = Bot.Initialize(); err != nil {
		log.Fatalf("Unable to initialize bot engine. Error: %s\n", err)
	}

	// Run the bot
	if err = Bot.Run(); err != nil {
		log.Fatalf("Unable to start bot engine. Error: %s\n", err)
	}

	// Wait for system interrupt to shut down the bot
	interrupt := system.WaitForInterrupt()
	s := fmt.Sprintf("Captured %v, shutdown requested.", interrupt)
	qstlog.Infoln(qstlog.Global, s)
	Bot.Stop()
	qstlog.Infoln(qstlog.Global, "Exiting.")
}
