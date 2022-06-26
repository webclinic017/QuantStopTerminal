package engine

import (
	"github.com/quantstop/quantstopterminal/internal/log"
	"sync"
)

type TraderSubsystem struct {
	Subsystem
}

func (s *TraderSubsystem) init(bot *Engine, name string) error {

	if err := s.Subsystem.init(bot, name, true); err != nil {
		return err
	}
	s.initialized = true
	log.Debugln(log.TraderLogger, s.name+MsgSubsystemInitialized)
	return nil
}

func (s *TraderSubsystem) start(wg *sync.WaitGroup) (err error) {
	if err = s.Subsystem.start(wg); err != nil {
		return err
	}

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

func (s *TraderSubsystem) run() error {
	/*db, err := s.bot.GetSQL()
	if err != nil {
		return err
	}

	e := models.CryptoExchange{}
	err = e.GetExchangeByName(db, "coinbasepro")
	if err != nil {
		log.Error(log.TraderLogger, err)
		return err
	}

	ex, err := exchange.GetExchange(base.CoinbasePro, base.NewAuth(
		e.AuthKey,
		e.AuthPassphrase,
		e.AuthSecret,
	),
	)
	log.Debugln(log.TraderLogger, ex.GetName())
	log.Debugln(log.TraderLogger, ex.Test())




	go trader.Run(db, s.bot.Webserver.Hub)*/

	return nil
}
