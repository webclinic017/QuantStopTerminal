package database

import (
	"context"
	"database/sql"
	"errors"
	"github.com/quantstop/quantstopterminal/internal/database/drivers"
	"sync"
)

// Instance holds all information for a database instance
type Instance struct {
	SQL       *sql.DB
	DataPath  string
	config    *Config
	connected bool
	m         sync.RWMutex
}

// Config holds all configurable parameters for the database subsystem
type Config struct {
	Enabled bool
	Verbose bool
	Driver  string
	drivers.ConnectionDetails
}

var (
	// DB Global Database Connection
	DB = &Instance{}

	// MigrationDir which folder to look in for current migrations
	//MigrationDir = filepath.Join("..", "..", "database", "migrations")

	// ErrNoDatabaseProvided error to display when no database is provided
	ErrNoDatabaseProvided = errors.New("no database provided")

	// ErrDatabaseSupportDisabled error to display when no database is provided
	ErrDatabaseSupportDisabled = errors.New("database support is disabled")

	// SupportedDrivers slice of supported database driver types
	//SupportedDrivers = []string{DBSQLite, DBSQLite3, DBPostgreSQL}

	// ErrFailedToConnect for when a database fails to connect
	ErrFailedToConnect = errors.New("database failed to connect")

	// ErrDatabaseNotConnected for when a database is not connected
	ErrDatabaseNotConnected = errors.New("database is not connected")

	// DefaultSQLiteDatabase is the default sqlite3 database name to use
	//DefaultSQLiteDatabase = "qstrader.db"

	// ErrNilInstance for when a database is nil
	ErrNilInstance = errors.New("database instance is nil")

	// ErrNilConfig for when a config is nil
	ErrNilConfig  = errors.New("received nil config")
	errNilSQL     = errors.New("database SQL connection is nil")
	errFailedPing = errors.New("unable to verify database is connected, failed ping")
)

const (
	// DBSQLite const string for sqlite across code base
	DBSQLite = "sqlite"

	// DBSQLite3 const string for sqlite3 across code base
	DBSQLite3 = "sqlite3"

	// DBPostgreSQL const string for PostgreSQL across code base
	DBPostgreSQL = "postgres"

	// DBMySQL const string for MySQL across code base
	DBMySQL = "mysql"
)

// IDatabase allows for the passing of a database struct
// without giving the receiver access to all functionality
type IDatabase interface {
	IsConnected() bool
	GetSQL() (*sql.DB, error)
	GetConfig() *Config
}

// ISQL allows for the passing of an SQL connection
// without giving the receiver access to all functionality
type ISQL interface {
	BeginTx(context.Context, *sql.TxOptions) (*sql.Tx, error)
	Exec(string, ...interface{}) (sql.Result, error)
	Query(string, ...interface{}) (*sql.Rows, error)
	QueryRow(string, ...interface{}) *sql.Row
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}
