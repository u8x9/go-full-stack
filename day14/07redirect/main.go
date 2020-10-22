package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func redirect(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "https://google.com")
}
func main() {
	r := gin.Default()
	r.GET("/google", redirect)
	r.Run(":8000")
}
