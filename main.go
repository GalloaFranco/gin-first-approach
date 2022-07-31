package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	// Default method return an Engine instance with Logger and Recovery middleware
	var server = gin.Default()
	var address = ":3000"

	// Context is the most important part of gin. It allows us to pass variables between middleware,
	// manage the flow, validate the JSON of a request and render a JSON response for example.
	server.GET("/hello", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Hello👋🏻",
		})
	})

	log.Fatalln(server.Run(address))
}
