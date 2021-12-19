package sqlite

/*import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/quantstop/quantstopterminal/internal/database"
	"path/filepath"
)



// Connect opens a connection to sqlite database and returns a pointer to database.DB
func Connect(db string) (*database.Instance, error) {
	if db == "" {
		return nil, database.ErrNoDatabaseProvided
	}

	databaseFullLocation := filepath.Join(database.DB.DataPath, db)
	dbConn, err := sql.Open("sqlite3", databaseFullLocation)
	if err != nil {
		return nil, err
	}

	err = database.DB.SetSQLiteConnection(dbConn)
	if err != nil {
		return nil, err
	}

	return database.DB, nil
}
*/
