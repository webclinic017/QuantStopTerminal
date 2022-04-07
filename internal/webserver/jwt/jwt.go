package jwt

import (
	"database/sql"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/quantstop/quantstopterminal/internal/database/models"
	"github.com/quantstop/quantstopterminal/internal/webserver/errors"
	"log"
	"net/http"
	"os"
	"time"
)

// jwt-cookie building and parsing
const cookieName = "gvsn"
const insecureSecret = "asd973hkalkjhx97asdh"

// tokens auto-refresh at the end of their lifetime,
// so long as the user hasn't been disabled in the interim
const tokenLifetime = time.Hour * 6

var hmacSecret []byte

func init() {
	hmacSecret = []byte(os.Getenv("API_SECRET"))
	if hmacSecret == nil {
		log.Fatal("No API_SECRET environment variable was found!")
	}
	if string(hmacSecret) == insecureSecret {
		log.Print("\n\n*** WARNING ***\nYour JWT isn't secure!\n")
		log.Print("You need to change your API_SECRET variable in .env (and restart your containers).\n\n")
	}
}

type userClaims struct {
	ID       uint32 `json:"id"`
	Username string `json:"username"`
	/*Email        string   `json:"email"`
	Status       string   `json:"status"`
	Verification string   `json:"verification"`
	Roles        []string `json:"roles"`*/
}

type claims struct {
	User *userClaims
	jwt.StandardClaims
}

// WriteUserCookie encodes a user's JWT and sets it as an httpOnly & Secure cookie
func WriteUserCookie(w http.ResponseWriter, u *models.User) {
	http.SetCookie(w, &http.Cookie{
		Name:     cookieName,
		Value:    encodeUser(u, time.Now()),
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   60 * 60 * 24 * 7, // one week
	})
}

// HandleUserCookie attempts to refresh an expired token if the user is still valid
func HandleUserCookie(db *sql.DB, w http.ResponseWriter, r *http.Request) (*models.User, error) {
	u, err := userFromCookie(r)

	// attempt refresh of expired token:
	if err == errors.ExpiredToken { //&& u.Status == models.UserStatusActive {
		user := models.User{}
		err = user.GetUserByUsername(db, u.Username) // todo: switch to id might be faster
		//err = env.DB().Debug().Model(models.User{}).Where("username = ?", u.Username).Take(&user).Error
		if err != nil {
			fmt.Println("this is the error getting the user: ", err)
			return wipeCookie(db, w)
		}

		/*if user.Status == models.UserStatusActive {
			WriteUserCookie(w, &user)
			return &user, nil
		} else {
			// their account isn't verified, log them out
			return wipeCookie(db, w)
		}*/
	}

	if err != nil {
		return nil, err
	}

	return u, err
}

func wipeCookie(db *sql.DB, w http.ResponseWriter) (*models.User, error) {
	u := &models.User{}
	WriteUserCookie(w, u)
	return u, nil
}

// userFromCookie builds a user object from a JWT, if it's valid
func userFromCookie(r *http.Request) (*models.User, error) {
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		log.Println(err)
	}

	var tokenString string
	if cookie != nil {
		tokenString = cookie.Value
	}

	if tokenString == "" {
		log.Println("token string empty")
		return &models.User{}, nil
	}

	return decodeUser(tokenString)
}

// encodeUser convert a user struct into a jwt
func encodeUser(u *models.User, t time.Time) (tokenString string) {

	// convert []Role to []string
	/*var roles []string
	for _, role := range u.Roles {
		roles = append(roles, role.Name)
	}

	for _, r := range roles {
		log.Println(r)
	}*/

	claims := claims{
		&userClaims{
			ID:       u.ID,
			Username: u.Username,
			/*Email:        u.Email,
			Status:       u.Status,
			Verification: u.Verification,
			Roles:        roles,*/
		},
		jwt.StandardClaims{
			IssuedAt:  t.Add(-time.Second).Unix(),
			ExpiresAt: t.Add(tokenLifetime).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// unhandled err here
	tokenString, err := token.SignedString(hmacSecret)
	if err != nil {
		log.Println("Error signing token", err)
	}
	return
}

// decodeUser converts a jwt into a user struct (or returns a zero-value user)
func decodeUser(tokenString string) (*models.User, error) {

	// try parsing token
	token, err := jwt.ParseWithClaims(tokenString, &claims{}, func(token *jwt.Token) (interface{}, error) {
		// is token.Method type of/can be converted to *jwt.SigningMethodHMAC ?
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSecret, nil
	})

	// errors
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				fmt.Println("ValidationError, token malformed.")
				return nil, errors.InvalidToken
			} else if ve.Errors&(jwt.ValidationErrorExpired) != 0 {
				// Token is expired
				fmt.Println("ValidationError, token expired.")
				return nil, errors.ExpiredToken
			} else if ve.Errors&(jwt.ValidationErrorNotValidYet) != 0 {
				// Token is not active yet
				fmt.Println("ValidationError, token not active yet.")
				return nil, errors.InvalidToken
			} else {
				fmt.Println("ValidationError, Couldn't handle this token:", err)
				return nil, errors.InvalidToken
			}
		} else {
			fmt.Println("Couldn't handle this token:", err)
			return nil, errors.InvalidToken
		}
	}

	// token valid?
	if token.Valid {
		return getUserFromToken(token), nil
	}
	return nil, errors.InvalidToken

}

func getUserFromToken(token *jwt.Token) *models.User {
	if c, ok := token.Claims.(*claims); ok {
		/*var roles []models.Role
		for _, role := range c.User.Roles {
			roles = append(roles, models.Role{
				Name: role,
			})
		}*/
		return &models.User{
			ID:       c.User.ID,
			Username: c.User.Username,
			/*Email:        c.User.Email,
			Status:       c.User.Status,
			Verification: c.User.Verification,
			Roles:        roles,*/
		}
	}

	return &models.User{}
}
