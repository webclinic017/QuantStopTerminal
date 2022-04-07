package handlers

import (
	"database/sql"
	"github.com/quantstop/quantstopterminal/internal/database/models"
	"github.com/quantstop/quantstopterminal/internal/webserver/write"
	"net/http"
)

/*func Signup(db *sql.DB, usr *models.User, w http.ResponseWriter, r *http.Request) http.HandlerFunc {

	// Decode the json request body into User struct
	decoder := json.NewDecoder(r.Body)
	var user models.User
	err := decoder.Decode(&user)
	if err != nil || &user == nil {
		return write.Error(errors.NoJSONBody)
	}

	// Set salt, and hash password with salt
	user.Salt = utils.GenerateRandomString(32)
	user.Password, err = utils.HashPassword(user.Password, user.Salt)
	if err != nil {
		return write.Error(err)
	}

	// Set user status, and generate random verification code
	user.Status = models.UserStatusUnverified
	user.Verification = utils.GenerateRandomString(32)

	// Prepare and validate the user data
	user.Prepare()
	err = user.Validate("")
	if err != nil {
		return write.Error(err)
	}

	// All good, try saving the user to the database
	_, err = user.SaveUser(env.DB())
	if err != nil {
		//todo: can we get more specific errors? do we even need to?
		//formattedError := formaterror.FormatError(err.Error())
		if isDupe(err) {
			return write.Error(errors.AlreadyRegistered)
		}
		return write.Error(err)
	}

	// User is saved, lets send the verification email
	// Note: when using dev mode, link is sent to log output
	err = env.Mailer().VerifyEmail(user.Email, user.Verification)
	if err != nil {
		return write.Error(err)
	}

	return write.Success()
}*/

/*func UpdatePassword(db *sql.DB, user *models.User, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	if user.Status != models.UserStatusActive {
		return write.Error(errors.RouteUnauthorized)
	}

	decoder := json.NewDecoder(r.Body)
	var u models.User
	err := decoder.Decode(&u)
	if err != nil || &u == nil {
		return write.Error(errors.NoJSONBody)
	}

	// salt and hash it
	u.Salt = utils.GenerateRandomString(32)
	u.Password, err = utils.HashPassword(u.Password, u.Salt)
	if err != nil {
		return write.Error(err)
	}

	// todo:
	err = env.DB().UpdateUserPassword(r.Context(), db.UpdateUserPasswordParams{
		ID:   user.ID,
		Pass: u.Pass,
		Salt: u.Salt,
	})
	if err != nil {
		return write.Error(err)
	}

	return write.Success()
}*/

type WhoamiResponse struct {
	ID       uint32 `json:"id"`
	Username string `json:"username"`
	//Email    string   `json:"email"`
	Roles []string `json:"roles"`
}

func Whoami(db *sql.DB, user *models.User, w http.ResponseWriter, r *http.Request) http.HandlerFunc {

	// We don't want to send all the data from the Role table for each role
	// Create a new array of strings, we will send just the role name
	/*var roles []string
	for _, role := range user.Roles {
		roles = append(roles, role.Name)
	}*/

	res := WhoamiResponse{
		ID:       user.ID,
		Username: user.Username,
		//Email:    user.Email,
		Roles: []string{
			"user",
			"moderator",
			"admin",
		},
	}
	return write.JSON(res)
}

type verifyRequest struct {
	Code string
}

/*func Verify(db *sql.DB, user *models.User, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	decoder := json.NewDecoder(r.Body)
	var req verifyRequest
	err := decoder.Decode(&req)
	if err != nil || &req == nil || req.Code == "" {
		return write.Error(errors.NoJSONBody)
	}

	u := &models.User{}
	u, err = u.FindUserByVerificationCode(env.DB(), req.Code)
	if err != nil {
		return write.Error(err)
	}

	if u.Status != models.UserStatusUnverified {
		return write.Error(errors.VerificationExpired)
	}

	u.Status = models.UserStatusActive

	err = u.SetUserStatusActive(env.DB(), u.ID)
	if err != nil {
		return write.Error(err)
	}

	jwt.WriteUserCookie(w, u)
	return write.JSON(&u)
}*/
