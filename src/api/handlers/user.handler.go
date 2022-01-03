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
	UpdateUserDetail(c echo.Context) error
	ChangePassword(c echo.Context) error
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
		return echo.ErrBadRequest
	}
	if err := h.service.Validate(newUser); err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, err.Error())
	}
	err := h.service.Register(&newUser)
	if err != nil {
		return echo.NewHTTPError(http.StatusConflict, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

func (h *userHandler) SignIn(c echo.Context) error {
	var user entities.UserSignIn
	if err := c.Bind(&user); err != nil {
		return echo.ErrBadRequest
	}
	if err := h.service.Validate(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	userUuid, err := h.service.SignIn(&user)
	if err != nil {
		return echo.ErrUnauthorized
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
	var user entities.UserSignIn
	if err := c.Bind(&user); err != nil {
		return echo.ErrBadRequest
	}
	if err := h.service.Validate(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	uid, err := h.service.SignOut(&user)
	if err != nil {
		return echo.ErrUnauthorized
	}
	if err := h.service.DeleteAuth(uid); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

func (h *userHandler) RefreshToken(c echo.Context) error {
	uid, isAuth := c.Get("uid").(string)
	if !isAuth {
		return echo.ErrUnauthorized
	}

	// delete previous refresh token
	h.service.DeleteAuth(uid)

	// regenerate access and refresh token
	ts, err := h.service.GenerateToken(uid)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	if err := h.service.CreateAuth(uid, ts); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{
		"access_token":  ts.AccessToken,
		"refresh_token": ts.RefreshToken,
	})
}

func (h *userHandler) UpdateUserDetail(c echo.Context) error {
	if isAuth := c.Get("isAuth"); isAuth == nil {
		return echo.ErrUnauthorized
	}
	return c.JSON(http.StatusOK, map[string]string{})
}

func (h *userHandler) ChangePassword(c echo.Context) error {
	if isAuth := c.Get("isAuth"); isAuth == nil {
		return echo.ErrUnauthorized
	}
	return c.JSON(http.StatusOK, map[string]string{})
}
