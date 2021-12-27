package routers

import (
	"covid_admission_api/database"
	"covid_admission_api/handlers"
	"covid_admission_api/middlewares"
	"covid_admission_api/repositories"
	"covid_admission_api/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
		user.POST("/register", userHandler.Register)
		user.POST("/sign-in", userHandler.SignIn)
		user.POST("/sign-out", userHandler.SignOut)
	}
	user.Use(authMiddleware.Auth)

	r.POST("/refreshtoken", userHandler.RefreshToken, authMiddleware.Refresh)

	r.Validator = services.NewValidateService()
	r.Use(middleware.Logger())

	return r

}
