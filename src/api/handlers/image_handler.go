package handlers

import (
	"covid_admission_api/services"

	"github.com/gin-gonic/gin"
)

type ImageHandler struct {
	imageService services.ImageService
}

func NewImageHandler(service *services.ImageService) *ImageHandler {
	return &ImageHandler{
		imageService: *service,
	}
}

func (handler *ImageHandler) ListAllImages(ctx *gin.Context) {

}

func (handler *ImageHandler) UploadImage(ctx *gin.Context) {

}

func (handler *ImageHandler) DeleteImage(ctx *gin.Context) {

}
