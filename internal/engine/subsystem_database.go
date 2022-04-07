package engine

import (
	"fmt"
	"github.com/quantstop/quantstopterminal/internal/database"
	"github.com/quantstop/quantstopterminal/internal/database/drivers/mysql"
	pgsql "github.com/quantstop/quantstopterminal/internal/database/drivers/postgres"
	sqlite "github.com/quantstop/quantstopterminal/internal/database/drivers/sqlite3"
	"github.com/quantstop/quantstopterminal/internal/database/seed"
	"github.com/quantstop/quantstopterminal/internal/log"
	"sync"
	"time"
)

type DatabaseSubsystem struct {
	Subsystem
	wg     sync.WaitGroup
	dbConn *database.Instance
}

// init sets config and params
func (s *DatabaseSubsystem) init(bot *Engine, name string) error {
	if err := s.Subsystem.init(bot, name); err != nil {
		return err
	}
	s.enabled = bot.Config.Database.Enabled
	s.dbConn = database.DB
	if err := s.dbConn.SetConfig(&s.bot.Config.Database); err != nil {
		return err
	}
	s.initialized = true
	log.Debugln(log.DatabaseLogger, s.name+MsgSubsystemInitialized)
	return nil
}

// start sets up the database subsystem to maintain an SQL connection
func (s *DatabaseSubsystem) start(wg *sync.WaitGroup) (err error) {
	if err = s.Subsystem.start(wg); err != nil {
		return err
	}

	if s.bot.Config.Database.Enabled {
		switch s.bot.Config.Database.Driver {
		case database.DBPostgreSQL:
			log.Debugf(log.DatabaseLogger,
				"Database subsystem attempting to establish database connection to host %s/%s utilising %s driver\n",
				s.bot.Config.Database.Host,
				s.bot.Config.Database.Database,
				s.bot.Config.Database.Driver)
			s.dbConn, err = pgsql.Connect(&s.bot.Config.Database)
		case database.DBSQLite, database.DBSQLite3:
			log.Debugf(log.DatabaseLogger,
				"Database subsystem attempting to establish database connection to %s utilising %s driver\n",
				s.bot.Config.Database.Database,
				s.bot.Config.Database.Driver)
			s.dbConn, err = sqlite.Connect(s.bot.Config.Database.Database)
		case database.DBMySQL:
			log.Debugf(log.DatabaseLogger,
				"Database subsystem attempting to establish database connection to host %s/%s utilising %s driver\n",
				s.bot.Config.Database.Host,
				s.bot.Config.Database.Database,
				s.bot.Config.Database.Driver)
			s.dbConn, err = mysql.Connect(&s.bot.Config.Database)
		default:
			return database.ErrNoDatabaseProvided
		}
		if err != nil {
			return fmt.Errorf("%w: %v Some features that utilise a database will be unavailable", database.ErrFailedToConnect, err)
		}
		log.Debugln(log.DatabaseLogger, s.name+MsgSubsystemStarted)
		s.started = true
		s.dbConn.SetConnected(true)
		err = seed.SeedDB(s.dbConn.SQL, s.bot.Config.Database.Driver)
		if err != nil {
			return err
		}
		wg.Add(1)
		s.wg.Add(1)
		go s.run(wg)
		return nil
	}

	return database.ErrDatabaseSupportDisabled
}

// stop attempts to shut down the subsystem
func (s *DatabaseSubsystem) stop() error {
	if err := s.Subsystem.stop(); err != nil {
		return err
	}

	s.started = false
	err := s.dbConn.CloseConnection()
	if err != nil {
		log.Errorf(log.DatabaseLogger, "Failed to close database: %v", err)
	}

	close(s.shutdown)
	s.wg.Wait()
	log.Debugln(log.DatabaseLogger, s.name+MsgSubsystemShutdown)
	return nil
}

// run this is the main loop for the subsystem
func (s *DatabaseSubsystem) run(wg *sync.WaitGroup) {

	t := time.NewTicker(time.Second * 5)

	defer func() {
		t.Stop()
		s.wg.Done()
		wg.Done()
		log.Debugln(log.DatabaseLogger, "Database subsystem shutdown.")
	}()

	// This lets the goroutine wait on communication from the channel
	// Docs: https://tour.golang.org/concurrency/5
	for {
		select {
		case <-s.shutdown: // if channel message is shutdown finish loop
			return
		case <-t.C: // on channel tick check the connection
			err := s.CheckConnection()
			if err != nil {
				log.Error(log.DatabaseLogger, "Database connection error:", err)
			}
		}
	}
}

// GetInstance returns a limited scoped database instance
func (s *DatabaseSubsystem) GetInstance() database.IDatabase {
	if s == nil || !s.started {
		return nil
	}
	return s.dbConn
}

// CheckConnection checks to make sure the database is connected
func (s *DatabaseSubsystem) CheckConnection() error {
	if s == nil {
		return fmt.Errorf("%s %w", "DatabaseSubsystem", ErrNilSubsystem)
	}
	if s.started == false {
		return fmt.Errorf("%s %w", "DatabaseSubsystem", ErrSubsystemNotStarted)
	}
	if !s.bot.Config.Database.Enabled {
		return database.ErrDatabaseSupportDisabled
	}
	if s.dbConn == nil {
		return database.ErrNoDatabaseProvided
	}

	if err := s.dbConn.Ping(); err != nil {
		s.dbConn.SetConnected(false)
		return err
	}

	if !s.dbConn.IsConnected() {
		log.Info(log.DatabaseLogger, "Database connection reestablished")
		s.dbConn.SetConnected(true)
	}
	return nil
}
