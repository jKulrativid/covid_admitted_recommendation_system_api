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

	storageRepo := repositories.NewStorageRepository()
	storageService := services.StorageService(storageRepo)
	storageHandler := handlers.NewStorageHandler(storageService)

	r := echo.New()
	user := r.Group("/user")
	{
		user.POST("/register", userHandler.Register)
		user.POST("/signin", userHandler.SignIn)
		user.POST("/signout", userHandler.SignOut)
		user.POST("/updateuserdetail", userHandler.UpdateUserDetail)
		user.POST("/changepassword", userHandler.ChangePassword)
	}
	user.Use(authMiddleware.Auth)

	storage := r.Group("/storage")
	{
		storage.POST("/uploadfiles", storageHandler.UploadFiles)
		storage.GET("/listallfiles", storageHandler.ListAllFiles)
		storage.POST("/deletefiles", storageHandler.DeleteFiles)
	}

	r.POST("/refreshtoken", userHandler.RefreshToken, authMiddleware.Refresh)

	r.Use(middleware.Logger())

	return r

}
