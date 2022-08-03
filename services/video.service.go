package services

import "github.com/GalloaFranco/gin-first-approach/entities"

type IVideoService interface {
	Save(video entities.Video) entities.Video
	FindAll() []entities.Video
}

// Struct that would implement IVideoService
type videoService struct {
	videos []entities.Video
}

// New Constructor function, to return instance of videoService (impl of interface)
// This function returns a value of interface type IVideoService, which contains a value of concrete type &videoService.
func New() IVideoService {
	return &videoService{}
}

// Save function impl
func (service *videoService) Save(video entities.Video) entities.Video {
	service.videos = append(service.videos, video)
	return video
}

// FindAll function impl
func (service *videoService) FindAll() []entities.Video {
	return service.videos
}
