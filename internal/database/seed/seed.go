package seed

import (
	"database/sql"
	"github.com/quantstop/quantstopterminal/internal/database/models"
)

// SeedDB will create the database tables if they do not exist, and create the default admin user.
func SeedDB(db *sql.DB, driver string) error {

	// check/create users table - also creates default admin
	if err := models.CreateUsersTable(db, driver); err != nil {
		return err
	}

	// check/create roles table - also creates default roles
	if err := models.CreateRolesTable(db, driver); err != nil {
		return err
	}

	// check/create users_roles table (association table) - also creates default admin roles
	if err := models.CreateUsersRolesTable(db, driver); err != nil {
		return err
	}

	return nil

}
