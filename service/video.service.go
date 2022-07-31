package service

import "github.com/GalloaFranco/gin-first-approach/entity"

type IVideoService interface {
	Save(video entity.Video) entity.Video
	FindAll() []entity.Video
}

// Struct that implement IVideoService
type videoService struct {
	videos []entity.Video
}

// New Constructor function, to return instance of video (impl of interface)
func New() IVideoService {
	return &videoService{}
}

// Save function impl
func (service *videoService) Save(video entity.Video) entity.Video {
	service.videos = append(service.videos, video)
	return video
}

// FindAll function impl
func (service *videoService) FindAll() []entity.Video {
	return service.videos
}
