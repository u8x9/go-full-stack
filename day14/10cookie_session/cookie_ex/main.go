package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// cookie 练习
// 模拟实现权限验证中间件
// 有2个路由，login 和 home
// login 用于设置 cookie
// home 是访问查看信息的请求
// 在请求 home 之前，先跑中间件代码，检验是否存在cookie

const cookieName = "user"

func checkLogin(c *gin.Context) {
	cookie, err := c.Cookie(cookieName)
	if err != nil {
		c.String(http.StatusUnauthorized, "请先登录")
		c.Abort()
		return
	}
	c.Set("user", cookie)
	c.Next()
}

func login(c *gin.Context) {
	if c.Request.Method == http.MethodGet {
		c.HTML(http.StatusOK, "login.html", nil)
		return
	}
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "u8x9" && password == "9x8u" {
		c.SetCookie(cookieName, username, 360, "/", "127.0.0.1", false, true)
		c.Redirect(http.StatusMovedPermanently, "/home")
	}
}

func home(c *gin.Context) {
	data := gin.H{
		"user": c.GetString("user"),
	}
	c.HTML(http.StatusOK, "home.html", data)
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/login", login)
	r.POST("/login", login)
	r.GET("/home", checkLogin, home)
	r.Run(":8000")
}
