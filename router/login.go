package router

import (
	"github.com/gin-gonic/gin"
	"go-gin/middle_ware"
)

func LoginRoute(e *gin.Engine)  {
	e.POST("/user/login", middle_ware.Auth())
}
