package helpers

import (
	"minity/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(config.GetEnv("JWT_SECRET", "secret_key"))

func GenerateToken(username string) string {
	// setting time expired token, in here, we will set to 60 minutes from now
	expirationTime := time.Now().Add(60 * time.Minute)

	// make JWT claims
	// subject value is username and ExpiresAt looking for expiration token
	claims := &jwt.RegisteredClaims{
		Subject:   username,
		ExpiresAt: jwt.NewNumericDate(expirationTime),
	}

	// make new token with new create claims
	// using HS256 algorithm for signing the token
	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtKey)

	// return the token in the form of string
	return token
}
