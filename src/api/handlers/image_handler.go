package handlers

import (
	"covid_admission_api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ImageHandler struct {
	imageService services.ImageService
}

func NewImageHandler(service services.ImageService) *ImageHandler {
	return &ImageHandler{
		imageService: service,
	}
}

func (handler *ImageHandler) ListAllImages(ctx *gin.Context) {
	token := 
	
}

func (handler *ImageHandler) UplaodImage(ctx *gin.Context) {
	imgFile, imgHeader, err := ctx.Request.FormFile("img")
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

}

func (handler *ImageHandler) DeleteImage(ctx *gin.Context) {

}
