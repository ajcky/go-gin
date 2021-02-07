package router

import (
	"github.com/gin-gonic/gin"
	"go-gin/middle_ware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	LoginRoute(r)
	//
	r.Use(middle_ware.CheckAuth())
	HelloRoute(r)
	return r
}

