package database

import (
	"database/sql"
	"fmt"
	"github.com/quantstop/quantstopterminal/internal/database/models"
	"github.com/quantstop/quantstopterminal/internal/log"
	"time"
)

// SetConfig safely sets the global database instance's config
func (i *Instance) SetConfig(cfg *Config) error {
	if i == nil {
		return ErrNilInstance
	}
	if cfg == nil {
		return ErrNilConfig
	}
	i.m.Lock()
	i.config = cfg
	i.m.Unlock()
	return nil
}

// SetSQLiteConnection safely sets the global database instance's connection to use SQLite
func (i *Instance) SetSQLiteConnection(con *sql.DB) error {
	if i == nil {
		return ErrNilInstance
	}
	if con == nil {
		return errNilSQL
	}
	i.m.Lock()
	defer i.m.Unlock()
	i.SQL = con
	i.SQL.SetMaxOpenConns(1)
	return nil
}

// SetPostgresConnection safely sets the global database instance's connection to use Postgres
func (i *Instance) SetPostgresConnection(con *sql.DB) error {
	if i == nil {
		return ErrNilInstance
	}
	if con == nil {
		return errNilSQL
	}
	if err := con.Ping(); err != nil {
		return fmt.Errorf("%w %s", errFailedPing, err)
	}
	i.m.Lock()
	defer i.m.Unlock()
	i.SQL = con
	i.SQL.SetMaxOpenConns(2)
	i.SQL.SetMaxIdleConns(1)
	i.SQL.SetConnMaxLifetime(time.Hour)
	return nil
}

// SetMySQLConnection safely sets the global database instance's connection to use SQLite
func (i *Instance) SetMySQLConnection(con *sql.DB) error {
	if i == nil {
		return ErrNilInstance
	}
	if con == nil {
		return errNilSQL
	}
	i.m.Lock()
	defer i.m.Unlock()
	i.SQL = con
	i.SQL.SetMaxOpenConns(10)
	i.SQL.SetMaxIdleConns(10)
	i.SQL.SetConnMaxLifetime(time.Hour)
	return nil
}

// SetConnected safely sets the global database instance's connected status
func (i *Instance) SetConnected(v bool) {
	if i == nil {
		return
	}
	i.m.Lock()
	i.connected = v
	i.m.Unlock()
}

// CloseConnection safely disconnects the global database instance
func (i *Instance) CloseConnection() error {
	if i == nil {
		return ErrNilInstance
	}
	if i.SQL == nil {
		return errNilSQL
	}
	i.m.Lock()
	defer i.m.Unlock()

	return i.SQL.Close()
}

// IsConnected safely checks the SQL connection status
func (i *Instance) IsConnected() bool {
	if i == nil {
		return false
	}
	i.m.RLock()
	defer i.m.RUnlock()
	return i.connected
}

// GetConfig safely returns a copy of the config
func (i *Instance) GetConfig() *Config {
	if i == nil {
		return nil
	}
	i.m.RLock()
	defer i.m.RUnlock()
	cpy := i.config
	return cpy
}

// Ping pings the database
func (i *Instance) Ping() error {
	if i == nil {
		return ErrNilInstance
	}
	if !i.IsConnected() {
		return ErrDatabaseNotConnected
	}
	i.m.RLock()
	defer i.m.RUnlock()
	if i.SQL == nil {
		return errNilSQL
	}
	return i.SQL.Ping()
}

// GetSQL returns the sql connection
func (i *Instance) GetSQL() (*sql.DB, error) {
	if i == nil {
		return nil, ErrNilInstance
	}
	if i.SQL == nil {
		return nil, errNilSQL
	}
	i.m.Lock()
	defer i.m.Unlock()
	resp := i.SQL
	return resp, nil
}

// SeedDB will create the database tables if they do not exist, and create the default admin user.
func (i *Instance) SeedDB() error {

	// ToDo: this is only for sqlite ... can we move into individual drivers?

	log.Debugln(log.DatabaseLogger, "SeedDB - Checking for users table ...")
	row := i.SQL.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name='users' LIMIT 1")
	var table interface{}
	if err := row.Scan(&table); err != nil {
		log.Debugln(log.DatabaseLogger, "SeedDB - Checking for users table ... Not found.")
		log.Debugln(log.DatabaseLogger, "SeedDB - Creating users table ... ")
		usersTable := `
create table if not exists users
(
    id integer primary key autoincrement,
    username varchar(255) not null,
    password varchar(100) not null,
    constraint username
        unique (username)
);
`
		_, err := i.SQL.Exec(usersTable)
		if err != nil {
			log.Errorf(log.DatabaseLogger, "SeedDB - Creating users table ... Failed. Error: %v", err)
		}
		log.Debugln(log.DatabaseLogger, "SeedDB - Creating users table ... Success!")
		log.Debugln(log.DatabaseLogger, "SeedDB - Creating default admin ... ")
		defaultUser := models.User{
			Username: "admin",
			Password: "admin",
		}
		err = defaultUser.CreateUser(i.SQL)
		if err != nil {
			return err
		}
		log.Debugln(log.DatabaseLogger, "SeedDB - Creating default admin ... Success! Finished SeedDB.")
		return nil
	}

	log.Debugln(log.DatabaseLogger, "SeedDB - Checking for users table ... Found! Finished SeedDB.")
	return nil
}
