package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/cookie", func(c *gin.Context) {
		// 从客户端获取cookie
		// c.Cookie 是对 c.Request.Cookie 的封装
		username, err := c.Cookie("username")
		if err != nil { // 如果客户端没有cookie，则设置
			username = "u8x9"
			// Http Header -> Set-Cookie: username=u8x9; Path=/; Domain=127.0.0.1; Max-Age=60; HttpOnly
			c.SetCookie("username", username, 60, "/", "127.0.0.1", false, true)
		}
		c.String(200, "username from cookie is: %q", username)
	})
	r.Run(":8000")
}
