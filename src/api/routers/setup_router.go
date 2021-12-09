package routers

import (
	"covid_admission_api/handlers"
	"covid_admission_api/repositories"
	"covid_admission_api/services"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {

	imageRepo := repositories.NewImageRepository()
	imageService := services.NewImageService(*imageRepo)
	imageHandler := handlers.NewImageHandler(*imageService)

	r := gin.Default()
	image := r.Group("/image")
	{
		image.GET("list-all", imageHandler.ListAllImages)
		image.POST("upload", imageHandler.UplaodImage)
		image.DELETE("delete/:id", imageHandler.DeleteImage)
	}

	return r

}
