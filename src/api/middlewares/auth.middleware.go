package middlewares

import (
	"covid_admission_api/services"

	"github.com/labstack/echo/v4"
)

type AuthMiddleware interface {
	Auth(next echo.HandlerFunc) echo.HandlerFunc
	Refresh(next echo.HandlerFunc) echo.HandlerFunc
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
		cookie, err := c.Cookie("access-token")
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
		c.Set("uid", uid)
		return next(c)
	}
}

// for authenticating refresh token
func (a *authMiddleware) Refresh(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("refresh-token")
		if err != nil {
			return echo.ErrUnauthorized
		}
		detail, err := a.service.ExtractMetadata(cookie.Value)
		if err != nil {
			return echo.ErrUnauthorized
		}
		uid, err := a.service.FetchAuth(detail)
		if err != nil {
			return echo.ErrUnauthorized
		}
		c.Set("uid", uid)
		return next(c)
	}

}
