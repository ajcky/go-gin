package controller

import "github.com/gin-gonic/gin"

func Hello(c *gin.Context){
	uid,ok := c.Get("uid")
	if ok {
		c.JSON(200,uid)
	}
}