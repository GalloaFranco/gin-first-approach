package controller

import (
	"github.com/GalloaFranco/gin-first-approach/entity"
	"github.com/GalloaFranco/gin-first-approach/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IVideoController interface {
	FindAll(context *gin.Context)
	Save(context *gin.Context)
}

type controller struct {
	service service.IVideoService
}

func New(service service.IVideoService) IVideoController {
	return &controller{
		service: service,
	}
}

func (controller *controller) FindAll(context *gin.Context) {
	context.JSON(http.StatusOK, controller.service.FindAll())
}

func (controller *controller) Save(context *gin.Context) {
	var video entity.Video
	// MustBindWith binds the passed struct pointer using the specified binding engine.
	if err := context.BindJSON(&video); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}
	video = controller.service.Save(video)
	context.JSON(http.StatusCreated, video)
}
