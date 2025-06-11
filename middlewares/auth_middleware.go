package middlewares

import (
	"minity/config"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(config.GetEnv("JWT_SECRET", "secret_key"))

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// take the header authorization from the request
		tokenString := c.GetHeader("Authorization")

		// if the token is empty, return a 401 Unauthorized response
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Token is required",
			})
			c.Abort()
			return
		}

		// delete the prefix "Bearer " from the token string
		// Header usually looks like "Bearer <token>"
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// make struct to hold the claims
		claims := &jwt.RegisteredClaims{}

		// Parse the token string and verify the signature using the jwtKey
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		// if the token is invalid or error when parsing
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid Token",
			})
			c.Abort()
			return
		}

		// save the sub claim (username) in the context
		c.Set("username", claims.Subject)

		// continue to the next handler
		c.Next()
	}
}
