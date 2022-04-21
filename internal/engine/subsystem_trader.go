package engine

import (
	"github.com/quantstop/quantstopterminal/internal/log"
	"github.com/quantstop/quantstopterminal/internal/trader"

	//"github.com/quantstop/quantstopterminal/internal/strategy"
	"sync"
)

type TraderSubsystem struct {
	Subsystem
}

func (s *TraderSubsystem) init(bot *Engine, name string) error {

	if err := s.Subsystem.init(bot, name); err != nil {
		return err
	}
	/*s.enabled = bot.Config*/
	s.enabled = true
	s.initialized = true
	log.Debugln(log.TraderLogger, s.name+MsgSubsystemInitialized)
	return nil
}

func (s *TraderSubsystem) start(wg *sync.WaitGroup) (err error) {
	if err = s.Subsystem.start(wg); err != nil {
		return err
	}

	go trader.Run()
	s.started = true
	log.Debugln(log.TraderLogger, s.name+MsgSubsystemStarted)
	return nil
}

func (s *TraderSubsystem) stop() error {

	if err := s.Subsystem.stop(); err != nil {
		return err
	}

	s.started = false
	log.Debugln(log.TraderLogger, s.name+MsgSubsystemShutdown)
	return nil
}
