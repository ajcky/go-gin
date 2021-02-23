package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin/config"
	"go-gin/model"
)

func Hello(c *gin.Context) {
	id := c.Query("id")
	girl := model.Girl{}
	err := config.DBS.Table("girl_star").Where("id = ?", id).Find(&girl).Error
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200, girl.Nickname)
}
