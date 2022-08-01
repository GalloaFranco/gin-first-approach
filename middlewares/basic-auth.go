package middlewares

import "github.com/gin-gonic/gin"

func Auth() gin.HandlerFunc { // HandlerFunc is the type expected for all the middlewares
	// BasicAuth returns a Basic HTTP Authorization middleware. It takes as argument a map[string]string where
	// the key is the user name and the value is the password.
	return gin.BasicAuth(gin.Accounts{
		"admin": "123",
	})
}
