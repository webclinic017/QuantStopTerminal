package engine

import (
	"fmt"
	"github.com/quantstop/quantstopterminal/internal/connectionmonitor"
	"github.com/quantstop/quantstopterminal/internal/log"
	"sync"
)

type ConnectionMonitor struct {
	Subsystem
	conn *connectionmonitor.Checker
}

func (s *ConnectionMonitor) init(bot *Engine, name string) error {

	if err := s.Subsystem.init(bot, name); err != nil {
		return err
	}

	if s.bot.Config.Internet.DNSList == nil {
		s.bot.Config.Internet.DNSList = connectionmonitor.DefaultDNSList
	}
	if s.bot.Config.Internet.PublicDomainList == nil {
		s.bot.Config.Internet.PublicDomainList = connectionmonitor.DefaultDomainList
	}
	if s.bot.Config.Internet.CheckInterval == 0 {
		s.bot.Config.Internet.CheckInterval = connectionmonitor.DefaultCheckInterval
	}
	s.enabled = bot.Config.Internet.Enabled
	s.initialized = true
	log.Debugln(log.ConnMonitor, s.name+MsgSubsystemInitialized)
	return nil
}

func (s *ConnectionMonitor) start(wg *sync.WaitGroup) (err error) {

	if err = s.Subsystem.start(wg); err != nil {
		return err
	}

	s.conn, err = connectionmonitor.New(s.bot.Config.Internet.DNSList,
		s.bot.Config.Internet.PublicDomainList,
		s.bot.Config.Internet.CheckInterval)
	if err != nil {
		s.started = false
		return err
	}

	s.started = true
	log.Debugln(log.ConnMonitor, s.name+MsgSubsystemStarted)
	return nil
}

func (s *ConnectionMonitor) stop() error {
	if err := s.Subsystem.stop(); err != nil {
		return err
	}

	if s.conn == nil {
		return fmt.Errorf("connection manager: %w", ErrNilSubsystem)
	}
	s.conn.Shutdown()

	s.started = false
	log.Debugln(log.ConnMonitor, s.name+MsgSubsystemShutdown)
	return nil
}
