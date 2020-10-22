package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "你好%s", name)
	})
	// http://127.0.0.1:8000/info/zhangsan/foo/bar/foobar/barfoo
	r.GET("/info/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		c.String(http.StatusOK, "%q - %q", name, action)
	})
	r.Run(":8000")
}
