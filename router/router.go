package router

import (
	"Assigment-3/controllers"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	r.GET("/status", controllers.GetStatus)

	return r
}
