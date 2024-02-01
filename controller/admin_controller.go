package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminController struct {
	NewAdminController(*gin.Engine)
}

func NewAdminController(router *gin.Engine) {
    AdminController := &AdminController{}
    router.GET("/admin/tenant", AdminController.Tenant)
}

func (admin *AdminController) Tenant(context *gin.Context) {
    context.JSON(http.StatusOK, gin.H{
        "message": "TODO: Implement me!",
    })
}
