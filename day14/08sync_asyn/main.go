package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func longAsync(c *gin.Context) {
	// 异步操作 不能直接使用原始context
	// 需要使用其副本
	cc := c.Copy()
	// 异步处理
	go func() {
		time.Sleep(3 * time.Second)
		log.Println("异步执行：", cc.Request.URL.Path)
	}()
}
func longSync(c *gin.Context) {
	time.Sleep(3 * time.Second)
	log.Println("同步执行：", c.Request.URL.Path)

}
func main() {
	r := gin.Default()
	r.GET("/long_async", longAsync) // 异步
	r.GET("/long_sync", longSync)   // 同步
	r.Run(":8000")
}
