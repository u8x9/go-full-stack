package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/u8x9/go-full-stack/day14/11ex/db"
	"github.com/u8x9/go-full-stack/day14/11ex/handler"
)

func init() {
	if err := db.Init(); err != nil {
		fmt.Println("Init database failed:", err)
		return
	}
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", handler.Index)
	r.GET("/add", handler.Add)
	r.POST("/add", handler.Add)
	r.GET("/del/:id", handler.Del)
	r.Run(":8000")
}
