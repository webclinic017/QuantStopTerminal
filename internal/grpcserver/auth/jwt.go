package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/quantstop/quantstopterminal/internal/log"
	"google.golang.org/grpc/metadata"
	"time"
)

func CreateToken(userName string) (tokenString string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":      "quantstopterminal",
		"aud":      "quantstopterminal",
		"nbf":      time.Now().Unix(),
		"exp":      time.Now().Add(time.Hour).Unix(),
		"sub":      "user",
		"username": userName,
	})
	tokenString, err = token.SignedString([]byte("ajaeyuiuop5n32o9s0g"))
	if err != nil {
		log.Errorf(log.GRPClog, "jwt error: %v", err)
		return "", err
	}
	return tokenString, nil
}

// Token  Custom authentication
type Token struct {
	Token string
}

func (c Token) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": c.Token,
	}, nil
}

func (c Token) RequireTransportSecurity() bool {
	return false
}

// Claims defines the struct containing the token claims
type Claims struct {
	jwt.StandardClaims

	// Username defines the identity of the user.
	Username string `json:"username"`
}

// getTokenFromContext reads incoming context and extracts auth token from metadata
func getTokenFromContext(ctx context.Context) (string, error) {

	// try and get the metadata from context
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", fmt.Errorf("ErrNoMetadataInContext")
	}

	// metadata type is map[string][]string of all request headers
	// token stored in authorization header
	// try and get token from header
	authHeader, ok := md["authorization"]
	if !ok || len(authHeader) == 0 {
		return "", fmt.Errorf("ErrNoAuthorizationInMetadata")
	}

	// success, authHeader is an array of strings, token stored in first slice
	return authHeader[0], nil
}

// CheckAuth verifies jwt
func CheckAuth(ctx context.Context) (username string, err error) {

	// try and get the token from incoming header
	tokenStr, err := getTokenFromContext(ctx)
	if err != nil {
		return "", err
	}

	// try and parse token into clientClaims
	var clientClaims Claims
	token, err := jwt.ParseWithClaims(tokenStr, &clientClaims, func(token *jwt.Token) (interface{}, error) {

		// check if algo matches
		if token.Header["alg"] != "HS256" {
			//panic("ErrInvalidAlgorithm")
			return "", errors.New("invalid token algorithm")
		}

		// return secret, no error
		return []byte("ajaeyuiuop5n32o9s0g"), nil
	})

	// return any error from token parsing
	if err != nil {
		return "", err
	}

	// return error if token isn't valid
	if !token.Valid {
		return "", err
	}

	// success return username
	return clientClaims.Username, nil
}
