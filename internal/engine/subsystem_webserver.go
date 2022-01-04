package engine

import (
	"github.com/quantstop/quantstopterminal/internal/log"
	"github.com/quantstop/quantstopterminal/internal/webserver"
	"sync"
)

type WebserverSubsystem struct {
	Subsystem
	wg     sync.WaitGroup
	server *webserver.Webserver
}

func (s *WebserverSubsystem) init(bot *Engine, name string) error {

	var err error

	if err = s.Subsystem.init(bot, name); err != nil {
		return err
	}

	s.server, err = webserver.CreateWebserver(bot, bot.Config.Webserver)
	if err != nil {
		log.Errorf(log.Global, "Error creating webserver: %v", err)
	}

	s.enabled = bot.Config.Webserver.Enabled
	s.initialized = true
	log.Debugln(log.Webserver, s.name+MsgSubsystemInitialized)
	return nil
}

func (s *WebserverSubsystem) start(wg *sync.WaitGroup) (err error) {
	if err = s.Subsystem.start(wg); err != nil {
		return err
	}

	s.server.SetupRoutes(s.bot.IsDevelopment)

	wg.Add(1)
	s.wg.Add(1)
	//go s.run()
	go s.server.StartWebServer(true, s.bot.IsDevelopment, s.shutdown, s.bot.Config.ConfigDir)

	s.started = true
	log.Debugln(log.Webserver, s.name+MsgSubsystemStarted)
	return nil
}

func (s *WebserverSubsystem) stop() error {

	if err := s.Subsystem.stop(); err != nil {
		return err
	}
	s.started = false
	close(s.shutdown)
	s.wg.Wait()
	log.Debugln(log.Webserver, s.name+MsgSubsystemShutdown)
	return nil
}
