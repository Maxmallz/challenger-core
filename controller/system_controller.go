package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SystemController struct {
	NewSystemController(*gin.Engine)
}

func NewSystemController(router *gin.Engine) {
    systemController := &SystemController{}
    router.GET("/ping", systemController.Ping)
}

func (system *SystemController) Ping(context *gin.Context) {
    context.JSON(http.StatusOK, gin.H{
        "message": "pong",
    })
}
