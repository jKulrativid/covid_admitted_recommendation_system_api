package routers

import (
	"covid_admission_api/database"
	"covid_admission_api/handlers"
	"covid_admission_api/middlewares"
	"covid_admission_api/repositories"
	"covid_admission_api/services"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

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

	r.POST("/refreshtoken", userHandler.RefreshToken)

	userEdit := r.Group("/user-edit")
	{
		userEdit.POST("/updata-username", userHandler.UpdateUsername)
	}
	r.Validator = &CustomValidator{validator: validator.New()}
	r.Use(middleware.Logger())
	r.Use(authMiddleware.Auth)

	return r

}
