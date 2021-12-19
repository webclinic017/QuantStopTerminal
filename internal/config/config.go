package config

import (
	"encoding/json"
	"fmt"
	"github.com/quantstop/quantstopterminal/internal/connectionmonitor"
	"github.com/quantstop/quantstopterminal/internal/database"
	"github.com/quantstop/quantstopterminal/internal/database/drivers"
	"github.com/quantstop/quantstopterminal/internal/ntpmonitor"
	"github.com/quantstop/quantstopterminal/internal/qstlog"
	"github.com/quantstop/quantstopterminal/pkg/system"
	"github.com/quantstop/quantstopterminal/pkg/system/convert"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"
)

const (
	DefaultNTPAllowedDifference         = 50000000
	DefaultNTPAllowedNegativeDifference = 50000000
	DefaultFileMode                     = os.FileMode(0755)
)

var (
	PreReleaseBlurb = "This version is pre-release and is not intended to be used as a production ready trading framework or bot - use at your own risk."
	Github          = "GitHub: https://github.com/QuantStop/QuantStopTerminal"
	Copyright       = fmt.Sprintf("Copyright (c) 2021-%d QuantStop.com", time.Now().Year())
	mutex           sync.Mutex
)

type Config struct {
	// Version types
	MajorVersion    string
	MinorVersion    string
	IsRelease       bool
	IsDevelopment   bool
	PrereleaseBlurb string
	GitHub          string
	Copyright       string

	// Config types
	ConfigDir       string
	WebServerPort   int
	GoMaxProcessors int
	Database        database.Config
	//Strategy 		strategy.Config
	NTP      ntpmonitor.Config
	Internet connectionmonitor.Config
	Logger   qstlog.Config
}

// DefaultFileMode controls the default permissions on any paths created by using MakePath.
//var DefaultFileMode = os.FileMode(0755)

func init() {
	findPaths()
}

// Refresh will rediscover the config paths, checking current environment
// variables again.
//
// This function is automatically called when the program initializes. If you
// change the environment variables at run-time, though, you may call the
// Refresh() function to reevaluate the config paths.
func Refresh() {
	findPaths()
}

// SystemConfig returns the system-wide configuration paths, with optional path
// components added to the end for vendor/application-specific settings.
func SystemConfig(folder ...string) []string {
	if len(folder) == 0 {
		return systemConfig
	}

	var paths []string
	for _, root := range systemConfig {
		p := append([]string{root}, filepath.Join(folder...))
		paths = append(paths, filepath.Join(p...))
	}

	return paths
}

// LocalConfig returns the local user configuration path, with optional
// path components added to the end for vendor/application-specific settings.
func LocalConfig(folder ...string) string {
	if len(folder) == 0 {
		return localConfig
	}

	return filepath.Join(localConfig, filepath.Join(folder...))
}

// LocalCache returns the local user cache folder, with optional path
// components added to the end for vendor/application-specific settings.
func LocalCache(folder ...string) string {
	if len(folder) == 0 {
		return localCache
	}

	return filepath.Join(localCache, filepath.Join(folder...))
}

// MakePath ensures that the full path you wanted, including vendor or
// application-specific components, exists. You can give this the output of
// any of the config path functions (SystemConfig, LocalConfig or LocalCache).
//
// In the event that the path function gives multiple answers, e.g. for
// SystemConfig, MakePath() will only attempt to create the sub-folders on
// the *first* path found. If this isn't what you want, you may want to just
// use the os.MkdirAll() functionality directly.
func makePath(paths ...string) error {
	if len(paths) >= 1 {
		err := os.MkdirAll(paths[0], DefaultFileMode)
		if err != nil {
			return err
		}
	}

	return nil
}

// SetupConfig will create the Config object and set the default data paths for the application.
func (c *Config) SetupConfig(MajorVersion string, MinorVersion string, IsRelease bool, IsDevelopment bool) error {

	// A common use case is to get a private config folder for your app to
	// place its settings files into, that are specific to the local user.
	configPath := LocalConfig("QuantStopTerminal")
	err := makePath(configPath) // Ensure it exists.
	if err != nil {
		return err
	}

	// Deal with a JSON configuration file in that folder.
	configFile := filepath.Join(configPath, "settings.json")

	// Does the file not exist?
	if _, err = os.Stat(configFile); os.IsNotExist(err) {

		// Setup default config
		c.ConfigDir = configPath
		c.WebServerPort = 8080
		c.GoMaxProcessors = -1
		c.Database = database.Config{
			Enabled: false,
			Verbose: false,
			Driver:  "mysql",
			ConnectionDetails: drivers.ConnectionDetails{
				Host:     "127.0.0.1",
				Port:     3306,
				Username: "docker",
				Password: "docker",
				Database: "docker",
				SSLMode:  "false",
			},
		}
		c.NTP = ntpmonitor.Config{
			Enabled: true,
			Verbose: false,
			Level:   0,
			Pool: []string{
				"pool.ntp.org:123",
			},
			AllowedDifference:         new(time.Duration),
			AllowedNegativeDifference: new(time.Duration),
		}
		c.Internet = connectionmonitor.Config{
			Enabled:          true,
			Initialized:      false,
			DNSList:          []string{"8.8.8.8", "8.8.4.4", "1.1.1.1", "1.0.0.1"},
			PublicDomainList: []string{"www.google.com", "www.cloudflare.com", "www.facebook.com"},
			CheckInterval:    time.Second * 3,
		}

		// Set default ntp settings
		*c.NTP.AllowedDifference = DefaultNTPAllowedDifference
		*c.NTP.AllowedNegativeDifference = DefaultNTPAllowedNegativeDifference

		// Load default logging config
		c.Logger = *qstlog.GenDefaultSettings()

		// Copy default logging config to global log config
		qstlog.RWM.Lock()
		qstlog.GlobalLogConfig = &c.Logger
		qstlog.RWM.Unlock()

		// Create the config file
		fh, err := os.Create(configFile)
		if err != nil {
			return err
		}
		defer func(fh *os.File) {
			_ = fh.Close()
		}(fh)

		// Write config to file in json format
		err = PrettyEncode(&c, fh)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		// Load the existing file.
		fh, err := os.Open(configFile)
		if err != nil {
			return err
		}
		defer func(fh *os.File) {
			_ = fh.Close()
		}(fh)

		decoder := json.NewDecoder(fh)
		err = decoder.Decode(&c)
		if err != nil {
			return err
		}
	}

	// Set flags
	c.MajorVersion = MajorVersion
	c.MinorVersion = MinorVersion
	c.IsRelease = IsRelease
	c.IsDevelopment = IsDevelopment
	c.PrereleaseBlurb = PreReleaseBlurb
	c.Copyright = Copyright
	c.GitHub = Github

	return nil
}

func PrettyEncode(data interface{}, out io.Writer) error {
	enc := json.NewEncoder(out)
	enc.SetIndent("", "    ")
	if err := enc.Encode(data); err != nil {
		return err
	}
	return nil
}

// CheckConfig will run private functions to verify the system config, and all subsystem configs are valid
func (c *Config) CheckConfig() error {
	err := c.checkLoggerConfig()
	if err != nil {
		return err
	}

	return nil
}

// CheckLoggerConfig checks to see logger values are present and valid in config
// if not, it creates a default instance of the logger
func (c *Config) checkLoggerConfig() error {
	mutex.Lock()
	defer mutex.Unlock()

	if c.Logger.Enabled == nil || c.Logger.Output == "" {
		c.Logger = *qstlog.GenDefaultSettings()
	}

	if c.Logger.AdvancedSettings.ShowLogSystemName == nil {
		c.Logger.AdvancedSettings.ShowLogSystemName = convert.BoolPtr(false)
	}

	if c.Logger.LoggerFileConfig != nil {
		if c.Logger.LoggerFileConfig.FileName == "" {
			c.Logger.LoggerFileConfig.FileName = "log.txt"
		}
		if c.Logger.LoggerFileConfig.Rotate == nil {
			c.Logger.LoggerFileConfig.Rotate = convert.BoolPtr(false)
		}
		if c.Logger.LoggerFileConfig.MaxSize <= 0 {
			qstlog.Warnf(qstlog.Global, "Logger rotation size invalid, defaulting to %v", qstlog.DefaultMaxFileSize)
			c.Logger.LoggerFileConfig.MaxSize = qstlog.DefaultMaxFileSize
		}
		qstlog.FileLoggingConfiguredCorrectly = true
	}
	qstlog.RWM.Lock()
	qstlog.GlobalLogConfig = &c.Logger
	qstlog.RWM.Unlock()

	logPath := c.GetDataPath("logs")
	err := system.CreateDir(logPath)
	if err != nil {
		return err
	}
	qstlog.LogPath = logPath

	return nil
}

// GetDataPath gets the data path for the given subpath
func (c *Config) GetDataPath(elem ...string) string {
	return filepath.Join(append([]string{c.ConfigDir}, elem...)...)
}

// GetVersion returns the version string
func (c *Config) GetVersion(short bool) string {
	versionStr := fmt.Sprintf("QuantstopTerminal v%s.%s %s %s",
		c.MajorVersion, c.MinorVersion, runtime.GOARCH, runtime.Version())

	if c.IsRelease {
		versionStr += " release.\n"
	} else {
		versionStr += " pre-release.\n"
		if !short {
			versionStr += c.PrereleaseBlurb + "\n"
		}
	}

	if c.IsDevelopment {
		versionStr += "Development mode: On\n"
	} else {
		versionStr += "Development mode: Off\n"
	}

	if short {
		return versionStr
	}
	versionStr += c.Copyright + "\n\n"
	versionStr += c.GitHub + "\n"
	//versionStr += c.Trello + "\n"
	//versionStr += c.Slack + "\n"
	//versionStr += c.Issues + "\n"
	return versionStr
}
