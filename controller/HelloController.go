package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin/config"
)

type Girl struct {
	ID int64 `gorm:"column:id" json:"id" form:"id"`
	Nickname string `gorm:"column:nickname" json:"nickname" form:"nickname"`
}

func Hello(c *gin.Context){
	id := c.Query("id")
	girl := Girl{}
	err := config.DBS.Table("girl_star").Where("id = ?", id).Find(&girl).Error
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200,girl.Nickname)
}