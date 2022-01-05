package services

import (
	"covid_admission_api/entities"
	"covid_admission_api/repositories"
	"mime/multipart"
	"strings"
)

type StorageService interface {
	UploadFiles(uid string, mf *multipart.Form) map[string]string
	uploadFile(f *entities.UploadFile) error
	ListAllFiles(uid string) ([]string, error)
	DeleteFiles(uid string, deleteList []string) map[string]string
	deleteFile(uid string, fileName string) error
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
func (s *storageService) uploadFile(f *entities.UploadFile) error {
	if err := s.validateFileType(f); err != nil {
		return err
	}
	if err := s.repo.UploadFile(f); err != nil {
		return err
	}
	return nil
}

func (s *storageService) UploadFiles(uid string, mf *multipart.Form) map[string]string {
	files := mf.File["upload"]
	results := make(map[string]string, len(files))
	for _, file := range files {
		uploadFile := entities.UploadFile{
			UploaderUid: uid,
			FileHeader:  file,
		}
		uploadResult := entities.StatusUploadSuccess
		if err := s.uploadFile(&uploadFile); err != nil {
			uploadResult = err.Error()
		}
		results[file.Filename] = uploadResult
	}
	return results
}

// list all file as map of filename and description
func (s *storageService) ListAllFiles(uid string) ([]string, error) {
	nameLists, err := s.repo.ListAllFileNames(uid)
	if err != nil {
		return nil, err
	}
	return nameLists, nil
}

// this service delete one file at a time
func (s *storageService) deleteFile(uid string, fileName string) error {
	if err := s.repo.DeleteFile(uid, fileName); err != nil {
		return err
	}
	return nil
}

func (s *storageService) DeleteFiles(uid string, deleteList []string) map[string]string {
	results := make(map[string]string, len(deleteList))
	for _, fileName := range deleteList {
		deleteResult := entities.StatusDeleteSuccess
		if err := s.deleteFile(uid, fileName); err != nil {
			deleteResult = err.Error()
		}
		results[fileName] = deleteResult
	}
	return results
}

func (s *storageService) validateFileType(f *entities.UploadFile) error {
	fileType := strings.Split(f.FileHeader.Filename, ".")[1]
	if !entities.AllowedFileType[fileType] {
		return entities.ErrorUnsupportedMediaType
	}
	return nil
}
