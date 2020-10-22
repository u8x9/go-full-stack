package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func login(c *gin.Context) {
	c.String(http.StatusOK, "user login")
}
func register(c *gin.Context) {
	c.String(http.StatusOK, "user register")
}
func info(c *gin.Context) {
	id := c.Param("id")
	c.String(http.StatusOK, "user info by %s", id)
}
func detail(c *gin.Context) {
	id := c.Param("id")
	c.String(http.StatusOK, "content detail by %s", id)
}
func addContent(c *gin.Context) {
	c.String(http.StatusOK, "add new content")
}
func editContent(c *gin.Context) {
	id := c.Param("id")
	c.String(http.StatusOK, "edit content %s", id)
}

func main() {
	r := gin.Default()
	// 路由组
	ug := r.Group("/user")
	{
		ug.POST("/login", login)
		ug.POST("/register", register)
		ug.GET("/info/:id", info)
	}
	cg := r.Group("/content")
	{
		cg.GET("/:id", detail)
		cg.POST("/", addContent)
		cg.PUT("/:id", editContent)
	}
	r.Run(":8000")
}
