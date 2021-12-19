package main

import (
	"fmt"
	"github.com/quantstop/quantstopterminal/internal/config"
	"github.com/quantstop/quantstopterminal/internal/engine"
	"github.com/quantstop/quantstopterminal/pkg/logger"
	"github.com/quantstop/quantstopterminal/pkg/system"
	"log"
	"path/filepath"
	"strings"
)

var (
	Bot           *engine.Engine
	Config        *config.Config
	MajorVersion  string
	MinorVersion  string
	IsRelease     string
	IsDevelopment string
)

func main() {

	var err error

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
	logger.Init()
	logger.SetupGlobalLogger()
	logger.SetupSubLoggers(Config.Logger.SubLoggers)

	// Print banner and version
	logger.Infof(logger.Global, "\n"+engine.Banner+"\n"+Config.GetVersion(false))

	// Print logger info
	logger.Debugln(logger.Global, "Logger initialized.")
	logger.Debugf(logger.Global, "Using config dir: %s\n", Config.ConfigDir)
	if strings.Contains(Config.Logger.Output, "file") {
		logger.Debugf(logger.Global, "Using log file: %s\n",
			filepath.Join(logger.LogPath, Config.Logger.LoggerFileConfig.FileName))
	}

	// Create the bot
	if Bot, err = engine.Create(Config); err != nil {
		logger.Fatalf(logger.Global, "Unable to create bot engine. Error: %s\n", err)
	}

	// Initialize the bot
	if err = Bot.Initialize(); err != nil {
		logger.Fatalf(logger.Global, "Unable to initialize bot engine. Error: %s\n", err)
	}

	// Run the bot
	if err = Bot.Run(); err != nil {
		logger.Fatalf(logger.Global, "Unable to start bot engine. Error: %s\n", err)
	}

	// Wait for system interrupt to shut down the bot
	interrupt := system.WaitForInterrupt()
	s := fmt.Sprintf("Captured %v, shutdown requested.", interrupt)
	logger.Infoln(logger.Global, s)
	//Bot.Stop()
	logger.Infoln(logger.Global, "Exiting.")
}
