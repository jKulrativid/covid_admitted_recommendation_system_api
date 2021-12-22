package handlers

import (
	"covid_admission_api/entities"
	"covid_admission_api/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	Register(c echo.Context) error
	SignIn(c echo.Context) error
	SignOut(c echo.Context) error
	RefreshToken(c echo.Context) error
	UpdateUsername(c echo.Context) error
}

type userHandler struct {
	service services.UserService
}

func NewUserHandler(us services.UserService) UserHandler {
	return &userHandler{
		service: us,
	}
}

func (h *userHandler) Register(c echo.Context) error {
	var newUser entities.UserRegister
	if err := c.Bind(&newUser); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(&newUser); err != nil {
		return err
	}
	err := h.service.Register(&newUser)
	if err != nil {
		return c.JSON(http.StatusConflict, err)
	}
	return c.NoContent(http.StatusNoContent)
}

func (h *userHandler) SignIn(c echo.Context) error {
	var user entities.UserSignIn
	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(&user); err != nil {
		return err
	}
	userUuid, err := h.service.SignIn(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
	ts, err := h.service.GenerateToken(userUuid)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	err = h.service.CreateAuth(userUuid, ts)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]string{
		"access_token":  ts.AccessToken,
		"refresh_token": ts.RefreshToken,
	})
}

func (h *userHandler) SignOut(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{})
}

func (h *userHandler) RefreshToken(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{})
}

func (h *userHandler) UpdateUsername(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{})
}
