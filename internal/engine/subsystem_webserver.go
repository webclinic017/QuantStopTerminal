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

	s.server, err = webserver.CreateWebserver(bot, bot.Config.Webserver, false)
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

	//wg.Add(1)
	//s.wg.Add(1)
	//go s.run()
	go s.server.StartWebServer(s.bot.IsDevelopment, s.shutdown)

	s.started = true
	log.Debugln(log.Webserver, s.name+MsgSubsystemStarted)
	return nil
}

func (s *WebserverSubsystem) stop() error {

	if err := s.Subsystem.stop(); err != nil {
		return err
	}
	close(s.shutdown)
	s.wg.Wait()
	s.started = false
	log.Debugln(log.Webserver, s.name+MsgSubsystemShutdown)
	return nil
}

// run this is the main loop for the subsystem
func (s *WebserverSubsystem) run() {

	//s.server.StartWebServer()
	// Start the Node client app (only for version "development")
	/*if s.bot.Config.IsDevelopment {
		go webserver.StartNodeDevelopmentServer()
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer func() {
		cancel()
		qstlog.Debugln(qstlog.Webserver, "Webserver subsystem shutdown.")
	}()

	_ = s.server.HttpServer.ListenAndServe()
	qstlog.Infof(qstlog.Webserver, "Web server started, listening on http://%v\n", ":8080")



	// Code below will only run if we get a message on the shutdown channel
	// Remember, we are inside a goroutine right now
	<-s.shutdown
	_ = s.server.HttpServer.Shutdown(ctx)*/
}
