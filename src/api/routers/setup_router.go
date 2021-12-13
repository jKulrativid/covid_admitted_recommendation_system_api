package routers

import (
	databases "covid_admission_api/database"
	"covid_admission_api/handlers"
	"covid_admission_api/repositories"
	"covid_admission_api/services"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {

	imageRepo := repositories.NewImageRepository() // TODO add conn field in ImageRepo
	imageService := services.NewImageService(*imageRepo)
	imageHandler := handlers.NewImageHandler(*imageService)

	userRepo := repositories.NewUserRepository(databases.DB, databases.RedisClient)
	userService := services.NewUserService(*userRepo)
	userHandler := handlers.NewUserHandler(*userService)

	r := gin.Default()
	image := r.Group("/image")
	{
		image.GET("list-all", imageHandler.ListAllImages)
		image.POST("upload", imageHandler.UploadImage)
		image.DELETE("delete/:id", imageHandler.DeleteImage)
	}

	user := r.Group("/user")
	{
		user.POST("register", userHandler.Register)
		user.POST("sign-in", userHandler.SignIn)
		user.POST("sign-out", userHandler.SignOut)
	}

	return r

}
