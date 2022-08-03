package service

import "github.com/GalloaFranco/gin-first-approach/entity"

type IVideoService interface {
	Save(video entity.Video) entity.Video
	FindAll() []entity.Video
}

// Struct that would implement IVideoService
type videoService struct {
	videos []entity.Video
}

// New Constructor function, to return instance of videoService (impl of interface)
// This function returns a value of interface type IVideoService, which contains a value of concrete type &videoService.
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
