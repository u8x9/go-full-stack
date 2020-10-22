package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func middleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		log.Println("中间件启动了～")
		// 设置变量到Context的key中
		c.Set("foo", "两只🐯")
		// 执行函数
		c.Next()
		status := c.Writer.Status()
		log.Println("中间件执行完毕：", status)
		//log.Println("中间件耗时：", time.Now().Sub(t))
		log.Println("中间件耗时：", time.Since(t))
	}
}
func middleWare2() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		log.Println("中间件2启动了～")
		// 设置变量到Context的key中
		c.Set("foo", "锄禾日当午")
		// 执行函数
		c.Next()
		status := c.Writer.Status()
		log.Println("中间件2执行完毕：", status)
		//log.Println("中间件耗时：", time.Now().Sub(t))
		log.Println("中间件2耗时：", time.Since(t))
	}
}

func main() {
	r := gin.Default()
	r.Use(middleWare())
	{
		r.GET("/mw", func(c *gin.Context) {
			v, ok := c.Get("foo")
			if !ok {
				log.Println("can not get value from context")
			}
			c.String(200, "Hello from middleWare: %v", v)
		})
		r.GET("/mw2", middleWare2(), func(c *gin.Context) {
			v, ok := c.Get("foo")
			if !ok {
				log.Println("can not get value from context")
			}
			c.String(200, "Hello from middleWare: %v", v)
		})
	}
	r.Run(":8000")
}
