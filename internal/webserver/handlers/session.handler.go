package handlers

import (
	"database/sql"
	"encoding/json"
	"github.com/quantstop/quantstopterminal/internal/database/models"
	"github.com/quantstop/quantstopterminal/internal/webserver/errors"
	"github.com/quantstop/quantstopterminal/internal/webserver/jwt"
	"github.com/quantstop/quantstopterminal/internal/webserver/utils"
	"github.com/quantstop/quantstopterminal/internal/webserver/write"
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

type signupRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
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

	if req.Username == "" || req.Password == "" {
		return write.Error(errors.InvalidInput)
	}

	user = &models.User{}
	err = user.GetUserByUsername(db, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("cannot get user by username")
			return write.Error(errors.FailedLogin)
		}
		return write.Error(errors.FailedLogin)
	}

	if !utils.CheckPasswordHash(req.Password, user.Salt, user.Password) {
		return write.Error(errors.FailedLogin)
	}

	jwt.WriteUserCookie(w, user)

	res := loginResponse{
		ID:       user.ID,
		Username: user.Username,
		Roles:    user.Roles,
	}
	return write.JSON(res)
}

func Logout(db *sql.DB, user *models.User, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	u := &models.User{}
	jwt.WriteUserCookie(w, u)
	return write.Success()
}

func Signup(db *sql.DB, user *models.User, w http.ResponseWriter, r *http.Request) http.HandlerFunc {

	if db == nil {
		return write.Error(errors.NilDBError)
	}

	decoder := json.NewDecoder(r.Body)
	req := signupRequest{}
	err := decoder.Decode(&req)
	if err != nil || &req == nil {
		return write.Error(errors.NoJSONBody)
	}

	if req.Username == "" || req.Password == "" {
		return write.Error(errors.InvalidInput)
	}

	user = &models.User{
		Username: req.Username,
	}

	// Set salt, and hash password with salt
	user.Salt = utils.GenerateRandomString(32)
	user.Password, err = utils.HashPassword(req.Password, user.Salt)
	if err != nil {
		return write.Error(err)
	}

	err = user.CreateUser(db)
	if err != nil {
		//todo: can we get more specific errors? do we even need to?
		return write.Error(err)
	}

	return write.Success()
}
