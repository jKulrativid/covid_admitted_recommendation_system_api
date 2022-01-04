package repositories

import (
	"covid_admission_api/entities"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

type StorageRepository interface {
	UploadFile(f *entities.UploadFile) error
	ListAllFileNames(uid string) ([]string, error)
	DeleteFile(uid string, fileName string) error
}

type storageRepository struct {
	storagePath string
}

func NewStorageRepository() StorageRepository {
	storagePath := os.Getenv("STORAGE_PATH")
	if storagePath == "" {
		storagePath = "storage"
	}
	if err := os.MkdirAll(storagePath, 0777); err != nil {
		panic(err)
	}
	return &storageRepository{storagePath: storagePath}
}

func (r *storageRepository) UploadFile(f *entities.UploadFile) error {
	dstPath := filepath.Join(r.storagePath, f.UploaderUid, f.FileHeader.Filename)
	if isExist, err := r.exists(dstPath); err != nil {
		return entities.ErrorInternalServer
	} else if isExist {
		return entities.ErrorConflict
	}
	src, err := f.FileHeader.Open()
	if err != nil {
		return entities.ErrorUnprocessableEntity
	}
	defer src.Close()

	dst, err := os.Create(dstPath)
	if err != nil {
		return entities.ErrorInternalServer
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return entities.ErrorInternalServer
	}
	return nil
}

func (r *storageRepository) ListAllFileNames(uid string) ([]string, error) {
	files, err := ioutil.ReadDir(filepath.Join(r.storagePath, uid))
	if err != nil {
		return nil, entities.ErrorUnprocessableEntity
	}
	fileLists := make([]string, len(files))
	for i, f := range files {
		fileLists[i] = f.Name()
	}
	return fileLists, nil
}

func (r *storageRepository) DeleteFile(uid string, fileName string) error {
	filePath := filepath.Join(r.storagePath, uid, fileName)
	exist, err := r.exists(filePath)
	if err != nil {
		return entities.ErrorUnprocessableEntity
	}
	if !exist {
		return entities.ErrorNotFound
	}
	if err := os.Remove(filePath); err != nil {
		return entities.ErrorUnprocessableEntity
	}
	return nil
}

func (r *storageRepository) exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
