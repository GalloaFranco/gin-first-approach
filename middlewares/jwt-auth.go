package middlewares

import (
	"github.com/GalloaFranco/gin-first-approach/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// AuthorizeJWT validates the token from the http request, returning a 401 if it's not valid
func AuthorizeJWT() gin.HandlerFunc {
	return func(context *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := context.GetHeader("Authorization")

		if authHeader == "" {
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		tokenString := authHeader[len(BEARER_SCHEMA):] // Slice from the (BEARER_SCHEMA length) character

		token, err := services.NewJWTService().ValidateToken(tokenString)

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claims[Name]: ", claims["name"])
			log.Println("Claims[Admin]: ", claims["admin"])
			log.Println("Claims[Issuer]: ", claims["iss"])
			log.Println("Claims[IssuedAt]: ", claims["iat"])
			log.Println("Claims[ExpiresAt]: ", claims["exp"])

		} else {
			log.Println("Error: ", err)
			context.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
