package main

import (
	"github.com/GalloaFranco/gin-first-approach/controllers"
	"github.com/GalloaFranco/gin-first-approach/docs"
	"github.com/GalloaFranco/gin-first-approach/middlewares"
	"github.com/GalloaFranco/gin-first-approach/repositories"
	"github.com/GalloaFranco/gin-first-approach/services"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
)

var (
	videoRepository repositories.IVideoRepository = repositories.NewVideoRepository()
	videoService    services.IVideoService        = services.New(videoRepository)
	videoController controllers.IVideoController  = controllers.NewVideoController(videoService)
	loginService    services.ILoginService        = services.NewLoginService()
	jwtService      services.IJWTService          = services.NewJWTService()
	loginController controllers.ILoginController  = controllers.NewLoginController(loginService, jwtService)
)

// @title Galloafranco - First approach to Gin Framework
// @version 1.0
// @description Gin first approach
// @BasePath /
// @schemes http
// @securityDefinitions.apikey bearerAuth
// @in header
// @name Authorization
func main() {

	// Default method return an Engine instance with Logger and Recovery middleware
	//var server = gin.Default()
	var server = gin.New()
	var address = ":8081"

	// server.Use(gin.Recovery(), middlewares.Logger(), middlewares.Auth())
	server.Use(gin.Recovery(), middlewares.Logger())
	defer videoRepository.CloseDB()

	// Context is the most important part of gin. It allows us to pass variables between middleware,
	// manage the flow, validate the JSON of a request and render a JSON response for example.
	server.GET("/hello", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Hello👋🏻",
		})
	})

	apiRoutes := server.Group(docs.SwaggerInfo.BasePath)
	{

		login := apiRoutes.Group("auth/")
		{
			login.POST("/login", loginController.Login)
		}

		videos := apiRoutes.Group("/videos")
		{
			videos.GET("", middlewares.AuthorizeJWT(), videoController.FindAll)

			videos.POST("", middlewares.AuthorizeJWT(), videoController.Save)
		}
	}
	server.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Fatalln(server.Run(address))
}
