package router

import (
	"github.com/gin-gonic/gin"
	"go-gin/controller"
)

func HelloRoute(e *gin.Engine) {
	e.GET("/hello", controller.Hello)
}
