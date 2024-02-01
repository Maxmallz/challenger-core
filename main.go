package main

import (
	"challenger-core/controller"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
    controller.NewSystemController(router)
	controller.NewAdminController(router)
	controller.NewUserController(router)

	fmt.Printf("Server started at :8080\n")
    http.ListenAndServe(":8080", router)
}