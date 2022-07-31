package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	var router = gin.Default()
	var address = ":3000"

	router.GET("/hello", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello!")
	})

	log.Fatalln(router.Run(address))
}
