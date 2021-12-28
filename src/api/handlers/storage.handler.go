package handlers

import (
	"covid_admission_api/entities"
	"covid_admission_api/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type StorageHandler interface {
	UploadFiles(c echo.Context) error
	ListAllFiles(c echo.Context) error
	DeleteFiles(c echo.Context) error
}

type storageHandler struct {
	service services.StorageService
}

func NewStorageHandler(s services.StorageService) StorageHandler {
	return &storageHandler{
		service: s,
	}
}

// support multiple files upload
func (h *storageHandler) UploadFiles(c echo.Context) error {
	if isAuth := c.Get("isAuth").(bool); !isAuth {
		return echo.NewHTTPError(http.StatusUnauthorized, entities.ErrorUnAuthorized)
	}
	return nil
}

// list all file to JSON
func (h *storageHandler) ListAllFiles(c echo.Context) error {
	return nil
}

// support multiple files delete
func (h *storageHandler) DeleteFiles(c echo.Context) error {
	return nil
}
