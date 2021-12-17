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
	var newUser entities.UserRegister
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
	var user entities.UserSignIn
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	userUuid, err := handler.userService.SignIn(&user)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}
	ts, err := handler.jwtService.GenerateToken(userUuid)
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnprocessableEntity)
		return
	}
	err = handler.userService.CreateAuth(userUuid, ts)
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

func (handler *UserHandler) RefreshToken(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{})
}

func (handler *UserHandler) UpdateUsername(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{})
}
