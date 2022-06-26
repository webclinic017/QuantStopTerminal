package engine

import (
	"errors"
	"fmt"
	"github.com/quantstop/quantstopterminal/internal/log"
	"sync"
)

const (

	// MsgSubsystemInitializing message to return when subsystem is initializing
	MsgSubsystemInitializing = " subsystem initializing..."

	// MsgSubsystemInitialized message to return when subsystem has initialized
	MsgSubsystemInitialized = " subsystem initializing... Success."

	// MsgSubsystemStarting message to return when subsystem is starting up
	MsgSubsystemStarting = " subsystem starting..."

	// MsgSubsystemStarted message to return when subsystem has started
	MsgSubsystemStarted = " subsystem starting... Success."

	// MsgSubsystemShuttingDown message to return when a subsystem is shutting down
	MsgSubsystemShuttingDown = " subsystem shutting down..."

	// MsgSubsystemShutdown message to return when a subsystem has shutdown
	MsgSubsystemShutdown = " subsystem shutting down ... Success"
)

var (
	// ErrSubsystemAlreadyStarted message to return when a subsystem is already started
	ErrSubsystemAlreadyStarted = errors.New("subsystem already started")

	// ErrSubsystemNotStarted message to return when subsystem not started
	ErrSubsystemNotStarted = errors.New("subsystem not started")

	// ErrNilSubsystem is returned when a subsystem hasn't had its Setup() func run
	ErrNilSubsystem = errors.New("subsystem not setup")

	// ErrSubsystemNotEnabled is returned when a subsystem can't be found
	ErrSubsystemNotEnabled = errors.New("subsystem not enabled")

	// ErrSubsystemNotInitialized is returned when a subsystem hasn't been initialized
	ErrSubsystemNotInitialized = errors.New("subsystem not initialized")

	// ErrSubsystemNotFound is returned when a subsystem can't be found
	ErrSubsystemNotFound = errors.New("subsystem not found")

	// ErrNilWaitGroup is returned when a subsystem has nil wait group
	ErrNilWaitGroup = errors.New("subsystem nil wait group received")

	// errNilEngine is returned when a subsystem has no bot to initialize with
	errNilEngine = errors.New("subsystem received nil engine")
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

// Subsystem The Subsystem struct is meant to be used as an abstract type
type Subsystem struct {
	name        string
	enabled     bool
	initialized bool
	started     bool
	shutdown    chan struct{}
	bot         *Engine
}

// init The function to initialize the Subsystem struct fields with values
func (sub *Subsystem) init(bot *Engine, name string, enabled bool) error {
	if sub == nil {
		return fmt.Errorf("%s subsystem %w", sub.name, ErrNilSubsystem)
	}
	if bot == nil {
		return fmt.Errorf("%s subsystem %w", sub.name, errNilEngine)
	}
	log.Debugln(log.SubsystemLogger, name+MsgSubsystemInitializing)
	sub.name = name
	sub.enabled = enabled
	sub.initialized = false
	sub.started = false
	sub.shutdown = make(chan struct{})
	sub.bot = bot
	return nil
}

// start The function to start the subsystem
func (sub *Subsystem) start(wg *sync.WaitGroup) error {
	if sub == nil {
		return fmt.Errorf("%s subsystem %w", sub.name, ErrNilSubsystem)
	}
	if sub.initialized == false {
		return fmt.Errorf("%s subsystem %w", sub.name, ErrSubsystemNotInitialized)
	}
	if wg == nil {
		return fmt.Errorf("%s subsystem %w", sub.name, ErrNilWaitGroup)
	}
	if sub.started == true {
		return fmt.Errorf("%s subsystem %w", sub.name, ErrSubsystemAlreadyStarted)
	}
	sub.started = false
	sub.shutdown = make(chan struct{})
	log.Debugln(log.SubsystemLogger, sub.name+MsgSubsystemStarting)
	return nil
}

// stop The function to stop the subsystem
func (sub *Subsystem) stop() error {
	if sub == nil {
		return fmt.Errorf("%s subsystem %w", sub.name, ErrNilSubsystem)
	}
	if sub.started == false {
		return fmt.Errorf("%s subsystem %w", sub.name, ErrSubsystemNotStarted)
	}
	sub.started = false
	log.Debugln(log.SubsystemLogger, sub.name+MsgSubsystemShuttingDown)

	return nil
}

// isRunning checks whether the subsystem is initialized and running
func (sub *Subsystem) isRunning() bool {
	if sub == nil {
		return false
	}
	return sub.started
}

// isEnabled checks whether the subsystem is enabled or not
func (sub *Subsystem) isEnabled() bool {
	return sub.enabled
}

// isInitialized checks whether the subsystem is initialized or not
func (sub *Subsystem) isInitialized() bool {
	return sub.initialized
}

// getName returns the subsystems name
func (sub *Subsystem) getName() string {
	return sub.name
}
