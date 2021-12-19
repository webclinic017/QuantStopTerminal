package engine

import (
	"github.com/quantstop/quantstopterminal/internal/config"
	"github.com/quantstop/quantstopterminal/internal/qstlog"
	"github.com/quantstop/quantstopterminal/internal/webserver"
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
	qstlog.Debugln(qstlog.Webserver, s.name+MsgSubsystemInitialized)
	return nil
}

func (s *WebserverSubsystem) start(wg *sync.WaitGroup) (err error) {
	if err = s.Subsystem.start(wg); err != nil {
		return err
	}
	webserver.StartHttpServer(s.config)
	//strategy.RunStrats()
	s.started = true
	qstlog.Debugln(qstlog.Webserver, s.name+MsgSubsystemStarted)
	return nil
}

func (s *WebserverSubsystem) stop() error {

	if err := s.Subsystem.stop(); err != nil {
		return err
	}

	s.started = false
	qstlog.Debugln(qstlog.Webserver, s.name+MsgSubsystemShutdown)
	return nil
}
