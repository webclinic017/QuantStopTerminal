package engine

import (
	"github.com/quantstop/quantstopterminal/internal/config"
	"github.com/quantstop/quantstopterminal/pkg/logger"
	//"github.com/quantstop/quantstopterminal/internal/strategy"
	"sync"
)

type StrategySubsystem struct {
	Subsystem
}

func (s *StrategySubsystem) init(config *config.Config, name string) error {

	if err := s.Subsystem.init(config, name); err != nil {
		return err
	}
	//s.enabled = config.Strategy.Enabled
	s.initialized = true
	logger.Debugln(logger.StrategyLogger, s.name+MsgSubsystemInitialized)
	return nil
}

func (s *StrategySubsystem) start(wg *sync.WaitGroup) (err error) {
	if err = s.Subsystem.start(wg); err != nil {
		return err
	}

	//strategy.RunStrats()
	s.started = true
	logger.Debugln(logger.StrategyLogger, s.name+MsgSubsystemStarted)
	return nil
}

func (s *StrategySubsystem) stop() error {

	if err := s.Subsystem.stop(); err != nil {
		return err
	}

	s.started = false
	logger.Debugln(logger.StrategyLogger, s.name+MsgSubsystemShutdown)
	return nil
}
