package crypto

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

// FIXME: Replace this secret with env vars!!!

var (
	SecretKey = []byte("secret")
)

// Generate a jwt token assigning a username to it's claims and return it.
func GenerateToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenStr, err := token.SignedString(SecretKey)
	if err != nil {
		log.Fatal("Error in Generating key")
		return "", err

	}
	return tokenStr, nil

}

// Parse a jwt token and return its username within the claim.
func ParseToken(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil

	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := claims["username"].(string)
		return username, nil

	} else {
		return "", err

	}

}
