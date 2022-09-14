package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	//TODO:Add Auth service here
}

func NewAuthController() AuthController {
	return &authController{}
}

func (a authController) Login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Hello Login")
}

func (a authController) Register(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Hello Register")
}
