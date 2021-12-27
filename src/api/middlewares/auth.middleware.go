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

// TODO fix middleware after change framework from Gin to Echo
func (a *authMiddleware) Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("token")
		c.Set("isAuth", false) // always set isAuth false before authentication
		if err != nil {
			return next(c)
		}
		accessDetail, err := a.service.ExtractMetadata(cookie.Value)
		if err != nil {
			return next(c)
		}
		uid, err := a.service.FetchAuth(accessDetail)
		if err != nil {
			return next(c)
		}
		c.Set("isAuth", true)
		c.Set("uid", uid)
		return next(c)
	}
}
