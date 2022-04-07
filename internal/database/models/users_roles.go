package models

import (
	"database/sql"
	"github.com/quantstop/quantstopterminal/internal/log"
)

type UserRole struct {
	UserID int
	RoleID int
}

func CreateUsersRolesTable(db *sql.DB, driver string) error {

	// todo: still only sqlite, dont like this too much as it is. could do a switch/case here with driver string parm ...

	log.Debugln(log.DatabaseLogger, "Checking for users_roles table ...")
	row := db.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name='users_roles' LIMIT 1")
	var table interface{}

	// returns err if no table is round
	if err := row.Scan(&table); err != nil {
		log.Debugln(log.DatabaseLogger, "Checking for users_roles table ... Not found.")
		log.Debugln(log.DatabaseLogger, "Creating users_roles table ... ")
		usersTable := `
create table if not exists users_roles
(
    user_id integer,
    role_id integer,
	foreign key (user_id) references users(id),
	foreign key (role_id) references roles(id)
);
`
		_, err := db.Exec(usersTable)
		if err != nil {
			log.Errorf(log.DatabaseLogger, "Creating users_roles table ... Failed. Error: %v", err)
			return err // todo: custom error?
		}
		log.Debugln(log.DatabaseLogger, "Creating users_roles table ... Success!")

		if err = CreateDefaultAdminRoles(db); err != nil {
			log.Errorf(log.DatabaseLogger, "Error creating default admin roles: %v", err)
			return err
		}

	}

	log.Debugln(log.DatabaseLogger, "Checking for users_roles table ... Found!")
	return nil
}

func CreateDefaultAdminRoles(db *sql.DB) error {

	for i := range defaultRoles {
		ur := UserRole{RoleID: i + 1, UserID: 1}
		if err := ur.CreateUserRole(db); err != nil {
			return err
		}
	}

	return nil
}

func (ur *UserRole) CreateUserRole(db *sql.DB) error {
	log.Debugln(log.DatabaseLogger, "Creating user role association ...")

	// the `Exec` method returns a `Result` type instead of a `Row`
	// we follow the same argument pattern to add query params
	result, err := db.Exec("INSERT INTO users_roles (user_id, role_id) VALUES ($1, $2)", ur.UserID, ur.RoleID)
	if err != nil {
		log.Errorf(log.DatabaseLogger, "could not insert row: %v", err)
		return err
	}

	// the `Result` type has special methods like `RowsAffected` which returns the
	// total number of affected rows reported by the database
	// In this case, it will tell us the number of rows that were inserted using
	// the above query
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Errorf(log.DatabaseLogger, "could not get affected rows: %v", err)
		return err
	}

	// we can log how many rows were inserted
	log.Debugln(log.DatabaseLogger, "User role association created. Inserted", rowsAffected, "rows")

	return nil
}
