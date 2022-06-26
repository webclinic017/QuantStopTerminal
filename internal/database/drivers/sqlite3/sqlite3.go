package sqlite

import (
	"database/sql"
	"github.com/quantstop/quantstopterminal/internal/database"
	//_ "github.com/mattn/go-sqlite3"
	_ "modernc.org/sqlite"
	"path/filepath"
)

// Connect opens a connection to sqlite database and returns a pointer to database.CoreDB
func Connect(name, db string) (*database.Instance, error) {
	if db == "" {
		return nil, database.ErrNoDatabaseProvided
	}

	databaseFullLocation := filepath.Join(database.CoreDB.DataPath, db)
	dbConn, err := sql.Open("sqlite", databaseFullLocation)
	if err != nil {
		return nil, err
	}

	switch name {
	case "core":
		err = database.CoreDB.SetSQLiteConnection(dbConn)
		if err != nil {
			return nil, err
		}
		return database.CoreDB, nil
	case "coinbase":
		err = database.CoinbaseDB.SetSQLiteConnection(dbConn)
		if err != nil {
			return nil, err
		}
		return database.CoinbaseDB, nil
	case "tdameritrade":
		err = database.TDAmeritradeDB.SetSQLiteConnection(dbConn)
		if err != nil {
			return nil, err
		}
		return database.TDAmeritradeDB, nil
	default:
		err = database.CoreDB.SetSQLiteConnection(dbConn)
		if err != nil {
			return nil, err
		}
		return database.CoreDB, nil
	}
}
