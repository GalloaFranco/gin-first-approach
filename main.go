package main

import (
	"github.com/GalloaFranco/gin-first-approach/controller"
	"github.com/GalloaFranco/gin-first-approach/middlewares"
	"github.com/GalloaFranco/gin-first-approach/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var (
	videoService    service.IVideoService       = service.New()
	videoController controller.IVideoController = controller.New(videoService)
)

func main() {
	// Default method return an Engine instance with Logger and Recovery middleware
	//var server = gin.Default()

	var server = gin.New()
	var address = ":3000"

	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.Auth())

	// Context is the most important part of gin. It allows us to pass variables between middleware,
	// manage the flow, validate the JSON of a request and render a JSON response for example.
	server.GET("/hello", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Hello👋🏻",
		})
	})

	server.GET("/videos", videoController.FindAll)

	server.POST("/video", videoController.Save)

	log.Fatalln(server.Run(address))
}
