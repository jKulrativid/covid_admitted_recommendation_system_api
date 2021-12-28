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
	if isAuth := c.Get("isAuth").(bool); !isAuth {
		return echo.ErrUnauthorized
	}
	uploaderUid := c.Get("uid").(string)
	multiForm, err := c.MultipartForm()
	if err != nil {
		return echo.ErrBadRequest
	}
	results := h.service.UploadFiles(uploaderUid, multiForm)
	return c.JSON(http.StatusOK, results)
}

// list all file name
func (h *storageHandler) ListAllFiles(c echo.Context) error {
	if isAuth := c.Get("isAuth").(bool); !isAuth {
		return echo.ErrUnauthorized
	}
	uid := c.Get("uid").(string)
	nameList, err := h.service.ListAllFiles(uid)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"list": nameList,
	})
}

// support multiple files delete
func (h *storageHandler) DeleteFiles(c echo.Context) error {
	if isAuth := c.Get("isAuth").(bool); !isAuth {
		return echo.ErrUnauthorized
	}
	uid := c.Get("uid").(string)
	fileList := []string{}
	if err := c.Bind(&fileList); err != nil {
		return echo.ErrBadRequest
	}
	results := h.service.DeleteFiles(uid, fileList)
	return c.JSON(http.StatusOK, results)
}
