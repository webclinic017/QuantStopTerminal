package engine

import (
	"github.com/quantstop/quantstopterminal/internal/log"
	//"github.com/quantstop/quantstopterminal/internal/strategy"
	"sync"
)

type StrategySubsystem struct {
	Subsystem
}

func (s *StrategySubsystem) init(bot *Engine, name string) error {

	if err := s.Subsystem.init(bot, name); err != nil {
		return err
	}
	//s.enabled = config.Strategy.Enabled
	s.initialized = true
	log.Debugln(log.StrategyLogger, s.name+MsgSubsystemInitialized)
	return nil
}

func (s *StrategySubsystem) start(wg *sync.WaitGroup) (err error) {
	if err = s.Subsystem.start(wg); err != nil {
		return err
	}

	//strategy.RunStrats()
	s.started = true
	log.Debugln(log.StrategyLogger, s.name+MsgSubsystemStarted)
	return nil
}

func (s *StrategySubsystem) stop() error {

	if err := s.Subsystem.stop(); err != nil {
		return err
	}

	s.started = false
	log.Debugln(log.StrategyLogger, s.name+MsgSubsystemShutdown)
	return nil
}
