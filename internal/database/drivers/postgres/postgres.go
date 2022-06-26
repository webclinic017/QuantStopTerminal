package postgres

import (
	"database/sql"
	"fmt"
	"github.com/quantstop/quantstopterminal/internal/database"
	"github.com/quantstop/quantstopterminal/internal/database/drivers"

	// import go libpq driver package
	_ "github.com/lib/pq"
)

// Connect opens a connection to Postgres database and returns a pointer to database.CoreDB
func Connect(name string, cfg *drivers.ConnectionDetails) (*database.Instance, error) {
	if cfg == nil {
		return nil, database.ErrNilConfig
	}

	if cfg.SSLMode == "" {
		cfg.SSLMode = "disable"
	}

	configDSN := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
		cfg.SSLMode)

	db, err := sql.Open(database.DBPostgreSQL, configDSN)
	if err != nil {
		return nil, err
	}
	switch name {
	case "core":
		err = database.CoreDB.SetPostgresConnection(db)
		if err != nil {
			return nil, err
		}
		return database.CoreDB, nil
	case "coinbase":
		err = database.CoinbaseDB.SetPostgresConnection(db)
		if err != nil {
			return nil, err
		}
		return database.CoinbaseDB, nil
	case "tdameritrade":
		err = database.TDAmeritradeDB.SetPostgresConnection(db)
		if err != nil {
			return nil, err
		}
		return database.TDAmeritradeDB, nil
	default:
		err = database.CoreDB.SetPostgresConnection(db)
		if err != nil {
			return nil, err
		}
		return database.CoreDB, nil
	}
}
