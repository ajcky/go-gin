package main

import (
	"fmt"
	"go-gin/config"
	"go-gin/router"
)

func main() {
	err := config.InitConfig()
	if err != nil {
		fmt.Println(err)
	}
	err = config.InitMysql()
	if err != nil {
		fmt.Println(err)
	}
	r := router.SetupRouter()
	err = r.Run(":80")
	if err != nil {
		fmt.Println(err)
	}
}
