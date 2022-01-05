package middlewares

import (
	"covid_admission_api/services"

	"github.com/labstack/echo/v4"
)

type AuthMiddleware interface {
	Auth(next echo.HandlerFunc) echo.HandlerFunc
}

type authMiddleware struct {
	service services.AuthService
}

func NewAuthMiddleware(s services.AuthService) AuthMiddleware {
	return &authMiddleware{
		service: s,
	}
}

// for authenticating access token
func (a *authMiddleware) Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		accessDetail, err := a.service.ExtractMetadata(token)
		if err != nil {
			return next(c)
		}
		uid, err := a.service.FetchAuth(accessDetail)
		if err != nil {
			return next(c)
		}
		c.Set("uid", uid)
		return next(c)
	}
}
