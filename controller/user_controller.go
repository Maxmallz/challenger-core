package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	NewUserController(*gin.Engine)
}

func NewUserController(router *gin.Engine) {
    userController := &UserController{}
    router.GET("/users", userController.GetAllUsers)
}

func (user *UserController) GetAllUsers(context *gin.Context) {
    context.JSON(http.StatusOK, gin.H{
        "message": "TODO: Implement me!",
    })
}
