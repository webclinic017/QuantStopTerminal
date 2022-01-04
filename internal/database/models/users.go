package models

import (
	"database/sql"
	"github.com/quantstop/quantstopterminal/internal/log"
)

type User struct {
	ID       uint32
	Username string
	Password string
}

func (u *User) CreateUser(db *sql.DB) error {

	log.Debugln(log.DatabaseLogger, "Creating user ...")

	// the `Exec` method returns a `Result` type instead of a `Row`
	// we follow the same argument pattern to add query params
	result, err := db.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", u.Username, u.Password)
	if err != nil {
		log.Errorf(log.DatabaseLogger, "could not insert row: %v", err)
	}

	// the `Result` type has special methods like `RowsAffected` which returns the
	// total number of affected rows reported by the database
	// In this case, it will tell us the number of rows that were inserted using
	// the above query
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Errorf(log.DatabaseLogger, "could not get affected rows: %v", err)
	}
	// we can log how many rows were inserted
	log.Debugln(log.DatabaseLogger, "User created. Inserted", rowsAffected, "rows")

	return nil
}

func (u *User) GetUserByUsername(db *sql.DB, username string) error {

	row := db.QueryRow("SELECT * FROM users WHERE username=$1 LIMIT 1", username)
	//user := models.User{}
	if err := row.Scan(&u.ID, &u.Username, &u.Password); err != nil {
		log.Errorf(log.DatabaseLogger, "could not get user by username: %v", err)
		return err
	}

	return nil
}

/*func (u *User) BeforeSave() error {
	hashedPassword, err := security.Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) Prepare() {
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
}

func (u *User) SaveUser(db *sql.DB) (*User, error) {

	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

// THE ONLY PERSON THAT NEED TO DO THIS IS THE ADMIN, SO I HAVE COMMENTED THE ROUTES, SO SOMEONE ELSE DONT VIEW THIS DETAILS.
func (u *User) FindAllUsers(db *gorm.DB) (*[]User, error) {
	var err error
	users := []User{}
	err = db.Debug().Model(&User{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}

func (u *User) FindUserByID(db *gorm.DB, uid uint32) (*User, error) {
	var err error
	err = db.Debug().Model(User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User Not Found")
	}
	return u, err
}

func (u *User) UpdateAUser(db *gorm.DB, uid uint32) (*User, error) {

	if u.Password != "" {
		// To hash the password
		err := u.BeforeSave()
		if err != nil {
			log.Fatal(err)
		}

		db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
			map[string]interface{}{
				"password":  u.Password,
				"email":     u.Email,
				"update_at": time.Now(),
			},
		)
	}

	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"email":     u.Email,
			"update_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &User{}, db.Error
	}

	// This is the display the updated user
	err := db.Debug().Model(&User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) UpdateAUserAvatar(db *gorm.DB, uid uint32) (*User, error) {
	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"avatar_path": u.AvatarPath,
			"update_at":   time.Now(),
		},
	)
	if db.Error != nil {
		return &User{}, db.Error
	}
	// This is the display the updated user
	err := db.Debug().Model(&User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) DeleteAUser(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).Delete(&User{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (u *User) UpdatePassword(db *gorm.DB) error {

	// To hash the password
	err := u.BeforeSave()
	if err != nil {
		log.Fatal(err)
	}

	db = db.Debug().Model(&User{}).Where("email = ?", u.Email).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"password":  u.Password,
			"update_at": time.Now(),
		},
	)
	if db.Error != nil {
		return db.Error
	}
	return nil
}*/
