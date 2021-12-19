package engine

import (
	"github.com/quantstop/quantstopterminal/internal/config"
	"sync"
)

// iSubsystem is the main interface for all subsystems
type iSubsystem interface {
	init(config *config.Config, name string) error
	start(wg *sync.WaitGroup) error
	stop() error
	isRunning() bool
	isEnabled() bool
	isInitialized() bool
	getName() string
}

func InitSubsystem(i iSubsystem, name string, config *config.Config) error {
	if err := i.init(config, name); err != nil {
		return err
	}
	return nil
}

func StartSubsystem(i iSubsystem, wg *sync.WaitGroup) error {
	if err := i.start(wg); err != nil {
		return err
	}
	return nil
}

func StopSubsystem(i iSubsystem) error {
	err := i.stop()
	if err != nil {
		return err
	}
	return nil
}
