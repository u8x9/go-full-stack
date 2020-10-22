package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	User     string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"pass" uri:"pass" xml:"pass" binding:"required"`
}

// bind2json 绑定json
func bind2json(c *gin.Context) {
	var loginForm LoginForm
	//if err := c.BindJSON(&loginForm); err != nil {
	if err := c.ShouldBindJSON(&loginForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, loginForm)
}
func bind2form(c *gin.Context) {
	var loginForm LoginForm
	if err := c.Bind(&loginForm); err != nil {
		c.String(http.StatusBadRequest, "bind form failed: %v", err)
		return
	}
	c.String(http.StatusOK, "%#v", loginForm)
}
func bind2uri(c *gin.Context) {
	var loginForm LoginForm
	if err := c.BindUri(&loginForm); err != nil {
		c.String(http.StatusBadRequest, "bind uri failed: %v", err)
		return
	}
	c.String(http.StatusOK, "%#v", loginForm)
}

// bind2xml 绑定 xml
func bind2xml(c *gin.Context) {
	var loginForm LoginForm
	if err := c.BindXML(&loginForm); err != nil {
		c.String(http.StatusBadRequest, "bind xml failed: %v", err)
		return
	}
	c.String(http.StatusOK, "%#v", loginForm)
}
func main() {
	r := gin.Default()
	r.POST("/json", bind2json)
	r.POST("/form", bind2form)
    r.GET("/uri/:user/:pass", bind2uri)

	/*
	   <login>
	       <user>张三</user>
	       <pass>123</pass>
	   </login>
	*/
	r.POST("/xml", bind2xml)
	r.Run(":8000")
}
