package models

import (
	"database/sql"
	"errors"
	"github.com/quantstop/quantstopterminal/internal/log"
	"github.com/quantstop/quantstopterminal/internal/webserver/utils"
)

type User struct {
	ID       uint32
	Username string
	Password string
	Salt     string
	Roles    []string
}

func CreateUsersTable(db *sql.DB, driver string) error {

	// todo: still only sqlite, dont like this too much as it is. could do a switch/case here with driver string parm ...

	log.Debugln(log.DatabaseLogger, "Checking for users table ...")
	row := db.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name='users' LIMIT 1")
	var table interface{}

	// returns err if no table is round
	if err := row.Scan(&table); err != nil {
		log.Debugln(log.DatabaseLogger, "Checking for users table ... Not found.")
		log.Debugln(log.DatabaseLogger, "Creating users table ... ")
		usersTable := `
create table if not exists users
(
    id integer primary key autoincrement,
    username varchar(255) not null,
    password varchar(100) not null,
    salt varchar(100) not null,
    constraint username
        unique (username)
);
`
		_, err := db.Exec(usersTable)
		if err != nil {
			log.Errorf(log.DatabaseLogger, "Creating users table ... Failed. Error: %v", err)
			return err // todo: custom error?
		}
		log.Debugln(log.DatabaseLogger, "Creating users table ... Success!")

		// check/create default admin
		if err := CreateDefaultAdmin(db); err != nil {
			log.Errorf(log.DatabaseLogger, "Error creating default admin: %v", err)
			return err
		}
	}

	log.Debugln(log.DatabaseLogger, "Checking for users table ... Found!")
	return nil
}

func CreateDefaultAdmin(db *sql.DB) error {

	// Check if default admin exists
	log.Debugln(log.DatabaseLogger, "Checking if default admin exists ...")
	if CheckDefaultAdminExists(db) {
		return nil
	}

	// Create default admin
	log.Debugln(log.DatabaseLogger, "Creating default admin ... ")
	defaultUser := User{
		Username: "admin",
		Password: "admin",
	}

	var err error

	// Set salt, and hash password with salt
	defaultUser.Salt = utils.GenerateRandomString(32)
	defaultUser.Password, err = utils.HashPassword(defaultUser.Password, defaultUser.Salt)
	if err != nil {
		return err
	}

	err = defaultUser.CreateUser(db)
	if err != nil {
		return err
	}
	log.Debugln(log.DatabaseLogger, "Creating default admin ... Success! Finished SeedDB.")

	return nil
}

func CheckDefaultAdminExists(db *sql.DB) bool {
	row := db.QueryRow("SELECT 1 FROM users WHERE id=$1 LIMIT 1", "1")
	u := &User{}
	if err := row.Scan(&u.ID, &u.Username, &u.Password, &u.Salt); err != nil {
		return false
	}
	return true
}

func (u *User) CreateUser(db *sql.DB) error {

	log.Debugln(log.DatabaseLogger, "Creating user ...")

	// the `Exec` method returns a `Result` type instead of a `Row`
	// we follow the same argument pattern to add query params
	result, err := db.Exec("INSERT INTO users (username, password, salt) VALUES ($1, $2, $3)", u.Username, u.Password, u.Salt)
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
	log.Debugln(log.DatabaseLogger, "User created. Inserted", rowsAffected, "rows")

	return nil
}

func (u *User) GetUserByUsername(db *sql.DB, username string) error {

	if username == "" {
		log.Errorf(log.DatabaseLogger, "username is nil")
		return errors.New("users model, cannot GetUserByUsername, username is nil")
	}

	if db == nil {
		log.Errorf(log.DatabaseLogger, "db is nil")
		return errors.New("users model, cannot GetUserByUsername, db is nil")
	}

	row := db.QueryRow("SELECT * FROM users WHERE username=$1 LIMIT 1", username)
	if err := row.Scan(&u.ID, &u.Username, &u.Password, &u.Salt); err != nil {
		log.Errorf(log.DatabaseLogger, "could not get user by username: %v", err)
		return err
	}

	return nil
}
