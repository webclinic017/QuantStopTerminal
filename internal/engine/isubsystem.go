package engine

import (
	"sync"
)

// iSubsystem is the main interface for all subsystems
type iSubsystem interface {
	init(bot *Engine, name string) error
	start(wg *sync.WaitGroup) error
	stop() error
	isRunning() bool
	isEnabled() bool
	isInitialized() bool
	getName() string
}
