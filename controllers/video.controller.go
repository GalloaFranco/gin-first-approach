package controllers

import (
	"github.com/GalloaFranco/gin-first-approach/entities"
	"github.com/GalloaFranco/gin-first-approach/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IVideoController interface {
	FindAll(context *gin.Context)
	Save(context *gin.Context)
}

type controller struct {
	service services.IVideoService
}

// NewVideoController Constructor function, to return instance of controllers (impl of interface)
// This function returns a value of interface type IVideoController, which contains a value of concrete type &controllers.
func NewVideoController(service services.IVideoService) IVideoController {
	return &controller{
		service: service,
	}
}

func (controller *controller) FindAll(context *gin.Context) {
	context.JSON(http.StatusOK, controller.service.FindAll())
}

func (controller *controller) Save(context *gin.Context) {
	var video entities.Video
	// MustBindWith binds the passed struct pointer using the specified binding engine.
	if err := context.BindJSON(&video); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}
	video = controller.service.Save(video)
	context.JSON(http.StatusCreated, video)
}
