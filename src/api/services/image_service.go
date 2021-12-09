package services

import (
	"covid_admission_api/entities"
	"covid_admission_api/repositories"
)

type ImageService struct {
	imageRepo repositories.ImageRepository
}

func NewImageService(repo repositories.ImageRepository) *ImageService {
	return &ImageService{
		imageRepo: repo,
	}
}

func (service *ImageService) ListAllImages() (images []entities.Image, err error) {
	var image []entities.Image

	return image, nil
}
