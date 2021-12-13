package handlers

import (
	"covid_admission_api/entities"
	"covid_admission_api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService services.UserService
	jwtService  services.JWTService
}

func NewUserHandler(us services.UserService, js services.JWTService) *UserHandler {
	return &UserHandler{
		userService: us,
		jwtService:  js,
	}
}

func (handler *UserHandler) Register(ctx *gin.Context) {
	var newUser entities.User
	if err := ctx.ShouldBind(&newUser); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return

	}
	err := handler.userService.Register(&newUser)
	if err != nil {
		ctx.AbortWithStatus(http.StatusConflict)
		return

	}
	ctx.JSON(http.StatusNoContent, gin.H{})

}

func (handler *UserHandler) SignIn(ctx *gin.Context) {
	var user entities.User
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return

	}
	if err := handler.userService.SignIn(&user); err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return

	}
	ts, err := handler.jwtService.GenerateToken(user.UserId)
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnprocessableEntity)
		return

	}
	err = handler.userService.CreateAuth(user.UserId, ts)
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnprocessableEntity)
		return

	}
	ctx.JSON(http.StatusOK, gin.H{
		"access_token":  ts.AccessToken,
		"refresh_token": ts.RefreshToken,
	})

}

func (handler *UserHandler) SignOut(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{})
}
