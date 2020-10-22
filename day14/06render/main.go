package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
)

func jsonRender(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
func xmlRender(c *gin.Context) {
	c.XML(http.StatusOK, gin.H{"name": "张三", "code": 0})
}
func yamlRender(c *gin.Context) {
	c.YAML(http.StatusOK, gin.H{"host": "127.0.0.1", "port": 12345, "paths": []string{"/foo", "/bar", "/foobar"}})
}

func protoBufRender(c *gin.Context) {
	reps := []int64{int64(1), int64(2)}
	label := "label"
	data := &protoexample.Test{
		Label: &label,
		Reps:  reps,
	}
	c.ProtoBuf(http.StatusOK, data)
}

func htmlRender(c *gin.Context) {
	data := gin.H{
		"title": "你好，世界🐯",
	}
	c.HTML(http.StatusOK, "index.html", data)
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/json", jsonRender)
	r.GET("/xml", xmlRender)
	r.GET("/yaml", yamlRender)
	r.GET("/protobuf", protoBufRender)
	r.GET("/html", htmlRender)
	r.Run(":8000")
}
