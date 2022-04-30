package engine

import (
	"github.com/quantstop/quantstopterminal/internal/database/models"
	"github.com/quantstop/quantstopterminal/internal/log"
	"github.com/quantstop/quantstopterminal/internal/trader"
	"github.com/quantstop/quantstopterminal/pkg/exchange"
	"github.com/quantstop/quantstopterminal/pkg/exchange/coinbasepro"

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
	db, err := s.bot.GetSQL()
	if err != nil {
		return err
	}

	e := models.CryptoExchange{}
	err = e.GetCryptoExchangeByName(db, "coinbasepro")
	if err != nil {
		log.Error(log.TraderLogger, err)
		return err
	}

	//Create a client instance
	exchange.Coinbasepro, err = coinbasepro.NewSandboxClient(
		&coinbasepro.Auth{
			Key:        e.AuthKey,
			Passphrase: e.AuthPassphrase,
			Secret:     e.AuthSecret,
		},
	)
	if err != nil {
		log.Error(log.TraderLogger, err)
		return err
	}

	go trader.Run(db, s.bot.Webserver.Hub)

	return nil
}
