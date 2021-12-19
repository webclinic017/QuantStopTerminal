package engine

import (
	"github.com/quantstop/quantstopterminal/internal/config"
	"github.com/quantstop/quantstopterminal/internal/webserver"
	"github.com/quantstop/quantstopterminal/pkg/logger"
	"sync"
)

type WebserverSubsystem struct {
	Subsystem
}

func (s *WebserverSubsystem) init(config *config.Config, name string) error {

	if err := s.Subsystem.init(config, name); err != nil {
		return err
	}
	s.initialized = true
	logger.Debugln(logger.WebserverLogger, s.name+MsgSubsystemInitialized)
	return nil
}

func (s *WebserverSubsystem) start(wg *sync.WaitGroup) (err error) {
	if err = s.Subsystem.start(wg); err != nil {
		return err
	}
	webserver.StartHttpServer(s.config)
	//strategy.RunStrats()
	s.started = true
	logger.Debugln(logger.WebserverLogger, s.name+MsgSubsystemStarted)
	return nil
}

func (s *WebserverSubsystem) stop() error {

	if err := s.Subsystem.stop(); err != nil {
		return err
	}

	s.started = false
	logger.Debugln(logger.WebserverLogger, s.name+MsgSubsystemShutdown)
	return nil
}
