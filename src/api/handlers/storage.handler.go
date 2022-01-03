package handlers

import (
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
	uid, isAuth := c.Get("uid").(string)
	if !isAuth {
		return echo.ErrUnauthorized
	}
	multiForm, err := c.MultipartForm()
	if err != nil {
		return echo.ErrBadRequest
	}
	results := h.service.UploadFiles(uid, multiForm)
	return c.JSON(http.StatusOK, results)
}

// list all file name
func (h *storageHandler) ListAllFiles(c echo.Context) error {
	uid, isAuth := c.Get("uid").(string)
	if !isAuth {
		return echo.ErrUnauthorized
	}
	files, err := h.service.ListAllFiles(uid)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"result": files,
	})
}

// support multiple files delete
func (h *storageHandler) DeleteFiles(c echo.Context) error {
	uid, isAuth := c.Get("uid").(string)
	if !isAuth {
		return echo.ErrUnauthorized
	}
	files := make([]string, 0)
	if err := c.Bind(files); err != nil {
		return echo.ErrBadRequest
	}
	results := h.service.DeleteFiles(uid, files)
	return c.JSON(http.StatusOK, results)
}
