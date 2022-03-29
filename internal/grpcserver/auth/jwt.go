package auth

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc/metadata"
	"time"
)

func CreateToken(userName string) (tokenString string) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":      "lora-app-server",
		"aud":      "lora-app-server",
		"nbf":      time.Now().Unix(),
		"exp":      time.Now().Add(time.Hour).Unix(),
		"sub":      "user",
		"username": userName,
	})
	tokenString, err := token.SignedString([]byte("verysecret"))
	if err != nil {
		panic(err)
	}
	return tokenString
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

// Claims defines the struct containing the token claims.
type Claims struct {
	jwt.StandardClaims

	// Username defines the identity of the user.
	Username string `json:"username"`
}

// Step1.  from  context  Of  metadata  in , Take out  token

func getTokenFromContext(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", fmt.Errorf("ErrNoMetadataInContext")
	}
	// md  The type is  type MD map[string][]string
	token, ok := md["authorization"]
	if !ok || len(token) == 0 {
		return "", fmt.Errorf("ErrNoAuthorizationInMetadata")
	}
	//  therefore ,token  Is an array of strings , We only use  token[0]
	return token[0], nil
}

func CheckAuth(ctx context.Context) (username string) {
	tokenStr, err := getTokenFromContext(ctx)
	if err != nil {
		panic("get token from context error")
	}
	var clientClaims Claims
	token, err := jwt.ParseWithClaims(tokenStr, &clientClaims, func(token *jwt.Token) (interface{}, error) {
		if token.Header["alg"] != "HS256" {
			panic("ErrInvalidAlgorithm")
		}
		return []byte("verysecret"), nil
	})
	if err != nil {
		panic("jwt parse error")
	}

	if !token.Valid {
		panic("ErrInvalidToken")
	}

	return clientClaims.Username
}
