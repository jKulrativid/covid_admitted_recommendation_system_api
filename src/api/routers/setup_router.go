package routers

import (
	databases "covid_admission_api/database"
	"covid_admission_api/handlers"
	"covid_admission_api/middlewares"
	"covid_admission_api/repositories"
	"covid_admission_api/services"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {

	jwtService := services.NewJWTService()

	userRepo := repositories.NewUserRepository(databases.DB, databases.RedisClient)
	userService := services.NewUserService(*userRepo)
	userHandler := handlers.NewUserHandler(*userService, *jwtService)

	r := gin.Default()
	user := r.Group("/user")
	{
		user.POST("register", userHandler.Register)
		user.POST("sign-in", userHandler.SignIn)
		user.POST("sign-out", userHandler.SignOut)
		user.POST("refresh-token", userHandler.RefreshToken)
	}

	userEdit := r.Group("/user-edit")
	userEdit.Use(middlewares.AuthorizeJWT(jwtService))
	{
		userEdit.POST("updata-username", userHandler.UpdateUsername)
	}

	return r

}
