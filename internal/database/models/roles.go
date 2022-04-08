package models

import (
	"database/sql"
	"github.com/quantstop/quantstopterminal/internal/log"
)

type Role struct {
	ID   int
	Name string
}

var defaultRoles = []Role{
	{Name: "user"},
	{Name: "moderator"},
	{Name: "admin"},
}

func CreateRolesTable(db *sql.DB, driver string) error {

	// todo: still only sqlite, dont like this too much as it is. could do a switch/case here with driver string parm ...

	log.Debugln(log.DatabaseLogger, "Checking for roles table ...")
	row := db.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name='roles' LIMIT 1")
	var table interface{}

	// returns err if no table is round
	if err := row.Scan(&table); err != nil {
		log.Debugln(log.DatabaseLogger, "Checking for roles table ... Not found.")
		log.Debugln(log.DatabaseLogger, "Creating roles table ... ")
		usersTable := `
			create table if not exists roles
			(
				id integer primary key autoincrement,
				name varchar(255) not null,
				constraint name
					unique (name)
			);
		`
		_, err := db.Exec(usersTable)
		if err != nil {
			log.Errorf(log.DatabaseLogger, "Creating roles table ... Failed. Error: %v", err)
			return err // todo: custom error?
		}
		log.Debugln(log.DatabaseLogger, "Creating roles table ... Success!")

		if err = CreateDefaultRoles(db); err != nil {
			log.Errorf(log.DatabaseLogger, "Error creating default roles: %v", err)
			return err
		}
	}

	log.Debugln(log.DatabaseLogger, "Checking for roles table ... Found!")
	return nil
}

func CreateDefaultRoles(db *sql.DB) error {

	for _, role := range defaultRoles {
		if err := role.CreateRole(db); err != nil {
			return err
		}
	}

	return nil
}

func (r *Role) CreateRole(db *sql.DB) error {

	log.Debugln(log.DatabaseLogger, "Creating role ...")

	result, err := db.Exec("INSERT INTO roles (name) VALUES ($1)", r.Name)
	if err != nil {
		log.Errorf(log.DatabaseLogger, "could not insert row: %v", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Errorf(log.DatabaseLogger, "could not get affected rows: %v", err)
		return err
	}

	log.Debugln(log.DatabaseLogger, "Role created. Inserted", rowsAffected, "rows")

	return nil
}
