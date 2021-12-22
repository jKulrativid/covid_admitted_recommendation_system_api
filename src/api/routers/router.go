package routers

import (
	"covid_admission_api/database"
	"covid_admission_api/handlers"
	"covid_admission_api/middlewares"
	"covid_admission_api/repositories"
	"covid_admission_api/services"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewRouter(db database.Database, rs database.RedisClient) *echo.Echo {

	authRepo := repositories.NewAuthRepo(rs)
	authService := services.NewAuthService(authRepo)
	authMiddleware := middlewares.NewAuthMiddleware(authService)

	userRepo := repositories.NewUserRepository(db, rs)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	r := echo.New()
	user := r.Group("/user")
	{
		user.POST("register", userHandler.Register)
		user.POST("sign-in", userHandler.SignIn)
		user.POST("sign-out", userHandler.SignOut)
	}

	r.POST("/refreshtoken", userHandler.RefreshToken)

	userEdit := r.Group("/user-edit")
	{
		userEdit.POST("updata-username", userHandler.UpdateUsername)
	}

	r.Use(middleware.Logger())
	r.Use(authMiddleware.Auth)

	return r

}