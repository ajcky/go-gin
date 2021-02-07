package main

import (
	"fmt"
	"go-gin/router"
)

func main() {
	r := router.SetupRouter()
	err := r.Run(":8090")
	if err != nil{
		fmt.Println(err)
	}
}