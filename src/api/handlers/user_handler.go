package handlers

import (
	"covid_admission_api/entities"
	"covid_admission_api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{
		userService: service,
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
	var newUser entities.User
	if err := ctx.ShouldBind(&newUser); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return

	}
	jwtToken, err := handler.userService.SignIn(&newUser)
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return

	}
	ctx.JSON(http.StatusOK, gin.H{
		"token": jwtToken,
	})

}

func (handler *UserHandler) SignOut(ctx *gin.Context) {

}
