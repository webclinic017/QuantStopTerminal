package mysql

import (
	"database/sql"
	"fmt"
	"github.com/quantstop/quantstopterminal/internal/database"
	"github.com/quantstop/quantstopterminal/internal/database/drivers"

	_ "github.com/go-sql-driver/mysql"
)

// Connect opens a connection to MySQL database and returns a pointer to database.CoreDB
func Connect(name string, cfg *drivers.ConnectionDetails) (*database.Instance, error) {
	if cfg == nil {
		return nil, database.ErrNilConfig
	}

	if cfg.SSLMode == "" {
		cfg.SSLMode = "disable"
	}

	configDSN := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
		"charset=utf8mb4&parseTime=True&loc=Local")

	db, err := sql.Open(database.DBMySQL, configDSN)
	if err != nil {
		return nil, err
	}

	switch name {
	case "core":
		err = database.CoreDB.SetMySQLConnection(db)
		if err != nil {
			return nil, err
		}
		return database.CoreDB, nil
	case "coinbase":
		err = database.CoinbaseDB.SetMySQLConnection(db)
		if err != nil {
			return nil, err
		}
		return database.CoinbaseDB, nil
	case "tdameritrade":
		err = database.TDAmeritradeDB.SetMySQLConnection(db)
		if err != nil {
			return nil, err
		}
		return database.TDAmeritradeDB, nil
	default:
		err = database.CoreDB.SetMySQLConnection(db)
		if err != nil {
			return nil, err
		}
		return database.CoreDB, nil
	}

}
