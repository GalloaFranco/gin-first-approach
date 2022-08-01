package controller

import (
	"github.com/GalloaFranco/gin-first-approach/entity"
	"github.com/GalloaFranco/gin-first-approach/service"
	"github.com/gin-gonic/gin"
)

type IVideoController interface {
	FindAll() []entity.Video
	Save(context *gin.Context) entity.Video
}

type controller struct {
	service service.IVideoService
}

func New(service service.IVideoService) IVideoController {
	return &controller{
		service: service,
	}
}

func (controller *controller) FindAll() []entity.Video {
	//TODO implement me
	panic("implement me")
}

func (controller *controller) Save(context *gin.Context) entity.Video {
	var video entity.Video
	context.BindJSON(&video)
	controller.service.Save(video)
	return video
}
