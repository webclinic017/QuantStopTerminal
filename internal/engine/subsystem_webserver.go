package engine

import (
	"fmt"
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

	s.server, err = webserver.CreateWebserver(bot, bot.Config.Webserver, s.bot.IsDevelopment)
	if err != nil {
		log.Errorf(log.Webserver, "Error creating webserver: %v", err)
		return err
	}

	//s.server.SetupRoutes(s.bot.IsDevelopment)

	s.enabled = bot.Config.Webserver.Enabled
	s.initialized = true
	log.Debugln(log.Webserver, s.name+MsgSubsystemInitialized)
	return nil
}

func (s *WebserverSubsystem) start(wg *sync.WaitGroup) (err error) {
	if err = s.Subsystem.start(wg); err != nil {
		return err
	}

	s.started = true
	log.Debugln(log.Webserver, s.name+MsgSubsystemStarted)
	wg.Add(1)
	s.wg.Add(1)
	go s.run(wg)

	// if dev mode, run node server
	/*if s.bot.IsDevelopment {
		go s.server.StartNodeDevelopmentServer()
	}*/

	return nil
}

func (s *WebserverSubsystem) stop() error {

	if err := s.Subsystem.stop(); err != nil {
		return err
	}
	s.started = false
	s.server.Shutdown()
	close(s.shutdown)
	s.wg.Wait()
	log.Debugln(log.Webserver, s.name+MsgSubsystemShutdown)
	return nil
}

// run this is the main loop for the subsystem
func (s *WebserverSubsystem) run(wg *sync.WaitGroup) {

	defer func() {
		s.wg.Done()
		wg.Done()
		log.Debugln(log.Webserver, "Webserver subsystem shutdown.")
	}()

	// run api server
	err := s.server.ListenAndServe(true, s.bot.Config.ConfigDir)
	if err != nil {
		err = fmt.Errorf("unexpected error from ListenAndServe: %w", err)
	}

}
