package main

import (
	"github.com/GalloaFranco/gin-first-approach/controllers"
	"github.com/GalloaFranco/gin-first-approach/middlewares"
	"github.com/GalloaFranco/gin-first-approach/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var (
	videoService    services.IVideoService       = services.New()
	videoController controllers.IVideoController = controllers.NewVideoController(videoService)
	loginService    services.ILoginService       = services.NewLoginService()
	jwtService      services.IJWTService         = services.NewJWTService()
	loginController controllers.ILoginController = controllers.NewLoginController(loginService, jwtService)
)

func main() {
	// Default method return an Engine instance with Logger and Recovery middleware
	//var server = gin.Default()

	var server = gin.New()
	var address = ":8081"

	// server.Use(gin.Recovery(), middlewares.Logger(), middlewares.Auth())
	server.Use(gin.Recovery(), middlewares.Logger())

	// Context is the most important part of gin. It allows us to pass variables between middleware,
	// manage the flow, validate the JSON of a request and render a JSON response for example.
	server.GET("/hello", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Hello👋🏻",
		})
	})

	server.POST("/login", loginController.Login)

	server.GET("/videos", middlewares.AuthorizeJWT(), videoController.FindAll)

	server.POST("/video", middlewares.AuthorizeJWT(), videoController.Save)

	log.Fatalln(server.Run(address))
}
