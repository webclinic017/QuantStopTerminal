package mysql

import (
	"database/sql"
	"fmt"
	"github.com/quantstop/quantstopterminal/internal/database"

	_ "github.com/go-sql-driver/mysql"
)

// Connect opens a connection to MySQL database and returns a pointer to database.DB
func Connect(cfg *database.Config) (*database.Instance, error) {
	if cfg == nil {
		return nil, database.ErrNilConfig
	}
	if !cfg.Enabled {
		return nil, database.ErrDatabaseSupportDisabled
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
	err = database.DB.SetMySQLConnection(db)
	if err != nil {
		return nil, err
	}
	return database.DB, nil
}
