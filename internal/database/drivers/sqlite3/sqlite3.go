package sqlite

import (
	"database/sql"
	"github.com/quantstop/quantstopterminal/internal/database"
	//_ "github.com/mattn/go-sqlite3"
	_ "modernc.org/sqlite"
	"path/filepath"
)

// Connect opens a connection to sqlite database and returns a pointer to database.DB
func Connect(db string) (*database.Instance, error) {
	if db == "" {
		return nil, database.ErrNoDatabaseProvided
	}

	databaseFullLocation := filepath.Join(database.DB.DataPath, db)
	dbConn, err := sql.Open("sqlite", databaseFullLocation)
	if err != nil {
		return nil, err
	}

	err = database.DB.SetSQLiteConnection(dbConn)
	if err != nil {
		return nil, err
	}

	return database.DB, nil
}
