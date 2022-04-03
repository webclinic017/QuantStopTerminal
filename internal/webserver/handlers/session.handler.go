package handlers

import (
	"database/sql"
	"encoding/json"
	"github.com/quantstop/quantstopterminal/internal/database/models"
	"github.com/quantstop/quantstopterminal/internal/webserver/errors"
	"github.com/quantstop/quantstopterminal/internal/webserver/jwt"
	"github.com/quantstop/quantstopterminal/internal/webserver/write"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginResponse struct {
	ID       uint32   `json:"id"`
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
}

func Login(db *sql.DB, user *models.User, w http.ResponseWriter, r *http.Request) http.HandlerFunc {

	if db == nil {
		return write.Error(errors.NilDBError)
	}

	decoder := json.NewDecoder(r.Body)
	req := loginRequest{}
	err := decoder.Decode(&req)
	if err != nil || &req == nil {
		return write.Error(errors.NoJSONBody)
	}

	//log.Println("Username: " + req.Username)
	//log.Println("Password: " + req.Password)

	if req.Username == "" || req.Password == "" {
		return write.Error(errors.InvalidInput)
	}

	//log.Println("getting user")
	user = &models.User{}
	err = user.GetUserByUsername(db, req.Username)
	if err != nil {
		//log.Println("failed lookup")
		if err == sql.ErrNoRows {
			return write.Error(errors.FailedLogin)
		}
		return write.Error(err)
	}

	/*u, err = u.GetUserRoles(env.DB())
	if err != nil {
		return write.Error(err)
	}*/

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password+user.Salt))
	if err != nil {
		log.Println("failed pw hash: " + err.Error())
		return write.Error(errors.FailedLogin)
	}

	jwt.WriteUserCookie(w, user)

	res := loginResponse{
		ID:       user.ID,
		Username: user.Username,
		Roles: []string{
			"user",
			"moderator",
			"admin",
		},
	}
	return write.JSON(res)
	//return write.Success()
}

func Logout(db *sql.DB, user *models.User, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	u := &models.User{}
	jwt.WriteUserCookie(w, u)
	return write.Success()
}
