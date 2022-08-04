package services

import (
	"github.com/GalloaFranco/gin-first-approach/entities"
	"github.com/GalloaFranco/gin-first-approach/repositories"
)

type IVideoService interface {
	Save(video entities.Video) entities.Video
	FindAll() []entities.Video
}

// Struct that would implement IVideoService
type videoService struct {
	videoRepository repositories.IVideoRepository
	videos          []entities.Video
}

// New Constructor function, to return instance of videoService (impl of interface)
// This function returns a value of interface type IVideoService, which contains a value of concrete type &videoService.
func New(repository repositories.IVideoRepository) IVideoService {
	return &videoService{
		videoRepository: repository,
	}
}

// Save function impl
func (service *videoService) Save(video entities.Video) entities.Video {
	//service.videos = append(service.videos, video)
	service.videoRepository.Save(video)
	return video
}

// FindAll function impl
func (service *videoService) FindAll() []entities.Video {
	//return service.videos
	return service.videoRepository.FindAll()
}
