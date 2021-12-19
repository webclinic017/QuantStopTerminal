package engine

import (
	"github.com/quantstop/quantstopterminal/internal/config"
	"github.com/quantstop/quantstopterminal/internal/connectionmonitor"
	"github.com/quantstop/quantstopterminal/internal/qstlog"
	"sync"
)

type ConnectionMonitor struct {
	Subsystem
	conn *connectionmonitor.Checker
}

func (s *ConnectionMonitor) init(config *config.Config, name string) error {

	if err := s.Subsystem.init(config, name); err != nil {
		return err
	}

	if s.config.Internet.DNSList == nil {
		s.config.Internet.DNSList = connectionmonitor.DefaultDNSList
	}
	if s.config.Internet.PublicDomainList == nil {
		s.config.Internet.PublicDomainList = connectionmonitor.DefaultDomainList
	}
	if s.config.Internet.CheckInterval == 0 {
		s.config.Internet.CheckInterval = connectionmonitor.DefaultCheckInterval
	}
	s.enabled = config.Internet.Enabled
	s.initialized = true
	qstlog.Debugln(qstlog.ConnMonitor, s.name+MsgSubsystemInitialized)
	return nil
}

func (s *ConnectionMonitor) start(wg *sync.WaitGroup) (err error) {

	if err = s.Subsystem.start(wg); err != nil {
		return err
	}

	s.conn, err = connectionmonitor.New(s.config.Internet.DNSList,
		s.config.Internet.PublicDomainList,
		s.config.Internet.CheckInterval)
	if err != nil {
		s.started = false
		return err
	}

	s.started = true
	qstlog.Debugln(qstlog.ConnMonitor, s.name+MsgSubsystemStarted)
	return nil
}

func (s *ConnectionMonitor) stop() error {
	if err := s.Subsystem.stop(); err != nil {
		return err
	}

	if s.conn == nil {
		//return fmt.Errorf("connection manager: %w", errConnectionCheckerIsNil)
	}
	s.conn.Shutdown()

	s.started = false
	qstlog.Debugln(qstlog.ConnMonitor, s.name+MsgSubsystemShutdown)
	return nil
}
