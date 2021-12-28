package services

import "covid_admission_api/repositories"

type StorageService interface {
}

type storageService struct {
	repo repositories.StorageRepository
}

func NewStorageService(r repositories.StorageRepository) StorageService {
	return &storageService{
		repo: r,
	}
}

// this service upload one file at a time
func (s *storageService) UploadFile() error {
	return nil
}

// list all file as map of filename and description
func (s *storageService) ListAllFiles() (map[string]interface{}, error) {
	return nil, nil
}

// this service delete one file at a time
func (s *storageService) DeleteFile() error {
	return nil
}

func (s *storageService) ValidateFile() error {
	return nil
}
